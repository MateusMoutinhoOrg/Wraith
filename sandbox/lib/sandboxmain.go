package lib

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// SandboxmainFactory fills the Sandboxmain field with a closure
// that serves as the CLI entry point.
func SandboxmainFactory(l *api.Lib) func(args []string) int {
	return func(args []string) int {
		// TODO: implement logic
		return 0
	}
}
