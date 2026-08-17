# 🚨 Alerts and Rules

> **Last scan:** 16-aug-2026 · **Open:** 3 🔴 · 3 🟡 · 1 ⚪ · 1 ℹ️

---

## 1. 🔔 Open Alerts

| Sev. | ID     | Alert                                                           | Fired on     | Rule  | Action                                        | Deadline |
| ---- | ------ | --------------------------------------------------------------- | ------------ | ----- | --------------------------------------------- | -------- |
| 🔴   | `A-01` | `Vices` blew its cap by 80% (R$ 90 / R$ 50)                     | 14-aug-2026  | R-03  | Freeze category until 01-sep                  | 18-aug   |
| 🔴   | `A-02` | 100% of income concentrated in 1 client                         | 01-aug-2026  | R-11  | Send 3 proposals                              | 31-aug   |
| 🔴   | `A-03` | Projected cash trough: R$ 289 on 24-aug                         | 16-aug-2026  | R-06  | Pull invoice forward or defer subscription    | 20-aug   |
| 🟡   | `A-04` | Spending pace 18 pp ahead of the calendar                       | 12-aug-2026  | R-02  | Reallocate envelopes (see `Budget.md` §3)     | 18-aug   |
| 🟡   | `A-05` | Runway stuck at 3.2 months for 3 months                         | 01-aug-2026  | R-05  | Execute the R$ 500 contribution               | 30-aug   |
| 🟡   | `A-06` | Expensive debt open: R$ 380 at 14%/mo                           | 01-aug-2026  | R-08  | Pay the bill in full                          | 21-aug   |
| ⚪   | `A-07` | Category `Study` with no movement for 2 months                  | 01-aug-2026  | R-04  | Reallocate R$ 50 or make a commitment         | 31-aug   |
| ℹ️   | `A-08` | `AddCategory Poker` returned an error — category already exists | 09-aug-2026  | R-13  | Use the existing category; no action needed   | —        |

```
Alerts by severity
🔴 Critical  ███████████████  3
🟡 Attention ███████████████  3
⚪ Info      █████            1
ℹ️ System    █████            1
```

---

## 2. 📜 Rule Catalog

### 📦 Budget

| ID   | Rule                                                         | Trigger                          | Sev. | Automatic action             |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-01 | Envelope above 80% of the cap                                | usage ≥ 80%                      | 🟡   | Notify                       |
| R-02 | Spending pace ahead of the calendar                          | %used − %month ≥ 10 pp           | 🟡   | Suggest reallocation         |
| R-03 | **Envelope over cap**                                        | usage > 100%                     | 🔴   | **Freeze the category**      |
| R-04 | Idle envelope                                                | 0 movements for 2 months         | ⚪   | Suggest reallocation         |
| R-14 | Global monthly cap exceeded                                  | expense > R$ 1,000               | 🔴   | Block discretionary spending |

### 💧 Liquidity

| ID   | Rule                                                         | Trigger                          | Sev. | Automatic action             |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-05 | Runway below target                                          | reserve < 6 months of expense    | 🟡   | Notify monthly               |
| R-06 | **Cash trough**                                              | projected balance < R$ 300       | 🔴   | Contingency plan             |
| R-07 | Projected negative cash                                      | projected balance < R$ 0         | 🔴   | Withdraw from the reserve    |
| R-15 | Overdue receivable                                           | invoice unpaid for 5 days        | 🔴   | Chase the client             |

### 💳 Debt

| ID   | Rule                                                         | Trigger                          | Sev. | Automatic action             |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-08 | **Expensive debt open**                                      | interest > 3%/mo and balance > R$ 0 | 🟡| Priority over contributions  |
| R-09 | High income commitment                                       | installments > 30% of income     | 🔴   | Forbid new debt              |
| R-10 | Installments with an incomplete reserve                      | reserve < 6 m and new installment| 🔴   | **Block the purchase**       |

### 💼 Business

| ID   | Rule                                                         | Trigger                          | Sev. | Automatic action             |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-11 | **Client concentration**                                     | largest client > 60% of income   | 🔴   | Prospecting target           |
| R-12 | Insufficient pipeline                                        | weighted < 300% of quarter target| 🟡   | Prospecting block on calendar|
| R-16 | MEI cap at risk                                              | 12-month revenue > 80% of R$ 81k | 🟡   | Plan a tax regime migration  |
| R-17 | Month closed in the red                                      | result < R$ 0                    | 🔴   | Cap review the following month |

