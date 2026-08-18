package commands

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Start creates a vault: a task file and a visualization config, copied from
// the defaults compiled into the binary. It is the first command anyone runs,
// and it never overwrites — a file already on disk is reported and left
// exactly as it is.
func Start(l *api.Lib, quiet bool) int {
	existed := l.Deps.IoLib.Exist(l.TaskPath) || l.Deps.IoLib.Exist(l.VisualizationPath)
	if err := l.Start(); err != nil {
		return Failure(l, err)
	}
	if quiet {
		return api.ExitOk
	}
	if existed {
		l.Deps.Printf(config.AlreadyStarted, l.TaskPath)
		return api.ExitOk
	}
	l.Deps.Printf(config.Started, l.TaskPath, l.VisualizationPath)
	return api.ExitOk
}
