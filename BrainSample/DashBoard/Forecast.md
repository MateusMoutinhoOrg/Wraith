# Forecast

> **Updated:** 18-aug-2026 · Horizon **sep-2026 → apr-2027** (8 months) · 6 recurrences · 0 installments in flight

[Dashboard](README.md) · [Credit Cards](Credit-Cards.md) · [Categories](Categories.md) · [Months](Months/README.md) · [Forecast](Forecast.md)

Everything below is today's balance rolled forward through commitments **you declared** — the
recurrences in §3, the installments and future-dated transactions in §4, and the card bills §5
derives from the card registry. Nothing here is a statistical guess. What that leaves out is
spelled out in §6, and it is not small.

---

## 1. What you will have, month by month

Balance at the **end** of each month.

| Month | `Bank` | `Cash` | `Emergency Savings` | Held in accounts | Owed on cards | **Net position** |
| ----- | -----: | -----: | ------------------: | ---------------: | ------------: | ---------------: |
| aug-2026 *(today)* | R$ 2,119 | R$ 75 | R$ 2,500 | R$ 4,694 | R$ 1,185 | **R$ 3,509** |
| sep-2026 | R$ 3,114 | R$ 75 | R$ 3,000 | R$ 6,189 | R$ 1,075 | **R$ 5,114** |
| oct-2026 | R$ 3,474 | R$ 75 | R$ 3,500 | R$ 7,049 | R$ 1,130 | **R$ 5,919** |
| nov-2026 | R$ 3,779 | R$ 75 | R$ 4,000 | R$ 7,854 | R$ 1,130 | **R$ 6,724** |
| dec-2026 | R$ 4,084 | R$ 75 | R$ 4,500 | R$ 8,659 | R$ 1,130 | **R$ 7,529** |
| jan-2027 | R$ 4,389 | R$ 75 | R$ 5,000 | R$ 9,464 | R$ 1,130 | **R$ 8,334** |
| feb-2027 | R$ 4,694 | R$ 75 | R$ 5,500 | R$ 10,269 | R$ 1,130 | **R$ 9,139** |
| mar-2027 | R$ 5,444 | R$ 75 | R$ 5,500 | R$ 11,019 | R$ 1,075 | **R$ 9,944** |
| apr-2027 | R$ 6,304 | R$ 75 | R$ 5,500 | R$ 11,879 | R$ 1,130 | **R$ 10,749** |

```
Held in accounts
aug-2026  █████████░░░░░░░░░░░░░░░  R$ 4,694  ← today
sep-2026  ████████████░░░░░░░░░░░░  R$ 6,189
oct-2026  ██████████████░░░░░░░░░░  R$ 7,049
nov-2026  ████████████████░░░░░░░░  R$ 7,854
dec-2026  █████████████████░░░░░░░  R$ 8,659
jan-2027  ███████████████████░░░░░  R$ 9,464
feb-2027  █████████████████████░░░  R$ 10,269
mar-2027  ██████████████████████░░  R$ 11,019
apr-2027  ████████████████████████  R$ 11,879
```

`Emergency Savings` stops growing after feb-2027 because the `Emergency reserve` recurrence
carries `end: 2027-02`. `Cash` never moves: no recurrence is declared against it.

---

## 2. Where `Bank` moves each month

`Bank` is the only account with both sides declared, so it is the one worth breaking out.

| Month | Opening | Income | Card bill paid | To savings | Other expenses | Closing |
| ----- | ------: | -----: | -------------: | ---------: | -------------: | ------: |
| sep-2026 | R$ 2,119 | +R$ 2,800 | -R$ 1,185 | -R$ 500 | -R$ 120 | **R$ 3,114** |
| oct-2026 | R$ 3,114 | +R$ 2,000 | -R$ 1,020 | -R$ 500 | -R$ 120 | **R$ 3,474** |
| nov-2026 | R$ 3,474 | +R$ 2,000 | -R$ 1,075 | -R$ 500 | -R$ 120 | **R$ 3,779** |
| dec-2026 | R$ 3,779 | +R$ 2,000 | -R$ 1,075 | -R$ 500 | -R$ 120 | **R$ 4,084** |
| jan-2027 | R$ 4,084 | +R$ 2,000 | -R$ 1,075 | -R$ 500 | -R$ 120 | **R$ 4,389** |
| feb-2027 | R$ 4,389 | +R$ 2,000 | -R$ 1,075 | -R$ 500 | -R$ 120 | **R$ 4,694** |
| mar-2027 | R$ 4,694 | +R$ 2,000 | -R$ 1,130 | R$ 0 | -R$ 120 | **R$ 5,444** |
| apr-2027 | R$ 5,444 | +R$ 2,000 | -R$ 1,020 | R$ 0 | -R$ 120 | **R$ 6,304** |

