#!/usr/bin/env bash
# Registers a provider version + implementation in Terrakube's private
# provider registry via the Terrakube API, pointing at artifacts already
# hosted somewhere reachable (e.g. Nexus raw repo, see scripts/nexus-upload.sh).
#
# Reference: https://docs.terrakube.io/user-guide/private-registry/using-providers
#
# Required env vars:
#   TERRAKUBE_API_URL  e.g. https://terrakube-api.example.com
#   TERRAKUBE_TOKEN    Terrakube personal access token (Bearer)
#   ORG_ID             Terrakube organization id
#   VERSION            Provider version, e.g. 7.0.0-SNAPSHOT-2e1dff4c
#   DOWNLOAD_BASE_URL  Base URL where dist artifacts are hosted, e.g.
#                      http://nexus.wibu.local/repository/terraform-providers/rancher2/7.0.0-SNAPSHOT-2e1dff4c
#   OS                 Target os, default: linux
#   ARCH               Target arch, default: amd64
#   GPG_KEY_ID         Long key id of the GPG key used to sign SHA256SUMS
#                      (gpg --list-secret-keys --keyid-format LONG). The
#                      SHA256SUMS.sig detached signature must already exist
#                      in dist/, e.g.:
#                        gpg --detach-sig --local-user <GPG_KEY_ID> \
#                          --output dist/<project>_<version>_SHA256SUMS.sig \
#                          dist/<project>_<version>_SHA256SUMS
#
# Optional:
#   SOURCE_NAME, SOURCE_URL, TRUST_SIGNATURE
#
# Requires: curl, jq, gpg
#
# Usage:
#   TERRAKUBE_API_URL=https://terrakube-api.example.com \
#   TERRAKUBE_TOKEN=*** ORG_ID=*** VERSION=7.0.0-SNAPSHOT-2e1dff4c \
#   DOWNLOAD_BASE_URL=http://nexus.wibu.local/repository/terraform-providers/rancher2/7.0.0-SNAPSHOT-2e1dff4c \
#   ./scripts/terrakube-register-provider.sh



set -euo pipefail

# Load local, gitignored config if present (see scripts/registry.env.example).
# shellcheck disable=SC1091
[ -f "$(dirname "${BASH_SOURCE[0]}")/registry.env" ] && source "$(dirname "${BASH_SOURCE[0]}")/registry.env"

: "${TERRAKUBE_API_URL:?TERRAKUBE_API_URL is required, e.g. https://terrakube-api.k8s.wibu.local}"
: "${TERRAKUBE_TOKEN:?TERRAKUBE_TOKEN is required}"
: "${ORG_ID:?ORG_ID is required}"
: "${VERSION:?VERSION is required, e.g. 7.0.0-SNAPSHOT-2e1dff4c}"
: "${DOWNLOAD_BASE_URL:?DOWNLOAD_BASE_URL is required (where dist artifacts are hosted)}"
: "${GPG_KEY_ID:?GPG_KEY_ID is required (long key id of the key used to sign SHA256SUMS, e.g. gpg --list-secret-keys --keyid-format LONG)}"

# Strip any accidental trailing slashes so we don't end up with double slashes in URLs.
ORG_ID="${ORG_ID%/}"

# Require an explicit scheme so curl doesn't guess wrong / silently hit the wrong protocol.
case "${TERRAKUBE_API_URL}" in
  http://*|https://*) ;;
  *) echo "TERRAKUBE_API_URL must include a scheme, e.g. https://${TERRAKUBE_API_URL}" >&2; exit 1 ;;
esac

command -v jq >/dev/null || { echo "jq is required" >&2; exit 1; }

PROJECT="terraform-provider-rancher2"
OS="${OS:-linux}"
ARCH="${ARCH:-amd64}"
API="${TERRAKUBE_API_URL%/}/api/v1"
AUTH_HEADER="Authorization: Bearer ${TERRAKUBE_TOKEN}"
CT_HEADER="Content-Type: application/vnd.api+json"

# Calls the API and validates the response is JSON before handing it to jq,
# so failures (HTML error pages, plain-text proxy/gateway errors, etc.) are
# reported clearly instead of tripping a confusing jq parse error.
api_call() {
  local resp status body
  resp="$(curl -sS -w '\n%{http_code}' "$@")"
  status="${resp##*$'\n'}"
  body="${resp%$'\n'*}"
  if [ "$status" -ge 400 ] || ! echo "$body" | jq -e . >/dev/null 2>&1; then
    echo "Request failed (HTTP ${status}). Raw response:" >&2
    echo "$body" >&2
    exit 1
  fi
  echo "$body"
}

ZIP_NAME="${PROJECT}_${VERSION}_${OS}_${ARCH}.zip"
SHASUMS_NAME="${PROJECT}_${VERSION}_SHA256SUMS"
SHASUMS_SIG_NAME="${SHASUMS_NAME}.sig"
DOWNLOAD_URL="${DOWNLOAD_BASE_URL%/}/${ZIP_NAME}"
SHASUMS_URL="${DOWNLOAD_BASE_URL%/}/${SHASUMS_NAME}"
SHASUMS_SIGNATURE_URL="${DOWNLOAD_BASE_URL%/}/${SHASUMS_SIG_NAME}"

