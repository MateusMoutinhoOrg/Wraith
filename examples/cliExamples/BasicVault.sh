#!/usr/bin/env bash
# BasicVault.sh — goes from an empty folder to a rendered vault: the registries,
# a first income, a transfer and a card purchase, then the dashboard they draw.
#
# Run it from the project root:
#   bash ./examples/cliExamples/BasicVault.sh
set -euo pipefail

# Build the binary into a scratch directory and run it in a vault of its own,
# so the example never touches a brain of yours.
workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

go build -o "$workdir/wraith" ./cmd/main
mkdir -p "$workdir/vault"
cd "$workdir/vault"
wraith() { "$workdir/wraith" "$@"; }

echo "== an empty folder becomes a vault"
wraith start
ls

echo
echo "== what a movement can be classified under"
wraith run AddCategory --category "Salary" --description "Salary income" --revenues true --expenses false
wraith run AddCategory --category "Electronics" --description "Electronics purchases" --revenues false --expenses true
wraith run AddCategory --category "Food" --description "Food expenses" --revenues false --expenses true
wraith run AddCategory --category "Investments" --description "Financial contributions" --revenues true --expenses true

echo
echo "== where the money sits, including one account that owes rather than holds"
wraith run AddAccount --account "Checking Account" --opening 1000
wraith run AddAccount --account "Brokerage" --opening 0
wraith run AddCreditCard --account "Credit Card" --limit 5000 --closing_day 10 --due_day 15

echo
echo "== what came in"
wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5000 --date "2026-08-05" --description "August Salary"

echo
echo "== a transfer is its two legs: one account down, the other up"
wraith run AddTransaction --account "Checking Account" --category "Investments" --amount -1000 --date "2026-08-10" --description "Monthly Contribution"
wraith run AddTransaction --account "Brokerage" --category "Investments" --amount 1000 --date "2026-08-10" --description "Monthly Contribution"

echo
echo "== what went out in cash, and what went out on the card in twelve parts"
wraith run AddTransaction --account "Checking Account" --category "Food" --amount -150 --date "2026-08-12" --description "Grocery Store"
wraith run AddTransaction --account "Credit Card" --category "Electronics" --amount -2400 --date "2026-08-12" --installments 12 --description "Phone Purchase"

echo
echo "== the vault those movements drew"
sed -n '/## 1. Position/,/## 2./p' DashBoard/README.md
