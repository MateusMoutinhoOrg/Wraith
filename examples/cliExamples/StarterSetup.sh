#!/usr/bin/env bash
# StarterSetup.sh — the setup an ordinary household starts from: the everyday
# categories, the two places money actually sits, and four months of living
# recorded against them.
#
# The vault is pinned to August 2026 by the three dashboard flags of `start`,
# so the transcript below is the same one every time it is run.
#
# Run it from the project root:
#   bash ./examples/cliExamples/StarterSetup.sh
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

echo "== an empty folder becomes a vault, opened on a month of its own choosing"
wraith start --prev-months 6 --future-months 4 --current-month 2026-08
ls

echo
echo "== basic categories for everyday life"
wraith run AddCategory --category "Income" --description "Salary and other income" --revenues true --expenses false
wraith run AddCategory --category "Housing" --description "Rent and utilities" --revenues false --expenses true
wraith run AddCategory --category "Food" --description "Groceries and dining" --revenues false --expenses true
wraith run AddCategory --category "Transport" --description "Public transit and fuel" --revenues false --expenses true
wraith run AddCategory --category "Leisure" --description "Entertainment and hobbies" --revenues false --expenses true
wraith run AddCategory --category "Cash withdrawal" --description "Money moved from the bank into the wallet" --revenues false --expenses false

echo
echo "== where the money sits: a bank account and a physical wallet"
wraith run AddAccount --account "Bank Account"
wraith run AddAccount --account "Wallet"

echo
echo "== an account starts empty: what it already holds is a movement like any other"
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false
wraith run AddTransaction --account "Bank Account" --category "Opening balance" --amount 1200 --date "2026-04-01" --description "Balance when the vault started"
wraith run AddTransaction --account "Wallet" --category "Opening balance" --amount 60 --date "2026-04-01" --description "Cash in hand when the vault started"

echo
echo "== April: pay in, the fixed lines out, and cash drawn into the wallet"
wraith run AddTransaction --account "Bank Account" --category "Income" --amount 3600 --date "2026-04-05" --description "April salary" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -1150 --date "2026-04-08" --description "April rent" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -142.30 --date "2026-04-19" --description "April electricity" --quiet
wraith run AddTransaction --account "Bank Account" --category "Food" --amount -486.20 --date "2026-04-11" --description "April supermarket" --quiet
wraith run AddTransaction --account "Bank Account" --category "Transport" --amount -96.00 --date "2026-04-06" --description "April transit pass" --quiet
wraith run AddTransaction --account "Bank Account" --category "Cash withdrawal" --amount -200 --date "2026-04-06" --description "Cash for the month" --quiet
wraith run AddTransaction --account "Wallet" --category "Cash withdrawal" --amount 200 --date "2026-04-06" --description "Cash for the month" --quiet
wraith run AddTransaction --account "Wallet" --category "Food" --amount -78.50 --date "2026-04-22" --description "Lunches and the bakery" --quiet
wraith run AddTransaction --account "Wallet" --category "Leisure" --amount -64.00 --date "2026-04-25" --description "Cinema and a round" --quiet
echo "April recorded"

echo
echo "== May and June, the same shape with different numbers"
wraith run AddTransaction --account "Bank Account" --category "Income" --amount 3600 --date "2026-05-05" --description "May salary" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -1150 --date "2026-05-08" --description "May rent" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -118.90 --date "2026-05-19" --description "May electricity" --quiet
wraith run AddTransaction --account "Bank Account" --category "Food" --amount -531.75 --date "2026-05-12" --description "May supermarket" --quiet
wraith run AddTransaction --account "Bank Account" --category "Transport" --amount -96.00 --date "2026-05-06" --description "May transit pass" --quiet
wraith run AddTransaction --account "Bank Account" --category "Cash withdrawal" --amount -200 --date "2026-05-06" --description "Cash for the month" --quiet
wraith run AddTransaction --account "Wallet" --category "Cash withdrawal" --amount 200 --date "2026-05-06" --description "Cash for the month" --quiet
wraith run AddTransaction --account "Wallet" --category "Food" --amount -92.10 --date "2026-05-21" --description "Lunches and the bakery" --quiet
wraith run AddTransaction --account "Wallet" --category "Leisure" --amount -110.00 --date "2026-05-30" --description "Concert ticket" --quiet

wraith run AddTransaction --account "Bank Account" --category "Income" --amount 3600 --date "2026-06-05" --description "June salary" --quiet
wraith run AddTransaction --account "Bank Account" --category "Income" --amount 1800 --date "2026-06-20" --description "Half of the yearly bonus" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -1150 --date "2026-06-08" --description "June rent" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -131.40 --date "2026-06-19" --description "June electricity" --quiet
wraith run AddTransaction --account "Bank Account" --category "Food" --amount -604.90 --date "2026-06-13" --description "June supermarket" --quiet
wraith run AddTransaction --account "Bank Account" --category "Transport" --amount -96.00 --date "2026-06-06" --description "June transit pass" --quiet
wraith run AddTransaction --account "Bank Account" --category "Leisure" --amount -420.00 --date "2026-06-27" --description "Weekend away" --quiet
echo "May and June recorded"

echo
echo "== July: the month the car needed work, which is what a category is for"
wraith run AddTransaction --account "Bank Account" --category "Income" --amount 3600 --date "2026-07-05" --description "July salary" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -1150 --date "2026-07-08" --description "July rent" --quiet
wraith run AddTransaction --account "Bank Account" --category "Housing" --amount -159.80 --date "2026-07-19" --description "July electricity" --quiet
wraith run AddTransaction --account "Bank Account" --category "Food" --amount -572.40 --date "2026-07-14" --description "July supermarket" --quiet
wraith run AddTransaction --account "Bank Account" --category "Transport" --amount -96.00 --date "2026-07-06" --description "July transit pass" --quiet
wraith run AddTransaction --account "Bank Account" --category "Transport" --amount -880.00 --date "2026-07-23" --description "Car repair"
echo "July recorded"

echo
echo "== August has only just opened: the salary has landed, nothing else has moved"
wraith run AddTransaction --account "Bank Account" --category "Income" --amount 3600 --date "2026-08-01" --description "August salary"

echo
echo "== the vault those accounts drew"
sed -n '/^| Indicator/,/^## 2\./p' DashBoard/README.md

echo
echo "== what each category has come to, this month and across the whole ledger"
sed -n '/| Category/,/^$/p' DashBoard/Categories.md
