# Task Guide

Every action in Wraith is performed by writing a task into [`Task.yaml`](../Task.yaml) and letting
the state machine pick it up. This page lists every available task and how to run one.

---

## 1. How to run a task

1. Pick a task from the tables below and open its **Guide** — every guide ends with a
   **Sample** section.
2. Copy the sample into [`Task.yaml`](../Task.yaml) at the project root and fill in your values.
3. Make sure `apply: true` is set — with `apply: false` the task is ignored (no error).
4. Trigger a tick:

```bash
./wraith tick              # run once
./wraith watch --time 1s   # or keep watching; every tick picks up Task.yaml
```

On success every visualization declared in [`Visualization.yaml`](../Visualization.yaml) is
re-rendered — in this vault that includes the tree under [`DashBoard/`](../DashBoard/README.md) —
and `apply` is reset to `false`. On failure an `Error.md` file is created with the details and
nothing is changed. Both commands take `--task`, `--visualization` and `--database` to point at a
file or a vault other than the defaults (`Task.yaml`, `Visualization.yaml` and `data`). See
[`Usage.md`](Usage.md) for the full tick workflow and
[`Visualization.md`](Visualization.md) for how to choose what gets rendered.

### Anatomy of `Task.yaml`

```yaml
name: AddTransaction   # which task to run — one of the names below
# ...task-specific fields (see each task's Guide)...
apply: true            # true = execute on the next tick
```

---

## 2. Available tasks

### Transactions — day-to-day

| Task           | Description                               | Links |
| -------------- | ----------------------------------------- | ----- |
| AddTransaction    | Record an income, an expense or one leg of a transfer | [Guide](../Tasks/AddTransaction.md) |
| ModifyTransaction | Modify any attribute of an existing transaction       | [Guide](../Tasks/ModifyTransaction.md) |

A purchase paid over several months is still one `AddTransaction` — give it `installments: N` and
it writes N monthly transactions at once, each landing in its own month.

### Recurrences — what is coming

| Task             | Description                                              | Links |
| ---------------- | -------------------------------------------------------- | ----- |
| AddRecurrence    | Declare a commitment that repeats every month             | [Guide](../Tasks/AddRecurrence.md) |
| RemoveRecurrence | Stop a recurring commitment                               | [Guide](../Tasks/RemoveRecurrence.md) |

A recurrence is the only task that describes the **future**. It moves no balance and writes no
transaction — it is a rule that [`DashBoard/Forecast.md`](../DashBoard/Forecast.md) reads to project
what each account will hold over the forecast horizon — `DashBoard.future-months`, 8 months here.
When the day arrives you still record what actually happened with `AddTransaction`.

### Categories

| Task           | Description                                   | Links |
| -------------- | --------------------------------------------- | ----- |
| AddCategory    | Add a category to classify transactions       | [Guide](../Tasks/AddCategory.md) |
| RemoveCategory | Remove a category                             | [Guide](../Tasks/RemoveCategory.md) |

### Accounts

| Task          | Description                | Links |
| ------------- | -------------------------- | ----- |
| AddAccount    | Add an account             | [Guide](../Tasks/AddAccount.md) |
| RemoveAccount | Remove an account          | [Guide](../Tasks/RemoveAccount.md) |

### Credit cards

| Task             | Description           | Links |
| ---------------- | --------------------- | ----- |
| AddCreditCard    | Add a credit card     | [Guide](../Tasks/AddCreditCard.md) |
| RemoveCreditCard | Remove a credit card  | [Guide](../Tasks/RemoveCreditCard.md) |

---

## 3. Moving money between your own accounts

There is no transfer task. A one-off transfer is two `AddTransaction`s sharing a **transfer category** — a
category created with `revenues: false` **and** `expenses: false`, which is how Wraith knows the
movement is neither income nor expense:

```yaml
name: AddTransaction        # money leaves the bank
account: Bank
category: Card Payment
description: August card bill
amount: -1286
date: 2026-09-05
apply: true
```

```yaml
name: AddTransaction        # ...and lands on the card
account: Nubank Card
category: Card Payment
description: August card bill
amount: 1286
date: 2026-09-05
apply: true
```

The pair nets to zero, so paying a card bill never shows up as an expense — the purchases were
already counted on the day they happened.

A transfer that repeats every month is the one exception to the two-task rule: a single
[`AddRecurrence`](../Tasks/AddRecurrence.md) with `to_account` filled in declares both legs at once.

---

## 4. What each dashboard page is fed by

The pages below are the tree the `DashBoard` visualization writes. Paths are relative to its `dest`
— `DashBoard/` in this vault, whatever you set in [`Visualization.yaml`](../Visualization.yaml)
elsewhere. How far back `Months/` goes and how far ahead `Forecast.md` looks are its `prev-months`
and `future-months` args; see [`Visualization.md`](Visualization.md).

| Dashboard file | Rendered from |
| -------------- | ------------- |
| [`DashBoard/README.md`](../DashBoard/README.md) | Everything — balances, the open month, and the index of every other page |
| [`DashBoard/Accounts.md`](../DashBoard/Accounts.md) | AddAccount, RemoveAccount, AddTransaction |
| [`DashBoard/Credit-Cards.md`](../DashBoard/Credit-Cards.md) | AddCreditCard, RemoveCreditCard, AddTransaction |
| [`DashBoard/Categories.md`](../DashBoard/Categories.md) | AddCategory, RemoveCategory, AddTransaction |
| [`DashBoard/Forecast.md`](../DashBoard/Forecast.md) | AddRecurrence, RemoveRecurrence, AddCreditCard, and any transaction dated ahead of today |
| [`DashBoard/Months/README.md`](../DashBoard/Months/README.md) | Every month that holds at least one transaction |
| `DashBoard/Months/<year>-<month>/DashBoard.md` | That month's result, accounts, categories and dated commitments |
| `DashBoard/Months/<year>-<month>/Statement.md` | Every transaction dated in that month |
| `DashBoard/Months/<year>-<month>/Accounts/<account>.md` | One account's statement for that month |

Every figure on those pages is computed from the five registries the tasks above write — accounts,
credit cards, categories, transactions and recurrences. If a number cannot be derived from them, it
does not belong on the dashboard. `Forecast.md` obeys the same rule: it projects declared
commitments, and the one place it uses a historical average is fenced off in its own section and
labelled as an estimate.

---

## 5. Rules to remember

- One task per tick — `Task.yaml` holds a single action.
- Dashboard files are **generated**: never edit them by hand, a hand edit is overwritten on the
  next tick.
- One category and one account per transaction — no splits.
- A positive `amount` is only valid in a category whose `revenues` is `true`; a negative one only
  where `expenses` is `true`. A transfer category accepts both because it counts as neither.
- A card purchase counts on its `date`; the money leaves the paying account when you record the
  bill payment.
- `payment_date` is when the money actually moves. Until that date arrives the transaction counts
  in the month's result but not in the account balance.
- Months are never created by hand — a month appears on the dashboard as soon as a transaction
  carries a `date` inside it. Installment parts count: a purchase in 12x opens the next 11 months.
  Whether it is *rendered* as a folder is a separate question, answered by `DashBoard.prev-months`.
- A recurrence never becomes a transaction on its own. Nothing is recorded until you record it —
  the forecast only says what you told it to expect.
- A card bill is never declared as a recurrence. Its amount changes every month, so the forecast
  derives it from the card's `closing_day`, `due_day` and whatever falls inside the window.
- The forecast horizon is `DashBoard.future-months` months after the open one, and moves with the
  calendar. Change the arg in `Visualization.yaml` and the next tick re-renders against it.
