package publicfunctions

import (
	task "github.com/MateusMoutinhoOrg/Wraith/sandbox/Tasks"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// PerformTaskFactory fills api.Lib.PerformTask with a closure that runs one
// task by name against the library's database. It is the whole of "do
// something": a tick calls it with what it read out of the task file, and the
// `run` command calls it with what it read off the command line.
//
// It writes no file and renders nothing. Re-rendering is a separate step, so
// a caller that wants both — a tick — asks for both.
func PerformTaskFactory(l *api.Lib) func(taskName string, entries map[string]any) error {
	return func(taskName string, entries map[string]any) error {
		return task.Run(l.Deps, l.Deps.KeepLib.NewDatabase(config.DatabaseProps(l.DatabasePath)), taskName, entries)
	}
}
