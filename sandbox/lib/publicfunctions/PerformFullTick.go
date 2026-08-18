package publicfunctions

import "github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"

// PerformFullTickFactory fills api.Lib.PerformFullTick with a closure that
// runs one whole tick of the state machine: the pending task first, then
// every visualization.
//
// The order matters and only goes one way. A failed task stops the tick
// before anything is rendered, so the pages on disk always describe data that
// actually exists — the last state that was written successfully, never a
// half-applied one.
func PerformFullTickFactory(l *api.Lib) func() error {
	return func() error {
		if err := l.PerformTaskTick(); err != nil {
			return err
		}
		return l.PerformVisualizationTick()
	}
}
