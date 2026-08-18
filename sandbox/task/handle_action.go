package task

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// HandleActionFactory fills the HandleAction field with a closure
// that executes the task with the given entries.
func HandleActionFactory(t *api.Task) func(entries map[string]any) error {
	return func(entries map[string]any) error {
		// TODO: implement logic
		return nil
	}
}
