# 💧 Cash Flow

> **Question this file answers:** *will I run out of money? When?*
> **Cash today (16-aug-2026):** R$ 1,100 · **Burn rate:** R$ 43.75/day · **Breathing room:** 25 days

---

## 1. 🚨 Cash Traffic Light

| Metric                          | Value      | Limit       | Status |
| ------------------------------- | ---------: | ----------: | ------ |
| Cash available today            | R$ 1,100   | ≥ R$ 1,000  | 🟡     |
| **Lowest projected balance (month)** | **-R$ 380**| ≥ R$ 300 | 🔴     |
| Trough date                     | 20-aug     | —           | 🔴     |
| Cash projected on 31-aug        | R$ 1,600   | ≥ R$ 1,000  | 🟢     |
| Open receivables                | R$ 1,500   | —           | 🟡     |
| Bills payable through 31-aug    | R$ 300     | —           | 🟢     |

> 🔴 **Liquidity alert:** between 20-aug and 25-aug cash goes negative by R$ 380.
> **Ways out:** pull Client A's invoice forward, defer the cloud subscription (R$ 60), or take R$ 400 from the reserve (return it on 26-aug).

---

## 2. 📆 Daily Projection — 2nd Half of August

| Date        | Description                  | Category   | Type    |    Value | Balance     | Level |
| ----------- | ---------------------------- | ---------- | ------- | -------: | ----------: | ----- |
| 16-aug-2026 | *Opening balance*            | —          | —       |        — |  R$ 1,100   | 🟡    |
| 17-aug-2026 | Cigarettes                   | `Vices`    | Expense |   -R$ 10 |  R$ 1,090   | 🟡    |
| 18-aug-2026 | Groceries (biweekly)         | `Food`     | Expense |  -R$ 120 |    R$ 970   | 🟡    |
| 19-aug-2026 | Fuel                         | `Transport`| Expense |   -R$ 40 |    R$ 930   | 🟡    |
| 20-aug-2026 | Cloud/VPS + domains          | `Business` | Expense |   -R$ 85 |    R$ 845   | 🟡    |
| 20-aug-2026 | DAS-MEI                      | `Business` | Expense |   -R$ 76 |    R$ 769   | 🟡    |
| 21-aug-2026 | Credit card bill             | `Debt`     | Expense |  -R$ 380 |    R$ 389   | 🔴    |
| 22-aug-2026 | Lunches for the week         | `Food`     | Expense |   -R$ 60 |    R$ 329   | 🔴    |
| 23-aug-2026 | Streaming                    | `Poker`    | Expense |   -R$ 30 |    R$ 299   | 🔴    |
| 24-aug-2026 | Cigarettes                   | `Vices`    | Expense |   -R$ 10 |    R$ 289   | 🔴    |
| 25-aug-2026 | **Client A invoice — project**| `Freelance`| Income |+R$ 1,500 |  R$ 1,789   | 🟢    |
| 27-aug-2026 | Groceries                    | `Food`     | Expense |   -R$ 90 |  R$ 1,699   | 🟢    |
| 28-aug-2026 | Fuel                         | `Transport`| Expense |   -R$ 30 |  R$ 1,669   | 🟢    |
| 30-aug-2026 | **Reserve contribution**     | `Reserve`  | Transfer|  -R$ 500 |  R$ 1,169   | 🟢    |
| 30-aug-2026 | **Investment contribution**  | `Invest`   | Transfer|  -R$ 200 |    R$ 969   | 🟡    |
| 31-aug-2026 | Misc / slack                 | —          | Expense |   -R$ 50 |    R$ 919   | 🟡    |
|             | **Closing balance**          |            |         |          | **R$ 919**  | 🟡    |

```
Cash curve — 2nd half
R$1,800 ┤                                    ▇▇▇▇
R$1,400 ┤                                    ████▇▇▇
R$1,000 ┼▇▇▇▇▇▇▇                             ███████▇▇▇
R$  600 ┤███████▇▇                           ██████████
R$  200 ┤██████████▇▇▇▇▇                     ██████████
R$    0 ┼───────────────▁▁▁▁▁────────────────██████████
        └ 16  18  20  21  22  24  25  27  29  31
                       ▲ trough: R$ 289 (day 24)
```

> ℹ️ The "-R$ 380" in the traffic light assumes the scenario **without** deferring the card bill; in the projection above the bill lands on day 21 and the real floor is **R$ 289** on 24-aug — a 6-day margin with no slack at all.

---

## 3. 📥 Open Receivables

| Client / Source  | Description          | Issued      | Due         |    Value | Prob. | Status      |
| ---------------- | -------------------- | ----------- | ----------- | -------: | ----: | ----------- |
| Client A         | Project — phase 2    | 20-aug-2026 | 25-aug-2026 | R$ 1,500 |   95% | 🟡 To issue |
| Client B         | Proposal sent        | —           | sep/2026    | R$ 1,200 |   40% | ⚪ Pipeline |
| Client C         | Prospecting          | —           | oct/2026    |   R$ 800 |   20% | ⚪ Pipeline |
| **Total**        | —                    | —           | —           | **R$ 3,500** | — | —      |
| **Weighted**     | —                    | —           | —           | **R$ 2,065** | — | —      |

