# CloseMonth

Closes the given month. Run on day 1 of the following month, after every transaction is in.

What it does, in order:

1. Freezes the month totals and writes them to `Year-Report.md` §2.
2. Charges pending recurring bills marked `auto: true` that were not recorded manually.
3. Resets `Month/Statement.md` and every `Month/Accounts/*.md` for the new month,
   carrying the closing balances as the new opening balances.
4. Applies budget-limit reviews (new categories, pending renames/merges).
5. Clears completed reallocations in `Budget.md`.

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `CloseMonth`                   |
| `month` |    ✅    | string | Month to close, `YYYY-MM`              |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

**Everything** — all files under `DashBoard/`.

## Errors

- `month` is not the current open month → `Error.md`
- Month already closed → `Error.md`

> ⚠️ Closing a month is not reversible by a task. Make sure the ledger is complete first.

## Sample

```yaml
name: CloseMonth
month: 2026-08
apply: true
```
