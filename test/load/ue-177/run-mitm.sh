#!/usr/bin/env bash
# Starts mitmdump as a reverse proxy in front of the controller on
# 127.0.0.1:8666 with the UE-177 fault-injection addon (faults.py).
# Switch faults with: echo slow_writes|html504|slow_jobs|none > /tmp/ue177-fault
# Stop it with: pkill -f "mitmdump --mode reverse"
set -euo pipefail
LOAD="$(cd "$(dirname "$0")" && pwd)"
mkdir -p "$LOAD/logs"
echo none > "${UE177_FAULT_FILE:-/tmp/ue177-fault}"
nohup mitmdump --mode "reverse:https://${TF_VAR_zedcloud_url#https://}" --listen-port 8666 \
  --set ssl_insecure=true -s "$LOAD/faults.py" > "$LOAD/logs/mitm.log" 2>&1 &
sleep 3
curl -sk -o /dev/null -w "mitmdump on 127.0.0.1:8666 -> HTTP %{http_code}\n" \
  -H "Authorization: Bearer $TF_VAR_zedcloud_token" "https://127.0.0.1:8666/api/v1/apps?next.pageSize=1"
