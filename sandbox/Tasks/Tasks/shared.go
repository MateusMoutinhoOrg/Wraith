// Package tasks holds every action the brain can perform, one task per file.
//
// A task is a value, not a function: it carries its name, what it is for, the
// fields it accepts, and the closure that runs it. Because the fields are
// declared rather than parsed by hand, one validator checks every task, the
// command line gets one `--flag` per field for free, and the Task-List
// visualization can document a task it has never heard of.
//
// Adding a task is adding one file here that returns an api.Task, and one
// line in sandbox/Tasks/run.go's TaskArray. Nothing else in the project has
// to learn about it. See docs/Tutorials/HandleTasks.md.
//
// A task may only write to the database it is handed. It has no file, no
// terminal and no clock beyond Deps.Now — which is what makes a task safe to
// run from a tick, from the command line, or from a test, with the same
// result.
//
// The database is the injected Keep library and nothing else: a task asks it
// for the schema it writes — the registries and their fields are the
// constants in sandbox/config — lists or looks up records on that schema, and
// inserts, updates or removes one.
package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
)

// The ground every task in this package shares: the words its fields are
// named with, and the handful of reads and writes more than one task repeats.
// Nothing that belongs to a single action lives here — a task file still
// carries the whole of what it alone does, so reading one file is reading
// that action.

// The field names the task guides use. They are constants because the same
// word names a key in Task.yaml, a `--flag` on the command line, and a column
// of a rendered page — so it is spelled once.
const (
	// AccountField names an account.
	AccountField = "account"
	// CategoryField names a category.
	CategoryField = "category"
	// DescriptionField is free text describing the record.
	DescriptionField = "description"
	// AmountField is a value, positive for income and negative for expense.
	AmountField = "amount"
	// DateField is the date a transaction counts on.
	DateField = "date"
	// IdField addresses an existing transaction.
	IdField = "id"
	// RevenuesField marks a category as accepting income.
	RevenuesField = "revenues"
	// ExpensesField marks a category as accepting expenses.
	ExpensesField = "expenses"
	// ParentField hangs a category under another one.
	ParentField = "parent"
)

// The two-way link between a movement and the registries it names, which is
// the one piece of behavior every transaction task repeats. A transaction
// stores the id of its account and of its category; the account and the
// category each store the movement's id in their own nested TransactionsDB.
// Keeping both sides of that link in one place is what makes "every
// transaction of this account" a read of one record, and what keeps the two
// sides from ever drifting apart.

// schemaOf returns one registry of the database a task was handed, failing in
// the same words for every task that cannot reach it.
func schemaOf(args api.HandleActionArgs, name string) (keepdeps.SchemaInstance, error) {
	instance, reachable := args.DataBase.GetSchema(name)
	if !reachable {
		return instance, errors.New("the " + name + " registry is unreachable")
	}
	return instance, nil
}

// linkTransaction writes a movement's id into an account's or a category's
// nested transaction registry.
func linkTransaction(owner keepdeps.SchemaItem, transactionId int64) error {
	_, failure := owner.NewSubItem(config.TransactionsDB, map[string]any{
		config.TransactionId: transactionId,
	})
	if failure != nil {
		return errors.New("the transaction could not be indexed on " +
			text(owner, config.NameField) + ": " + failure.Message)
	}
	return nil
}

// unlinkTransaction removes a movement's id from an account's or a category's
// nested transaction registry. A link that is not there is not an error:
// removing a movement is over when no registry points at it any more.
func unlinkTransaction(owner keepdeps.SchemaItem, transactionId int64) error {
	for _, link := range owner.ListAll(config.TransactionsDB) {
		if number(link, config.TransactionId) != transactionId {
			continue
		}
		if failure := link.Remove(); failure != nil {
			return errors.New("the transaction could not be unindexed from " +
				text(owner, config.NameField) + ": " + failure.Message)
		}
	}
	return nil
}

// transactionCount reports how many movements an account or a category still
// holds, read off its own nested registry rather than off the whole ledger.
func transactionCount(owner keepdeps.SchemaItem) int {
	return len(owner.ListAll(config.TransactionsDB))
}

// text reads a text field of a stored record, as "" when the field is absent
// or holds something else.
func text(record keepdeps.SchemaItem, field string) string {
	value, err := record.Get(field)
	if err != nil {
		return ""
	}
	stored, ok := value.(string)
	if !ok {
		return ""
	}
	return stored
}

// number reads a whole-number field of a stored record, as 0 when the field
// is absent or holds something else.
func number(record keepdeps.SchemaItem, field string) int64 {
	value, err := record.Get(field)
	if err != nil {
		return 0
	}
	switch stored := value.(type) {
	case int64:
		return stored
	case int:
		return int64(stored)
	}
	return 0
}
