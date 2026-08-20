# Track Transactions

## Description
Records a real month in the financial brain shipped with this repository: accounts, categories, what came in, what went out, and the commitments the forecast reads. Creating the vault first is [StartABrain.md](/docs/Tutorials/StartABrain.md); the mechanics of running any task at all are [RunTasks.md](/docs/Tutorials/RunTasks.md).

### Rules
- A positive `amount` needs a category with `revenues: true`; a negative one needs `expenses: true`.
- A category with **both** false is a **transfer category** — it counts as neither income nor expense, and is how money moving between your own accounts is recorded.
- Every expense is paid in full on the day it happens. A movement carries one `date`, and that date answers both questions the pages ask: the month it counts in, and the day the money moves. There is nothing owed and nothing to settle later.
- A movement may still be dated ahead of today. It counts in its own month and reaches the balance on its date, which is what the forecast projects it into.

---

## Workflow

1. Create the places money sits. An account is born empty — what it already holds goes in as a transaction in step 3.

```bash
wraith run AddAccount --account Bank
wraith run AddAccount --account Cash
```

2. Create the categories transactions are classified under.

```bash
wraith run AddCategory --category Salary --description "What comes in" \
  --revenues true --expenses false
wraith run AddCategory --category Food --description "Groceries and eating out" \
  --revenues false --expenses true
wraith run AddCategory --category Savings --description "Moving money to the pot" \
  --revenues false --expenses false
wraith run AddCategory --category "Opening balance" \
  --description "What an account already held when it was added" \
  --revenues false --expenses false
```

The last two are transfer categories. A transfer category accepts a positive amount and a negative one, because it is neither income nor expense — which is also what makes it the right classification for money an account already held.

3. Put in what each account already holds, dated the day you start counting from.

```bash
wraith run AddTransaction --account Bank --category "Opening balance" \
  --amount 3000 --date 2026-08-01 --description "Balance when the vault started"
wraith run AddTransaction --account Cash --category "Opening balance" \
  --amount 120 --date 2026-08-01 --description "Cash in hand when the vault started"
```

4. Record what came in and what went out.

```bash
wraith run AddTransaction --account Bank --category Salary \
  --amount 2200 --date 2026-08-05 --description "August salary"
wraith run AddTransaction --account Bank --category Food \
  --amount -32.90 --date 2026-08-18 --description Market
```

5. Move money between two of your own accounts. It is two transactions sharing the transfer category — one account down, the other up — so they net to zero and neither total moves.

```bash
wraith run AddTransaction --account Bank --category Savings \
  --amount -500 --date 2026-08-20 --description "Set aside"
wraith run AddTransaction --account Cash --category Savings \
  --amount 500 --date 2026-08-20 --description "Set aside"
```

6. Declare what repeats. A recurrence moves no money and writes no transaction: it is a rule the forecast reads.

```bash
wraith run AddRecurrence --description Rent --account Bank --category Food \
  --amount -1200 --day 10 --start 2026-08
```

7. Correct a line you got wrong. The `Id` is the one the statement shows.

```bash
wraith run ModifyTransaction --id 4 --amount -40 --description "Market, corrected"
```

8. Read the result:

| Page | What it answers |
|------|-----------------|
| `DashBoard/README.md` | Where you stand today |
| `DashBoard/Categories.md` | Every category, what it accepts, and what it has come to |
| `DashBoard/Accounts/Bank.md` | One account: what it holds, how the open month is going on it, and the menu of every month it has moved in |
| `DashBoard/Months/2026-08/Statement.md` | Every movement dated in August, with its id |
| `DashBoard/Months/2026-08/Accounts/Bank.md` | What that account moved in August, in the order the money moved |
| `DashBoard/Months/README.md` | Every month in three tables — the ones before this one, the open one, and the ones ahead the declared commitments add up to |
