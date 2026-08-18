package visualizer

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
)

// New builds an api.Visualizer object, storing the injected deps on it and
// running every factory over it to fill its function fields.
func New(d deps.Deps, name string, outs string, description string) api.Visualizer {
	v := api.Visualizer{
		Deps:        d,
		Name:        name,
		Outs:        outs,
		Description: description,
	}
	
	v.HandleVisualizer = HandleVisualizerFactory(&v)
	
	return v
}
