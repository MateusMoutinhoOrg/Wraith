# AddTransaction

Records a transaction in the ledger. One category and one account per transaction —
no exceptions, no splits.

## Fields

| Field         | Required | Type   | Description                                            |
| ------------- | :------: | ------ | ------------------------------------------------------ |
| `name`        |    ✅    | string | Must be `AddTransaction`                               |
| `date`        |    ✅    | date   | `YYYY-MM-DD`, inside the current month                 |
| `description` |    ✅    | string | e.g. `Supermarket`, `Client payment`                   |
| `category`    |    ✅    | string | Existing category (see [`../../DashBoard/Categories.md`](../../DashBoard/Categories.md)) |
| `account`     |    ✅    | string | Account ID: `BANK`, `CASH`, `CARD`, …                  |
| `value`       |    ✅    | number | Negative = expense, positive = income                  |
| `apply`       |    ✅    | bool   | Set `true` to execute on the next tick                 |

## Renders

`Month/Statement.md`, `Month/Accounts/<Account>.md`, `Month/DashBoard.md`, `Budget.md`,
`Accounts.md`, `README.md`

## Errors

- Unknown `category` or `account` → `Error.md`
- Sign not allowed by the category (`positive`/`negative` flags) → `Error.md`
- `date` outside the current month → `Error.md` (use `CloseMonth` first)

> Card purchases count as expense on the purchase date; the cash leaves the bank only on the
> bill due date. Transfers between own accounts are **not** transactions — use `AddTransfer`.
