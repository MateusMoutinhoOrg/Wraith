#!/usr/bin/env bash
# DriveItWithTaskFile.sh — drives the state machine through Task.yaml instead of
# the command line, including what an armed action that fails looks like.
#
# Run it from the project root:
#   bash ./examples/cliExamples/DriveItWithTaskFile.sh
set -uo pipefail

# Build the binary into a scratch directory and run it in a vault of its own,
# so the example never touches a brain of yours.
workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

go build -o "$workdir/wraith" ./cmd/main
mkdir -p "$workdir/vault"
cd "$workdir/vault"
wraith() { "$workdir/wraith" "$@"; }

echo "== start writes the two files a brain is driven by"
wraith start
ls

echo
echo "== the task file arrives disarmed, so nothing runs until you say so"
cat Task.yaml
wraith tick

echo
echo "== write an action into it and arm it"
cat > Task.yaml <<'YAML'
name: AddAccount
account: Bank
opening: 1200
apply: true
YAML
wraith tick

echo
echo "== the tick disarmed it again, so the same action never runs twice"
grep '^apply:' Task.yaml
wraith tick

echo
echo "== arm the next action against the account the first one created"
cat > Task.yaml <<'YAML'
name: AddCategory
category: Food
description: Groceries and eating out
revenues: false
expenses: true
apply: true
YAML
wraith tick

cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Food
amount: -84.20
date: 2026-08-18
description: Grocery store
apply: true
YAML
wraith tick

echo
echo "== an armed action that fails writes Error.md and changes nothing"
cat > Task.yaml <<'YAML'
name: AddTransaction
account: Savings
category: Food
amount: -10
date: 2026-08-19
apply: true
YAML
wraith tick
echo "exit code: $?"
cat Error.md

echo
echo "== the ledger still holds only the movement that succeeded"
sed -n '1,20p' DashBoard/Accounts/Bank.md
