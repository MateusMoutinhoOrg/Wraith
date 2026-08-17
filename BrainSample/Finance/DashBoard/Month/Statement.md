# Consolidated Statement — August 2026

> All transactions, all accounts · generated on day 16 · 27 transactions · **Result so far: -R$ 200**
> Per-account statements: [`Accounts/Bank.md`](Accounts/Bank.md) · [`Accounts/Cash.md`](Accounts/Cash.md) · [`Accounts/Credit-Card.md`](Accounts/Credit-Card.md)

---

## 1. Summary by account

| Account                | Opening      | Transactions | Change      | Current      | Status |
| ---------------------- | -----------: | -----------: | ----------: | -----------: | :----: |
| Cash                   |       R$ 300 |           14 |    -R$ 163  |      R$ 137  |   🔴   |
| Bank                   |     R$ 1,000 |            9 |     -R$ 37  |      R$ 963  |   🟡   |
| Credit card (owed)     |      -R$ 380 |            6 |       R$ 0  |     -R$ 380  |   🟡   |
| **Available**          | **R$ 1,300** |            — | **-R$ 200** | **R$ 1,100** |   🟡   |
| **Net (minus card)**   |   **R$ 920** |            — | **-R$ 200** |   **R$ 720** |   🔴   |

---

## 2. Ledger (01-aug → 16-aug)

| Date   | Description               | Category       | Account     | Value       | Available    |
| ------ | ------------------------- | -------------- | ----------- | ----------: | -----------: |
| 01-aug | Rent + utilities          | `Home`         | Card        |     -R$ 300 |     R$ 1,300 |
| 02-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |     R$ 1,290 |
| 02-aug | Supermarket               | `Food`         | Bank        |      -R$ 45 |     R$ 1,245 |
| 03-aug | Fuel                      | `Transport`    | Card        |      -R$ 15 |     R$ 1,245 |
| 04-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |     R$ 1,235 |
| 04-aug | Bakery                    | `Food`         | Cash        |      -R$ 18 |     R$ 1,217 |
| 05-aug | Cloud hosting             | `Business`     | Card        |      -R$ 20 |     R$ 1,217 |
| 05-aug | *July card bill payment*  | `Card-Payment` | Bank → Card |   *-R$ 380* |       R$ 837 |
| 06-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |       R$ 827 |
| 06-aug | Poker — buy-in            | `Poker`        | Cash        |      -R$ 50 |       R$ 777 |
| 07-aug | Supermarket               | `Food`         | Bank        |      -R$ 32 |       R$ 745 |
| 08-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |       R$ 735 |
| 08-aug | Ride app                  | `Transport`    | Bank        |      -R$ 12 |       R$ 723 |
| 09-aug | Poker — cash out          | `Poker`        | Cash        |      +R$ 30 |       R$ 753 |
| 10-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |       R$ 743 |
| 10-aug | Lunch out                 | `Food`         | Card        |      -R$ 25 |       R$ 743 |
| 11-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |       R$ 733 |
| 11-aug | **Client payment**        | `Freelance`    | Bank        | **+R$ 500** |     R$ 1,233 |
| 12-aug | Supermarket               | `Food`         | Bank        |      -R$ 28 |     R$ 1,205 |
| 12-aug | Ride app                  | `Transport`    | Bank        |      -R$ 10 |     R$ 1,195 |
| 13-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |     R$ 1,185 |
| 13-aug | Domain renewal            | `Business`     | Card        |      -R$ 20 |     R$ 1,185 |
| 14-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |     R$ 1,175 |
| 14-aug | Poker — buy-in            | `Poker`        | Cash        |      -R$ 25 |     R$ 1,150 |
| 15-aug | Supermarket               | `Food`         | Bank        |      -R$ 22 |     R$ 1,128 |
| 15-aug | Fuel                      | `Transport`    | Bank        |       -R$ 8 |     R$ 1,120 |
| 16-aug | Cigarettes                | `Vices`        | Cash        |      -R$ 10 |     R$ 1,110 |
| 16-aug | Snack                     | `Food`         | Cash        |      -R$ 10 |     R$ 1,100 |

> "Available" = cash + bank. Card purchases do not change it — they increase the bill and leave the
> bank only when the bill is paid.

---

## 3. Totals by category

