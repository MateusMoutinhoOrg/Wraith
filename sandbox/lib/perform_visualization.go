package lib

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// PerformVisualizationFactory fills the PerformVisualization field with a closure
// that executes a specific visualizer by name.
func PerformVisualizationFactory(l *api.Lib) func(visualizerName string) (string, error) {
	return func(visualizerName string) (string, error) {
		// TODO: implement logic
		return "", nil
	}
}