September's income is R$ 2,800 because the R$ 800 `Client B invoice` — recorded on 18-aug with
`payment_date: 2026-09-02` — settles that month on top of the retainer. It is a fact, not a
projection; see §4.

---

## 3. Declared recurrences

| Recurrence | Account | Category | Amount | Day | From | Until | Occurrences left |
| ---------- | ------- | -------- | -----: | --: | ---- | ----- | ---------------: |
| `Client A retainer` | [`Bank`](Accounts/Bank.md) | [`Freelance`](Categories.md) | +R$ 2,000 | 11 | sep-2026 | — | open-ended |
| `Rent` | [`Nubank Card`](Credit-Cards.md) | [`Home`](Categories.md) | -R$ 900 | 1 | sep-2026 | — | open-ended |
| `Internet` | [`Bank`](Accounts/Bank.md) | [`Home`](Categories.md) | -R$ 120 | 17 | sep-2026 | — | open-ended |
| `Cloud hosting` | [`Nubank Card`](Credit-Cards.md) | [`Business`](Categories.md) | -R$ 120 | 5 | sep-2026 | — | open-ended |
| `Streaming` | [`Nubank Card`](Credit-Cards.md) | [`Leisure`](Categories.md) | -R$ 55 | 30 | sep-2026 | — | open-ended |
| `Emergency reserve` | `Bank` → `Emergency Savings` | [`Reserve`](Categories.md) | R$ 500 | 6 | sep-2026 | feb-2027 | 6 |

Monthly total: **+R$ 2,000** in, **-R$ 1,195** out, **R$ 500** moved between your own accounts.

`Emergency reserve` is a transfer: one recurrence with `to_account` filled in, expanded into two
legs that cancel out — it is neither income nor an expense, it just changes which account holds
the money. `Streaming` falls on day 30, which is clamped to the 28th in february 2027.

---

## 4. Already scheduled

Facts the registry already holds with a date in the future. These are not projections — they are
transactions that exist.

| Date | What | Account | Amount | Source |
| ---- | ---- | ------- | -----: | ------ |
| 02-sep-2026 | `Client B invoice` settles | [`Bank`](Months/2026-08/Accounts/Bank.md) | +R$ 800 | `AddTransaction.payment_date` |

**Installments in flight:** none. A purchase recorded with
[`installments`](../Tasks/AddTransaction.md) writes one transaction per month up front, so every
remaining part would be listed here and would already be counted in §1 and §5.

---

## 5. Card bills

A bill is never a recurrence — its amount changes every month. It is derived: recurring purchases
and installment parts falling inside the billing window add up to the bill that closes on
`closing_day`, and the whole of it leaves the paying account on `due_day`.

`Nubank Card` — closes day 28, due day 5. The paying account is `Bank`, the account that settled
the last two bills.

| Bill | Purchase window | Closes | What lands in it | Amount | Due | Paid from |
| ---- | --------------- | -----: | ---------------- | -----: | --: | --------- |
| aug-2026 | 29-jul → 28-aug | 28-aug | already purchased | R$ 1,185 | 05-sep | `Bank` |
| sep-2026 | 29-aug → 28-sep | 28-sep | Rent, Cloud hosting | R$ 1,020 | 05-oct | `Bank` |
| oct-2026 | 29-sep → 28-oct | 28-oct | Streaming (30-sep), Rent, Cloud hosting | R$ 1,075 | 05-nov | `Bank` |
| nov-2026 | 29-oct → 28-nov | 28-nov | Streaming, Rent, Cloud hosting | R$ 1,075 | 05-dec | `Bank` |
| dec-2026 | 29-nov → 28-dec | 28-dec | Streaming, Rent, Cloud hosting | R$ 1,075 | 05-jan | `Bank` |
| jan-2027 | 29-dec → 28-jan | 28-jan | Streaming, Rent, Cloud hosting | R$ 1,075 | 05-feb | `Bank` |
| feb-2027 | 29-jan → 28-feb | 28-feb | Streaming ×2, Rent, Cloud hosting | R$ 1,130 | 05-mar | `Bank` |
| mar-2027 | 01-mar → 28-mar | 28-mar | Rent, Cloud hosting | R$ 1,020 | 05-apr | `Bank` |
| apr-2027 | 29-mar → 28-apr | 28-apr | Streaming (30-mar), Rent, Cloud hosting | R$ 1,075 | 05-may | `Bank` |

Two windows are worth reading twice. February 2027 catches **two** `Streaming` charges — 30-jan
and 28-feb, the latter being day 30 clamped to the end of a 28-day month — which is why that bill
is R$ 1,130. March then opens on the 1st with nothing before it and comes out R$ 1,020.

