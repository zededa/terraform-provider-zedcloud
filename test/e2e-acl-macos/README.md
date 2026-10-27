# e2e-acl-macos — UE-203 interface ACLs on a virtual EVE node, against a local zedcloud

Checks that the port maps Terraform configures on `zedcloud_application_instance`
are what EVE gets, on one virtual EVE node running on this Mac (vfkit + zedamigo),
onboarded to the local controller with zedcloud #8548 (ENG-3000).

The node setup (`nodes.tf`, `cloud_objects.tf`, `vars.tf` up to the stages) is copied
from [`test/e2e-snapshot-macos`](../e2e-snapshot-macos); read its README and
[`test/e2e-cluster-macos`](../e2e-cluster-macos)'s for the prerequisites and macOS traps.
The project has interface ordering enabled.

`app.tf` deploys CirrOS VM instances of an app whose bundle has the ENG-3000 eth0 rules:
allow 0.0.0.0/0, 9022->22, 9080->80.

## Setup

```shell
cp e2e-acl-macos.tfrc.example e2e-acl-macos.tfrc      # edit the mirror path
cd ../../v2 && CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o /tmp/p . && \
  cp /tmp/p ~/.terraform.d/plugins/registry.opentofu.org/zededa/zedcloud/99.0.0/darwin_arm64/terraform-provider-zedcloud_v99.0.0
cd ../test/e2e-acl-macos
ZC_RUN_ID=acl1 ./run.sh init
```

## Stages

`ZC_RUN_ID` must be fresh for every node build. `./check.sh` prints, per instance, the
stored ACLs, the controller status, the ACLs in EVE's `AppInstanceConfig`, and EVE's
port-map iptables rules.

| Command | What it checks |
|---|---|
| `./run.sh apply -var stage=0` | node onboarded |
| `./run.sh apply -var stage=1` | instance X with the bundle's 9022->22 edited to 8022->22 |
| `./run.sh apply -var stage=2` | + instance Y on the same node with 7022->22, 7080->80 |
| `./run.sh apply -var stage=2 -var x_ssh_port=6022` | ACL update of a running instance |
| `./run.sh destroy -var stage=2 -var x_ssh_port=6022` | teardown |

## Results, 2026-10-27 (EVE 17.0.0-lts-kvm-arm64, local zedcloud with #8548, provider `ue-203-app-instance-acls`)

Node onboarded in about 5 minutes, including the installer build (`RUN_STATE_ONLINE`, `ADMIN_STATE_REGISTERED`).

| Stage | Controller stores | EVE `AppInstanceConfig` | EVE iptables DNAT on 192.168.127.2 | Instance status | Plan after |
|---|---|---|---|---|---|
| 1: X = allow, 8022->22, 9080->80 | as configured | allow, 8022->22, 9080->80 | 8022->X:22, 9080->X:80; **no 9022** | `RUN_STATE_ONLINE`, no errors | No changes |
| 2: + Y = allow, 7022->22, 7080->80 | X unchanged; Y as configured | Y: allow, 7022->22, 7080->80 | + 7022->Y:22, 7080->Y:80 | both `RUN_STATE_ONLINE`, **no "overlapping portmaps"** | No changes |
| 3: X 8022 -> 6022 | X: allow, 6022->22, 9080->80 | X: allow, 6022->22, 9080->80 | 6022->X:22 replaces 8022 | both `RUN_STATE_ONLINE` | No changes |

Stage 3's plan showed exactly one change, `value = "8022" -> "6022"` on X's ACL, and the
update completed in 2 s. No "Interface ACLs were not stored as configured" warning.

Before #8548 (ENG-3000 retest), X would have received 9022->22 as well, and Y would have
received 9022 and 9080 and failed with `overlapping portmaps`. Before UE-203 the stage 3
change would have planned "No changes" and never reached the device.

Not checked: traffic through the port maps. gvproxy only forwards the node's ssh ports to
the Mac, so the check stops at the iptables rules EVE installs.