⚠️ **Single point of failure:** 95% of September's cash depends on one receivable. If Client A is 10 days late, the month closes negative.

---

## 4. 📤 Bills Payable

| Due date    | Description          | Category   |  Value | Auto | Status     |
| ----------- | -------------------- | ---------- | -----: | ---- | ---------- |
| 20-aug-2026 | Cloud / VPS          | `Business` |  R$ 60 | ✅   | 🟡 Scheduled|
| 20-aug-2026 | DAS-MEI              | `Business` |  R$ 76 | ✅   | 🟡 Scheduled|
| 21-aug-2026 | Credit card bill     | `Debt`     | R$ 380 | ❌   | 🔴 Manual  |
| 23-aug-2026 | Streaming            | `Poker`    |  R$ 30 | ✅   | 🟡 Scheduled|
| 05-sep-2026 | Rent                 | `Home`     | R$ 300 | ✅   | ⚪ Future  |
| 10-sep-2026 | Internet             | `Home`     |  R$ 90 | ✅   | ⚪ Future  |
| **Total through 31-aug** | —       | —          | **R$ 546** | — | —      |

---

## 5. 🗓️ Rolling 12-Month Cash Flow

| Month     | Inflows   | Outflows  | Net        | Ending cash | Status |
| --------- | --------: | --------: | ---------: | ----------: | ------ |
| sep-2025  |  R$ 1,400 |    R$ 950 |   +R$ 450  |    R$ 1,900 | 🟢     |
| oct-2025  |  R$ 1,600 |    R$ 980 |   +R$ 620  |    R$ 2,520 | 🟢     |
| nov-2025  |  R$ 1,300 |  R$ 1,050 |   +R$ 250  |    R$ 2,770 | 🟢     |
| dec-2025  |  R$ 2,100 |  R$ 1,400 |   +R$ 700  |    R$ 3,470 | 🟢     |
| jan-2026  |  R$ 1,500 |    R$ 900 |   +R$ 600  |    R$ 2,070 | 🟢     |
| feb-2026  |  R$ 1,700 |    R$ 880 |   +R$ 820  |    R$ 1,890 | 🟢     |
| mar-2026  |  R$ 1,800 |  R$ 1,000 |   +R$ 800  |    R$ 1,690 | 🟢     |
| apr-2026  |  R$ 1,500 |    R$ 920 |   +R$ 580  |    R$ 1,470 | 🟡     |
| may-2026  |  R$ 1,800 |    R$ 950 |   +R$ 850  |    R$ 1,620 | 🟢     |
| jun-2026  |  R$ 2,200 |  R$ 1,100 | +R$ 1,100  |    R$ 1,920 | 🟢     |
| jul-2026  |  R$ 1,000 |  R$ 1,050 |    -R$ 50  |    R$ 1,170 | 🔴     |
| aug-2026* |  R$ 2,000 |  R$ 1,000 | +R$ 1,000  |    R$ 919   | 🟡     |

> Cash falls even with a positive result because contributions to `Reserve`/`Invest` leave the cash account — this is **healthy**, the money migrated into net worth (see [`Net-Worth.md`](Net-Worth.md)).

```
Monthly net
sep ███▌      oct ████▉      nov ██        dec █████▌
jan ████▊     feb ██████▌    mar ██████▍   apr ████▋
may ██████▊   jun ████████▊  jul ▍🔴       aug ████████
```

---

## 6. 🧯 Contingency Plan

| Scenario                         | Cash impact      | Plan                                                    |
| -------------------------------- | ---------------: | ------------------------------------------------------- |
| Client A is 10 days late         |        -R$ 1,500 | Withdraw R$ 800 from the reserve, return it in 5 days    |
| Client A cancels                 |        -R$ 1,500 | Freeze `Vices`+`Poker`+`Study`, activate pipeline B/C   |
| Emergency expense of R$ 1,000    |        -R$ 1,000 | The reserve covers it — that's exactly what it's for     |
| Income drops for 3 months        |        -R$ 3,000 | Reserve covers 3.2 months; after that, cut fixed costs   |
| Equipment breaks                 |        -R$ 4,000 | "Laptop replacement" budget covers R$ 900; rest financed |

**Emergency withdrawal order:** ① Cash → ② Reserve → ③ Liquid investments → ④ Credit (last resort).

---

## 🔗 Related
[`Month-Report.md`](Month-Report.md) · [`Budget.md`](Budget.md) · [`Net-Worth.md`](Net-Worth.md) · [`Business-MEI.md`](Business-MEI.md)
