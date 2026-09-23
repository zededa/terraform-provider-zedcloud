#!/bin/bash
# Compare the node's EdgeNodeCerts against its device cert, and try the exact
# check the controller performs.
#
# The controller's EdgeNodeCert verification requires every attest cert to be
# SIGNED BY the device cert. EVE's dom0 has no openssl, so pull the PEMs out and
# inspect them here.
set -u
cd "$(dirname "$0")"
OUT=${OUT:-/tmp/eve-certs}
mkdir -p "$OUT"

for f in /persist/certs/ecdh.cert.pem /persist/certs/attest.cert.pem /config/device.cert.pem; do
  n=$(basename "$f")
  bash eve.sh "cat $f" 2>/dev/null | sed -n '/BEGIN CERTIFICATE/,/END CERTIFICATE/p' > "$OUT/$n"
  printf '%-22s %s bytes\n' "$n" "$(wc -c < "$OUT/$n" | tr -d ' ')"
done

echo
for n in device.cert.pem ecdh.cert.pem attest.cert.pem; do
  echo "--- $n"
  openssl x509 -in "$OUT/$n" -noout -subject -issuer -dates 2>&1 | sed 's/^/    /'
done

echo
echo "=== does each EdgeNodeCert verify against the device cert? ==="
for n in ecdh.cert.pem attest.cert.pem; do
  # -partial_chain so a non-CA device cert can still act as the trust anchor,
  # which is what the controller's signature check effectively does.
  R=$(openssl verify -partial_chain -CAfile "$OUT/device.cert.pem" "$OUT/$n" 2>&1)
  printf '%-18s %s\n' "$n" "$R"
done
