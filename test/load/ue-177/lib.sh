# shellcheck shell=bash
# Helpers shared by the scenario scripts. Source it; set VARIANT (new|old)
# and PREFIX first.
LOAD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PAR="${PAR:-160}"
RESULTS="$LOAD/logs/results.tsv"
API="https://${TF_VAR_zedcloud_url#https://}/api"
[ -f "$RESULTS" ] || printf 'time\tvariant\tprefix\tstep\texit\tseconds\terrors\ttop_error\n' > "$RESULTS"

tf() {
  (cd "$LOAD/$VARIANT" && TF_CLI_CONFIG_FILE="$LOAD/$VARIANT.tfrc" terraform "$@")
}

# record STEP EXIT SECONDS LOG
record() {
  local step=$1 rc=$2 secs=$3 log=$4 errs top
  # -no-color prints diagnostics without the box drawing: "Error: <summary>"
  errs=$(grep -Ec '^(│ )?Error: ' "$log" || true)
  top=$(grep -E '^(│ )?Error: ' "$log" | sed 's/^│ *//' \
        | sed -E 's/[0-9a-f]{8}-[0-9a-f-]{27}/<id>/g; s/inst-[0-9]+/inst-N/g; s/app-[0-9]+/app-N/g' \
        | sort | uniq -c | sort -rn | head -1 | sed 's/^ *//' | cut -c1-150 || true)
  printf '%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n' "$(date +%T)" "$VARIANT" "$PREFIX" "$step" "$rc" "$secs" "$errs" "$top" >> "$RESULTS"
  printf '%-14s exit=%s %4ss errors=%-4s %s\n' "$step" "$rc" "$secs" "$errs" "$top"
}

# step NAME terraform-args...  (apply/destroy/plan with -parallelism)
step() {
  local name=$1; shift
  local log="$LOAD/logs/$VARIANT-$PREFIX-$name.log" t0 rc
  t0=$(date +%s)
  set +e
  tf "$@" -no-color -input=false -parallelism="$PAR" -var "prefix=$PREFIX" > "$log" 2>&1
  rc=$?
  set -e
  record "$name" "$rc" "$(( $(date +%s) - t0 ))" "$log"
  return 0
}

# plan_changes NAME vars... -> records how many resources a plan wants to change
plan_changes() {
  local name=$1; shift
  local log="$LOAD/logs/$VARIANT-$PREFIX-$name.log" t0 rc summary
  t0=$(date +%s)
  set +e
  tf plan -detailed-exitcode -no-color -input=false -parallelism="$PAR" -var "prefix=$PREFIX" "$@" > "$log" 2>&1
  rc=$?
  set -e
  record "$name" "$rc" "$(( $(date +%s) - t0 ))" "$log"
  summary=$(grep -E '^Plan:|No changes' "$log" | head -1 || true)
  echo "               ${summary:-<no plan summary, see $log>}"
  printf '\t\t\t%s-summary\t\t\t\t%s\n' "$name" "$summary" >> "$RESULTS"
}

# api_count KIND -> number of objects on the controller whose name starts with PREFIX
api_count() {
  local path
  case $1 in
    inst) path="v1/apps/instances" ;;
    app) path="v1/apps" ;;
  esac
  curl -sk -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
    "$API/$path?next.pageSize=1000" \
    | jq -r --arg p "$PREFIX" '[.list[]? | select(.name | startswith($p))] | length'
}

controller_counts() {
  local label=$1 a i
  a=$(api_count app || true); i=$(api_count inst || true)
  echo "               controller ($label): apps=$a instances=$i"
  printf '\t%s\t%s\tcontroller-%s\t\t\tapps=%s instances=%s\t\n' "$VARIANT" "$PREFIX" "$label" "$a" "$i" >> "$RESULTS"
}
