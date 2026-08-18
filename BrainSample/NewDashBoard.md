
## New DashBoard
Problem: the current Dashboard is irrealistic,its not covering the possible actions of the user (Tasks).


## New Dasboard expected:
- Aligned with the user tasks
 - (kpis,docs,everything) must be possible with the user given information.
- The user must be able to navigate through the documentation structure by the dashboard
- the documentation must cover the Categories , like showing the current categories,etc..
- the documentation must cover Accounts, like showing the current accounts,etc..
- the documentation must cover Creddit Cards 
- it must be possible to navegate over transactions, from current month, and from prewies month. 
- the documentation must cover the **future**: how much each account will hold in each of the
  next 8 months.
  - the same rule holds — it can only project what the user declared. That needs a new task
    ([`AddRecurrence`](Tasks/AddRecurrence.md)) for monthly commitments and an `installments`
    field on [`AddTransaction`](Tasks/AddTransaction.md); a card bill is derived from the card
    registry, never declared.
  - a projection must never silently mix declared facts with statistical guesses. Whatever the
    user did **not** declare has to be shown as a gap, in its own section, labelled as an
    estimate — a forecast that quietly ignores R$ 850/month of groceries is worse than no
    forecast.
  - rendered at [`DashBoard/Forecast.md`](DashBoard/Forecast.md).
