# 🧭 Financial Dashboard — Wraith

> **Last update:** 16-aug-2026 · **Currency:** BRL (R$)
> **Scope:** Individual + MEI (Wraith Software)

---

## ⚡ One-Screen View

| Block                     | Number             | Status | Where to look                   |
| ------------------------- | ------------------ | ------ | ------------------------------- |
| Month result (actual)     | **-R$ 200**        | 🔴     | [Month-Report](Month-Report.md) |
| Month result (projected)  | **+R$ 1,000**      | 🟢     | [Month-Report](Month-Report.md) |
| Budget used               | 70% (R$ 700/1,000) | 🟡     | [Budget](Budget.md)             |
| Cash today                | R$ 1,100           | 🟡     | [Cash-Flow](Cash-Flow.md)       |
| Lowest projected balance  | -R$ 380 (20-aug)   | 🔴     | [Cash-Flow](Cash-Flow.md)       |
| Net worth                 | **R$ 18,820**      | 🟡     | [Net-Worth](Net-Worth.md)       |
| Total debt                | R$ 5,180           | 🟡     | [Net-Worth](Net-Worth.md)       |
| Year 2026 accumulated     | +R$ 5,700          | 🟢     | [Year-Report](Year-Report.md)   |
| Health Score              | **48 / 100**       | 🟡     | [KPIs](KPIs.md)                 |
| FIRE progress             | 2.5%               | ⚪      | [Long-Term](Long-Term.md)       |
| Open alerts               | 3 🔴 · 3 🟡        | 🔴     | [Alerts-Rules](Alerts-Rules.md) |

```
Year 2026  Income   █████████████▌░░░░░░  56%  (R$ 13,500 / R$ 24,000)
           Expense  ██████████████▍░░░░░  64%  (R$  7,800 / R$ 12,100)
           Time     ████████████▍░░░░░░░  62%  (16-aug / 31-dec)
```

---

## 📁 Dashboard Map

### 🔵 Short term — the current month
| File                                     | What it answers                                                           | Frequency  |
| -------------------------------------- | ------------------------------------------------------------------------- | --------- |
| [`Month-Results.md`](Month-Results.md)   | Raw month numbers (actual vs. projected)                                  | Daily      |
| [`Month-Balance.md`](Month-Balance.md)   | Statement — every transaction of the month, line by line                  | Daily   |
| [`Month-Report.md`](Month-Report.md)     | **Full month report** — KPIs, budget, actions                             | Weekly     |
| [`Cash-Flow.md`](Cash-Flow.md)           | Will I run out of money? When?                                            | Weekly     |

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
| **Every day**   | Record transactions in `Month-Balance.md`                                   |
| **Every Mon.**  | Run `Cash-Flow.md`, check `Alerts-Rules.md`, update `Month-Report.md`        |
| **Day 1**       | Close the month: `Month-Results.md` → `Year-Report.md` → `KPIs.md` → reset statement |
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
| Projection  | `*` suffix on the number                          |
| Status      | 🟢 ok · 🟡 attention · 🔴 immediate action · ⚪ inactive |
| Trend       | ↗ rising · ↘ falling · → stable                   |

> 📌 *Fictional dashboard, used as a model for personal + small-business financial management.*
