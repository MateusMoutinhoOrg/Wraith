# RemoveCreditCard

## Description
Removes a credit card from the registry by its unique ID.

---

## Fields

| Field   | Required | Type   | Description                                  |
| ------- | :------: | ------ | -------------------------------------------- |
| `name`  |    ✅     | string | Must be `RemoveCreditCard`                   |
| `id`    |    ✅     | string | Unique identifier for the card, e.g. `NUBANK`|
| `apply` |    ✅     | bool   | Set `true` to execute on the next tick       |

---

## Errors

- Credit card not found → `Error.md`
- Unknown `type` → `Error.md`

---

## Sample

```yaml
name: RemoveCreditCard
id: NUBANK
apply: true
```
