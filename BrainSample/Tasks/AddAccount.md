# AddAccount

Adds an account to the registry. 

## Fields

| Field       | Required | Type   | Description                                          |
| ----------- | :------: | ------ | ---------------------------------------------------- |
| `name`      |    ✅     | string | Must be `AddAccount`                                 |
| `account`   |    ✅     | string | Display name, e.g. `Vacation savings`                |
| `opening`   |    ✅     | number | Opening balance in R$ (credit card: negative = owed) |
| `apply`     |    ✅     | bool   | Set `true` to execute on the next tick               |


## Errors

- `account` already exists → `Error.md`
- Unknown `type` → `Error.md`

## Sample

```yaml
name: AddAccount
account: Vacation savings
opening: 0            # opening balance in R$ (credit cards: amount already owed, negative)
apply: true
```
