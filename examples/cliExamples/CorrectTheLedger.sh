#!/usr/bin/env bash
# CorrectTheLedger.sh — fixing what was recorded wrong in a ledger that
# already holds a few months: correcting a movement by its id, moving one that
# was filed under the wrong account, removing one that never happened, and
# clearing out a registry entry nothing points at any more.
#
# The vault is pinned to August 2026 by the three dashboard flags of `start`,
# so the transcript below is the same one every time it is run.
#
# Run it from the project root:
#   bash ./examples/cliExamples/CorrectTheLedger.sh
set -uo pipefail

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

echo "== a ledger with three months in it, the last one typed in a hurry"
wraith start --prev-months 4 --future-months 3 --current-month 2026-08 > /dev/null
wraith run AddAccount --account Bank
wraith run AddAccount --account Wallet
wraith run AddCategory --category Salary --description "Monthly pay" --revenues true --expenses false
wraith run AddCategory --category Food --description "Groceries and eating out" --revenues false --expenses true
wraith run AddCategory --category Transport --description "Fuel and public transit" --revenues false --expenses true
wraith run AddCategory --category Travel --description "Trips" --revenues false --expenses true
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false

echo
echo "== what the two accounts already held"
wraith run AddTransaction --account Bank --category "Opening balance" --amount 1000 --date 2026-05-01 --description "Balance when the vault started" --quiet
wraith run AddTransaction --account Wallet --category "Opening balance" --amount 120 --date 2026-05-01 --description "Cash in hand when the vault started" --quiet

echo
echo "== May and June, recorded carefully"
wraith run AddTransaction --account Bank --category Salary --amount 4200 --date 2026-05-05 --description "May pay" --quiet
wraith run AddTransaction --account Bank --category Food --amount -412.30 --date 2026-05-12 --description "May supermarket" --quiet
wraith run AddTransaction --account Bank --category Transport --amount -96.00 --date 2026-05-18 --description "May transit pass" --quiet
wraith run AddTransaction --account Bank --category Salary --amount 4200 --date 2026-06-05 --description "June pay" --quiet
wraith run AddTransaction --account Bank --category Food --amount -388.75 --date 2026-06-14 --description "June supermarket" --quiet
wraith run AddTransaction --account Bank --category Transport --amount -128.00 --date 2026-06-21 --description "June fuel" --quiet
echo "May and June recorded"

echo
echo "== July, recorded in a hurry: one amount slipped, one line is in the wrong"
echo "== account, and one was entered twice"
wraith run AddTransaction --account Bank --category Salary --amount 4200 --date 2026-07-05 --description "July pay" --quiet
wraith run AddTransaction --account Wallet --category Food --amount -32.90 --date 2026-07-11 --description "Bakery" --quiet
wraith run AddTransaction --account Bank --category Food --amount -890 --date 2026-07-13 --description "Supermarket" --quiet
wraith run AddTransaction --account Bank --category Travel --amount -240 --date 2026-07-15 --description "Train tickets" --quiet
wraith run AddTransaction --account Bank --category Food --amount -54.20 --date 2026-07-19 --description "Supermarket" --quiet
echo "July recorded"

echo
echo "== August has only just opened: the pay has landed, nothing else has"
wraith run AddTransaction --account Bank --category Salary --amount 4200 --date 2026-08-01 --description "August pay"

echo
echo "== every movement carries the id a correction is made by"
sed -n '/| Date | Id/,/^$/p' DashBoard/Months/2026-07/Statement.md

echo
echo "== correcting the amount that was typed with a digit too many"
wraith run ModifyTransaction --id 11 --amount -89 --description "Supermarket, corrected"

echo
echo "== a correction can move it too — wrong account, wrong date"
wraith run ModifyTransaction --id 12 --account Wallet --date 2026-07-16
grep '| 12 |' DashBoard/Months/2026-07/Statement.md

echo
echo "== removing the one that was entered twice and never happened"
wraith run RemoveTransaction --id 13

echo
echo "== an id that is not there is refused, and nothing is written"
wraith run RemoveTransaction --id 99
echo "exit code: $?"

echo
echo "== the trip was cancelled and refunded, so the category has nothing left to"
echo "== classify — but a registry entry only goes once nothing points at it"
wraith run RemoveCategory --category Travel
echo "exit code: $?"

echo
echo "== so remove the movement first, then the entry itself"
wraith run RemoveTransaction --id 12
wraith run RemoveCategory --category Travel

echo
echo "== what July holds after the corrections"
sed -n '/| Date | Id/,/^$/p' DashBoard/Months/2026-07/Statement.md

echo
echo "== and what the months come to now that they are right"
sed -n '/^## 1\./,/^## 2\./p' DashBoard/Months/README.md
