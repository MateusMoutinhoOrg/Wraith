#!/usr/bin/env bash
# BasicVault.sh — goes from an empty folder to a rendered vault: the two
# registries, half a year of movements, a transfer written as its two legs,
# and the position they draw.
#
# The vault is pinned to August 2026 by the three dashboard flags of `start`,
# so the transcript below is the same one every time it is run.
#
# Run it from the project root:
#   bash ./examples/cliExamples/BasicVault.sh
set -euo pipefail

# Build the binary into a scratch directory and run it in a vault of its own,
# so the example never touches a brain of yours and never depends on which
# version happens to be installed.
workdir="$(mktemp -d)"
trap 'rm -rf "$workdir"' EXIT

PROJECT_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
go build -o "$workdir/wraith" "$PROJECT_ROOT/cmd/main"

VAULT_DIR="$PROJECT_ROOT/WraithSample"
rm -rf "$VAULT_DIR"
mkdir -p "$VAULT_DIR"
cd "$VAULT_DIR"

wraith() { "$workdir/wraith" "$@"; }

echo "== an empty folder becomes a vault, opened on a month of its own choosing"
wraith start --prev-months 6 --future-months 6 --current-month 2026-08
ls

echo
echo "== what a movement can be classified under"
wraith run AddCategory --category "Salary" --description "Monthly pay" --revenues true --expenses false
wraith run AddCategory --category "Housing" --description "Rent and utilities" --revenues false --expenses true
wraith run AddCategory --category "Food" --description "Groceries and eating out" --revenues false --expenses true
wraith run AddCategory --category "Transport" --description "Fuel and public transit" --revenues false --expenses true
wraith run AddCategory --category "Electronics" --description "Electronics purchases" --revenues false --expenses true
wraith run AddCategory --category "Investments" --description "Money moved into the brokerage" --revenues false --expenses false
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false

echo
echo "== where the money sits"
wraith run AddAccount --account "Checking Account"
wraith run AddAccount --account "Brokerage"

echo
echo "== an account is born empty, so what it already held is recorded as a movement"
wraith run AddTransaction --account "Checking Account" --category "Opening balance" --amount 1200 --date "2026-03-01" --description "Balance when the vault started"
wraith run AddTransaction --account "Brokerage" --category "Opening balance" --amount 4200 --date "2026-03-01" --description "Balance when the vault started"

echo
echo "== March: the rhythm every month repeats — pay in, rent out, living costs"
wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5000 --date "2026-03-05" --description "March pay" --quiet
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-03-10" --description "March rent" --quiet
wraith run AddTransaction --account "Checking Account" --category "Food" --amount -612.40 --date "2026-03-14" --description "March groceries" --quiet
wraith run AddTransaction --account "Checking Account" --category "Transport" --amount -128.00 --date "2026-03-18" --description "March transit pass" --quiet
echo "March recorded"

echo
echo "== a transfer is its two legs: one account down, the other up"
wraith run AddTransaction --account "Checking Account" --category "Investments" --amount -1000 --date "2026-03-20" --description "March contribution"
wraith run AddTransaction --account "Brokerage" --category "Investments" --amount 1000 --date "2026-03-20" --description "March contribution"

echo
echo "== April, May, June and July, each one the same shape with different numbers"
wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5000 --date "2026-04-05" --description "April pay" --quiet
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-04-10" --description "April rent" --quiet
wraith run AddTransaction --account "Checking Account" --category "Food" --amount -548.90 --date "2026-04-13" --description "April groceries" --quiet
wraith run AddTransaction --account "Checking Account" --category "Transport" --amount -128.00 --date "2026-04-18" --description "April transit pass" --quiet
wraith run AddTransaction --account "Checking Account" --category "Investments" --amount -1000 --date "2026-04-20" --description "April contribution" --quiet
wraith run AddTransaction --account "Brokerage" --category "Investments" --amount 1000 --date "2026-04-20" --description "April contribution" --quiet

wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5000 --date "2026-05-05" --description "May pay" --quiet
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-05-10" --description "May rent" --quiet
wraith run AddTransaction --account "Checking Account" --category "Food" --amount -703.15 --date "2026-05-15" --description "May groceries" --quiet
wraith run AddTransaction --account "Checking Account" --category "Transport" --amount -128.00 --date "2026-05-18" --description "May transit pass" --quiet
wraith run AddTransaction --account "Checking Account" --category "Investments" --amount -1000 --date "2026-05-20" --description "May contribution" --quiet
wraith run AddTransaction --account "Brokerage" --category "Investments" --amount 1000 --date "2026-05-20" --description "May contribution" --quiet

wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5000 --date "2026-06-05" --description "June pay" --quiet
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-06-10" --description "June rent" --quiet
wraith run AddTransaction --account "Checking Account" --category "Food" --amount -581.05 --date "2026-06-12" --description "June groceries" --quiet
wraith run AddTransaction --account "Checking Account" --category "Transport" --amount -128.00 --date "2026-06-18" --description "June transit pass" --quiet
wraith run AddTransaction --account "Checking Account" --category "Investments" --amount -1000 --date "2026-06-20" --description "June contribution" --quiet
wraith run AddTransaction --account "Brokerage" --category "Investments" --amount 1000 --date "2026-06-20" --description "June contribution" --quiet

wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5200 --date "2026-07-05" --description "July pay, after the raise" --quiet
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-07-10" --description "July rent" --quiet
wraith run AddTransaction --account "Checking Account" --category "Food" --amount -659.70 --date "2026-07-16" --description "July groceries" --quiet
wraith run AddTransaction --account "Checking Account" --category "Transport" --amount -128.00 --date "2026-07-18" --description "July transit pass" --quiet
wraith run AddTransaction --account "Checking Account" --category "Investments" --amount -1000 --date "2026-07-20" --description "July contribution" --quiet
wraith run AddTransaction --account "Brokerage" --category "Investments" --amount 1000 --date "2026-07-20" --description "July contribution" --quiet
echo "April to July recorded"

echo
echo "== the month that broke the rhythm: one large purchase, in June"
wraith run AddTransaction --account "Checking Account" --category "Electronics" --amount -2400 --date "2026-06-22" --description "Phone purchase"

echo
echo "== August has only just opened: pay and rent both land on the first"
wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5200 --date "2026-08-01" --description "August pay"
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-08-01" --description "August rent"

echo
echo "== what is already dated ahead of the open month is the whole forecast"
wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5200 --date "2026-09-05" --description "September pay" --quiet
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-09-10" --description "September rent" --quiet
wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5200 --date "2026-10-05" --description "October pay" --quiet
wraith run AddTransaction --account "Checking Account" --category "Housing" --amount -1800 --date "2026-10-10" --description "October rent" --quiet
echo "September and October recorded"

echo
echo "== the vault those movements drew"
sed -n '/^| Indicator/,/^## 2\./p' DashBoard/README.md

echo
echo "== how the open month compares with the one before it"
sed -n '/^## 2\./,/^## 3\./p' DashBoard/README.md

echo
echo "== the forecast: today's position rolled through what is already dated ahead"
sed -n '/^## 3\./,/^The whole projection/p' DashBoard/README.md

echo
echo "== the months that hold a movement, each with a folder of its own"
ls DashBoard/Months
