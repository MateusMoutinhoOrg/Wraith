# Commands

## Description
Every command, flag, value format and exit code of the `wraith` interface. The interface itself is `api.Lib.Sandboxmain` — one field of the library — so everything below is produced inside the closed sandbox. Running your first action is [RunTasks.md](/docs/Tutorials/RunTasks.md); choosing what gets rendered is [ChooseVisualizations.md](/docs/Tutorials/ChooseVisualizations.md).

---

## Commands

Every command follows the same shape:

```bash
wraith <command> [arguments] [flags]
```

Arguments are positional and required unless the table says otherwise; flags are named, order-free, and fall back to a default.

| Command | Arguments | Writes | Purpose |
| --- | --- | --- | --- |
| `start` | — | `Task.yaml`, `Visualization.yaml`, every `dest` | Create a vault where there was none, then render it |
| `tick` | — | the data, every `dest` | Run the pending task, then render everything |
| `watch` | — | the data, every `dest` | Run a tick on an interval, until interrupted |
| `run` | `<task-name>` | the data, every `dest` | Run one task from the command line |
| `render` | `<visualization-name>` | one `dest` | Render one visualization to disk |
| `tasks` | — | nothing | List every task the binary carries |
| `visualizations` | — | nothing | List every visualization it can render |
| `help` | — | nothing | Print the usage screen |
| `version` | — | nothing | Print the interface version |

### `start`

```bash
wraith start --prev-months 6 --future-months 12 --current-month 2026-08
```

Writes the two files a brain is driven by, copied from the defaults compiled into the binary, and immediately runs a tick to render them. It **never overwrites**: a file already on disk is reported and left as it is, and a tick is not performed if the files already existed.

The three flags above are written into the `args:` of the `DashBoard` entry of the `Visualization.yaml` it creates:

| Flag | Default | Written as |
| --- | --- | --- |
| `--prev-months <count>` | `3` | `prev-months`, how many months back of `Months/` are drawn |
| `--future-months <count>` | `12` | `future-months`, how far ahead the forecast projects |
| `--current-month <YYYY-MM>` | the month today falls in | `current-month`, the month the dashboard renders as the open one |

They are a starting point, not a setting: what they write is a line in a file you own from then on. `--current-month` in particular **pins** the dashboard to that month — delete the line for the vault to follow the calendar again. A count that is not a whole number of zero or more, and a month not written as `YYYY-MM`, are usage errors, and nothing is created.

Because `start` never overwrites, the flags do nothing in a vault that already has a config: the file they would be written into is the one already there.

### `tick`

Performs a single tick of the state machine — see [The Tick Workflow](#the-tick-workflow).

### `watch`

```bash
wraith watch --time 1s
```

Runs a tick every `--time`, until interrupted. A failing tick prints its error and does not stop the loop; the task file is disarmed by the tick itself, so a broken action costs one tick rather than the session.

### `run`

```bash
wraith run AddTransaction --account Bank --category Food --amount -32.90 --date 2026-08-18
```

Runs one task without touching `Task.yaml`, then re-renders every visualization exactly as a tick does. One `--flag` per field the task declares; required fields must be given. An unknown name, a missing required field or a field of the wrong type is reported and **nothing** is written.

### `render`

```bash
wraith render DashBoard --future-months 24 --current-month 2026-03
```

Renders one visualization and writes it to its `dest`, without executing `Task.yaml` and without touching any other entry. Where it writes and what its args default to come from the matching entry in `Visualization.yaml`; a flag given here overrides that entry for this invocation only, and the file is never edited.

`--dest` is **required** when the name is declared nowhere in the config, and is what picks the entry when the same name is declared twice. `enabled: false` does not block an explicit invocation: asking for an entry by name is the decision to render it.

---

## Flags

| Flag | Default | Description |
| --- | --- | --- |
| `--task <path>` | `Task.yaml` | The task file a tick reads and resets |
| `--visualization <path>` | `Visualization.yaml` | The config a tick renders from |
| `--database <path>` | `data` | The folder inside the vault the data lives in |
| `--dest <path>` | the entry's `dest` | Where a single `render` writes |
| `--time <interval>` | — | How long `watch` sleeps between ticks. Required by `watch` |
| `--prev-months <count>` | `3` | Months of history the vault `start` creates renders |
| `--future-months <count>` | `12` | Months ahead the vault `start` creates forecasts |
| `--current-month <YYYY-MM>` | this month | The month the vault `start` creates renders as the open one |
| `--<field> <value>` | the declared default | One flag per field a task or a visualization declares |
| `-h`, `--help` | — | Print the usage screen and exit |
| `-v`, `--version` | — | Print the interface version and exit |
| `-q`, `--quiet` | — | Print only listings and errors |

The three path flags point the same binary at another vault without moving a file. Every `dest` stays relative to the vault root, not to `--database`.

---

## Values

| Kind | Written as | Examples |
| --- | --- | --- |
| Text | Bare, or quoted when it carries a space | `Bank`, `"Checking Account"` |
| Number | Decimal, negative for money going out | `3000`, `-32.90` |
| Switch | `true` or `false` | `--revenues true` |
| Date | `YYYY-MM-DD` | `2026-08-18` |
| Month | `YYYY-MM` | `2026-08` |
| Interval | A Go duration | `1s`, `500ms`, `2m` |

A value written into `Task.yaml` and the same value passed as a flag reach a task identically — the command line has only text, and the same validator turns it into a number or a switch either way.

---

## Exit Codes

| Code | Meaning |
| --- | --- |
| `0` | The command did what it was asked to |
| `1` | The command failed — an unknown task, an invalid field, a file that could not be written |
| `2` | The command line could not be understood, and was answered with the usage screen |

---

## The Tick Workflow

1. Read the task file. If it does not exist, stop — a vault that has not been started is not a broken vault, and this is not an error.
2. Check `apply`. If it is `false`, stop. Not an error either: that is an action waiting to be armed.
3. Check the task's name and its fields. An unknown task, an unknown field or a field of the wrong type writes `Error.md` and stops.
4. Execute the task. On failure, write `Error.md` and stop — **nothing** is changed.
5. Read the visualization config, validating every entry before writing a single file. A broken config renders nothing rather than half a vault.
6. Render every enabled entry, in order, writing each under its `dest`.
7. Set `apply` back to `false`, so the same action never runs twice.

A folder `dest` is written **into**, never emptied: a file a visualization no longer produces is left where it is rather than deleted.
