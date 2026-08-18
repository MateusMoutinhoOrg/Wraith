# `api.Lib.PerformTaskTick`

**Type:** Field

## Signature

```go
PerformTaskTick func() error
```

## Description

Runs the task half of a tick: read `Lib.TaskPath`, decide whether there is anything to do, do it, and disarm the file either way.

Two outcomes are deliberately **not** errors:

- **No task file at all.** A vault that has not been started yet is not a broken vault.
- **`apply: false`.** That is an action waiting to be armed, which is the normal state between two edits.

Otherwise the task runs. On failure the error is written to `Error.md` in the vault, the task file is disarmed, and the error is returned — so an action that cannot work is not retried on every tick of a `watch` loop.

## Returns

| Type | Description |
| :--- | :--- |
| `error` | The failure the task reported, or a task file that could not be read. `nil` when the task succeeded, and when there was nothing to do. |
