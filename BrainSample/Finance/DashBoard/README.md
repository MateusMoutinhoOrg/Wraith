# Financial Dashboard

> **Updated:** 16-aug-2026 · **Currency:** BRL (R$) · **Scope:** personal + Wraith Software (MEI)

---

## 1. Position

| Indicator                       |           Value | Status |
| ------------------------------- | --------------: | :----: |
| Available balance (cash + bank) |    **R$ 2,250** |   🟡   |
| Credit card outstanding         |        R$ 1,350 |   🟡   |
| **Net available**               |      **R$ 900** |   🔴   |
| Month result (day 16)           |         -R$ 900 |   🔴   |
| Month result — projected        |      +R$ 2,000* |   🟢   |
| Budget used                     | 70% of R$ 3,000 |   🟡   |
| Net worth                       |       R$ 10,000 |   🟡   |
| Total debt                      |        R$ 7,750 |   🟡   |
| Saved year-to-date              |       +R$ 7,550 |   🟢   |

```
Year to date
Income        ███████████▏░░░░░░░░  56%   R$ 30,000 of R$ 54,000 planned
Expenses      ████████████▍░░░░░░░  62%   R$ 22,450 of R$ 36,050 planned
Year elapsed  ████████████▍░░░░░░░  62%
```

---

## 2. Active alerts

| Status | Alert                                            |
| :----: | ------------------------------------------------ |
|   🔴   | `Vices` over budget — R$ 180 spent, limit R$ 100 |
|   🟡   | Net available dips to ~R$ 400 around 20-aug      |

Full list in [`Month/DashBoard.md`](Month/DashBoard.md).

---

## 3. Structure

| File                                                    | Contents                                          | Review    |
| ------------------------------------------------------- | ------------------------------------------------- | --------- |
| [`Accounts.md`](Accounts.md)                            | Account registry, balances, links to statements   | Weekly    |
| [`Categories.md`](Categories.md)                        | Category registry, limits, month usage            | Monthly   |
| [`Budget.md`](Budget.md)                                | Active budgets, allocation targets, recurring bills | Monthly |
| [`Month/DashBoard.md`](Month/DashBoard.md)              | Current month: result, balances, alerts           | Weekly    |
| [`Month/Statement.md`](Month/Statement.md)              | Consolidated ledger — every transaction           | Daily     |
| [`Month/Accounts/Bank.md`](Month/Accounts/Bank.md)      | Bank account statement (isolated)                 | Weekly    |
| [`Month/Accounts/Cash.md`](Month/Accounts/Cash.md)      | Cash statement (isolated)                         | Weekly    |
| [`Month/Accounts/Credit-Card.md`](Month/Accounts/Credit-Card.md) | Credit card statement (isolated)         | Weekly    |
| [`Net-Worth.md`](Net-Worth.md)                          | Assets, liabilities, investments, goals           | Monthly   |
| [`Year-Report.md`](Year-Report.md)                      | Month-by-month results + business figures         | Monthly   |

---

## 4. Routine

| When         | Task                                                                       |
| ------------ | -------------------------------------------------------------------------- |
| Daily        | Record each transaction in [`Month/Statement.md`](Month/Statement.md)      |
| Monday       | Review [`Month/DashBoard.md`](Month/DashBoard.md) and the active alerts    |
| Day 5        | Pay the card bill in full from the bank account                            |
| Day 1        | Close the previous month: totals → `Year-Report.md`, reset the statements  |

---

## 5. Conventions

| Notation      | Meaning                                                                    |
| ------------- | -------------------------------------------------------------------------- |
| 🟢 🟡 🔴 ⚪    | on track · attention · action required · no activity                       |
| `*`           | projected value, not yet realized                                          |
| Category      | classification of a transaction (`Food`, `Business`, …) — see [`Categories.md`](Categories.md) |
| Transfer      | movement between own accounts — never counted as income or expense         |
| Card purchase | counted as expense on purchase date; leaves the bank on the bill due date  |

> *Fictional dashboard, used as a model for personal + small-business financial management. All files are generated — do not edit by hand.*
