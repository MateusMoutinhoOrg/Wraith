package task

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
)

// New builds an api.Task object, storing the injected deps on it and
// running every factory over it to fill its function fields.
func New(d deps.Deps, name string, description string) api.Task {
	t := api.Task{
		Deps:        d,
		Name:        name,
		Description: description,
	}
	
	t.HandleAction = HandleActionFactory(&t)
	
	return t
}
