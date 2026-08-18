package lib

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// PerformTaskTickFactory fills the PerformTaskTick field with a closure
// that runs the periodic task tick operations.
func PerformTaskTickFactory(l *api.Lib) func() error {
	return func() error {
		// TODO: implement logic
		return nil
	}
}
