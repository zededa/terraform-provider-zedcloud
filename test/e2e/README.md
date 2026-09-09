# test/e2e — virtual EVE node on the Berlin lab, onboarded to alpha

Creates a virtual EVE-OS edge node as a QEMU/KVM VM on the Berlin lab host
`h-m-dl20`, onboards it into the **alpha** Zedcloud cluster, and blocks until the
controller reports it `RUN_STATE_ONLINE`.

This is iteration 1 of
[`docs/design/e2e-virtual-eve-node-testing.md`](../../docs/design/e2e-virtual-eve-node-testing.md):
prove the node lifecycle. Running the acceptance suite against the node is iteration 2.

**Validated end-to-end on 2026-08-19.** ~3.5 minutes from `apply` to `ONLINE`.

## What it creates

| Provider | Resources |
|---|---|
| `zedcloud` | `brand`, `model`, `project`, `network`, `edgenode` |
| `zedamigo` | `host_reservation`, `disk_image`, `eve_installer`, `installed_edge_node`, `edge_node`, `wait_until` |

Every object name carries `var.run_id`, so concurrent runs and leftovers cannot
collide — necessary on alpha, which holds ~233 devices and ~334 projects belonging
to other people.

## Endpoints (verified)

| Purpose | Host |
|---|---|
| Control / REST API (`zedcloud` provider) | `zedcontrol.alpha.zededa.net` |
| Device API (EVE `/config/server`) | `zedcloud.alpha.zededa.net` |

