#!/bin/bash
# Ask the controller what it thinks of the node, and try SSH. No secrets printed.
set -u
cd "$(dirname "$0")"
export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"
DEV=$(tofu output -raw device_id 2>/dev/null)
PORT=$(tofu output -raw eve_ssh_port 2>/dev/null)
H="${ZC_HOST:-zedcloud.local.zededa.net}"

echo "device=$DEV  ssh_port=$PORT"

echo
echo "=== controller view ==="
curl -sk --max-time 15 -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
  "https://$H/api/v1/devices/id/$DEV/status" \
  | python3 -c "
import sys,json
d=json.load(sys.stdin)
for k in ('runState','adminState','lastRebootTime','eveVersion','serialno'):
    if k in d: print(f'  {k}: {d[k]}')
oc=d.get('optionalCapabilities')
print('  optionalCapabilities:', oc)
" 2>&1 | head -12

echo
echo "=== is EVE talking at all (any device object activity)? ==="
curl -sk --max-time 15 -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
  "https://$H/api/v1/devices/id/$DEV" \
  | python3 -c "
import sys,json
d=json.load(sys.stdin)
print('  adminState:', d.get('adminState'))
print('  serialno  :', d.get('serialno'))
print('  onboarding:', (d.get('onboardingKey') or '')[:8]+'...')
" 2>&1 | head -8

echo
echo "=== ssh :$PORT ==="
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
    -o ConnectTimeout=5 -o BatchMode=yes -o LogLevel=ERROR \
    -i ~/.ssh/eve_probe_ed25519 -o IdentitiesOnly=yes \
    -p "$PORT" root@localhost 'echo SSH_OK; cat /etc/hosts | grep -i zededa' 2>&1 | head -5
