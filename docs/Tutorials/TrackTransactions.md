# Track Transactions

## Description
Records a real month in the financial brain shipped with this repository: accounts, categories, what came in, what went out, a purchase split into installments, and a card bill. Creating the vault first is [StartABrain.md](/docs/Tutorials/StartABrain.md); the mechanics of running any task at all are [RunTasks.md](/docs/Tutorials/RunTasks.md).

### Rules
- A positive `amount` needs a category with `revenues: true`; a negative one needs `expenses: true`.
- A category with **both** false is a **transfer category** — it counts as neither income nor expense, and is how money moving between your own accounts is recorded.
- A card purchase counts on its `date`. The money leaves your bank when you record the bill payment.
- A movement carries two dates, and the pages read them apart. `date` is the month it **counts** in: the month's result, its statement and every category total are built from it. `payment_date` is the day the money actually **moves**: every balance, every account page and the account figures of a month are built from it. Left out, `payment_date` is the `date` and the two questions have the same answer.
- An invoice dated in august and paid in september is therefore august income on `Months/2026-08/DashBoard.md`, and a september movement on `Accounts/Bank.md`.

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

5. Pay the card bill. One task writes both legs — the money leaving the account, and the same amount arriving on the card — under the transfer category, so they net to zero.

```bash
wraith run PayCreditCardBill --card "Nubank Card" --account Bank \
  --category "Card Payment" --date 2026-09-05 --amount 350
```

Leave `amount` out and it pays exactly what the card's closed statements still ask for:

```bash
wraith run PayCreditCardBill --card "Nubank Card" --account Bank \
  --category "Card Payment" --date 2026-09-05
```

Paying a bill never shows up as an expense — the purchases were already counted on the day they happened. Money arriving on a card pays its oldest statement still asking for something, and what is left of it runs on to the next one, so one payment can settle two bills.

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
| `DashBoard/Accounts/Bank.md` | One account: what it holds, how the open month is going on it, and the menu of every month it has moved in — all of it in payment dates |
| `DashBoard/Months/2026-08/Statement.md` | Every movement dated in August, with its id |
| `DashBoard/Months/2026-08/Accounts/Bank.md` | What that account moved in August, in the order the money moved |
| `DashBoard/Credit-Cards.md` | What is outstanding, how much limit is left, and what each card is asking for |
| `DashBoard/Bills/Nubank-Card.md` | One card statement by statement: what each cycle charged, what has been paid against it, and what is left |
| `DashBoard/Pending.md` | Everything still waiting to be paid — the bills to pay, the statements still open, and the movements dated ahead |
| `DashBoard/Months/README.md` | Every month in three tables — the ones before this one, the open one, and the ones ahead the declared commitments add up to |
