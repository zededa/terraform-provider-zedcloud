"""Edit an instance's eth0 ACLs through the API, the way Cortex does (S5).

Usage: python3 api_edit.py NAME LPORT APP_PORT
Replaces the eth0 ACLs with: allow 0.0.0.0/0, LPORT->APP_PORT (no ids).
"""
import json
import os
import ssl
import sys
import urllib.request

name, lport, app_port = sys.argv[1], sys.argv[2], int(sys.argv[3])
base = "https://%s/api" % os.environ["TF_VAR_zedcloud_url"]
hdr = {"Authorization": "Bearer " + os.environ["TF_VAR_zedcloud_token"], "Content-Type": "application/json"}
ctx = ssl._create_unverified_context()  # local cluster, private CA


def call(method, path, body=None):
    data = json.dumps(body).encode() if body is not None else None
    req = urllib.request.Request(base + path, data=data, headers=hdr, method=method)
    with urllib.request.urlopen(req, context=ctx) as r:
        return json.load(r)


inst = call("GET", "/v1/apps/instances/name/" + name)
for i in inst["interfaces"]:
    if i["intfname"] == "eth0":
        i["acls"] = [
            {"matches": [{"type": "ip", "value": "0.0.0.0/0"}]},
            {"matches": [{"type": "protocol", "value": "tcp"}, {"type": "lport", "value": lport},
                         {"type": "ip", "value": "0.0.0.0/0"}],
             "actions": [{"portmap": True, "mapparams": {"port": app_port}}]},
        ]
resp = call("PUT", "/v1/apps/instances/id/" + inst["id"], inst)
print("PUT", name, "->", resp.get("objectId") or resp)
