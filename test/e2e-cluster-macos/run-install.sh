#!/bin/bash
# Build the EVE-OS installer and run the install, WITHOUT touching zedcloud.
#
# zedamigo_installed_edge_node.INSTALL depends only on zedamigo_eve_installer
# and zedamigo_disk_image, so this needs no valid controller token -- it can
# run while a fresh one is being fetched. It also produces the soft_serial that
# the zedcloud_edgenode record needs, and a bootable node to test controller
# reachability from.
set -eu
cd "$(dirname "$0")"
export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"
export TF_VAR_run_id="${TF_VAR_run_id:-mac1}"
export TF_VAR_zedcloud_token="${TF_VAR_zedcloud_token:-placeholder}"
exec tofu apply -no-color -auto-approve -target=zedamigo_installed_edge_node.INSTALL
