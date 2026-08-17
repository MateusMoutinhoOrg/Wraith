package api

import (
	"github.com/MateusMoutinhoOrg/Agnos-Cli/sandbox/contracts/deps"
)

type Task struct {
	name         string
	description  string
	HandleAction func(deps deps.Deps, entries map[string]any) error
}

type Visualizer struct {
	name             string
	outs             string
	description      string
	HandleVisualizer func(deps deps.Deps) string
}

type Lib struct {
	Deps           deps.Deps
	Tasks          []Task
	Visualizations []Visualizer
	workDir        string

	PerforTask          func(taskName string, entries map[string]any) error
	PerforVisualization func(visualizerName string) (string, error)

	PerforTaskTick           func() error
	PerformVisualizationTick func() error
	PerformFullTick          func() error

	Sandboxmain func(args []string) int
}
