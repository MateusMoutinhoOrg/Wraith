# `api.HandleVisualizationArgs`

**Type:** Struct

## Definition

```go
type HandleVisualizationArgs struct {
	Deps     deps.Deps
	DataBase keepdeps.KeepDatabase
	Entries  map[string]any
}
```

## Description

Everything a [visualization](/docs/References/PublicApi/api.Visualizer.md) is handed when it renders. It is the same database a task writes, read here and never written to, plus the clock — which is how a page knows which month is the open one.

`Entries` is the visualization's args, already validated and already filled with the declared defaults.

## Fields

| Field | Description |
| :--- | :--- |
| [`Deps deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | The dependency set the library was built with. |
| [`DataBase keepdeps.KeepDatabase`](/docs/References/PublicApi/keepdeps.Lib.md) | The registries it reads. |
| `Entries map[string]any` | Its args, keyed by the name they carry under `args:`. |
