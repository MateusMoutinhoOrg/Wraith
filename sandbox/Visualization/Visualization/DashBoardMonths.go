package visualizations

// The month tree of the DashBoard visualization: the index of every month —
// closed, open and projected — and the three pages each recorded one gets.
//
// A month is never created by hand. It appears as soon as a transaction
// carries a date inside it.
//
// One date is all a movement has: what the month was worth and what its
// accounts moved are the same set of lines, read two ways.

import (
	"sort"
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// monthsIndex writes Months/README.md: the calendar of the vault, in three
// tables — the months already closed, the month still open, and the months
// ahead the movements dated forward reach.
//
// The third table runs past the forecast horizon when a movement is already
// dated further out, so a month that owns a folder always owns a row here
// too: this page is the only index the month folders are reached from.
func monthsIndex(state ledger.State, months []int64, ahead int) api.VisualizationRender {
	p := &page{}
	open := state.OpenMonth()
	p.heading(1, "Months")
	p.line(navigationAt("../"))
	p.blank()
	p.rule()

	p.heading(2, "1. Months before "+utils.PrettyMonth(open))
	if !writeMonthRows(p, state, months, func(month int64) bool { return month < open }) {
		p.line("No earlier month holds a movement. Only the months from `prev-months` back " +
			"are written.")
	}
	p.blank()
	p.rule()

	p.heading(2, "2. "+utils.PrettyMonth(open))
	if !writeMonthRows(p, state, months, func(month int64) bool { return month == open }) {
		p.line("Nothing is recorded in the open month yet. Record one with the " +
			"`AddTransaction` task and it appears here by itself.")
	}
	p.blank()
	p.rule()

	p.heading(2, "3. The next "+strconv.Itoa(ahead)+" months")
	projections := state.Forecast(ahead)
	p.table("Month", ">Income", ">Expenses", ">Held", "Pages")
	for _, projection := range projections {
		p.row(utils.PrettyMonth(projection.Month), signed(projection.Income),
			signed(projection.Expenses), money(projection.Held),
			monthPages(months, projection.Month))
	}
	horizon := open
	if len(projections) > 0 {
		horizon = projections[len(projections)-1].Month
	}
	for _, month := range months {
		if month <= horizon {
			continue
		}
		result := state.MonthResult(month)
		p.row(utils.PrettyMonth(month), signed(result.Income), signed(result.Expenses),
			"—", monthPages(months, month))
	}
	p.blank()
	p.line("Every figure above is a transaction you already dated ahead of today. Nothing " +
		"here is an average of your past, and nothing is projected on your behalf. A month " +
		"ahead only carries pages once a movement is dated inside it, and the months past " +
		"the horizon are listed with their recorded lines alone.")
	return p.render("Months/README.md")
}

// writeMonthRows writes the recorded-month table for the months `keep`
// accepts, newest first, and reports whether any row was written — a table
// with no row is never opened, so the caller can say why instead.
func writeMonthRows(p *page, state ledger.State, months []int64, keep func(int64) bool) bool {
	written := false
	for index := len(months) - 1; index >= 0; index-- {
		month := months[index]
		if !keep(month) {
			continue
		}
		if !written {
			p.table("Month", ">Income", ">Expenses", ">Result", ">Movements", "Pages")
			written = true
		}
		result := state.MonthResult(month)
		p.row(utils.PrettyMonth(month), signed(result.Income), signed(result.Expenses),
			"**"+signed(result.Total())+"**", strconv.Itoa(result.Count),
			monthPages(months, month))
	}
	return written
}

// monthPages links the pages of a month, when it is one of the rendered ones.
func monthPages(months []int64, month int64) string {
	if !containsMonth(months, month) {
		return "—"
	}
	folder := utils.MonthText(month)
	return "[month](" + folder + "/DashBoard.md) · [statement](" + folder + "/Statement.md)"
}

// monthPage writes Months/<month>/DashBoard.md: what the month came to, and
// where it landed.
func monthPage(state ledger.State, month int64) api.VisualizationRender {
	p := &page{}
	folder := utils.MonthText(month)
	result := state.MonthResult(month)
	p.heading(1, utils.PrettyMonth(month))
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
	p.line("What each account moved this month.")
	p.blank()
	p.table("Account", ">Moved this month", ">Balance at month end", "Movements")
	for _, account := range state.Accounts {
		flow := state.AccountFlow(account.Name, month)
		if flow.Count == 0 {
			continue
		}
		p.row("["+account.Name+"](Accounts/"+slug(account.Name)+".md)", signed(flow.Net()),
			money(state.BalanceOn(account, utils.DateIn(month, 31))),
			"[all time](../../"+accountPath(account)+")")
	}
	p.blank()
	p.rule()

	p.heading(2, "3. Categories")

	var incomeCategories, expenseCategories, volatileCategories []ledger.Category
	for _, category := range state.Categories {
		total := state.CategoryTotal(category.Name, month)
		if total == 0 {
			continue
		}
		if category.IsTransfer() || (category.Revenues && category.Expenses) {
			volatileCategories = append(volatileCategories, category)
		} else if category.Revenues {
			incomeCategories = append(incomeCategories, category)
		} else {
			expenseCategories = append(expenseCategories, category)
		}
	}

	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}

	sort.Slice(incomeCategories, func(i, j int) bool {
		return abs(state.CategoryTotal(incomeCategories[i].Name, month)) > abs(state.CategoryTotal(incomeCategories[j].Name, month))
	})
	sort.Slice(expenseCategories, func(i, j int) bool {
		return abs(state.CategoryTotal(expenseCategories[i].Name, month)) > abs(state.CategoryTotal(expenseCategories[j].Name, month))
	})
	sort.Slice(volatileCategories, func(i, j int) bool {
		return abs(state.CategoryTotal(volatileCategories[i].Name, month)) > abs(state.CategoryTotal(volatileCategories[j].Name, month))
	})

	if len(incomeCategories) > 0 {
		p.heading(3, "Income")
		p.table("Category", ">This month")
		for _, category := range incomeCategories {
			p.row(category.Name, signed(state.CategoryTotal(category.Name, month)))
		}
		p.blank()
	}

	if len(expenseCategories) > 0 {
		p.heading(3, "Expenses")
		p.table("Category", ">This month")
		for _, category := range expenseCategories {
			p.row(category.Name, signed(state.CategoryTotal(category.Name, month)))
		}
		p.blank()
	}

	if len(volatileCategories) > 0 {
		p.heading(3, "Volatile")
		p.table("Category", ">This month")
		for _, category := range volatileCategories {
			p.row(category.Name, signed(state.CategoryTotal(category.Name, month)))
		}
		p.blank()
	}

	return p.render("Months/" + folder + "/DashBoard.md")
}

