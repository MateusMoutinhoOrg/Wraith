# Task Guide

Every action in Brain is performed by writing a task into [`Task.yaml`](Task.yaml) and letting
the state machine pick it up. This page lists every available task and how to run one.

---

## 1. How to run a task

1. Pick a task from the tables below and open its **Guide** — every guide ends with a
   **Sample** section.
2. Copy the sample into [`Task.yaml`](Task.yaml) at the project root and fill in your values.
3. Make sure `apply: true` is set — with `apply: false` the task is ignored (no error).
4. Trigger a tick:

```bash
./brain tick              # run once
./brain watch --time 1s   # or keep watching; every tick picks up Task.yaml
```

On success the dashboards under [`DashBoard/`](DashBoard/README.md) are re-rendered and
`apply` is reset to `false`. On failure an `Error.md` file is created with the details and
nothing is changed. See [`Usage.md`](Usage.md) for the full tick workflow.

### Anatomy of `Task.yaml`

```yaml
name: AddTransaction   # which task to run — one of the names below
# ...task-specific fields (see each task's Guide)...
apply: true            # true = execute on the next tick
```

---

## 2. Available tasks

### Transactions & transfers — day-to-day

| Task              | Description                                            | Links |
| ----------------- | ------------------------------------------------------ | ----- |
| AddTransaction    | Record an income or expense in the ledger              | [Guide](Tasks/AddTransaction.md) |
| RemoveTransaction | Remove a wrong or duplicated transaction               | [Guide](Tasks/RemoveTransaction.md) |
| AddTransfer       | Move money between own accounts (card bill, reserve)   | [Guide](Tasks/AddTransfer.md) |
| RemoveTransfer    | Remove a wrong transfer                                | [Guide](Tasks/RemoveTransfer.md) |

### Categories & budget

| Task            | Description                                              | Links |
| --------------- | -------------------------------------------------------- | ----- |
| AddCategory     | Add a category to classify transactions                  | [Guide](Tasks/AddCategory.md) |
| RemoveCategory  | Remove a category, migrating its transactions            | [Guide](Tasks/RemoveCategory.md) |
| SetBudget       | Set or update a category's spending limit                | [Guide](Tasks/SetBudget.md) |
| AddReallocation | Move budget between categories inside the current month  | [Guide](Tasks/AddReallocation.md) |

### Accounts

| Task          | Description                                          | Links |
| ------------- | ---------------------------------------------------- | ----- |
| AddAccount    | Add an account (bank, cash, card, savings)           | [Guide](Tasks/AddAccount.md) |
| RemoveAccount | Remove an empty account and its statement            | [Guide](Tasks/RemoveAccount.md) |

### Recurring bills

| Task                | Description                                       | Links |
| ------------------- | ------------------------------------------------- | ----- |
| AddRecurringBill    | Register a monthly bill (rent, subscriptions…)    | [Guide](Tasks/AddRecurringBill.md) |
| RemoveRecurringBill | Remove a cancelled bill                           | [Guide](Tasks/RemoveRecurringBill.md) |

### Net worth & goals

| Task            | Description                                          | Links |
| --------------- | ---------------------------------------------------- | ----- |
| AddAsset        | Add a non-account asset (car, equipment…)            | [Guide](Tasks/AddAsset.md) |
| RemoveAsset     | Remove a sold or written-off asset                   | [Guide](Tasks/RemoveAsset.md) |
| AddLiability    | Add a debt (loans, financing)                        | [Guide](Tasks/AddLiability.md) |
| RemoveLiability | Remove a paid-off debt                               | [Guide](Tasks/RemoveLiability.md) |
| AddGoal         | Add a financial goal with target and deadline        | [Guide](Tasks/AddGoal.md) |
| RemoveGoal      | Remove an achieved or abandoned goal                 | [Guide](Tasks/RemoveGoal.md) |

### Maintenance

| Task       | Description                                                       | Links |
| ---------- | ----------------------------------------------------------------- | ----- |
| CloseMonth | Close the month: totals → Year-Report, reset statements           | [Guide](Tasks/CloseMonth.md) |
| Render     | Re-render all dashboards without executing any action             | [Guide](Tasks/Render.md) |

---

## 3. What each dashboard is fed by

| Dashboard file | Rendered from |
| -------------- | ------------- |
| [`DashBoard/README.md`](DashBoard/README.md) | Everything — top-level position and alerts |
| [`DashBoard/Month/Statement.md`](DashBoard/Month/Statement.md) | AddTransaction, RemoveTransaction, AddTransfer, RemoveTransfer |
| [`DashBoard/Month/Accounts/*.md`](DashBoard/Accounts.md) | Transactions and transfers, per account |
| [`DashBoard/Month/DashBoard.md`](DashBoard/Month/DashBoard.md) | Transactions, budgets, recurring bills |
| [`DashBoard/Accounts.md`](DashBoard/Accounts.md) | AddAccount, RemoveAccount, transactions, transfers |
| [`DashBoard/Categories.md`](DashBoard/Categories.md) | AddCategory, RemoveCategory, SetBudget |
| [`DashBoard/Budget.md`](DashBoard/Budget.md) | SetBudget, AddReallocation, AddRecurringBill, RemoveRecurringBill |
| [`DashBoard/Net-Worth.md`](DashBoard/Net-Worth.md) | AddAsset/RemoveAsset, AddLiability/RemoveLiability, AddGoal/RemoveGoal |
| [`DashBoard/Year-Report.md`](DashBoard/Year-Report.md) | CloseMonth |

---

## 4. Rules to remember

- One task per tick — `Task.yaml` holds a single action.
- Dashboard files are **generated**: never edit them by hand (a hand edit is overwritten on
  the next tick; use `Render` to force a clean re-render).
- Transfers between own accounts are never income or expense — use `AddTransfer`, not
  `AddTransaction`.
- One category and one account per transaction — no splits.
- Category renames/merges and limit reviews happen at month close (`CloseMonth`), never mid-month.
