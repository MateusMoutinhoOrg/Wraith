package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/vault"
)

// StartFactory fills api.Lib.Start with a closure that creates a vault: a
// task file and a visualization config, copied from the defaults compiled
// into the binary. It then runs a tick to render the initial state.
//
// The options are the choices the created config is written with: an entry
// named in api.StartOptions.VisualizationArgs is written with those args
// instead of the default's, which is how the interface's `--prev-months`,
// `--future-months` and `--current-month` reach the file. The zero value
// writes the defaults exactly as they are.
//
// It never overwrites. A file already on disk is left exactly as it is — the
// options included, since the file they would be written into is the one
// already there — so running `wraith start` in a vault that is already going
// is harmless rather than destructive, which matters because it is the first
// command anyone runs and the easiest one to run twice.
func StartFactory(l *api.Lib) func(options api.StartOptions) error {
	return func(options api.StartOptions) error {
		wroteTask, err := vault.WriteAsset(l.Deps, vault.StartTaskAsset, l.TaskPath)
		if err != nil {
			return err
		}
		wroteVis, err := vault.WriteStartVisualization(l.Deps, l.VisualizationPath,
			options.VisualizationArgs)
		if err != nil {
			return err
		}
		if wroteTask || wroteVis {
			_, err = l.PerformFullTick()
			return err
		}
		return nil
	}
}
