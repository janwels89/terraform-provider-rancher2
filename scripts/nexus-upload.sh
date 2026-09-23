#!/usr/bin/env bash
# Uploads goreleaser dist/ artifacts for a provider version to a Nexus raw
# hosted repository, so they can be referenced by Terrakube's private
# provider registry (see: https://docs.terrakube.io/user-guide/private-registry/using-providers).
#
# Prerequisite: create a "raw (hosted)" repository in Nexus first
# (Administration -> Repositories -> Create repository -> raw (hosted)).
#
# Required env vars:
#   NEXUS_URL      Base URL of Nexus, e.g. http://nexus.wibu.local
#   NEXUS_REPO     Name of the raw hosted repository, e.g. terraform-providers
#   NEXUS_USER     Nexus username
#   NEXUS_PASSWORD Nexus password or API token
#   VERSION        Provider version to upload, e.g. 7.0.0-SNAPSHOT-2e1dff4c
#
# Usage:
#   NEXUS_URL=http://nexus.wibu.local NEXUS_REPO=terraform-providers \
#   NEXUS_USER=admin NEXUS_PASSWORD=*** VERSION=7.0.0-SNAPSHOT-2e1dff4c \
#   ./scripts/nexus-upload.sh

set -euo pipefail

# Load local, gitignored config if present (see scripts/registry.env.example).
# shellcheck disable=SC1091
[ -f "$(dirname "${BASH_SOURCE[0]}")/registry.env" ] && source "$(dirname "${BASH_SOURCE[0]}")/registry.env"

: "${NEXUS_URL:?NEXUS_URL is required, e.g. http://nexus.wibu.local}"
: "${NEXUS_REPO:?NEXUS_REPO is required, e.g. terraform-providers}"
: "${NEXUS_USER:?NEXUS_USER is required}"
: "${NEXUS_PASSWORD:?NEXUS_PASSWORD is required}"
: "${VERSION:?VERSION is required, e.g. 7.0.0-SNAPSHOT-2e1dff4c}"

PROJECT="terraform-provider-rancher2"
DIST_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/dist"
DEST_PATH="rancher2/${VERSION}"

shopt -s nullglob
files=(
  "${DIST_DIR}/${PROJECT}_${VERSION}"_*.zip
  "${DIST_DIR}/${PROJECT}_${VERSION}_SHA256SUMS"
  "${DIST_DIR}/${PROJECT}_${VERSION}_SHA256SUMS.sig"
  "${DIST_DIR}/${PROJECT}_${VERSION}_manifest.json"
)

# nullglob only drops the *.zip glob when it matches nothing -- the three
# literal (non-glob) paths above stay in the array even if those files don't
# exist, so filter down to what's really there before checking for "nothing
# to upload" (otherwise this check can never fire, and a typo'd VERSION
# silently "succeeds" with zero uploads).
existing=()
for f in "${files[@]}"; do
  [ -f "$f" ] && existing+=("$f")
done
files=("${existing[@]}")

if [ ${#files[@]} -eq 0 ]; then
  echo "No artifacts found for version ${VERSION} in ${DIST_DIR}" >&2
  echo "Run 'goreleaser release --snapshot --clean --skip=sign,publish,validate' first." >&2
  exit 1
fi

for f in "${files[@]}"; do
  fname="$(basename "$f")"
  url="${NEXUS_URL%/}/repository/${NEXUS_REPO}/${DEST_PATH}/${fname}"
  echo "Uploading ${fname} -> ${url}"
  curl -sSf -u "${NEXUS_USER}:${NEXUS_PASSWORD}" --upload-file "$f" "$url"
done

echo
echo "Done. Artifacts available under:"
echo "  ${NEXUS_URL%/}/repository/${NEXUS_REPO}/${DEST_PATH}/"

# --- Network mirror protocol files -----------------------------------------
# `tofu providers mirror` always speaks the origin registry protocol, which
# nexus.wibu.local doesn't implement, so we hand-write the two static JSON
# files the network mirror protocol actually needs (index.json + a
# per-version archives file) straight from the SHA256SUMS above, and drop
# them next to the zips at <nexus-host>/rancher/rancher2/ (the mirror path
# is <hostname>/<namespace>/<type>/, confirmed via TF_LOG=trace -- the
# hostname segment is NOT optional, OpenTofu builds the request path from
# the full provider source address incl. host), flat, lowercase version --
# see https://opentofu.org/docs/internals/provider-network-mirror-protocol/.
SHASUMS="${DIST_DIR}/${PROJECT}_${VERSION}_SHA256SUMS"
if [ -f "$SHASUMS" ]; then
  command -v jq >/dev/null || { echo "jq is required for the mirror JSON step" >&2; exit 1; }

  MIRROR_VERSION="$(echo "$VERSION" | tr '[:upper:]' '[:lower:]')"
  MIRROR_HOST="$(echo "${NEXUS_URL#*://}" | cut -d/ -f1)"
  MIRROR_DEST_PATH="${MIRROR_HOST}/rancher/rancher2"
  MIRROR_DIR="$(mktemp -d)"
  trap 'rm -rf "$MIRROR_DIR"' EXIT

  echo "{\"versions\":{\"${MIRROR_VERSION}\":{}}}" | jq . > "${MIRROR_DIR}/index.json"

  archives="{}"
  while read -r hash fname; do
    case "$fname" in *.zip) ;; *) continue ;; esac
    plat="$(echo "$fname" | sed -E "s/^${PROJECT}_${VERSION}_([a-zA-Z0-9]+_[a-zA-Z0-9]+)\.zip\$/\1/")"
    archives="$(jq --arg plat "$plat" --arg url "$fname" --arg hash "$hash" \
      '.[$plat] = {url: $url, hashes: ["zh:" + $hash]}' <<<"$archives")"
  done < "$SHASUMS"
  jq -n --argjson archives "$archives" '{archives: $archives}' > "${MIRROR_DIR}/${MIRROR_VERSION}.json"

  echo
  echo "Uploading network mirror files + zips -> ${NEXUS_URL%/}/repository/${NEXUS_REPO}/${MIRROR_DEST_PATH}/"
  for f in "${MIRROR_DIR}/index.json" "${MIRROR_DIR}/${MIRROR_VERSION}.json" "${DIST_DIR}/${PROJECT}_${VERSION}"_*.zip; do
    [ -f "$f" ] || continue
    fname="$(basename "$f")"
    url="${NEXUS_URL%/}/repository/${NEXUS_REPO}/${MIRROR_DEST_PATH}/${fname}"
    echo "Uploading ${fname} -> ${url}"
    curl -sSf -u "${NEXUS_USER}:${NEXUS_PASSWORD}" --upload-file "$f" "$url"
  done
fi
