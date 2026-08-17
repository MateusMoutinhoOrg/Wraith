# RemoveTransfer

Removes a transfer between own accounts. Matched by `date` + `from` + `to` + `value`.

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `RemoveTransfer`               |
| `date`  |    ✅    | date   | Date of the transfer                   |
| `from`  |    ✅    | string | Source account ID                      |
| `to`    |    ✅    | string | Destination account ID                 |
| `value` |    ✅    | number | Exact amount                           |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

Same set as `AddTransfer` — balances on both accounts are recalculated.

## Errors

- No transfer matches → `Error.md`
- More than one transfer matches → `Error.md` (nothing is removed)
