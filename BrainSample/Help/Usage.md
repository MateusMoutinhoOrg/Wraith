# Wraith

Wraith is a data-visualization system that applies actions to a database and renders the
current state of that database through several visualization dashboards.

## Table of Contents

- [Overview](#overview)
- [Commands](#commands)
  - [watch](#watch)
  - [tick](#tick)
  - [run](#run)
  - [visualize](#visualize)
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

Every command follows the same shape:

```bash
./wraith <command> [arguments] [flags]
```

Arguments are positional and required unless the table below says otherwise; flags are named,
order-free, and fall back to a default. A flag written `--<arg>` stands for a whole family: one flag
per arg the named task or visualization declares.

| Command | Arguments | Writes | Purpose |
| --- | --- | --- | --- |
| [`watch`](#watch) | — | database, every `dest` | Runs a tick on an interval |
| [`tick`](#tick) | — | database, every `dest` | Runs one tick |
| [`run`](#run) | `<task-name>` | database, every `dest` | Runs one task from the command line |
| [`visualize`](#visualize) | `<visualization-name>` | one `dest` | Renders one visualization to disk |
| [`render`](#render) | `<visualization-name>` | nothing | Previews one visualization in the terminal |

### watch

```bash
./wraith watch --time <interval> [flags]
```

Runs a [tick](#tick) every `<interval>`, until interrupted.

**Arguments**

| Argument | Required | Description | Example |
| --- | --- | --- | --- |
| — | — | Takes no positional argument | — |

**Flags**

| Flag | Default | Description | Example |
| --- | --- | --- | --- |
| `--time` | — | Interval between ticks. Required. | `--time 1s`, `--time 500ms`, `--time 2m` |
| `--task` | `Task.yaml` | Task file every tick reads | `--task Inbox/Task.yaml` |
| `--visualization` | `Visualization.yaml` | Visualization config every tick renders from | `--visualization Archive.yaml` |
| `--database` | `data` | Folder the database lives in | `--database vaults/home` |

**Example**

```bash
./wraith watch --time 1s --task Inbox/Task.yaml
```

The three path flags mean the same thing as in [`tick`](#tick), and are read once at startup:
every tick of the loop uses them.

### tick

```bash
./wraith tick [flags]
```

Performs a single tick of the state machine: executes the pending action declared in `Task.yaml`
and renders every visualization declared in `Visualization.yaml`. See
[Tick Workflow](#tick-workflow).

**Arguments**

| Argument | Required | Description | Example |
| --- | --- | --- | --- |
| — | — | Takes no positional argument | — |

**Flags**

| Flag | Default | Description | Example |
| --- | --- | --- | --- |
| `--task` | `Task.yaml` | Which task file to read and reset | `--task Inbox/Task.yaml` |
| `--visualization` | `Visualization.yaml` | Which visualization config to render from | `--visualization Archive.yaml` |
| `--database` | `data` | Which folder the database lives in | `--database vaults/home` |

**Example**

```bash
./wraith tick --database vaults/home
```

The three point the tick at a different vault without moving a file: `--task` replaces
`Task.yaml` everywhere the [Tick Workflow](#tick-workflow) and the [Procedures](#procedures)
name it — it is the file read, reset to `apply: false`, and created with defaults when missing.
`--visualization` replaces `Visualization.yaml` the same way, including the create-with-defaults
step. `--database` is the folder read and written by the action, created if it does not exist.
Every `dest` in the visualization config stays relative to the vault root, not to `--database`.

### run

```bash
./wraith run <task-name> [flags]
```

Runs a task directly from the command line, without editing `Task.yaml`. This is the preferred way
to drive Wraith programmatically (scripts, automations, other tools). On success the database is
written and every visualization is re-rendered, exactly as in a [tick](#tick).

**Arguments**

| Argument | Required | Description | Example |
| --- | --- | --- | --- |
| `<task-name>` | yes | Which task to run — any name in [`Task.md`](Task.md) | `AddTransaction` |

**Flags**

| Flag | Default | Description | Example |
| --- | --- | --- | --- |
| `--<field>` | — | One flag per field the task declares. Required fields must be given. | `--amount 100` |
| `--visualization` | `Visualization.yaml` | Which visualization config to render from | `--visualization Archive.yaml` |
| `--database` | `data` | Which folder the database lives in | `--database vaults/home` |

**Example**

```bash
./wraith run AddTransaction --amount 100 --date 2026-08-18 --description test
```

Rules:

- The name must be a known task. An unknown name, an unknown field, a missing required field or a
  field of the wrong type is reported and nothing is written.
- `Task.yaml` is neither read nor reset — the task comes entirely from the command line.

### visualize

```bash
./wraith visualize <visualization-name> [flags]
```

Renders a single visualization and **writes it to its `dest`**, without executing `Task.yaml` and
without touching any other entry. It is to `Visualization.yaml` what [`run`](#run) is to
`Task.yaml`: the whole tick narrowed down to one thing, driven from the command line.

The destination and the defaults come from the matching entry in `Visualization.yaml`; any flag
given on the command line overrides that entry's value for this invocation only — the file is never
edited.

**Arguments**

| Argument | Required | Description | Example |
| --- | --- | --- | --- |
| `<visualization-name>` | yes | Which visualization to render — any name in [`Visualization.md`](Visualization.md) | `DashBoard` |

**Flags**

| Flag | Default | Description | Example |
| --- | --- | --- | --- |
| `--dest` | the entry's `dest` | Where to write. Required when the name is not declared. | `--dest Archive` |
| `--<arg>` | the entry's value | One flag per arg the visualization declares | `--future-months 12` |
| `--visualization` | `Visualization.yaml` | Which visualization config to look the name up in | `--visualization Archive.yaml` |
| `--database` | `data` | Which folder the database is read from | `--database vaults/home` |

**Example**

```bash
./wraith visualize DashBoard --future-months 12
```

Rules:

- The name must be a known visualization. An unknown name, an unknown arg or an arg of the wrong
  type is reported and nothing is written.
- If the name is not declared in `Visualization.yaml`, `--dest` is required. If it is declared more
  than once with different destinations, `--dest` is what picks which one to write.
- `enabled: false` does not block an explicit invocation — asking for the entry by name is the
  decision to render it.
- `dest` obeys the same rules as in a tick: it stays inside the vault, folders are created as
  needed, and a folder `dest` is written into, never emptied.
- The database is only read. `Task.yaml` is neither executed nor reset.

### render

```bash
./wraith render <visualization-name> [flags]
```

Renders a single visualization to standard output, without writing anything to `dest` and without
touching the database. Use it to preview an entry — or a different set of args — before putting it
into `Visualization.yaml`. For a visualization that renders a folder, it prints the tree it would
write and every file below it.

It takes the same argument and the same arg flags as [`visualize`](#visualize) — the two differ only
in where the output lands, the terminal or `dest`. Because nothing is written, `render` needs no
`dest` at all: a name that appears nowhere in `Visualization.yaml` still previews fine.

**Arguments**

| Argument | Required | Description | Example |
| --- | --- | --- | --- |
| `<visualization-name>` | yes | Which visualization to preview — any name in [`Visualization.md`](Visualization.md) | `DashBoard` |

**Flags**

| Flag | Default | Description | Example |
| --- | --- | --- | --- |
| `--<arg>` | the entry's value | One flag per arg the visualization declares | `--future-months 12` |
| `--visualization` | `Visualization.yaml` | Which visualization config to read defaults from | `--visualization Archive.yaml` |
| `--database` | `data` | Which folder the database is read from | `--database vaults/home` |

**Example**

```bash
./wraith render DashBoard --future-months 12
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

Wraith renders no page on its own. Everything it writes is one entry you declared in
`Visualization.yaml`, so the shape of your vault is yours to choose:

```yaml
- name: DashBoard
  args:
    prev-months: 3
    future-months: 8
  dest: DashBoard

- name: Task-List
  dest: Tasks

- name: Help
  dest: Help
```

| Field | Required | Description |
| --- | --- | --- |
| `name` | yes | Which visualization to render — `DashBoard`, `Task-List`, `Help`, `Usage` |
| `dest` | yes | Where to write it, relative to the vault root. Missing folders are created. |
| `args` | no | Per-visualization options, such as `DashBoard.future-months`. |
| `enabled` | no | `false` silences the entry without deleting it. Defaults to `true`. |

A visualization renders either **one file** or **a whole folder**, and `dest` follows: a file path
for the first kind, a folder path for the second. A folder visualization owns the tree it writes —
`DashBoard` produces `README.md`, `Forecast.md` and a `Months/<month>/` subtree under whatever
folder you point it at. Its `args` are what bound that expansion, so a page set never needs one
entry per file.

Three entries are what produce this whole vault: this page is not one of them, it is part of the
`Help` tree. `Usage` also exists as a standalone file visualization, for a vault that wants the
command reference somewhere else and none of the other guides — but it cannot be declared *and*
have `Help` write over the same place, since destinations may not nest.

### Visualization Workflow

1. Attempt to read `Visualization.yaml`.
   - If the file does not exist, create it with the default entries and use those.
2. Validate every entry: a known `name`, a `dest` inside the vault, known `args` of the right type,
   and no two entries whose destinations collide or nest inside one another.
   - If any entry is invalid → [Show Error](#show-error) and [Stop Execution](#stop-execution).
     Nothing is rendered — a broken config leaves the vault as it was, never half-written.
3. For each enabled entry, in order, render it against the current state and write it to `dest`,
   overwriting whatever was there.
   - A folder `dest` is written into, never emptied: files the visualization no longer produces are
     left untouched rather than deleted.

Entries removed from the list simply stop being refreshed; nothing is deleted. Rendered files are
generated output — a hand edit is overwritten on the next tick.

The full catalog of visualizations, the tree each one writes and their `args` is in
[`Visualization.md`](Visualization.md).

## Procedures

### Show Error

Creates an `Error.md` file containing information about the error.

### Stop Execution

1. Verify that `Task.yaml` exists.
   - If it does not exist, create it with default values.
2. Set `apply` to `false` in `Task.yaml`.
