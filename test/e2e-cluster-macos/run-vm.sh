#!/bin/bash
# Create the cloud objects and boot the node, but STOP before the waits.
#
# Deliberate checkpoint: zedamigo_wait_until.onboarded has a 35-minute budget,
# and the single most likely way for this config to fail is the guest not
# reaching the controller at var.controller_host_ip. Far cheaper to boot the
# node, SSH in and curl the device API than to discover it after 35 minutes of
# polling.
set -eu
cd "$(dirname "$0")"
export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"
export TF_VAR_run_id="${TF_VAR_run_id:-mac1}"
exec tofu apply -no-color -auto-approve -target=zedamigo_edge_node.VM
