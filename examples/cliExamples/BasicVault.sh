rm -rf WraithSample/
mkdir WraithSample
cd WraithSample
wraith start


#!/bin/bash

echo "Inicializando o cenário..."

# Criando categorias
echo "Criando categorias..."
wraith run AddCategory --category "Salário" --description "Recebimentos de salário" --revenues true --expenses false
wraith run AddCategory --category "Eletrônicos" --description "Compras de eletrônicos" --revenues false --expenses true
wraith run AddCategory --category "Alimentação" --description "Gastos com comida" --revenues false --expenses true
wraith run AddCategory --category "Investimentos" --description "Aportes financeiros" --revenues false --expenses false

# Criando contas
echo "Criando contas..."
wraith run AddAccount --account "Conta Corrente" --opening 1000
wraith run AddAccount --account "Corretora" --opening 0
wraith run AddCreditCard --account "Cartão de Crédito" --limit 5000 --closing_day 10 --due_day 15

# Registrando recebimentos
echo "Registrando recebimento de salário..."
wraith run AddTransaction --account "Conta Corrente" --category "Salário" --amount 5000 --date "2026-08-05" --description "Salário de Agosto"

# Registrando investimentos (transferência de Conta Corrente para Corretora)
echo "Registrando investimentos..."
wraith run AddTransaction --account "Conta Corrente" --category "Investimentos" --amount -1000 --date "2026-08-10" --description "Aporte Mensal"
wraith run AddTransaction --account "Corretora" --category "Investimentos" --amount 1000 --date "2026-08-10" --description "Aporte Mensal"

# Registrando despesas à vista
echo "Registrando despesas à vista..."
wraith run AddTransaction --account "Conta Corrente" --category "Alimentação" --amount -150 --date "2026-08-12" --description "Supermercado"

# Registrando despesa parcelada no cartão de crédito
echo "Registrando compra parcelada no cartão..."
wraith run AddTransaction --account "Cartão de Crédito" --category "Eletrônicos" --amount -2400 --date "2026-08-12" --installments 12 --description "Compra de Celular"

echo "Cenário emulado criado com sucesso!"
