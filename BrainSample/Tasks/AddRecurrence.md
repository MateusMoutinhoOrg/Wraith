# AddRecurrence

Declares a commitment that repeats every month — a salary, rent, a subscription, a standing
transfer into savings. It is the only way the registry learns about money that has not been
recorded yet, and therefore the only thing [`DashBoard/Forecast.md`](../DashBoard/Forecast.md)
is allowed to project with.

A recurrence is **not** a transaction. It does not touch any balance and it never appears in a
month's ledger. It is a rule the forecast reads. When the day actually arrives you still record
what happened with [`AddTransaction`](AddTransaction.md) — with the real amount, which is rarely
the declared one to the cent.

## Fields

| Field         | Required | Type   | Description                                                                          |
| ------------- | :------: | ------ | ------------------------------------------------------------------------------------ |
| `name`        |    ✅     | string | Must be `AddRecurrence`                                                              |
| `description` |    ✅     | string | Identifies the recurrence — must be unique, it is how `RemoveRecurrence` addresses it |
| `account`     |    ✅     | string | The account the money leaves from or arrives in                                       |
| `to_account`  |    ❌     | string | Destination account. Only valid when `category` is a transfer category               |
| `category`    |    ✅     | string | The category of the recurrence                                                       |
| `amount`      |    ✅     | number | Amount per occurrence (positive = income, negative = expense)                         |
| `day`         |    ✅     | number | Day of the month it falls on (1-31)                                                  |
| `start`       |    ✅     | string | First month it applies (`YYYY-MM`)                                                   |
| `end`         |    ❌     | string | Last month it applies (`YYYY-MM`). `null` or omitted = open-ended                     |
| `apply`       |    ✅     | bool   | Set `true` to execute on the next tick                                               |

### `day` on short months

A `day` larger than the month has is clamped to that month's last day: `day: 30` falls on
28-feb in a common year, 29-feb in a leap year, and 30-apr. It is never skipped and never
spills into the next month.

### `to_account` — recurring transfers

A one-off transfer is two `AddTransaction`s sharing a transfer category (see
[`Help/Task.md` §3](../Help/Task.md)). A recurring one is a single `AddRecurrence` with
`to_account` filled in: the forecast expands it into both legs, `-amount` on `account` and
`+amount` on `to_account`, so the pair nets to zero and is counted as neither income nor expense.

`to_account` is rejected when `category` is not a transfer category — a salary has no
destination account.

## Errors

- `description` already used by another recurrence → `Error.md`
- `account` or `to_account` not found → `Error.md`
- `category` not found → `Error.md`
- Positive `amount` in a non-transfer category whose `revenues` is `false` → `Error.md`
- Negative `amount` in a non-transfer category whose `expenses` is `false` → `Error.md`
- Non-positive `amount` on a transfer category — a transfer is declared once, as the amount that
  leaves `account` → `Error.md`
- `to_account` given on a non-transfer category → `Error.md`
- `to_account` missing on a transfer category → `Error.md`
- `to_account` equal to `account` → `Error.md`
- `day` outside 1-31 → `Error.md`
- Invalid `start` / `end` format, or `end` earlier than `start` → `Error.md`

## What you do not declare

The card bill is **not** a recurrence. Its amount changes every month, so the forecast derives it
from the card registry instead: recurring purchases and installments landing inside the billing
window add up to the bill that closes on `closing_day`, and the whole of it leaves the paying
account on `due_day`. See [`Forecast.md` §4](../DashBoard/Forecast.md).

## Sample

```yaml
name: AddRecurrence
description: Client A retainer   # unique — this is the recurrence's name
account: Bank
category: Freelance
amount: 2000                     # positive = income
day: 11                          # falls on the 11th of every month
start: 2026-09                   # first month it applies
end: null                        # open-ended
apply: true
```

A subscription charged to the credit card:

```yaml
name: AddRecurrence
description: Streaming
account: Nubank Card
category: Leisure
amount: -55
day: 30                          # clamped to the 28th in february
start: 2026-09
end: null
apply: true
```

A standing transfer into savings that stops after six months:

```yaml
name: AddRecurrence
description: Emergency reserve
account: Bank
to_account: Emergency Savings    # transfer category → both legs are generated
category: Reserve
amount: 500
day: 6
start: 2026-09
end: 2027-02                     # last month it applies
apply: true
```
