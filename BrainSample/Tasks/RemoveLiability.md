# RemoveLiability

Removes a debt from the net-worth sheet (paid off or renegotiated into a new one).

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `RemoveLiability`              |
| `debt`  |    ✅    | string | Exact debt name                        |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

`Net-Worth.md`, `README.md`

## Errors

- Debt does not exist → `Error.md`

## Sample

```yaml
name: RemoveLiability
debt: Car loan
apply: true
```
