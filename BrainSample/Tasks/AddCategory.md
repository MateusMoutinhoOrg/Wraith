# AddCategory

Adds a category to the registry. 

## Fields

| Field         | Required | Type   | Description                                          |
| ------------- | :------: | ------ | ---------------------------------------------------- |
| `name`        |    ✅     | string | Must be `AddCategory`                                |
| `category`    |    ✅     | string | Display name, e.g. `Food`                            |
| `description` |    ✅     | string | Description of the category                          |
| `revenues`    |    ✅     | bool   | Set `true` if this category contains revenues        |
| `expenses`    |    ✅     | bool   | Set `true` if this category contains expenses        |
| `parent`      |    ❌     | string | Parent category name, allowing it to be a child      |
| `apply`       |    ✅     | bool   | Set `true` to execute on the next tick               |


## Errors

- `category` already exists → `Error.md`
- Parent category not found → `Error.md`
- Unknown `type` → `Error.md`

## Sample

```yaml
name: AddCategory
category: Food
description: All food-related transactions
revenues: false
expenses: true
parent: null
apply: true
```
