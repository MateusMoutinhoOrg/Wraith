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

// Shared ground every task in this package stands on: the words its fields
// are named with, the checks each one repeats, and the few lines reading a
// stored record costs.

import (
	"errors"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
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

// KeyWidth is how many digits a transaction's storage key is padded to, so
// the keys of one registry sort the way they were written.
const KeyWidth = 6

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
	if strings.Contains(text, utils.Separator) {
		return "", errors.New(key + " may not contain " + utils.Separator)
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
	parsed, parseErr := utils.ParseDate(text)
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
	parsed, parseErr := utils.ParseMonth(text)
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
	return utils.Cents(amount), nil
}

// schema asks the database for one registry, reporting a registry the
// injected library hands nothing back for — which can only happen if
// config.DatabaseProps and the name asked for have drifted apart.
func schema(args api.HandleActionArgs, name string) (keepdeps.SchemaInstance, error) {
	instance, ok := args.DataBase.GetSchema(name)
	if !ok {
		return keepdeps.SchemaInstance{}, errors.New("the " + name + " registry is unreachable")
	}
	return instance, nil
}

// records returns every record of one registry, so a task can check what
// naming something would break. An unreachable registry reads as no records:
// a task that cannot see a registry has nothing to refuse over.
func records(args api.HandleActionArgs, name string) []keepdeps.SchemaItem {
	instance, err := schema(args, name)
	if err != nil {
		return nil
	}
	stored, failure := instance.ListAll()
	if failure != nil {
		return nil
	}
	return stored
}

// find returns the record of one registry carrying the given name.
func find(args api.HandleActionArgs, schemaName string, key string) (keepdeps.SchemaItem, bool) {
	instance, err := schema(args, schemaName)
	if err != nil {
		return keepdeps.SchemaItem{}, false
	}
	return instance.FindByKey(config.NameField, key)
}

// insert writes one record into a registry, reporting a name already taken in
// the words the guides use.
func insert(args api.HandleActionArgs, schemaName string, subject string, fields map[string]any) error {
	instance, err := schema(args, schemaName)
	if err != nil {
		return err
	}
	_, failure := instance.NewItem(fields)
	if failure == nil {
		return nil
	}
	if failure.Type == keepdeps.KeyConflict {
		return errors.New(subject + " already exists")
	}
	return errors.New(subject + " could not be stored: " + failure.Message)
}

// remove deletes one record of a registry by name, reporting a name that is
// not there.
func remove(args api.HandleActionArgs, schemaName string, subject string, key string) error {
	record, found := find(args, schemaName, key)
	if !found {
		return errors.New(subject + " not found: " + key)
	}
	if failure := record.Remove(); failure != nil {
		return errors.New(subject + " could not be removed: " + failure.Message)
	}
	return nil
}

// nextKey returns a storage key no record of the registry carries yet, built
// from how many records it already holds. Keys are never reused, so a key
// already taken is stepped over rather than filled in.
func nextKey(instance keepdeps.SchemaInstance) string {
	stored, failure := instance.ListAll()
	if failure != nil {
		stored = nil
	}
	candidate := int64(len(stored)) + 1
	for {
		key := utils.Pad(candidate, KeyWidth)
		if _, taken := instance.FindByKey(config.NameField, key); !taken {
			return key
		}
		candidate++
	}
}

// requireAccount returns the stored record of one account, reporting an
// account the registry does not carry.
func requireAccount(args api.HandleActionArgs, accountName string) (keepdeps.SchemaItem, error) {
	record, found := find(args, config.AccountSchema, accountName)
	if !found {
		return keepdeps.SchemaItem{}, errors.New("account not found: " + accountName)
	}
	return record, nil
}

// requireCategory returns the stored record of one category, reporting a
// category the registry does not carry.
func requireCategory(args api.HandleActionArgs, categoryName string) (keepdeps.SchemaItem, error) {
	record, found := find(args, config.CategorySchema, categoryName)
	if !found {
		return keepdeps.SchemaItem{}, errors.New("category not found: " + categoryName)
	}
	return record, nil
}

// isCard reports whether a stored account record is a credit card.
func isCard(account keepdeps.SchemaItem) bool {
	return number(account, config.KindField) == config.KindCard
}

// isTransfer reports whether a stored category record is a transfer category
// — one that is neither revenue nor expense, and is therefore how a movement
// between two of your own accounts is recorded.
func isTransfer(category keepdeps.SchemaItem) bool {
	return number(category, config.RevenuesField) != 1 &&
		number(category, config.ExpensesField) != 1
}

// checkSign enforces the one rule that keeps a ledger honest: money may only
// arrive in a category that accepts income, and may only leave through one
// that accepts expenses. A transfer category — neither revenue nor expense —
// accepts both, because it counts as neither.
func checkSign(category keepdeps.SchemaItem, amount int64) error {
	if isTransfer(category) {
		return nil
	}
	categoryName := text(category, config.NameField)
	if amount > 0 && number(category, config.RevenuesField) != 1 {
		return errors.New("a positive amount needs a category with revenues: true — " +
			categoryName + " does not accept income")
	}
	if amount < 0 && number(category, config.ExpensesField) != 1 {
		return errors.New("a negative amount needs a category with expenses: true — " +
			categoryName + " does not accept expenses")
	}
	return nil
}

// text reads a Key field of a stored record, returning "" when the field is
// absent or does not hold a string.
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

// number reads an Int field of a stored record, returning 0 when the field is
// absent or does not hold a whole number.
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

// detail reads one part of a record's packed detail key — the positions are
// the constants in sandbox/config, one set per registry.
func detail(record keepdeps.SchemaItem, count int, index int) string {
	return utils.Part(text(record, config.DetailField), count, index)
}

// movement is one transaction as a task handles it, before it is stored and
// after it is read back. It is the shape AddTransaction composes, spread over
// installments and written, and the shape ModifyTransaction overlays its
// changes onto.
type movement struct {
	// Key is the record's unique storage key.
	Key string
	// Account is the account the movement happened on.
	Account string
	// Category is the category it is classified under.
	Category string
	// Description is the free text it carries.
	Description string
	// Amount is the value in cents — positive for income, negative for an
	// expense.
	Amount int64
	// Date is the date it counts on, as yyyymmdd.
	Date int64
	// PaymentDate is the date the money actually moves, as yyyymmdd.
	PaymentDate int64
}

// readMovement reads one stored record as a movement.
func readMovement(record keepdeps.SchemaItem) movement {
	return movement{
		Key:         text(record, config.NameField),
		Account:     detail(record, config.TransactionParts, config.TransactionAccount),
		Category:    detail(record, config.TransactionParts, config.TransactionCategory),
		Description: detail(record, config.TransactionParts, config.TransactionDescription),
		Amount:      number(record, config.AmountField),
		Date:        number(record, config.DateField),
		PaymentDate: number(record, config.PaymentDateField),
	}
}

// movementFields composes the stored fields of one movement, packing its
// account, category and description into the detail key beside its own unique
// storage key.
func movementFields(key string, m movement) map[string]any {
	return map[string]any{
		config.NameField:        key,
		config.DetailField:      utils.Pack(key, m.Account, m.Category, m.Description),
		config.AmountField:      m.Amount,
		config.DateField:        m.Date,
		config.PaymentDateField: m.PaymentDate,
	}
}
