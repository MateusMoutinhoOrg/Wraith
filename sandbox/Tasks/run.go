// Package task is the registry of every action the brain can perform: the
// array of tasks the binary carries, and the switcher that runs one by name.
//
// It is the seam between "what a brain can do" and everything that asks it to
// do something. A tick reads a name out of Task.yaml, the command line takes
// one as an argument, and both arrive here — so a task written once is
// reachable from every direction, and a name nothing declares is one error
// message rather than several.
package task

import (
	"errors"

	tasks "github.com/MateusMoutinhoOrg/Wraith/sandbox/Tasks/Tasks"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
)

// TaskArray is every task the binary carries, in the order the guides list
// them: the day-to-day ledger first, then what is coming, then the registries
// they lean on.
//
// This is the one place a task is registered. Adding a file to
// sandbox/Tasks/Tasks and a line here is the whole of adding an action —
// the command line, the tick, and the Task-List visualization all read this
// array rather than a list of their own.
func TaskArray() []api.Task {
	return []api.Task{
		tasks.AddTransaction(),
		tasks.ModifyTransaction(),
		tasks.RemoveTransaction(),
		tasks.AddRecurrence(),
		tasks.RemoveRecurrence(),
		tasks.AddCategory(),
		tasks.RemoveCategory(),
		tasks.AddAccount(),
		tasks.RemoveAccount(),
		tasks.AddCreditCard(),
		tasks.RemoveCreditCard(),
	}
}

// Find returns the task carrying the given name. ok is false when no task
// does.
func Find(name string) (api.Task, bool) {
	for _, candidate := range TaskArray() {
		if candidate.Name == name {
			return candidate, true
		}
	}
	return api.Task{}, false
}

// Names returns every task name, for the message an unknown one is answered
// with.
func Names() []string {
	names := []string{}
	for _, candidate := range TaskArray() {
		names = append(names, candidate.Name)
	}
	return names
}

// Run executes one task by name against the given database: it looks the name
// up, checks the fields it was given against the ones the task declares, fills
// in the declared defaults, and hands the result to the task's HandleAction.
//
// Validation happens here rather than inside each task, so every task reports
// a missing field, an unknown field and a field of the wrong type in the same
// words — and a task written tomorrow gets that for free.
func Run(d deps.Deps, database keepdeps.KeepDatabase, name string, values map[string]any) error {
	found, ok := Find(name)
	if !ok {
		return errors.New("unknown task: " + name)
	}
	if err := entries.Validate(found.Fields, values); err != nil {
		return errors.New(name + ": " + err.Error())
	}
	err := found.HandleAction(api.HandleActionArgs{
		Deps:     d,
		DataBase: database,
		Entries:  entries.WithDefaults(found.Fields, values),
	})
	if err != nil {
		return errors.New(name + ": " + err.Error())
	}
	return nil
}
