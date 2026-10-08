# Design: NFR-19 — Edge app logos (and other artifacts) via Terraform

**Ticket:** [NFR-19](https://zededa.atlassian.net/browse/NFR-19) — *Terraform provider: Set a logo for an Edge app*
**Customer:** Speedcast. They manage many tenants and today set the logo on every Edge app in every tenant by hand in the UI.
**Repos:** `zededa/terraform-provider-zedcloud` @ `d97d6b8c` (origin/main), `zededa/zedcloud` @ `f956feff70` (main),
`zededa/zededa-services` (UI, `zedui-dev/`); UI references are relative to `zedui-dev/src/`
**Date:** 2026-09-30 (revised 2026-10-06 with UI findings, §4.6, and controller checks, §4.7)
**Author:** Ivan Curachin
**Status:** Implemented on `nfr-19-edge-app-logo` ([UE-176](https://zededa.atlassian.net/browse/UE-176)),
2026-10-07. See §9 for where the implementation differs from this design. Controller fixes remain
open in [ENG-3009](https://zededa.atlassian.net/browse/ENG-3009).

**Headline: the provider needs new work, and zedcloud needs no new endpoint.** The controller
already has a public, documented upload API (the ArtifactManager API, `/v1/artifacts`). The
Edge app manifest already has a logo field (`manifestJSON.desc.logo`), and the provider already
exposes that field (`zedcloud_application.manifest.desc.logo`). The gap is that the provider has
no way to upload a file and get back an artifact id to put in that field. Separately, zedcloud has
a few bugs and weak spots in the upload path. We should fix them, but they don't block this work (§6.2).

---

## 1. Problem Statement

A published Edge app shows a logo in the UI (app details page, edit page, marketplace card).
The only way to set it today is to upload it by hand in the UI, once per app **per tenant**. For
a customer that provisions tenants and apps with Terraform, that one manual step is what keeps
the rollout from being fully automated.

The ticket's comments assumed there is no API for this. There is one: it is the generic artifact
upload API, which isn't named "logo", so a search of the swagger for "logo" doesn't find it (§4.1).

The customer's proposal is to *"provide a means to set a logo as a Terraform resource and
use that as an object for the edge app logo"*. That is the design below.

### 1.1 Verified ground truth

| Thing | Actually | Where |
|---|---|---|
| Logo field on an Edge app | `map<string,string> logo = 11` on `Details` (`manifestJSON.desc`) | `zedcloud/libs/zmsg/zcommon/manifest.proto:96` |
| What the value holds | An **artifact id**, `<uuid>_<name>`. Not a URL or base64 | `zedcloud/tests/apiTest/tests/test_app_bundle.py:420-462`, `tests/apiTest/testData/manifest.dockerCompose4.json:15` |
| Map key the UI writes | Always the fixed key **`"logo"`**: `desc.logo = { logo: <artifactId> }` | UI `features/marketplace/services/edge-apps.ts:1012` |
| Map key the UI reads | `"logo"` if present, otherwise the first non-empty value whose key isn't `"zfill"`. Only one logo is ever used, and **http(s) URLs are not rendered** | UI `edge-apps.ts:960-971` |
| Server-side validation of `desc.logo` | **None.** The value is stored as an opaque string (only `licenseList`/`agreementList` are validated) | `zedcloud/srvs/seine/appproc.go:377, 647, 775-830` |
| Upload API | `ArtifactManager` service, `/v1/artifacts`, public, in the published swagger, no internal marker | `zedcloud/libs/zmsg/zapiservices/zedge_storage_service.proto:838-1040` |
| Who serves it | gilas (HTTP handlers), which passes the file over gRPC to niles, which stores it in S3/Azure under `<enterpriseId>/<artifactId>` | `srvs/gilas/artifact_handlers.go:17-114`, `srvs/niles/artifact_handlers.go` |
| Access control | `ACCESS_TYPE_EDGE_APP`, the same permission needed to manage apps | `srvs/gilas/artifact_handlers.go:55, 82` |
| Provider schema | `zedcloud_application.manifest.desc.logo` is a `TypeMap` of strings and round-trips through create/read | `v2/schemas/details.go:210-217`, `:110-120`, `:153`, `:167` |
| Provider artifact client | **None in v2.** A legacy v1 generated client and resource exist but aren't registered, and would be rejected by gilas anyway (they send base64 JSON) | `client/artifact_manager/gen/`, `resources/gen/artifact_manager_resource.go` |
| Other resources with the same logo pattern | `SysModel.logo`, `SysBrand.logo`: both `map<string,string>` of artifact ids | `zedcloud/libs/zmsg/device/sysmodel.proto:111, 305`; `v2/schemas/sys_model.go:355`, `v2/schemas/sys_brand.go:201` |

So the only thing a user can do today is paste an artifact id, created some other way, into
`desc.logo`. That works, but it doesn't solve the problem.

---

## 2. Proposed Solution

Add a new, general-purpose **`zedcloud_artifact`** resource. It uploads a local file (or inline
base64 content) to the ArtifactManager API and exposes the resulting artifact `id`. Users put that
id into `zedcloud_application.manifest.desc.logo`, and the same resource also covers
screenshots, custom license uploads, brand logos and model logos.

```hcl
resource "zedcloud_artifact" "app_logo" {
  name        = "myapp-logo.png"
  source      = "${path.module}/assets/myapp-logo.png"
  source_hash = filesha256("${path.module}/assets/myapp-logo.png")

  lifecycle { create_before_destroy = true }
}

resource "zedcloud_application" "myapp" {
  name = "myapp"
  # ...
  manifest {
    # ...
    desc {
      app_category = "APP_CATEGORY_OPERATING_SYSTEM"
      logo = {
        logo = zedcloud_artifact.app_logo.id   # key must be "logo" (§4.6)
      }
    }
  }
}
```

Because the provider is configured per tenant and artifacts belong to the enterprise that
uploaded them, a module that pairs a `zedcloud_artifact` with a `zedcloud_application`
gives the customer one logo per tenant with no manual step.

---

## 3. Alternatives Considered

| Option | Why not |
|---|---|
| **A. Do nothing; document "paste an artifact id into `desc.logo`".** | Still needs a manual or scripted upload per tenant, which is exactly the problem the customer has. |
| **B. A `logo_file` convenience attribute on `zedcloud_application`** that uploads and fills `desc.logo` behind the scenes. | It ties the upload's lifecycle to the app. It can't be shared between apps, and it doesn't extend to screenshots, licenses, brands or models without adding a `*_file` attribute to each. It also makes `desc.logo` partly provider-managed, so the diff has to be suppressed to avoid a perpetual change. More code and less general than C. |
| **C. Standalone `zedcloud_artifact` resource (chosen).** | Matches the customer's proposal, can be referenced from anywhere, and each resource maps to one API object. Its lifecycle (create, upload, delete) is ordinary Terraform. |
| **D. A `zedcloud_artifact` data source only** (look up existing artifacts). | Doesn't upload anything. Looking artifacts up by name is also useless, because the UI names every logo artifact `"logo"` (§4.6). Rejected. |
| **E. Put the image inline as a data URL in `desc.logo`.** | The backend would accept it (no validation), but the UI expects an artifact id and would break. It would also make the manifest very large. Rejected. |

---

## 4. Interface / API Design

### 4.1 Controller API used (already exists)

| Operation | Route | Request | Response |
|---|---|---|---|
| CreateArtifact | `POST /api/v1/artifacts` | JSON `{"name":"myapp-logo.png"}` | `200`, `Artifact{id:"<uuid>_myapp-logo.png", name}` |
| UploadArtifact | `PUT /api/v1/artifacts/id/{id}/upload/chunked` | raw bytes; headers `Content-Type: application/octet-stream`, `Content-Range: bytes 0-<N>/<N>` (exclusive end, see below) | `202` |
| GetArtifactSignedUrl | `GET /api/v1/artifacts/id/{id}/url` | none | `200` `Artifact{signedUrl, ttl}`; **`307` when not found** (§4.7) |
| GetArtifactStream | `GET /api/v1/artifacts/id/{id}` | none | `200` file bytes (`Content-Type: application/x-proto-binary`); `307` when not found |
| QueryArtifacts | `GET /api/v1/artifacts` | none | list. **Unusable: times out with a 504 at nginx** (§4.7). Not used |
| DeleteArtifact | `DELETE /api/v1/artifacts/id/{id}` | none | `200` always, including for missing or already-deleted ids |

Constraints the client must respect:

- **Name:** 3–256 characters, `[a-zA-Z0-9][a-zA-Z0-9_.-]+` (swagger `Artifact.name`, checked by `zutils.ValidateName` in niles).
- **Content type:** the upload body **must** be `application/octet-stream` (`libs/hutils/httputils_octet_stream.go:87-100`). The published swagger doesn't say so; it declares the body as a JSON `string/byte`. Because of that, a go-swagger generated client can't be used as-is (§6.2 item 5).
- **Content-Range end is exclusive:** the UI sends `bytes <start>-<start+len>/<total>`, not the RFC 7233 inclusive `end-1` form (UI `lib/api/base.ts:1329-1345`). On the controller, both forms worked (§4.7), because gilas only logs the end value. The provider uses the exclusive form, as the UI does.
- **Single chunk only:** send the whole file in one `PUT`. Multi-chunk uploads are broken today (§6.2 item 1). Logos and screenshots are small, so this costs nothing.
- **niles, not gilas, detects completion:** niles uploads to object storage once the bytes it has received reach `fileSize`. gilas returns `202` without waiting for niles's result (§6.2 item 2), so the client should confirm the artifact is readable before reporting success.
- **Metadata stripping:** niles strips EXIF/XMP from JPEG/PNG before storing (`srvs/niles/image_metadata.go`, ZEDCLOUD-2365). This **re-encodes** the image: a 70-byte PNG came back as 74 bytes (§4.7). The provider must never compare stored bytes or hashes against the source.
- **Not-found is `307`, not `404`:** gilas returns `307 Temporary Redirect` with no `Location` header for a missing artifact (`srvs/gilas/artifact_handlers.go:191, 258`). The client must map 307 to "not found". Go's `net/http` returns a 3xx response that has no `Location` as-is, so it won't try to follow it.

### 4.2 New provider resource: `zedcloud_artifact`

| Attribute | Type | Mode | Notes |
|---|---|---|---|
| `name` | string | Required, ForceNew | Validated against the regex above. Only becomes the id suffix; it is **not** the map key (the key is always `"logo"`, §4.6). The UI uses `"logo"`, `"license"` or `"agreement"`, and any valid name works. |
| `source` | string | Optional, ForceNew | Path to a local file. `ExactlyOneOf` with `content_base64`. |
| `content_base64` | string | Optional, ForceNew | For generated content or `filebase64()`. Store only a hash in state, not the content itself (use a `StateFunc`), so state stays small. |
| `source_hash` | string | Optional, ForceNew | Recommended with `source`, set to `filesha256(...)`. Changing it re-uploads. Same pattern as `aws_s3_object.source_hash`. |
| `size` | int | Computed | Bytes uploaded. |
| `id` | string | Computed | Artifact id, `<uuid>_<name>`. This is what goes into `desc.logo` / `screenshotList` / brand/model `logo`. |

Deliberately **not** exposed: `signed_url`. It expires after the controller's TTL, so storing
it would produce a diff on every plan. If there's demand, it can be a data source attribute
with no diff.

No `Update` function: artifacts can't be changed after upload, so every input is `ForceNew`.
The docs will recommend `create_before_destroy` so that when a logo is replaced, the app is
pointed at the new artifact before the old one is deleted. (SDK v2 can't force this from inside the provider.)

**Import:** `terraform import zedcloud_artifact.x <id>`. On import, the provider sets `name`
from the id suffix. `source`/`source_hash` stay unset, so the user must add them to config.
After that, `ignore_changes` or a matching hash is needed, or the next apply re-uploads.
This limitation will be documented.

### 4.3 CRUD behaviour

- **Create:**
  1. Read the bytes from `source` or `content_base64`.
  2. Call `CreateArtifact(name)` to get an `id`.
  3. Call `UploadArtifact(id, bytes)` with `Content-Range: bytes 0-<N>/<N>`.
  4. Poll `GetArtifactSignedUrl(id)` every 1s until it returns `200`, bounded by the create timeout (`timeouts { create }`, default 1 minute; see §9). It took about 2s on the controller. A `307` means the artifact is either not stored yet or **failed silently**: the controller doesn't let us tell the two apart (§6.2 item 2), so hitting the timeout is the only failure signal.
  5. `d.SetId(id)`.
  6. If the upload or the poll fails, call `DeleteArtifact(id)` on a best-effort basis and return an error.
- **Read:** `GetArtifactSignedUrl(id)`. `200` means the artifact exists. `307` (or `404`, in case ENG-3009 changes it) means it's gone, so clear the id and Terraform plans a re-create. Treat any other status as an error. Don't use `QueryArtifacts`, because it times out.
- **Delete:** `DeleteArtifact(id)`. It always returns 200, so anything that isn't 2xx is an error. Since a successful DELETE says nothing about whether the artifact existed, no read-back is needed.

### 4.4 Provider-side client

This will be a small, hand-written package, `v2/client/artifact/`, with `Create`, `Upload`, `GetSignedURL`
and `Delete`. `GetSignedURL` returns a typed `ErrNotFound` for 307 and 404. It uses the provider's existing go-openapi transport and bearer auth
(`v2/resources/provider.go:154-177`). The upload op sets `ConsumesMediaTypes: []string{"application/octet-stream"}`
and passes `[]byte`/`io.Reader` as the body; go-openapi's built-in `ByteStreamProducer` handles it.
The package gets wired into `v2/client/zedcloud_api.go` like the other sub-clients.

The retry client retries only GETs (`provider.go`). That suits us: `Create` and `Upload` must
not be retried blindly, because a retry would create an orphan artifact.

This is hand-written rather than regenerated because the published swagger has the wrong content
type (§4.1), and v2 clients are already hand-maintained. The precedent for a hand-written facade is
`v2/schemas/enterprise_white_labeling.go` (NFR-163).

### 4.5 Changes to existing schema

- `v2/schemas/details.go:210`: improve the `logo` description (and those of `screenshot_list` and
  `license_list`). Say that the value is an artifact id, give the recommended key, and link to `zedcloud_artifact`.
- `v2/schemas/details.go`: fix the latent bugs found during analysis:
  - `DetailsModel(d)` calls `GetOk` with camelCase keys (`"agreementList"`, `"licenseList"`, `"screenshotList"`). This is harmless today because only the `FromMap` path is used.
  - `DetailsModelFromMap` does unchecked `m["category"].(string)` / `m["os"].(string)` type assertions.
- Add a plan-time validator on `desc.logo`, because the UI behaviour in §4.6 makes these mistakes silent (the logo just doesn't render):
  - error if a value looks like an http(s) or `data:` URL;
  - warn if there is more than one entry, or if the only key isn't `"logo"`.

### 4.6 What the UI does (verified in `zededa-services/zedui-dev`)

| Aspect | UI behaviour | Consequence for the provider |
|---|---|---|
| Writing the logo | `POST /artifacts {name:"logo"}`, then one chunked PUT, then read-modify-write of the app with `desc.logo = {logo: id}`. Removing the logo sets `desc.logo = null` (`edge-apps.ts:918-937, 994-1030`) | The documented and recommended form is `logo = { logo = <id> }` |
| Reading the logo | Uses the `"logo"` key, otherwise the first non-empty value (except `"zfill"`). Never treats the value as a URL (`edge-apps.ts:960-983`) | Validator in §4.5. Values that are URLs won't render |
| Showing the image | `GET /artifacts/id/{id}/url` → `<img src=signedUrl>`. On failure it falls back to inline `desc.svg` or a generated tile (`MarketplaceItemVisual.tsx:82-93, 193-237`) | A missing artifact is a cosmetic failure, not an error |
| Upload limits | PNG/JPEG only, at most 5 MB (`CloudEdgeAppDetails.tsx:1084-1093, 1381`). The UI's colour sampling assumes PNG or JPEG | Document the limits. The generic resource doesn't enforce them, but the logo example uses PNG |
| Chunking | 50 MB chunks, exclusive Content-Range end, `Content-Type: application/octet-stream`, plus an extra `data: {"name":"logo"}` header that the server ignores (`lib/api/base.ts:1299-1375`) | Logos are always one chunk, so the multi-chunk bug never affects the UI either |
| Upload completion | No polling. The UI saves the app as soon as the PUT returns 2xx | The provider's poll (§4.3) is stricter than the UI. Keep it, with a short timeout |
| Old artifacts | **Never deleted.** Not on replace, not on remove, not on app delete. The UI never calls `DELETE /artifacts` | Orphans already pile up from UI use. Deleting on Terraform destroy is an improvement |
| App duplicate | The copy shares the same logo artifact id (`CloudEdgeAppDetails.tsx:980-985`) | **Destroying a Terraform-managed artifact can break the logo of an app someone duplicated in the UI.** Document this, and consider an opt-out `retain_on_destroy` attribute (§8) |
| Create wizard | No logo step: `logo: null` (`EdgeAppCreationWizard.tsx:1261`) | Terraform will be the only way to set a logo when the app is created |
| Other logos | Brand and model logos are read the same way (`"logo"` key, signed URL), with no upload UI. The enterprise white-label logo (`$ztag.entp.zui.ux.logo`) holds an artifact id **or** an http(s) URL, and the UI handles both (`admin.ts:1288`, `useTenantBrandingHydration.ts:22-26`). Limits there: under 1 MB; PNG, JPEG, SVG or WebP | `zedcloud_artifact` ids also work for `zedcloud_brand`/`zedcloud_model` `logo = { logo = id }` and for `zedcloud_enterprise.white_labeling.logo_url` (an id is about 45 chars, which fits the 3–256 limit). The enterprise UI also uploads a separate monochrome variant (`...logo.mono`), which is out of scope |
| licenseList | Keys are `CUSTOM_UPLOAD`, `CUSTOM_UPLOAD_2`, …; values are artifact ids from `POST {name:"license"}`. The UI shows a label and never resolves the artifact | `zedcloud_artifact` covers custom license uploads too |
| screenshotList | Not supported by the UI | Out of scope |

---

### 4.7 Verified against a controller (local cluster, 2026-10-06)

Run with `curl` against `zedcontrol.local.zededa.net`, using a 70-byte 1×1 PNG. All test artifacts were deleted afterwards.

**Upload**

| Case | Result |
|---|---|
| `Content-Range: bytes 0-70/70` (exclusive end, UI form) | `202`; `/url` returns `200` after ~2s |
| `Content-Range: bytes 0-69/70` (inclusive end) | `202`; `/url` returns `200` after ~2s. gilas ignores the end value |
| `Content-Type: application/json` (what a go-swagger client sends) | **`500`**, not 400/415 |
| Two chunks, `bytes 0-35/70` then `bytes 35-70/70` | Both `202`, but `/url` was **still `307` after 4 minutes**. Confirms §6.2 items 1 and 2 |
| Stored bytes compared with the source | Valid PNG with the same header, but IDAT re-compressed: **74 bytes, not 70**. Same result through the signed URL |

**Existence**

| Case | `GET /url` | `GET /id/{id}` (stream) |
|---|---|---|
| Uploaded | `200`, and the signed S3 URL fetches with `200` | `200` |
| Never existed | `307` | `307` |
| Created, never uploaded | `307` | — |
| Uploaded, then deleted | `307` | `307` |

| Other | Result |
|---|---|
| `DELETE` an existing id | `200` |
| `DELETE` an already-deleted id, or one that never existed | `200` |
| `GET /artifacts` (list) | No response; nginx returned **`504` after 120s**, on both attempts |

---

## 5. Affected Components

### 5.1 `terraform-provider-zedcloud` (required)

| Change | Files |
|---|---|
| Artifact client | `v2/client/artifact/*.go` (new), `v2/client/zedcloud_api.go` |
| Resource | `v2/resources/artifact.go` (new), `v2/resources/provider.go` (`ResourcesMap`) |
| Schema | `v2/schemas/artifact.go` (new), `v2/schemas/details.go` (descriptions and bug fixes) |
| Tests | `v2/resources/artifact_test.go`, `v2/resources/application_test.go`, `v2/resources/testdata/artifact/`, `v2/resources/testdata/application/create_with_logo.{tf,yaml}` plus a small PNG fixture, `v2/client/artifact/*_test.go` |
| Docs | `v2/examples/resources/zedcloud_artifact/resource.tf`, then run `make docs` (`docs/resources/artifact.md`, refreshed `docs/resources/application.md`) |

Rough size: **3–5 days**, including acceptance tests against a dev enterprise.

### 5.2 `zedcloud` (recommended, not blocking)

See §6.2. All of these are small and could ship in one PR owned by the storage (gilas/niles) team.

---

## 6. Security Considerations

### 6.1 Provider

- **Auth:** the provider reuses the existing bearer token. The token needs the same Edge-app permission it already needs for `zedcloud_application`. There is no new credential handling.
- **Local file access:** the provider reads only the path the user gives in `source`, the same trust model as `manifest_file` (`v2/resources/application.go:331`).
- **State:** stores the id, name, hash and size. It never stores the file content (`content_base64` is hashed by the `StateFunc`).
- **Tenant isolation:** the controller stores artifacts under `<enterpriseId>/<artifactId>` and applies the caller's scope. The provider adds nothing on top.
- **PII:** niles already strips image metadata, and the provider relies on that.

### 6.2 Controller issues found during analysis

1. **Multi-chunk uploads are broken.** `sendArtifactChunk` never sets `StartRange` on the gRPC request (`srvs/gilas/artifact_helpers.go:116-117`). As a result, niles sees `StartRange == 0` on every chunk and deletes the partial file (`srvs/niles/artifact_handlers.go:120`). Fix: add `StartRange: startRange`. The provider avoids the bug by uploading in one chunk.
2. **The upload returns 202 without knowing whether it worked.** gilas calls `stream.CloseSend()` and never calls `CloseAndRecv()`, so errors from niles are lost: EXIF stripping failing, or the S3/Azure upload failing. Fix: receive the response and map failures to 4xx/5xx. Until then the provider polls (§4.3).
3. **A malformed `Content-Range` header crashes the handler.** `readOctetStreamLength` indexes split results without checking them (`libs/hutils/httputils_octet_stream.go:65-79`). A missing or bad header causes an index-out-of-range panic. Fix: validate the header and return 400.
4. **No upload size limit.** `ReadOctetStream` calls `readOctetStream(..., 0)`, which means no limit, and the body is read fully into memory. Fix: set a limit for artifacts (for example 10 MB) and return 413.
5. **The swagger is wrong for the upload.** The upload op in `zedge_storage_service.proto` declares no `consumes`, so generated clients send JSON/base64 and gilas rejects them. Fix: add `consumes: "application/octet-stream"` to the openapiv2 operation option and regenerate the swagger.
6. **Optional:** validate `desc.logo` in `seine` `validateManifest`. Accept only `ValidateName`-valid artifact ids, and **reject URLs**, because the UI can't render a URL logo (§4.6). (The UI notes that some backends already reject the legacy `zfill` entry with "Name field contains invalid characters", so a check like this may already exist in some form. Confirm before adding one.)

7. **The JSON upload body returns `500`.** A wrong `Content-Type` should be a `415` (or `400`) with a message, not an internal error (§4.7).
8. **Not-found is `307 Temporary Redirect`** with no `Location` header, on both `/url` and the stream (`srvs/gilas/artifact_handlers.go:191, 258`). It should be `404`. The provider accepts both, so this change is safe for it, but check UI callers first (the UI treats any failure as "no logo").
9. **`GET /artifacts` (list) times out:** a `504` from nginx after 120s on the local cluster. It's probably an unbounded scan of the bucket or tenant prefix. Needs pagination or a bound, or the endpoint should be removed.

None of these change the API contract the provider depends on. Changing 307 to 404 is the only one the provider notices, and it already handles both.

---

## 7. Testing Strategy

- **Unit (no network):**
  - Artifact client: the request has `Content-Type: application/octet-stream`, `Content-Range: bytes 0-<N>/<N>`, a raw body, and the correct path. `GetSignedURL` maps `307`/`404` to `ErrNotFound`, and other non-2xx responses to errors. `Delete` accepts `200`. Test with `httptest.Server`.
  - Schema: `name` validation, `ExactlyOneOf(source, content_base64)`, hashing of `content_base64` in the `StateFunc`.
  - `details.go` fixes: expand/flatten round-trip for `logo`, `screenshot_list`, `license_list`.
- **Acceptance (`TF_ACC=1`, `make test-run case=...`):**
  - `TestArtifact_Create`: upload a small PNG fixture, check that `id` has the form `<uuid>_<name>` and that the signed URL can be fetched.
  - `TestArtifact_Replace`: change `source_hash`, which should re-create the artifact with a new id and delete the old one.
  - `TestArtifact_Import`: import, then plan with `ignore_changes` or a matching hash, and expect no diff.
  - `TestApplication_CreateWithLogo`: create an artifact and an app with `desc.logo = { logo = id }`. Read the app back and assert `manifest.desc.logo["logo"] == artifact.id`, mirroring the zedcloud API test `test_003a_add_drop_logo_artifact_to_app`.
  - Use the existing `__SUFFIX__` fixture tokens (`docs/design/ue-139-fixture-suffix-tokens.md`) so artifact names don't collide on the shared enterprise.
- **Manual:** after apply, open the app in the UI and check the logo renders on the details page and the marketplace card.

---

## 8. Open Questions

### Resolved (UI code and controller checks, 2026-10-06)

- ~~Which map key does the UI read?~~ **`"logo"`**, with a fallback to the first value (§4.6). Because the key is fixed, a typed `logo_artifact_id` shortcut is not worth adding; the validator in §4.5 is enough.
- ~~More than one logo entry?~~ **No.** Only one entry is used, so the provider warns when there is more than one.
- ~~Does the UI clean up old artifacts?~~ **No, never.** Orphans already pile up, and Terraform destroy is an improvement.
- ~~App profiles?~~ **No logos** in the profile UI or on app-instance pages. Out of scope.
- ~~Existence check?~~ Settled on the controller (§4.7): use `GET /url`. `200` means it exists; `307` means missing, deleted, never uploaded, or a failed upload. The list endpoint is unusable.
- ~~Content-Range form?~~ Both forms work, and the provider uses the exclusive one the UI sends (§4.7).

### Still open

1. ~~**Shared artifacts and destroy.**~~ Decided: added `retain_on_destroy` (default false). See §9.
2. ~~**Restrict content for logos?**~~ Decided: the docs describe the UI's limits (PNG/JPEG, up to 5 MB), and the provider enforces only a hard 10 MB cap. It does not check file types, because the resource is generic. See §9.
3. **Adopting existing logos.** Speedcast already has logos uploaded by hand. Looking them up by name is useless, because the UI names every logo artifact `"logo"`. The simplest way to adopt one: read the id from the app (`manifest.desc.logo.logo` on the existing `zedcloud_application` resource or data source), or `terraform import zedcloud_artifact.x <id>`. A dedicated data source is probably unnecessary.
4. **Controller fixes (§6.2): who owns them, and when?** They are tracked in [ENG-3009](https://zededa.atlassian.net/browse/ENG-3009), currently unassigned. Provider work is tracked in [UE-176](https://zededa.atlassian.net/browse/UE-176).

---

## 9. Implementation notes (UE-176, 2026-10-07)

This section lists what was built and where it differs from §2–§7.

**Files**

| Area | Files |
|---|---|
| Client | `v2/client/artifact/client.go` (+ `client_test.go`), wired into `v2/client/zedcloud_api.go` |
| Resource | `v2/resources/artifact.go`, registered in `v2/resources/provider.go` |
| Schema | `v2/schemas/artifact.go`, `v2/schemas/details_logo.go` (+ test), `v2/schemas/details.go` |
| Tests | `v2/resources/artifact_test.go`; fixtures in `v2/resources/testdata/artifact/` and `testdata/application/{create_with_logo,logo_url_rejected}.tf` |
| Docs | `docs/resources/artifact.md` (generated), `v2/examples/resources/zedcloud_artifact/{resource.tf,import.sh}` |

**Differences from the design**

- **Create timeout:** the default is 1 minute (via `timeouts { create }`), not 30s. The poll interval is 1s, and on the controller the artifact appeared after about 2s.
- **`retain_on_destroy`:** added (resolves §8 Q1). The resource therefore has an `Update` function, which only touches state, so that this flag can change without replacing the artifact.
- **Size cap:** content must be 1 byte to 10 MiB (`zschema.ArtifactMaxSize`), checked before calling the API. There is no file-type check (resolves §8 Q2).
- **`GetSignedURL` treats a `200` with an empty `signedUrl` as not found.** It also maps the retrying HTTP client's "giving up on 404" transport error to `ErrNotFound`, because the provider's retry policy retries GET 404s.
- **`Delete` also accepts `404`**, so the client keeps working after ENG-3009 item 8 lands.
- **`desc.logo` validator:** it errors on http(s) and `data:` URLs, and warns when there is more than one entry or the only key isn't `logo`. Values not known at plan time (an artifact id before create) are skipped.
- **`details.go`:** besides the description changes, the map and scalar type assertions in `DetailsModelFromMap` no longer panic on missing or mistyped values. The camelCase `GetOk` keys in the unused `DetailsModel(d)` are fixed.
- **Fixtures:** new `__TESTDATA__` placeholder (the absolute path of `./testdata`), because acceptance tests run Terraform in a temp directory. It is registered in `v2/testing/fixtures_test.go`.

**Verified on the local controller (2026-10-07)**

| Test | Result |
|---|---|
| `TestArtifact_CRUD`: create from file, replace via `content_base64` (new id, old one deleted), import | PASS |
| `TestArtifact_RetainOnDestroy`: artifact still exists after destroy | PASS |
| `TestApplication_CreateWithLogo`: API returns `desc.logo == {logo: <artifact id>}` | PASS |
| `TestApplication_LogoURLRejected`: plan fails on a URL logo | PASS |
| `TestApplication_Create`, `TestApplication_Create_FromFile`: regression check after the `details.go` changes | PASS |
| Unit: client (incl. 307 without `Location`), validator, `waitForArtifact`, `InternalValidate` | PASS |

**Not done here**

- Rendering in the UI was not checked by eye: the acceptance test destroys its app. Apply the example and open the app in the console to confirm the logo shows on the details page and the marketplace card.
- The `zedcloud_enterprise` example still says `white_labeling.logo_url` must be a URL. The console also accepts an artifact id there (§4.6). Update it separately.
