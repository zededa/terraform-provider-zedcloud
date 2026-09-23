#!/bin/bash
# Wrapper that runs tofu in this directory with a CLEAN variable environment.
#
#   ./run.sh plan
#   ./run.sh apply -auto-approve
#   ./run.sh destroy -auto-approve
#
# WHY THIS EXISTS -- a trap that cost an hour on 2026-09-21.
#
# TF_VAR_* is global to the shell, not scoped to a directory. A ~/.zshrc set up
# for test/e2e-cluster against the Berlin lab exports, among others:
#
#   TF_VAR_eve_cluster_host=zedcloud.alpha.zededa.net
#   TF_VAR_zedamigo_target=10.208.13.94
#   TF_VAR_node_disk_devs=[]
#   TF_VAR_run_id=<the lab run>
#
# Those override this config's defaults with no warning whatsoever. The first
# build here came out with /config/server = zedcloud.alpha.zededa.net and
# /config/hosts = "192.168.127.254 zedcloud.alpha.zededa.net" -- so the node
# dialled the LOCAL controller's address while presenting alpha's hostname in
# SNI, the local ingress had no route or certificate for it, and the node sat
# in RUN_STATE_PROVISIONED with nothing anywhere saying why. Terraform reports
# no diagnostic for an env var that merely supplies a variable.
#
# The installer bakes these values, so getting it wrong costs a full rebuild.
#
# terraform.tf also carries a precondition that fails the plan when
# eve_cluster_host and zedcloud_url are not in the same domain, which catches
# this in seconds rather than in an hour. Belt and braces.
set -eu
cd "$(dirname "$0")"

# Drop every TF_VAR_ this config does not want, whatever the shell says.
for v in $(env | sed -n 's/^\(TF_VAR_[A-Za-z0-9_]*\)=.*/\1/p'); do
  case "$v" in
    TF_VAR_zedcloud_token) : ;;   # the only one worth inheriting
    *) unset "$v" ;;
  esac
done

export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"

: "${ZC_RUN_ID:?set ZC_RUN_ID to a FRESH id for every build, e.g. ZC_RUN_ID=mac2}"
export TF_VAR_run_id="$ZC_RUN_ID"

# The token must be the LOCAL one. TF_VAR_zedcloud_token_local is what ~/.zshrc
# calls it; prefer that when it is set, so a stale alpha token in
# TF_VAR_zedcloud_token cannot be used by accident.
if [ -n "${TF_VAR_zedcloud_token_local:-}" ]; then
  export TF_VAR_zedcloud_token="$TF_VAR_zedcloud_token_local"
fi

# Fail fast rather than after a 35-minute onboarding timeout.
CODE=$(curl -sk -o /dev/null -w '%{http_code}' --max-time 10 \
  -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
  "https://${TF_VAR_zedcloud_url:-zedcontrol.local.zededa.net}/api/v1/enterprises")
if [ "$CODE" != "200" ]; then
  echo "token check failed: HTTP $CODE from the local controller." >&2
  echo "  401 with \"Session Cache miss\" means the token predates the last" >&2
  echo "  redeploy -- re-login to the local UI and re-export." >&2
  exit 1
fi

exec tofu "$@"
