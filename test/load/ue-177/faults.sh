#!/usr/bin/env bash
# Fault-injection run through mitmproxy (faults.py), reproducing the ENG-3015 /
# UE-177 failure modes on a fast local controller:
#   ./faults.sh new|old [N]
# Needs mitmdump listening on 127.0.0.1:8666 (see run-mitm.sh).
set -euo pipefail
VARIANT=$1; N=${2:-40}
PREFIX="ue177f-$VARIANT-$(date +%H%M%S)"
source "$(dirname "$0")/lib.sh"          # API= points at the real controller
FAULT_FILE=${UE177_FAULT_FILE:-/tmp/ue177-fault}
export TF_VAR_zedcloud_url=127.0.0.1:8666 TF_INSECURE_SKIP_TLS_VERIFY=true

fault() { echo "$1" > "$FAULT_FILE"; echo "-- fault: $1"; }
v() { echo -var "n_apps=$1" -var "n_inst=$2"; }

# delete_oob K: delete K of this run's instances behind Terraform's back
delete_oob() {
  local ids
  ids=$(curl -sk -H "Authorization: Bearer $TF_VAR_zedcloud_token" "$API/v1/apps/instances?next.pageSize=1000" \
        | jq -r --arg p "$PREFIX" '[.list[]? | select(.name | startswith($p))][0:'"$1"'][] | .id')
  for id in $ids; do
    curl -sk -o /dev/null -H "Authorization: Bearer $TF_VAR_zedcloud_token" -X DELETE "$API/v1/apps/instances/id/$id"
  done
  echo "               deleted $(echo "$ids" | grep -c .) instances out of band"
}

echo "== $VARIANT faults  prefix=$PREFIX  apps=$N instances=$N parallelism=$PAR =="
fault none
step  A-setup         apply -auto-approve $(v "$N" 0)
fault slow_writes
step  B-inst-create   apply -auto-approve $(v "$N" "$N")
controller_counts     B
fault none
step  C-reapply       apply -auto-approve $(v "$N" "$N")
controller_counts     C
fault html504
plan_changes D-plan-504 $(v "$N" "$N")
fault none
delete_oob 10
plan_changes E-plan-oob $(v "$N" "$N")
step  E-converge      apply -auto-approve $(v "$N" "$N")
controller_counts     E
fault slow_writes
step  F-inst-delete   apply -auto-approve $(v "$N" 0)
fault none
controller_counts     F
plan_changes G-plan-after $(v "$N" 0)
if [ "$VARIANT" = new ]; then
  fault slow_jobs
  step  I-create-slowjob apply -auto-approve $(v "$N" 20) -var inst_timeout=10s
  fault none
  controller_counts     I
  step  I-converge      apply -auto-approve $(v "$N" 20)
  controller_counts     I2
  step  I-delete        apply -auto-approve $(v "$N" 0)
fi
fault slow_writes
step  H-apps-delete   apply -auto-approve $(v 0 0)
fault none
controller_counts     H
plan_changes H-plan-after $(v 0 0)
step  Z-teardown      destroy -auto-approve $(v 0 0)
controller_counts     Z
echo "PREFIX=$PREFIX"
