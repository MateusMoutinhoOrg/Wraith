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
		"opening": 2500,
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
