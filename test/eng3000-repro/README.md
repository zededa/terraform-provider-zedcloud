# ENG-3000 repro: bundle ACLs vs instance ACLs (Terraform side)

Ticket: https://zededa.atlassian.net/browse/ENG-3000 (comment 106827).
Run on 2026-10-26 against the local zedcloud (zedcontrol.local.zededa.net,
zedcloud 5d8e5eaa9) with the provider built from this checkout.

No EVE node needed: this checks what the controller stores. The bundle's eth0
rules are the ticket's: allow 0.0.0.0/0, 9022->22, 9080->80.

| Case | Instance ACLs in TF | Stored by controller | Plan after |
|---|---|---|---|
| a: edits a bundle rule | allow, 8022->22, 9080->80 | 8022->22, allow, 9022->22, 9080->80 | No changes |
| b: acls omitted | none | allow, 9022->22, 9080->80 | No changes |
| c: echoes the bundle | allow, 9022->22, 9080->80 | allow, 9022->22, 9080->80 | No changes |

Findings:

1. Case a reproduces ENG-3000: both 8022 and 9022 forward to 22 (see
   `ai_a.json`, a direct API GET). seine always re-appends manifest ACEs
   (`instanceOwnedACLs` / `applyMatchedACLReplaces`).
2. The provider never shows the drift and never sends ACL edits:
   `interfaces` has `DiffSuppressFunc: suppressAppInterfacesChangeAfterCreation`
   (v2/schemas/schema_helpers.go, from CI-201, 70c2ab39). Changing 8022 to
   7022 in config still planned "No changes" (`plan2.log`).

## Run

```sh
cd v2 && CGO_ENABLED=0 go build -o ../test/eng3000-repro/.bin/terraform-provider-zedcloud . && cd -
export TF_CLI_CONFIG_FILE=$PWD/repro.tfrc   # dev_overrides -> .bin (absolute path, edit for your user)
tofu apply      # uses TF_VAR_zedcloud_url / TF_VAR_zedcloud_token
tofu destroy
```

`terraform.tfstate` and `terraform.log` are gitignored and stay local only.
