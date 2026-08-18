# Wraith

Wraith is a data-visualization system that applies actions to a database and renders the
current state of that database through several visualization dashboards.

## Table of Contents

- [Overview](#overview)
- [Commands](#commands)
  - [watch](#watch)
  - [tick](#tick)
  - [run](#run)
  - [render](#render)
- [Tick Workflow](#tick-workflow)
- [Visualization](#visualization)
- [Procedures](#procedures)

## Overview

Wraith works as a small state machine driven by two files, `Task.yaml` and `Visualization.yaml`:

1. A task is declared in `Task.yaml` (or passed directly through the `run` command).
2. Wraith executes the task against the database.
3. Every visualization declared in `Visualization.yaml` is re-rendered with the new state.

`Task.yaml` decides what changes; `Visualization.yaml` decides what you get to see.

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
`Task.yaml` and renders every visualization declared in `Visualization.yaml`. See
[Tick Workflow](#tick-workflow).

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

### render

```bash
./wraith render <visualization-name> [args...]
```

Renders a single visualization to standard output, without writing its `dest` file and without
touching the database. Use it to preview an entry — or a different set of `args` — before putting
it into `Visualization.yaml`.

Example:

```bash
./wraith render ForeCast --months 12
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
5. [Render the visualizations](#visualization).

## Visualization

Wraith renders no page on its own. Every file it writes is one entry you declared in
`Visualization.yaml`, so the set of dashboards you see is yours to choose:

```yaml
- name: Usage
  dest: Usage.md

- name: DashBoard
  dest: DashBoard/README.md

- name: ForeCast
  args:
    months: 8
  dest: DashBoard/Forecast.md
```

| Field | Required | Description |
| --- | --- | --- |
| `name` | yes | Which visualization to render — `DashBoard`, `ForeCast`, `Accounts`, `Month`, … |
| `dest` | yes | Where to write it, relative to the vault root. Missing folders are created. |
| `args` | no | Per-visualization options, such as `ForeCast.months`. |
| `enabled` | no | `false` silences the entry without deleting it. Defaults to `true`. |

### Visualization Workflow

1. Attempt to read `Visualization.yaml`.
   - If the file does not exist, create it with the default entries and use those.
2. Validate every entry: a known `name`, a `dest` inside the vault, known `args` of the right type,
   and no two entries writing the same `dest`.
   - If any entry is invalid → [Show Error](#show-error) and [Stop Execution](#stop-execution).
     Nothing is rendered — a broken config leaves the vault as it was, never half-written.
3. For each enabled entry, in order, render it against the current state and write it to `dest`,
   overwriting whatever was there.
   - A `dest` carrying a placeholder (`{month}`, `{account}`) renders once per matching item.

Files removed from the list simply stop being refreshed; they are not deleted. Rendered files are
generated output — a hand edit is overwritten on the next tick.

The full list of visualizations and their `args` is in [`Visualization.md`](Visualization.md).

## Procedures

### Show Error

Creates an `Error.md` file containing information about the error.

### Stop Execution

1. Verify that `Task.yaml` exists.
   - If it does not exist, create it with default values.
2. Set `apply` to `false` in `Task.yaml`.
