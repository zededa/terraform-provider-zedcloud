# UE-203: interface ACL test matrix (local cluster)

Checks the UE-203 provider changes (`ue-203-app-instance-acls`) against a
controller with zedcloud #8548 (ENG-3000). No EVE node: this checks what the
controller stores and what the provider plans.

Design: `docs/design/eng-3000-app-instance-acl-defaults.md`.

The app bundle's eth0 rules are allow 0.0.0.0/0, 9022->22, 9080->80. Each
`sN.tfvars` is one step; `./run.sh sN` plans, applies, plans again and dumps
the stored ACLs (`acls.sh`). `api_edit.py` edits ACLs through the API the way
Cortex does.

```sh
cd v2 && CGO_ENABLED=0 go build -o ../test/eng3000-repro/.bin/terraform-provider-zedcloud . && cd -
cd test/ue-203-acls   # repro.tfrc: dev_overrides -> that binary (absolute path, edit for your user)
./run.sh s1 ...       # uses TF_VAR_zedcloud_url / TF_VAR_zedcloud_token
TF_CLI_CONFIG_FILE=$PWD/repro.tfrc tofu destroy -var-file=s6.tfvars
```

## Results, 2026-10-27 (local cluster restarted with #8548)

Instances: a, b, c on a project with interface ordering **on**; d on one with it **off**.

| Step | Change | Plan | Stored after apply | Plan after |
|---|---|---|---|---|
| s1 | create: a = allow, 8022->22, 9080->80; b = acls omitted; c = bundle copy; d (off) = allow, 8022->22 | create | a as written; b and c = bundle rules once; d = 8022->22 + bundle rules appended (ordering off, unchanged behaviour) | no instance changes |
| s2 | a: 8022 -> 7022 | one ACL change on a (`value = "8022" -> "7022"`) | a: allow, 7022->22, 9080->80 | clean |
| s3 | b: `vminfo.cpus` 1 -> 2, acls still omitted | b cpus only | b ACLs unchanged (bundle defaults) | see note 2 |
| s4 | a: `acls` block removed; d (off): 8022 -> 7022 | no ACL change for a or d | a keeps 7022 rules; d unchanged | no ACL changes |
| s5 | b edited through the API (allow, 6022->22), then plan with b's acls omitted | no change for b | b keeps the API edit | — |
| s5 | b lists allow, 9022->22, 9080->80 | ACL change back to config | b as written | clean |
| s6 | `async_writes = false` (v1 routes); c drops 9022->22 | one ACL change on c | c: allow, 9080->80 | clean |

No step printed the "Interface ACLs were not stored as configured" warning.

Not covered here:

- Instances stored with merged lists before #8548: the cluster restart cleared
  the database, and a controller with #8548 no longer creates such lists on
  ordering-enabled projects. Covered by unit tests only.
- The read-back warning against a controller without #8548: unit tests only.

Notes (unrelated to UE-203, seen while running):

1. `project_access_list` on `zedcloud_application`, `zedcloud_datastore` and
   `zedcloud_image` with two projects: the controller returns the ids in a
   varying order, so plans keep reordering them.
2. `vminfo.cpus` on the instance: the change from 1 to 2 is planned again after
   every apply.
3. `manifest_info` is left out of the config: its `details` change whenever
   the app changes, which shows up as instance drift.
