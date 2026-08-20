# Start a Brain

## Description
Turns an empty folder into a working vault: the two files a brain is driven by, and the pages it renders. Installing the binary first is [InstallCli.md](/docs/Tutorials/InstallCli.md); running your first action afterwards is [RunTasks.md](/docs/Tutorials/RunTasks.md).

### Rules
- A vault is a **folder**. The brain you drive is the folder you run `wraith` in, so a second brain is a second directory and nothing more.
- `wraith start` never overwrites. Running it twice in the same folder reports what is already there and changes nothing.

---

## Workflow

1. Make the folder your brain will live in, and go into it.

```bash
mkdir my-brain
cd my-brain
```

2. Create the two files the brain is driven by, and render the first pages.

```bash
wraith start
```

You now have `Task.yaml` — the one action waiting to run — and `Visualization.yaml` — the list of pages you want written. Because `start` also performs an immediate tick, those pages are already on disk.

3. Look at what appeared:

```bash
ls
# DashBoard/  Help/  Task.yaml  Tasks/  Visualization.yaml
```

| Folder | Written by | What it holds |
|--------|-----------|---------------|
| `DashBoard/` | the `DashBoard` visualization | Your position, one page per account, one folder per month, and the forecast on the month index |
| `Tasks/` | the `Task-List` visualization | One reference page per action the binary carries |
| `Help/` | the `Help` visualization | The command reference, the task guide, the visualization guide |

Every one of those folders exists because a line in `Visualization.yaml` asked for it. Delete the line and the folder stops being refreshed; see [ChooseVisualizations.md](/docs/Tutorials/ChooseVisualizations.md).

4. Record something, so the pages have something to show.

```bash
wraith run AddAccount --account Bank
```

5. Open `DashBoard/README.md`. Your account is on it, and the row that names it opens `DashBoard/Accounts/Bank.md` — that account's own page.

6. Leave a watcher running while you work, so saving `Task.yaml` in an editor is all it takes to apply an action.

```bash
wraith watch --time 1s
```

Every rendered page is **generated**. A hand edit is overwritten on the next tick — everything you want to keep goes into a task.
