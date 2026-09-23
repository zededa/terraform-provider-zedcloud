#!/bin/bash
# SSH into the most recently started node, without needing tofu outputs
# (which are unavailable while an apply is still running).
#
#   ./eve.sh 'eve status'
#   ./eve.sh            # interactive
set -u
BASE="$HOME/Library/Application Support/zedamigo/edge_nodes"
D=$(ls -dt "$BASE"/*/ 2>/dev/null | head -1)

# Ask the live gvproxy process, not the logs: 0.13.1 does not record the
# forwards in start_vm.bash, and the three ports it opens are, in order,
# ssh (22), 10022 and 10080 -- take the lowest.
#
# NOTE: `lsof -p <pid> -iTCP` ORs the two filters. The `-a` is required to AND
# them, otherwise you get every listening socket on the machine.
#
# Do NOT glob over this directory with grep: it contains disk0.raw (100 GB)
# and unix sockets that block on open.
GVPID=$(cat "$D/gvproxy.pid" 2>/dev/null || true)
PORT=$(lsof -nP -a -p "${GVPID:-0}" -iTCP -sTCP:LISTEN 2>/dev/null \
       | awk 'NR>1{sub(/.*:/,"",$9); print $9}' | sort -n | head -1)

if [ -z "${PORT:-}" ]; then
  echo "no gvproxy listener found for $D (pid ${GVPID:-none})" >&2
  exit 1
fi
echo "# node dir: $D" >&2
echo "# ssh port: $PORT" >&2

exec ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
     -o ConnectTimeout=5 -o LogLevel=ERROR \
     -i ~/.ssh/eve_probe_ed25519 -o IdentitiesOnly=yes \
     -p "$PORT" root@localhost "$@"
