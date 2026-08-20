#!/usr/bin/env bash
# MonthlyBudget.sh — declares the commitments that repeat every month and reads
# the forecast they produce on the month index.
#
# Run it from the project root:
#   bash ./examples/cliExamples/MonthlyBudget.sh
set -euo pipefail

rm -rf WraithSample

mkdir WraithSample
cd WraithSample/

echo "== the account the month runs through, and what its lines are classified under"
wraith start > /dev/null
wraith run AddAccount --account Bank --opening 2500
wraith run AddCategory --category Salary --description "Monthly pay" --revenues true --expenses false
wraith run AddCategory --category Housing --description "Rent and utilities" --revenues false --expenses true
wraith run AddCategory --category Subscriptions --description "Recurring services" --revenues false --expenses true

echo
echo "== what comes in every month, open-ended because there is no end date to it"
wraith run AddRecurrence --description "Salary" --account Bank --category Salary --amount 5200 --day 5 --start 2026-08

echo
echo "== what goes out every month, on the day it falls on"
wraith run AddRecurrence --description "Rent" --account Bank --category Housing --amount -1800 --day 10 --start 2026-08
wraith run AddRecurrence --description "Electricity" --account Bank --category Housing --amount -160 --day 20 --start 2026-08

echo
echo "== a commitment that ends: a subscription paid up to december only"
wraith run AddRecurrence --description "Streaming" --account Bank --category Subscriptions --amount -39.90 --day 3 --start 2026-08 --end 2026-12

echo
echo "== the forecast: every month the commitments above reach"
sed -n '/## 3. The next/,/Nothing here is an average/p' DashBoard/Months/README.md

echo
echo "== a longer horizon, rendered on its own without touching the config"
wraith render DashBoard --future-months 18 > /dev/null
sed -n '/## 3. The next/,/^---$/p' DashBoard/Months/README.md | grep -c '^| ' |
	sed 's/^/forecast rows now: /'

echo
echo "== stopping one is naming it back — it leaves the commitments and the forecast"
wraith run RemoveRecurrence --recurrence "Streaming"
sed -n '/## 4. The commitments/,/A recurrence never/p' DashBoard/Months/README.md
