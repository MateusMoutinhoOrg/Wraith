# AddReallocation

Moves budget from one category to another inside the current month. This is the only way a
category that exceeded its limit can unfreeze before the 1st.

## Fields

| Field    | Required | Type   | Description                                |
| -------- | :------: | ------ | ------------------------------------------ |
| `name`   |    ✅    | string | Must be `AddReallocation`                  |
| `amount` |    ✅    | number | R$ moved between limits                    |
| `from`   |    ✅    | string | Category that gives up budget              |
| `to`     |    ✅    | string | Category that receives budget              |
| `reason` |    ✅    | string | Why — shown in `Budget.md` §2              |
| `apply`  |    ✅    | bool   | Set `true` to execute on the next tick     |

## Renders

`Budget.md`, `Month/DashBoard.md`, `Categories.md`

## Errors

- Unknown `from` or `to` category → `Error.md`
- `from` has less remaining budget than `amount` → `Error.md`

## Sample

```yaml
name: AddReallocation
amount: 80
from: Study
to: Vices
reason: Cover the overspend
apply: true
```
