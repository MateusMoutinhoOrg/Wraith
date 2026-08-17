# 📐 KPIs — Dictionary and Historical Series

> **Reference:** aug/2026 (partial, day 16) · Every indicator has a formula, a target and history.

---

## 1. 🩺 Financial Health Score

**48 / 100** — `████▊░░░░░` 🟡 *Fragile, but recoverable*

| Pillar             | Weight | Score | Contribution | Read                                    |
| ------------------ | -----: | ----: | -----------: | --------------------------------------- |
| 💧 Liquidity       |    25% |    30 |         7.5  | Reserve covers only 3.2 months          |
| 📦 Budget          |    20% |    62 |        12.4  | 5 of 7 envelopes within cap             |
| 🐖 Savings         |    25% |    20 |         5.0  | Negative rate in the partial month      |
| 🌐 Diversification |    15% |    25 |         3.8  | 100% of income from 1 client            |
| 🎯 Discipline      |    15% |    70 |        10.5  | Entries up to date, few overruns        |
| **Total**          | 100%   |    —  |   **48.2**   | 🟡                                       |

**Bands:** 0–39 🔴 critical · 40–59 🟡 fragile · 60–79 🟢 healthy · 80–100 💎 solid

```
Score  jan  feb  mar  apr  may  jun  jul  aug
       52   55   58   56   61   66   44   48
       ██▌  ██▊  ██▉  ██▊  ███  ███▎ ██▏  ██▍
```

---

## 2. 📖 Indicator Dictionary

### 💧 Liquidity and Safety

| KPI                    | Formula                                | Current | Target  | Status | Trend |
| ---------------------- | -------------------------------------- | ------: | ------: | ------ | ----- |
| **Reserve Runway**     | Reserve ÷ average monthly expense      | 3.2 m   | ≥ 6 m   | 🟡     | ↗     |
| **Immediate Liquidity**| Cash ÷ expenses of the next 30 days    | 1.10x   | ≥ 1.5x  | 🟡     | ↘     |
| **Cash Trough**        | Lowest projected balance in the month  | -R$ 380 | ≥ R$ 0  | 🔴     | ↘     |
| **Days of Breathing Room** | Cash ÷ daily burn rate             | 25 d    | ≥ 45 d  | 🟡     | →     |

### 🐖 Accumulation

| KPI                     | Formula                              | Current| Target | Status | Trend |
| ----------------------- | ------------------------------------ | -----: | -----: | ------ | ----- |
| **Savings Rate**        | (Income − Expense) ÷ Income          | -40.0% | ≥ 30%  | 🔴     | ↘     |
| **Coverage Ratio**      | Income ÷ Expense                     |  0.71x | ≥ 1.5x | 🔴     | ↘     |
| **Δ Net Worth (month)** | (NW_end − NW_start) ÷ NW_start       |    -2% | ≥ +2%  | 🔴     | ↘     |
| **Effective Contribution** | Amount invested ÷ contribution target |  0% | 100%   | 🔴     | →     |

### 📦 Spending Control

| KPI                       | Formula                                | Current | Target  | Status | Trend |
| ------------------------- | -------------------------------------- | ------: | ------: | ------ | ----- |
| **Budget Adherence**      | Envelopes within cap ÷ total           |   5 / 7 |   7 / 7 | 🟡     | ↘     |
| **Fixed Cost / Income**   | Fixed ÷ projected income               |   29.1% |   ≤ 50% | 🟢     | →     |
| **Discretionary**         | (Vices + Poker) ÷ total expense        |   19.3% |   ≤ 15% | 🟡     | ↗     |
| **Average Ticket**        | Expense ÷ no. of transactions          | R$ 25.9 | ≤ R$ 30 | 🟢     | ↘     |
| **No-spend Days**         | Count of days with R$ 0 of outflow     |  4 / 16 |    ≥ 8  | 🟡     | ↘     |
| **Pace vs. Calendar**     | % budget used − % of month elapsed     | +18 pp  | ≤ 0 pp  | 🔴     | ↗     |

### 💼 Business (MEI)

| KPI                        | Formula                             | Current | Target  | Status | Trend |
| -------------------------- | ----------------------------------- | ------: | ------: | ------ | ----- |
| **Client Concentration**   | Largest client ÷ revenue            |    100% |   ≤ 60% | 🔴     | ↗     |
| **Operating Margin**       | (Revenue − costs) ÷ revenue         |   85.0% |   ≥ 70% | 🟢     | →     |
| **Income Volatility**      | Standard deviation ÷ mean (12m)     |    ±38% |   ≤ 20% | 🔴     | ↗     |
| **MEI Cap Usage**          | 12-month revenue ÷ R$ 81,000        |   16.7% |  ≤ 100% | 🟢     | ↗     |

