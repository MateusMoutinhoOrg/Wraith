# `api.Task`

**Type:** Struct

## Definition

```go
type Task struct {
	Name         string
	Description  string
	Fields       []Field
	HandleAction func(args HandleActionArgs) error
}
```

## Description

One action a brain can perform. A task is a **value, not a function**: it carries its name, what it is for, the fields it accepts, and the closure that runs it.

Declaring the fields rather than parsing them is what makes the rest free: one validator checks every task, the command line gets one `--flag` per field, and the `Task-List` visualization documents a task it has never heard of. A task that declares a field it does not use, or uses one it did not declare, is caught the first time it runs.

`HandleAction` is handed a [`HandleActionArgs`](/docs/References/PublicApi/api.HandleActionArgs.md) and may write **only** to the database inside it. It has no file, no terminal, and no clock beyond `Deps.Now` — which is what makes the same task safe to run from a tick, from the command line, or from a test.

Tasks live one per file under [sandbox/Tasks/Tasks/](/sandbox/Tasks/Tasks/) and are registered in `TaskArray`; adding one is [HandleTasks.md](/docs/Tutorials/HandleTasks.md).

## Fields

| Field | Description |
| :--- | :--- |
| `Name string` | The identifier — the value `Task.yaml.name` carries, and the word `wraith run` takes. |
| `Description string` | What the task does, as the listings and the generated guide show it. |
| [`Fields []Field`](/docs/References/PublicApi/api.Field.md) | The entries it accepts, in the order they are documented. |
| `HandleAction func(HandleActionArgs) error` | Runs the task. Validate first, write last: a failure must change nothing. |
