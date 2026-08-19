package tasks

import (
	"errors"
	"strconv"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// modifyTransactionFields declares what the task accepts.
func modifyTransactionFields() []api.Field {
	return []api.Field{
		{Name: IdField, Type: api.NumberField, Required: true,
			Description: "The id of the transaction, as the statement shows it"},
		{Name: AccountField, Type: api.TextField,
			Description: "Move it to another account"},
		{Name: CategoryField, Type: api.TextField,
			Description: "Classify it under another category"},
		{Name: DescriptionField, Type: api.TextField,
			Description: "Rewrite what it was"},
		{Name: AmountField, Type: api.NumberField,
			Description: "Correct the amount"},
		{Name: DateField, Type: api.TextField,
			Description: "Correct the date it counts on, as YYYY-MM-DD"},
		{Name: PaymentDateField, Type: api.TextField,
			Description: "Correct when the money actually moves, as YYYY-MM-DD"},
	}
}

// modifyTransactionAction finds the stored movement by its id, overlays the
// fields the task was given onto it, and writes it back.
func modifyTransactionAction(args api.HandleActionArgs) error {
	id, err := entries.Whole(args.Entries, IdField)
	if err != nil {
		return err
	}
	ledger, reachable := args.DataBase.GetSchema(config.TransactionSchema)
	if !reachable {
		return errors.New("the " + config.TransactionSchema + " registry is unreachable")
	}
	stored, failure := ledger.ListAll()
	if failure != nil {
		stored = nil
	}
	record := keepdeps.SchemaItem{}
	found := false
	for _, candidate := range stored {
		if candidate.Id == id {
			record = candidate
			found = true
			break
		}
	}
	if !found {
		return errors.New("transaction not found: " + strconv.FormatInt(id, 10))
	}
	// The record's account, category and description are packed into one
	// detail key beside its own storage key.
	key := modifyTransactionText(record, config.NameField)
	detail := modifyTransactionText(record, config.DetailField)
	current := modifyTransactionMovement{
		Account:     utils.Part(detail, config.TransactionParts, config.TransactionAccount),
		Category:    utils.Part(detail, config.TransactionParts, config.TransactionCategory),
		Description: utils.Part(detail, config.TransactionParts, config.TransactionDescription),
		Amount:      modifyTransactionNumber(record, config.AmountField),
		Date:        modifyTransactionNumber(record, config.DateField),
		PaymentDate: modifyTransactionNumber(record, config.PaymentDateField),
	}
	updated, err := modifyTransactionApply(args, current)
	if err != nil {
		return err
	}
	fields := map[string]any{
		config.NameField: key,
		config.DetailField: utils.Pack(key, updated.Account, updated.Category,
			updated.Description),
		config.AmountField:      updated.Amount,
		config.DateField:        updated.Date,
		config.PaymentDateField: updated.PaymentDate,
	}
	for field, value := range fields {
		if writeErr := record.Update(field, value); writeErr != nil {
			return errors.New("transaction " + strconv.FormatInt(id, 10) +
				" could not be updated: " + writeErr.Message)
		}
	}
	return nil
}

// modifyTransactionMovement is one stored movement as the task reads it and
// writes it back.
type modifyTransactionMovement struct {
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

// modifyTransactionApply overlays the fields the task was given onto the
// stored movement, checking each one the way AddTransaction checks it.
func modifyTransactionApply(args api.HandleActionArgs,
	current modifyTransactionMovement) (modifyTransactionMovement, error) {
	updated := current
	accounts, reachable := args.DataBase.GetSchema(config.AccountSchema)
	if !reachable {
		return updated, errors.New("the " + config.AccountSchema + " registry is unreachable")
	}
	categories, reachable := args.DataBase.GetSchema(config.CategorySchema)
	if !reachable {
		return updated, errors.New("the " + config.CategorySchema + " registry is unreachable")
	}
	if entries.Present(args.Entries, AccountField) {
		accountName, err := entries.Text(args.Entries, AccountField)
		if err != nil {
			return updated, err
		}
		if accountName == "" {
			return updated, errors.New(AccountField + " is required")
		}
		if strings.Contains(accountName, utils.Separator) {
			return updated, errors.New(AccountField + " may not contain " + utils.Separator)
		}
		if _, found := accounts.FindByKey(config.NameField, accountName); !found {
			return updated, errors.New("account not found: " + accountName)
		}
		updated.Account = accountName
	}
	if entries.Present(args.Entries, CategoryField) {
		categoryName, err := entries.Text(args.Entries, CategoryField)
		if err != nil {
			return updated, err
		}
		if categoryName == "" {
			return updated, errors.New(CategoryField + " is required")
		}
		if strings.Contains(categoryName, utils.Separator) {
			return updated, errors.New(CategoryField + " may not contain " + utils.Separator)
		}
		if _, found := categories.FindByKey(config.NameField, categoryName); !found {
			return updated, errors.New("category not found: " + categoryName)
		}
		updated.Category = categoryName
	}
	if entries.Present(args.Entries, DescriptionField) {
		description, err := entries.Text(args.Entries, DescriptionField)
		if err != nil {
			return updated, err
		}
		if strings.Contains(description, utils.Separator) {
			return updated, errors.New(DescriptionField + " may not contain " + utils.Separator)
		}
		updated.Description = description
	}
	if entries.Present(args.Entries, AmountField) {
		amount, err := entries.Number(args.Entries, AmountField)
		if err != nil {
			return updated, err
		}
		cents := utils.Cents(amount)
		if cents == 0 {
			return updated, errors.New(AmountField + " may not be zero")
		}
		updated.Amount = cents
	}
	if entries.Present(args.Entries, DateField) {
		dateText, err := entries.Text(args.Entries, DateField)
		if err != nil {
			return updated, err
		}
		when, parseErr := utils.ParseDate(dateText)
		if parseErr != nil {
			return updated, errors.New(DateField +
				" must be a date written as YYYY-MM-DD, not " + dateText)
		}
		// A payment date that was never given a life of its own follows the
		// date it was copied from.
		if updated.PaymentDate == current.Date {
			updated.PaymentDate = when
		}
		updated.Date = when
	}
	if entries.Present(args.Entries, PaymentDateField) {
		paymentText, err := entries.Text(args.Entries, PaymentDateField)
		if err != nil {
			return updated, err
		}
		payment, parseErr := utils.ParseDate(paymentText)
		if parseErr != nil {
			return updated, errors.New(PaymentDateField +
				" must be a date written as YYYY-MM-DD, not " + paymentText)
		}
		updated.PaymentDate = payment
	}
	category, found := categories.FindByKey(config.NameField, updated.Category)
	if !found {
		return updated, errors.New("category not found: " + updated.Category)
	}
	// The corrected movement still has to obey the sign its category accepts.
	revenues := modifyTransactionNumber(category, config.RevenuesField)
	expenses := modifyTransactionNumber(category, config.ExpensesField)
	transfer := revenues != 1 && expenses != 1
	if !transfer && updated.Amount > 0 && revenues != 1 {
		return updated, errors.New("a positive amount needs a category with revenues: true — " +
			updated.Category + " does not accept income")
	}
	if !transfer && updated.Amount < 0 && expenses != 1 {
		return updated, errors.New("a negative amount needs a category with expenses: true — " +
			updated.Category + " does not accept expenses")
	}
	return updated, nil
}

// modifyTransactionText reads a text field of a stored record, as "" when the
// field is absent or holds something else.
func modifyTransactionText(record keepdeps.SchemaItem, field string) string {
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

// modifyTransactionNumber reads a whole-number field of a stored record, as 0
// when the field is absent or holds something else.
func modifyTransactionNumber(record keepdeps.SchemaItem, field string) int64 {
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

// ModifyTransaction returns the task that corrects a movement already in the
// ledger. Every field but `id` is optional: what you give is overwritten,
// what you leave out stays as it is. The id is the one a statement shows
// beside each line.
//
// One part of a split purchase is an ordinary transaction like any other, so
// this is how a single part is corrected — the other parts are untouched.
func ModifyTransaction() api.Task {
	return api.Task{
		Name:         "ModifyTransaction",
		Description:  "Correct a transaction already in the ledger",
		Fields:       modifyTransactionFields(),
		HandleAction: modifyTransactionAction,
	}
}
