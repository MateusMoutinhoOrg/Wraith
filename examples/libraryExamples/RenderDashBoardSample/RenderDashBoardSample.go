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
	l := wraithlib.New(deps, "data")

	if err := l.PerformTask("AddAccount", map[string]any{
		"account": "Emergency Savings",
	}); err != nil {
		panic(err)
	}

	// An account is born empty: what it already held is a movement like any
	// other, so it needs a category and a date of its own.
	if err := l.PerformTask("AddCategory", map[string]any{
		"category":    "Opening balance",
		"description": "What an account already held when it was added",
		"revenues":    false,
		"expenses":    false,
	}); err != nil {
		panic(err)
	}
	if err := l.PerformTask("AddTransaction", map[string]any{
		"account":     "Emergency Savings",
		"category":    "Opening balance",
		"amount":      2500,
		"date":        "2026-08-01",
		"description": "Balance when the vault started",
	}); err != nil {
		panic(err)
	}

	// The args map is the same one a Visualization.yaml entry's `args:` block
	// decodes to. Anything omitted falls back to the declared default.
	renders, err := l.PerformVisualization("DashBoard", map[string]any{
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
