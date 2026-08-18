#!/usr/bin/env bash
#
# StartAVault.sh — create a brain from nothing and watch it fill itself in.
#
# Shows the shortest possible path from an empty folder to a rendered vault:
# `start` writes the two files a brain is driven by, `tick` applies whatever
# `Task.yaml` holds, and every page under DashBoard/ appears by itself.
#
# Run it from the project root:
#   bash ./examples/cliExamples/StartAVault.sh

set -euo pipefail

workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

go build -o "$workdir/wraith" ./cmd/main
mkdir -p "$workdir/vault"
cd "$workdir/vault"

wraith() { "$workdir/wraith" "$@"; }

echo "== A brain starts as two files"
wraith start
ls

echo
echo "== The first tick renders the vault, empty as it is"
wraith tick
find . -type f -name '*.md' | sort | head -12

echo
echo "== Give it something to hold"
wraith run AddAccount --account Bank --opening 3000
wraith run AddCategory --category Food --description Groceries --revenues false --expenses true
wraith run AddTransaction --account Bank --category Food --amount -32.90 --date 2026-08-18 --description Market

echo
echo "== And the dashboard has redrawn itself"
sed -n '1,14p' DashBoard/README.md