DIST_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/dist"
SHASUMS_LOCAL="${DIST_DIR}/${SHASUMS_NAME}"
if [ ! -f "$SHASUMS_LOCAL" ]; then
  echo "Could not find ${SHASUMS_LOCAL} locally to read the checksum from." >&2
  exit 1
fi
SHASUM="$(awk -v f="$ZIP_NAME" '$2==f{print $1}' "$SHASUMS_LOCAL")"
if [ -z "$SHASUM" ]; then
  echo "Could not find checksum for ${ZIP_NAME} in ${SHASUMS_LOCAL}" >&2
  exit 1
fi

if [ ! -f "${DIST_DIR}/${SHASUMS_SIG_NAME}" ]; then
  echo "Could not find ${DIST_DIR}/${SHASUMS_SIG_NAME}" >&2
  echo "Sign it first: gpg --detach-sig --local-user ${GPG_KEY_ID} --output ${DIST_DIR}/${SHASUMS_SIG_NAME} ${SHASUMS_LOCAL}" >&2
  exit 1
fi

ASCII_ARMOR="$(gpg --armor --export "${GPG_KEY_ID}")"
if [ -z "$ASCII_ARMOR" ]; then
  echo "Could not export public key for ${GPG_KEY_ID}" >&2
  exit 1
fi

SOURCE_NAME="${SOURCE_NAME:-Internal}"
SOURCE_URL="${SOURCE_URL:-${DOWNLOAD_BASE_URL}}"
TRUST_SIGNATURE="${TRUST_SIGNATURE:-5.0,6.0}"

echo "==> Ensuring provider 'rancher2' exists in org ${ORG_ID}"
providers_resp="$(api_call "${API}/organization/${ORG_ID}/provider" -H "${AUTH_HEADER}")"
PROVIDER_ID="$(echo "$providers_resp" | jq -r '.data[]? | select(.attributes.name=="rancher2") | .id' | head -n1)"
if [ -n "$PROVIDER_ID" ]; then
  echo "    provider already exists: ${PROVIDER_ID}"
else
  provider_resp="$(api_call -X POST "${API}/organization/${ORG_ID}/provider" \
    -H "${AUTH_HEADER}" -H "${CT_HEADER}" \
    -d '{"data":{"type":"provider","attributes":{"name":"rancher2","description":"Rancher2 terraform provider"}}}')"
  PROVIDER_ID="$(echo "$provider_resp" | jq -r '.data.id')"
  echo "    provider id: ${PROVIDER_ID}"
fi

echo "==> Ensuring version ${VERSION} exists"
versions_resp="$(api_call "${API}/organization/${ORG_ID}/provider/${PROVIDER_ID}/version" -H "${AUTH_HEADER}")"
VERSION_ID="$(echo "$versions_resp" | jq -r --arg v "$VERSION" '.data[]? | select(.attributes.versionNumber==$v) | .id' | head -n1)"
if [ -n "$VERSION_ID" ]; then
  echo "    version already exists: ${VERSION_ID}"
else
  version_resp="$(api_call -X POST "${API}/organization/${ORG_ID}/provider/${PROVIDER_ID}/version" \
    -H "${AUTH_HEADER}" -H "${CT_HEADER}" \
    -d "{\"data\":{\"type\":\"version\",\"attributes\":{\"versionNumber\":\"${VERSION}\",\"protocols\":\"5.0,6.0\"}}}")"
  VERSION_ID="$(echo "$version_resp" | jq -r '.data.id')"
  echo "    version id: ${VERSION_ID}"
fi

echo "==> Creating implementation ${OS}/${ARCH}"
implementation_payload="$(jq -n \
  --arg os "$OS" --arg arch "$ARCH" --arg filename "$ZIP_NAME" \
  --arg downloadUrl "$DOWNLOAD_URL" --arg shasumsUrl "$SHASUMS_URL" \
  --arg shasumsSignatureUrl "$SHASUMS_SIGNATURE_URL" --arg shasum "$SHASUM" \
  --arg keyId "$GPG_KEY_ID" --arg asciiArmor "$ASCII_ARMOR" \
  --arg trustSignature "$TRUST_SIGNATURE" --arg source "$SOURCE_NAME" --arg sourceUrl "$SOURCE_URL" \
  '{data:{type:"implementation",attributes:{os:$os,arch:$arch,filename:$filename,downloadUrl:$downloadUrl,shasumsUrl:$shasumsUrl,shasumsSignatureUrl:$shasumsSignatureUrl,shasum:$shasum,keyId:$keyId,asciiArmor:$asciiArmor,trustSignature:$trustSignature,source:$source,sourceUrl:$sourceUrl}}}')"
implementation_resp="$(api_call -X POST "${API}/organization/${ORG_ID}/provider/${PROVIDER_ID}/version/${VERSION_ID}/implementation" \
  -H "${AUTH_HEADER}" -H "${CT_HEADER}" \
  -d "${implementation_payload}")"
echo "$implementation_resp" | jq .

echo
echo "Done. In your terraform code use:"
echo "  source  = \"<your-terrakube-registry-host>/${ORG_ID}/rancher2\""
echo "  version = \"${VERSION}\""
echo
echo "NOTE: no GPG signature was submitted (shasumsSignatureUrl/keyId/asciiArmor)."
echo "If Terraform enforces signature verification, 'terraform init' may reject"
echo "this provider. Sign the release with a GPG key and re-run with those"
echo "fields added if you hit trust/signature errors."
