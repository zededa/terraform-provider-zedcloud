#!/bin/bash
# Has the node actually joined the cluster, as the DEVICE object reports it?
# ./clusternode.sh <device-id>
set -u
bash "$(dirname "$0")/devstatus.sh" "${1:?usage: clusternode.sh <device-id>}" >/dev/null 2>&1
python3 - <<'PY'
import json
d = json.load(open('/tmp/devstatus.json'))
print('runState          ', d.get('runState'))
print('adminState        ', d.get('adminState'))
print('clusterID         ', d.get('clusterID'))
print('lastUpdate        ', d.get('lastUpdate'))
print('optionalCapabilities', json.dumps(d.get('optionalCapabilities'))[:300])
print('clusterNodeStatus ', json.dumps(d.get('clusterNodeStatus'))[:500])
PY
