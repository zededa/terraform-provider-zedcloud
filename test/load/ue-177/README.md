# UE-177 load and fault-injection harness

Reproduces the CI-1004 / ENG-3015 conditions against a controller and compares the
registry provider **2.8.0** (`old`) with a build of this repository (`new`). Design and
results: [`docs/design/ue-177-high-parallelism-robustness.md`](../../../docs/design/ue-177-high-parallelism-robustness.md) §6.1.

The setup mirrors the ticket: N edge apps (VM) and N app instances (`activate = false`)
spread over two registered-only edge nodes, applied with `-parallelism=160`. The instances
are deliberately minimal (no `vminfo`, `logs`, `manifest_info`, `tags`), so the plan right
after apply shows any server-filled-field drift.

## Requirements

- `terraform`, `go`, `jq`, `curl`; `mitmproxy` for the fault scenarios (`brew install mitmproxy`)
- `TF_VAR_zedcloud_url` and `TF_VAR_zedcloud_token` for the target controller
  (for a local controller, `.vscode/set_api_token.sh` sets both)
- Internet access for `terraform init` of the 2.8.0 provider

## Run

```bash
cd test/load/ue-177
./setup.sh                      # builds the provider, renders new/ and old/, inits 2.8.0

./ab.sh old 200 120             # A/B: create, re-apply, drift plan, delete, teardown
./ab.sh new 200 120

./run-mitm.sh                   # mitmdump on 127.0.0.1:8666 in front of the controller
./faults.sh old 40              # slow writes, HTML 504s, out-of-band deletes
./faults.sh new 40              # ... plus slow jobs against a 10s timeouts.create
pkill -f "mitmdump --mode reverse"

./cleanup.sh                    # deletes leftover ue177* apps and instances via the API
```

`PAR=60 ./ab.sh ...` changes the parallelism. Each step prints its exit code, duration,
number of errors and the most frequent error. Everything is also appended to
`logs/results.tsv`, and each step's full Terraform output is in `logs/`.

## Faults (`faults.py`)

The active fault is read from `/tmp/ue177-fault` (or `$UE177_FAULT_FILE`) on every request.
`faults.sh` switches it between steps.

| Fault | Effect |
|---|---|
| `slow_writes` | Holds the response of every app or app-instance write for `$UE177_DELAY` s (35). The controller has already done the work, so a client that gives up first sees exactly ENG-3015. 2.8.0's fixed deadline is 30 s. |
| `html504` | Answers 30% of app and app-instance reads with a Cloudflare-style HTML 504 page. |
| `slow_jobs` | Holds job-status responses, so async writes outlive a short `timeouts.create` (taint path). |
| `none` | Pass-through. |

## Notes

- On a local minikube controller, the `istio-proxy` sidecar of the `znginx` pod (128 MiB)
  is OOM-killed by the request bursts at parallelism 160. Every request then fails with
  503 "upstream connect error". For the published comparison it was raised to 512 MiB:
  ```bash
  kubectl -n zedcloud patch deploy znginx --type merge \
    -p '{"spec":{"template":{"metadata":{"annotations":{"sidecar.istio.io/proxyMemoryLimit":"512Mi"}}}}}'
  ```
  To revert, remove the annotation:
  ```bash
  kubectl -n zedcloud patch deploy znginx --type json \
    -p '[{"op":"remove","path":"/spec/template/metadata/annotations/sidecar.istio.io~1proxyMemoryLimit"}]'
  ```
- When 2.8.0 is stuck on 404s (the bug it reproduces), its own `destroy` cannot finish. Run
  `./cleanup.sh <prefix>`, `terraform state rm` the apps and instances in `old/`, then destroy.
