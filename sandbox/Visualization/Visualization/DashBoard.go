package visualizations

// The DashBoard visualization: the whole financial picture, as a folder tree.
// It is the one visualization that reads every registry, and the reason the
// rest of the brain exists — a task is only worth running because this is
// what it changes.

import (
	"sort"
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// The name the DashBoard is registered under, the args it declares, and what
// they default to.
const (
	// DashBoardName is the name the catalog registers the DashBoard under,
	// and the name an entry of `Visualization.yaml` asks for it by.
	DashBoardName = "DashBoard"
	// PrevMonthsArg is how many months back of `Months/` are written.
	PrevMonthsArg = "prev-months"
	// FutureMonthsArg is how far ahead the forecast looks.
	FutureMonthsArg = "future-months"
	// DefaultPrevMonths is three months of history.
	DefaultPrevMonths = 3
	// DefaultFutureMonths is eight months of horizon.
	DefaultFutureMonths = 8
	// CurrentMonthArg is the month the dashboard treats as the open one,
	// written YYYY-MM.
	CurrentMonthArg = "current-month"
	// DefaultCurrentMonth is empty, which is the month today falls in — the
	// actual one, read off the injected clock.
	DefaultCurrentMonth = ""
)

// DashBoard returns the visualization that writes the financial vault:
//
//	DashBoard/
//	├── README.md          where you stand today
//	├── Categories.md      every category and what it has cost
//	├── Accounts/          one page per account, with its month menu
//	└── Months/            one folder per month that holds a movement,
//	                       indexed by a page that carries the forecast too
//
// Every figure on those pages is computed from the four registries the tasks
// write. Nothing on them is editable: a hand edit is overwritten on the next
// tick, which is the point — the pages are a view, and the tasks are the
// truth.
func DashBoard() api.Visualizer {
	return api.Visualizer{
		Name:        DashBoardName,
		Description: "The full financial vault: position, registries, months and the forecast",
		Folder:      true,
		Args: []api.Field{
			{Name: PrevMonthsArg, Type: api.NumberField,
				Description: "How many months back of Months/ to write",
				Default:     int64(DefaultPrevMonths)},
			{Name: FutureMonthsArg, Type: api.NumberField,
				Description: "How many months ahead the forecast projects",
				Default:     int64(DefaultFutureMonths)},
			{Name: CurrentMonthArg, Type: api.TextField,
				Description: "The month to render as the open one, as YYYY-MM " +
					"(defaults to the month today falls in)",
				Default: DefaultCurrentMonth},
		},
		HandleVisualizer: func(args api.HandleVisualizationArgs) ([]api.VisualizationRender, error) {
			state := ledger.Load(args.Deps, args.DataBase)
			state = state.AsOf(monthArg(args.Entries, CurrentMonthArg, state.OpenMonth()))
			previous := wholeArg(args.Entries, PrevMonthsArg, DefaultPrevMonths)
			ahead := wholeArg(args.Entries, FutureMonthsArg, DefaultFutureMonths)
			months := state.RenderedMonths(previous)

			renders := []api.VisualizationRender{
				overview(state, months, ahead),
				categoriesPage(state),
				monthsIndex(state, months, ahead),
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
	return "[Dashboard](" + prefix + "README.md) · [Categories](" + prefix +
		"Categories.md) · [Months](" + prefix + "Months/README.md)"
}

// overview writes README.md: where you stand today, how the open month is
// going, and where everything else is.
func overview(state ledger.State, months []int64, ahead int) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Dashboard")
	p.line("> **Updated:** " + utils.PrettyDate(state.Today) + " · **Registry:** " +
		count(len(state.Accounts), "account", "accounts") + ", " +
		count(len(state.Categories), "category", "categories") + ", " +
		count(len(state.Transactions), "transaction", "transactions"))
	p.blank()
	p.line(navigation())
	p.blank()
	p.rule()

	futureExpenses := int64(0)
	for _, transaction := range state.Transactions {
		if transaction.Date > state.Today && transaction.Amount < 0 && !state.IsTransfer(transaction) {
			futureExpenses += transaction.Amount
		}
	}
	netWorth := state.Held() + futureExpenses

	p.heading(2, "1. Position on "+utils.PrettyDate(state.Today))
	p.table("Indicator", ">Value", "Where it comes from")
	p.row("**Balance in accounts**", "**"+money(state.Held())+"**",
		"every movement dated up to today")
	p.row("**Net Worth**", "**"+money(netWorth)+"**",
		"balance minus all future expenses")
	p.row("Movements recorded", strconv.Itoa(len(state.Transactions)),
		"every line the ledger holds")
	p.blank()

	if len(state.Accounts) > 0 {
		p.table("Account", ">Balance", "Share of the money you hold")
		held := state.Held()
		for _, account := range state.Accounts {
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
	p.line("Today's position rolled forward through the movements already dated ahead of " +
		"today. Nothing else is projected.")
	p.blank()
	p.table("Month", ">Income", ">Expenses", ">Held in accounts")
	for _, projection := range state.Forecast(ahead) {
		p.row(utils.PrettyMonth(projection.Month), signed(projection.Income),
			signed(projection.Expenses), money(projection.Held))
	}
	p.blank()
	p.line("The whole projection, month by month: [Months/README.md](Months/README.md)")
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

	var incomeCategories, expenseCategories, volatileCategories []ledger.Category
	for _, category := range state.Categories {
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
		return abs(state.CategoryTotal(incomeCategories[i].Name, open)) > abs(state.CategoryTotal(incomeCategories[j].Name, open))
	})
	sort.Slice(expenseCategories, func(i, j int) bool {
		return abs(state.CategoryTotal(expenseCategories[i].Name, open)) > abs(state.CategoryTotal(expenseCategories[j].Name, open))
	})
	sort.Slice(volatileCategories, func(i, j int) bool {
		return abs(state.CategoryTotal(volatileCategories[i].Name, open)) > abs(state.CategoryTotal(volatileCategories[j].Name, open))
	})

	if len(incomeCategories) > 0 {
		p.heading(2, "Income")
		p.table("Category", "Parent", "What it is", ">This month", ">Prev-Month", ">All time", ">Movements")
		for _, category := range incomeCategories {
			p.row(category.Name, dash(category.Parent),
				dash(category.Description),
				signed(state.CategoryTotal(category.Name, open)),
				signed(state.CategoryTotal(category.Name, utils.AddMonths(open, -1))),
				signed(state.CategoryTotal(category.Name, 0)),
				strconv.Itoa(state.CategoryCount(category.Name)))
		}
		p.blank()
	}

	if len(expenseCategories) > 0 {
		p.heading(2, "Expenses")
		p.table("Category", "Parent", "What it is", ">This month", ">Prev-Month", ">All time", ">Movements")
		for _, category := range expenseCategories {
			p.row(category.Name, dash(category.Parent),
				dash(category.Description),
				signed(state.CategoryTotal(category.Name, open)),
				signed(state.CategoryTotal(category.Name, utils.AddMonths(open, -1))),
				signed(state.CategoryTotal(category.Name, 0)),
				strconv.Itoa(state.CategoryCount(category.Name)))
		}
		p.blank()
	}

	if len(volatileCategories) > 0 {
		p.heading(2, "Volatile")
		p.table("Category", "Parent", "What it is", ">This month", ">Prev-Month", ">All time", ">Movements")
		for _, category := range volatileCategories {
			p.row(category.Name, dash(category.Parent),
				dash(category.Description),
				signed(state.CategoryTotal(category.Name, open)),
				signed(state.CategoryTotal(category.Name, utils.AddMonths(open, -1))),
				signed(state.CategoryTotal(category.Name, 0)),
				strconv.Itoa(state.CategoryCount(category.Name)))
		}
		p.blank()
	}

	p.line("A category accepting neither income nor expense is a **transfer category**: it is " +
		"how money moving between two of your own accounts is recorded without counting as " +
		"either.")
	return p.render("Categories.md")
}

