# Dashboard

> **Updated:** 18-aug-2026 · **Currency:** BRL (R$) · **Registry:** 3 accounts, 1 credit card, 14 categories, 63 transactions, 6 recurrences

[Dashboard](README.md) · [Accounts](Accounts.md) · [Credit Cards](Credit-Cards.md) · [Categories](Categories.md) · [Months](Months/README.md) · [Forecast](Forecast.md)

---

## 1. Position on 18-aug-2026

| Indicator | Value | Where it comes from |
| --------- | ----: | ------------------- |
| Balance in accounts | **R$ 4,694** | opening balances + every transaction |
| Owed on credit cards | R$ 1,185 | card opening + purchases − payments |
| **Net position** | **R$ 3,509** | what you hold − what you owe |
| Card limit available | R$ 3,815 of R$ 5,000 | `AddCreditCard.limit` − outstanding |
| Pending settlement | +R$ 800 | transactions with a future `payment_date` |

| Account                                                           |  Balance | Share of the money you hold |
| ----------------------------------------------------------------- | -------: | --------------------------- |
| [Emergency Savings](Months/2026-08/Accounts/Emergency-Savings.md) | R$ 2,500 | `███████████░░░░░░░░░` 53%  |
| [Bank](Months/2026-08/Accounts/Bank.md)                           | R$ 2,119 | `█████████░░░░░░░░░░░` 45%  |
| [Cash](Months/2026-08/Accounts/Cash.md)                           |    R$ 75 | `█░░░░░░░░░░░░░░░░░░░` 2%   |

---

## 2. August 2026 so far

| Line | This month | Previous month | Change |
| ---- | ---------: | -------------: | -----: |
| Income | +R$ 2,200 | +R$ 2,680 | -R$ 480 |
| Expenses | -R$ 2,110 | -R$ 2,261 | +R$ 151 |
| **Result** | **+R$ 90** | +R$ 419 | -R$ 329 |
| Transactions | 32 | 31 | +1 |

Full month: [`Months/2026-08/DashBoard.md`](Months/2026-08/DashBoard.md) · ledger: [`Months/2026-08/Statement.md`](Months/2026-08/Statement.md)

---

## 3. The next 8 months

Today's balance rolled forward through the commitments you declared — 6 recurrences, the card
bills derived from them, and one invoice still to settle.

| Month | Held in accounts | Net position | Month | Held in accounts | Net position |
| ----- | ---------------: | -----------: | ----- | ---------------: | -----------: |
| sep-2026 | R$ 6,189 | R$ 5,114 | jan-2027 | R$ 9,464 | R$ 8,334 |
| oct-2026 | R$ 7,049 | R$ 5,919 | feb-2027 | R$ 10,269 | R$ 9,139 |
| nov-2026 | R$ 7,854 | R$ 6,724 | mar-2027 | R$ 11,019 | R$ 9,944 |
| dec-2026 | R$ 8,659 | R$ 7,529 | apr-2027 | R$ 11,879 | R$ 10,749 |

⚠️ Those figures count declared commitments only. About **R$ 850 a month** of real spending —
groceries, transport, the daily habits — has no recurrence covering it, and once you subtract it
the balance stays flat around R$ 5,200 instead of climbing. The breakdown is in
[`Forecast.md` §6](Forecast.md).

Full projection: [`Forecast.md`](Forecast.md)

---

## 4. Where to go next

| Document | Answers | Fed by |
| -------- | ------- | ------ |
| [`Accounts.md`](Accounts.md) | Which accounts exist and what each one holds | `AddAccount`, `RemoveAccount`, `AddTransaction` |
| [`Credit-Cards.md`](Credit-Cards.md) | Limit, outstanding, when each bill closes and is due | `AddCreditCard`, `RemoveCreditCard`, `AddTransaction` |
| [`Categories.md`](Categories.md) | How spending is classified and what each category costs | `AddCategory`, `RemoveCategory`, `AddTransaction` |
| [`Months/README.md`](Months/README.md) | Every closed and open month | `AddTransaction` |
| [`Forecast.md`](Forecast.md) | What each account holds in each of the next 8 months | `AddRecurrence`, `RemoveRecurrence`, `AddTransaction`, `AddCreditCard` |

```
DashBoard/
├── README.md              ← you are here
├── Accounts.md
├── Credit-Cards.md
├── Categories.md
├── Forecast.md            ← the next 8 months
└── Months/
    ├── README.md
    ├── 2026-08/   ← current month
    │   ├── DashBoard.md
    │   ├── Statement.md
    │   └── Accounts/
    │       ├── Bank.md
    │       ├── Cash.md
    │       ├── Emergency-Savings.md
    │       └── Nubank-Card.md
    └── 2026-07/
        ├── DashBoard.md
        ├── Statement.md
        └── Accounts/
            ├── Bank.md
            ├── Cash.md
            ├── Emergency-Savings.md
            └── Nubank-Card.md
```

---

## 5. Why these pages and not others

That tree is not built in. Every file above exists because it is declared as an entry in
[`../Visualization.yaml`](../Visualization.yaml) — a name, a destination, and optional arguments:

```yaml
- name: DashBoard
  dest: DashBoard/README.md

- name: ForeCast
  args:
    months: 8
  dest: DashBoard/Forecast.md
```

| To… | Do this in `Visualization.yaml` |
| --- | --- |
| Stop seeing a page | Delete its entry, or set `enabled: false` on it |
| Add a page back | Add an entry with its `name` and a `dest` |
| Move a page | Change its `dest` — the layout of the vault is yours |
| Change the forecast horizon | Change `ForeCast.args.months` (this vault: 8) |
| Preview a change first | `./wraith render ForeCast --months 12` |

Nothing here is rendered that you did not ask for, and a page you remove stops being refreshed on
the next tick. The catalog of every available visualization and its arguments is in
[`../Visualization.md`](../Visualization.md).

---

## 6. Conventions

| Notation | Meaning |
| -------- | ------- |
| Income | A positive `amount` in a category whose `revenues` is `true` |
| Expense | A negative `amount` in a category whose `expenses` is `true` |
| Transfer | Any amount in a category with `revenues: false` **and** `expenses: false` — moving money between your own accounts is neither income nor expense |
| Pending | `payment_date` is later than today: the transaction is recorded but has not moved money yet, so it is excluded from balances |
| Card purchase | Counts in the month of `date`; the cash leaves the bank when you record the bill payment |
| Recurrence | A declared monthly commitment. It moves no balance and writes no transaction — it only tells [`Forecast.md`](Forecast.md) what is coming |
| Installment | One purchase recorded as N monthly transactions. Every part is real from the moment it is recorded |

Every figure on these pages is computed from the five registries the tasks write: accounts, credit cards, categories, transactions and recurrences. Nothing here needs information you were never asked for.

> Generated by `./wraith tick` from the tasks in [`../Task-help.md`](../Task-help.md), laid out by
> [`../Visualization.yaml`](../Visualization.yaml) — do not edit by hand.
