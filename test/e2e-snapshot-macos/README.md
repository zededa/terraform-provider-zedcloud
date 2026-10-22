# e2e-snapshot-macos — PLE-576 snapshots on a virtual EVE node, against a local zedcloud

Exercises `zedcloud_app_instance_snapshot`, `zedcloud_app_instance_snapshot_activation`,
`data.zedcloud_app_instance_snapshots` and the `allow_snapshots` / `max_snapshots` /
`snapshot_on_update` attributes of `zedcloud_application_instance` on one virtual EVE node
running on this Mac (vfkit + zedamigo), onboarded to the local minikube controller.

Derived from [`test/e2e-cluster-macos`](../e2e-cluster-macos) — read its README first. All of
its traps apply here (leaked `TF_VAR_*`, the trailing newline in `additional_hosts`, no
`eve_nuke_disks`, the private CA). Differences:

- EVE is the **KVM** flavour (`17.0.0-lts-kvm-arm64`). The controller refuses snapshots on
  KubeVirt devices, which the `-k-` flavour is.
- No kube readiness and no cluster: snapshots are not supported for edge-node-cluster app
  instances (PLE-576 D3).
- `app.tf` deploys a small VM app (CirrOS arm64 qcow2 from download.cirros-cloud.net) and
  drives the snapshot resources in stages.

## Setup

Same prerequisites as `e2e-cluster-macos` (vfkit, qemu, Docker, zedamigo 0.13.1, the probe
SSH key, `~/zedcloud-local/certs/zededa-root-ca.pem`). Then:

```shell
cp e2e-snapshot-macos.tfrc.example e2e-snapshot-macos.tfrc      # edit the mirror path
cd ../../v2 && CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o /tmp/p . && \
  cp /tmp/p ~/.terraform.d/plugins/registry.opentofu.org/zededa/zedcloud/99.0.0/darwin_arm64/terraform-provider-zedcloud_v99.0.0
cd ../test/e2e-snapshot-macos
TF_CLI_CONFIG_FILE=$PWD/e2e-snapshot-macos.tfrc tofu init
```

After every provider rebuild, drop the `zededa/zedcloud` block from `.terraform.lock.hcl`
and re-run `tofu init -upgrade`: the lock records the binary's checksum.

## Stages

`ZC_RUN_ID` must be fresh for every node build.

| Command | What it checks |
|---|---|
| `./run.sh apply -var stage=0` | node onboarded (~4 min incl. installer build) |
| `./run.sh apply -var stage=1` then `plan` | VM app running; `allow_snapshots = true`; `max_snapshots` reads back 1 with no diff |
| `./run.sh apply -var stage=2` | immediate snapshot `CREATED` (~1 min; EVE restarts the app) |
| `./run.sh apply -var stage=3 -var 'app_description=after'` | update leaves the snapshot alone |
| `./run.sh apply -var stage=4 -var 'app_description=after'` then `plan` | rollback `ACTIVE`; app instance shows `description` drift back to the snapshot's |
| `... -var rollback_trigger=2` | second rollback to the same snapshot is **refused** |
| `-var stage=3 -var keep_snapshot_resource=false -var snapshot_on_update=true` | flag-only change sends no update |
| `... -var 'app_description=v2' -var purge_counter=1` | update sent with `?makeSnapshot=true`; app-update snapshot recorded with the pre-update config |
| `./run.sh destroy -var stage=4 -var keep_snapshot_resource=false -var rollback_snapshot_name=x` | teardown |

`./eve.sh '<cmd>'` runs a command on the node. EVE's own view of snapshots is in
`/run/zedagent/AppInstanceConfig/*.json` (`Snapshot`, `RollbackCmd`) and
`/run/zedmanager/AppInstanceStatus/*.json` (`SnapStatus`).

## Results, 2026-10-21 (M5 Pro, EVE 17.0.0-lts-kvm-arm64, local zedcloud)

Verified end to end:

- the node reports `apiCapability = API_CAPABILITY_REPORT_TPM_EVENTLOG` (21). EVE reports the
  **highest** level it implements, so the capability check compares levels (>= 13). zcli's
  exact-name test would refuse this node.
- immediate snapshot: create → `CREATED` in ~1 min, read, import, delete.
- rollback: `ACTIVE` in ~30 s; the controller config returns to the snapshot's; the next plan
  shows the drift (D4). `apply -refresh-only` records it in state, but the plan still reverts
  it until the configuration matches.
- `allow_snapshots`, `max_snapshots` (1 → 2) round trip with clean plans; `snapshot_on_update`
  flag-only change sends nothing; an update sends `?makeSnapshot=true`.

Found on the device, and handled in the provider:

1. **A snapshot can be rolled back to only once.** purus always sends EVE
   `RollbackCmd{Counter: 1}` (`srvs/purus/appinst.go`), and EVE rolls back only when the
   counter changes. A second activate makes seine rewrite the app config to the snapshot's while
   the volumes stay, and ganges sits in `ROLLING_BACK` forever. The activation resource now
   refuses `ACTIVE` / `ROLLING_BACK` / `ROLLBACK_FAILED` snapshots before sending anything.
2. **Only `CREATED` snapshots can be rolled back to.** Activating a `CREATING` app-update
   snapshot made EVE reject the app config (`Error restoring config from snapshot ... not
   available`), so the app stopped receiving config until the snapshot was deleted. Refused now.
3. **App-update snapshots stay `CREATING`** until an update that replaces the app's volumes:
   description and cpu changes did not take them; a purge did (EVE moved both requested snapshots
   to `AvailableSnapshots` at the purge's config version).

Controller / EVE behaviour to raise in zedcloud (not provider bugs):

- the RollbackCmd counter (item 1).
- the snapshot's stored config looked stale: snapshots taken at revision 3 (`description = v2`)
  carry `description = v1`. Needs a check of what seine copies into the snapshot.
- ganges answered 404 for the state of an immediate snapshot taken while app-update snapshots were
  pending on the device, and never recovered (`processDeviceSnapshots` removes stored snapshots the
  device does not report yet). The provider times out with "the controller never reported a state".
- EVE kept a request for a snapshot that the controller had deleted in
  `SnapStatus.RequestedSnapshots`; with `max_snapshots = 1` that blocked every later app-update
  snapshot ("Too many snapshots requested").
- the `purge` counter reads back 0, so a configured `purge { counter = 1 }` plans an update on
  every run (pre-existing; with `snapshot_on_update` each of those takes a snapshot).
