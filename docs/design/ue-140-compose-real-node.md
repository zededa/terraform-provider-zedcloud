# UE-140 — Un-skip `TestApplicationInstance_CreateCompose` against a real node

Status: implementation plan
Ticket: [UE-140](https://zededa.atlassian.net/browse/UE-140) (parent [UE-131](https://zededa.atlassian.net/browse/UE-131))
Ref: [`e2e-virtual-eve-node-testing.md`](./e2e-virtual-eve-node-testing.md) §6.3–§6.5, [`ue-139-fixture-suffix-tokens.md`](./ue-139-fixture-suffix-tokens.md)

> **Revision note.** §1.4 and §6 were rewritten after a working customer configuration
> (`1-create-infra.tf`, Edge Field Day demo) became available. It is the only known example of
> compose apps deployed through this provider against real hardware, and it corrected three
> things this plan originally got wrong — see §1.4. Where the customer config and my reading of
> seine disagree, the customer config wins: it demonstrably works.

---

## 0. Headline: the ticket splits in two, and the split is not where it looks

UE-140 as written reads like a three-line change — re-gate, repoint the fixture, add a wait.
Two of those three are indeed small. The third is not, and the reason is entirely on the
ZedCloud side.

A Docker Compose app instance in ZedCloud is **not a thing EVE ever receives**. It has no
config message, no volumes, no interfaces, no drives. It exists only as a DB row plus an entry
inside an *encrypted Patch Envelope* attached to a **separate, mandatory, one-per-device
`DEPLOYMENT_TYPE_DOCKER_RUNTIME` app instance**. That runtime VM boots a ZEDEDA-built QCOW2
image containing the `zrun` agent; `zrun` reads the envelope, pulls the compose images from a
container-registry datastore, runs `docker compose up`, and reports a `NestedAppInstanceStatus`
back up through the *runtime's* info message. Only then does the compose instance's `swState`
become anything other than `SW_STATE_UNSPECIFIED`.

So "assert the compose app reaches `SW_STATE_RUNNING` on a real node" transitively requires a
working zrun runtime VM on that node. That is a genuinely new capability for the e2e lab, and
it has one hard external dependency (§3.1) and one hard topology conflict (§3.2) that no amount
of Go or HCL in this repo can resolve.

**Recommendation: land the `CI`-degating work and the compose *API* path now (Phase A, small,
self-contained, satisfies two of the three acceptance criteria), and split the real-node
compose run into its own ticket (Phase B) blocked on locating the zrun runtime bundle and on a
second NIC.**
Phase B is §6.4 "Phase 2" of the parent design doc arriving early and wearing a different hat.

---

## 1. What the code actually says

### 1.1 Provider side — verified inventory

| Fact | Location |
|---|---|
| `TestApplicationInstance_CreateCompose`, gated on `os.Getenv("CI") != ""` | `v2/resources/application_instance_test.go:65-114` (skip at 69–71) |
| `TestApplicationInstance_RealNode` — the working pattern to copy | `v2/resources/application_instance_real_node_test.go:44-102` |
| `testhelper.RealNode(t) RealNodeInfo` — **exists** | `v2/testing/realnode.go:69-112` |
| `testhelper.WaitForAppInstanceSwState(t, id, want, timeout)` — **exists** | `v2/testing/realnode.go:245-284` |
| `testhelper.WaitForRunState` (named in design doc §6.3) | **ABSENT** — only the unexported `waitNodeOnline` |
| `ZEDCLOUD_ACC_REAL_NODE` (`testhelper.EnvRealNode`) | `v2/testing/realnode.go:32` |
| `ZEDCLOUD_ACC_ELEVATED` | **ABSENT from all code** — design-doc proposal only |
| `create_real_node.tf` and its `__NODE_ID__` / `__SUFFIX__` contract | `v2/resources/testdata/application_instance/create_real_node.tf:1-29` |
| Fixture token allow-list, enforced by a failing test | `v2/testing/fixtures_test.go:22-27`, `TestFixtureTokensAreKnown:35` |
| `getAcKind`: `APP_TYPE_DOCKER_COMPOSE` → `"ComposeManifest"` | `v2/resources/application.go:344-364` |
| Compose payload lives on the **application manifest**, not the instance | `v2/schemas/vm_manifest.go:594-603`; `v2/models/app_instance.go:977-993` marks both fields `validate.ReadOnly` |

Six `os.Getenv("CI")` sites, all `!= ""`:

| File | Line | Test |
|---|---|---|
| `v2/resources/application_instance_test.go` | 69 | `TestApplicationInstance_CreateCompose` |
| `v2/resources/enterprise_test.go` | 23 / 76 / 131 / 167 | `TestEnterprise_Create`, `_WhiteLabeling`, `_WhiteLabelingControllerHostURL`, `_WhiteLabelingLegacyAttributes` |
| `v2/resources/profile_deployment_service_resource_test.go` | 22 | `TestProfileDeployment_Create` |

`CI` is set at `.github/workflows/test.yml:11` and — importantly — at
`.github/workflows/e2e.yml:273`, on the acceptance step itself. GitHub Actions also sets it
implicitly in every job.

### 1.2 ZedCloud side — what a compose instance requires

Server-side validation, all 400s, from `srvs/seine/appproc.go` and `srvs/seine/appinstproc.go`:

**Compose *bundle*** (`acKind: "ComposeManifest"`) must have: `deploymentType =
DEPLOYMENT_TYPE_STAND_ALONE`; exactly one of `dockerComposeYamlText` (valid base64) or
`dockerComposeTarImageName`; empty `vmmode`, `interfaces`, `resources`; `enablevnc` false;
`cpuPinningEnabled` false; `runtimeVersion` and `runtimeProtocolVersion` left `Unknown`; at
most one cloud-init variable group, no BOOLEAN/DROPDOWN variables (`appproc.go:415-582`).
Separately, **no** `images` refs (`appproc.go:110-113`), and every non-`FILE` cloud-init
variable needs a non-empty default (`appproc.go:2391-2400`). Every `datastoreIdList` entry
must be `DATASTORE_TYPE_CONTAINERREGISTRY` and must grant the bundle's projects
(`appproc.go:173-199`).

**Compose *instance*** must omit `vminfo`, `drives`, `interfaces`, `startDelayInSeconds`,
`clusterID`, `edgeNodeCluster`, `remoteConsole`, `allowSnapshots`; `maxSnapshots` must be 0
(`appinstproc.go:5462-5605`).

**Hard prerequisite** — `fillMfestInfo`, `appinstproc.go:1506-1534`:

```go
if adata.AppType == int(zcommon.AppType_APP_TYPE_DOCKER_COMPOSE) {
    // check if runtime instance running on device, if not, don't allow creation of compose app inst
    exists, err := lookupAppInstsOfGivenDeploymentType(ctx, txCtx, adata.EnterpriseId,
        adata.DeviceId, zcommon.DeploymentType_DEPLOYMENT_TYPE_DOCKER_RUNTIME)
    if err != nil || exists == nil { /* 400 */ }
    adata.RuntimeInstanceID = exists.Id
    // ... then validateDockerComposeImageAndDatastore (1522): the tar image must be
    // IMAGE_STATUS_READY and the datastore DATASTORE_STATUS_ACTIVE.
```

**Runtime instance** (`appinstproc.go:1100-1146`, `srvs/seine/runtimes.go:16-88`):

- device must report `API_CAPABILITY_ENCRYPTED_PATCH_ENVELOPE` (= 14,
  `libs/zmsg/eveproto/proto/info/info.proto:676`), else *"device %s is not capable of running
  docker runtime"*
- exactly one runtime per device
- **adapter 0 → a `NETKINDSWITCH` network instance; adapter 1 → a `NETKINDLOCAL` network
  instance with `Port == ""` (airgapped)**; no direct-attach adapters
- `drives[0]` references an `IMAGE_TYPE_DOCKER_RUNTIME` image, status READY; a blank
  platform-managed drive is injected at index 1 (zrun hardcodes `vdb` →
  `/var/lib/zrun-data`, `runtimes.go:145`)
- runtime bundle needs `runtimeVersion: RuntimeVersion_V1`, `runtimeProtocolVersion:
  RuntimeProtocolVersion_V1`, non-zero `PersistentRuntimeSizeBytes`, ≥2 non-direct-attach
  interfaces

**Status path** — `srvs/ganges/device.go:2054-2131` → `srvs/yamuna/appinstproc.go:197-220`.
The compose row's `swState` is whatever zrun reports, unmarshalled from an
`appstatus.AppInstanceStatus` in the row's `Opaque` column. Empty until zrun speaks.

No feature flag, no enterprise toggle, no licence gate exists for compose — the capability
check is the only gate.

### 1.3 The existing `create_compose.tf` fixture is synthetic and cannot work as-is

`v2/resources/testdata/application_instance/create_compose.tf` (426 lines) fabricates its own
project → brand → model → edgenode → datastores → images → apps → NIs → instances. Specific
defects for a real run:

1. **`activate` is unset**, and `create_compose.yaml:1` is `activate: false` — nothing is ever
   pushed to a device today. The test asserts an API round-trip only.
2. **The base64 in `docker_compose_yaml_text` (line 286) decodes to a Dockerfile, not a
   compose file** (`FROM alpine:latest` / `RUN docker login …`). seine never parses it, so the
   API accepts it; zrun would not.
3. The runtime image `image_rel_url = "test_tf_provider-runtime-image"` on `fake-datastore.com`
   does not exist.

**Not a defect** (an earlier revision of this plan called it one): `test_compose_image` points
at the *file* datastore rather than the container registry. That is correct. The compose TAR is
an ordinary file artifact and lives on an HTTP/SFTP datastore; the container-registry datastore
exists to supply credentials for the images referenced *inside* the compose YAML. seine
constrains only the `datastoreIdList` entries to `CONTAINERREGISTRY`
(`appinstproc.go:541-578`), never the tar image's own datastore. The customer config does
exactly the same thing — its four `IMAGE_TYPE_DOCKER_COMPOSE_TAR` images all sit on
`DATASTORE_TYPE_HTTP`.

### 1.4 The customer configuration — what a working compose deployment actually looks like

`1-create-infra.tf` (Edge Field Day demo, two Proxmox nodes) deploys compose apps through this
provider against real hardware. It is worth more than any amount of reading seine, because it
shows which of seine's many requirements the *provider* satisfies implicitly. Four findings,
in descending order of how much they change this plan.

**(a) The runtime app bundle is referenced by bare UUID, not created.**

```hcl
resource "zedcloud_application_instance" "tf-dc-1" {
  name     = "TF-RUNTIME-1"
  app_id   = "bc317ad0-ad0a-469d-8ad5-e9c3c2223729"   # not a resource reference
  activate = true
  ...
}
```

The customer never declares a `zedcloud_application` for the runtime, and never declares the
`IMAGE_TYPE_DOCKER_RUNTIME` image. They point at a bundle that already exists in their
enterprise. Combined with seine's `Origin_ORIGIN_GLOBAL` import machinery
(`appproc.go:1257, 1963, 2049`, "Fetching bundle imported from global enterprise"), the
strong implication is that **the zrun runtime bundle is a ZEDEDA-published global/marketplace
bundle that gets imported, not built.**

This substantially downgrades blocker §3.1. The question changes from *"can someone publish a
zrun image for us"* — open-ended, another team's roadmap — to *"what is the runtime bundle's
UUID on alpha, and is it visible to the test enterprise"*, which is a lookup. It may still turn
out that alpha has no such bundle, but that is a far cheaper question to answer. **Ask it
first.**

**(b) The runtime instance omits `drives` entirely.**

Contradicting an earlier reading of `validateRuntimeInstanceDrives` (`runtimes.go:259-281`),
which does require a boot drive: drives are auto-filled from the bundle manifest's `images`
when the instance supplies none (`appinstproc.go:1482-1504`, *"if Drives are provided in
instance just it, ignore the manifest"*), and validation runs afterwards at
`appinstproc.go:1542`. The blank platform-managed drive is then injected at index 1
(`:1559-1572`). So the fixture writes no `drives` block at all, and must not — setting
`runtime_internal` from the client is a 400.

**(c) `datastore_id_list` is optional.** `appproc.go:173` guards the whole block with
`if len(uInput.GetDatastoreIdList()) != 0`. It carries registry credentials, so it is needed
only for private images. The customer creates a `TF-DOCKER-HUB` container-registry datastore
but does **not** attach it to any compose bundle — their compose YAML pulls public images. A
test fixture using a public image can skip the datastore entirely, removing a resource, a
credential, and the project-access-list subset rule (`appproc.go:185-198`) from the fixture.

**(d) Interfaces, networks and cloud-init — the concrete shape.**

```hcl
# switch NI on a DEDICATED second physical port; eth0 stays management
resource "zedcloud_network_instance" "node1_ni_switch_rt_1" {
  kind = "NETWORK_INSTANCE_KIND_SWITCH"
  type = "NETWORK_INSTANCE_DHCP_TYPE_UNSPECIFIED"
  port = "eth1"
}
resource "zedcloud_network_instance" "node1_ni_local_rt_1" {
  kind = "NETWORK_INSTANCE_KIND_LOCAL"
  type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  port = ""                      # airgapped
  ip { gateway = "10.20.0.1"  subnet = "10.20.0.0/16"
       dhcp_range { start = "10.20.0.100"  end = "10.20.0.200" } }
}

# on the runtime instance:
interfaces { intfname = "eth1-ext"  intforder = 1  netinstname = <switch NI> }
interfaces { intfname = "eth2-int"  intforder = 2  netinstname = <local NI>
             ipaddr   = "10.20.0.201" }          # static, outside the DHCP range
custom_config {
  add = true   override = true   allow_storage_resize = true
  field_delimiter = "###"   name = "cloud-config"
  template = base64encode(file("./cinit/rt-1.txt"))
}
```

Four things a fixture author would otherwise get wrong: the logical interface names are
`eth1-ext` / `eth2-int` (they come from the runtime *bundle's* manifest, not from the node's
adapters); `intforder` is explicit and load-bearing, since seine validates adapter 0 as the
switch and adapter 1 as the local NI by position (`runtimes.go:76-87`); the local interface
takes a **static** `ipaddr` deliberately placed outside the NI's DHCP range; and the runtime
instance carries a **cloud-init `custom_config`**, which this plan had not accounted for at
all.

### 1.4b The runtime cloud-init is mandatory, and it hides the plan's worst failure mode

`rt-1.txt` (the customer's runtime cloud-init) settles what the runtime bundle expects. Three
things matter.

**(i) The static IP is the status channel — and it is duplicated.** The cloud-init writes a
systemd drop-in for the zrun agent:

```sh
- echo "[Service]" > /etc/systemd/system/zrun-agent.service.d/zrun-env.conf
- echo 'Environment=COLLECTSTATSIPADDR="10.20.0.201"' >> /etc/systemd/system/zrun-agent.service.d/zrun-env.conf
```

That address is the *same literal* as the runtime instance's `eth2-int` `ipaddr` in
`1-create-infra.tf`, and it is what seine resolves `$zri.system.edge-instance.ip` from — the
runtime's `CollectStatsIPAddr` (`appinstutils.go:41-49`,
`docker_compose_app_inst.go:218-252`). So the IP appears in two places, in two different
languages, and **nothing validates that they agree.**

This is the single nastiest hazard in Phase B. If they diverge, everything upstream succeeds —
bundle created, instance created, runtime boots, Terraform goes green — and the compose
instance's `swState` simply stays `SW_STATE_UNSPECIFIED` forever, because zrun's status never
arrives. `WaitForAppInstanceSwState` then reports a bare timeout with `last: swState=
runState=` and no error, since `errDescriptions(st.ErrInfo)` has nothing to report
(`realnode.go:264-272`). Someone will spend a day on that.

Mitigations, in the fixture: derive both from one place rather than typing the literal twice —
the cleanest is a `__RUNTIME_IP__` fixture token (added to `knownTokens`) interpolated into
both the `ipaddr` attribute and the `base64encode(templatefile(...))` cloud-init. Failing that,
at minimum put a comment at each site naming the other. And consider having
`WaitForAppInstanceSwState`, or a Phase B wrapper, special-case "still UNSPECIFIED after N
polls" with a message pointing at `COLLECTSTATSIPADDR` — a timeout that explains its own most
likely cause is worth the ten lines.

**(ii) The runtime image ships a contract; cloud-init is not optional.** `rt-1.txt` calls
`/home/zededa/bin/disk-partition.sh` (partitions the vdb persistent disk that
`runtimes.go:145` pins at index 1), `/home/zededa/bin/setup_dcompose.sh`,
`/home/zededa/bin/last_things.sh`, then `systemctl enable/start zrun-agent`. It also writes
into `/etc/systemd/system/zrun-agent.service.d/` without creating it. All of that presupposes a
purpose-built ZEDEDA QCOW2 — more evidence for §3.1's "published bundle, not one you build" —
and it means **a runtime instance with an empty `custom_config` boots and does nothing.** Any
Phase B fixture must carry a real cloud-init.

**(iii) Compose apps can be coupled to the runtime's cloud-init — so pick one that is not.**
The customer's cloud-init pre-creates `/var/lib/zrun-data/{ignition,mosquitto}` and seeds
`mosquitto.conf`, because their compose files bind-mount those host paths
(`ignition_mosquitto.yaml:36`, `:50`). A test compose app that needs host state has a hidden
dependency on cloud-init contents, which is a bad thing to debug.

**So the test's compose app should use a public image, no bind mounts, no named volumes, and
no `${VAR}` interpolation** (the customer's `IGNITION_UID=${UID}` relies on env plumbed through
seine's `EnvData`). A single `nginx:stable-alpine` with a port mapping satisfies all of that
and still exercises the full chain: patch envelope → registry pull → container running.

One hygiene note: `rt-1.txt` contains a real SHA-512 password hash (and names the plaintext in
a comment), and the compose files carry gateway admin credentials. **Do not copy any of it into
a repo fixture.** Write a fresh cloud-init that does the structural work — partition, set up
compose, set `COLLECTSTATSIPADDR`, start the agent — with an SSH key instead of a password, or
no interactive login at all.

The compose app and instance shapes confirm §6's sketch. The bundle is as expected
(`ac_kind = "ComposeManifest"`, `app_type = "APP_TYPE_DOCKER_COMPOSE"`,
`deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"`,
`docker_compose_yaml_text = base64encode(file("./yaml-files/grafana.yaml"))` — note
`base64encode(file(...))`, far more readable in a fixture than a pasted blob). The instance
carries only `name`, `title`, `description`, `app_id`, `app_type`, `device_id`,
`logs { access }`, `manifest_info { transition_action = "INSTANCE_TA_NONE" }` and a
no-op `custom_config` — no interfaces, drives or vminfo, exactly as
`appinstproc.go:5462-5605` demands.

**(e) Corroboration for §3.2.** The customer's model declares both `eth0` and `eth1`, the node
assigns `eth0` to management and `eth1` to `ADAPTER_USAGE_APP_SHARED`, and the switch NI binds
`eth1`. Nobody puts the switch NI on the management port. Treat the second NIC as mandatory,
not as a thing to try shortcuts around.

### 1.5 CI wiring

`.github/workflows/e2e.yml:269-306` runs `go test -run TestApplicationInstance_RealNode .` —
a hard filter. Un-skipping `CreateCompose` changes nothing in CI until that regex widens, and
`CI: "true"` on line 273 means the current guard would skip it anyway.

---

## 2. Scope decision

| Phase | Content | Blocked on |
|---|---|---|
| **A** | Retire `CI` as a capability gate; add `ZEDCLOUD_ACC_ELEVATED`; keep compose as an API-only test, fixed and honest | nothing |
| **B** | Real compose app reaching `SW_STATE_RUNNING` on the virtual EVE node | §3.1 locate the runtime bundle, §3.2 second NIC |

Phase A satisfies AC #2 and AC #3 outright and converts AC #1 into a well-scoped follow-up.

---

## 3. Phase B blockers — resolve before committing to a date

### 3.1 The zrun runtime bundle has to be located (downgraded from "published")

The runtime app instance boots a real QCOW2. The Python API suite sidesteps this entirely: it
uploads `image.zededa_greengrass_config.img` as a stand-in against a fake S3 datastore and
drives a **simulated** device (`zdevice.DeviceSimulate`), so nothing ever downloads or boots
(`tests/apiTest/tests/test_app_instance.py:3311-3474`). Neither repo contains a published zrun
image name, URL, tag, or datastore — the only `zrun-image` strings in either repo are
placeholder literals in seine's own unit tests (`srvs/seine/runtimes_test.go:21`,
`runtime_volumes_test.go:21`).

But per §1.4(a), the customer does not build one either — they reference an existing bundle by
UUID, and seine has global-enterprise bundle import. So this is most likely a **lookup, not a
publishing request**.

**Action, in order:**

1. Find the runtime bundle on alpha — list app bundles filtered to
   `deploymentType = DEPLOYMENT_TYPE_DOCKER_RUNTIME`, and check whether a global-origin one is
   importable into the test enterprise. If yes, most of this blocker evaporates and Phase B
   needs only §3.2.
2. If not, ask the zrun/runtime team for the image location, a known-good
   `runtimeVersion`/`runtimeProtocolVersion` pair, and the expected boot time on a 4-CPU/8 GB
   QEMU guest.
3. Either way, get the runtime cloud-init template (§1.4(d)) — the customer's `rt-1.txt` or an
   equivalent. A runtime that boots misconfigured fails in a way that will read as a test bug.

Also note the fixture consequence if the bundle is imported rather than created: the test
cannot reference it as a Terraform resource, so the bundle UUID arrives as a **new fixture
token**. `knownTokens` (`v2/testing/fixtures_test.go:22`) admits only `SUFFIX` and `NODE_ID`
and `TestFixtureTokensAreKnown` fails the build on anything else — so adding, say,
`__RUNTIME_APP_ID__` means editing that map *and* plumbing the value through
`MustGetTestInputWithVars` and the workflow env. Cheap, but it must be deliberate.

### 3.2 The runtime needs two NICs; the e2e VM has one (hard, but §6.4 already scoped it)

The runtime demands adapter 0 on a **switch** NI and adapter 1 on an **airgapped local** NI.
Only adapter 0 needs a physical port, so on paper one NIC suffices. In practice `test/e2e/`
gives the VM a single QEMU SLIRP NIC which is also the *management* uplink
(`test/e2e/node.tf`; `create_real_node.tf:80-82` — *"NOTE eth0, not eth1: the VM is created
with a single QEMU SLIRP NIC"*), and `cluster_objects.tf` declares only `eth0` in
`io_member_list`. Binding a switch NI to the sole management port is how you lose the node.

This is exactly design-doc §6.4 Phase 2: `zedamigo_bridge` + `zedamigo_tap` +
`zedamigo_dhcp_server` via `extra_qemu_args`, which needs `use_sudo = true` and a fourth
sudoers allow-list entry for the bootstrapped provider binary. Interface names are
kernel-global, so the new resources must carry `__SUFFIX__`.

**Action:** confirm with the EVE/lab side whether a switch NI on a shared management port is
survivable; assume not, and budget the second NIC.

### 3.3 Runtime boot time (soft)

Serialized chain before the compose instance can report anything: runtime image download →
volume creation → cloud-init (partitions vdb, sets up compose, starts `zrun-agent`) → runtime
`SW_STATE_RUNNING` → patch-envelope pickup → registry pull → compose `SW_STATE_RUNNING`. The
existing real-node test already allows 15 minutes for a single `nginx:stable-alpine`. Budget
30–40 min for the compose chain and raise the job's `timeout-minutes: 60`.

Wait on the **runtime** instance first and fail there if it does not come up. A single wait on
the compose instance cannot distinguish "runtime never booted" from "zrun cannot report"
(§1.4b(i)) from "registry pull is slow" — and all three present identically as
`swState=UNSPECIFIED` until the deadline.

### 3.4 API capability of the lab EVE build (cheap to check, do it first)

`test/e2e/vars.tf` pins `eve_tag = "16.0.1-lts-kvm-amd64"`. Capability 14 is old enough that
this should be fine, but it is a one-command check against the already-onboarded node and it
invalidates Phase B outright if it fails:

```sh
curl -sH "Authorization: Bearer $TF_VAR_zedcloud_token" \
  "https://zedcontrol.alpha.zededa.net/api/v1/devices/id/$DEVICE_ID/status" \
  | jq '.apiCapability'   # need >= 14 (API_CAPABILITY_ENCRYPTED_PATCH_ENVELOPE)
```

---

## 4. Phase A — implementation steps

### A1. Add the capability-gate accessors

`v2/testing/realnode.go`, alongside `EnvRealNode:32`:

```go
// EnvElevated gates tests that need operator/distributor permissions the
// ordinary CI service account does not hold -- child-enterprise creation and
// white-labeling. Kept separate from EnvRealNode: they are different
// capabilities and conflating them is what made CI a catch-all in the first place.
const EnvElevated = "ZEDCLOUD_ACC_ELEVATED"

// Elevated skips the calling test unless EnvElevated is enabled.
func Elevated(t *testing.T) {
	t.Helper()
	if !envEnabled(EnvElevated) {
		t.Skipf("%s is not set; skipping test that requires elevated enterprise permissions", EnvElevated)
	}
}
```

Reuses the existing `envEnabled` (`realnode.go:194`, accepts `1`/`true`/`yes`).

### A2. Re-gate the four enterprise tests

`v2/resources/enterprise_test.go` — replace the identical three-line block at 23, 76, 131, 167:

```go
	testhelper.Elevated(t)
```

Drop the now-unused `os` import if nothing else in the file uses it. **Do not** assume the skip
reason: the ticket is right that it is undocumented. File a short note or a follow-up
confirming *why* these fail (child-enterprise create permission is the hypothesis) so the next
person does not inherit the same mystery. Leave `TestEnterpriseResource_InternalValidate:64`
ungated — it is a pure schema test.

### A3. Leave `TestProfileDeployment_Create` alone, but stop it from using `CI`

Per the ticket and design doc, this is a real flake, not a capability gap. Re-gating it behind
a new env var would launder a known bug. But it must also stop reading `CI`, or AC #3 is not
met. Use an unconditional skip that names the ticket:

```go
// Skipped: fails in CI, passes locally -- a real flake, tracked in <UE-XXX>.
// Deliberately NOT behind a capability flag: that would hide a bug behind a
// feature switch. Delete this line when the flake is fixed.
func TestProfileDeployment_Create(t *testing.T) {
	t.Skip("flaky, see <UE-XXX>")
```

**File the flake ticket first** so the comment can cite a real key.

### A4. Fix the compose fixture's real defects, keep it API-only

`v2/resources/testdata/application_instance/create_compose.tf`:

- line 286: replace the Dockerfile base64 with a real compose file. Follow the customer's
  idiom and keep it a readable sibling file rather than a blob —
  `docker_compose_yaml_text = base64encode(file("./testdata/application_instance/compose/nginx.yaml"))`
  — so the next reader can see what is being deployed without decoding anything:

```yaml
services:
  web:
    image: nginx:stable-alpine
    ports: ["8080:80"]
```

  (Leave `test_compose_image`'s datastore alone — it is correct, see §1.3.)

- add a header comment stating plainly what this fixture does and does not prove: it exercises
  the controller's compose validation path (`appproc.go:415-582`) with `activate = false`; it
  does **not** deploy anything. Point at Phase B for the real thing.

Rename the test to match, so the name stops over-promising:

```go
// TestApplicationInstance_CreateCompose exercises the controller-side validation
// and round-trip for a docker-compose application bundle and instance. It does
// not deploy: activate is false and the node is fabricated. Deploying a compose
// app for real requires a docker-runtime instance on an onboarded device --
// see docs/design/ue-140-compose-real-node.md and <Phase B ticket>.
func TestApplicationInstance_CreateCompose(t *testing.T) {
```

and **delete the `CI` skip entirely** — with the fixture fixed there is no reason this cannot
run in the ordinary PR job. If it turns out it cannot (see §5), gate it on `TF_ACC` like the
rest of the suite rather than reintroducing a capability flag.

### A5. Assert the compose payload instead of ignoring it

`testApplicationInstanceAttributes` (`application_instance_test.go:164-177`) lists
`"DockerComposeYamlText"` at line 174 in `ignoredFields` — the one field that makes this a
compose test. It is ignored because `models.AppInstance` marks it read-only and the controller
populates it from the bundle (`appinstproc.go:1531-1533`). That is a reason to assert it
*equals the bundle's value*, not to skip it.

The plumbing is already there and was verified: `zedcloud_application_instance` uses
`zschema.ApplicationInstance()` (`v2/resources/application_instance.go:23,33` →
`v2/schemas/application_instance.go:510`), which declares `docker_compose_yaml_text` as
`Computed` at line 775 and populates it in `SetApplicationInstanceResourceData` at line 460.
The bundle side is `Optional` at `v2/schemas/vm_manifest.go:600`. So a plain attribute-pair
check works — no custom check func needed:

```go
resource.TestCheckResourceAttrPair(
    "zedcloud_application_instance.test_tf_appinst_compose", "docker_compose_yaml_text",
    "zedcloud_application.test_tf_app_compose", "manifest.0.docker_compose_yaml_text",
),
```

Leave the shared `ignoredFields` list alone — it is used by both tests in the file and the
non-compose one has nothing to compare.

### A6. Grep-clean and lock it in

```sh
grep -rn 'Getenv("CI")' --include='*.go' .   # must return nothing
```

Then add a guard test next to the existing fixture guards in `v2/testing/fixtures_test.go`,
which already has the "fail the build on drift" idiom:

```go
// TestNoCIEnvGating fails if CI is reintroduced as a capability gate. CI says
// "an automated runner is executing this", never "this environment can do X" --
// conflating the two is what made four unrelated tests share one switch.
func TestNoCIEnvGating(t *testing.T) { /* walk ../resources, parse *_test.go, flag os.Getenv("CI") */ }
```

### A7. Workflow changes

`.github/workflows/e2e.yml`:

- **delete `CI: "true"` at line 273.** It is now load-bearing for nothing and actively
  misleading.
- leave `-run TestApplicationInstance_RealNode` as-is for Phase A; widen it in Phase B.

`.github/workflows/test.yml`:

- `CI: true` at line 11 can stay (Actions sets it anyway) but no test reads it after A6.
- the four enterprise tests and the profile-deployment test now skip via their own mechanisms;
  the compose test now runs. Confirm the job stays green — see §5.

---

## 5. Risk in Phase A: the compose test may need controller state the PR job lacks

`create_compose.tf` creates a brand, a model, an edgenode with a fabricated serial, two
datastores and two applications on the shared alpha enterprise. It has presumably not run in CI
since the skip was added, so **the reason for the original skip is itself unverified** — the
message says only "Skipping docker compose test for CI environment".

**Do this before writing any code**: run the test as-is against alpha from a workstation and
see what actually happens.

```sh
cd v2/resources
TF_ACC=1 TF_CLI_CONFIG_FILE=$PWD/../../dev.tfrc \
TF_VAR_zedcloud_url=zedcontrol.alpha.zededa.net TF_VAR_zedcloud_token='<token>' \
ZEDCLOUD_TEST_SUFFIX=_ue140local \
  go test -v -timeout 20m -run TestApplicationInstance_CreateCompose .
```

Three outcomes:

- **passes** — the skip was cargo-cult; A4/A6 are straightforwardly correct.
- **fails on permissions** (brand/model/enterprise creation) — it belongs behind
  `ZEDCLOUD_ACC_ELEVATED` after all, and A4's "delete the skip entirely" becomes
  `testhelper.Elevated(t)`. Note this is a *different* answer from the ticket's and worth
  reporting back.
- **fails on validation** — a real provider or fixture bug is hiding behind the skip. Fix it
  in A4 and say so in the PR; this is the most valuable possible outcome.

Whatever happens, record it in the PR. The point of UE-140 is to stop tests being skipped for
unrecorded reasons, so replacing one undocumented skip with another would miss it.

---

## 6. Phase B — sketch, for the follow-up ticket

Only once §3.1 and §3.2 clear.

1. **`test/e2e/node.tf`**: add the second NIC (§6.4 Phase 2 — `zedamigo_bridge` +
   `zedamigo_tap` + `zedamigo_dhcp_server`, `use_sudo = true`, `__SUFFIX__`-carrying names,
   plus the fourth sudoers entry). Add `eth1` to `io_member_list` in `cluster_objects.tf`.
2. **New fixture** `testdata/application_instance/create_compose_real_node.tf`, modelled on
   `create_real_node.tf`, using `__NODE_ID__` / `__SUFFIX__` only — the token allow-list at
   `v2/testing/fixtures_test.go:22` admits nothing else, and `TestFixtureTokensAreKnown` fails
   the build on drift. If the fixture needs the project or model id, either add the token to
   `knownTokens` *and* plumb it from `test/e2e/outputs.tf` (`project_id` and `model_id` are
   already exposed but unconsumed), or create the objects in the fixture.

   Contents, in dependency order, following `1-create-infra.tf` rather than inventing a shape:

   - switch NI on the **second** physical port (`NETWORK_INSTANCE_KIND_SWITCH`,
     `type = "NETWORK_INSTANCE_DHCP_TYPE_UNSPECIFIED"`); airgapped local NI
     (`NETWORK_INSTANCE_KIND_LOCAL`, `type = "…DHCP_TYPE_V4"`, `port = ""`, explicit
     `ip { gateway/subnet/dhcp_range }`)
   - runtime app bundle: **prefer referencing the existing one** (§3.1) over declaring it. Only
     if no importable bundle exists, declare it — `DEPLOYMENT_TYPE_DOCKER_RUNTIME`,
     `RuntimeVersion_V1`, `RuntimeProtocolVersion_V1`, non-zero
     `persistent_runtime_size_bytes`, two non-direct-attach interfaces, boot image first
   - runtime app instance: `activate = true`, **no `drives` block** (auto-filled, §1.4(b)), the
     two `interfaces` with explicit `intforder` 1 and 2 and the bundle's logical names, static
     `ipaddr` on the local one outside the DHCP range, and a **real cloud-init
     `custom_config`** — partition vdb, run the image's compose setup, write
     `COLLECTSTATSIPADDR`, start `zrun-agent` (§1.4b(ii)). Write a fresh one; do not copy the
     customer's credentials.
   - the static IP must be **one value used twice** — `ipaddr` and `COLLECTSTATSIPADDR`.
     Interpolate a `__RUNTIME_IP__` token into both rather than typing it in two languages
     (§1.4b(i)); adding the token means editing `knownTokens` at
     `v2/testing/fixtures_test.go:22`.
   - compose app bundle: `ac_kind = "ComposeManifest"`,
     `deployment_type = "DEPLOYMENT_TYPE_STAND_ALONE"`,
     `docker_compose_yaml_text = base64encode(file("./testdata/.../compose/<name>.yaml"))` —
     keep the YAML a readable sibling file, not a pasted blob; **omit `datastore_id_list`**
     (§1.4(c)) and use a public image with no bind mounts, no named volumes and no `${VAR}`
     interpolation (§1.4b(iii))
   - compose app instance: `name`, `title`, `app_id`, `app_type`, `device_id`, `activate`,
     `logs { access = true }`, `manifest_info { transition_action = "INSTANCE_TA_NONE" }`,
     `max_snapshots = 0`. Nothing else — `vminfo`, `drives`, `interfaces`, `cluster_id`,
     `start_delay_in_seconds`, `remote_console`, `allow_snapshots` are each an independent 400.

   Note the asymmetry that bites: `runtime_version` / `runtime_protocol_version` must be **set**
   on the runtime bundle (`appproc.go:702-716`) and **left Unknown** on the compose bundle
   (`appproc.go:567-581`). And `datastore_id_list`, if used at all, is accepted *only* for
   compose apps (`create_real_node.tf:105-107`).

3. **New test** `TestApplicationInstance_ComposeRealNode` in
   `application_instance_real_node_test.go`: `testhelper.RealNode(t)` →
   `MustGetTestInputWithVars` → wait for the **runtime** instance to reach
   `SWStateSWSTATERUNNING` → then the **compose** instance. Both via
   `testhelper.WaitForAppInstanceSwState`. Mixing `testhelper.Suffix()` (env-only) with
   `RealNode`'s generated fallback suffix produces mismatched names — use `node.Suffix`
   throughout, as `application_instance_real_node_test.go:49-55` does.
4. **Richer assertion, cheaply available**: `models.AppInstStatusMsg.ContainerStatusList`
   (`v2/models/app_inst_status_msg.go:128`, elements typed `ComposeContainerStatus` whose
   `ContainerState` is at `compose_container_status.go:32`) carries per-service state. A helper asserting every
   container is `Running` is a much stronger signal than aggregate `swState`, and
   `testhelper.apiGet` already fetches the status message that carries it.
5. **Workflow**: widen the `-run` regex to
   `-run 'TestApplicationInstance_(RealNode|ComposeRealNode)'`, raise `-timeout 40m` → `90m`
   (design doc §6.5) and `timeout-minutes: 60` → `120`. Keep the suite sequential — no
   `t.Parallel()`, one shared device.
6. **Teardown**: compose instances block deletion of the runtime instance
   (`appinstproc.go:3019-3041`), so destroy order matters. Terraform's dependency graph gets
   this right if and only if the compose instance genuinely depends on the runtime instance —
   add an explicit `depends_on` rather than relying on it falling out.

---

## 7. Acceptance criteria — mapped

| AC | Phase | Status after Phase A |
|---|---|---|
| `CreateCompose` passes against the real node and asserts running state | B | Deferred, with blockers named and a fixture that is at least honest about what it proves |
| Suite still green with no real node (tests skip) | A | Met — `RealNode`/`Elevated` both `t.Skip` when unset |
| No remaining use of `CI` as a capability gate | A | Met, and guarded against regression by A6 |

---

## 8. Open questions to raise before starting

1. **Does an importable `DEPLOYMENT_TYPE_DOCKER_RUNTIME` bundle exist on alpha, and what is its
   UUID?** (§3.1. Now a lookup rather than a publishing request — ask this first, it is the
   cheapest question that can unblock Phase B.) If not: where is the zrun
   `IMAGE_TYPE_DOCKER_RUNTIME` image, and can the lab node pull it?
2. ~~What goes in the runtime's cloud-init template?~~ **Answered** — see §1.4b. Remaining
   sub-question: are `disk-partition.sh` / `setup_dcompose.sh` / `last_things.sh` and the
   `zrun-agent` unit a stable contract of the published runtime image, or specific to the build
   this customer used? Worth one confirmation, since the Phase B cloud-init depends on them.
3. Can a switch network instance share the sole management port on a SLIRP node, or is the
   second NIC mandatory? (§1.4(e) says assume mandatory; worth one confirmation from the EVE
   side before budgeting §6.4 Phase 2.)
4. Why do the four `TestEnterprise_*` tests fail under CI credentials? Confirm rather than
   assume before writing the `ZEDCLOUD_ACC_ELEVATED` skip message.
5. Does `create_compose.tf` pass today against alpha? (§5 — cheap, and it may change the plan.)
