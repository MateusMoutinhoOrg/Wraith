#!/usr/bin/env bash
# MonthlyBudget.sh — dates a month's fixed lines forward and reads the forecast
# they produce on the month index. The forecast has one ingredient and only
# one: transactions already dated ahead of today.
#
# Run it from the project root:
#   bash ./examples/cliExamples/MonthlyBudget.sh
set -euo pipefail

rm -rf WraithSample

mkdir WraithSample
cd WraithSample/

echo "== the account the month runs through, and what its lines are classified under"
wraith start > /dev/null
wraith run AddAccount --account Bank
wraith run AddCategory --category Salary --description "Monthly pay" --revenues true --expenses false
wraith run AddCategory --category Housing --description "Rent and utilities" --revenues false --expenses true
wraith run AddCategory --category Subscriptions --description "Recurring services" --revenues false --expenses true
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false

echo
echo "== what the account already holds"
wraith run AddTransaction --account Bank --category "Opening balance" --amount 2500 --date 2026-08-01 --description "Balance when the vault started"

echo
echo "== this month's fixed lines, as they already happened"
wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-08-05 --description "August pay"
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-08-10 --description "August rent"

echo
echo "== next month's, dated forward — that is what the forecast reads"
wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-09-05 --description "September pay"
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-09-10 --description "September rent"
wraith run AddTransaction --account Bank --category Housing --amount -160 --date 2026-09-20 --description "September electricity"
wraith run AddTransaction --account Bank --category Subscriptions --amount -39.90 --date 2026-09-03 --description "September streaming"

echo
echo "== and the month after that"
wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-10-05 --description "October pay"
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-10-10 --description "October rent"

echo
echo "== the forecast: every month a movement is dated into"
sed -n '/## 3. The next/,/Nothing here is an average/p' DashBoard/Months/README.md

echo
echo "== a longer horizon, rendered on its own without touching the config"
wraith render DashBoard --future-months 18 > /dev/null
sed -n '/## 3. The next/,/^Every figure above/p' DashBoard/Months/README.md | grep -c '^| ' |
	sed 's/^/forecast rows now: /'

echo
echo "== a line you no longer expect is removed like any other — it was never a rule"
wraith run RemoveTransaction --id 9
sed -n '/## 3. The next/,/Nothing here is an average/p' DashBoard/Months/README.md
