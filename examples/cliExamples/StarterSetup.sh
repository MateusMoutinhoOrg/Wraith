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
wraith run AddAccount --account "Bank Account" --opening 0
wraith run AddAccount --account "Wallet" --opening 0

echo
echo "== the vault those accounts drew"
sed -n '/## 1. Position/,/## 2./p' DashBoard/README.md
