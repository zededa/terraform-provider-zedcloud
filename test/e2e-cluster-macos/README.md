# e2e-cluster-macos — a 1-node edge-node cluster on this Mac, against a local zedcloud

A vfkit/Apple-Silicon sibling of [`test/e2e-cluster`](../e2e-cluster). Same
goal — form a `zedcloud_edgenode_cluster` over a real virtual EVE node so that
CI-834 and CI-709 are reproducible — but with no Berlin lab and no VPN in the
loop.

**Status: written, not yet applied.** Two things are unverified and either
could sink it; see [Unknowns](#unknowns).

---

## Why a separate directory

`test/e2e-cluster` is proven on `h-m-dl20` and is the config a PR is being
raised for. macOS differs from it structurally, not by a flag:

| | Linux (`test/e2e-cluster`) | macOS (here) |
|---|---|---|
| Hypervisor | QEMU/KVM | vfkit (Virtualization.framework) |
| Host reservations | `zedamigo_host_reservation` | not supported — plain vars |
| Node networking | SLIRP + taps on bridges | gvproxy only |
| Max nodes | 1 or 3+ | **1** — no bridge/tap/dhcp resources exist |
| Installer format | `iso` → `installer_iso` | `raw` → `installer_raw` |
| Serial | chosen in Terraform, stamped via SMBIOS | **EVE's soft serial, read back** |
| vTPM | available | none — Virtualization.framework has no TPM device |
| CPU pinning | yes | no (`taskset` is Linux-only) |
| Controller | alpha | local minikube deployment |

Folding all of that behind a boolean would make the Linux config unreadable
and would put an unverified path into the PR that carries the verified one.

---

## The inverted serial flow

This is the one thing most likely to confuse someone reading both configs.

On Linux the serial is chosen in `zedcloud_edgenode` and flows **downwards**:
QEMU stamps it with `-smbios type=1,serial=…` and EVE reads it back through
`dmidecode` as the hardware serial. The cloud object is the source of truth.

Apple's Virtualization.framework cannot set an SMBIOS serial. EVE therefore
generates a **soft serial** — a UUID derived from the device certificate — at
install time. So the dependency edge reverses:

```
Linux:   zedcloud_edgenode  ->  installed_edge_node  ->  edge_node
macOS:   installed_edge_node  ->  zedcloud_edgenode  ->  edge_node
```

`zedamigo_installed_edge_node.INSTALL.serial_no` is a throwaway
(`"1234567890"`, as in upstream's macOS example); the value that matters is
`soft_serial`, and it feeds `zedcloud_edgenode.EN.serialno`.

---

## Pointing EVE at a private, loopback-only controller

Three separate problems, three separate attributes.

**1. The name resolves to 127.0.0.1.** `zedcloud.local.zededa.net` is an
`/etc/hosts` entry on the Mac, and `minikube tunnel` binds the istio ingress
LoadBalancer on loopback. A guest can use neither. So the installer bakes an
`/etc/hosts` entry of its own:

```hcl
additional_hosts = "192.168.127.254 zedcloud.local.zededa.net"
```

`192.168.127.254` is gvproxy's host address on its default `192.168.127.0/24`
network. Traffic to it is dialled by the gvproxy process on the Mac, so it
lands on the Mac's loopback and reaches the tunnel — **no rebinding of
`minikube tunnel` required**, if the theory holds.

**2. The TLS certificate is signed by a private CA.** Verified 2026-09-21:

```
0 s:CN=zedcloud.local.zededa.net
  i:CN=Zededa Inc. Intermediat CA1
1 s:CN=Zededa Inc. Intermediat CA1
  i:CN=Zededa Inc. Root CA
Verify return code: 0 (ok)   [against ~/zedcloud-local/certs/zededa-root-ca.pem]
```

so `tls_ca` gets that root. This is exactly what killed the earlier
local-controller attempt recorded in
`docs/design/ue-156-3-node-cluster-e2e.md`: EVE's pre-onboarding TLS check
failed against the private CA and said essentially nothing about why.

**3. Config is signed.** `object_signing_ca` defaults to the same root. That
is the expected arrangement but is **not confirmed** — see Unknowns.

---

## Prerequisites

```shell
brew install vfkit qemu            # qemu-img is required on macOS too
# Docker Desktop running (the installer is built by running lfedge/eve)

curl -fsSL https://raw.githubusercontent.com/andrei-zededa/terraform-provider-zedamigo/v0.13.1/install.sh \
  | bash -s -- --binary-only 0.13.1

ssh-keygen -t ed25519 -N "" -f ~/.ssh/eve_probe_ed25519
```

Nested virtualisation on macOS needs **Apple M3 or later** (EVE starts VMs of
its own inside the VM).

Stage the PR's provider binary into the mirror, and point tofu at the CLI
config:

```shell
cd ../../v2 && CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 \
  go build -o terraform-provider-zedcloud .
M=~/.terraform.d/plugins/registry.opentofu.org/zededa/zedcloud/99.0.0/darwin_arm64
mkdir -p "$M" && cp -f terraform-provider-zedcloud "$M/terraform-provider-zedcloud_v99.0.0"

cd ../test/e2e-cluster-macos
cp e2e-cluster-macos.tfrc.example e2e-cluster-macos.tfrc   # edit the path
export TF_CLI_CONFIG_FILE="$PWD/e2e-cluster-macos.tfrc"
```

Re-copy the binary after every rebuild.

## The token

`TF_VAR_zedcloud_token` must be minted against the **local** controller. An
alpha token fails with a recognisable error:

```
HTTP 401   error[0].details = "Session Cache miss"
```

The local controller keeps sessions in memory, so redeploying it — which
ArgoCD does routinely — invalidates every token issued before. Re-login to the
local UI and re-export. Check before spending an hour on a build:

```shell
curl -sk -o /dev/null -w '%{http_code}\n' \
  -H "Authorization: Bearer $TF_VAR_zedcloud_token" \
  https://zedcloud.local.zededa.net/api/v1/enterprises
```

## Run

**Always go through `./run.sh`.** It scrubs `TF_VAR_*` before calling tofu, and
that is not optional — see the first trap below.

```shell
tofu init                                   # once
ZC_RUN_ID=mac1 ./run.sh apply -auto-approve
```

`ZC_RUN_ID` must be **fresh for every build**. Reinstalling EVE against an
already-registered device object leaves the node stuck forever on
`Waiting for DeviceName from controller...`.

Rough timings on an M5 (18 cores, 48 GB, Docker Desktop running the local
zedcloud alongside):

| step | time |
|---|---|
| build the installer (`docker run lfedge/eve`) | ~2 min (first time +1.1 GB pull) |
| install onto the 100 GB image | ~1.5 min |
| boot → `ADMIN_STATE_REGISTERED` | ~3 min |
| Kubernetes stack up | tens of minutes (`kube_ready_timeout` = 90m) |

There is a useful checkpoint partway: `./run-vm.sh` stops after booting the
node, before the waits, so you can check `./status.sh` and `./probe-node.sh`
without committing to the long barriers.

---

## CI-752 — the cluster-scoped instance repro

`ci752_repro.tf` adds a volume instance and a network instance that supply only
`edge_node_cluster`, so the controller picks their node. It is opt-in and it
does **not** wait on `cluster_ready`, which blocker B below makes unreachable
here — `deviceId` is assigned on the config plane, so the cluster object
existing is the whole prerequisite. Build in two stages:

```shell
ZC_RUN_ID=<fresh> ./run.sh apply -auto-approve -var repro_ci752=true \
  -target=zedcloud_edgenode_cluster.CLUSTER          # node, kube, cluster object
ZC_RUN_ID=<same>  ./run.sh apply -auto-approve -var repro_ci752=true \
  -target=zedcloud_volume_instance.CI752 \
  -target=zedcloud_network_instance.CI752
ZC_RUN_ID=<same>  ./run.sh plan -var repro_ci752=true \
  -target=zedcloud_volume_instance.CI752 \
  -target=zedcloud_network_instance.CI752            # must say "No changes"
```

Measured 2026-09-23 against the local controller, EVE `17.0.0-lts-k-arm64`:
onboard 1m0s, kube_ready 4m25s, cluster object 20s. Both instances came back
with `deviceId` populated and the plan was clean; reverting the #234 schema and
re-planning against the same objects reproduced the ticket's
`- device_id = "…" -> null` on both. Two description edits then exercised the
update path — revision 1 → 2 → 3, no 409, `deviceId` preserved.

The header of `ci752_repro.tf` carries the analysis and the raw-API check that
keeps the assertion from going vacuous.

---

## Helper scripts

| script | what it does |
|---|---|
| `run.sh` | tofu with a scrubbed `TF_VAR_*` environment and an up-front token check |
| `run-install.sh` | installer + install only; **needs no valid token** |
| `run-vm.sh` | up to a booted node, stopping before the waits |
| `status.sh` | what the controller thinks of the node, plus an SSH attempt |
| `probe-node.sh` | asks the node over SSH whether it can see the controller |
| `console-exec.sh` | runs a command over the **serial console** — works before onboarding, when SSH does not |
| `diag-node.sh` | post-onboarding diagnostics over SSH |
| `console-grep.sh` | the handful of interesting greps over the serial console log |

`console-exec.sh` is the important one. vfkit exposes the serial port as a PTY
and EVE auto-logs-in root on `hvc0`, so writing to that PTY and reading
zedamigo's console log gives a usable shell on a node that is not yet
reachable any other way:

```shell
./console-exec.sh 'cat /config/server; ip -4 addr show; netstat -tlnp'
```

---

## Traps found the hard way

Five of these cost real time on 2026-09-21, and four of them fail **silently**
— the apply reports success and the node simply never arrives.

**1. `TF_VAR_*` leaks in from your shell, with no diagnostic.**
A `~/.zshrc` set up for `test/e2e-cluster` exports
`TF_VAR_eve_cluster_host=zedcloud.alpha.zededa.net`,
`TF_VAR_zedamigo_target=10.208.13.94`, `TF_VAR_node_disk_devs=[]` and a
`TF_VAR_run_id` for the lab. Those silently override this config's defaults —
Terraform says nothing about an env var that merely supplies a variable. The
first build here baked `/config/server = zedcloud.alpha.zededa.net` and an
`/etc/hosts` entry pointing *alpha's name* at the *local* controller, so the
node dialled the right address with the wrong SNI, the local ingress had no
route or certificate for it, and the node sat in `RUN_STATE_PROVISIONED`
saying nothing. Hence `run.sh`, plus a precondition in `nodes.tf` that fails
the plan when `eve_cluster_host` and `zedcloud_url` are in different domains.

**2. `eve_nuke_disks` destroys the installer on macOS.**
The Linux harness passes `eve_nuke_disks=vda,sda,vdb,sdb`, safe there because
the installer is an ISO. Here vfkit attaches the target as `vda` and **the
installer raw as `vdb`** — so that setting wipes the installer's own CONFIG
partition mid-install. The install still reports success, `install_success` is
`true`, and the installer `.raw` provably contains the custom config, but the
installed disk contains none of it and the node boots with EVE's stock
`/config` (`server = zedcloud.zededa.net`, no keys, no CAs). The only hint
anywhere is `EVE install server : ` with an empty value in the install console
log. Do not add it back; a fresh disk image is created every run anyway.

**3. `additional_hosts` needs a trailing newline.**
zedamigo writes the string to `/config/hosts` verbatim and EVE **appends** its
own entry once it has a device UUID. Without a final newline you get

```
192.168.127.254 zedcloud.local.zededa.net127.0.0.1 9aa846b3-...
```

and the name becomes `zedcloud.local.zededa.net127.0.0.1`. The node registers
fine — that happens before EVE appends — and then goes silent forever:
`ADMIN_STATE_REGISTERED`, `runState` stuck at `RUN_STATE_PROVISIONED`, and
`curl` from inside pillar returning `000`. The config uses a heredoc so the
newline cannot be edited away.

**4. sshd is unreachable before onboarding.** sshd *is* listening on `0.0.0.0:22`
inside the guest from early boot, but gvproxy's forward gets `connection
refused` until the node has onboarded and EVE opens it up. So
`kex_exchange_identification: read: Connection reset by peer` on a fresh node
is expected, not a fault. Use `console-exec.sh` until the node is registered.

**5. `/config` is mounted read-only** on a running node, so you cannot patch
`/config/hosts` or `/config/server` in place to avoid a rebuild. Anything baked
by the installer costs a full rebuild to change.

Inherited from the Linux harness, unchanged here:

- a 30 GB disk gives `/persist` ~5.1 GB, below Longhorn's 25%-free floor → the
  node goes into DiskPressure, multus is evicted, kubevirt cycles forever, and
  EVE never advertises `hvTypeKubevirt`. Hence the 100 GB default.
- `RUN_STATE_SUSPECT` is not an error — it means info received, metrics
  pending — and the cluster API gates on `AdminState`, never on `runState`.
- the probe's SSH key must be the one baked into the **installer**;
  `debug.enable.ssh` does not apply on first boot.
- `zedamigo_eve_installer` rejects updates rather than being `ForceNew`, and
  `-replace` on it alone silently does not reinstall the node — replace the
  whole chain. A fresh `ZC_RUN_ID` is the reliable way.
- **EVE never logs kube activity to the serial console.** k3s can run for half
  an hour leaving no trace there, so absence of `k3s`/`kubelet`/`longhorn` in
  that log is not evidence of anything.
- sweep stray VMs **scoped to `$HOME`**, or you kill other users' VMs on a
  shared host: `pkill -f "$HOME/.local/state/zedamigo/edge_nodes/"`.

---

## What is verified

Measured on 2026-09-21, M5 Mac, EVE `17.0.0-lts-k-arm64`, local zedcloud
`21.0.1-17-g94098c4fc`:

- EVE-K **arm64 installs and boots under vfkit**, with nested virtualisation
  (`--nested`) and no vTPM.
- The guest gets `192.168.127.2/24` from gvproxy, default via
  `192.168.127.1`, and **reaches the Mac's loopback services at
  `192.168.127.254`** — so `minikube tunnel` needs no rebinding.
- A **private CA works**: `tls_ca` and `object_signing_ca` from
  `~/zedcloud-local/certs/zededa-root-ca.pem` are enough for EVE to trust the
  local ingress.
- The **stock lf-edge onboarding key** (`5d0767ee-…`) is accepted by the local
  controller; the node reached `ADMIN_STATE_REGISTERED` about three minutes
  after boot.
- The **soft-serial flow works**: `zedamigo_installed_edge_node.soft_serial`
  feeds `zedcloud_edgenode.serialno`, and the controller matched the device.

- **k3s comes up on arm64 under vfkit.** `kubectl get nodes` →
  `tf-encm-en-mac4 Ready control-plane v1.34.2+k3s1`, about 20 minutes after
  boot. `/persist` is 74 GB with 66 GB free off the 100 GB image.
- **The cluster forms on the controller.** `zedcloud_edgenode_cluster` is
  created and the controller writes the CI-834 surface back onto the device:

  ```
  clusterInterface: eth0
  edgeNodeCluster:  clusterPrefix 10.244.244.1/28, seedNodeIp 10.244.244.1,
                    isMaster true, seedNodeId = this device
  ```

## Two blockers found, both real

**A. the controller's EdgeNodeCert verification rejects a TPM-less node (zedcloud, hard blocker).**

the controller's attestation handler (the controller's EdgeNodeCert verification, added as a hardening change) requires every EdgeNodeCert in an attest request to be signed by
the device cert. It is unconditional — no feature flag. On this node the controller
logged, every few seconds:

```
EdgeNodeCert verification for device <uuid>:
  cert 0 (type CERT_TYPE_DEVICE_ECDH_EXCHANGE):
  signature does not chain to device cert: x509: ECDSA verification failure
```

(and the same for `CERT_TYPE_DEVICE_RESTRICTED_SIGNING`). EVE responds by
setting `EdgeNodeCertsRefused` and dropping into `RUN_STATE_MAINTENANCE_MODE`,
then clearing it, then setting it again — flapping indefinitely. While it
flaps, EVE's own config fetch reports `configStatus: fail`, so nothing
progresses.

The cause is an ordering problem, visible in the file timestamps on the node:

```
/persist/certs/ecdh.cert.pem     notBefore 11:53:14   issuer  O=The Linux Foundation, CN=EVE
/persist/certs/attest.cert.pem   notBefore 11:53:14   issuer  O=The Linux Foundation, CN=EVE
/config/device.cert.pem          notBefore 11:55:07   subject O=The Linux Foundation, CN=EVE
```

EVE mints the EdgeNodeCerts at first boot and a *new* device cert two minutes
later during onboarding. The issuer *name* matches; the *key* does not — hence
"ECDSA verification failure" rather than "unknown authority". With a TPM the
device key is TPM-resident and stable across that transition, which is why the
Berlin lab (swtpm) never hits this and Apple's Virtualization.framework — which
provides no vTPM at all — always does.

Workaround that recovers a running node, verified here:

```shell
./eve.sh 'rm -f /persist/certs/{ecdh,attest}.{cert,key}.pem; reboot'
```

EVE regenerates them against the current device cert, attestation passes,
maintenance mode clears for good, and `configStatus` becomes `success`.

Worth a zedcloud ticket: either EVE should regenerate its EdgeNodeCerts when
the device cert changes, or the check needs to tolerate this flow. Note
`tests/k6/zedcloud/lib/cert.go` deliberately signs its child certs with the
device key, so the test suite encodes the assumption that EVE always does.

**B. The cluster config never reaches EVE (open).**

With the node healthy, `configStatus: success`, and the controller holding a
fully-populated `edgeNodeCluster` for it, EVE still never receives the cluster
section:

```
/persist/status/zedagent/          contains only PatchEnvelopeInfoList
                                   -- no EdgeNodeClusterConfig
grep -c EdgeNodeClusterConfig <node log>   -> 0
kubectl get nodes -o wide          INTERNAL-IP 192.168.127.2, not 10.244.244.1
device status: optionalCapabilities null, clusterNodeStatus empty
cluster status: RUN_STATE_UNPROVISIONED, nodes: []
```

So `zedamigo_wait_until.cluster_ready` cannot pass. Recreating the cluster with
`-replace=zedcloud_edgenode_cluster.CLUSTER` after the node was healthy did not
change it. The cluster object exists and is correct in the REST view; the push
to the device is what is missing. The `edgeNodeCluster` block carries
`kubernetesMode: false` and `clusterType: 0`, which may or may not be related.

This is the next thing to chase, and it is a controller-side question rather
than anything this config can fix.
