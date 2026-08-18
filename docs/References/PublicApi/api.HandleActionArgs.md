# `api.HandleActionArgs`

**Type:** Struct

## Definition

```go
type HandleActionArgs struct {
	Deps     deps.Deps
	DataBase keepdeps.KeepDatabase
	Entries  map[string]any
}
```

## Description

Everything a [task](/docs/References/PublicApi/api.Task.md) is handed when it runs — and, just as importantly, everything it is **not**. A task reaches storage through `DataBase` and nothing else, which is what keeps it from touching a file, a socket, or the terminal. Whatever it cannot reach, it cannot break.

`Entries` is the task's fields, already validated against its declaration and already filled with the declared defaults, so a task reads its input without checking whether it is there.

## Fields

| Field | Description |
| :--- | :--- |
| [`Deps deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | The dependency set — the clock a record is stamped with, and little else a task needs. |
| [`DataBase keepdeps.KeepDatabase`](/docs/References/PublicApi/keepdeps.Lib.md) | The registries it may write. |
| `Entries map[string]any` | Its fields, keyed by the name they carry in `Task.yaml`. |