// statementPage writes Months/<month>/Statement.md: every movement dated in
// the month, in order.
func statementPage(state ledger.State, month int64) api.VisualizationRender {
	p := &page{}
	folder := utils.MonthText(month)
	p.heading(1, "Statement — "+utils.PrettyMonth(month))
	p.line("[Month](DashBoard.md) · [Months](../README.md) · [Dashboard](../../README.md)")
	p.blank()
	p.rule()
	movements := state.In(month)
	if len(movements) == 0 {
		p.line("No movement is dated in this month.")
		return p.render("Months/" + folder + "/Statement.md")
	}
	p.table("Date", ">Id", "Account", "Category", "Description", ">Amount")
	running := int64(0)
	for _, transaction := range movements {
		running += transaction.Amount
		p.row(utils.PrettyDate(transaction.Date), idText(transaction.Id),
			"["+transaction.Account+"](../../"+accountPathOf(state, transaction.Account)+")",
			transaction.Category, dash(transaction.Description),
			signed(transaction.Amount))
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
func accountMonthPage(state ledger.State, month int64, account ledger.Account) api.VisualizationRender {
	p := &page{}
	folder := utils.MonthText(month)
	p.heading(1, account.Name+" — "+utils.PrettyMonth(month))
	p.line("[" + account.Name + "](../../../" + accountPath(account) + ") · " +
		"[Month](../DashBoard.md) · [Statement](../Statement.md) · " +
		"[Dashboard](../../../README.md)")
	p.blank()
	p.rule()

	movements := ledger.OfAccount(state.In(month), account.Name)
	opening := state.BalanceOn(account, utils.DateIn(utils.AddMonths(month, -1), 31))
	p.table("Line", ">Value")
	p.row("Balance carried in", money(opening))
	p.row("Movements this month", strconv.Itoa(len(movements)))
	p.blank()

	if len(movements) == 0 {
		p.line("Nothing moved on this account in this month.")
		return p.render("Months/" + folder + "/Accounts/" + slug(account.Name) + ".md")
	}
	accountMovements(p, movements, opening)
	return p.render("Months/" + folder + "/Accounts/" + slug(account.Name) + ".md")
}
