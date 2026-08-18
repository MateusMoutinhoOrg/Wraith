// Package visualization is the registry of every renderer the brain carries:
// the array of visualizations, and the switcher that renders one by name.
//
// It is the mirror of sandbox/Tasks: a tick reads names out of
// Visualization.yaml, the command line takes one as an argument, and both
// arrive here. Writing the bytes to disk happens one layer up, in
// sandbox/lib — a visualization returns files, it does not create them.
package visualization

import (
	"errors"

	visualizations "github.com/MateusMoutinhoOrg/Wraith/sandbox/Visualization/Visualization"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
)

// VisualizationArray is every visualization the binary can render. It is the
// catalog declared beside the renderers themselves, so the Help guide and
// this switcher can never disagree about what exists.
func VisualizationArray() []api.Visualizer {
	return visualizations.Catalog()
}

// Find returns the visualization carrying the given name. ok is false when no
// visualization does.
func Find(name string) (api.Visualizer, bool) {
	for _, candidate := range VisualizationArray() {
		if candidate.Name == name {
			return candidate, true
		}
	}
	return api.Visualizer{}, false
}

// Names returns every visualization name, for the message an unknown one is
// answered with.
func Names() []string {
	names := []string{}
	for _, candidate := range VisualizationArray() {
		names = append(names, candidate.Name)
	}
	return names
}

// Run renders one visualization by name against the given database and hands
// back the files it produced. It looks the name up, checks the args it was
// given against the ones the visualization declares, fills in the declared
// defaults, and calls HandleVisualizer.
//
// Nothing is written here. A caller that means to write — a tick, or the
// `render` command — takes these files and puts them under the entry's dest.
func Run(d deps.Deps, database keepdeps.KeepDatabase, name string, args map[string]any) ([]api.VisualizationRender, error) {
	found, ok := Find(name)
	if !ok {
		return nil, errors.New("unknown visualization: " + name)
	}
	if err := entries.Validate(found.Args, args); err != nil {
		return nil, errors.New(name + ": " + err.Error())
	}
	renders, err := found.HandleVisualizer(api.HandleVisualizationArgs{
		Deps:     d,
		DataBase: database,
		Entries:  entries.WithDefaults(found.Args, args),
	})
	if err != nil {
		return nil, errors.New(name + ": " + err.Error())
	}
	return renders, nil
}
