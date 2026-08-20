#!/usr/bin/env bash
# ChooseWhatIsRendered.sh — the other half of the interface: listing what the
# binary carries, rendering one visualization on its own, overriding where it
# writes, and turning an entry off in Visualization.yaml.
#
# Run it from the project root:
#   bash ./examples/cliExamples/ChooseWhatIsRendered.sh
set -euo pipefail

# Build the binary into a scratch directory and run it in a vault of its own,
# so the example never touches a brain of yours.
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
wraith start > /dev/null
wraith run AddAccount --account Bank --opening 3000 --quiet
wraith run AddCategory --category Food --description "Groceries" --revenues false --expenses true --quiet
wraith run AddTransaction --account Bank --category Food --amount -76.40 --date 2026-08-12 --description "Supermarket" --quiet
ls

echo
echo "== the config decides what a tick draws, and where each page lands"
cat Visualization.yaml

echo
echo "== one visualization on its own, without running the pending task"
wraith render Task-List
ls Tasks | head -5

echo
echo "== an arg given here overrides the config for this invocation only"
wraith render DashBoard --future-months 24
grep '^## 3' DashBoard/README.md
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
    future-months: 8
  dest: DashBoard

- name: Task-List
  dest: Tasks
  enabled: false

- name: Help
  dest: Help
YAML
rm -rf Tasks
cat > Task.yaml <<'YAML'
name: AddTransaction
account: Bank
category: Food
amount: -18.50
date: 2026-08-15
description: Bakery
apply: true
YAML
wraith tick
ls

echo
echo "== asking for it by name renders it anyway: that request is the decision"
wraith render Task-List
ls Tasks | head -3
