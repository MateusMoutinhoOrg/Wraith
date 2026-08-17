package visualizer

import (
	"github.com/MateusMoutinhoOrg/Agnos-Cli/sandbox/contracts/api"
)

// HandleVisualizerFactory fills the HandleVisualizer field with a closure
// that produces the visualization output.
func HandleVisualizerFactory(v *api.Visualizer) func() string {
	return func() string {
		// TODO: implement logic
		return ""
	}
}
