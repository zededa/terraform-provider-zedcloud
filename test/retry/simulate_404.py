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
