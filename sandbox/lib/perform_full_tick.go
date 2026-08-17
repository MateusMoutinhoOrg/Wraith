package lib

import (
	"github.com/MateusMoutinhoOrg/Agnos-Cli/sandbox/contracts/api"
)

// PerformFullTickFactory fills the PerformFullTick field with a closure
// that runs both task and visualization tick operations.
func PerformFullTickFactory(l *api.Lib) func() error {
	return func() error {
		// TODO: implement logic
		return nil
	}
}
