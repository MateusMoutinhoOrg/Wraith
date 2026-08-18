# `api.Lib.PerformFullTick`

**Type:** Field

## Signature

```go
PerformFullTick func() error
```

## Description

Runs one whole tick of the state machine: [`PerformTaskTick`](/docs/References/PublicApi/api.PerformTaskTick.md) first, then [`PerformVisualizationTick`](/docs/References/PublicApi/api.PerformVisualizationTick.md).

The order matters and only goes one way. A failed task stops the tick **before** anything is rendered, so the pages on disk always describe data that actually exists — the last state that was written successfully, never a half-applied one.

It is what `wraith tick` calls, and what `wraith watch` calls on every interval.

## Returns

| Type | Description |
| :--- | :--- |
| `error` | The first failure of the two halves. `nil` when the whole tick succeeded. |
