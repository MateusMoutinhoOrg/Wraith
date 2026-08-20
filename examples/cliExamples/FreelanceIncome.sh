#!/usr/bin/env bash
# FreelanceIncome.sh — an irregular income across several accounts: invoices
# landing when they land, a category tree, money set aside for tax, and a
# transfer written as its two legs.
#
# Run it from the project root:
#   bash ./examples/cliExamples/FreelanceIncome.sh
set -euo pipefail

rm -rf WraithSample

mkdir WraithSample
cd WraithSample/

echo "== where the money sits: what it arrives in, what it is kept in"
wraith start > /dev/null
wraith run AddAccount --account Bank --opening 800
wraith run AddAccount --account "Tax Reserve" --opening 0

echo
echo "== a category tree: one parent, one child per source of work"
wraith run AddCategory --category Clients --description "Everything invoiced" --revenues true --expenses false
wraith run AddCategory --category "Client Acme" --description "Retainer work" --revenues true --expenses false --parent Clients
wraith run AddCategory --category "Client Wayne" --description "Project work" --revenues true --expenses false --parent Clients
wraith run AddCategory --category Transfers --description "Money moving between my own accounts" --revenues true --expenses true
wraith run AddCategory --category Tools --description "Software and hardware for the work" --revenues false --expenses true

echo
echo "== invoices land when they land — no two months alike"
wraith run AddTransaction --account Bank --category "Client Acme" --amount 3200 --date 2026-08-04 --description "August retainer"
wraith run AddTransaction --account Bank --category "Client Wayne" --amount 5400 --date 2026-08-14 --description "Project milestone"

echo
echo "== the cost of working"
wraith run AddTransaction --account Bank --category Tools --amount -49.90 --date 2026-08-07 --description "Design tool"

echo
echo "== setting tax aside is one movement out and one movement in, sharing a category"
wraith run AddTransaction --account Bank --category Transfers --amount -2580 --date 2026-08-17 --description "Tax set aside"
wraith run AddTransaction --account "Tax Reserve" --category Transfers --amount 2580 --date 2026-08-17 --description "Tax set aside"

echo
echo "== what each account holds"
sed -n '/## 1. Position/,/## 2./p' DashBoard/README.md

echo
echo "== what each client brought in, the parent totalling its children"
sed -n '/| Category/,/^$/p' DashBoard/Categories.md | head -20
