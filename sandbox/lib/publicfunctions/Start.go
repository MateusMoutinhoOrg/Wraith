package publicfunctions

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/vault"
)

// StartFactory fills api.Lib.Start with a closure that creates a vault: a
// task file and a visualization config, copied from the defaults compiled
// into the binary.
//
// It never overwrites. A file already on disk is left exactly as it is, so
// running `wraith start` in a vault that is already going is harmless rather
// than destructive — which matters, because it is the first command anyone
// runs and the easiest one to run twice.
func StartFactory(l *api.Lib) func() error {
	return func() error {
		if _, err := vault.WriteAsset(l.Deps, vault.StartTaskAsset, l.TaskPath); err != nil {
			return err
		}
		_, err := vault.WriteAsset(l.Deps, vault.StartVisualizationAsset, l.VisualizationPath)
		return err
	}
}
