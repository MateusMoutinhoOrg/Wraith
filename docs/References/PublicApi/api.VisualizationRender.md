# `api.VisualizationRender`

**Type:** Struct

## Definition

```go
type VisualizationRender struct {
	Path    string
	Content []byte
}
```

## Description

One file a [visualization](/docs/References/PublicApi/api.Visualizer.md) produced: where it goes, and what it holds. A renderer hands these back rather than writing them, so the decision to put bytes on disk belongs to the caller.

`Path` is relative to the entry's `dest` and slash-separated — `Months/2026-08/Statement.md`. A visualization whose `Folder` is `false` returns a single render with an **empty** `Path`, because `dest` is the file itself.

## Fields

| Field | Description |
| :--- | :--- |
| `Path string` | The file's path below `dest`. Empty for a file visualization. |
| `Content []byte` | The bytes to write. |
