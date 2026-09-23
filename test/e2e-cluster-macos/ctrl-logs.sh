#!/bin/bash
# Pull attestation/EdgeNodeCert errors out of the controller's own logs, for
# when the node reports RUN_STATE_MAINTENANCE_MODE and you need to know why.
#
# The deployment names to inspect are specific to your controller install and
# are deliberately not hardcoded here -- pass them:
#
#   ZC_SERVICES="<svc> <svc>" ./ctrl-logs.sh
#
# Find them with:  kubectl -n "$ZC_NS" get deploy -o name
set -u
NS="${ZC_NS:-zedcloud}"
: "${ZC_SERVICES:?set ZC_SERVICES to the controller deployment names to inspect, space-separated}"

# Pods are named <release>-<service>-<hash>; the app label is not just the
# service name, so select by name prefix instead.
for app in $ZC_SERVICES; do
  POD=$(kubectl -n "$NS" get pods -o name 2>/dev/null | grep -m1 -- "-$app-")
  echo "=== $app  ($POD) ==="
  [ -z "$POD" ] && continue
  kubectl -n "$NS" logs "$POD" --tail=6000 --since=60m 2>/dev/null \
    | grep -iE 'edgenodecert|attest|does not chain|invalid state for device|signature' \
    | tail -12
done
