#!/bin/bash
# Grep the running node's serial console log for the handful of things that
# actually matter early on. Remember: EVE never logs kube activity here, so
# absence of k3s/kubelet/longhorn in this file is not evidence of anything.
set -u
LOG="${1:-}"
if [ -z "$LOG" ]; then
  cd "$(dirname "$0")"
  export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"
  LOG=$(tofu output -raw run_console_log)
fi
echo "log: $LOG  ($(wc -l < "$LOG") lines)"

echo
echo "=== ssh / sshd ==="
grep -aiE 'sshd|authorized_keys|debug\.enable\.ssh' "$LOG" | tail -10

echo
echo "=== network / dhcp / addressing ==="
grep -aiE 'dhcp|eth0.*(up|addr)|192\.168\.127|10\.0\.2\.' "$LOG" | tail -12

echo
echo "=== controller contact ==="
grep -aiE 'zedcloud\.local|register|onboard|uuid|/api/v2' "$LOG" | tail -20

echo
echo "=== TLS / certificate problems ==="
grep -aiE 'x509|certificate|tls|verify failed|unknown authority' "$LOG" | tail -15

echo
echo "=== errors ==="
grep -aiE 'error|fatal|panic|refused|timeout|no route' "$LOG" | tail -20
