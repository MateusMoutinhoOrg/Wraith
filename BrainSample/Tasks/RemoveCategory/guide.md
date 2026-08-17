# RemoveCategory

Removes a category from the system. If the category has transactions in the current month,
they must be migrated to another category via `migrate_to`.

## Fields

| Field        | Required | Type   | Description                                                  |
| ------------ | :------: | ------ | ------------------------------------------------------------ |
| `name`       |    ✅    | string | Must be `RemoveCategory`                                     |
| `category`   |    ✅    | string | Category to remove                                           |
| `migrate_to` |    ❌    | string | Category that inherits existing transactions. Required if the category has transactions this month |
| `apply`      |    ✅    | bool   | Set `true` to execute on the next tick                       |

## Renders

`Categories.md`, `Budget.md`, `Month/DashBoard.md`, `Month/Statement.md`

## Errors

- Category does not exist → `Error.md`
- Category has transactions and `migrate_to` is missing or invalid → `Error.md`

> Rule: renaming or merging categories is done at month close, never mid-month — see
> [`../../DashBoard/Categories.md`](../../DashBoard/Categories.md) §4.
