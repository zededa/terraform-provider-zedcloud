#!/bin/bash
# tofu in this directory with a CLEAN TF_VAR_* environment (see
# test/e2e-cluster-macos/run.sh for why that is not optional).
#
#   ZC_RUN_ID=acl1 ./run.sh apply -auto-approve -var stage=1
set -eu
cd "$(dirname "$0")"

for v in $(env | sed -n 's/^\(TF_VAR_[A-Za-z0-9_]*\)=.*/\1/p'); do
  case "$v" in
    TF_VAR_zedcloud_token|TF_VAR_zedcloud_token_local|TF_VAR_vm_image) : ;;
    *) unset "$v" ;;
  esac
done

export TF_CLI_CONFIG_FILE="$PWD/e2e-acl-macos.tfrc"

: "${ZC_RUN_ID:?set ZC_RUN_ID to a FRESH id for every node build, e.g. ZC_RUN_ID=acl1}"
export TF_VAR_run_id="$ZC_RUN_ID"

if [ -n "${TF_VAR_zedcloud_token_local:-}" ]; then
  export TF_VAR_zedcloud_token="$TF_VAR_zedcloud_token_local"
fi

CODE=$(curl -sk -o /dev/null -w '%{http_code}' --max-time 10 \
  -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
  "https://zedcontrol.local.zededa.net/api/v1/enterprises")
if [ "$CODE" != "200" ]; then
  echo "token check failed: HTTP $CODE from the local controller." >&2
  exit 1
fi

exec tofu "$@"
