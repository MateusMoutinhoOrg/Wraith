#!/usr/bin/env bash
#
# DriveItWithTaskFile.sh — the state machine, driven the way a person drives it.
#
# Everything here happens through `Task.yaml`: write an action, arm it with
# `apply: true`, run a tick. It also shows what a failure looks like — an
# `Error.md` in the vault, and nothing changed.
#
# Run it from the project root:
#   bash ./examples/cliExamples/DriveItWithTaskFile.sh

set -uo pipefail

workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

go build -o "$workdir/wraith" ./cmd/main
mkdir -p "$workdir/vault"
cd "$workdir/vault"

wraith() { "$workdir/wraith" "$@"; }

wraith start > /dev/null

echo "== One action per tick, written into Task.yaml"
cat > Task.yaml <<'TASK'
name: AddAccount
account: Bank
opening: 1500
apply: true
TASK
wraith tick

echo
echo "== The tick disarms the file, so the same action never runs twice"
cat Task.yaml

echo
echo "== A task that cannot work changes nothing, and says why"
cat > Task.yaml <<'TASK'
name: AddTransaction
account: Ghost
category: Food
amount: -10
date: 2026-08-18
apply: true
TASK
wraith tick
echo "exit code: $?"
cat Error.md

echo
echo "== The account that does exist is still exactly as it was"
sed -n '1,12p' DashBoard/Accounts.md
