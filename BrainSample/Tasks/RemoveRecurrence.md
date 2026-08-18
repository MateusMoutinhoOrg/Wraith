# RemoveRecurrence

## Description
Removes a recurrence from the registry by its `description`.

Removing a recurrence only changes the future: the projection in
[`DashBoard/Forecast.md`](../DashBoard/Forecast.md) stops counting it from the next tick on.
Transactions already recorded are untouched — a recurrence never created any.

To stop a commitment on a known date without losing it from the registry, prefer setting `end`
on the recurrence instead of removing it.

---

## Fields

| Field        | Required | Type   | Description                                                    |
| ------------ | :------: | ------ | -------------------------------------------------------------- |
| `name`       |    ✅     | string | Must be `RemoveRecurrence`                                     |
| `recurrence` |    ✅     | string | The `description` of the recurrence, e.g. `Emergency reserve`   |
| `apply`      |    ✅     | bool   | Set `true` to execute on the next tick                          |

---

## Errors

- Recurrence not found → `Error.md`

---

## Sample

```yaml
name: RemoveRecurrence
recurrence: Emergency reserve
apply: true
```
