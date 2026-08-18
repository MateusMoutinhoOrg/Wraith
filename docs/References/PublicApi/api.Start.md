# `api.Lib.Start`

**Type:** Field

## Signature

```go
Start func() error
```

## Description

Creates a vault: a task file at `Lib.TaskPath` and a visualization config at `Lib.VisualizationPath`, copied from the defaults compiled into the binary and read through [`deps.Deps.EmbedDeps`](/docs/References/PublicApi/embeddeps.Lib.md).

It **never overwrites**. A file already on disk is left exactly as it is, so running it in a vault that is already going is harmless rather than destructive — which matters, because it is the first thing anyone runs and the easiest thing to run twice.

The defaults themselves live in [assets/start/](/assets/start/); changing what a new vault starts life with is editing those files, not this field.

## Returns

| Type | Description |
| :--- | :--- |
| `error` | A default missing from the binary, or a file that could not be written. |
