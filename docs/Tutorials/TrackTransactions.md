# Track Transactions

## Description
Records a real month in the financial brain shipped with this repository: accounts, categories, what came in, what went out, a purchase split into installments, and a card bill. Creating the vault first is [StartABrain.md](/docs/Tutorials/StartABrain.md); the mechanics of running any task at all are [RunTasks.md](/docs/Tutorials/RunTasks.md).

### Rules
- A positive `amount` needs a category with `revenues: true`; a negative one needs `expenses: true`.
- A category with **both** false is a **transfer category** — it counts as neither income nor expense, and is how money moving between your own accounts is recorded.
- A card purchase counts on its `date`. The money leaves your bank when you record the bill payment.

---

## Workflow

1. Create the places money sits.

```bash
wraith run AddAccount --account Bank --opening 3000
wraith run AddAccount --account Cash --opening 120
wraith run AddCreditCard --account "Nubank Card" \
  --limit 5000 --closing_day 25 --due_day 5 --opening -150.50
```

2. Create the categories transactions are classified under.

```bash
wraith run AddCategory --category Salary --description "What comes in" \
  --revenues true --expenses false
wraith run AddCategory --category Food --description "Groceries and eating out" \
  --revenues false --expenses true
wraith run AddCategory --category "Card Payment" --description "Moving money to the card" \
  --revenues false --expenses false
```

The third one is the transfer category. It accepts a positive amount and a negative one, because it is neither income nor expense.

3. Record what came in and what went out.

```bash
wraith run AddTransaction --account Bank --category Salary \
  --amount 2200 --date 2026-08-05 --description "August salary"
wraith run AddTransaction --account Bank --category Food \
  --amount -32.90 --date 2026-08-18 --description Market
```

4. Split a purchase over twelve months. It is still **one** task: `amount` is the total, and each part lands in its own month.

```bash
wraith run AddTransaction --account "Nubank Card" --category Food \
  --amount -4200 --date 2026-08-20 --description Laptop --installments 12
```

Twelve transactions are written at once. The remainder goes to the first part, so the parts add back up to the total exactly.

5. Pay the card bill: two transactions sharing the transfer category, netting to zero.

```bash
wraith run AddTransaction --account Bank --category "Card Payment" \
  --amount -350 --date 2026-09-05 --description "August card bill"
wraith run AddTransaction --account "Nubank Card" --category "Card Payment" \
  --amount 350 --date 2026-09-05 --description "August card bill"
```

Paying a bill never shows up as an expense — the purchases were already counted on the day they happened.

6. Declare what repeats. A recurrence moves no money and writes no transaction: it is a rule the forecast reads.

```bash
wraith run AddRecurrence --description Rent --account Bank --category Food \
  --amount -1200 --day 10 --start 2026-08
```

7. Correct a line you got wrong. The `Id` is the one the statement shows.

```bash
wraith run ModifyTransaction --id 2 --amount -40 --description "Market, corrected"
```

8. Read the result:

| Page | What it answers |
|------|-----------------|
| `DashBoard/README.md` | Where you stand today |
| `DashBoard/Months/2026-08/Statement.md` | Every movement dated in August, with its id |
| `DashBoard/Credit-Cards.md` | What is outstanding, and how much limit is left |
| `DashBoard/Forecast.md` | What the declared commitments add up to |
