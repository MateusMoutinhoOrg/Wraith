# 🧭 Financial Dashboard — Wraith

> **Last update:** 16-aug-2026 · **Currency:** BRL (R$)
> **Scope:** Individual + MEI (Wraith Software)

---

## ⚡ One-Screen View

| Block                     | Number             | Status | Where to look                            |
| ------------------------- | ------------------ | ------ | ---------------------------------------- |
| Month result (actual)     | **-R$ 200**        | 🔴     | [Month/DashBoard](Month/DashBoard.md)    |
| Month result (projected)  | **+R$ 1,000**      | 🟢     | [Month/DashBoard](Month/DashBoard.md)    |
| Budget used               | 70% (R$ 700/1,000) | 🟡     | [Budget](Budget.md)                      |
| Liquid today (cash+bank)  | R$ 1,100           | 🟡     | [Month/All-Balance](Month/All-Balance.md)|
| 💵 Cash · 🏦 Bank           | R$ 137 · R$ 963    | 🔴     | [Month/All-Balance](Month/All-Balance.md)|
| 💳 Card bill (due 05-sep)  | R$ 380             | 🟡     | [Month/Credit-Card-Balance](Month/Credit-Card-Balance.md) |
| Net position              | **R$ 720**         | 🔴     | [Month/All-Balance](Month/All-Balance.md)|
| Lowest projected balance  | -R$ 380 (20-aug)   | 🔴     | [Cash-Flow](Cash-Flow.md)                |
| Net worth                 | **R$ 18,820**      | 🟡     | [Net-Worth](Net-Worth.md)       |
| Total debt                | R$ 5,180           | 🟡     | [Net-Worth](Net-Worth.md)       |
| Year 2026 accumulated     | +R$ 5,700          | 🟢     | [Year-Report](Year-Report.md)   |
| Health Score              | **48 / 100**       | 🟡     | [KPIs](KPIs.md)                 |
| FIRE progress             | 2.5%               | ⚪      | [Long-Term](Long-Term.md)       |
| Open alerts               | 4 🔴 · 5 🟡        | 🔴     | [Alerts-Rules](Alerts-Rules.md)          |

```
Year 2026  Income   █████████████▌░░░░░░  56%  (R$ 13,500 / R$ 24,000)
           Expense  ██████████████▍░░░░░  64%  (R$  7,800 / R$ 12,100)
           Time     ████████████▍░░░░░░░  62%  (16-aug / 31-dec)
```

---

## 📁 Dashboard Map

### 🔵 Short term — the current month (`Month/`)

One folder per month. Each account keeps its own statement; `All-Balance` consolidates them and
`DashBoard` turns them into KPIs, budget and actions.

| File                                                                    | Account       | What it answers                                            | Frequency |
| ----------------------------------------------------------------------- | ------------- | ---------------------------------------------------------- | --------- |
| [`Month/DashBoard.md`](Month/DashBoard.md)                              | all           | **Full month report** — accounts, KPIs, budget, actions     | Weekly    |
| [`Month/All-Balance.md`](Month/All-Balance.md)                          | all           | Consolidated statement — every transaction, tagged by account | Daily   |
| [`Month/Cash-Balance.md`](Month/Cash-Balance.md)                        | 💵 Cash        | What the wallet holds, and where it leaks                   | Daily     |
| [`Month/Bank-Balance.md`](Month/Bank-Balance.md)                        | 🏦 Bank        | Income, essentials and transfers in the checking account    | Daily     |
| [`Month/Credit-Card-Balance.md`](Month/Credit-Card-Balance.md)          | 💳 Credit-Card | How much is owed, when the cycle closes, when it is due     | Daily     |
| [`Cash-Flow.md`](Cash-Flow.md)                                          | all           | Will I run out of money? When?                              | Weekly    |

### 🟢 Planning — how much I can spend
| File                                       | What it answers                                     | Frequency  |
| ------------------------------------------ | --------------------------------------------------- | ---------- |
| [`Budget.md`](Budget.md)                   | **Envelope budget** — caps per category             | Monthly    |
| [`Categories.md`](Categories.md)           | Chart of accounts + category catalog and rules      | Quarterly  |
| [`Alerts-Rules.md`](Alerts-Rules.md)       | Rules that fire an automatic alert                  | Quarterly  |

### 🟡 Medium term — quarter and year
| File                                         | What it answers                                   | Frequency  |
| -------------------------------------------- | ------------------------------------------------- | ---------- |
| [`Quarter-Report.md`](Quarter-Report.md)     | Quarterly close and trends                        | Quarterly  |
| [`Year-Report.md`](Year-Report.md)           | 2026 consolidated month by month + closing forecast| Monthly   |
| [`Business-MEI.md`](Business-MEI.md)         | Is the business profitable? MEI cap, DAS, clients | Monthly    |
| [`Goals.md`](Goals.md)                       | Year goals and progress                           | Monthly    |

### 🔴 Long term — net worth
| File                               | What it answers                                      | Frequency  |
| ---------------------------------- | ---------------------------------------------------- | ---------- |
| [`Net-Worth.md`](Net-Worth.md)     | **Assets, liabilities, debts and portfolio**         | Monthly    |
| [`Long-Term.md`](Long-Term.md)     | 2026–2031 plan, financial independence, scenarios    | Semiannual |
| [`KPIs.md`](KPIs.md)               | Indicator dictionary + historical series             | Monthly    |

---

## 🔄 Maintenance Routine

| When            | What to do                                                                  |
| --------------- | --------------------------------------------------------------------------- |
| **Every day**   | Record each transaction in **its account's** statement under `Month/`, then mirror it into `Month/All-Balance.md` |
| **Every Mon.**  | Run `Cash-Flow.md`, check `Alerts-Rules.md`, update `Month/DashBoard.md`     |
| **Day 5**       | Pay the credit card bill **in full** from the bank (`Card-Payment` transfer) |
| **Day 1**       | Close the month: reconcile the 3 accounts → `Month/DashBoard.md` → `Year-Report.md` → `KPIs.md` → carry closing balances forward |
| **Day 1 (month)**| Review caps in `Budget.md`; reallocate idle envelopes                      |
| **End of qtr.** | `Quarter-Report.md` + review `Goals.md`                                     |
| **End of year** | Final `Year-Report.md` + recalibrate `Long-Term.md`                         |

---

## 🧾 Conventions (apply to every file)

| Item        | Standard                                          |
| ----------- | ------------------------------------------------- |
| Date        | `DD-mmm-YYYY` (e.g., `16-aug-2026`)               |
| Value       | `R$ 1,234.56` · outflows always signed negative   |
| Category    | `Capitalized-single-word` per `Categories.md`     |
| Account     | 💵 `Cash` · 🏦 `Bank` · 💳 `Credit-Card` — exactly 1 per transaction |
| Transfer    | Two legs that net to zero. Never an expense, never touches a budget envelope |
| Card spend  | Expense on the **purchase** date; hits liquid only on the **payment** date |
| Projection  | `*` suffix on the number                          |
| Status      | 🟢 ok · 🟡 attention · 🔴 immediate action · ⚪ inactive |
| Trend       | ↗ rising · ↘ falling · → stable                   |

> 📌 *Fictional dashboard, used as a model for personal + small-business financial management.*