### 💳 Indebtedness

| KPI                       | Formula                          | Current| Target | Status | Trend |
| ------------------------- | -------------------------------- | -----: | -----: | ------ | ----- |
| **Debt / Net Worth**      | Liabilities ÷ assets             | 21.6%  | ≤ 30%  | 🟢     | ↘     |
| **Income Commitment**     | Installments ÷ monthly income    |  9.0%  | ≤ 30%  | 🟢     | ↘     |
| **Expensive Debt**        | Balance with interest > 3%/mo    | R$ 380 |  R$ 0  | 🟡     | ↘     |

---

## 3. 📈 2026 Historical Series

| KPI                    | jan   | feb   | mar   | apr   | may   | jun   | jul   | aug*  | Target |
| ---------------------- | ----: | ----: | ----: | ----: | ----: | ----: | ----: | ----: | -----: |
| Income (R$)            | 1,500 | 1,700 | 1,800 | 1,500 | 1,800 | 2,200 | 1,000 | 2,000 | 2,000  |
| Expense (R$)           |   900 |   880 | 1,000 |   920 |   950 | 1,100 | 1,050 | 1,000 | 1,000  |
| Result (R$)            |  +600 |  +820 |  +800 |  +580 |  +850 |+1,100 |   -50 |+1,000 | +1,000 |
| Savings rate           |  40%  |  48%  |  44%  |  39%  |  47%  |  50%  |  -5%  |  50%  |  ≥30%  |
| Coverage (x)           | 1.67  | 1.93  | 1.80  | 1.63  | 1.89  | 2.00  | 0.95  | 2.00  |  ≥1.5  |
| Runway (months)        |  1.2  |  1.8  |  2.3  |  2.6  |  2.9  |  3.3  |  3.1  |  3.2  |   ≥6   |
| Adherence (envelopes)  |  6/7  |  7/7  |  6/7  |  7/7  |  6/7  |  5/7  |  4/7  |  5/7  |   7/7  |
| Health score           |   52  |   55  |   58  |   56  |   61  |   66  |   44  |   48  |  ≥60   |

`*` projected

```
Savings Rate 2026 (target 30% = ─────)
50% ┤                    ██              ██
40% ┤ ██  ██  ██  ██  ██  ██              ██
30% ┼─██──██──██──██──██──██──────────────██──
20% ┤ ██  ██  ██  ██  ██  ██              ██
10% ┤ ██  ██  ██  ██  ██  ██              ██
 0% ┼─██──██──██──██──██──██───▁▁─────────██──
-10%┤                          ██
    └ jan feb mar apr may jun jul aug
```

---

## 4. 🔍 Month Diagnosis

| Type          | Observation                                                                   |
| ------------- | ----------------------------------------------------------------------------- |
| 🔴 Problem    | **July broke the streak.** Income fell 55% and expenses didn't follow.        |
| 🔴 Problem    | Client concentration rose from 60% → 100%: existential risk to income.        |
| 🟡 Risk       | Runway stuck at ~3.2 months for 3 months — the contributions aren't happening.|
| 🟢 Positive   | Fixed costs controlled at 29% of income, well below the 50% cap.              |
| 🟢 Positive   | Average ticket falling — large impulse purchases have disappeared.            |
| 💡 Lever      | Halving `Vices` raises the savings rate by ~2.3 pp/month.                     |
| 💡 Lever      | A 2nd client at R$ 1,000/month drops volatility from ±38% to ~±18%.          |

---

## 5. 🎚️ KPI Targets for dec/2026

| KPI                | Today | Dec target | Gap      | How to close it                        |
| ------------------ | ----: | ---------: | -------: | -------------------------------------- |
| Health score       |    48 |         65 | +17 pts  | Runway + diversification               |
| Runway             | 3.2 m |      6.0 m | +2.8 m   | R$ 700/month contributed for 4 months  |
| Savings rate       |  -40% |       +35% | +75 pp   | Close aug in the black and keep it there|
| Concentration      |  100% |        60% | -40 pp   | Close 2 clients by oct                 |
| Adherence          |   5/7 |        7/7 | +2       | Freeze `Vices`, reallocate `Study`     |
| Expensive debt     | R$ 380|       R$ 0 | -R$ 380  | Pay off the card bill in sep           |

---

## 🔗 Related
[`Month-Report.md`](Month-Report.md) · [`Year-Report.md`](Year-Report.md) · [`Net-Worth.md`](Net-Worth.md) · [`Business-MEI.md`](Business-MEI.md)
