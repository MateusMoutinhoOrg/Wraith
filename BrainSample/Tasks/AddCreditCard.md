# AddCreditCard

Adds a credit card to the registry. 

## Fields

| Field         | Required | Type   | Description                                          |
| ------------- | :------: | ------ | ---------------------------------------------------- |
| `name`        |    ✅     | string | Must be `AddCreditCard`                              |
| `account`     |    ✅     | string | Display name, e.g. `Nubank Card`                     |
| `limit`       |    ✅     | number | Total credit limit in R$                             |
| `closing_day` |    ✅     | number | Day of the month the bill closes (1-31)              |
| `due_day`     |    ✅     | number | Day of the month the bill is due (1-31)              |
| `opening`     |    ❌     | number | Amount already owed in R$ (negative number)          |
| `apply`       |    ✅     | bool   | Set `true` to execute on the next tick               |


## Errors

- `account` already exists → `Error.md`
- Invalid `closing_day` or `due_day` → `Error.md`

## Sample

```yaml
name: AddCreditCard
account: Nubank Card
limit: 5000           # total credit limit in R$
closing_day: 25       # day the bill closes
due_day: 5            # day the bill is due
opening: -150.50      # amount already owed (negative)
apply: true
```
