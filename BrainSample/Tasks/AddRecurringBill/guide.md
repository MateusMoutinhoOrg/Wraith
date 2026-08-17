# AddRecurringBill

Registers a recurring monthly bill. Recurring bills appear in `Budget.md` §4 and feed the
projections in `Month/DashBoard.md` (upcoming movements).

## Fields

| Field      | Required | Type   | Description                                 |
| ---------- | :------: | ------ | ------------------------------------------- |
| `name`     |    ✅    | string | Must be `AddRecurringBill`                  |
| `bill`     |    ✅    | string | Bill name — unique, e.g. `Internet`         |
| `category` |    ✅    | string | Existing expense category                   |
| `day`      |    ✅    | number | Day of the month (1–31) the bill is charged |
| `value`    |    ✅    | number | Positive amount in R$                       |
| `auto`     |    ❌    | bool   | Charged automatically. Default `false`      |
| `note`     |    ❌    | string | Free note, e.g. `cancel candidate`          |
| `apply`    |    ✅    | bool   | Set `true` to execute on the next tick      |

## Renders

`Budget.md` (§4), `Month/DashBoard.md` (§4 upcoming movements)

## Errors

- Bill name already exists → `Error.md`
- Unknown `category`, or `day` outside 1–31 → `Error.md`

> The bill is a projection, not a transaction — the actual charge still enters the ledger via
> `AddTransaction` (or automatically at month close if `auto: true`).
