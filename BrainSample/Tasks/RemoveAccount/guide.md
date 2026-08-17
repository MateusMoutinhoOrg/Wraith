# RemoveAccount

Removes an account from the registry and deletes its monthly statement file.

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `RemoveAccount`                |
| `id`    |    ✅    | string | ID of the account to remove            |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

`Accounts.md`, `README.md`; removes `Month/Accounts/<Name>.md`

## Errors

- Account does not exist → `Error.md`
- Balance is not zero → `Error.md` (transfer the balance out first with `AddTransfer`)
- Account has transactions in the current month → `Error.md` (close the month first)
