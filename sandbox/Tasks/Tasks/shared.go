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

// The only ground every task in this package shares: the words its fields are
// named with. There is no shared behavior here on purpose — each task file
// carries the whole of what it does, so reading one file is reading the whole
// action, and changing one task can never change another.

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