| Category    | Cash        | Bank        | Card        | Total       | Count | Share of spending |
| ----------- | ----------: | ----------: | ----------: | ----------: | ----: | ----------------: |
| `Freelance` |        R$ 0 |     +R$ 500 |        R$ 0 |     +R$ 500 |     1 |        — (income) |
| `Home`      |        R$ 0 |        R$ 0 |     -R$ 300 |     -R$ 300 |     1 |             42.9% |
| `Food`      |      -R$ 28 |     -R$ 127 |      -R$ 25 |     -R$ 180 |     7 |             25.7% |
| `Vices`     |      -R$ 90 |        R$ 0 |        R$ 0 |      -R$ 90 |     9 |             12.9% |
| `Poker`     |      -R$ 45 |        R$ 0 |        R$ 0 |      -R$ 45 |     3 |             10.7% |
| `Transport` |        R$ 0 |      -R$ 30 |      -R$ 15 |      -R$ 45 |     4 |              6.4% |
| `Business`  |        R$ 0 |        R$ 0 |      -R$ 40 |      -R$ 40 |     2 |              5.7% |
| **Total**   | **-R$ 163** | **+R$ 343** | **-R$ 380** | **-R$ 200** | **27**|          **100%** |

```
Spending by category
Home       ██████████████████████  42.9%
Food       █████████████░░░░░░░░░  25.7%
Vices      ██████▌░░░░░░░░░░░░░░░  12.9%
Poker      █████▌░░░░░░░░░░░░░░░░  10.7%
Transport  ███▎░░░░░░░░░░░░░░░░░░   6.4%
Business   ██▉░░░░░░░░░░░░░░░░░░░   5.7%
```

*The bank column excludes the R$ 380 card payment (transfer); including it, the bank moved -R$ 37.*

---

## 4. Transfers between own accounts

| Date        | From | To      | Value    | Note                                            |
| ----------- | ---- | ------- | -------: | ----------------------------------------------- |
| 05-aug-2026 | Bank | Card    |   R$ 380 | July bill, paid in full — no interest           |
| 30-aug-2026 | Bank | Savings | R$ 500*  | Monthly reserve ([`../Net-Worth.md`](../Net-Worth.md)) |
| 05-sep-2026 | Bank | Card    | R$ 380*  | August bill                                     |

> Transfers are not income or expense — counting the card payment as an expense would charge July's
> purchases twice.

---

## 5. Projection (17-aug → 31-aug)

| Date      | Description               | Account | Value          | Confidence |
| --------- | ------------------------- | ------- | -------------: | ---------- |
| 25-aug    | Client A invoice          | Bank    |    +R$ 1,500*  | 🟡 likely  |
| 17→31-aug | Groceries at current pace | Bank    |      -R$ 100*  | 🟢 high    |
| 17→31-aug | Snacks at current pace    | Cash    |       -R$ 50*  | 🟢 high    |
| 17→31-aug | Cigarettes at current pace| Cash    |       -R$ 70*  | 🔴 high    |
| 17→31-aug | Transport at current pace | Bank    |       -R$ 40*  | 🟢 high    |
| 20-aug    | Cloud + DAS tax           | Bank    |      -R$ 136*  | 🟢 certain |
| **Projected month result** |          |         | **+R$ 1,000*** |            |

| Account            | Current      | Projected 31-aug |
| ------------------ | -----------: | ---------------: |
| Cash               |       R$ 137 |           R$ 17* |
| Bank               |       R$ 963 |        R$ 1,803* |
| Credit card (owed) |      -R$ 380 |         -R$ 400* |
| **Available**      | **R$ 1,100** |      **R$ 1,820*** |

---

## 6. Observations

| Status | Observation                                                            |
| :----: | ---------------------------------------------------------------------- |
|   🔴   | `Vices` purchases on 9 of 16 days — R$ 10 each, ~R$ 3,650/year pace    |
|   🔴   | Single income event in the month — full dependency on one client       |
|   🟡   | `Food` split across 7 small purchases in 3 accounts — no shopping list |
|   🟡   | Over half of the month's spending is on the card, due 05-sep           |
|   🟢   | Discretionary spending is cash-only, self-limiting                     |
|   🟢   | No overdraft, card paid in full — zero fees and interest               |

> *Fictional document, used as a model for personal + small-business financial management.*
