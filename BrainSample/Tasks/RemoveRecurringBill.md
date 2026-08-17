# RemoveRecurringBill

Removes a recurring bill from the registry (e.g. a cancelled subscription). Past transactions
are untouched.

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `RemoveRecurringBill`          |
| `bill`  |    ✅    | string | Exact bill name as shown in `Budget.md` §4 |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

`Budget.md` (§4), `Month/DashBoard.md` (§4 upcoming movements)

## Errors

- Bill does not exist → `Error.md`

## Sample

```yaml
name: RemoveRecurringBill
bill: Streaming
apply: true
```
