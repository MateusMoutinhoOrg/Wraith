# RemoveTransaction

Removes a transaction from the ledger (e.g. a duplicate or a typo). The transaction is
matched by `date` + `description` + `value`.

## Fields

| Field         | Required | Type   | Description                              |
| ------------- | :------: | ------ | ---------------------------------------- |
| `name`        |    ✅    | string | Must be `RemoveTransaction`              |
| `date`        |    ✅    | date   | Date of the transaction to remove        |
| `description` |    ✅    | string | Exact description as shown in the ledger |
| `value`       |    ✅    | number | Exact value, with sign                   |
| `apply`       |    ✅    | bool   | Set `true` to execute on the next tick   |

## Renders

Same set as `AddTransaction` — all balances and budget usage are recalculated.

## Errors

- No transaction matches → `Error.md`
- More than one transaction matches → `Error.md` (nothing is removed; disambiguate manually)
