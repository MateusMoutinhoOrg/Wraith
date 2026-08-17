# 🏦 Bank Balance — August / 2026

> **Account:** `Bank` — checking account · **Type:** Liquid asset
> **Period:** 01-aug-2026 → 31-aug-2026 · **Partial close:** 16-aug-2026 (52% of the month)
> **Status:** 🟡 **-R$ 37 in 9 movements** — the only account that receives income, and the only one that pays the card.

---

## 1. 📌 Account Summary

| Line                         | Value       | Movements | Status |
| ---------------------------- | ----------: | --------: | ------ |
| Opening balance (01-aug)     |  R$ 1,000   |         — | 🟡     |
| **(+) Inflows**              |    R$ 500   |         1 | 🔴     |
| **(−) Outflows (expenses)**  |    R$ 157   |         7 | 🟢     |
| **(−) Outflows (transfers)** |    R$ 380   |         1 | 🟡     |
| **(=) Account result**       |  **-R$ 37** |         9 | 🟡     |
| **Closing balance (day 16)** | **R$ 963**  |         — | 🟡     |

> ℹ️ Of the R$ 537 that left the account, **R$ 380 (71%) was not an expense** — it was the payment of
> July's credit card bill, a transfer to [`Credit-Card-Balance.md`](Credit-Card-Balance.md).

---

## 2. 📒 Statement — Realized (01-aug → 16-aug)

| Date        | Transaction                  | Category    |       Value |   Balance |
| ----------- | ---------------------------- | ----------- | ----------: | --------: |
| 01-aug-2026 | *Opening balance*            | —           |           — |  R$ 1,000 |
| 02-aug-2026 | Supermarket — weekly shop    | `Food`      |      -R$ 45 |    R$ 955 |
| 05-aug-2026 | **Card bill — jul cycle**    | `Card-Payment`  |    **-R$ 380** |  R$ 575 |
| 07-aug-2026 | Supermarket                  | `Food`      |      -R$ 32 |    R$ 543 |
| 08-aug-2026 | Ride app                     | `Transport` |      -R$ 12 |    R$ 531 |
| 11-aug-2026 | **Client invoice — MEI**     | `Freelance` | **+R$ 500** |  R$ 1,031 |
| 12-aug-2026 | Supermarket                  | `Food`      |      -R$ 28 |  R$ 1,003 |
| 12-aug-2026 | Ride app                     | `Transport` |      -R$ 10 |    R$ 993 |
| 15-aug-2026 | Supermarket                  | `Food`      |      -R$ 22 |    R$ 971 |
| 15-aug-2026 | Fuel                         | `Transport` |       -R$ 8 |    R$ 963 |

> 📉 **Low point:** R$ 531 on 08-aug, three days before the client invoice cleared.

---

## 3. 🔁 Transfers Touching This Account

| Date        | Direction | Counterparty                | Value    | Note                          |
| ----------- | --------- | --------------------------- | -------: | ----------------------------- |
| 05-aug-2026 | ➡️ out    | `Credit-Card`               | -R$ 380  | July bill paid in full — no interest |
| 05-sep-2026 | ➡️ out    | `Credit-Card`               | -R$ 380\*| August bill, due day 5        |
| 30-aug-2026 | ➡️ out    | `Reserve` (Net-Worth)       | -R$ 500\*| Automatic monthly contribution|

> ⚠️ Transfers are **not expenses**. They move money between accounts and are excluded from every
> category and budget total in [`DashBoard.md`](DashBoard.md).

---

## 4. 🧮 Totals by Category (day 16)

| Category    | Mov. | Inflows | Outflows |      Net | % of account expense |
| ----------- | ---: | ------: | -------: | -------: | -------------------: |
| `Freelance` |    1 |  R$ 500 |    R$ 0  | +R$ 500  |                   —  |
| `Food`      |    4 |   R$ 0  |   R$ 127 | -R$ 127  |               80.9%  |
| `Transport` |    3 |   R$ 0  |    R$ 30 |  -R$ 30  |               19.1%  |
| `Card-Payment`  |    1 |   R$ 0  |   R$ 380 | -R$ 380  |         *excluded*   |
| **Total**   | **9** | **R$ 500** | **R$ 537** | **-R$ 37** | **100%** |

```
Bank expense distribution (transfers excluded)
Food        ████████████████████████  80.9%
Transport   █████▋                    19.1%
```

---

## 5. 📅 Scheduled — Not Yet Realized (17-aug → 31-aug)

| Date        | Transaction                 | Category    |        Value | Confidence |
| ----------- | --------------------------- | ----------- | -----------: | ---------- |
| 25-aug-2026 | **Client receivable — MEI** | `Freelance` | +R$ 1,500\*  | 🟡 Medium  |
| 17→31-aug   | Supermarket at pace         | `Food`      |   -R$ 100\*  | 🟢 High    |
| 17→31-aug   | Transport at pace           | `Transport` |    -R$ 40\*  | 🟢 High    |
| 28-aug-2026 | DAS — MEI                   | `Business`  |    -R$ 20\*  | 🟢 Certain |
| 30-aug-2026 | Reserve contribution        | `Reserve`    |   -R$ 500\*  | 🟢 Certain |
| **Total**   |                             |             | **+R$ 840\*** | —         |

**Projected closing balance (31-aug): R$ 1,803\*** 🟢 — *entirely dependent on the day-25 receivable.*
Without it: **R$ 303\*** 🔴, not enough to cover the R$ 380 card bill due 05-sep.

---

## 6. 🚩 Reading This Account

| Finding                                                          | Impact                 | Status |
| ---------------------------------------------------------------- | ---------------------- | ------ |
| Single inflow — 100% of income lands from one client              | Concentration risk     | 🔴     |
| Card bill consumed 38% of the opening balance on day 5            | Front-loaded outflow   | 🟡     |
| Balance never went negative — no overdraft used                   | No overdraft fee       | 🟢     |
| Only 7 expense movements — essentials are consolidated here       | Low fragmentation      | 🟢     |
| Sep-05 card bill (R$ 380) not yet funded if day 25 slips          | Cascade to card debt   | 🔴     |

---

## 🔗 Related

[`All-Balance.md`](All-Balance.md) · [`Cash-Balance.md`](Cash-Balance.md) · [`Credit-Card-Balance.md`](Credit-Card-Balance.md) · [`DashBoard.md`](DashBoard.md) · [`../Cash-Flow.md`](../Cash-Flow.md)

---

## 📖 Legend

| Symbol | Meaning                                        |
| ------ | ---------------------------------------------- |
| 🟢     | On target / healthy                            |
| 🟡     | Attention — between 70% and 100% of the limit  |
| 🔴     | Off target / requires immediate action         |
| ⚪     | Inactive / no movement                         |
| `*`    | Projected value, not realized                  |

> 📌 *Fictional document, used as a model for personal + small-business financial management.*
