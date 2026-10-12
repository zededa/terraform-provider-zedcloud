#!/usr/bin/env bash
# A/B load run, the CI-1004 / ENG-3015 sequence at -parallelism=$PAR:
#   ./ab.sh new|old [N_APPS] [N_INST] [extra -var args for new...]
set -euo pipefail
VARIANT=$1; N_APPS=${2:-200}; N_INST=${3:-120}; shift $(( $# < 3 ? $# : 3 ))
PREFIX="ue177-$VARIANT-$(date +%H%M%S)"
source "$(dirname "$0")/lib.sh"
EXTRA=("$@")
v() { echo -var "n_apps=$1" -var "n_inst=$2" "${EXTRA[@]+"${EXTRA[@]}"}"; }

echo "== $VARIANT  prefix=$PREFIX  apps=$N_APPS instances=$N_INST parallelism=$PAR =="
step setup        apply -auto-approve $(v 0 0)
step apps-create  apply -auto-approve $(v "$N_APPS" 0)
step inst-create  apply -auto-approve $(v "$N_APPS" "$N_INST")
controller_counts after-create
step reapply      apply -auto-approve $(v "$N_APPS" "$N_INST")
plan_changes drift-plan $(v "$N_APPS" "$N_INST")
step inst-delete  apply -auto-approve $(v "$N_APPS" 0)
step apps-delete  apply -auto-approve $(v 0 0)
controller_counts after-delete
plan_changes post-delete-plan $(v 0 0)
step teardown     destroy -auto-approve $(v 0 0)
controller_counts after-teardown
echo "PREFIX=$PREFIX"
