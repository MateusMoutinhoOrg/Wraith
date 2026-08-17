# AddGoal

Adds a financial goal to `Net-Worth.md` §5. Progress bars are rendered from
`current / target` on every tick.

## Fields

| Field      | Required | Type   | Description                              |
| ---------- | :------: | ------ | ---------------------------------------- |
| `name`     |    ✅    | string | Must be `AddGoal`                        |
| `goal`     |    ✅    | string | Goal description — unique                |
| `target`   |    ✅    | number | Target value in R$ (or % for ratio goals)|
| `current`  |    ❌    | number | Starting value. Default `0`              |
| `deadline` |    ✅    | date   | `YYYY-MM`                                |
| `apply`    |    ✅    | bool   | Set `true` to execute on the next tick   |

## Renders

`Net-Worth.md` (§5)

## Errors

- Goal already exists → `Error.md`
- `target` ≤ 0 → `Error.md`

## Sample

```yaml
name: AddGoal
goal: Net worth R$ 20,000
target: 20000
current: 10000
deadline: 2027-06
apply: true
```
