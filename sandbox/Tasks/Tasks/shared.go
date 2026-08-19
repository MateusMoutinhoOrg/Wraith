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
