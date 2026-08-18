package lib

import (
	task "github.com/MateusMoutinhoOrg/Wraith/sandbox/Tasks"
	visualization "github.com/MateusMoutinhoOrg/Wraith/sandbox/Visualization"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/publicfunctions"
)

// New builds the api.Lib entry point, storing the injected deps and the
// database path on it and running every lib factory over it to fill its
// function fields. Adding a function field to api.Lib means adding its
// factory call here.
//
// The two registries — every task the binary carries and every visualization
// it can render — are read once, here, out of the switchers that declare
// them. Everything downstream reads them off the returned api.Lib, so there
// is exactly one answer in the process to "what can this brain do".
func New(d deps.Deps, databasePath string) api.Lib {
	if databasePath == "" {
		databasePath = config.DefaultDatabasePath
	}
	l := api.Lib{
		Deps:              d,
		DatabasePath:      databasePath,
		TaskPath:          config.DefaultTaskPath,
		VisualizationPath: config.DefaultVisualizationPath,
		Tasks:             task.TaskArray(),
		Visualizations:    visualization.VisualizationArray(),
	}
	l.PerformTask = publicfunctions.PerformTaskFactory(&l)
	l.PerformVisualization = publicfunctions.PerformVisualizationFactory(&l)
	l.PerformTaskTick = publicfunctions.PerformTaskTickFactory(&l)
	l.PerformVisualizationTick = publicfunctions.PerformVisualizationTickFactory(&l)
	l.PerformFullTick = publicfunctions.PerformFullTickFactory(&l)
	l.Start = publicfunctions.StartFactory(&l)
	l.Sandboxmain = publicfunctions.SandboxmainFactory(&l)
	return l
}
