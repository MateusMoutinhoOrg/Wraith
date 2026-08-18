# AddAccount

Adds an account to the registry. Active accounts with `statement: true` get an isolated
monthly statement under `Month/Accounts/`.

## Fields

| Field       | Required | Type   | Description                                          |
| ----------- | :------: | ------ | ---------------------------------------------------- |
| `name`      |    ✅     | string | Must be `AddAccount`                                 |
| `account`   |    ✅     | string | Display name, e.g. `Vacation savings`                |
| `type`      |    ✅     | string | `Cash` or `Credit card`                              |
| `opening`   |    ✅     | number | Opening balance in R$ (credit card: negative = owed) |
| `statement` |    ❌     | bool   | Create `Month/Accounts/<Name>.md`. Default `true`    |
| `apply`     |    ✅     | bool   | Set `true` to execute on the next tick               |

## Renders

`Accounts.md`, `README.md`; creates `Month/Accounts/<Name>.md` when `statement: true`

## Errors

- `id` already exists → `Error.md`
- Unknown `type` → `Error.md`

## Sample

```yaml
name: AddAccount
id: SAVE2
account: Vacation savings
type: Savings         # Checking | Cash | Credit card | Savings
opening: 0            # opening balance in R$ (credit cards: amount already owed, negative)
statement: true       # create an isolated monthly statement under Month/Accounts/
apply: true
```
