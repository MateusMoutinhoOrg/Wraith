#!/usr/bin/env bash
# CorrectTheLedger.sh — fixing what was recorded wrong: correcting a movement
# by its id, removing one that never happened, and clearing out a registry
# entry nothing points at any more.
#
# Run it from the project root:
#   bash ./examples/cliExamples/CorrectTheLedger.sh
set -uo pipefail

rm -rf WraithSample

mkdir WraithSample
cd WraithSample/

echo "== a ledger with a few movements in it, one of them typed wrong"
wraith start > /dev/null
wraith run AddAccount --account Bank
wraith run AddAccount --account Wallet
wraith run AddCategory --category Food --description "Groceries and eating out" --revenues false --expenses true
wraith run AddCategory --category Travel --description "Trips" --revenues false --expenses true
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false
wraith run AddTransaction --account Bank --category "Opening balance" --amount 1000 --date 2026-08-01 --description "Balance when the vault started"
wraith run AddTransaction --account Wallet --category "Opening balance" --amount 120 --date 2026-08-01 --description "Cash in hand when the vault started"
wraith run AddTransaction --account Bank --category Food --amount -32.90 --date 2026-08-11 --description "Bakery"
wraith run AddTransaction --account Bank --category Food --amount -890 --date 2026-08-13 --description "Supermarket"
wraith run AddTransaction --account Bank --category Travel --amount -240 --date 2026-08-15 --description "Train tickets"

echo
echo "== every movement carries the id a correction is made by"
sed -n '/| Date | Id/,/^$/p' DashBoard/Months/2026-08/Statement.md

echo
echo "== correcting the amount that was typed with a digit too many"
wraith run ModifyTransaction --id 4 --amount -89 --description "Supermarket, corrected"

echo
echo "== a correction can move it too — wrong category, wrong account, wrong date"
wraith run ModifyTransaction --id 5 --account Wallet --date 2026-08-16

echo
echo "== removing one that never happened at all"
wraith run RemoveTransaction --id 3

echo
echo "== an id that is not there is refused, and nothing is written"
wraith run RemoveTransaction --id 99
echo "exit code: $?"

echo
echo "== a registry entry only goes once nothing is classified under it"
wraith run RemoveCategory --category Travel
echo "exit code: $?"

echo
echo "== so remove what points at it first, then the entry itself"
wraith run RemoveTransaction --id 5
wraith run RemoveCategory --category Travel

echo
echo "== what the month holds after the corrections"
sed -n '/| Date | Id/,/^$/p' DashBoard/Months/2026-08/Statement.md
