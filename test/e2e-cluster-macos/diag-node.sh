#!/bin/bash
# Post-onboarding diagnostics over SSH. Use once the node is REGISTERED (sshd
# is unreachable before that -- use console-exec.sh instead).
set -u
cd "$(dirname "$0")"
export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"
PORT="${1:-$(tofu output -raw eve_ssh_port)}"

e() {
  ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
      -o ConnectTimeout=5 -o BatchMode=yes -o LogLevel=ERROR \
      -i ~/.ssh/eve_probe_ed25519 -o IdentitiesOnly=yes \
      -p "$PORT" root@localhost "$@"
}

echo "=== controller reachability from the node ==="
e 'eve exec pillar curl -s -o /dev/null -w "ping=%{http_code} ip=%{remote_ip}\n" --max-time 10 https://$(cat /config/server)/api/v2/edgedevice/ping' 2>&1 | head -3

echo
echo "=== /etc/hosts inside pillar (where additional_hosts matters) ==="
e 'eve exec pillar cat /etc/hosts' 2>&1 | head -10

echo
echo "=== zedagent / config fetch errors, last 25 ==="
e 'eve exec pillar sh -c "grep -ahiE \"error|fail|refused|x509|certificate|timeout\" /persist/newlog/collect/*.gz 2>/dev/null | tail -5"' 2>&1 | tail -6
e 'cat /persist/log/*.log 2>/dev/null | grep -aiE "zedagent|getconfig|x509|certificate|refused" | tail -20' 2>&1 | tail -20

echo
echo "=== eve status ==="
e 'eve status 2>&1 | head -25' 2>&1 | head -25

echo
echo "=== kube markers ==="
e 'ls -l /var/lib/*initialized 2>&1; eve exec kube ls /var/lib/ 2>&1 | head' 2>&1 | head -10
