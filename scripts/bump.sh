#!/usr/bin/env bash
set -euo pipefail

if ! command -v git >/dev/null 2>&1; then
  echo "git is required" >&2
  exit 1
fi

if [ $# -ne 1 ]; then
  echo "usage: scripts/bump.sh <major|minor|patch>" >&2
  exit 1
fi

part="${1:-patch}"
if [[ "$part" != "major" && "$part" != "minor" && "$part" != "patch" ]]; then
  echo "invalid bump part: $part (use major|minor|patch)" >&2
  exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
  echo "working tree is not clean" >&2
  exit 1
fi

latest_tag="$(git tag --list 'v*' --sort=-v:refname | head -n1)"
if [ -z "$latest_tag" ]; then
  latest_tag="v0.0.0"
fi

version="${latest_tag#v}"
IFS='.' read -r major minor patch <<<"$version"

case "$part" in
major)
  major=$((major + 1))
  minor=0
  patch=0
  ;;
minor)
  minor=$((minor + 1))
  patch=0
  ;;
patch)
  patch=$((patch + 1))
  ;;
esac

next_tag="v${major}.${minor}.${patch}"

echo "latest: ${latest_tag}"
echo "next:   ${next_tag}"

read -r -p "create and push ${next_tag}? [y/N] " answer
if [[ "$answer" != "y" && "$answer" != "Y" ]]; then
  echo "aborted"
  exit 0
fi

git tag -a "$next_tag" -m "$next_tag"
git push origin "$next_tag"
echo "pushed ${next_tag}"
