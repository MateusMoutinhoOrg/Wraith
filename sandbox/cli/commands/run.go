package commands

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Run executes one task straight from the command line, without touching the
// task file, and then re-renders every visualization exactly as a tick does.
// It is the preferred way to drive this brain from a script, an automation,
// or another tool.
//
// The task file is neither read nor reset: the action comes entirely from the
// flags, one per field the task declares.
func Run(l *api.Lib, quiet bool) int {
	name, err := l.Deps.VerbLib.GetNextStringArg()
	if err != nil || name == "" {
		return UsageError(l, config.MissingTaskName)
	}
	declared, found := find(l.Tasks, name)
	if !found {
		return UsageError(l, config.UnknownTask, name)
	}
	if err := l.PerformTask(declared.Name, Flags(l, declared.Fields)); err != nil {
		return Failure(l, err)
	}
	if err := l.PerformVisualizationTick(); err != nil {
		return Failure(l, err)
	}
	if !quiet {
		l.Deps.Printf(config.Ran, declared.Name)
	}
	return api.ExitOk
}

// find returns the task carrying the given name, out of the ones the library
// was built with.
func find(tasks []api.Task, name string) (api.Task, bool) {
	for _, candidate := range tasks {
		if candidate.Name == name {
			return candidate, true
		}
	}
	return api.Task{}, false
}
