# 🏷️ Categories

> Chart of accounts for the dashboard. Every transaction in `Month/All-Balance.md` **must** have
> exactly **1 category** and **1 account** — see §5 for the account registry.

---

## 1. 🌳 Chart of Accounts

```
FINANCES
├── 📥 INCOME
│   ├── Freelance ......... invoiced company/MEI services
│   ├── Salary ............ employment contract (inactive)
│   ├── Yield ............. interest, dividends, cashback
│   └── Extra ............. item sales, reimbursement
│
├── 📤 EXPENSES
│   ├── 🔒 Essentials
│   │   ├── Home .......... rent, utilities, internet
│   │   ├── Food .......... groceries, meals
│   │   └── Transport ..... fuel, ride apps, maintenance
│   ├── 💼 Business
│   │   └── Business ...... cloud, domains, DAS, tools
│   ├── 🎈 Discretionary
│   │   ├── Poker ......... leisure, betting, streaming
│   │   ├── Vices ......... cigarettes, drinks
│   │   └── Study ......... courses, books
│   └── 💳 Debts
│       └── Debt .......... interest and amortization
│
└── 🔄 TRANSFERS (do not affect the result)
    ├── Reserve ........... emergency fund contribution
    ├── Invest ............ portfolio contribution
    └── Card-Payment ...... settling the credit card bill
```

> 🔄 A transfer moves money **between two accounts**. It has two legs that net to zero, it never
> hits an expense envelope, and it never changes the month result (rule R-21 in
> [`Alerts-Rules.md`](Alerts-Rules.md)).

---

## 2. 📇 Category Catalog

| Category     | Emoji | Group          | `positive` | `negative` | Cap/month | Active | Created on  |
| ------------ | ----- | -------------- | ---------- | ---------- | --------: | ------ | ----------- |
| `Freelance`  | 💼    | Income         | ✅ true    | ❌ false   |         — | ✅     | 03-jan-2026 |
| `Yield`      | 🪙    | Income         | ✅ true    | ❌ false   |         — | ✅     | 03-jan-2026 |
| `Extra`      | 🎁    | Income         | ✅ true    | ❌ false   |         — | ✅     | 11-feb-2026 |
| `Home`       | 🏠    | Essential      | ❌ false   | ✅ true    |    R$ 300 | ✅     | 03-jan-2026 |
| `Food`       | 🍽️    | Essential      | ❌ false   | ✅ true    |    R$ 250 | ✅     | 03-jan-2026 |
| `Transport`  | 🚗    | Essential      | ❌ false   | ✅ true    |    R$ 120 | ✅     | 03-jan-2026 |
| `Business`   | 💻    | Business       | ❌ false   | ✅ true    |    R$ 150 | ✅     | 05-jan-2026 |
| `Poker`      | 🃏    | Discretionary  | ✅ true    | ✅ true    |     R$ 80 | ✅     | 18-jan-2026 |
| `Vices`      | 🚬    | Discretionary  | ❌ false   | ✅ true    |     R$ 50 | ✅     | 22-jan-2026 |
| `Study`      | 📚    | Discretionary  | ❌ false   | ✅ true    |     R$ 50 | ✅     | 02-feb-2026 |
| `Debt`       | 💳    | Debt           | ❌ false   | ✅ true    |    R$ 120 | ✅     | 14-mar-2026 |
| `Reserve`    | 🛡️    | Transfer       | ✅ true    | ✅ true    |         — | ✅     | 03-jan-2026 |
| `Invest`     | 📈    | Transfer       | ✅ true    | ✅ true    |         — | ✅     | 03-jan-2026 |
| `Card-Payment`| 🔄   | Transfer       | ✅ true    | ✅ true    |         — | ✅     | 16-aug-2026 |
| `Salary`     | 🧾    | Income         | ✅ true    | ❌ false   |         — | ⚪ no  | 03-jan-2026 |

**Bidirectional categories** (`positive` **and** `negative` = true): `Poker`, `Reserve`, `Invest`, `Card-Payment`.
They are the only ones that accept both inflows and outflows in the same grouping.

---

## 3. 📊 Performance by Category (aug/2026, day 16)

| Category     | Mov. | Inflows  | Outflows | Net      | % of expense | Main account   | vs. jul  |
| ------------ | ---: | -------: | -------: | -------: | -----------: | -------------- | -------- |
| `Freelance`  |    1 |   R$ 500 |     R$ 0 | +R$ 500  |            — | 🏦 Bank 100%    | ↘ -50%   |
| `Home`       |    1 |     R$ 0 |   R$ 300 | -R$ 300  |        42.9% | 💳 Card 100%    | → 0%     |
| `Food`       |    7 |     R$ 0 |   R$ 180 | -R$ 180  |        25.7% | 🏦 Bank 71%     | ↗ +12%   |
| `Vices`      |    9 |     R$ 0 |    R$ 90 |  -R$ 90  |        12.9% | 💵 Cash 100%    | ↗ +80% 🔴|
| `Poker`      |    3 |    R$ 30 |    R$ 75 |  -R$ 45  |        10.7% | 💵 Cash 100%    | ↘ -10%   |
| `Transport`  |    4 |     R$ 0 |    R$ 45 |  -R$ 45  |         6.4% | 🏦 Bank 67%     | ↘ -20%   |
| `Business`   |    2 |     R$ 0 |    R$ 40 |  -R$ 40  |         5.7% | 💳 Card 100%    | → 0%     |
| `Study`      |    0 |     R$ 0 |     R$ 0 |    R$ 0  |         0.0% | —              | ⚪ idle  |
| `Debt`       |    0 |     R$ 0 |     R$ 0 |    R$ 0  |         0.0% | —              | ⚪       |
| `Card-Payment`|   2 |   R$ 380 |   R$ 380 |    R$ 0  |    *excluded*| 🏦 Bank → 💳     | → 0%     |

