#!/bin/bash
# Follow the node by NAME rather than by Terraform output, so it works while
# an apply is still in flight (outputs are not written until it finishes).
#
#   ./watch.sh mac4
set -u
RUN="${1:?usage: watch.sh <run_id>}"
H="${ZC_HOST:-zedcontrol.local.zededa.net}"
T="${TF_VAR_zedcloud_token_local:-$TF_VAR_zedcloud_token}"

curl -sk --max-time 20 -H "Authorization: Bearer $T" \
  "https://$H/api/v1/devices/status-config" \
  | python3 -c "
import sys, json
run = '$RUN'
d = json.load(sys.stdin)
rows = [x for x in d.get('list', []) if run in (x.get('name') or '')]
if not rows:
    print('no device matching', run)
for x in rows:
    print(x.get('name'))
    print('  id        :', x.get('id'))
    print('  adminState:', x.get('adminState'))
    print('  runState  :', x.get('runState'))
    print('  clusterId :', (x.get('edgeNodeCluster') or {}).get('id'))
" 2>&1 | head -20
