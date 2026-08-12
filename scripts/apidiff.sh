#!/usr/bin/env bash
# Fail if the exported API of ./githubevents changed incompatibly vs a base ref.
set -euo pipefail

BASE="${1:-origin/main}"
PKG="./githubevents"

command -v apidiff >/dev/null 2>&1 || { echo "apidiff not installed; run 'make tools'"; exit 1; }

work="$(mktemp -d)"
base_wt="$work/base"
cleanup() { git worktree remove --force "$base_wt" >/dev/null 2>&1 || true; rm -rf "$work"; }
trap cleanup EXIT

git worktree add --quiet --detach "$base_wt" "$BASE"
( cd "$base_wt" && go mod download && apidiff -w "$work/base.api" "$PKG" )

out="$(apidiff "$work/base.api" "$PKG")"
printf '%s\n' "$out"
if printf '%s\n' "$out" | grep -q '^Incompatible changes:'; then
  echo ">> incompatible API changes detected"
  exit 1
fi
echo ">> API compatible"
