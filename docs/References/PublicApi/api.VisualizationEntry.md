# `api.VisualizationEntry`

**Type:** Struct

## Definition

```go
type VisualizationEntry struct {
	Name    string
	Dest    string
	Args    map[string]any
	Enabled bool
}
```

## Description

One line of `Visualization.yaml`, decoded: which [visualization](/docs/References/PublicApi/api.Visualizer.md) is asked for, where it writes, and what it writes with. What is not in that file is not rendered, so this is the only thing standing between a renderer and a page on disk.

`Dest` is relative to the vault root and must stay inside it. Two entries may not write to the same place, and one `Dest` may not sit inside another's — the second render would erase part of the first.

## Fields

| Field | Description |
| :--- | :--- |
| `Name string` | The visualization asked for. |
| `Dest string` | Where it writes, relative to the vault root. Missing folders are created. |
| `Args map[string]any` | Per-entry options, overriding the catalog defaults. |
| `Enabled bool` | Whether a tick renders it. Defaults to `true`. |
