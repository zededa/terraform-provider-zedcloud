#!/usr/bin/env bash
# Builds the branch provider and renders the two Terraform roots:
#   new/  -> this branch (dev_overrides), with timeouts {} and provider knobs
#   old/  -> registry zededa/zedcloud 2.8.0, the version in ENG-3015 / UE-177
set -euo pipefail
LOAD="$(cd "$(dirname "$0")" && pwd)"
REPO="${REPO:-$(cd "$LOAD/../../.." && pwd)}"
mkdir -p "$LOAD/bin" "$LOAD/new" "$LOAD/old" "$LOAD/logs"

echo "building provider from $(git -C "$REPO" rev-parse --abbrev-ref HEAD)@$(git -C "$REPO" rev-parse --short HEAD)"
go build -C "$REPO/v2" -o "$LOAD/bin/terraform-provider-zedcloud"

python3 - "$LOAD" <<'PY'
import sys, pathlib
load = pathlib.Path(sys.argv[1])
tmpl = (load / "module.tf.tmpl").read_text()
timeouts = """timeouts {
    create = var.inst_timeout
    update = var.inst_timeout
    delete = var.inst_timeout
  }"""
new_extra = '''
variable "request_timeout" {
  type    = string
  default = "60s"
}
variable "async_writes" {
  type    = bool
  default = true
}
provider "zedcloud" {
  request_timeout = var.request_timeout
  async_writes    = var.async_writes
}
'''
(load / "new" / "main.tf").write_text(tmpl.replace("@@TIMEOUTS@@", timeouts) + new_extra)
old = tmpl.replace("@@TIMEOUTS@@", "").replace(
    'source = "zededa/zedcloud"', 'source  = "zededa/zedcloud"\n      version = "2.8.0"')
(load / "old" / "main.tf").write_text(old + '\nprovider "zedcloud" {}\n')
PY

cat > "$LOAD/new.tfrc" <<EOF
provider_installation {
  dev_overrides {
    "zededa/zedcloud" = "$LOAD/bin"
  }
  direct {}
}
EOF
printf 'provider_installation {\n  direct {}\n}\n' > "$LOAD/old.tfrc"

(cd "$LOAD/old" && TF_CLI_CONFIG_FILE="$LOAD/old.tfrc" terraform init -input=false -no-color >/dev/null && echo "old: initialised (2.8.0)")
echo "setup done"
