package lib

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// PerformTaskFactory fills the PerformTask field with a closure
// that executes a specific task by name.
func PerformTaskFactory(l *api.Lib) func(taskName string, entries map[string]any) error {
	return func(taskName string, entries map[string]any) error {
		// TODO: implement logic
		return nil
	}
}
