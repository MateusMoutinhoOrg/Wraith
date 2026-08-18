package lib

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
)

// New builds the api.Lib entry point, storing the injected deps on it and
// running every lib factory over it to fill its function fields. Adding a
// function field to api.Lib means adding its factory call here.
func New(d deps.Deps) api.Lib {
	l := api.Lib{Deps: d}
	
	l.PerformTask = PerformTaskFactory(&l)
	l.PerformVisualization = PerformVisualizationFactory(&l)
	l.PerformTaskTick = PerformTaskTickFactory(&l)
	l.PerformVisualizationTick = PerformVisualizationTickFactory(&l)
	l.PerformFullTick = PerformFullTickFactory(&l)
	l.Sandboxmain = SandboxmainFactory(&l)
	
	return l
}
