package visualizations

// The catalog: every visualization this package declares, in the order the
// guides list them.
//
// It lives beside the visualizations rather than in the switcher next door
// because the Help visualization renders the catalog itself — a guide that
// listed the renderers by hand would be one more thing to keep in step.

import "github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"

// Catalog returns every visualization the binary can render. Adding a file to
// this package and a line here is the whole of adding a renderer:
// sandbox/Visualization/run.go, the command line and the guides all read this
// list rather than one of their own.
func Catalog() []api.Visualizer {
	return []api.Visualizer{
		DashBoard(),
		TaskList(),
		Help(),
		Usage(),
	}
}
