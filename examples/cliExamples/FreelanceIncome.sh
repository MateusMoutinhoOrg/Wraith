#!/usr/bin/env bash
# FreelanceIncome.sh — an irregular income across several accounts: half a
# year of invoices landing when they land, a category tree per client, money
# set aside for tax as a transfer written in two legs, and the quarterly bill
# it is there to pay.
#
# The vault is pinned to August 2026 by the three dashboard flags of `start`,
# so the transcript below is the same one every time it is run.
#
# Run it from the project root:
#   bash ./examples/cliExamples/FreelanceIncome.sh
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

echo "== where the money sits: what it arrives in, what it is kept in"
wraith start --prev-months 6 --future-months 4 --current-month 2026-08 > /dev/null
wraith run AddAccount --account Bank
wraith run AddAccount --account "Tax Reserve"

echo
echo "== a category tree: one parent, one child per source of work"
wraith run AddCategory --category Clients --description "Everything invoiced" --revenues true --expenses false
wraith run AddCategory --category "Client Acme" --description "Retainer work" --revenues true --expenses false --parent Clients
wraith run AddCategory --category "Client Wayne" --description "Project work" --revenues true --expenses false --parent Clients
wraith run AddCategory --category "Client Stark" --description "Occasional consulting" --revenues true --expenses false --parent Clients
wraith run AddCategory --category Transfers --description "Money moving between my own accounts" --revenues false --expenses false
wraith run AddCategory --category Tools --description "Software and hardware for the work" --revenues false --expenses true
wraith run AddCategory --category Taxes --description "What is actually paid to the revenue service" --revenues false --expenses true
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false

echo
echo "== what the bank already held, put in as the movement it is"
wraith run AddTransaction --account Bank --category "Opening balance" --amount 800 --date 2026-03-01 --description "Balance when the vault started"

echo
echo "== March and April: the retainer is steady, everything else is not"
wraith run AddTransaction --account Bank --category "Client Acme" --amount 3200 --date 2026-03-04 --description "March retainer" --quiet
wraith run AddTransaction --account Bank --category "Client Wayne" --amount 4100 --date 2026-03-19 --description "Project milestone one" --quiet
wraith run AddTransaction --account Bank --category Tools --amount -49.90 --date 2026-03-07 --description "Design tool" --quiet
wraith run AddTransaction --account Bank --category Transfers --amount -2190 --date 2026-03-25 --description "Tax set aside" --quiet
wraith run AddTransaction --account "Tax Reserve" --category Transfers --amount 2190 --date 2026-03-25 --description "Tax set aside" --quiet

wraith run AddTransaction --account Bank --category "Client Acme" --amount 3200 --date 2026-04-06 --description "April retainer" --quiet
wraith run AddTransaction --account Bank --category Tools --amount -49.90 --date 2026-04-07 --description "Design tool" --quiet
wraith run AddTransaction --account Bank --category Transfers --amount -960 --date 2026-04-27 --description "Tax set aside" --quiet
wraith run AddTransaction --account "Tax Reserve" --category Transfers --amount 960 --date 2026-04-27 --description "Tax set aside" --quiet
echo "March and April recorded"

echo
echo "== May: a month with no project work at all, which is the point of the reserve"
wraith run AddTransaction --account Bank --category "Client Acme" --amount 3200 --date 2026-05-05 --description "May retainer" --quiet
wraith run AddTransaction --account Bank --category Tools --amount -49.90 --date 2026-05-07 --description "Design tool" --quiet
wraith run AddTransaction --account Bank --category Tools --amount -1890 --date 2026-05-18 --description "New laptop" --quiet
wraith run AddTransaction --account Bank --category Transfers --amount -960 --date 2026-05-26 --description "Tax set aside" --quiet
wraith run AddTransaction --account "Tax Reserve" --category Transfers --amount 960 --date 2026-05-26 --description "Tax set aside" --quiet
echo "May recorded"

echo
echo "== June and July: two clients at once, and the quarterly bill the reserve pays"
wraith run AddTransaction --account Bank --category "Client Acme" --amount 3200 --date 2026-06-04 --description "June retainer" --quiet
wraith run AddTransaction --account Bank --category "Client Wayne" --amount 5400 --date 2026-06-16 --description "Project milestone two" --quiet
wraith run AddTransaction --account Bank --category "Client Stark" --amount 1250 --date 2026-06-23 --description "Two days of consulting" --quiet
wraith run AddTransaction --account Bank --category Tools --amount -49.90 --date 2026-06-07 --description "Design tool" --quiet
wraith run AddTransaction --account Bank --category Transfers --amount -2955 --date 2026-06-26 --description "Tax set aside" --quiet
wraith run AddTransaction --account "Tax Reserve" --category Transfers --amount 2955 --date 2026-06-26 --description "Tax set aside" --quiet

wraith run AddTransaction --account Bank --category "Client Acme" --amount 3200 --date 2026-07-06 --description "July retainer" --quiet
wraith run AddTransaction --account Bank --category "Client Stark" --amount 1900 --date 2026-07-21 --description "Three days of consulting" --quiet
wraith run AddTransaction --account Bank --category Tools --amount -49.90 --date 2026-07-07 --description "Design tool" --quiet
wraith run AddTransaction --account Bank --category Transfers --amount -1530 --date 2026-07-27 --description "Tax set aside" --quiet
wraith run AddTransaction --account "Tax Reserve" --category Transfers --amount 1530 --date 2026-07-27 --description "Tax set aside" --quiet
echo "June and July recorded"

echo
echo "== the reserve is not savings: the quarterly bill leaves from it"
wraith run AddTransaction --account "Tax Reserve" --category Taxes --amount -4110 --date 2026-07-20 --description "Second quarter tax"

echo
echo "== August has only just opened: the retainer has landed, nothing else has"
wraith run AddTransaction --account Bank --category "Client Acme" --amount 3200 --date 2026-08-01 --description "August retainer"

echo
echo "== an invoice already agreed is dated into the month it is due"
wraith run AddTransaction --account Bank --category "Client Wayne" --amount 6200 --date 2026-09-15 --description "Project milestone three, agreed"

echo
echo "== what each account holds"
sed -n '/^| Indicator/,/^## 2\./p' DashBoard/README.md

echo
echo "== what each client brought in, the parent totalling its children"
sed -n '/| Category/,/^$/p' DashBoard/Categories.md

echo
echo "== half a year of uneven months, read off the month index"
sed -n '/^## 1\./,/^## 2\./p' DashBoard/Months/README.md
