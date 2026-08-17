# Brain Workflow

This document describes the available `brain` commands and the execution workflow.

## Commands

### Watch

```bash
./brain watch --time 1s
```

Runs a [tick](#tick) every `<time>` interval (e.g. `1s`).

### Tick

```bash
./brain tick
```

Performs a single tick of the state machine: executes the pending actions and renders the visualization elements.

## Tick Workflow

For each project, the tick performs the following steps:

1. Attempt to read `<Project>/Task.yaml`.
   - If the file does not exist → [Stop Execution](#stop-execution).
2. Validate `Task.yaml.name`.
   - If it does not contain a valid action → [Show Error](#show-error) and [Stop Execution](#stop-execution).
3. Check `Task.yaml.apply`.
   - If `apply == false` → [Stop Execution](#stop-execution).
     No error is shown, since this is not an error condition.
4. Execute the action defined in `Task.yaml.name`.
   - If the action fails → [Show Error](#show-error) and [Stop Execution](#stop-execution).
5. Render all Markdown files in `<Project>/Dashboard`.

## Procedures

### Show Error

Creates an `Error.md` file containing information about the error.

### Stop Execution

1. Verify that `<Project>/Task.yaml` exists.
   - If it does not exist, create it with default values.
2. Set `apply` to `false` in `<Project>/Task.yaml`.
