#!/usr/bin/env bash
# MonthlyBudget.sh — dates a household's fixed lines forward and reads the
# forecast they produce. The forecast has one ingredient and only one:
# transactions already dated ahead of today.
#
# The vault is pinned to August 2026 by the three dashboard flags of `start`,
# so the transcript below is the same one every time it is run.
#
# Run it from the project root:
#   bash ./examples/cliExamples/MonthlyBudget.sh
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

echo "== the account the month runs through, and what its lines are classified under"
wraith start --prev-months 3 --future-months 6 --current-month 2026-08 > /dev/null
wraith run AddAccount --account Bank
wraith run AddCategory --category Salary --description "Monthly pay" --revenues true --expenses false
wraith run AddCategory --category Housing --description "Rent and utilities" --revenues false --expenses true
wraith run AddCategory --category Subscriptions --description "Recurring services" --revenues false --expenses true
wraith run AddCategory --category Insurance --description "Policies paid once a year" --revenues false --expenses true
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false

echo
echo "== what the account already held when the vault was started"
wraith run AddTransaction --account Bank --category "Opening balance" --amount 2500 --date 2026-05-01 --description "Balance when the vault started"

echo
echo "== three months already closed, each the same handful of lines"
wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-05-05 --description "May pay" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-05-10 --description "May rent" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -148.60 --date 2026-05-20 --description "May electricity" --quiet
wraith run AddTransaction --account Bank --category Subscriptions --amount -39.90 --date 2026-05-03 --description "May streaming" --quiet

wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-06-05 --description "June pay" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-06-10 --description "June rent" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -132.20 --date 2026-06-20 --description "June electricity" --quiet
wraith run AddTransaction --account Bank --category Subscriptions --amount -39.90 --date 2026-06-03 --description "June streaming" --quiet

wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-07-05 --description "July pay" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-07-10 --description "July rent" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -171.40 --date 2026-07-20 --description "July electricity" --quiet
wraith run AddTransaction --account Bank --category Subscriptions --amount -39.90 --date 2026-07-03 --description "July streaming" --quiet
echo "May, June and July recorded"

echo
echo "== August has only just opened: pay and rent both land on the first"
wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-08-01 --description "August pay"
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-08-01 --description "August rent"

echo
echo "== the months ahead, dated forward — that, and nothing else, is the forecast"
wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-09-05 --description "September pay" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-09-10 --description "September rent" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -160 --date 2026-09-20 --description "September electricity" --quiet
wraith run AddTransaction --account Bank --category Subscriptions --amount -39.90 --date 2026-09-03 --description "September streaming" --quiet

wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-10-05 --description "October pay" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-10-10 --description "October rent" --quiet
wraith run AddTransaction --account Bank --category Subscriptions --amount -39.90 --date 2026-10-03 --description "October streaming" --quiet

wraith run AddTransaction --account Bank --category Salary --amount 5200 --date 2026-11-05 --description "November pay" --quiet
wraith run AddTransaction --account Bank --category Housing --amount -1800 --date 2026-11-10 --description "November rent" --quiet
wraith run AddTransaction --account Bank --category Subscriptions --amount -39.90 --date 2026-11-03 --description "November streaming" --quiet
echo "September to November recorded"

echo
echo "== the yearly bill you know is coming is one line, dated into the month it lands"
wraith run AddTransaction --account Bank --category Insurance --amount -2150 --date 2026-11-14 --description "Home insurance, renewed yearly"

echo
echo "== the forecast: every month a movement is dated into"
sed -n '/^## 3\./,/^The whole projection/p' DashBoard/README.md

echo
echo "== the same walk, month by month, on the month index"
sed -n '/## 3. The next/,/Nothing here is an average/p' DashBoard/Months/README.md

echo
echo "== a longer horizon, rendered on its own without touching the config"
wraith render DashBoard --future-months 12 > /dev/null
sed -n '/^## 3\./,/^The whole projection/p' DashBoard/README.md | grep -c '^| [a-z][a-z][a-z]-' |
	sed 's/^/forecast rows now: /'

echo
echo "== the config still asks for six, so the next tick goes back to six"
grep 'future-months' Visualization.yaml
wraith tick > /dev/null
sed -n '/^## 3\./,/^The whole projection/p' DashBoard/README.md | grep -c '^| [a-z][a-z][a-z]-' |
	sed 's/^/forecast rows now: /'

echo
echo "== November holds the bill, and the ledger names the id it is removed by"
sed -n '/| Date | Id/,/^$/p' DashBoard/Months/2026-11/Statement.md

echo
echo "== a line you no longer expect is removed like any other — it was never a rule"
wraith run RemoveTransaction --id 26
sed -n '/^## 3\./,/^The whole projection/p' DashBoard/README.md
