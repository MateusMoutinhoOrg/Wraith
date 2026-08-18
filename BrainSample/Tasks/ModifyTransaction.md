# ModifyTransaction

Modifies an existing transaction in the ledger. Any provided optional field overwrites the current value; omitted fields remain unchanged.

## Fields

| Field          | Required | Type   | Description                                                                  |
| -------------- | :------: | ------ | ---------------------------------------------------------------------------- |
| `name`         |    ✅     | string | Must be `ModifyTransaction`                                                  |
| `id`           |    ✅     | number | Unique identifier of the transaction to modify                               |
| `account`      |    ❌     | string | New account for the transaction                                              |
| `category`     |    ❌     | string | New category for the transaction                                             |
| `description`  |    ❌     | string | New description of the transaction                                           |
| `amount`       |    ❌     | number | New amount (positive for income, negative for expense)                       |
| `date`         |    ❌     | string | New transaction date (YYYY-MM-DD)                                            |
| `payment_date` |    ❌     | string | New payment/receipt date (YYYY-MM-DD). If null, considers the same as `date` |
| `apply`        |    ✅     | bool   | Set `true` to execute on the next tick                                       |

## Errors

- `id` not found → `Error.md`
- `account` not found → `Error.md`
- `category` not found → `Error.md`
- Invalid date format → `Error.md`

## Sample

```yaml
name: ModifyTransaction
id: 1
account: Savings Account
category: Food
description: Updated grocery shopping
amount: -75.00
date: 2026-08-18
payment_date: null
apply: true
```
