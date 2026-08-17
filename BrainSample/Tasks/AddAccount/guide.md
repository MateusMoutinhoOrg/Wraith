# AddAccount

Adds an account to the registry. Active accounts with `statement: true` get an isolated
monthly statement under `Month/Accounts/`.

## Fields

| Field       | Required | Type   | Description                                             |
| ----------- | :------: | ------ | ------------------------------------------------------- |
| `name`      |    ✅    | string | Must be `AddAccount`                                    |
| `id`        |    ✅    | string | Short unique ID in caps, e.g. `BANK`, `SAVE2`           |
| `account`   |    ✅    | string | Display name, e.g. `Vacation savings`                   |
| `type`      |    ✅    | string | `Checking`, `Cash`, `Credit card` or `Savings`          |
| `opening`   |    ✅    | number | Opening balance in R$ (credit card: negative = owed)    |
| `statement` |    ❌    | bool   | Create `Month/Accounts/<Name>.md`. Default `true`       |
| `apply`     |    ✅    | bool   | Set `true` to execute on the next tick                  |

## Renders

`Accounts.md`, `README.md`; creates `Month/Accounts/<Name>.md` when `statement: true`

## Errors

- `id` already exists → `Error.md`
- Unknown `type` → `Error.md`
