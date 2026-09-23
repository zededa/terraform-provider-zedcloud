#!/bin/bash
# Ask the node directly whether it can see the controller.
#
#   ./probe-node.sh [ssh_port]
#
# Answers the one question that decides whether this whole config can work:
# is var.controller_host_ip an address the guest can actually reach? Run it
# after run-vm.sh and before the full apply -- the onboard barrier has a
# 35-minute budget and this takes seconds.
set -u
PORT="${1:-$(cd "$(dirname "$0")" && TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc" tofu output -raw eve_ssh_port 2>/dev/null)}"
HOST_NAME="${2:-zedcloud.local.zededa.net}"

ssh_eve() {
  ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
      -o ConnectTimeout=5 -o BatchMode=yes -o LogLevel=ERROR \
      -i ~/.ssh/eve_probe_ed25519 -o IdentitiesOnly=yes \
      -p "$PORT" root@localhost "$@"
}

echo "=== port $PORT: reachable? ==="
ssh_eve 'uname -srm; cat /run/eve-release 2>/dev/null' 2>&1 | head -4

echo
echo "=== /etc/hosts (additional_hosts should appear here) ==="
ssh_eve 'cat /etc/hosts' 2>&1 | head -12

echo
echo "=== /config/server ==="
ssh_eve 'cat /config/server' 2>&1 | head -2

echo
echo "=== CAs baked into /config ==="
ssh_eve 'ls -l /config/*.pem 2>&1' 2>&1 | head -6

echo
echo "=== guest's own addressing ==="
ssh_eve 'ip -4 addr show; ip route' 2>&1 | grep -E 'inet |default' | head -8

echo
echo "=== can the guest reach the controller? ==="
ssh_eve "curl -s -o /dev/null -w 'ping http=%{http_code} via %{remote_ip}\n' --max-time 10 https://$HOST_NAME/api/v2/edgedevice/ping" 2>&1 | head -3

echo
echo "=== does TLS verify against the baked-in CA? ==="
ssh_eve "curl -s -o /dev/null -w 'tls http=%{http_code}\n' --cacert /config/v2tlsbaseroot-certificates.pem --max-time 10 https://$HOST_NAME/api/v2/edgedevice/ping" 2>&1 | head -3

echo
echo "=== has EVE tried to register? ==="
ssh_eve 'grep -aiE "register|onboard|ping.*fail|x509|certificate" /persist/newlog/keepSentQueue/* /persist/log/* 2>/dev/null | tail -15' 2>&1 | tail -15
