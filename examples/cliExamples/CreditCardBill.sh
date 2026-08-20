#!/usr/bin/env bash
# CreditCardBill.sh — puts a credit card in the vault, spends on it, splits a
# purchase into installments, and reads the bill it produces.
#
# Run it from the project root:
#   bash ./examples/cliExamples/CreditCardBill.sh
set -euo pipefail

rm -rf WraithSample

mkdir WraithSample
cd WraithSample/

echo "== a vault with somewhere to pay the bill from"
wraith start > /dev/null
wraith run AddAccount --account Bank --opening 4000
wraith run AddCategory --category Electronics --description "Devices and gadgets" --revenues false --expenses true
wraith run AddCategory --category Food --description "Groceries and eating out" --revenues false --expenses true

echo
echo "== a card is an account with a limit and two days: when it closes, when it is due"
wraith run AddCreditCard --account "Nubank Card" --limit 5000 --closing_day 10 --due_day 17

echo
echo "== spending on the card, charged to the card rather than to the bank"
wraith run AddTransaction --account "Nubank Card" --category Food --amount -142.30 --date 2026-08-06 --description "Supermarket"
wraith run AddTransaction --account "Nubank Card" --category Food --amount -67.90 --date 2026-08-14 --description "Restaurant"

echo
echo "== a purchase split into installments: one charge, twelve monthly parts"
wraith run AddTransaction --account "Nubank Card" --category Electronics --amount -2400 --date 2026-08-12 --installments 12 --description "Phone"

echo
echo "== what the card owes, and how the parts land across the months ahead"
sed -n '1,40p' DashBoard/Credit-Cards.md

echo
echo "== the same card as its own page"
sed -n '1,30p' DashBoard/Accounts/Nubank-Card.md
