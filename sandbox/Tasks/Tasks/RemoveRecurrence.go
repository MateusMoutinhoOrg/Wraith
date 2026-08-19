package tasks

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// removeRecurrenceFields declares what the task accepts.
func removeRecurrenceFields() []api.Field {
	return []api.Field{
		{Name: RecurrenceField, Type: api.TextField, Required: true,
			Description: "The description the recurrence was declared with"},
	}
}

// removeRecurrenceAction runs the task against the database it is handed.
func removeRecurrenceAction(args api.HandleActionArgs) error {
	description, err := name(args.Entries, RecurrenceField)
	if err != nil {
		return err
	}
	return remove(args, config.RecurrenceSchema, "recurrence", description)
}

// RemoveRecurrence returns the task that stops a recurring commitment,
// addressing it by the description it was declared with.
//
// Removing one only changes the future: the forecast stops counting it from
// the next tick on, and no transaction is touched — a recurrence never
// created one. To stop a commitment on a date you already know, prefer
// setting `end` on it instead of removing it, so the vault keeps the record
// of what you were once committed to.
func RemoveRecurrence() api.Task {
	return api.Task{
		Name:         "RemoveRecurrence",
		Description:  "Stop a recurring commitment",
		Fields:       removeRecurrenceFields(),
		HandleAction: removeRecurrenceAction,
	}
}
