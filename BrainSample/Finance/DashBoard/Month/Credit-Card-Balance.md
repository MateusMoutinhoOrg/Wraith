# 💳 Credit Card Balance — August / 2026

> **Account:** `Credit-Card` — Mastercard · **Type:** Liability (revolving credit)
> **Limit:** R$ 2,000 · **Cycle closes:** day 28 · **Due:** day 5 of the next month
> **Period:** 01-aug-2026 → 31-aug-2026 · **Partial close:** 16-aug-2026 (52% of the month)
> **Status:** 🟡 **R$ 380 owed** — paid in full every month so far, so no interest has been charged.

---

## 1. 📌 Account Summary

| Line                              | Value       | Movements | Status |
| --------------------------------- | ----------: | --------: | ------ |
| Opening bill (01-aug, jul cycle)  |    R$ 380   |         — | 🟡     |
| **(+) Purchases in the month**    |    R$ 380   |         5 | 🟡     |
| **(−) Payments received**         |    R$ 380   |         1 | 🟢     |
| **(=) Outstanding bill (day 16)** | **R$ 380**  |         6 | 🟡     |
| Limit used                        | 19% (R$ 380 / R$ 2,000) | — | 🟢 |
| Interest charged this month       |      R$ 0   |         — | 🟢     |

> ⚠️ On this card the invoice rate is **14%/month (≈ 380%/year)** — the most expensive debt in
> [`../Net-Worth.md`](../Net-Worth.md). It costs nothing **only** while the bill is paid in full.

---

## 2. 📒 Statement — Realized (01-aug → 16-aug)

*A purchase increases the bill; a payment reduces it. `Balance` = amount owed.*

| Date        | Transaction               | Category    |       Value |  Balance |
| ----------- | ------------------------- | ----------- | ----------: | -------: |
| 01-aug-2026 | *Opening bill — jul cycle*| —           |           — |   R$ 380 |
| 01-aug-2026 | Rent + utilities          | `Home`      |     +R$ 300 |   R$ 680 |
| 03-aug-2026 | Fuel                      | `Transport` |      +R$ 15 |   R$ 695 |
| 05-aug-2026 | Cloud hosting — Wraith    | `Business`  |      +R$ 20 |   R$ 715 |
| 05-aug-2026 | **Payment — jul bill**    | `Card-Payment`  | **-R$ 380** |   R$ 335 |
| 10-aug-2026 | Lunch out                 | `Food`      |      +R$ 25 |   R$ 360 |
| 13-aug-2026 | Domain renewal            | `Business`  |      +R$ 20 |   R$ 380 |

---

## 3. 🔁 Transfers Touching This Account

| Date        | Direction | Counterparty | Value    | Note                                |
| ----------- | --------- | ------------ | -------: | ----------------------------------- |
| 05-aug-2026 | ⬅️ in     | `Bank`       | R$ 380   | July bill settled in full 🟢        |
| 05-sep-2026 | ⬅️ in     | `Bank`       | R$ 380\* | August bill — **must** be funded    |

> 💡 Paying the bill is a **transfer**, not an expense. The expense was already booked on the
> purchase date, in its own category.

---

## 4. 🧮 Totals by Category (day 16)

| Category    | Mov. | Purchases |     Net | % of card spend |
| ----------- | ---: | --------: | ------: | --------------: |
| `Home`      |    1 |    R$ 300 | -R$ 300 |           78.9% |
| `Business`  |    2 |     R$ 40 |  -R$ 40 |           10.5% |
| `Food`      |    1 |     R$ 25 |  -R$ 25 |            6.6% |
| `Transport` |    1 |     R$ 15 |  -R$ 15 |            3.9% |
| **Total**   | **5** | **R$ 380** | **-R$ 380** | **100%** |

```
Card spend distribution
Home        ██████████████████████  78.9%
Business    ██▉                     10.5%
Food        █▉                       6.6%
Transport   █                        3.9%
```

> 🟢 **Zero `Vices` and zero `Poker` on the card** — discretionary spending is confined to
> [`Cash-Balance.md`](Cash-Balance.md), which keeps it visible and capped by the wallet.

---

## 5. 📅 Scheduled — Not Yet Realized (17-aug → 31-aug)

| Date        | Transaction                       | Category | Value     | Confidence |
| ----------- | --------------------------------- | -------- | --------: | ---------- |
| 28-aug-2026 | *Cycle closes* — bill for sep     | —        |         — | 🟢 Certain |
| 30-aug-2026 | Card interest (only if revolving) | `Debt`   | -R$ 20\*  | 🟡 Medium  |

**Projected bill at close (28-aug): R$ 380\*** · **due 05-sep-2026.**
Add **R$ 20\*** of interest if the bill is not settled in full.

---

## 6. 🚩 Reading This Account

| Finding                                                        | Impact                         | Status |
| --------------------------------------------------------------- | ------------------------------ | ------ |
| Bill paid in full in aug — R$ 0 interest                        | Free 30-day float              | 🟢     |
| Limit at 19% — plenty of headroom                               | No utilization pressure        | 🟢     |
| 79% of the bill is a single fixed cost (`Home`)                 | Predictable, easy to fund      | 🟢     |
| Sep-05 payment depends on the day-25 receivable                 | One slip ⇒ 14%/mo revolving 🔴 | 🔴     |

> 🔻 **Single point of failure:** if Client A's R$ 1,500 slips past 05-sep, the R$ 380 bill revolves
> and the cheapest month turns into the most expensive debt in the portfolio.

---

## 🔗 Related

[`All-Balance.md`](All-Balance.md) · [`Bank-Balance.md`](Bank-Balance.md) · [`Cash-Balance.md`](Cash-Balance.md) · [`DashBoard.md`](DashBoard.md) · [`../Net-Worth.md`](../Net-Worth.md)

---

## 📖 Legend

| Symbol | Meaning                                        |
| ------ | ---------------------------------------------- |
| 🟢     | On target / healthy                            |
| 🟡     | Attention — between 70% and 100% of the limit  |
| 🔴     | Off target / requires immediate action         |
| `*`    | Projected value, not realized                  |

> 📌 *Fictional document, used as a model for personal + small-business financial management.*
