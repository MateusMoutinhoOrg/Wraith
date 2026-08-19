package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
)

// removeRecurrenceFields declares what the task accepts.
func removeRecurrenceFields() []api.Field {
	return []api.Field{
		{Name: RecurrenceField, Type: api.TextField, Required: true,
			Description: "The description the recurrence was declared with"},
	}
}

// removeRecurrenceAction deletes one recurrence, addressed by the description
// it was declared with.
func removeRecurrenceAction(args api.HandleActionArgs) error {
	description, err := entries.Text(args.Entries, RecurrenceField)
	if err != nil {
		return err
	}
	if description == "" {
		return errors.New(RecurrenceField + " is required")
	}
	recurrences, reachable := args.DataBase.GetSchema(config.RecurrenceSchema)
	if !reachable {
		return errors.New("the " + config.RecurrenceSchema + " registry is unreachable")
	}
	record, found := recurrences.FindByKey(config.NameField, description)
	if !found {
		return errors.New("recurrence not found: " + description)
	}
	if failure := record.Remove(); failure != nil {
		return errors.New("recurrence could not be removed: " + failure.Message)
	}
	return nil
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
