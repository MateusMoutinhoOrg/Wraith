# `api.Field`

**Type:** Struct

## Definition

```go
type Field struct {
	Name        string
	Type        int
	Required    bool
	Description string
	Default     any
}

const (
	TextField = iota
	NumberField
	BoolField
)
```

## Description

One entry a [task](/docs/References/PublicApi/api.Task.md) or a [visualization](/docs/References/PublicApi/api.Visualizer.md) declares. It is the single source for four separate things: the validation a task file is checked against, the `--flag` the command line accepts, the default filled in when the entry is omitted, and the row the generated guide writes.

Values arrive as text from a command line and as typed scalars from a YAML file. The same coercion handles both, so `--amount -32.90` and `amount: -32.90` reach a task identically. `Default` is only read for a field that is not required.

## Fields

| Field | Description |
| :--- | :--- |
| `Name string` | The key in `Task.yaml`, and the `--flag` on the command line. |
| `Type int` | `TextField`, `NumberField`, or `BoolField`. |
| `Required bool` | Whether the task refuses to run without it. |
| `Description string` | What it means, as the generated guide shows it. |
| `Default any` | The value used when it is omitted. Only read for an optional field. |
