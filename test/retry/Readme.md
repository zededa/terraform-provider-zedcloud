# How to generate a 404 error using mitmproxy and test retry logic in terraform

## 1. Install mitmproxy

### macOS (with Homebrew):
```bash
brew install mitmproxy
```

### Ubuntu/Debian:
```bash
sudo apt update && sudo apt install mitmproxy
```

---

## 2. Add Fault Injection Using mitmproxy Scripting

```python
import re
from mitmproxy import http
from mitmproxy.http import Response

# Match /api/v1/devices/id/{uuid}
uuid_path_re = re.compile(r"^/api/v1/devices/id/[0-9a-fA-F-]{36}$")

# Match /api/v1/devices/name/{name} (1–250 non-slash chars)
name_path_re = re.compile(r"^/api/v1/devices/name/[^/]{1,250}$")

# Count matching GET requests
request_count = 0

def response(flow: http.HTTPFlow) -> None:
    global request_count

    if flow.request.method != "GET":
        return

    path = flow.request.path
    if uuid_path_re.match(path) or name_path_re.match(path):
        request_count += 1

        if request_count % 5 == 0:
            print(f"✔️ [ALLOW #{request_count}] {path}")
            return
        else:
            print(f"❌ [SIMULATED 404 #{request_count}] {path}")
            flow.response = Response.make(
                404,
                b'{"code":404,"message":"Not Found"}',
                {"Content-Type": "application/json"}
            )
```

---

## 3. Start mitmproxy as a Reverse Proxy Using Fault Injection Script

Run the following command from your terminal:
```bash
mitmproxy --mode reverse:https://zedcontrol.local.zededa.net --listen-port 8666 -s simulate_404.py
```

This will:
- Accept connections on `http://127.0.0.1:8666`
- Forward them to `https://zedcontrol.local.zededa.net`
- Intercept and log full request/response details
- Inject a `404` error for four consecutive requests to:
  - `/api/v1/devices/id/<your-device-id>`
  - `/api/v1/devices/name/<your-device-name>`
- Everty fifth request will be allowed through.

---

## 4. Test it with curl

Run the following command to test:
```bash
curl -v http://127.0.0.1:8666/api/v1/devices/id/<your-device-id> \
  -H "Authorization: Bearer <your-token>"
```

---

## 5. Run you terraform script to test the retry logic
```bash
terraform plan -out=out.plan
terraform apply out.plan
```
