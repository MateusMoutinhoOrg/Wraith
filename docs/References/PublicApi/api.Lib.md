# `api.Lib`

**Type:** Struct

## Definition

```go
type Lib struct {
	Deps                     deps.Deps
	DatabasePath             string
	TaskPath                 string
	VisualizationPath        string
	Tasks                    []Task
	Visualizations           []Visualizer
	PerformTask              func(taskName string, entries map[string]any) error
	PerformVisualization     func(name string, entries map[string]any) ([]VisualizationRender, error)
	PerformTaskTick          func() error
	PerformVisualizationTick func() error
	PerformFullTick          func() error
	Start                    func(options StartOptions) error
	Sandboxmain              func(args []string) int
}
```

## Description

The library entry point, returned by [`lib.New`](/docs/References/PublicApi/lib.New.md). It is a **state machine over a folder**: a task changes the data, and every visualization renders it. It is exposed as a struct of function fields — `lib.New` stores the injected [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) and the database path on it, then runs the factories in `sandbox/lib/publicfunctions/`, each of which fills one field with a closure reading the struct at call time. Calling a field reads exactly like calling a method. See [StructContracts.md](/docs/References/StructContracts.md).

The two registries — [`Tasks`](/docs/References/PublicApi/api.Task.md) and [`Visualizations`](/docs/References/PublicApi/api.Visualizer.md) — are read once at construction out of the arrays that declare them, so a caller can ask what the brain is allowed to do rather than hard-coding a list.

`Sandboxmain` is the same idea taken one step further: the whole command-line interface is a field of this struct, so the installed binary in [cmd/main](/cmd/main/) holds no behavior of its own. A Go caller that wants the library rather than the CLI simply never calls it.

`Deps` is exported because the library's own factories read it, but it is **read-only after construction**: the closures already captured the struct they were built over. Patch the `deps.Deps` value before calling `lib.New`.

The three path fields are writable, and are what the interface's `--task`, `--visualization` and `--database` flags overwrite for one invocation.

## Fields

| Field | Description |
| :--- | :--- |
| [`Deps deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | The dependency set injected by `lib.New`; read-only after construction. |
| `DatabasePath string` | The folder inside the vault the registries are persisted in. Required on construction. |
| `TaskPath string` | The task file a tick reads and resets. Defaults to `Task.yaml`. |
| `VisualizationPath string` | The config a tick renders from. Defaults to `Visualization.yaml`. |
| [`Tasks []Task`](/docs/References/PublicApi/api.Task.md) | Every action the binary carries. |
| [`Visualizations []Visualizer`](/docs/References/PublicApi/api.Visualizer.md) | Every renderer the binary carries. |
| [`PerformTask func(...) error`](/docs/References/PublicApi/api.PerformTask.md) | Runs one task by name against the database. |
| [`PerformVisualization func(...) ([]VisualizationRender, error)`](/docs/References/PublicApi/api.PerformVisualization.md) | Renders one visualization by name and hands back the files. |
| [`PerformTaskTick func() error`](/docs/References/PublicApi/api.PerformTaskTick.md) | Runs the task half of a tick, reading and disarming the task file. |
| [`PerformVisualizationTick func() error`](/docs/References/PublicApi/api.PerformVisualizationTick.md) | Renders every enabled entry of the config and writes each to its `dest`. |
| [`PerformFullTick func() error`](/docs/References/PublicApi/api.PerformFullTick.md) | Runs one whole tick: the task first, then every visualization. |
| [`Start func(options StartOptions) error`](/docs/References/PublicApi/api.Start.md) | Writes a default task file and config, without overwriting. |
| [`Sandboxmain func(args []string) int`](/docs/References/PublicApi/api.Sandboxmain.md) | The command-line interface; returns the exit code. |
