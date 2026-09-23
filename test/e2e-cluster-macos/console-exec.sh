#!/bin/bash
# Run a command on the node through the SERIAL CONSOLE, not SSH.
#
#   ./console-exec.sh 'ip -4 addr show'
#
# Why this exists: on macOS vfkit exposes the serial port as a PTY and EVE
# auto-logs-in root on hvc0, so the console is a usable shell even when sshd
# is not listening -- which is exactly the situation where you need it most.
# zedamigo's tailer already reads that PTY into serial_console_run.log, so we
# write the command in and read the answer out of the log.
#
# Not interactive, and output ordering is best-effort. For anything more than
# a few commands, fix SSH instead.
set -u
CMD="${1:?usage: console-exec.sh '<shell command>'}"
cd "$(dirname "$0")"
export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"

LOG=$(tofu output -raw run_console_log)
DIR=$(dirname "$LOG")
PTY=$(grep -h 'pty path' "$DIR"/*vfkit_stderr.log | tail -1 | sed -E 's/.*pty path: ([^)]*)\).*/\1/')

if [ -z "$PTY" ] || [ ! -e "$PTY" ]; then
  echo "no usable PTY found (looked in $DIR/*vfkit_stderr.log)" >&2
  exit 1
fi

MARK="__CEX_$$__"
BEFORE=$(wc -l < "$LOG")

printf '\n%s\necho %s_END\n' "$CMD" "$MARK" > "$PTY"
sleep 4
tail -n "+$((BEFORE + 1))" "$LOG"