The projection assumes every bill is paid in full on its due date. Available limit stays between
R$ 3,870 and R$ 3,980 of R$ 5,000 across the whole horizon — no month comes close to it.

---

## 6. What this projection does not include

**Read this before trusting §1.** The forecast only knows what was declared, and your day-to-day
spending is not declared. These are the categories that moved money in jul-2026 and aug-2026
with **no recurrence covering them**:

| Category | jul-2026 | aug-2026 | Monthly average | Declared? |
| -------- | -------: | -------: | --------------: | --------- |
| [`Food`](Categories.md) | -R$ 545 | -R$ 450 | **-R$ 498** | no |
| [`Vices`](Categories.md) | -R$ 200 | -R$ 180 | **-R$ 190** | no |
| [`Transport`](Categories.md) | -R$ 145 | -R$ 90 | **-R$ 118** | no |
| [`Business`](Categories.md) beyond `Cloud hosting` | -R$ 76 | -R$ 40 | **-R$ 58** | no |
| [`Study`](Categories.md) | R$ 0 | -R$ 60 | **-R$ 30** | no |
| [`Leisure`](Categories.md) own spending | -R$ 55 | R$ 0 | **-R$ 28** | no |
| [`Poker`](Categories.md) net | +R$ 80 | -R$ 100 | **-R$ 10** | no |
| [`Extra`](Categories.md) | R$ 0 | +R$ 150 | **+R$ 75** | no |
| | | | **≈ -R$ 850 / month** | |

Subtract that and the picture changes completely:

| Month | §1 says you hold | Minus undeclared spending | Reality check |
| ----- | ---------------: | ------------------------: | ------------: |
| sep-2026 | R$ 6,189 | -R$ 850 | **R$ 5,339** |
| oct-2026 | R$ 7,049 | -R$ 1,700 | **R$ 5,349** |
| nov-2026 | R$ 7,854 | -R$ 2,550 | **R$ 5,304** |
| dec-2026 | R$ 8,659 | -R$ 3,400 | **R$ 5,259** |
| jan-2027 | R$ 9,464 | -R$ 4,250 | **R$ 5,214** |
| feb-2027 | R$ 10,269 | -R$ 5,100 | **R$ 5,169** |
| mar-2027 | R$ 11,019 | -R$ 5,950 | **R$ 5,069** |
| apr-2027 | R$ 11,879 | -R$ 6,800 | **R$ 5,079** |

The right-hand column is the only number on this page that is an estimate — a two-month average
projected forward, shown so §1 is not mistaken for a promise. It says the flat truth: on today's
declared commitments plus today's habits, you do not get richer over the next 8 months. Money
moves from `Bank` into `Emergency Savings` and that is roughly all that happens.

Three things close the gap, in order of how much they buy you:

1. Declare the fixed part of what is missing — a grocery budget, a transport budget — with
   [`AddRecurrence`](../Tasks/AddRecurrence.md). Every recurrence you add moves a line out of
   this section and into §1.
2. Record purchases split over months with
   [`installments`](../Tasks/AddTransaction.md) so the parts land in the months they fall in
   instead of appearing as a surprise.
3. Give the recurrences you already have an honest `end` — a contract that finishes in january
   should say so, otherwise §1 quietly assumes it is forever.

---

## 7. What you can do here

| Want to | Task |
| ------- | ---- |
| Declare a repeating income or expense | [`AddRecurrence`](../Tasks/AddRecurrence.md) |
| Stop a repeating commitment | [`RemoveRecurrence`](../Tasks/RemoveRecurrence.md) |
| Record a purchase split over months | [`AddTransaction`](../Tasks/AddTransaction.md) with `installments` |
| Change when a commitment ends | [`AddRecurrence`](../Tasks/AddRecurrence.md) with a new `end` |

---

## 8. Conventions

| Notation  | Meaning                                                                                |
| --------- | -------------------------------------------------------------------------------------- |
| Horizon   | The 8 months after the open one — `DashBoard.future-months`. It moves with the calendar |
| Declared  | Comes from a recurrence, an installment part, or a `payment_date` in the future        |
| Derived   | Computed from the registry rather than declared — card bills are the only case         |
| Estimated | A historical average. Confined to §6 and never mixed into §1                           |
| Clamping  | A `day` larger than the month has falls on that month's last day, never the next month |

> Generated by `./wraith tick` as part of the `DashBoard` visualization declared in
> [`../Visualization.yaml`](../Visualization.yaml), from the tasks in
> [`../Help/Task.md`](../Help/Task.md) — do not edit by hand.
