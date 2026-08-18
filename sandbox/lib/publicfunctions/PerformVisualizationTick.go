package publicfunctions

import (
	visualization "github.com/MateusMoutinhoOrg/Wraith/sandbox/Visualization"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/vault"
)

// PerformVisualizationTickFactory fills api.Lib.PerformVisualizationTick with
// a closure that renders every enabled entry of the visualization config and
// writes each one under its destination.
//
// The config is validated in full before a single file is written, so a
// broken entry leaves the vault exactly as it was rather than half rewritten.
// Entries are rendered in order, and every one of them reads the same data —
// no entry can see another's output.
func PerformVisualizationTickFactory(l *api.Lib) func() error {
	return func() error {
		entries, err := vault.ReadEntries(l.Deps, l.VisualizationPath, visualization.VisualizationArray())
		if err != nil {
			vault.WriteError(l.Deps, config.ErrorPath, "reading "+l.VisualizationPath, err)
			return err
		}
		for _, entry := range entries {
			if !entry.Enabled {
				continue
			}
			declared, found := visualization.Find(entry.Name)
			if !found {
				continue
			}
			renders, err := l.PerformVisualization(entry.Name, entry.Args)
			if err != nil {
				vault.WriteError(l.Deps, config.ErrorPath, "rendering "+entry.Name, err)
				return err
			}
			if err := vault.Write(l.Deps, entry.Dest, declared.Folder, renders); err != nil {
				vault.WriteError(l.Deps, config.ErrorPath, "writing "+entry.Name, err)
				return err
			}
		}
		return nil
	}
}
