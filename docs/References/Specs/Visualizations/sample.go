//go:build ignore

// This file is an illustrative sample, not part of the build.
package visualizations

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
)

// The args this visualization declares, and what they default to.
const (
	// TopArg is how many accounts the page lists.
	TopArg = "top"
	// DefaultTop is five, which fits on a screen.
	DefaultTop = 5
)

// Summary returns the visualization that writes where you stand, on a single
// page. It is a file visualization: its dest is the file itself, so it hands
// back exactly one render with an empty path.
func Summary() api.Visualizer {
	return api.Visualizer{
		Name:        "Summary",
		Description: "Where you stand, on a single page",
		Folder:      false,
		Args: []api.Field{
			{Name: TopArg, Type: api.NumberField,
				Description: "How many accounts to list",
				Default:     int64(DefaultTop)},
		},
		HandleVisualizer: func(args api.HandleVisualizationArgs) ([]api.VisualizationRender, error) {
			// The whole registry is read once, so no two lines of the page
			// can disagree about what the data says.
			state := ledger.Load(args.Deps, args.DataBase)
			top := wholeArg(args.Entries, TopArg, DefaultTop)
			return []api.VisualizationRender{summaryPage(state, top)}, nil
		},
	}
}

// summaryPage writes the page itself. Every figure is asked of the ledger;
// nothing is computed here.
func summaryPage(state ledger.State, top int) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Summary")
	p.table("Account", ">Balance")
	for index, account := range state.PlainAccounts() {
		if index >= top {
			break
		}
		p.row(account.Name, money(state.Balance(account)))
	}
	p.blank()
	p.line("**Net position:** " + money(state.Net()))

	// An empty path: a file visualization is written at its dest.
	return p.render("")
}
