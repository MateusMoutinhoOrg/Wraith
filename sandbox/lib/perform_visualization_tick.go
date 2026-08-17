package lib

import (
	"github.com/MateusMoutinhoOrg/Agnos-Cli/sandbox/contracts/api"
)

// PerformVisualizationTickFactory fills the PerformVisualizationTick field with a closure
// that runs the periodic visualization tick operations.
func PerformVisualizationTickFactory(l *api.Lib) func() error {
	return func() error {
		// TODO: implement logic
		return nil
	}
}
