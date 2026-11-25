#!/bin/bash
# Show, for every app instance in state: what the controller stores, the
# controller's status (run state, errors), and what EVE received.
#
#   ZC_RUN_ID=acl1 ./check.sh
set -u
cd "$(dirname "$0")"
IDS=$(./run.sh output -json app_instance_ids 2>/dev/null)
API="https://zedcontrol.local.zededa.net/api"
TOKEN=${TF_VAR_zedcloud_token_local:-$TF_VAR_zedcloud_token}
get() { curl -sk --max-time 15 -H "Authorization: Bearer $TOKEN" "$API$1"; }

for k in $(printf '%s' "$IDS" | jq -r 'keys[]?'); do
  id=$(printf '%s' "$IDS" | jq -r --arg k "$k" '.[$k]')
  echo "== instance $k ($id)"
  get "/v1/apps/instances/id/$id" | jq -r '.interfaces[] | "stored \(.intfname): " +
    ([.acls[]? | if ([.actions[]? | select(.portmap)] | length) > 0
       then "\((.matches[] | select(.type=="lport") | .value))->\(.actions[0].mapparams.port)"
       else "allow \((.matches[] | select(.type=="ip") | .value))" end] | join(", "))'
  get "/v1/apps/instances/id/$id/status" | jq -c '{runState, errInfo: [.errInfo[]?.description], netStatus: [.netStatusList[]? | {ifName, ipAddrs: .ipAddrs}]}' 2>/dev/null
  echo "-- EVE AppInstanceConfig ACLs"
  ./eve.sh "cat /run/zedagent/AppInstanceConfig/$id.json" 2>/dev/null |
    jq -c '[.AppNetAdapterList[]? | {Name, ACLs: [.ACLs[]? | {m: [.Matches[] | "\(.Type)=\(.Value)"] | join(" "), a: [.Actions[]? | select(.PortMap) | .TargetPort]}]}]' 2>/dev/null ||
    echo "(no config on the device yet)"
done
echo "-- EVE port maps (iptables nat, all instances)"
./eve.sh "iptables -t nat -S 2>/dev/null | grep -iE 'dport|DNAT' | head -40; nft list ruleset 2>/dev/null | grep -iE 'dnat' | head -40" 2>/dev/null
echo "-- EVE errors"
./eve.sh "grep -h 'overlapping portmaps' /persist/newlog/keepSentQueue/* /persist/log/* 2>/dev/null | tail -3; cat /run/zedrouter/AppNetworkStatus/*.json 2>/dev/null | jq -c '{DisplayName, Error, ErrorTime}'" 2>/dev/null
