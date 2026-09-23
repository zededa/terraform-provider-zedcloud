#!/bin/bash
# Cluster readiness as the controller sees it. Works while an apply is running.
#
#   ./clusterstatus.sh [run_id]
#
# --cacert, not -k: the local control API uses a private certificate and a bare
# curl silently returns nothing.
set -u
RUN="${1:-mac}"
H="${ZC_HOST:-zedcontrol.local.zededa.net}"
T="${TF_VAR_zedcloud_token:?set TF_VAR_zedcloud_token}"
CA="${ZC_CA:-$HOME/zedcloud-local/certs/zededa-root-ca.pem}"

ID=$(curl -s --cacert "$CA" --max-time 20 -H "Authorization: Bearer $T" \
       "https://$H/api/v1/cluster" \
     | python3 -c "
import sys, json
run = '$RUN'
d = json.load(sys.stdin)
for c in d.get('list', []):
    if run in (c.get('name') or ''):
        print(c['id']); break
" 2>/dev/null)

if [ -z "${ID:-}" ]; then
  echo "no cluster matching '$RUN'" >&2
  exit 1
fi
echo "cluster $ID"

curl -s --cacert "$CA" --max-time 20 -H "Authorization: Bearer $T" \
  "https://$H/api/v1/cluster/id/$ID/status" \
  | python3 -c "
import sys, json
d = json.load(sys.stdin)
print('runState:', d.get('runState'))
for n in d.get('nodes', []):
    print(f\"  {n.get('name')}  {n.get('readyState')}\")
print(json.dumps(d)[:400] if not d.get('nodes') else '')
"
