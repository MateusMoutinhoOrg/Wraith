# RemoveGoal

Removes a goal from `Net-Worth.md` §5 (achieved or abandoned).

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `RemoveGoal`                   |
| `goal`  |    ✅    | string | Exact goal description                 |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

`Net-Worth.md` (§5)

## Errors

- Goal does not exist → `Error.md`

## Sample

```yaml
name: RemoveGoal
goal: Net worth R$ 20,000
apply: true
```
