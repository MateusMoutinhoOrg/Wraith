package visualizations

// The account tree of the DashBoard visualization: one page per account and
// one per credit card, each carrying the menu of every month that account has
// moved in and where the open month stands on it.
//
// There is no index page above them. The dashboard already lists every account
// with what it holds, and every one of those rows is a link — an index would
// have been a second copy of the same table, kept in step by hand.

import (
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// accountPath is where one account's own page is written, below the entry's
// dest. Every link to an account anywhere in the tree is built from it, so the
// folder is named in one place.
func accountPath(account ledger.Account) string {
	return "Accounts/" + slug(account.Name) + ".md"
}

// accountPathOf is the same path, found by the name a movement carries. A name
// no account answers to cannot be reached from the registry, so it is written
// as if it had a page: a movement always names an account that exists, and a
// task is what would have to have gone wrong for it not to.
func accountPathOf(state ledger.State, name string) string {
	account, _ := state.Account(name)
	if account.Name == "" {
		account.Name = name
	}
	return accountPath(account)
}

// settlement renders when a movement actually reaches the balance — the day it
// is dated on, unless it carries a payment date of its own.
func settlement(transaction ledger.Transaction) string {
	if transaction.PaymentDate == transaction.Date {
		return "on the date"
	}
	return utils.PrettyDate(transaction.PaymentDate)
}

// accountPage writes Accounts/<account>.md: where one account stands today,
// how the open month is going on it, and the menu of every month it has moved
// in. The months carrying a page of their own are the ones in `rendered`; the
// rest are listed with their figures and no link, because the dashboard was
// asked for fewer months of history than the account has.
func accountPage(state ledger.State, account ledger.Account, rendered []int64) api.VisualizationRender {
	p := &page{}
	p.heading(1, account.Name)
	p.line("> " + accountKind(account) + " · **Updated:** " + utils.PrettyDate(state.Today))
	p.blank()
	p.line(navigationAt("../"))
	p.blank()
	p.rule()

	accountPosition(p, state, account)
	accountOpenMonth(p, state, account)
	accountMenu(p, state, account, rendered)
	return p.render(accountPath(account))
}

// accountKind renders in words what the account is, so a page never has to be
// read twice to know whether the balance on it is money held or money owed.
func accountKind(account ledger.Account) string {
	if account.IsCard() {
		return "**Credit card** — the balance on it is what you owe"
	}
	return "**Account** — the balance on it is money you hold"
}

// accountPosition writes section 1: what the account holds today, and what it
// is still waiting on.
func accountPosition(p *page, state ledger.State, account ledger.Account) {
	p.heading(2, "1. Where it stands today")
	p.table("Indicator", ">Value", "Where it comes from")
	if account.IsCard() {
		owed := state.Owed(account)
		p.row("Outstanding", "**"+money(owed)+"**", "every purchase − every bill payment")
		p.row("Limit", money(account.Limit), "what the card was declared with")
		p.row("Available", money(account.Limit-owed), "limit − outstanding")
		p.row("Closes", utils.PrettyDate(utils.DateIn(state.OpenMonth(), account.ClosingDay)),
			"day "+strconv.FormatInt(account.ClosingDay, 10)+" of the month")
		p.row("Due", utils.PrettyDate(utils.DateIn(cardDueMonth(state, account), account.DueDay)),
			"day "+strconv.FormatInt(account.DueDay, 10)+" of the month")
	} else {
		p.row("Balance", "**"+money(state.Balance(account))+"**",
			"every settled movement on this account")
	}
	p.row("Pending settlement", signed(state.PendingOf(account)),
		"movements with a payment date still ahead")
	p.row("Movements recorded", strconv.Itoa(state.AccountFlow(account.Name, 0).Count),
		"every movement on this account, all time")
	p.blank()
	p.rule()
}

// cardDueMonth is the month a card's current bill falls due in — the next one
// when the card is due before it closes.
func cardDueMonth(state ledger.State, card ledger.Account) int64 {
	if card.DueDay < card.ClosingDay {
		return utils.AddMonths(state.OpenMonth(), 1)
	}
	return state.OpenMonth()
}

// accountOpenMonth writes section 2: how the month the vault is currently
// writing into is going on this account, movement by movement.
func accountOpenMonth(p *page, state ledger.State, account ledger.Account) {
	open := state.OpenMonth()
	flow := state.AccountFlow(account.Name, open)
	opening := state.BalanceOn(account, utils.DateIn(utils.AddMonths(open, -1), 31))
	p.heading(2, "2. "+utils.PrettyMonth(open)+" so far")
	p.table("Line", ">Value")
	p.row("Balance carried in", money(opening))
	p.row("Money in", signed(flow.In))
	p.row("Money out", signed(flow.Out))
	p.row("**Net movement**", "**"+signed(flow.Net())+"**")
	p.row("Balance at month end", money(state.BalanceOn(account, utils.DateIn(open, 31))))
	p.row("Movements", strconv.Itoa(flow.Count))
	p.blank()

	movements := ledger.OfAccount(state.In(open), account.Name)
	if len(movements) == 0 {
		p.line("Nothing has moved on this account this month yet.")
		p.blank()
		p.rule()
		return
	}
	accountMovements(p, movements, opening)
	p.rule()
}

// accountMovements writes one account's movements as a table with a running
// balance, starting from what it carried in.
func accountMovements(p *page, movements []ledger.Transaction, opening int64) {
	p.table("Date", ">Id", "Category", "Description", ">Amount", ">Balance", "Settles")
	running := opening
	for _, transaction := range movements {
		running += transaction.Amount
		p.row(utils.PrettyDate(transaction.Date), idText(transaction.Id),
			transaction.Category, dash(transaction.Description),
			signed(transaction.Amount), money(running), settlement(transaction))
	}
	p.blank()
	p.line("The `Id` column is what a `ModifyTransaction` or `RemoveTransaction` task addresses " +
		"a line by. A row that has not settled yet is already in the running balance above but " +
		"not in the account's balance — that is the difference between the two.")
	p.blank()
}

// accountMenu writes section 3: the menu of every month this account has moved
// in, newest first, each linking to that month's page for it when the
// dashboard was asked to render that far back.
func accountMenu(p *page, state ledger.State, account ledger.Account, rendered []int64) {
	months := state.AccountMonths(account.Name)
	p.heading(2, "3. Menu — every month this account has moved in")
	if len(months) == 0 {
		p.line("Nothing has moved on this account yet. Record a movement with the " +
			"`AddTransaction` task and its month appears here by itself.")
		return
	}
	p.table("Month", ">In", ">Out", ">Net", ">Movements", ">Balance at month end", "Page")
	for index := len(months) - 1; index >= 0; index-- {
		month := months[index]
		flow := state.AccountFlow(account.Name, month)
		p.row(utils.PrettyMonth(month), signed(flow.In), signed(flow.Out),
			"**"+signed(flow.Net())+"**", strconv.Itoa(flow.Count),
			money(state.BalanceOn(account, utils.DateIn(month, 31))),
			monthPageLink(account, month, rendered))
	}
	p.blank()
	p.line("A month is listed as soon as a movement on this account carries a date inside it — " +
		"future months included. A month older than the `prev-months` the dashboard was asked " +
		"for keeps its figures here but has no page of its own to open.")
}

// monthPageLink renders the menu's link to one month's page for one account,
// or an em dash when that month falls outside what the dashboard rendered.
func monthPageLink(account ledger.Account, month int64, rendered []int64) string {
	if !containsMonth(rendered, month) {
		return dash("")
	}
	folder := "../Months/" + utils.MonthText(month)
	return "[open](" + folder + "/Accounts/" + slug(account.Name) + ".md) · " +
		"[month](" + folder + "/DashBoard.md)"
}
