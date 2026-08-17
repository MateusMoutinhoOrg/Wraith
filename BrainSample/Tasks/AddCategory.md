# AddCategory

Adds a new category to the system. Every transaction carries exactly one category, so a
category must exist before any transaction can use it.

## Fields

| Field      | Required | Type   | Description                                                        |
| ---------- | :------: | ------ | ------------------------------------------------------------------ |
| `name`     |    ✅    | string | Must be `AddCategory`                                              |
| `category` |    ✅    | string | Category name — unique, e.g. `Poker`                               |
| `kind`     |    ✅    | string | `Fixed`, `Essential`, `Discretionary`, `Debt service`, `Income` or `Transfer` |
| `positive` |    ✅    | bool   | Category may register positive values (income)                     |
| `negative` |    ✅    | bool   | Category may register negative values (expenses)                   |
| `limit`    |    ❌    | number | Monthly limit in R$. Defaults to `0` (new categories are reviewed on the 1st) |
| `apply`    |    ✅    | bool   | Set `true` to execute on the next tick                             |

## Renders

`Categories.md`, `Budget.md`, `Month/DashBoard.md`

## Errors

- Category already exists → `Error.md`
- `positive` and `negative` both `false` → `Error.md` (category would be unusable)

## Sample

```yaml
name: AddCategory
category: Poker
kind: Discretionary   # Fixed | Essential | Discretionary | Debt service | Income | Transfer
positive: true        # may register income (e.g. poker winnings)
negative: true        # may register expenses
limit: 200            # monthly limit in R$ (0 = no limit yet, reviewed on the 1st)
apply: true
```
