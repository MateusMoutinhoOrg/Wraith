# Render

Re-renders every Markdown file under `DashBoard/` from the current database state, without
executing any action. Useful after an update of the `brain` binary or if a dashboard file was
edited or deleted by hand (they are generated — hand edits are overwritten).

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `Render`                       |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

**Everything** — all files under `DashBoard/`.

## Errors

None expected; rendering failures produce `Error.md`.

## Sample

```yaml
name: Render
apply: true
```
