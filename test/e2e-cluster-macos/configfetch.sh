#!/bin/bash
# Why is EVE not applying the config it is being sent?
#
# The tell is EVE's own "configStatus":"fail" while info POSTs succeed: the
# config response is SIGNED, and EVE verifies the signature against
# /config/root-certificate.pem (the installer's object_signing_ca). If that CA
# did not sign the controller's object-signing certificate, EVE discards every
# config it fetches and says nothing the controller can see.
set -u
cd "$(dirname "$0")"

echo "=== getconfig / signature failures, deduped ==="
bash eve.sh 'grep -ahoE "\"msg\":\"[^\"]{0,160}(getconfig|GetConfig|signature|verify|signed|certificate)[^\"]{0,160}\"" /persist/newlog/collect/current.device.log 2>/dev/null | sort -u | tail -25' 2>&1 \
  | grep -vi deprecation | cut -c1-240

echo
echo "=== does EVE have the controller signing cert? ==="
bash eve.sh 'ls -l /persist/certs/ /config/root-certificate.pem 2>&1' 2>&1 | grep -vi deprecation