```
Expense distribution
Home        ██████████████████████  42.9%
Food        █████████████░░░░░░░░░  25.7%
Vices       ██████▌░░░░░░░░░░░░░░░  12.9%
Poker       █████▌░░░░░░░░░░░░░░░░  10.7%
Transport   ███▎░░░░░░░░░░░░░░░░░░   6.4%
Business    ██▉░░░░░░░░░░░░░░░░░░░   5.7%
```

---

## 4. 🔥 Frequency Ranking (no. of transactions)

| Pos | Category    | Mov. | Avg. ticket | Read                                        |
| --: | ----------- | ---: | ----------: | ------------------------------------------- |
|  1st| `Vices`     |    9 |     R$ 10.0 | 🔴 Daily micro-spending — biggest leak      |
|  2nd| `Food`      |    7 |     R$ 25.7 | 🟡 Fragmented purchases, no shopping list   |
|  3rd| `Transport` |    4 |     R$ 11.3 | 🟢 Healthy                                  |
|  4th| `Poker`     |    3 |     R$ 25.0 | 🟢 Within the envelope                      |
|  5th| `Business`  |    2 |     R$ 20.0 | 🟢 Recurring                                |
|  6th| `Home`      |    1 |    R$ 300.0 | 🟢 Single fixed cost                        |

> 💡 **Insight:** `Vices` costs R$ 10/day. That's **R$ 3,650/year** at the current pace — more than a year of rent.

---

## 5. 🏦 Account Registry

Every transaction is booked to exactly one account. Accounts answer *where the money is*;
categories answer *what it was for*. The two dimensions are independent.

| Account        | Emoji | Type      | Statement                                                    | Opening 01-aug | Day 16   | Purpose                            |
| -------------- | ----- | --------- | ------------------------------------------------------------ | -------------: | -------: | ---------------------------------- |
| `Cash`         | 💵    | Asset     | [`Month/Cash-Balance.md`](Month/Cash-Balance.md)               |       R$ 300   |  R$ 137  | Discretionary spending only        |
| `Bank`         | 🏦    | Asset     | [`Month/Bank-Balance.md`](Month/Bank-Balance.md)               |     R$ 1,000   |  R$ 963  | Income, essentials, transfers      |
| `Credit-Card`  | 💳    | Liability | [`Month/Credit-Card-Balance.md`](Month/Credit-Card-Balance.md) |      -R$ 380   | -R$ 380  | Fixed + recurring, paid in full    |
| **Consolidated**| 🧾   | —         | [`Month/All-Balance.md`](Month/All-Balance.md)                 |     **R$ 920** | **R$ 720** | Net position across all accounts |

### Account rules

1. **One account per transaction**, alongside its single category (rule R-18).
2. **A card purchase is an expense on the purchase date**, not on the payment date. It raises the
   bill and leaves the liquid balance untouched until the bill is paid.
3. **Paying the card is a `Card-Payment` transfer**, never an expense — booking it as one would
   double-count the original purchases.
4. **Envelopes are account-agnostic.** `Food` consumes the same cap whether paid in cash, by bank
   or on the card.
5. **The wallet is not refilled mid-month.** Running dry is the intended constraint on `Vices` and
   `Poker`, not an incident.
6. **Liquid total** = `Cash` + `Bank`. **Net position** = liquid − the card bill.

---

## 6. 📐 Categorization Rules

1. **One category per transaction.** If two fit, the more specific one wins (`Business` > `Home`).
2. **A duplicate category is an error.** Creating an existing category returns `Error: Category <X> already exists`.
3. **Transfers don't affect the result** — `Reserve`, `Invest` and `Card-Payment` move money between accounts, they don't consume it.
4. **Every new category** requires a cap defined before the first entry.
5. **A category with no movement for 3 months** is deactivated (`Active: ⚪ no`), not deleted — this preserves history.
6. **Renaming a category is forbidden.** Deactivate the old one and create a new one, so closed months aren't corrupted.
7. **Reimbursements** are entered as a positive value on the **same category** as the original expense, never as income.

---

## 7. 🧾 Change Log

| Date        | Operation      | Category   | Result                                    |
| ----------- | -------------- | ---------- | ----------------------------------------- |
| 14-mar-2026 | AddCategory    | `Debt`     | ✅ OK                                     |
| 02-feb-2026 | AddCategory    | `Study`    | ✅ OK                                     |
| 11-feb-2026 | AddCategory    | `Extra`    | ✅ OK                                     |
| 09-aug-2026 | AddCategory    | `Poker`    | ❌ `Error: Category Poker already exists` |
| 12-aug-2026 | RemoveCategory | `Salary`   | ⚪ Deactivated (unused since jan)         |
| 16-aug-2026 | AddCategory    | `Card-Payment` | ✅ OK — transfer group, bidirectional |
| 16-aug-2026 | AddAccount     | `Cash` · `Bank` · `Credit-Card` | ✅ OK — accounts split out of the single statement |

---

## 🔗 Related
[`Budget.md`](Budget.md) · [`Month/All-Balance.md`](Month/All-Balance.md) · [`Alerts-Rules.md`](Alerts-Rules.md)
