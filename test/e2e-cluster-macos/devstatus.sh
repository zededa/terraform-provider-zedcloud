#!/bin/bash
# Dump the interesting fields of a device's status. ./devstatus.sh <device-id>
set -u
ID="${1:?usage: devstatus.sh <device-id>}"
H="${ZC_HOST:-zedcontrol.local.zededa.net}"
# TF_VAR_zedcloud_token is the one to trust: a ~/.zshrc may also define
# TF_VAR_zedcloud_token_local, and it was observed STALE while the plain one
# was current -- which shows up as HTTP 401 "Session Cache miss" from these
# helpers while the apply itself works fine.
T="${TF_VAR_zedcloud_token:?set TF_VAR_zedcloud_token}"

curl -sk --max-time 20 -H "Authorization: Bearer $T" \
  "https://$H/api/v1/devices/id/$ID/status" > /tmp/devstatus.json

python3 - <<'PY'
import json
d = json.load(open('/tmp/devstatus.json'))
for k in ('runState', 'adminState', 'eveVersion', 'maintenanceMode',
          'maintenanceModeReason', 'maintenanceModeReasons', 'bootReason',
          'lastRebootReason', 'lastRebootTime', 'dormantTime'):
    if k in d:
        print(f'{k} = {d[k]}')
print('optionalCapabilities =', d.get('optionalCapabilities'))

# RUN_STATE_MAINTENANCE_MODE means EVE refused to bring the vault up. The
# reason lives in dataSecInfo / attestState, not in runState.
for k in ('dataSecInfo', 'attestState', 'devError', 'deviceRebootReason'):
    if d.get(k):
        print(f'{k} = {json.dumps(d[k])[:600]}')
print()
print('all keys:', sorted(d.keys()))
PY
