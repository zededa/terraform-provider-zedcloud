#!/bin/bash
# Local test script for registry check workflow

set -e

# Configuration
PROVIDER="zededa/zedcloud"
GITHUB_REPO="zededa/terraform-provider-zedcloud"
VERSION="2.6.0"
PLATFORM="${PLATFORM:-linux}"
ARCH="${ARCH:-amd64}"
CURL_OPTS="-sf --max-time 30 --retry 3 --retry-delay 2"
CURL_OPTS_REDIRECT="-sfL --max-time 30 --retry 3 --retry-delay 2"

# Derived values
PROVIDER_NAME=$(echo "$PROVIDER" | cut -d'/' -f2)
FILENAME="terraform-provider-${PROVIDER_NAME}_${VERSION}_${PLATFORM}_${ARCH}.zip"

echo "🔍 Testing $PLATFORM/$ARCH..."
echo "Provider: $PROVIDER"
echo "GitHub Repo: $GITHUB_REPO"
echo "Filename: $FILENAME"
echo ""

# Test 1: GitHub SHA256SUMS
echo "=== Test 1: Fetching GitHub SHA256SUMS ==="
SHASUMS_URL="https://github.com/${GITHUB_REPO}/releases/download/v${VERSION}/terraform-provider-${PROVIDER_NAME}_${VERSION}_SHA256SUMS"
echo "URL: $SHASUMS_URL"

GITHUB_SHA=$(curl $CURL_OPTS_REDIRECT "$SHASUMS_URL" | grep "$FILENAME" | awk '{print $1}')
if [ -z "$GITHUB_SHA" ]; then
    echo "❌ Failed to fetch GitHub SHA"
    exit 1
else
    echo "✅ GitHub SHA: $GITHUB_SHA"
fi
echo ""

# Test 2: Terraform Registry
echo "=== Test 2: Fetching Terraform Registry ==="
TF_URL="https://registry.terraform.io/v1/providers/${PROVIDER}/${VERSION}/download/${PLATFORM}/${ARCH}"
echo "URL: $TF_URL"

TF_RESP=$(curl $CURL_OPTS "$TF_URL" 2>/dev/null || echo "{}")
TF_SHA=$(echo "$TF_RESP" | jq -r '.shasum // "none"')
if [ "$TF_SHA" = "none" ]; then
    echo "❌ Failed to fetch Terraform SHA"
    exit 1
else
    echo "✅ Terraform SHA: $TF_SHA"
fi
echo ""

# Test 3: OpenTofu Registry
echo "=== Test 3: Fetching OpenTofu Registry ==="
TOFU_URL="https://registry.opentofu.org/v1/providers/${PROVIDER}/${VERSION}/download/${PLATFORM}/${ARCH}"
echo "URL: $TOFU_URL"

TOFU_RESP=$(curl $CURL_OPTS "$TOFU_URL" 2>/dev/null || echo "{}")
TOFU_SHA=$(echo "$TOFU_RESP" | jq -r '.shasum // "none"')
if [ "$TOFU_SHA" = "none" ]; then
    echo "❌ Failed to fetch OpenTofu SHA"
    exit 1
else
    echo "✅ OpenTofu SHA: $TOFU_SHA"
fi
echo ""

# Test 4: Compare checksums
echo "=== Test 4: Comparing Checksums ==="
echo "GitHub:    $GITHUB_SHA"
echo "Terraform: $TF_SHA"
echo "OpenTofu:  $TOFU_SHA"
echo ""

VALID=true
if [ "$GITHUB_SHA" != "$TF_SHA" ]; then
    echo "❌ Terraform SHA mismatch!"
    VALID=false
fi

if [ "$GITHUB_SHA" != "$TOFU_SHA" ]; then
    echo "❌ OpenTofu SHA mismatch!"
    VALID=false
fi

if [ "$VALID" = "true" ]; then
    echo "✅ All checksums match!"
    exit 0
else
    echo "❌ Checksum validation failed"
    exit 1
fi
