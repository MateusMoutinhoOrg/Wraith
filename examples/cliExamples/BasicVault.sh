rm -rf WraithSample/
mkdir WraithSample
cd WraithSample
wraith start


#!/bin/bash

echo "Initializing the scenario..."

# Creating categories
echo "Creating categories..."
wraith run AddCategory --category "Salary" --description "Salary income" --revenues true --expenses false
wraith run AddCategory --category "Electronics" --description "Electronics purchases" --revenues false --expenses true
wraith run AddCategory --category "Food" --description "Food expenses" --revenues false --expenses true
wraith run AddCategory --category "Investments" --description "Financial contributions" --revenues false --expenses false

# Creating accounts
echo "Creating accounts..."
wraith run AddAccount --account "Checking Account" --opening 1000
wraith run AddAccount --account "Brokerage" --opening 0
wraith run AddCreditCard --account "Credit Card" --limit 5000 --closing_day 10 --due_day 15

# Recording income
echo "Recording salary income..."
wraith run AddTransaction --account "Checking Account" --category "Salary" --amount 5000 --date "2026-08-05" --description "August Salary"

# Recording investments (transfer from Checking Account to Brokerage)
echo "Recording investments..."
wraith run AddTransaction --account "Checking Account" --category "Investments" --amount -1000 --date "2026-08-10" --description "Monthly Contribution"
wraith run AddTransaction --account "Brokerage" --category "Investments" --amount 1000 --date "2026-08-10" --description "Monthly Contribution"

# Recording cash expenses
echo "Recording cash expenses..."
wraith run AddTransaction --account "Checking Account" --category "Food" --amount -150 --date "2026-08-12" --description "Grocery Store"

# Recording installment purchase on credit card
echo "Recording installment purchase on credit card..."
wraith run AddTransaction --account "Credit Card" --category "Electronics" --amount -2400 --date "2026-08-12" --installments 12 --description "Phone Purchase"

echo "Emulated scenario created successfully!"
