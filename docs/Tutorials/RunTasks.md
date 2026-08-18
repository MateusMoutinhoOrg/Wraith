# Run Tasks

## Description
Runs one action against your brain, the two ways there are: writing it into `Task.yaml`, or passing it straight on the command line. Creating the vault first is [StartABrain.md](/docs/Tutorials/StartABrain.md); a worked example of a whole month is [TrackTransactions.md](/docs/Tutorials/TrackTransactions.md).

### Rules
- One task per tick. `Task.yaml` holds a single action.
- `apply: false` is not an error. It is an action waiting to be armed, which is the normal state between two edits.
- A failed task changes **nothing** and writes `Error.md`. There is no half-applied task.

---

## Workflow

1. See what your brain can do. The list comes from the binary itself, so it can never name a task that does not exist.

```bash
wraith tasks
```

2. Open the task's page under `Tasks/` — `Tasks/AddTransaction.md` — and copy its **Sample**.

3. Paste it over `Task.yaml`, fill in your values, and set `apply: true`.

```yaml
name: AddTransaction
account: Bank
category: Food
description: Market
amount: -32.90
date: 2026-08-18
apply: true
```

4. Run a tick.

```bash
wraith tick
```

The action is applied, every visualization is re-rendered, and `apply` goes back to `false` — so the same action never runs twice.

5. Or skip the file entirely. Every field a task declares is also a flag, which is how a script drives a brain:

```bash
wraith run AddTransaction \
  --account Bank --category Food --amount -32.90 \
  --date 2026-08-18 --description Market
```

`wraith run` neither reads nor resets `Task.yaml`. On success it re-renders exactly as a tick does.

6. When something is wrong, read `Error.md` in the vault root. It names what failed and why:

```
AddTransaction: account not found: Ghost
```

Fix the cause and run another tick. Nothing was changed in the meantime.

7. To stop repeating yourself, leave a watcher running and just save the file:

```bash
wraith watch --time 1s
```

A failing tick does not stop the loop — the failure is printed, `Error.md` is written, and the next interval comes around.
