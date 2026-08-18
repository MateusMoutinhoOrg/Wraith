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
package tasks

// Shared ground every task in this package stands on: the words its fields
// are named with, and the checks each one repeats.

import (
	"errors"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// The field names the task guides use. They are constants because the same
// word names a key in Task.yaml, a `--flag` on the command line, and a column
// of a rendered page — so it is spelled once.
const (
	// AccountField names an account or a credit card.
	AccountField = "account"
	// ToAccountField names the destination account of a recurring transfer.
	ToAccountField = "to_account"
	// CategoryField names a category.
	CategoryField = "category"
	// RecurrenceField names a recurrence, by its description.
	RecurrenceField = "recurrence"
	// DescriptionField is free text describing the record.
	DescriptionField = "description"
	// AmountField is a value, positive for income and negative for expense.
	AmountField = "amount"
	// DateField is the date a transaction counts on.
	DateField = "date"
	// PaymentDateField is the date the money actually moves.
	PaymentDateField = "payment_date"
	// InstallmentsField splits one purchase into monthly parts.
	InstallmentsField = "installments"
	// IdField addresses an existing transaction.
	IdField = "id"
	// OpeningField is the balance an account starts at.
	OpeningField = "opening"
	// LimitField is a credit card's total limit.
	LimitField = "limit"
	// ClosingDayField is the day of the month a card's bill closes.
	ClosingDayField = "closing_day"
	// DueDayField is the day of the month a card's bill is due.
	DueDayField = "due_day"
	// RevenuesField marks a category as accepting income.
	RevenuesField = "revenues"
	// ExpensesField marks a category as accepting expenses.
	ExpensesField = "expenses"
	// ParentField hangs a category under another one.
	ParentField = "parent"
	// DayField is the day of the month a recurrence falls on.
	DayField = "day"
	// StartField is the first month a recurrence applies.
	StartField = "start"
	// EndField is the last month a recurrence applies.
	EndField = "end"
)

// name reads a required text field and refuses the one character the storage
// keys are packed with, so a name can always be read back the way it was
// written.
func name(values map[string]any, key string) (string, error) {
	text, err := entries.Text(values, key)
	if err != nil {
		return "", err
	}
	if text == "" {
		return "", errors.New(key + " is required")
	}
	if strings.Contains(text, store.Separator) {
		return "", errors.New(key + " may not contain " + store.Separator)
	}
	return text, nil
}

// optionalText reads a text field that may be absent, refusing the packing
// character the same way.
func optionalText(values map[string]any, key string) (string, error) {
	if !entries.Present(values, key) {
		return "", nil
	}
	return name(values, key)
}

// date reads a date field written as YYYY-MM-DD.
func date(values map[string]any, key string) (int64, error) {
	text, err := entries.Text(values, key)
	if err != nil {
		return 0, err
	}
	parsed, parseErr := store.ParseDate(text)
	if parseErr != nil {
		return 0, errors.New(key + " must be a date written as YYYY-MM-DD, not " + text)
	}
	return parsed, nil
}

// month reads a month field written as YYYY-MM.
func month(values map[string]any, key string) (int64, error) {
	text, err := entries.Text(values, key)
	if err != nil {
		return 0, err
	}
	parsed, parseErr := store.ParseMonth(text)
	if parseErr != nil {
		return 0, errors.New(key + " must be a month written as YYYY-MM, not " + text)
	}
	return parsed, nil
}

// dayOfMonth reads a field holding a day of the month, 1 to 31.
func dayOfMonth(values map[string]any, key string) (int64, error) {
	day, err := entries.Whole(values, key)
	if err != nil {
		return 0, err
	}
	if day < 1 || day > 31 {
		return 0, errors.New(key + " must be a day of the month, from 1 to 31")
	}
	return day, nil
}

// cents reads an amount field into the whole number of cents it stands for.
func cents(values map[string]any, key string) (int64, error) {
	amount, err := entries.Number(values, key)
	if err != nil {
		return 0, err
	}
	return store.Cents(amount), nil
}

// insert writes one record into a registry, reporting a name already taken
// in the words the guides use.
func insert(args api.HandleActionArgs, schema string, subject string, fields map[string]any) error {
	instance, ok := store.Schema(args.DataBase, schema)
	if !ok {
		return errors.New("the " + schema + " registry is unreachable")
	}
	_, err := instance.NewItem(fields)
	if err == nil {
		return nil
	}
	if err.Type == keepdeps.KeyConflict {
		return errors.New(subject + " already exists")
	}
	return errors.New(subject + " could not be stored: " + err.Message)
}

// remove deletes one record of a registry by name, reporting a name that is
// not there.
func remove(args api.HandleActionArgs, schema string, subject string, key string) error {
	record, found := store.Find(args.DataBase, schema, key)
	if !found {
		return errors.New(subject + " not found: " + key)
	}
	if err := record.Remove(); err != nil {
		return errors.New(subject + " could not be removed: " + err.Message)
	}
	return nil
}

// requireAccount reports an account the registry does not carry.
func requireAccount(args api.HandleActionArgs, accountName string) (store.Account, error) {
	account, found := store.FindAccount(args.DataBase, accountName)
	if !found {
		return store.Account{}, errors.New("account not found: " + accountName)
	}
	return account, nil
}

// requireCategory reports a category the registry does not carry.
func requireCategory(args api.HandleActionArgs, categoryName string) (store.Category, error) {
	category, found := store.FindCategory(args.DataBase, categoryName)
	if !found {
		return store.Category{}, errors.New("category not found: " + categoryName)
	}
	return category, nil
}

// checkSign enforces the one rule that keeps a ledger honest: money may only
// arrive in a category that accepts income, and may only leave through one
// that accepts expenses. A transfer category — neither revenue nor expense —
// accepts both, because it counts as neither.
func checkSign(category store.Category, amount int64) error {
	if category.IsTransfer() {
		return nil
	}
	if amount > 0 && !category.Revenues {
		return errors.New("a positive amount needs a category with revenues: true — " +
			category.Name + " does not accept income")
	}
	if amount < 0 && !category.Expenses {
		return errors.New("a negative amount needs a category with expenses: true — " +
			category.Name + " does not accept expenses")
	}
	return nil
}
