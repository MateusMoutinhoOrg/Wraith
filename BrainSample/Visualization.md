# Visualization Guide

A task changes the database. A **visualization** renders it. Which pages Wraith writes, where it
writes them, and how each one is shaped is not built in — it is declared by you in
[`Visualization.yaml`](Visualization.yaml).

If a visualization is not listed there, it is never rendered. That is the whole mechanic: you
choose the dashboards you want to see.

---

## 1. How `Visualization.yaml` works

The file is a **list** of entries. Every entry asks for one visualizer, rendered to one destination:

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
| ----- | -------- | ----------- |
| `name` | yes | A visualizer from the [catalog](#3-visualizer-catalog). Unknown names are an error. |
| `dest` | yes | Where to write it, relative to the vault root. Missing folders are created. |
| `args` | no | Per-visualizer options. Anything omitted falls back to the default in the catalog. |
| `enabled` | no | `false` silences the entry without deleting it. Defaults to `true`. |

Entries are rendered top to bottom on every tick, each one overwriting its `dest` with the current
state of the database.

---

## 2. Running visualizations

Every tick re-renders the whole list after the task has been applied:

```bash
./wraith tick              # apply Task.yaml, then render Visualization.yaml
./wraith watch --time 1s   # same thing, on a loop
```

To render one entry to the terminal without writing any file — useful to preview a change to
`args` before committing to it:

```bash
./wraith render ForeCast
./wraith render ForeCast --months 12
```

`render` never touches `dest`. See [`Usage.md`](Usage.md) for the full tick workflow.

---

## 3. Visualizer catalog

### The whole picture

| Visualizer | Args | Renders |
| ---------- | ---- | ------- |
| `DashBoard` | — | Current position, the open month, the next months, and the index of every other page |
| `ForeCast` | `months` | One row per future month: what each account holds and the net position |

`ForeCast.months` sets the horizon — how many months after the open one to project. Default `8`.
Recurrences, card bills and future-dated transactions are what it reads; see
[`Tasks/AddRecurrence.md`](Tasks/AddRecurrence.md).

### The registries

| Visualizer | Args | Renders |
| ---------- | ---- | ------- |
| `Accounts` | — | Every account, its balance and its share of the money you hold |
| `CreditCards` | — | Every card: limit, outstanding, closing day, due day |
| `Categories` | — | Every category, what it classifies and what it has cost |

### Months

| Visualizer | Args | Renders |
| ---------- | ---- | ------- |
| `Months` | `last` | The index of every month holding at least one transaction |
| `Month` | `month` | One month's result, accounts, categories and dated commitments |
| `Statement` | `month` | Every transaction dated in one month |
| `AccountStatement` | `month`, `account` | One account's movements inside one month |

`month` accepts a fixed `2026-08`, or `open` for whichever month is currently open — with `open`
the entry follows the calendar instead of pinning a date. `last` caps how many months the index
lists, newest first; omit it to list them all.

Month visualizers are the reason `dest` accepts placeholders (see [§4](#4-fanning-one-entry-over-many-files)).

### The documentation

| Visualizer | Args | Renders |
| ---------- | ---- | ------- |
| `Usage` | — | The command reference and tick workflow — [`Usage.md`](Usage.md) |
| `TaskHelp` | — | The catalog of every task — [`Task-help.md`](Task-help.md) |
| `VisualizationHelp` | — | This page |

These read the same registries as the rest: the task tables in `Task-help.md` list the tasks the
binary actually carries, so they cannot drift from it.

---

## 4. Fanning one entry over many files

A `dest` may carry placeholders. When it does, the entry renders once per matching item instead of
once in total:

| Placeholder | Expands to |
| ----------- | ---------- |
| `{month}` | Every month the `month` arg selects — `2026-08`, `2026-07`, … |
| `{account}` | Every account in the registry, slugged — `Emergency-Savings` |

```yaml
- name: Month
  args:
    month: all
  dest: DashBoard/Months/{month}/DashBoard.md

- name: AccountStatement
  args:
    month: all
  dest: DashBoard/Months/{month}/Accounts/{account}.md
```

Those two entries are what produce the per-month tree under
[`DashBoard/Months/`](DashBoard/Months/README.md). Drop them and the tree stops being written; give
`month` a fixed `2026-08` and only that folder is.

A `dest` without a placeholder that resolves to more than one item is an error — Wraith will not
silently overwrite the same file eight times.

---

## 5. Shaping the vault

Some worked examples of the mechanic.

**Only the numbers, none of the docs** — a vault driven by scripts has no use for rendered guides:

```yaml
- name: DashBoard
  dest: DashBoard/README.md
```

**A second forecast, further out** — the same visualizer twice with different args and destinations:

```yaml
- name: ForeCast
  args:
    months: 8
  dest: DashBoard/Forecast.md

- name: ForeCast
  args:
    months: 24
  dest: DashBoard/Forecast-Long.md
```

**A flat vault** — `dest` is a free path, so the folder layout is yours:

```yaml
- name: DashBoard
  dest: Home.md

- name: Accounts
  dest: Accounts.md

- name: Month
  args:
    month: open
  dest: This-Month.md
```

**Paused, not deleted** — keep the entry and its args around while it is quiet:

```yaml
- name: ForeCast
  args:
    months: 8
  dest: DashBoard/Forecast.md
  enabled: false
```

---

## 6. Rules to remember

- What is not in `Visualization.yaml` is not rendered. There are no implicit pages.
- Rendered files are **generated**: a hand edit is overwritten on the next tick. Everything you
  want to keep goes into a task, never into a dashboard.
- Removing an entry stops the file from being refreshed — it does **not** delete what is already on
  disk. Delete the stale file yourself, or it will sit there frozen at its last render.
- Two entries writing the same `dest` is an error: the second would erase the first.
- `dest` stays inside the vault. A path climbing out of it with `../` is rejected.
- An unknown `name`, an unknown arg, or an arg of the wrong type creates `Error.md` and leaves every
  file untouched — a broken config renders nothing rather than half a vault.
- If `Visualization.yaml` does not exist, a tick creates it with the default entries: `Usage`,
  `DashBoard` and `ForeCast`.
- Order is only cosmetic. Every entry reads the same state, so no entry can see another's output.
