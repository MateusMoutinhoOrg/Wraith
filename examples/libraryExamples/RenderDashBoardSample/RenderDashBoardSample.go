package main

// RenderDashBoardSample — rendering a visualization to bytes, not to disk.
//
// api.Lib.PerformVisualization hands back the files a visualization produced
// and writes nothing itself. That is what lets the same renderer serve a
// tick, a `wraith render`, and a program like this one that only wants the
// markdown — to serve over HTTP, to diff against yesterday's, to paste into a
// message.
//
// Run it with:
//   go run ./examples/libraryExamples/RenderDashBoardSample/RenderDashBoardSample.go

import (
	"fmt"
	"os"
	"time"

	wraithadapter "github.com/MateusMoutinhoOrg/Wraith/adapters/standard"
	wraithlib "github.com/MateusMoutinhoOrg/Wraith/sandbox"
	wraithtypes "github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

func main() {
	vault, err := os.MkdirTemp("", "wraith-sample")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(vault)

	deps := wraithadapter.New(vault)

	// Every effect the library has is a function field on deps, the clock
	// included — so a program that wants a reproducible page replaces it
	// rather than waiting for the right day. Nothing inside the sandbox can
	// tell the difference. Driving the CLI instead, where the clock is the
	// real one, the same pinning is `--current-month`.
	deps.Now = func() time.Time {
		return time.Date(2026, time.August, 20, 12, 0, 0, 0, time.UTC)
	}

	l := wraithlib.New(deps, "data")

	for _, account := range []string{"Checking", "Emergency Savings"} {
		if err := l.PerformTask("AddAccount", map[string]any{
			"account": account,
		}); err != nil {
			panic(err)
		}
	}

	// An account is born empty: what it already held is a movement like any
	// other, so it needs a category and a date of its own.
	categories := []map[string]any{
		{"category": "Opening balance", "revenues": false, "expenses": false,
			"description": "What an account already held when it was added"},
		{"category": "Salary", "revenues": true, "expenses": false,
			"description": "Monthly pay"},
		{"category": "Housing", "revenues": false, "expenses": true,
			"description": "Rent and utilities"},
		{"category": "Savings", "revenues": false, "expenses": false,
			"description": "Money moved between my own accounts"},
	}
	for _, category := range categories {
		if err := l.PerformTask("AddCategory", category); err != nil {
			panic(err)
		}
	}

	// Four months of movements, so the rendered tree has months to show.
	// The dates are written out rather than derived from the clock, and the
	// clock is pinned above, so this sample renders the same bytes whenever
	// it is run.
	transactions := []map[string]any{
		{"account": "Checking", "category": "Opening balance", "amount": 1200,
			"date": "2026-05-01", "description": "Balance when the vault started"},
		{"account": "Emergency Savings", "category": "Opening balance", "amount": 2500,
			"date": "2026-05-01", "description": "Balance when the vault started"},

		{"account": "Checking", "category": "Salary", "amount": 4800,
			"date": "2026-05-05", "description": "May pay"},
		{"account": "Checking", "category": "Housing", "amount": -1650,
			"date": "2026-05-10", "description": "May rent"},
		{"account": "Checking", "category": "Savings", "amount": -600,
			"date": "2026-05-20", "description": "May set aside"},
		{"account": "Emergency Savings", "category": "Savings", "amount": 600,
			"date": "2026-05-20", "description": "May set aside"},

		{"account": "Checking", "category": "Salary", "amount": 4800,
			"date": "2026-06-05", "description": "June pay"},
		{"account": "Checking", "category": "Housing", "amount": -1650,
			"date": "2026-06-10", "description": "June rent"},
		{"account": "Checking", "category": "Savings", "amount": -600,
			"date": "2026-06-20", "description": "June set aside"},
		{"account": "Emergency Savings", "category": "Savings", "amount": 600,
			"date": "2026-06-20", "description": "June set aside"},

		{"account": "Checking", "category": "Salary", "amount": 4800,
			"date": "2026-07-05", "description": "July pay"},
		{"account": "Checking", "category": "Housing", "amount": -1650,
			"date": "2026-07-10", "description": "July rent"},
		{"account": "Checking", "category": "Savings", "amount": -600,
			"date": "2026-07-20", "description": "July set aside"},
		{"account": "Emergency Savings", "category": "Savings", "amount": 600,
			"date": "2026-07-20", "description": "July set aside"},

		// August has only just opened: pay and rent both land on the first.
		{"account": "Checking", "category": "Salary", "amount": 4800,
			"date": "2026-08-01", "description": "August pay"},
		{"account": "Checking", "category": "Housing", "amount": -1650,
			"date": "2026-08-01", "description": "August rent"},

		// Already dated ahead of the open month, which is the whole forecast.
		{"account": "Checking", "category": "Salary", "amount": 4800,
			"date": "2026-09-05", "description": "September pay"},
		{"account": "Checking", "category": "Housing", "amount": -1650,
			"date": "2026-09-10", "description": "September rent"},
	}
	for _, transaction := range transactions {
		if err := l.PerformTask("AddTransaction", transaction); err != nil {
			panic(err)
		}
	}

	// The args map is the same one a Visualization.yaml entry's `args:` block
	// decodes to. Anything omitted falls back to the declared default.
	renders, err := l.PerformVisualization("DashBoard", map[string]any{
		"prev-months":   4,
		"future-months": 4,
	})
	if err != nil {
		panic(err)
	}

	// Each render is one file the visualization would write, and where it
	// would go below the entry's dest.
	var overview wraithtypes.VisualizationRender
	for _, render := range renders {
		fmt.Printf("%6d bytes  %s\n", len(render.Content), render.Path)
		if render.Path == "README.md" {
			overview = render
		}
	}

	fmt.Println()
	fmt.Println(string(overview.Content))
}
