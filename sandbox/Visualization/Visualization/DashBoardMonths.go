package visualizations

// The month tree of the DashBoard visualization: the index of every month
// that holds a movement, and the three pages each one gets.
//
// A month is never created by hand. It appears as soon as a transaction
// carries a date inside it — which is why a purchase split into twelve parts
// opens the next eleven months by itself.

import (
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// monthsIndex writes Months/README.md: one row per rendered month, with what
// it came to.
func monthsIndex(state ledger.State, months []int64) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Months")
	p.line("[Dashboard](../README.md) · [Accounts](../Accounts.md) · " +
		"[Credit Cards](../Credit-Cards.md) · [Categories](../Categories.md) · " +
		"[Forecast](../Forecast.md)")
	p.blank()
	p.rule()
	if len(months) == 0 {
		p.line("No month holds a movement yet. Record one with the `AddTransaction` task and " +
			"its month appears here by itself.")
		return p.render("Months/README.md")
	}
	p.table("Month", ">Income", ">Expenses", ">Result", ">Movements", "Pages")
	for index := len(months) - 1; index >= 0; index-- {
		month := months[index]
		result := state.MonthResult(month)
		folder := store.MonthText(month)
		p.row(store.PrettyMonth(month), signed(result.Income), signed(result.Expenses),
			"**"+signed(result.Total())+"**", strconv.Itoa(result.Count),
			"[month]("+folder+"/DashBoard.md) · [statement]("+folder+"/Statement.md)")
	}
	return p.render("Months/README.md")
}

// monthPage writes Months/<month>/DashBoard.md: what the month came to, where
// it landed, and what it is still waiting on.
func monthPage(state ledger.State, month int64) api.VisualizationRender {
	p := &page{}
	folder := store.MonthText(month)
	result := state.MonthResult(month)
	p.heading(1, store.PrettyMonth(month))
	p.line("[Months](../README.md) · [Statement](Statement.md) · [Dashboard](../../README.md)")
	p.blank()
	p.rule()

	p.heading(2, "1. Result")
	p.table("Line", ">Value")
	p.row("Income", signed(result.Income))
	p.row("Expenses", signed(result.Expenses))
	p.row("**Result**", "**"+signed(result.Total())+"**")
	p.row("Movements", strconv.Itoa(result.Count))
	p.blank()
	p.rule()

	p.heading(2, "2. Accounts")
	movements := state.In(month)
	p.table("Account", ">Moved this month", ">Balance at month end", "Statement")
	for _, account := range state.Accounts {
		inMonth := ledger.OfAccount(movements, account.Name)
		if len(inMonth) == 0 {
			continue
		}
		moved := int64(0)
		for _, transaction := range inMonth {
			moved += transaction.Amount
		}
		p.row(account.Name, signed(moved),
			money(state.BalanceOn(account, store.DateIn(month, 31))),
			"[Accounts/"+slug(account.Name)+".md](Accounts/"+slug(account.Name)+".md)")
	}
	p.blank()
	p.rule()

	p.heading(2, "3. Categories")
	p.table("Category", ">This month")
	for _, category := range state.Categories {
		total := state.CategoryTotal(category.Name, month)
		if total == 0 {
			continue
		}
		p.row(category.Name, signed(total))
	}
	p.blank()
	p.rule()

	p.heading(2, "4. Commitments dated in this month")
	due := state.DueIn(month)
	if len(due) == 0 {
		p.line("Nothing declared for this month.")
		return p.render("Months/" + folder + "/DashBoard.md")
	}
	p.table("Date", "Commitment", "Account", ">Amount")
	for _, entry := range due {
		p.row(store.PrettyDate(entry.Date), entry.Description, entry.Account,
			signed(entry.Amount))
	}
	p.blank()
	p.line("These are declared commitments, not recorded movements — they change no balance " +
		"until you record what actually happened.")
	return p.render("Months/" + folder + "/DashBoard.md")
}

// statementPage writes Months/<month>/Statement.md: every movement dated in
// the month, in order.
func statementPage(state ledger.State, month int64) api.VisualizationRender {
	p := &page{}
	folder := store.MonthText(month)
	p.heading(1, "Statement — "+store.PrettyMonth(month))
	p.line("[Month](DashBoard.md) · [Months](../README.md) · [Dashboard](../../README.md)")
	p.blank()
	p.rule()
	movements := state.In(month)
	if len(movements) == 0 {
		p.line("No movement is dated in this month.")
		return p.render("Months/" + folder + "/Statement.md")
	}
	p.table("Date", ">Id", "Account", "Category", "Description", ">Amount", "Settles")
	running := int64(0)
	for _, transaction := range movements {
		running += transaction.Amount
		settles := "on the date"
		if transaction.PaymentDate != transaction.Date {
			settles = store.PrettyDate(transaction.PaymentDate)
		}
		p.row(store.PrettyDate(transaction.Date), store.IdText(transaction.Id),
			transaction.Account, transaction.Category, dash(transaction.Description),
			signed(transaction.Amount), settles)
	}
	p.blank()
	p.line("**Movements:** " + strconv.Itoa(len(movements)) + " · **Net movement:** " +
		signed(running))
	p.blank()
	p.line("The `Id` column is what a `ModifyTransaction` task addresses a line by.")
	return p.render("Months/" + folder + "/Statement.md")
}

// accountMonthPage writes Months/<month>/Accounts/<account>.md: one account's
// movements inside one month, with a running balance.
func accountMonthPage(state ledger.State, month int64, account store.Account) api.VisualizationRender {
	p := &page{}
	folder := store.MonthText(month)
	p.heading(1, account.Name+" — "+store.PrettyMonth(month))
	p.line("[Month](../DashBoard.md) · [Statement](../Statement.md) · " +
		"[Dashboard](../../../README.md)")
	p.blank()
	p.rule()

	movements := ledger.OfAccount(state.In(month), account.Name)
	opening := state.BalanceOn(account, store.DateIn(store.AddMonths(month, -1), 31))
	p.table("Line", ">Value")
	p.row("Balance carried in", money(opening))
	if account.IsCard() {
		p.row("Limit", money(account.Limit))
		p.row("Outstanding today", money(state.Owed(account)))
	}
	p.row("Movements this month", strconv.Itoa(len(movements)))
	p.blank()

	if len(movements) == 0 {
		p.line("No movement on this account in this month.")
		return p.render("Months/" + folder + "/Accounts/" + slug(account.Name) + ".md")
	}
	p.table("Date", ">Id", "Category", "Description", ">Amount", ">Balance")
	running := opening
	for _, transaction := range movements {
		running += transaction.Amount
		p.row(store.PrettyDate(transaction.Date), store.IdText(transaction.Id),
			transaction.Category, dash(transaction.Description),
			signed(transaction.Amount), money(running))
	}
	p.blank()
	p.line("The balance column follows the movements in date order. What a movement settles " +
		"on — its `payment_date` — is what decides when it reaches the balance on the " +
		"dashboard.")
	return p.render("Months/" + folder + "/Accounts/" + slug(account.Name) + ".md")
}
