package visualizations

// The DashBoard visualization: the whole financial picture, as a folder tree.
// It is the one visualization that reads every registry, and the reason the
// rest of the brain exists — a task is only worth running because this is
// what it changes.

import (
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// The args the DashBoard declares, and what they default to.
const (
	// PrevMonthsArg is how many months back of `Months/` are written.
	PrevMonthsArg = "prev-months"
	// FutureMonthsArg is how far ahead the forecast looks.
	FutureMonthsArg = "future-months"
	// DefaultPrevMonths is three months of history.
	DefaultPrevMonths = 3
	// DefaultFutureMonths is eight months of horizon.
	DefaultFutureMonths = 8
)

// DashBoard returns the visualization that writes the financial vault:
//
//	DashBoard/
//	├── README.md          where you stand today
//	├── Credit-Cards.md    every card, its limit and its bill
//	├── Categories.md      every category and what it has cost
//	├── Forecast.md        what the declared commitments add up to
//	├── Accounts/          one page per account and card, with its month menu
//	└── Months/            one folder per month that holds a movement
//
// Every figure on those pages is computed from the five registries the tasks
// write. Nothing on them is editable: a hand edit is overwritten on the next
// tick, which is the point — the pages are a view, and the tasks are the
// truth.
func DashBoard() api.Visualizer {
	return api.Visualizer{
		Name:        "DashBoard",
		Description: "The full financial vault: position, registries, months and the forecast",
		Folder:      true,
		Args: []api.Field{
			{Name: PrevMonthsArg, Type: api.NumberField,
				Description: "How many months back of Months/ to write",
				Default:     int64(DefaultPrevMonths)},
			{Name: FutureMonthsArg, Type: api.NumberField,
				Description: "How many months ahead the forecast projects",
				Default:     int64(DefaultFutureMonths)},
		},
		HandleVisualizer: func(args api.HandleVisualizationArgs) ([]api.VisualizationRender, error) {
			state := ledger.Load(args.Deps, args.DataBase)
			previous := wholeArg(args.Entries, PrevMonthsArg, DefaultPrevMonths)
			ahead := wholeArg(args.Entries, FutureMonthsArg, DefaultFutureMonths)
			months := state.RenderedMonths(previous)

			renders := []api.VisualizationRender{
				overview(state, months, ahead),
				cardsPage(state),
				categoriesPage(state),
				forecastPage(state, ahead),
				monthsIndex(state, months),
			}
			for _, account := range state.Accounts {
				renders = append(renders, accountPage(state, account, months))
			}
			for _, month := range months {
				renders = append(renders, monthPage(state, month), statementPage(state, month))
				for _, account := range state.Accounts {
					if len(ledger.OfAccount(state.In(month), account.Name)) == 0 {
						continue
					}
					renders = append(renders, accountMonthPage(state, month, account))
				}
			}
			return renders, nil
		},
	}
}

// navigation is the line of links every top-level page of the tree carries,
// so no page is more than one click from any other. There is no `Accounts`
// link on it: an account is reached from the row that names it on the
// dashboard, which is the only place the list of them is written.
func navigation() string { return navigationAt("") }

// navigationAt is that same line, written from a page sitting `prefix` away
// from the root of the tree — `../` one folder down, `../../` two.
func navigationAt(prefix string) string {
	return "[Dashboard](" + prefix + "README.md) · [Credit Cards](" + prefix +
		"Credit-Cards.md) · [Categories](" + prefix + "Categories.md) · [Months](" + prefix +
		"Months/README.md) · [Forecast](" + prefix + "Forecast.md)"
}

// overview writes README.md: where you stand today, how the open month is
// going, and where everything else is.
func overview(state ledger.State, months []int64, ahead int) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Dashboard")
	p.line("> **Updated:** " + utils.PrettyDate(state.Today) + " · **Registry:** " +
		count(len(state.PlainAccounts()), "account", "accounts") + ", " +
		count(len(state.Cards()), "credit card", "credit cards") + ", " +
		count(len(state.Categories), "category", "categories") + ", " +
		count(len(state.Transactions), "transaction", "transactions") + ", " +
		count(len(state.Recurrences), "recurrence", "recurrences"))
	p.blank()
	p.line(navigation())
	p.blank()
	p.rule()

	p.heading(2, "1. Position on "+utils.PrettyDate(state.Today))
	p.table("Indicator", ">Value", "Where it comes from")
	p.row("Balance in accounts", "**"+money(state.Held())+"**",
		"opening balances + every settled movement")
	p.row("Owed on credit cards", money(state.TotalOwed()),
		"card opening + purchases − payments")
	p.row("**Net position**", "**"+money(state.Net())+"**", "what you hold − what you owe")
	p.row("Pending settlement", signed(state.Pending()),
		"movements with a payment date still ahead")
	p.blank()

	if len(state.PlainAccounts()) > 0 {
		p.table("Account", ">Balance", "Share of the money you hold")
		held := state.Held()
		for _, account := range state.PlainAccounts() {
			balance := state.Balance(account)
			p.row("["+account.Name+"]("+accountPath(account)+")", money(balance),
				"`"+utils.Bar(balance, held)+"` "+utils.Percent(balance, held))
		}
		p.blank()
	}
	p.rule()

	open := state.OpenMonth()
	result := state.MonthResult(open)
	before := state.MonthResult(utils.AddMonths(open, -1))
	p.heading(2, "2. "+utils.PrettyMonth(open)+" so far")
	p.table("Line", ">This month", ">Previous month", ">Change")
	p.row("Income", signed(result.Income), signed(before.Income),
		signed(result.Income-before.Income))
	p.row("Expenses", signed(result.Expenses), signed(before.Expenses),
		signed(result.Expenses-before.Expenses))
	p.row("**Result**", "**"+signed(result.Total())+"**", signed(before.Total()),
		signed(result.Total()-before.Total()))
	p.row("Transactions", strconv.Itoa(result.Count), strconv.Itoa(before.Count),
		strconv.Itoa(result.Count-before.Count))
	p.blank()
	if containsMonth(months, open) {
		folder := "Months/" + utils.MonthText(open)
		p.line("Full month: [`" + folder + "/DashBoard.md`](" + folder + "/DashBoard.md) · " +
			"ledger: [`" + folder + "/Statement.md`](" + folder + "/Statement.md)")
		p.blank()
	}
	p.rule()

	p.heading(2, "3. The next "+strconv.Itoa(ahead)+" months")
	p.line("Today's position rolled forward through what you declared — " +
		count(len(state.Recurrences), "recurrence", "recurrences") +
		" and the card bills derived from them.")
	p.blank()
	p.table("Month", ">Held in accounts", ">Net position")
	for _, projection := range state.Forecast(ahead) {
		p.row(utils.PrettyMonth(projection.Month), money(projection.Held), money(projection.Net()))
	}
	p.blank()
	p.line("The whole projection, month by month: [Forecast.md](Forecast.md)")
	return p.render("README.md")
}

