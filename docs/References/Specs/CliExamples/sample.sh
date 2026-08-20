#!/usr/bin/env bash
# This file is an illustrative sample, not part of the build.
#
# <Name>.sh — one sentence naming the goal this script demonstrates.
#
# Run it from the project root:
#   bash ./examples/cliExamples/<Name>.sh
set -euo pipefail

# Build the binary into a scratch directory and run it in a vault of its own,
# so the example never touches a brain of yours.
workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

go build -o "$workdir/wraith" ./cmd/main
mkdir -p "$workdir/vault"
cd "$workdir/vault"
wraith() { "$workdir/wraith" "$@"; }

echo "== what the commands below show"
wraith start
wraith run AddAccount --account Bank

echo
echo "== what the next commands show"
wraith tick
sed -n '1,12p' DashBoard/README.md
