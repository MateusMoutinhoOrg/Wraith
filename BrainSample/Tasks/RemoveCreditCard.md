# RemoveCreditCard

## Description
Removes a credit card from the registry by its name.

---

## Fields

| Field   | Required | Type   | Description                                  |
| ------- | :------: | ------ | -------------------------------------------- |
| `name`    |    ✅     | string | Must be `RemoveCreditCard`                   |
| `account` |    ✅     | string | Display name, e.g. `Nubank Card`             |
| `apply`   |    ✅     | bool   | Set `true` to execute on the next tick       |

---

## Errors

- Credit card not found → `Error.md`
- Unknown `type` → `Error.md`

---

## Sample

```yaml
name: RemoveCreditCard
account: Nubank Card
apply: true
```
