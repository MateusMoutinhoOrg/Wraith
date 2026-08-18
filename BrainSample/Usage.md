# Wraith

Wraith is a data-visualization system that applies actions to a database and renders the
current state of that database through several visualization dashboards.

## Table of Contents

- [Overview](#overview)
- [Commands](#commands)
  - [watch](#watch)
  - [tick](#tick)
  - [run](#run)
- [Tick Workflow](#tick-workflow)
- [Procedures](#procedures)

## Overview

Wraith works as a small state machine driven by a single file, `Task.yaml`:

1. A task is declared in `Task.yaml` (or passed directly through the `run` command).
2. Wraith executes the task against the database.
3. Every Markdown file inside `DashBoard/` is re-rendered with the new state.

## Commands

### watch

```bash
./wraith watch --time 1s
```

Runs a [tick](#tick) every `<time>` interval.

| Flag | Description | Example |
| --- | --- | --- |
| `--time` | Interval between ticks | `1s`, `500ms`, `2m` |

### tick

```bash
./wraith tick
```

Performs a single tick of the state machine: executes the pending action declared in
`Task.yaml` and renders the visualization elements. See [Tick Workflow](#tick-workflow).

### run

```bash
./wraith run <task-name> [entries...]
```

Runs a task directly from the command line, without editing `Task.yaml`. This is the
preferred way to drive Wraith programmatically (scripts, automations, other tools).

Example:

```bash
./wraith run AddTransaction --amount 100 --date 2026-08-18 --description test
```

## Tick Workflow

1. Attempt to read `Task.yaml`.
   - If the file does not exist → [Stop Execution](#stop-execution).
2. Validate `Task.yaml.name`.
   - If it does not contain a valid action → [Show Error](#show-error) and
     [Stop Execution](#stop-execution).
3. Check `Task.yaml.apply`.
   - If `apply == false` → [Stop Execution](#stop-execution).
     No error is shown, since this is not an error condition.
4. Execute the action defined in `Task.yaml.name`.
   - If the action fails → [Show Error](#show-error) and [Stop Execution](#stop-execution).
5. Render all Markdown files in `DashBoard/`.

## Procedures

### Show Error

Creates an `Error.md` file containing information about the error.

### Stop Execution

1. Verify that `Task.yaml` exists.
   - If it does not exist, create it with default values.
2. Set `apply` to `false` in `Task.yaml`.
