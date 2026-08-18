# RemoveCategory

## Description
Removes a category from the registry by its name.

---

## Fields

| Field       | Required | Type   | Description                                          |
| ----------- | :------: | ------ | ---------------------------------------------------- |
| `name`      |    ✅     | string | Must be `RemoveCategory`                             |
| `category`  |    ✅     | string | Display name of the category to remove, e.g. `Food`  |
| `apply`     |    ✅     | bool   | Set `true` to execute on the next tick               |

---

## Errors

- Category not found → `Error.md`
- Unknown `type` → `Error.md`

---

## Sample

```yaml
name: RemoveCategory
category: Food
apply: true
```
