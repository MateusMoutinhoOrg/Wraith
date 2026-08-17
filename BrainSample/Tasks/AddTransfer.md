# AddTransfer

Records a movement between your own accounts. Transfers are **never** income or expense and
never consume budget — counting a card payment as an expense would charge the purchases twice.

## Fields

| Field   | Required | Type   | Description                                    |
| ------- | :------: | ------ | ---------------------------------------------- |
| `name`  |    ✅    | string | Must be `AddTransfer`                          |
| `date`  |    ✅    | date   | `YYYY-MM-DD`                                   |
| `from`  |    ✅    | string | Source account ID, e.g. `BANK`                 |
| `to`    |    ✅    | string | Destination account ID, e.g. `SAVE`, `CARD`    |
| `value` |    ✅    | number | Positive amount moved                          |
| `note`  |    ❌    | string | e.g. `July bill, paid in full`                 |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick         |

## Common uses

| Movement            | `from` | `to`   | Shown in the ledger as |
| ------------------- | ------ | ------ | ---------------------- |
| Pay the card bill   | `BANK` | `CARD` | `Card-Payment`         |
| Monthly reserve     | `BANK` | `SAVE` | `Reserve`              |
| Invest              | `BANK` | `INV`  | `Invest`               |

## Renders

`Month/Statement.md` (§4 Transfers), both account statements, `Accounts.md`, `README.md`

## Errors

- Unknown `from` or `to` account → `Error.md`
- `from` equals `to`, or `value` ≤ 0 → `Error.md`

## Sample

```yaml
name: AddTransfer
date: 2026-08-30
from: BANK
to: SAVE
value: 1500
note: Monthly reserve
apply: true
```
