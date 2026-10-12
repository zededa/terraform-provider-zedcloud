"""mitmproxy addon: UE-177 fault injection in front of a local controller.

The active fault is read from the file named by $UE177_FAULT_FILE on every
request, so it can be switched without restarting mitmdump:

  none         pass everything through
  slow_writes  hold the RESPONSE of every app / app-instance write for
               $UE177_DELAY seconds (default 35). The controller has already
               done the work, so a client that gives up first sees exactly
               ENG-3015: "context deadline exceeded" for a change that
               happened.
  html504      answer 30% of app / app-instance reads with a Cloudflare-style
               HTML 504 page (UE-177 item 5).
  slow_jobs    hold job-status responses for $UE177_DELAY seconds, so async
               writes outlive a short resource timeout (taint path).
"""
import asyncio
import os
import random
import re

from mitmproxy import http

FAULT_FILE = os.environ.get("UE177_FAULT_FILE", "/tmp/ue177-fault")
DELAY = float(os.environ.get("UE177_DELAY", "35"))

APP_WRITE = re.compile(r"^/api/v[12]/apps(/instances)?(/id/[^/?]+)?/?(\?.*)?$")
APP_READ = re.compile(r"^/api/v1/apps(/instances)?/(id|name)/[^/?]+$")
JOB_READ = re.compile(r"^/api/v1/jobs/id/")

HTML_504 = (b"<!DOCTYPE html><html><head><title>504 Gateway time-out</title>"
            b"<style>body{font-family:sans-serif}</style></head><body>"
            b"<h1>504 Gateway time-out</h1><p>The web server reported a gateway time-out error.</p>"
            b"<p>cloudflare</p></body></html>")


def fault() -> str:
    try:
        with open(FAULT_FILE) as f:
            return f.read().strip()
    except OSError:
        return "none"


def request(flow: http.HTTPFlow) -> None:
    if fault() == "html504" and flow.request.method == "GET" and APP_READ.match(flow.request.path):
        if random.random() < 0.3:
            flow.response = http.Response.make(504, HTML_504, {"Content-Type": "text/html"})


async def response(flow: http.HTTPFlow) -> None:
    f = fault()
    path, method = flow.request.path, flow.request.method
    if f == "slow_writes" and method in ("POST", "PUT", "DELETE") and APP_WRITE.match(path):
        await asyncio.sleep(DELAY)
    elif f == "slow_jobs" and method == "GET" and JOB_READ.match(path):
        await asyncio.sleep(DELAY)
