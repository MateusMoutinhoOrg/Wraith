# `api.Visualizer`

**Type:** Struct

## Definition

```go
type Visualizer struct {
	Name             string
	Description      string
	Folder           bool
	Args             []Field
	HandleVisualizer func(args HandleVisualizationArgs) ([]VisualizationRender, error)
}
```

## Description

One renderer a brain can write. Like a [`Task`](/docs/References/PublicApi/api.Task.md) it is a value: its name, what it shows, whether it owns a folder or a single file, the args it accepts, and the closure that renders it.

`HandleVisualizer` is handed a [`HandleVisualizationArgs`](/docs/References/PublicApi/api.HandleVisualizationArgs.md) and **returns files**; it never writes one. Putting bytes on disk happens one layer up, which is what lets the same renderer serve a tick, a `wraith render`, and a caller that only wants the markdown. It reads the database and never writes to it.

`Folder` decides how its renders are placed: `true` writes each one below the entry's `dest`, `false` writes a single render **at** `dest` — a file visualization returns one render with an empty `Path`.

Visualizations live one per file under [sandbox/Visualization/Visualization/](/sandbox/Visualization/Visualization/) and are registered in `Catalog`; adding one is [HandleVisualizations.md](/docs/Tutorials/HandleVisualizations.md).

## Fields

| Field | Description |
| :--- | :--- |
| `Name string` | The identifier — the value a `Visualization.yaml` entry's `name` carries. |
| `Description string` | What it displays, as the listings and the generated catalog show it. |
| `Folder bool` | `true` when it owns a whole tree below `dest`, `false` when it is one file. |
| [`Args []Field`](/docs/References/PublicApi/api.Field.md) | The options it accepts, with their defaults. |
| `HandleVisualizer func(HandleVisualizationArgs) ([]VisualizationRender, error)` | Produces one render per file it writes. |
