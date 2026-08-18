# RemoveAccount

## Description
Removes an account from the registry by its name.

---

## Fields

| Field       | Required | Type   | Description                                          |
| ----------- | :------: | ------ | ---------------------------------------------------- |
| `name`      |    ✅     | string | Must be `RemoveAccount`                              |
| `account`   |    ✅     | string | Display name, e.g. `Vacation savings`                |
| `apply`     |    ✅     | bool   | Set `true` to execute on the next tick               |

---

## Errors

- Account not found → `Error.md`
- Unknown `type` → `Error.md`

---

## Sample

```yaml
name: RemoveAccount
account: Vacation savings
apply: true
```
