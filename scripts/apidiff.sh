#!/usr/bin/env bash
# Fail if the exported API of ./githubevents changed incompatibly vs a base ref.
#
# A go-github major bump (e.g. v89 -> v90) re-exports every event type under a
# new import path, which apidiff reports as incompatible even though nothing in
# our own surface changed. Those version-only diffs are filtered out here so the
# gate still catches real breaks (renames, signature changes, removals, and any
# real field change go-github makes to an event type) without blocking routine
# dependency bumps.
set -euo pipefail

BASE="${1:-origin/main}"
PKG="./githubevents"

# real_incompatible_changes reads apidiff output on stdin and prints only the
# incompatible changes that are NOT purely a go-github version-path bump. A
# "changed from X to Y" line whose X and Y are identical once the go-github
# major (/vN/) is normalized away is a version-only change and is dropped.
real_incompatible_changes() {
  local line norm rest from to in_incompat=0
  while IFS= read -r line; do
    case "$line" in
      "Incompatible changes:") in_incompat=1; continue ;;
      "Compatible changes:") in_incompat=0; continue ;;
      "") in_incompat=0; continue ;;
    esac
    [ "$in_incompat" -eq 1 ] || continue
    case "$line" in "- "*) ;; *) continue ;; esac

    norm="$(printf '%s' "$line" | sed 's#go-github/v[0-9]\{1,\}#go-github/vN#g')"
    case "$norm" in
      *" changed from "*" to "*)
        rest="${norm#* changed from }"
        from="${rest% to *}"
        to="${rest#* to }"
        [ "$from" = "$to" ] && continue
        ;;
    esac
    printf '%s\n' "$line"
  done
}

# When sourced for testing, stop here.
[ "${APIDIFF_LIB:-0}" = 1 ] && return 0

command -v apidiff >/dev/null 2>&1 || { echo "apidiff not installed; run 'make tools'"; exit 1; }

work="$(mktemp -d)"
base_wt="$work/base"
cleanup() { git worktree remove --force "$base_wt" >/dev/null 2>&1 || true; rm -rf "$work"; }
trap cleanup EXIT

git worktree add --quiet --detach "$base_wt" "$BASE"
( cd "$base_wt" && go mod download && apidiff -w "$work/base.api" "$PKG" )

out="$(apidiff "$work/base.api" "$PKG")"
printf '%s\n' "$out"

real="$(printf '%s\n' "$out" | real_incompatible_changes)"
if [ -n "$real" ]; then
  echo ">> incompatible API changes detected (beyond go-github version bump):"
  printf '%s\n' "$real"
  exit 1
fi
echo ">> API compatible (go-github version-only changes ignored)"
