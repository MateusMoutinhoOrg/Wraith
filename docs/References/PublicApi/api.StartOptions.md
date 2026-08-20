# `api.StartOptions`

**Type:** Struct

## Definition

```go
type StartOptions struct {
	VisualizationArgs map[string]map[string]any
}
```

## Description

The choices a caller creating a vault makes about the vault it creates — the argument [`api.Lib.Start`](/docs/References/PublicApi/api.Start.md) takes. Every field is optional: the zero value writes the defaults compiled into the binary exactly as they are, which is what a caller that only wants a vault passes.

`VisualizationArgs` is keyed by the visualization an entry of the created `Visualization.yaml` asks for, and carries the `args:` block that entry is written with. An arg the map does not name keeps the value the default config carries, and an entry it does not name is written untouched — so overriding one number never rewrites the rest of the file.

It is written into the file, not applied to one render: the created config is a file the caller owns from then on, and editing that file is how those args are changed afterwards. Nothing is written at all if the config is already on disk, since `Start` never overwrites.

The command-line interface fills it from [`wraith start`](/docs/References/Commands.md#start)'s `--prev-months`, `--future-months` and `--current-month`, all three under the `DashBoard` key.

## Fields

| Field | Description |
| :--- | :--- |
| `VisualizationArgs map[string]map[string]any` | The `args:` written into the created config, keyed by visualization name. |
