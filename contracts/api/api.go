// Package api defines the output contracts of the sandbox, representing
// the objects and functions the library exposes to callers.
package api

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
)

// Task represents an executable action within the library.
// Its behavior is filled by the factories in sandbox/task/.
type Task struct {
	// Deps is the dependency set propagated from the lib that created
	// this object, carried here so its factories can reach it.
	Deps deps.Deps
	
	// Name is the identifier of the task.
	Name string
	
	// Description explains what the task does.
	Description string
	
	// HandleAction executes the task with the given entries.
	// It is filled by task.HandleActionFactory.
	HandleAction func(entries map[string]any) error
}

// Visualizer represents a component that generates visual output.
// Its behavior is filled by the factories in sandbox/visualizer/.
type Visualizer struct {
	// Deps is the dependency set propagated from the lib that created
	// this object, carried here so its factories can reach it.
	Deps deps.Deps
	
	// Name is the identifier of the visualizer.
	Name string
	
	// Outs specifies the output format or target.
	Outs string
	
	// Description explains what the visualizer displays.
	Description string
	
	// HandleVisualizer produces the visualization output.
	// It is filled by visualizer.HandleVisualizerFactory.
	HandleVisualizer func() string
}

// Lib is the entry point handed back by lib.New. Every object it
// creates carries the same deps it was built with.
type Lib struct {
	// Deps is the dependency set injected by lib.New, carried here so
	// every factory-built function field can reach it.
	Deps deps.Deps
	
	// Tasks is a collection of available tasks.
	Tasks []Task
	
	// Visualizations is a collection of available visualizers.
	Visualizations []Visualizer
	
	// WorkDir is the working directory for operations.
	WorkDir string

	// PerformTask executes a specific task by name.
	// It is filled by lib.PerformTaskFactory.
	PerformTask func(taskName string, entries map[string]any) error
	
	// PerformVisualization executes a specific visualizer by name.
	// It is filled by lib.PerformVisualizationFactory.
	PerformVisualization func(visualizerName string) (string, error)

	// PerformTaskTick runs the periodic task tick operations.
	// It is filled by lib.PerformTaskTickFactory.
	PerformTaskTick func() error
	
	// PerformVisualizationTick runs the periodic visualization tick operations.
	// It is filled by lib.PerformVisualizationTickFactory.
	PerformVisualizationTick func() error
	
	// PerformFullTick runs both task and visualization tick operations.
	// It is filled by lib.PerformFullTickFactory.
	PerformFullTick func() error

	// Sandboxmain is the CLI entry point.
	// It is filled by lib.SandboxmainFactory.
	Sandboxmain func(args []string) int
}
