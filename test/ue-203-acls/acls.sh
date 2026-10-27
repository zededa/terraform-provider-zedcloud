#!/bin/bash
# Print the stored interface ACLs of edge app instances, by name.
# Usage: ./acls.sh NAME... (uses TF_VAR_zedcloud_url / TF_VAR_zedcloud_token)
here=$(cd "$(dirname "$0")" && pwd)
for n in "$@"; do
  curl -sk -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
    "https://$TF_VAR_zedcloud_url/api/v1/apps/instances/name/$n" | python3 "$here/acls.py" "$n"
done
