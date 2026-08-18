# Tasks Specification

## Description
Defines the required shape of a task file in `sandbox/Tasks/Tasks/<Name>.go` — one action the brain can perform. A task is a **value**: it declares what it is called, what it accepts, and the closure that runs it. The renderers that show the result are governed by [Visualizations](/docs/References/Specs/Visualizations/Specs.md) instead; the workflow of adding one is [HandleTasks.md](/docs/Tutorials/HandleTasks.md).

### Rules
- One task per file, in package `tasks`, named after the task: `AddNote.go` declaring `func AddNote() api.Task`.
- The function takes no arguments and returns a fully populated `api.Task`. It reads no state and touches nothing — building the value is all it does.
- `Name` matches the file name and the function name, so a task can be found from any of the three.
- Every field the task accepts is declared in `Fields`, with its type, whether it is required, and a `Description` a beginner can act on. A field the task reads but did not declare is rejected before `HandleAction` is reached; a field it declared but never reads is dead weight.
- `HandleAction` writes through `args.DataBase` and **nothing else**. It must not touch a file, a socket, or the terminal, and must not read the clock except through `args.Deps.Now`.
- `HandleAction` validates every field **before** writing anything: a task that fails must leave the registries exactly as it found them.
- Errors are returned, never printed, and are written for the person who typed the field — name the field, say what was expected, and show what was given.
- Checks more than one task repeats belong in `shared.go` beside them, so every task reports the same failure in the same words.
- Adding, renaming, or deleting a task requires updating `TaskArray` in [sandbox/Tasks/run.go](/sandbox/Tasks/run.go) in the same commit. Nothing else has to learn about it: the command line, the tick and the `Task-List` visualization all read that array.

## Structure
1. **Package clause**: `package tasks`.
2. **Imports**: `sandbox/contracts/api`, plus the internal packages the action needs — usually `sandbox/lib/store`, and `sandbox/lib/entries` when a field is optional.
3. **Doc comment**: what the task is for, in the words its user would use, and any rule that is not obvious from the fields.
4. **Constructor**: `func <Name>() api.Task` returning the value, with `Fields` listed in the order the guide documents them and `HandleAction` last.
5. **Helpers** *(optional)*: unexported functions this task alone needs, below the constructor.

> **Note**: For a concrete example, refer to [sample.go](./sample.go).
