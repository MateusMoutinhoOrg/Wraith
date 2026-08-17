# SetBudget

Sets or updates the spending limit of a category. Setting `limit: 0` removes the limit
(the category shows as ⚪ until reviewed).

## Fields

| Field      | Required | Type   | Description                                  |
| ---------- | :------: | ------ | -------------------------------------------- |
| `name`     |    ✅    | string | Must be `SetBudget`                          |
| `category` |    ✅    | string | Existing expense category                    |
| `limit`    |    ✅    | number | Limit in R$. `0` removes the limit           |
| `period`   |    ❌    | string | `month` (default) or `year`                  |
| `apply`    |    ✅    | bool   | Set `true` to execute on the next tick       |

## Renders

`Budget.md`, `Categories.md`, `Month/DashBoard.md`, `README.md`

## Errors

- Unknown `category` → `Error.md`
- Negative `limit` → `Error.md`

> Rules: limits are reviewed on the 1st of each month; the limit is per category regardless of
> which account paid — see [`../../DashBoard/Budget.md`](../../DashBoard/Budget.md) §5.
