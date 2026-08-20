#!/usr/bin/env bash
# ChooseWhatIsRendered.sh — the other half of the interface: listing what the
# binary carries, rendering one visualization on its own, overriding where it
# writes, and turning an entry off in Visualization.yaml.
#
# The vault is pinned to August 2026 by the three dashboard flags of `start`,
# so the transcript below is the same one every time it is run.
#
# Run it from the project root:
#   bash ./examples/cliExamples/ChooseWhatIsRendered.sh
set -euo pipefail

# Build the binary into a scratch directory and run it in a vault of its own,
# so the example never touches a brain of yours and never depends on which
# version happens to be installed.
workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

go build -o "$workdir/wraith" ./cmd/main
mkdir -p "$workdir/vault"
cd "$workdir/vault"
wraith() { "$workdir/wraith" "$@"; }

echo "== what this binary carries, before a vault exists at all"
wraith tasks
wraith visualizations

echo
echo "== a vault with something in it to look at"
echo "== the three dashboard flags of start are written into the config it creates"
wraith start --prev-months 6 --future-months 6 --current-month 2026-08 > /dev/null
wraith run AddAccount --account Bank --quiet
wraith run AddCategory --category Salary --description "Monthly pay" --revenues true --expenses false --quiet
wraith run AddCategory --category Food --description "Groceries" --revenues false --expenses true --quiet
wraith run AddCategory --category Housing --description "Rent and utilities" --revenues false --expenses true --quiet
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false --quiet
wraith run AddTransaction --account Bank --category "Opening balance" --amount 3000 --date 2026-03-01 --description "Balance when the vault started" --quiet
for month in 03 04 05 06 07; do
	wraith run AddTransaction --account Bank --category Salary --amount 4400 --date "2026-$month-05" --description "Pay" --quiet
	wraith run AddTransaction --account Bank --category Housing --amount -1500 --date "2026-$month-10" --description "Rent" --quiet
	wraith run AddTransaction --account Bank --category Food --amount -76.40 --date "2026-$month-12" --description "Supermarket" --quiet
done
wraith run AddTransaction --account Bank --category Salary --amount 4400 --date 2026-08-01 --description "August pay" --quiet
wraith run AddTransaction --account Bank --category Salary --amount 4400 --date 2026-09-05 --description "September pay" --quiet
ls

echo
echo "== the config decides what a tick draws, and where each page lands"
cat Visualization.yaml

echo
echo "== the months it drew, one folder each, back as far as prev-months asked"
ls DashBoard/Months

echo
echo "== one visualization on its own, without running the pending task"
wraith render Task-List
ls Tasks | head -5

echo
echo "== an arg given here overrides the config for this invocation only"
wraith render DashBoard --future-months 12
sed -n '/^## 3\./,/^The whole projection/p' DashBoard/README.md |
	grep -c '^| [a-z][a-z][a-z]-' | sed 's/^/forecast rows now: /'
grep 'future-months' Visualization.yaml

echo
echo "== --dest sends the same page somewhere else, leaving the original alone"
wraith render DashBoard --dest Snapshot
ls Snapshot

echo
echo "== turning an entry off keeps it out of every tick from here on"
cat > Visualization.yaml <<'YAML'
- name: DashBoard
  args:
    prev-months: 3
    future-months: 3
    current-month: 2026-08
  dest: DashBoard

- name: Task-List
  dest: Tasks
  enabled: false

- name: Help
  dest: Help
YAML
rm -rf Tasks DashBoard
cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Food
amount: -18.50
date: 2026-08-01
description: Bakery
apply: true
YAML
wraith tick
ls

echo
echo "== and the config it was ticked with is the one that decided the months"
ls DashBoard/Months

echo
echo "== asking for it by name renders it anyway: that request is the decision"
wraith render Task-List
ls Tasks | head -3
