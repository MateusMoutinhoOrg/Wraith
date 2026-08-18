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
| `apply`        |    ✅     | bool   | Set `true` to execute on the next tick                                       |

## Errors

- `account` not found → `Error.md`
- `category` not found → `Error.md`
- Invalid date format → `Error.md`

## Sample

```yaml
name: AddTransaction
account: Checking Account
category: Food
description: Grocery shopping
amount: -50.00
date: 2026-08-18
payment_date: null
apply: true
```
