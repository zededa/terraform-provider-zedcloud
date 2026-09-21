#!/bin/bash
# Release ORPHANED zedamigo reservations. Run ON the zedamigo target.
#
# WHY THIS IS NEEDED
#
# zedamigo_host_reservation releases its claim on destroy. An apply that is
# killed, or a state file that is discarded, leaves the claim behind -- and the
# reservation tree is HOST-GLOBAL, not per-state, so the next run fails
# admission control:
#
#   Error: insufficient free CPUs under
#   /var/lib/zedamigo/reservations/cpus/unit: requested 6, only 2 free
#
# with no VM running and nothing in any state file. Measured on h-m-dl20
# 2026-09-21: 14 of 16 CPUs and 48 GB of RAM were held by three dead
# reservation ids from earlier runs. The error names neither the holder nor the
# path to the claim, so it reads like the host is busy.
#
# The claim files are plain text and a claim is released by truncating one.
# Each records: <reservation-id> <user> <user> <host> <host> <directory>.
#
# SAFETY
#
# The lab is shared. This only touches claims whose contents name BOTH the
# invoking user AND that user's harness directory, and it refuses to run while
# any of that user's VM processes are alive. It never matches on reservation id
# alone and never touches another user's claim -- the same discipline that the
# $HOME scoping in the pkill patterns exists for (see "Sweeping orphans" in
# README.md).
set -eu

ME="$(id -un)"
MYDIR="${ZEDAMIGO_HARNESS_DIR:-$HOME/e2e-cluster}"
ROOT="${ZEDAMIGO_RESERVATIONS:-/var/lib/zedamigo/reservations}"
DRY="${DRY_RUN:-0}"

echo "user=$ME  dir=$MYDIR  root=$ROOT  dry_run=$DRY"

# zedamigo_host_reservation is Linux-only (it needs util-linux flock and the
# shared tree), so on macOS there is nothing to release and no tree to read.
# Bail out rather than printing a summary computed from an unmatched glob --
# "free cpus: 1 / 1" is not a reassuring thing to leave lying around.
if [ ! -d "$ROOT/cpus/unit" ]; then
  echo "no reservation tree at $ROOT -- run this ON the Linux zedamigo target." >&2
  echo "(reservations do not exist on macOS; nothing to release here.)" >&2
  exit 0
fi

# `pgrep -c` and `pgrep -a` are GNU extensions: BSD/macOS pgrep has neither and
# exits with a usage error, which `|| true` would then swallow into a count of
# zero -- a guard that always passes. `| wc -l` is portable.
live=$(pgrep -f "$HOME/.local/state/zedamigo/edge_nodes/" 2>/dev/null | wc -l | tr -d ' ')
if [ "${live:-0}" != "0" ]; then
  echo "REFUSING: $live of your VM processes are still running; tofu destroy first" >&2
  ps -o pid=,command= -p "$(pgrep -f "$HOME/.local/state/zedamigo/edge_nodes/" | tr '\n' ',' | sed 's/,$//')" 2>/dev/null | head -5
  exit 1
fi

released=0
skipped=0
for f in $(find "$ROOT" -type f ! -name '.lock' 2>/dev/null); do
  [ -s "$f" ] || continue
  body=$(cat "$f" 2>/dev/null || true)
  case "$body" in
    *"$ME"*"$MYDIR"*)
      id=$(printf '%s' "$body" | cut -f1)
      if [ "$DRY" = "1" ]; then
        echo "would release  $f  ($id)"
      elif : > "$f" 2>/dev/null; then
        echo "released  $f  ($id)"
      else
        echo "NOT WRITABLE  $f  -- are you in the zedamigo_reservations group?" >&2
      fi
      released=$((released + 1))
      ;;
    *)
      echo "skipped (not yours)  $f  -> $(printf '%s' "$body" | cut -f2)"
      skipped=$((skipped + 1))
      ;;
  esac
done

echo
echo "released=$released skipped=$skipped"

free_cpus=0
total_cpus=0
for f in "$ROOT"/cpus/unit/*; do
  total_cpus=$((total_cpus + 1))
  [ -s "$f" ] || free_cpus=$((free_cpus + 1))
done
echo "free cpus: $free_cpus / $total_cpus"
echo "ram units still claimed: $(find "$ROOT/ram" -type f ! -name '.lock' -size +0 2>/dev/null | wc -l)"
