package publicfunctions

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/vault"
)

// PerformTaskTickFactory fills api.Lib.PerformTaskTick with a closure that
// runs the task half of a tick: read the task file, decide whether there is
// anything to do, do it, and disarm the file either way.
//
// The order is the whole contract, and it is worth reading once:
//
//  1. No task file at all — nothing to do, and no error. A vault that has not
//     been started yet is not a broken vault.
//  2. `apply: false` — nothing to do, and no error either. That is a task
//     waiting to be armed, which is the normal state between two edits.
//  3. Otherwise run it. On failure, write Error.md, disarm the file, and
//     report — so a task that cannot work is not retried on every tick of a
//     `watch` loop.
func PerformTaskTickFactory(l *api.Lib) func() error {
	return func() error {
		pending, err := vault.ReadTask(l.Deps, l.TaskPath)
		if errors.Is(err, vault.ErrNoTask) {
			return nil
		}
		if err != nil {
			vault.WriteError(l.Deps, config.ErrorPath, "reading "+l.TaskPath, err)
			return err
		}
		if !pending.Apply {
			return nil
		}
		if pending.Name == "" {
			failure := errors.New(l.TaskPath + " carries no `name` — it must name one of the " +
				"tasks `wraith tasks` lists")
			vault.WriteError(l.Deps, config.ErrorPath, "reading "+l.TaskPath, failure)
			vault.ResetApply(l.Deps, l.TaskPath, pending)
			return failure
		}
		failure := l.PerformTask(pending.Name, pending.Fields())
		if resetErr := vault.ResetApply(l.Deps, l.TaskPath, pending); resetErr != nil && failure == nil {
			failure = resetErr
		}
		if failure != nil {
			vault.WriteError(l.Deps, config.ErrorPath, "running "+pending.Name, failure)
			return failure
		}
		vault.ClearError(l.Deps, config.ErrorPath)
		return nil
	}
}
