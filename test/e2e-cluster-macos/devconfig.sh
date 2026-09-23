#!/bin/bash
# The device CONFIG object (not /status): this is where cluster_interface and
# edgeNodeCluster live -- the CI-834 surface.
set -u
ID="${1:?usage: devconfig.sh <device-id>}"
H="${ZC_HOST:-zedcontrol.local.zededa.net}"
T="${TF_VAR_zedcloud_token:?set TF_VAR_zedcloud_token}"
CA="${ZC_CA:-$HOME/zedcloud-local/certs/zededa-root-ca.pem}"

curl -s --cacert "$CA" --max-time 20 -H "Authorization: Bearer $T" \
  "https://$H/api/v1/devices/id/$ID" > /tmp/devconfig.json

python3 - <<'PY'
import json
d = json.load(open('/tmp/devconfig.json'))
print('name            ', d.get('name'))
print('adminState      ', d.get('adminState'))
print('clusterInterface', d.get('clusterInterface'))
enc = d.get('edgeNodeCluster')
if enc:
    enc = dict(enc)
    if 'token' in enc:
        enc['token'] = '<redacted>'
print('edgeNodeCluster ', json.dumps(enc))
PY
