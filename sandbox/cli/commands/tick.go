package commands

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Tick runs one whole tick of the state machine: the pending task, then every
// visualization. A failure is reported here and written to Error.md by the
// tick itself, so a person watching a terminal and a person watching the
// vault see the same thing.
func Tick(l *api.Lib, quiet bool) int {
	if err := l.PerformFullTick(); err != nil {
		return Failure(l, err)
	}
	if !quiet {
		l.Deps.Printf("%s", config.Ticked)
	}
	return api.ExitOk
}
