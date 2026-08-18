# Choose Visualizations

## Description
Decides which pages your brain writes and where they go, by editing `Visualization.yaml`. Running the actions that change the underlying data is [RunTasks.md](/docs/Tutorials/RunTasks.md); adding a renderer that does not exist yet is [HandleVisualizations.md](/docs/Tutorials/HandleVisualizations.md).

### Rules
- What is not in `Visualization.yaml` is not rendered. There are no implicit pages.
- A folder `dest` is written **into**, never emptied. A page that stops being produced sits there frozen until you delete it.
- Two entries may not write to the same place, and one `dest` may not sit inside another — the second render would erase part of the first.
- A `dest` stays inside the vault. A path climbing out of it with `..` is rejected.

---

## Workflow

1. See what your brain can render:

```bash
wraith visualizations
```

2. Open `Visualization.yaml`. It is a **list**: every entry asks for one visualization, rendered to one destination.

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
|-------|----------|-------------|
| `name` | yes | A visualization from `wraith visualizations`. An unknown name is an error. |
| `dest` | yes | Where to write it, relative to the vault root. Missing folders are created. |
| `args` | no | Per-visualization options. Anything omitted falls back to its default. |
| `enabled` | no | `false` silences the entry without deleting it. Defaults to `true`. |

3. Move a tree by renaming its `dest`. The whole thing appears somewhere else on the next tick:

```yaml
- name: DashBoard
  dest: Home
```

The old folder is left where it is — delete it yourself.

4. Ask for the same visualization twice, with different args and different destinations:

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

5. Silence an entry without losing its settings:

```yaml
- name: Help
  dest: Help
  enabled: false
```

6. Render one entry on its own, without executing `Task.yaml` and without touching any other entry:

```bash
wraith render DashBoard --future-months 24
```

Flags override that entry's args for this invocation only; the file is never edited. `--dest` overrides where it writes, and is **required** for a visualization your config never declared.

7. Run a tick to apply the whole config:

```bash
wraith tick
```

A broken config renders **nothing** rather than half a vault: every entry is validated before the first file is written.
