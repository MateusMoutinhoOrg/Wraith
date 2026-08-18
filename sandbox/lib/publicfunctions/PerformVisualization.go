package publicfunctions

import (
	visualization "github.com/MateusMoutinhoOrg/Wraith/sandbox/Visualization"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// PerformVisualizationFactory fills api.Lib.PerformVisualization with a
// closure that renders one visualization by name and hands back the files it
// produced.
//
// Nothing is written to disk here. A caller that means to write takes these
// files and puts them under a destination it chose — which is what lets the
// same call serve a tick, a `render` command, and a caller that only wants
// the bytes.
func PerformVisualizationFactory(l *api.Lib) func(name string, entries map[string]any) ([]api.VisualizationRender, error) {
	return func(name string, entries map[string]any) ([]api.VisualizationRender, error) {
		return visualization.Run(l.Deps, store.Database(l.Deps, l.DatabasePath), name, entries)
	}
}