### 🧾 Data Integrity

| ID   | Rule                                                         | Trigger                               | Sev. | Automatic action                   |
| ---- | ------------------------------------------------------------ | ------------------------------------- | ---- | ---------------------------------- |
| R-13 | Duplicate category                                           | `AddCategory` on an existing category | ℹ️   | `Error: Category X already exists` |
| R-18 | Transaction without a category                               | uncategorized line in the statement   | 🟡   | Block the month close              |
| R-19 | Statement out of date                                        | no entry for 3 days                   | 🟡   | Reminder                           |
| R-20 | Statement balance ≠ bank balance                             | divergence > R$ 1                     | 🔴   | Reconcile before closing           |

---

## 3. 🧊 Active Freezes

| Category  | Frozen on    | Releases on | Reason                          | Rule  |
| --------- | ------------ | ----------- | ------------------------------- | ----- |
| `Vices`   | 18-aug-2026  | 01-sep-2026 | 80% overrun on the cap          | R-03  |

> A frozen category **accepts no new entries** until the release date. The only exception is a formal reallocation recorded in [`Budget.md`](Budget.md) §3.

---

## 4. 🧯 Escalation Matrix

| Situation                                   | Level | Who decides  | Action                                            |
| ------------------------------------------- | ----- | ------------ | ------------------------------------------------- |
| 1 envelope over cap                         | 🟡    | Automatic    | Freeze category, reallocate from an idle envelope |
| 2+ envelopes over cap                       | 🔴    | Review       | Freeze all discretionary ones until day 1         |
| Projected cash < R$ 300                     | 🔴    | Review       | Defer non-essential expenses, pull receivable forward |
| Projected cash < R$ 0                       | 🔴    | Emergency    | Withdraw from the reserve (return within ≤ 30 days) |
| Month closed in the red                     | 🔴    | Review       | Cut 15% off the cap the following month           |
| 2 months in a row in the red                | 🔴    | Emergency    | Freeze all discretionary spending for 60 days     |
| Loss of the main client                     | 🔴    | Emergency    | Survival mode: fixed costs + food + transport only|

**Survival mode** = monthly cap of R$ 700 (`Home` R$ 300 + `Food` R$ 250 + `Transport` R$ 100 + `Business` R$ 50), everything else frozen. Extends the runway from 3.2 → 4.6 months.

---

## 5. 📋 Month-Close Checklist

Run on the 1st of every month, in order:

- [ ] Reconcile `Month-Balance.md` against the bank statement (rule R-20)
- [ ] Verify that every transaction has a category (rule R-18)
- [ ] Consolidate totals in `Month-Results.md`
- [ ] Update `Month-Report.md` with the closed numbers
- [ ] Record the month in `Year-Report.md` and `KPIs.md`
- [ ] Update balances in `Net-Worth.md`
- [ ] Review caps and reallocations in `Budget.md`
- [ ] Release frozen categories whose freeze has expired
- [ ] Run every rule in this file and log new alerts
- [ ] Mark progress in `Goals.md`
- [ ] Archive the statement and reset `Month-Balance.md`

---

## 6. 🗂️ Resolved Alert History

| ID     | Alert                                     | Opened      | Closed      | How it was resolved              |
| ------ | ----------------------------------------- | ----------- | ----------- | -------------------------------- |
| `A-00` | Overdraft in use                          | 05-jan-2026 | 18-feb-2026 | Paid off with February's result  |
| `J-04` | July closed in the red                    | 01-aug-2026 | —           | 🔴 Still under review            |
| `J-02` | `Food` blew its cap in July               | 28-jul-2026 | 01-aug-2026 | Cap renewed in the new month     |
| `M-07` | Uncategorized transaction in May          | 31-may-2026 | 01-jun-2026 | Classified as `Transport`        |

---

## 🔗 Related
[`Budget.md`](Budget.md) · [`Cash-Flow.md`](Cash-Flow.md) · [`Categories.md`](Categories.md) · [`Month-Report.md`](Month-Report.md)
