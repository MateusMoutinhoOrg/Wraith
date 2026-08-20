#!/usr/bin/env bash

rm -rf WraithSample

mkdir WraithSample
cd WraithSample/

echo "== an empty folder becomes a vault"
wraith start 
ls

echo
echo "== basic categories for everyday life"
wraith run AddCategory --category "Food" --description "Groceries and dining" --revenues false --expenses true
wraith run AddCategory --category "Transport" --description "Public transit and fuel" --revenues false --expenses true
wraith run AddCategory --category "Housing" --description "Rent and utilities" --revenues false --expenses true
wraith run AddCategory --category "Leisure" --description "Entertainment and hobbies" --revenues false --expenses true
wraith run AddCategory --category "Income" --description "Salary and other income" --revenues true --expenses false

echo
echo "== where the money sits: a bank account and a physical wallet"
wraith run AddAccount --account "Bank Account"
wraith run AddAccount --account "Wallet"

echo
echo "== an account starts empty: what it already holds is a movement like any other"
wraith run AddCategory --category "Opening balance" --description "What an account already held when it was added" --revenues false --expenses false
wraith run AddTransaction --account "Bank Account" --category "Opening balance" --amount 1200 --date 2026-08-01 --description "Balance when the vault started"
wraith run AddTransaction --account "Wallet" --category "Opening balance" --amount 60 --date 2026-08-01 --description "Cash in hand when the vault started"

echo
echo "== the vault those accounts drew"
sed -n '/## 1. Position/,/## 2./p' DashBoard/README.md
