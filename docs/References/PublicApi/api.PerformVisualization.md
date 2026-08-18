# `api.Lib.PerformVisualization`

**Type:** Field

## Signature

```go
PerformVisualization func(name string, entries map[string]any) ([]VisualizationRender, error)
```

## Description

Renders one [visualization](/docs/References/PublicApi/api.Visualizer.md) by name and hands back the files it produced, as [`VisualizationRender`](/docs/References/PublicApi/api.VisualizationRender.md) values.

**Nothing is written to disk.** A caller that means to write takes these files and puts them under a destination it chose — which is what lets the same call serve a tick, a `wraith render`, and a program that only wants the markdown.

The `entries` map is the same one a `Visualization.yaml` entry's `args:` block decodes to. It is validated against the args the visualization declares, and anything omitted falls back to the declared default.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `name` | `string` | The visualization to render. Must be one of `Lib.Visualizations`. |
| `entries` | `map[string]any` | Its args, keyed by the name they carry under `args:`. |

## Returns

| Type | Description |
| :--- | :--- |
| [`[]VisualizationRender`](/docs/References/PublicApi/api.VisualizationRender.md) | One entry per file the visualization produced. |
| `error` | An unknown name, an invalid arg, or a failure the visualization reported. |