`zedcontrold.alpha.zededa.net` **does not resolve** — do not use it. Both real hosts
sit behind Cloudflare with a Google Trust Services cert for `*.alpha.zededa.net`, so
no `tls_ca` and no `additional_hosts` are needed (unlike the local-cluster setup,
where EVE's pre-onboarding TLS check failed silently against a private Zededa CA).

## Prerequisites

The lab host is on `10.208.13.94`, reachable only via the Berlin office VPN.

1. Start the VPN container (from `zededa-berlin-lab/berlin-vpn-in-a-container`).
2. Reach the host through squid's CONNECT proxy. Note that macOS `nc -X connect`
   misparses squid's response — use `socat`:

   ```sh
   LAB() { ssh -o ProxyCommand='socat - PROXY:127.0.0.1:%h:%p,proxyport=3128' \
                -i ~/.ssh/ivan-berlin-lab ivan@10.208.13.94 "$@"; }
   ```

   (If the container publishes `11022`, `-o ProxyJump=root@localhost:11022` also works.)

3. Install the zedamigo provider on the host. It is not on a public registry, so
   cross-compile it and drop it in the filesystem mirror `e2e.tfrc` points at:

   ```sh
   cd ~/go/src/github.com/andrei-zededa/terraform-provider-zedamigo
   GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath \
     -ldflags "-s -w -X main.version=0.13.1" \
     -o /tmp/terraform-provider-zedamigo_v0.13.1 .

   M=.terraform.d/plugins/localhost/andrei-zededa/zedamigo/0.13.1/linux_amd64
   LAB "mkdir -p ~/$M"
   scp -o ProxyCommand='socat - PROXY:127.0.0.1:%h:%p,proxyport=3128' \
       -i ~/.ssh/ivan-berlin-lab /tmp/terraform-provider-zedamigo_v0.13.1 \
       ivan@10.208.13.94:"~/$M/"
   LAB "chmod +x ~/$M/*"
   ```

   The binary is statically linked, which is what makes it run on NixOS.

4. Point `zededa/zedcloud` at a locally built provider. `terraform.tf` declares
   no version constraint on purpose — this config is meant to exercise the
   provider **you are changing**, not whatever the registry currently serves —
   so your `e2e.tfrc` needs a `dev_overrides` entry:

   ```hcl
   provider_installation {
     dev_overrides {
       "zededa/zedcloud" = "/home/<you>/e2e-alpha-provider"
     }
     filesystem_mirror {
       path    = "/home/<you>/.terraform.d/plugins"
       include = ["localhost/andrei-zededa/zedamigo"]
     }
     direct {
       exclude = ["localhost/andrei-zededa/zedamigo"]
     }
   }
   ```

   `dev_overrides` must come first, and the directory is scanned for a plain
   `terraform-provider-zedcloud`. `make build` emits a version-suffixed name,
   so link it:

   ```sh
   make build
   ln -sf "$(ls -1t v2/terraform-provider-zedcloud_* | head -1)" \
          v2/terraform-provider-zedcloud
   ```

   Cross-compile for the lab host if you are building on macOS:
   `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -C ./v2 -o ../provider-linux .`

   Expect `tofu` to print "Provider development overrides are in effect" on
   every command. That warning is the confirmation it worked — without it, you
   are silently testing a released provider. CI generates the equivalent config
   and hard-fails if the binary is missing, rather than falling back.

## Running it

OpenTofu runs **on the lab host** in this iteration, so the zedamigo `target` is
localhost and no `ssh {}` block is needed. Moving to a GitHub-hosted runner later
means setting `target` + `ssh {}` in `terraform.tf` and changing nothing else —
every path in this config is already a path on the target.

```sh
rsync -az -e "ssh -o ProxyCommand='socat - PROXY:127.0.0.1:%h:%p,proxyport=3128' \
  -i $HOME/.ssh/ivan-berlin-lab" ./ ivan@10.208.13.94:e2e-alpha/

LAB "cd ~/e2e-alpha \
  && export TF_CLI_CONFIG_FILE=\$PWD/e2e.tfrc \
            TF_VAR_zedcloud_token='<api-token>' \
            TF_VAR_run_id='it1' \
            TF_VAR_edge_node_ssh_pub_key='$(cat ~/.ssh/ivan-berlin-lab.pub)' \
  && tofu init -input=false && tofu apply -auto-approve"
```

For long applies, detach on the host and poll — note the braces, without them `&`
backgrounds the whole `cd && export && ...` chain and the log lands in `$HOME`:

```sh
LAB "cd ~/e2e-alpha && export ... && { setsid nohup tofu apply -auto-approve \
  -no-color > apply.log 2>&1 < /dev/null & } ; sleep 8; tail -8 apply.log"
```

Teardown:

```sh
LAB "cd ~/e2e-alpha && export TF_CLI_CONFIG_FILE=\$PWD/e2e.tfrc TF_VAR_... \
  && tofu destroy -auto-approve"
```

Destroying matters twice over: it stops the QEMU process **and** releases the
capacity reservation, which is durable and survives reboots.

## Measured timings (h-m-dl20, 16 cores / 62 GB, EVE 16.0.1)

| Step | Duration |
|---|---|
| `eve_installer` (image cached) | 12s |
| `installed_edge_node` (install VM, synchronous) | 1m34s |
| `edge_node` (QEMU detached) | 0s |
| `wait_until` boot → register → `RUN_STATE_ONLINE` | 1m21s (5 attempts) |
| **total** | **~3m10s** |

Far cheaper than the 30–60 min the design doc assumed, which came from a macOS/arm64
emulated run. This makes a per-PR trigger considerably more viable and weakens the
argument for a warm node pool.

## Known issues

**Capacity reservation lock is single-user.** On `h-m-dl20`,
`/var/lib/zedamigo/reservations/.lock` is mode `0644` owned by `andrei`, while the
directory is group-writable (`drwxrwsr-x root:zedamigo_reservations`). The provider
opens the lock for writing, so every group member except `andrei` fails:

```
za-host-reservation: /var/lib/zedamigo/reservations/.lock: Permission denied
cannot open lock file (is /var/lib/zedamigo/reservations writable?)
```

Cooperative reservation therefore works for exactly one user. `var.reservations_path`
currently points at a per-user tree as a workaround, which keeps the resource
exercised but provides **no cross-user coordination**. The real fix is `0664`/`g+w`
on `.lock`, declared in the lab's `configuration.nix` so it survives a
`nixos-rebuild`. Switch `reservations_path` back to the default once that lands.

**Installer-time network warnings are benign.** The install console reports
`eth0 with dhcp is NOT working properly` and `eth0 with static configuration is not
working`, then completes successfully and the node onboards normally. The installer's
connectivity probe appears to run before SLIRP is usable.

## Running the real-node acceptance test against it

Once the node is up, `TestApplicationInstance_RealNode` deploys an nginx container onto
it and waits until the device reports `SW_STATE_RUNNING`. That test runs the provider
in-process, so it runs from anywhere with access to the alpha API — it does **not** need
to run on the lab host:

```sh
cd v2/resources
TF_ACC=1 \
TF_CLI_CONFIG_FILE=$PWD/../../dev.tfrc \
TF_VAR_zedcloud_url=zedcontrol.alpha.zededa.net \
TF_VAR_zedcloud_token='<api-token>' \
ZEDCLOUD_ACC_REAL_NODE=1 \
ZEDCLOUD_TEST_NODE_NAME=tf_e2e_en_it1 \
ZEDCLOUD_TEST_SUFFIX=_it7 \
  go test -v -timeout 40m -run TestApplicationInstance_RealNode .
```

Takes about a minute. Without `ZEDCLOUD_ACC_REAL_NODE` the test skips, so the default
suite is unaffected. Use a fresh `ZEDCLOUD_TEST_SUFFIX` per run.

## Notes

- The onboarding key `5d0767ee-0547-4569-b530-387e526f8cb9` **is** authorized on
  alpha — the 403 `register` → "device not found" blocker documented against the
  local cluster did not occur here.
- `grub_cfg` must match `serial_type`. Default `serial_type` is `virtio`, whose
  console is `hvc0`. A mismatch produces an empty console log, so
  `installed_edge_node.success` never flips true — and it fails silently.
- `installed_edge_node` has **no timeout**: if the install hangs, the apply hangs.
- The two `swtpm` / `genisoimage` "executable not found" warnings are expected; those
  resources are not used.
- `eve_ssh_port` is forwarded on the **lab host**, not locally. Reach EVE with
  `LAB "ssh -p <port> root@localhost"` or your own `ssh -L` through the proxy.
