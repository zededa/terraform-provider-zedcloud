# Design: NFR-19 — Edge app logos (and other artifacts) via Terraform

**Ticket:** [NFR-19](https://zededa.atlassian.net/browse/NFR-19) — *Terraform provider: Set a logo for an Edge app*
**Customer:** Speedcast. They manage many tenants and today set the logo on every Edge app in every tenant by hand in the UI.
**Repos:** `zededa/terraform-provider-zedcloud` @ `d97d6b8c` (origin/main), `zededa/zedcloud` @ `f956feff70` (main)
**Date:** 2026-09-30
**Author:** Ivan Curachin
**Status:** Draft

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
| What the value holds | An **artifact id**, `<uuid>_<name>`, keyed by the artifact name. Not a URL or base64 | `zedcloud/tests/apiTest/tests/test_app_bundle.py:420-462`, `tests/apiTest/testData/manifest.dockerCompose4.json:15` |
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
        (zedcloud_artifact.app_logo.name) = zedcloud_artifact.app_logo.id
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
| **D. A `zedcloud_artifact` data source only** (look up existing artifacts). | Useful later for adopting logos uploaded in the UI, but it doesn't upload anything. Out of scope; see §8. |
| **E. Put the image inline as a data URL in `desc.logo`.** | The backend would accept it (no validation), but the UI expects an artifact id and would break. It would also make the manifest very large. Rejected. |

---

## 4. Interface / API Design

### 4.1 Controller API used (already exists)

| Operation | Route | Request | Response |
|---|---|---|---|
| CreateArtifact | `POST /api/v1/artifacts` | JSON `{"name":"myapp-logo.png"}` | `200`, `Artifact{id:"<uuid>_myapp-logo.png", name}` |
| UploadArtifact | `PUT /api/v1/artifacts/id/{id}/upload/chunked` | raw bytes; headers `Content-Type: application/octet-stream`, `Content-Range: bytes 0-<N-1>/<N>` | `202` |
| GetArtifactSignedUrl | `GET /api/v1/artifacts/id/{id}/url` | none | `Artifact{signedUrl, ttl}` |
| GetArtifactStream | `GET /api/v1/artifacts/id/{id}` | none | file bytes |
| QueryArtifacts | `GET /api/v1/artifacts` | none | list |
| DeleteArtifact | `DELETE /api/v1/artifacts/id/{id}` | none | none |

Constraints the client must respect:

- **Name:** 3–256 characters, `[a-zA-Z0-9][a-zA-Z0-9_.-]+` (swagger `Artifact.name`, checked by `zutils.ValidateName` in niles).
- **Content type:** the upload body **must** be `application/octet-stream` (`libs/hutils/httputils_octet_stream.go:87-100`). The published swagger doesn't say so; it declares the body as a JSON `string/byte`. Because of that, a go-swagger generated client can't be used as-is (§6.2 item 5).
- **Single chunk only:** send the whole file in one `PUT`. Multi-chunk uploads are broken today (§6.2 item 1). Logos and screenshots are small, so this costs nothing.
- **niles, not gilas, detects completion:** niles uploads to object storage once the bytes it has received reach `fileSize`. gilas returns `202` without waiting for niles's result (§6.2 item 2), so the client should confirm the artifact is readable before reporting success.
- **Metadata stripping:** niles strips EXIF/XMP from JPEG/PNG before storing (`srvs/niles/image_metadata.go`, ZEDCLOUD-2365). The stored bytes can therefore differ from the source file, so the provider must not compare hashes of the stored bytes against the source.

### 4.2 New provider resource: `zedcloud_artifact`

| Attribute | Type | Mode | Notes |
|---|---|---|---|
| `name` | string | Required, ForceNew | Validated against the regex above. Also a natural map key for `desc.logo`. |
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
  3. Call `UploadArtifact(id, bytes)` with `Content-Range: bytes 0-<N-1>/<N>`.
  4. Poll `GetArtifactSignedUrl(id)`, up to the create timeout (default 2m), until it succeeds.
  5. `d.SetId(id)`.
  6. If the upload or the poll fails, call `DeleteArtifact(id)` on a best-effort basis and return an error.
- **Read:** check `GetArtifactSignedUrl(id)` (or look the id up with `QueryArtifacts` if that turns out to be cheaper or more precise). On 404, clear the id so Terraform plans a re-create.
- **Delete:** `DeleteArtifact(id)`; treat 404 as success.

### 4.4 Provider-side client

This will be a small, hand-written package, `v2/client/artifact/`, with `Create`, `Upload`, `GetSignedURL`,
`Query` and `Delete`. It uses the provider's existing go-openapi transport and bearer auth
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
- Optional: add a plan-time validator on `desc.logo` values (id pattern or http(s) URL). It would be advisory only, since the backend accepts anything.

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
6. **Optional:** validate `desc.logo` / `desc.screenshotList` in `seine` `validateManifest`, the same way licenses already are (URL or `ValidateName` artifact id).

None of these change the API contract the provider depends on.

---

## 7. Testing Strategy

- **Unit (no network):**
  - Artifact client: the request has `Content-Type: application/octet-stream`, `Content-Range: bytes 0-<N-1>/<N>`, a raw body, and the correct path. Test with `httptest.Server`.
  - Schema: `name` validation, `ExactlyOneOf(source, content_base64)`, hashing of `content_base64` in the `StateFunc`.
  - `details.go` fixes: expand/flatten round-trip for `logo`, `screenshot_list`, `license_list`.
- **Acceptance (`TF_ACC=1`, `make test-run case=...`):**
  - `TestArtifact_Create`: upload a small PNG fixture, check that `id` has the form `<uuid>_<name>` and that the signed URL can be fetched.
  - `TestArtifact_Replace`: change `source_hash`, which should re-create the artifact with a new id and delete the old one.
  - `TestArtifact_Import`: import, then plan with `ignore_changes` or a matching hash, and expect no diff.
  - `TestApplication_CreateWithLogo`: create an artifact and an app with `desc.logo = { name = id }`. Read the app back and assert `manifest.desc.logo[name] == artifact.id`, mirroring the zedcloud API test `test_003a_add_drop_logo_artifact_to_app`.
  - Use the existing `__SUFFIX__` fixture tokens (`docs/design/ue-139-fixture-suffix-tokens.md`) so artifact names don't collide on the shared enterprise.
- **Manual:** after apply, open the app in the UI and check the logo renders on the details page and the marketplace card. This is the only check for §8 Q1.

---

## 8. Open Questions

1. **Which map key does the UI read?** Is it the artifact name, a fixed key such as `"logo"`, or the first entry? API tests use the artifact name as the key. Confirm against a real app by setting a logo in the UI and running `GET /v1/apps/id/{id}`, then inspecting `manifestJSON.desc.logo`. The answer decides the recommended key in the docs, and whether the provider should also offer a typed `logo_artifact_id` shortcut.
2. **Does the UI support more than one logo entry?** If not, should the provider warn when `len(desc.logo) > 1`?
3. **Garbage collection.** When the UI replaces a logo, does it delete the old artifact, or is it left orphaned? Terraform's destroy handles artifacts Terraform manages. Artifacts uploaded in the UI are out of scope.
4. **Adopting existing logos.** Speedcast already has logos uploaded by hand. Should we add a `zedcloud_artifact` data source (look up by name via `QueryArtifacts`), so they can reference existing artifacts without re-uploading? It's cheap to add and could be a follow-up.
5. **Existence check.** Is `GetArtifactSignedUrl` the right way to check an artifact exists, or does it return a URL even for a missing object? Verify on dev. Otherwise use `QueryArtifacts` or a `HEAD` on the stream.
6. **Controller fixes (§6.2): who owns them, and when?** Items 1–4 are small. Do they go with this work or into the storage team's backlog?
7. **App profiles.** `zedge_app_profile_service.swagger.json:2766` has the same `Details.logo`. Is there a profile resource in scope that should get the same documentation?