// containsMonth reports whether a month is one of the rendered ones.
func containsMonth(months []int64, month int64) bool {
	for _, candidate := range months {
		if candidate == month {
			return true
		}
	}
	return false
}

// cardsPage writes Credit-Cards.md: every card, what is outstanding on it,
// and how much of its limit is still available.
func cardsPage(state ledger.State) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Credit Cards")
	p.line(navigation())
	p.blank()
	p.rule()
	if len(state.Cards()) == 0 {
		p.line("No credit card yet. Add one with the `AddCreditCard` task.")
		return p.render("Credit-Cards.md")
	}
	p.table("Card", ">Limit", ">Outstanding", ">Available", "Closes", "Due")
	for _, card := range state.Cards() {
		owed := state.Owed(card)
		closeMonth := state.OpenMonth()
		dueMonth := closeMonth
		if card.DueDay < card.ClosingDay {
			dueMonth = utils.AddMonths(closeMonth, 1)
		}
		p.row("["+card.Name+"]("+accountPath(card)+")", money(card.Limit), money(owed),
			money(card.Limit-owed),
			utils.PrettyDate(utils.DateIn(closeMonth, card.ClosingDay)),
			utils.PrettyDate(utils.DateIn(dueMonth, card.DueDay)))
	}
	p.blank()
	p.line("A purchase counts on the day it happens. The money leaves your account when you " +
		"record the bill payment — two `AddTransaction`s sharing a transfer category.")
	return p.render("Credit-Cards.md")
}

// categoriesPage writes Categories.md: what each category classifies, what it
// accepts, and what it has come to.
func categoriesPage(state ledger.State) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Categories")
	p.line(navigation())
	p.blank()
	p.rule()
	if len(state.Categories) == 0 {
		p.line("No category yet. Add one with the `AddCategory` task.")
		return p.render("Categories.md")
	}
	open := state.OpenMonth()
	p.table("Category", "Accepts", "Parent", "What it is", ">This month", ">All time",
		">Movements")
	for _, category := range state.Categories {
		p.row(category.Name, accepts(category), dash(category.Parent),
			dash(category.Description),
			signed(state.CategoryTotal(category.Name, open)),
			signed(state.CategoryTotal(category.Name, 0)),
			strconv.Itoa(state.CategoryCount(category.Name)))
	}
	p.blank()
	p.line("A category accepting neither income nor expense is a **transfer category**: it is " +
		"how money moving between two of your own accounts is recorded without counting as " +
		"either.")
	return p.render("Categories.md")
}

// accepts renders in words what a category is allowed to classify.
func accepts(category ledger.Category) string {
	if category.IsTransfer() {
		return "transfers"
	}
	if category.Revenues && category.Expenses {
		return "income and expenses"
	}
	if category.Revenues {
		return "income"
	}
	return "expenses"
}

// forecastPage writes Forecast.md: the declared commitments, and the position
// they roll today's balances forward to.
func forecastPage(state ledger.State, ahead int) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Forecast")
	p.line(navigation())
	p.blank()
	p.rule()

	p.heading(2, "1. The next "+strconv.Itoa(ahead)+" months")
	p.table("Month", ">Income", ">Expenses", ">Card bills", ">Held", ">Net position")
	for _, projection := range state.Forecast(ahead) {
		p.row(utils.PrettyMonth(projection.Month), signed(projection.Income),
			signed(projection.Expenses), money(projection.Bills),
			money(projection.Held), money(projection.Net()))
	}
	p.blank()
	p.line("Every figure above is something you declared: a recurrence, a transaction already " +
		"dated ahead, or a card bill derived from what the card owes when its cycle closes. " +
		"Nothing here is an average of your past.")
	p.blank()
	p.rule()

	p.heading(2, "2. The commitments it reads")
	if len(state.Recurrences) == 0 {
		p.line("No recurrence declared. Add one with the `AddRecurrence` task and this page " +
			"starts projecting it.")
		return p.render("Forecast.md")
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
	return p.render("Forecast.md")
}
