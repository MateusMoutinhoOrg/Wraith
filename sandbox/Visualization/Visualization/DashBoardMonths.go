package visualizations

// The month tree of the DashBoard visualization: the index of every month —
// closed, open and projected — and the three pages each recorded one gets.
//
// A month is never created by hand. It appears as soon as a transaction
// carries a date inside it, or settles inside it — which is why a purchase
// split into twelve parts opens the next eleven months by itself.
//
// The month reads the ledger through both dates, and the pages keep the two
// apart. The result, the statement and the categories are read through the
// `date` a movement counts on: they are what the month was worth. The
// accounts section and the per-account pages are read through `payment_date`:
// they are what actually moved.

import (
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// monthsIndex writes Months/README.md: the calendar of the vault, in three
// tables — the months already closed, the month still open, and the months
// ahead the declared commitments reach — followed by the commitments those
// projections read.
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

	p.heading(2, "1. Months before this one")
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
	p.table("Month", ">Income", ">Expenses", ">Card bills", ">Held", ">Net position", "Pages")
	for _, projection := range projections {
		p.row(utils.PrettyMonth(projection.Month), signed(projection.Income),
			signed(projection.Expenses), money(projection.Bills), money(projection.Held),
			money(projection.Net()), monthPages(months, projection.Month))
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
			"—", "—", "—", monthPages(months, month))
	}
	p.blank()
	p.line("Every figure above is something you declared: a recurrence, a transaction already " +
		"dated ahead, or a card bill derived from what the card owes when its cycle closes. " +
		"Nothing here is an average of your past. A month ahead only carries pages once a " +
		"movement is dated inside it — an installment does that by itself, and the months it " +
		"opens past the horizon are listed with their recorded lines alone.")
	p.blank()
	p.rule()

	p.heading(2, "4. The commitments the projection reads")
	if len(state.Recurrences) == 0 {
		p.line("No recurrence declared. Add one with the `AddRecurrence` task and the months " +
			"ahead start projecting it.")
		return p.render("Months/README.md")
	}
	p.table("Recurrence", "Account", "Category", ">Amount", "Day", "From", "Until")
	for _, recurrence := range state.Recurrences {
		destination := recurrence.Account
		if recurrence.ToAccount != "" {
			destination = recurrence.Account + " → " + recurrence.ToAccount
		}
		until := "open-ended"
		if recurrence.End != 0 {
			until = utils.MonthText(recurrence.End)
		}
		p.row(recurrence.Description, destination, recurrence.Category,
			signed(recurrence.Amount), strconv.FormatInt(recurrence.Day, 10),
			utils.MonthText(recurrence.Start), until)
	}
	p.blank()
	p.line("A recurrence never becomes a transaction on its own. When the day arrives you " +
		"record what actually happened with `AddTransaction`.")
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

// monthPage writes Months/<month>/DashBoard.md: what the month came to, where
// it landed, and what it is still waiting on.
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
	p.line("What each account actually moved this month — read through payment dates, " +
		"so the figures agree with the balances.")
	p.blank()
	p.table("Account", ">Moved this month", ">Balance at month end", "Movements")
	for _, account := range state.Accounts {
		flow := state.AccountFlow(account.Name, month)
		if flow.Count == 0 {
			continue
		}
		p.row("["+account.Name+"](../../"+accountPath(account)+")", signed(flow.Net()),
			money(state.BalanceOn(account, utils.DateIn(month, 31))),
			"[this month](Accounts/"+slug(account.Name)+".md)")
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
		p.row(utils.PrettyDate(entry.Date), entry.Description, entry.Account,
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
	p.table("Date", ">Id", "Account", "Category", "Description", ">Amount", "Settles")
	running := int64(0)
	for _, transaction := range movements {
		running += transaction.Amount
		p.row(utils.PrettyDate(transaction.Date), idText(transaction.Id),
			"["+transaction.Account+"](../../"+accountPathOf(state, transaction.Account)+")",
			transaction.Category, dash(transaction.Description),
			signed(transaction.Amount), settlement(transaction))
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

	movements := ledger.OfAccount(state.SettledIn(month), account.Name)
	opening := state.BalanceOn(account, utils.DateIn(utils.AddMonths(month, -1), 31))
	p.table("Line", ">Value")
	p.row("Balance carried in", money(opening))
	if account.IsCard() {
		p.row("Limit", money(account.Limit))
		p.row("Outstanding today", money(state.Owed(account)))
	}
	p.row("Movements this month", strconv.Itoa(len(movements)))
	p.blank()

	if len(movements) == 0 {
		p.line("Nothing moved on this account in this month.")
		return p.render("Months/" + folder + "/Accounts/" + slug(account.Name) + ".md")
	}
	accountMovements(p, movements, opening)
	return p.render("Months/" + folder + "/Accounts/" + slug(account.Name) + ".md")
}
