#!/usr/bin/env bash
# Deletes every app instance, then every edge app, whose name starts with
# PREFIX (default: ue177, which covers ue177-* and ue177f-* runs), straight through the API. For leftovers of runs
# whose state no longer matches the controller.
set -uo pipefail
PREFIX=${1:-ue177}
API="https://${TF_VAR_zedcloud_url#https://}/api"
H=(-H "Authorization: Bearer $TF_VAR_zedcloud_token")

del_all() { # path kind
  local ids n=0
  ids=$(curl -sk "${H[@]}" "$API/$1?next.pageSize=1000" \
        | jq -r --arg p "$PREFIX" '.list[]? | select(.name | startswith($p)) | .id')
  for id in $ids; do
    curl -sk -o /dev/null "${H[@]}" -X DELETE "$API/$1/id/$id" & n=$((n+1))
    (( n % 20 == 0 )) && wait
  done
  wait
  echo "$2: deleted $n"
}
del_all v1/apps/instances instances
sleep 5
del_all v1/apps apps
