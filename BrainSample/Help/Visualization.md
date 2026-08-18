# Visualization Guide

A task changes the database. A **visualization** renders it. Which pages Wraith writes, where it
writes them, and how each one is shaped is not built in — it is declared by you in
[`Visualization.yaml`](../Visualization.yaml).

A visualization is not a page: it is a **renderer that owns a destination**. Give it a name, its
args, and where to put it — Wraith writes the whole thing, be that one file or an entire folder
tree. If a visualization is not listed in `Visualization.yaml`, it is never rendered. That is the
whole mechanic: you choose what your vault is made of.

---

## 1. How `Visualization.yaml` works

The file is a **list** of entries. Every entry asks for one visualization, rendered to one
destination:

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
| ----- | -------- | ----------- |
| `name` | yes | A visualization from the [catalog](#3-visualization-catalog). Unknown names are an error. |
| `dest` | yes | Where to write it, relative to the vault root. Missing folders are created. |
| `args` | no | Per-visualization options. Anything omitted falls back to the default in the catalog. |
| `enabled` | no | `false` silences the entry without deleting it. Defaults to `true`. |

Entries are rendered top to bottom on every tick, each one overwriting its `dest` with the current
state of the database.

---

## 2. What `dest` means

Every visualization declares whether it renders **a single file** or **a whole folder**. The
catalog says which, and `dest` follows:

| Kind | `dest` is | Wraith writes |
| ---- | --------- | ------------- |
| File | The path of the file, extension included — `Usage.md` | That one file |
| Folder | The path of the folder, no extension — `DashBoard` | Every page the visualization produces, in the tree it defines below `dest` |

A folder visualization owns its internal layout: the file names below `dest` and how deep they nest
are part of the visualization, not of your config. `DashBoard` writes `README.md`, `Forecast.md` and
a `Months/<month>/` subtree under whatever folder you point it at — rename `dest` to `Home` and the
same tree appears under `Home/` instead.

Because the destination is a whole tree, one entry is enough for a page set that used to need one
line per file. There are no `{month}` or `{account}` placeholders: the visualization expands its own
tree from the database, and its args are what control how far that expansion goes.

Renaming a `dest` does not move the old tree. Wraith writes the new one and leaves the previous
folder where it is — delete it yourself.

---

## 3. Visualization catalog

| Visualization | Kind | Args | Renders |
| ------------- | ---- | ---- | ------- |
| `DashBoard` | folder | `prev-months`, `future-months` | The full financial vault: position, registries, one folder per month, and the forecast |
| `Task-List` | folder | — | One reference page per task the binary carries |
| `Help` | folder | — | The guides: the command reference, how tasks work and how visualizations work |
| `Usage` | file | — | The command reference and the tick workflow, on its own |

### `DashBoard`

The whole picture, as a folder tree:

```
DashBoard/
├── README.md                 current position, the open month, the index of every other page
├── Accounts.md               every account, its balance and its share of the money you hold
├── Credit-Cards.md           every card: limit, outstanding, closing day, due day
├── Categories.md             every category, what it classifies and what it has cost
├── Forecast.md               one row per future month: what each account holds and the net position
└── Months/
    ├── README.md             the index of every rendered month
    └── <year>-<month>/
        ├── DashBoard.md      that month's result, accounts, categories and dated commitments
        ├── Statement.md      every transaction dated in that month
        └── Accounts/
            └── <account>.md  one account's movements inside that month
```

| Arg | Default | Description |
| --- | ------- | ----------- |
| `prev-months` | `3` | How many months back of `Months/` to render, counting from the open month. Months holding no transaction are skipped. |
| `future-months` | `8` | The forecast horizon — how many months after the open one `Forecast.md` projects. |

`prev-months` is what bounds the tree: raise it and older months reappear as folders, lower it and
the tree shrinks — already-written folders stay on disk, frozen at their last render, until you
delete them. `future-months` reads recurrences, card bills and future-dated transactions; see
[`../Tasks/AddRecurrence.md`](../Tasks/AddRecurrence.md).

### `Task-List`

One page per task, named after it — `AddTransaction.md`, `RemoveCategory.md`, … Each page is the
task's fields, its rules and a copy-ready sample. The list comes from the binary itself, so it
cannot drift from the tasks Wraith actually accepts.

### `Help`

Every guide, as a folder tree:

```
Help/
├── Usage.md            every command and the tick workflow
├── Task.md             how to run a task, and the catalog of every one of them
└── Visualization.md    this page
```

Like `Task-List`, all three are generated from the registries the binary carries, so no guide can
describe a command, a task or a visualization that does not exist.

### `Usage`

[`Usage.md`](Usage.md) on its own, for a vault that wants the command reference somewhere else and
none of the other guides:

```yaml
- name: Usage
  dest: README.md
```

It is the same page `Help` writes, which is exactly why the two cannot both aim at it: an entry
whose `dest` falls inside another entry's folder is rejected. Declare `Usage` separately only when
`Help` is not writing to the same place — or is not declared at all.

---

## 4. Previewing without writing

To render one entry to the terminal without writing anything to disk — useful to check a change to
`args` before committing to it:

```bash
./wraith render DashBoard
./wraith render DashBoard --future-months 24
```

For a folder visualization, `render` prints the tree it would write and each file below it. It never
touches `dest`. See [`Usage.md`](Usage.md) for the full tick workflow.

---

## 5. Shaping the vault

Some worked examples of the mechanic.

**Only the numbers, none of the docs** — a vault driven by scripts has no use for rendered guides:

```yaml
- name: DashBoard
  dest: DashBoard
```

**A flat vault** — `dest` is a free path, so where each tree lands is yours:

```yaml
- name: DashBoard
  dest: .

- name: Help
  dest: Docs/Guides
```

**The same visualization twice** — different args, different destinations:

```yaml
- name: DashBoard
  args:
    future-months: 8
  dest: DashBoard

- name: DashBoard
  args:
    prev-months: 36
    future-months: 24
  dest: Archive
```

**Paused, not deleted** — keep the entry and its args around while it is quiet:

```yaml
- name: Help
  dest: Help
  enabled: false
```

---

## 6. Rules to remember

- What is not in `Visualization.yaml` is not rendered. There are no implicit pages.
- Rendered files are **generated**: a hand edit is overwritten on the next tick. Everything you
  want to keep goes into a task, never into a rendered page.
- A folder `dest` is **written into, never emptied**. Wraith overwrites the files the visualization
  produces and leaves anything else alone — a page that stops being produced sits there frozen at
  its last render until you delete it.
- Removing an entry stops its tree from being refreshed — it does **not** delete what is on disk.
- Two entries writing the same `dest` is an error, and so is a `dest` nested inside another entry's
  folder: the second render would erase part of the first.
- `dest` stays inside the vault. A path climbing out of it with `../` is rejected.
- An unknown `name`, an unknown arg, or an arg of the wrong type creates `Error.md` and leaves every
  file untouched — a broken config renders nothing rather than half a vault.
- If `Visualization.yaml` does not exist, a tick creates it with the default entries: `DashBoard`,
  `Task-List` and `Help`.
- Order is only cosmetic. Every entry reads the same state, so no entry can see another's output.
