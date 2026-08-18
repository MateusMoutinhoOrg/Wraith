# AddTransaction

Records an income or expense in the ledger.

## Fields

| Field          | Required | Type   | Description                                                                  |
| -------------- | :------: | ------ | ---------------------------------------------------------------------------- |
| `name`         |    ✅     | string | Must be `AddTransaction`                                                     |
| `account`      |    ✅     | string | The account where the transaction occurred                                   |
| `category`     |    ✅     | string | The category of the transaction                                              |
| `description`  |    ❌     | string | Description of the transaction                                               |
| `amount`       |    ✅     | number | Transaction amount (positive for income, negative for expense)               |
| `date`         |    ✅     | string | Transaction date (YYYY-MM-DD)                                                |
| `payment_date` |    ❌     | string | Payment/receipt date (YYYY-MM-DD). If null, considers the same as `date`     |
| `installments` |    ❌     | number | Split `amount` into this many monthly parts (2-72). If null, a single part    |
| `apply`        |    ✅     | bool   | Set `true` to execute on the next tick                                       |

## Installments

`installments: N` records **one purchase as N transactions**, one per month. `amount` is the
total, not the value of each part.

- Part 1 falls on `date`; part *k* falls on the same day of the *k-1*-th following month, clamped
  to that month's last day when it does not exist (a purchase on the 31st lands on 30-apr).
- Each part gets `amount / N`, rounded to cents. Any remainder goes to the **first** part, so the
  parts always add back up to `amount` exactly.
- Every part is a real transaction: it counts in the month it falls in, appears in that month's
  [`Statement.md`](../DashBoard/Months/README.md), and — on a credit card — lands in the billing
  window its date belongs to.
- Because the future parts are already recorded, [`Forecast.md`](../DashBoard/Forecast.md)
  counts them as scheduled facts, not as a projection.

A purchase of R$ 4,200 in 12x on 20-aug-2026 becomes:

| Part  | Date   | Amount    |
| ----- | ------ | --------: |
| 1/12  | 20-aug | -R$ 350   |
| 2/12  | 20-sep | -R$ 350   |
| …     | …      | …         |
| 12/12 | 20-jul | -R$ 350   |

Use [`ModifyTransaction`](ModifyTransaction.md) to change or clear a single part — the other
parts are independent transactions and are not affected.

## Errors

- `account` not found → `Error.md`
- `category` not found → `Error.md`
- Invalid date format → `Error.md`
- `installments` outside 2-72, or not a whole number → `Error.md`
- `installments` combined with `payment_date` — each part settles on its own date → `Error.md`

## Sample

```yaml
name: AddTransaction
account: Checking Account
category: Food
description: Grocery shopping
amount: -50.00
date: 2026-08-18
payment_date: null
installments: null
apply: true
```

Split into twelve monthly parts:

```yaml
name: AddTransaction
account: Nubank Card
category: Business
description: Notebook
amount: -4200.00      # the total, not the value of each part
date: 2026-08-20
installments: 12      # 12 × -R$ 350, from aug-2026 to jul-2027
apply: true
```
