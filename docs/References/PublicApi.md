# Public API

## Description
Index of every public-facing entry of the library, grouped by its role in the system. Callers hold **structs of function fields** declared in `sandbox/contracts/api` and `sandbox/contracts/deps`; the **factories** that fill those fields live inside `sandbox/` and are unreachable from outside it. See [StructContracts.md](/docs/References/StructContracts.md).

---

## Entry Points

### [lib.New](/docs/References/PublicApi/lib.New.md)
Injects a `deps.Deps` and a database path into the library, and returns an `api.Lib`.

### [standard.New](/docs/References/PublicApi/standard.New.md)
Creates a `deps.Deps` using the standard adapter — a real clock and sleep, a Keep database on disk, the filesystem, and the assets compiled into the binary.

---

## The State Machine

### [api.Lib](/docs/References/PublicApi/api.Lib.md)
The library entry point, returned by `lib.New`. A state machine over a folder: a task changes the data, and every visualization renders it.

### [api.Lib.PerformTask](/docs/References/PublicApi/api.PerformTask.md)
Runs one task by name against the database.

### [api.Lib.PerformVisualization](/docs/References/PublicApi/api.PerformVisualization.md)
Renders one visualization by name and hands back the files it produced, writing nothing.

### [api.Lib.PerformTaskTick](/docs/References/PublicApi/api.PerformTaskTick.md)
Runs the task half of a tick: read the task file, apply what it holds, disarm it.

### [api.Lib.PerformVisualizationTick](/docs/References/PublicApi/api.PerformVisualizationTick.md)
Renders every enabled entry of the config and writes each one to its destination.

### [api.Lib.PerformFullTick](/docs/References/PublicApi/api.PerformFullTick.md)
Runs one whole tick — the task first, then every visualization.

### [api.Lib.Start](/docs/References/PublicApi/api.Start.md)
Writes a default task file and config, creating a vault where there was none.

### [api.Lib.Sandboxmain](/docs/References/PublicApi/api.Sandboxmain.md)
Runs the whole command-line interface over an argument vector and returns the process exit code.

---

## Declarations

### [api.Task](/docs/References/PublicApi/api.Task.md)
One action the brain can perform: its name, its fields, and the closure that runs it.

### [api.Visualizer](/docs/References/PublicApi/api.Visualizer.md)
One renderer the brain can write: its name, its args, and the closure that renders it.

### [api.Field](/docs/References/PublicApi/api.Field.md)
One entry a task or a visualization declares — the source of its validation, its flag, and its documentation.

### [api.HandleActionArgs](/docs/References/PublicApi/api.HandleActionArgs.md)
Everything a task is handed when it runs, and nothing more.

### [api.HandleVisualizationArgs](/docs/References/PublicApi/api.HandleVisualizationArgs.md)
Everything a visualization is handed when it renders.

### [api.VisualizationRender](/docs/References/PublicApi/api.VisualizationRender.md)
One file a visualization produced: where it goes, and what it holds.

### [api.VisualizationEntry](/docs/References/PublicApi/api.VisualizationEntry.md)
One line of `Visualization.yaml`, decoded.

### [api.StartOptions](/docs/References/PublicApi/api.StartOptions.md)
The choices a created vault is written with — the argument `api.Lib.Start` takes.

---

## Dependency Contracts

### [deps.Deps](/docs/References/PublicApi/deps.Deps.md)
The dependency contract every adapter must satisfy — one function field per injectable behavior.

### [verbdeps.Lib](/docs/References/PublicApi/verbdeps.Lib.md)
The sandbox's copy of the embedded Verb argv-parser library.

### [keepdeps.Lib](/docs/References/PublicApi/keepdeps.Lib.md)
The sandbox's copy of the embedded Keep schema-database library.

### [embeddeps.Lib](/docs/References/PublicApi/embeddeps.Lib.md)
The sandbox's copy of the embedded-asset library.

### [iodeps.Lib](/docs/References/PublicApi/iodeps.Lib.md)
The sandbox's copy of the filesystem library — how a tick reads its files and writes its pages.

### [requestdeps.Request](/docs/References/PublicApi/requestdeps.Request.md)
One HTTP request under construction, handed back by `deps.Deps.NewRequest`.

### [requestdeps.Response](/docs/References/PublicApi/requestdeps.Response.md)
One HTTP response, handed back by `requestdeps.Request.Fetch`.
