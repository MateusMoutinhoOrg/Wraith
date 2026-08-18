# `api.Lib.PerformVisualizationTick`

**Type:** Field

## Signature

```go
PerformVisualizationTick func() error
```

## Description

Renders every enabled entry of `Lib.VisualizationPath` and writes each one under its `dest`.

The config is validated **in full** before a single file is written — known names, destinations inside the vault, no two entries colliding or nesting — so a broken entry leaves the vault exactly as it was rather than half rewritten. Entries render in order, and every one of them reads the same data, so no entry can see another's output.

A missing config file is created from the default compiled into the binary and then read, so a vault always renders something rather than reporting a missing file. A folder `dest` is written **into**, never emptied.

## Returns

| Type | Description |
| :--- | :--- |
| `error` | A config that could not be read or validated, a visualization that failed, or a file that could not be written. |
