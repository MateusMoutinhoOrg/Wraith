#!/usr/bin/env bash
# DriveItWithTaskFile.sh — drives the state machine through Task.yaml instead
# of the command line: one armed action per tick, across several months, plus
# what an armed action that fails looks like.
#
# The vault is pinned to August 2026 by the three dashboard flags of `start`,
# so the transcript below is the same one every time it is run.
#
# Run it from the project root:
#   bash ./examples/cliExamples/DriveItWithTaskFile.sh
set -uo pipefail

# Build the binary into a scratch directory and run it in a vault of its own,
# so the example never touches a brain of yours and never depends on which
# version happens to be installed.
workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

go build -o "$workdir/wraith" ./cmd/main
mkdir -p "$workdir/vault"
cd "$workdir/vault"
wraith() { "$workdir/wraith" "$@"; }

echo "== start writes the two files a brain is driven by"
wraith start --prev-months 4 --future-months 3 --current-month 2026-08
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
category: Salary
description: Monthly pay
revenues: true
expenses: false
apply: true
YAML
wraith tick

cat > Task.yaml <<'YAML'
name: AddCategory
category: Food
description: Groceries and eating out
revenues: false
expenses: true
apply: true
YAML
wraith tick

echo
echo "== one tick per movement, across the months the ledger covers"
cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Salary
amount: 4800
date: 2026-06-05
description: June pay
apply: true
YAML
wraith tick

cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Food
amount: -512.60
date: 2026-06-18
description: June supermarket
apply: true
YAML
wraith tick

cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Salary
amount: 4800
date: 2026-07-05
description: July pay
apply: true
YAML
wraith tick

cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Food
amount: -463.15
date: 2026-07-16
description: July supermarket
apply: true
YAML
wraith tick

echo
echo "== a movement dated ahead of the open month is written the same way"
cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Salary
amount: 4800
date: 2026-09-05
description: September pay
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
echo "== the ledger still holds only the movements that succeeded"
sed -n '/^## 1\./,/^## 2\./p' DashBoard/Accounts/Bank.md

echo
echo "== every month those ticks wrote into"
sed -n '/^## 1\./,/^## 2\./p' DashBoard/Months/README.md
