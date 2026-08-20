package tasks

import (
	"errors"
	"strconv"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
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
	}
}

// modifyTransactionAction finds the stored movement by its id, overlays the
// fields the task was given onto it, and writes it back — moving its id
// between the nested registries of the accounts and categories it names when
// one of those two changes.
func modifyTransactionAction(args api.HandleActionArgs) error {
	id, err := entries.Whole(args.Entries, IdField)
	if err != nil {
		return err
	}
	ledger, err := schemaOf(args, config.TransactionSchema)
	if err != nil {
		return err
	}
	record, found := ledger.FindById(id)
	if !found {
		return errors.New("transaction not found: " + strconv.FormatInt(id, 10))
	}
	// The record's account, category and description are packed into one
	// detail key beside its own storage key; which account and which category
	// those names stand for is held as their ids.
	key := text(record, config.NameField)
	detail := text(record, config.DetailField)
	current := modifyTransactionMovement{
		Account:     utils.Part(detail, config.TransactionParts, config.TransactionAccount),
		Category:    utils.Part(detail, config.TransactionParts, config.TransactionCategory),
		Description: utils.Part(detail, config.TransactionParts, config.TransactionDescription),
		Amount:      number(record, config.AmountField),
		Date:        number(record, config.DateField),
		AccountId:   number(record, config.AccountID),
		CategoryId:  number(record, config.CategoryID),
	}
	updated, err := modifyTransactionApply(args, current)
	if err != nil {
		return err
	}
	fields := map[string]any{
		config.NameField: key,
		config.DetailField: utils.Pack(key, updated.Account, updated.Category,
			updated.Description),
		config.AmountField: updated.Amount,
		config.DateField:   updated.Date,
		config.AccountID:   updated.AccountId,
		config.CategoryID:  updated.CategoryId,
	}
	for field, value := range fields {
		if writeErr := record.Update(field, value); writeErr != nil {
			return errors.New("transaction " + strconv.FormatInt(id, 10) +
				" could not be updated: " + writeErr.Message)
		}
	}
	return modifyTransactionReindex(args, record.Id, current, updated)
}

// modifyTransactionReindex moves the movement's id out of the account and the
// category it used to point at and into the ones it points at now. A side
// that did not change is left alone, so a correction that only rewrites a
// description touches no index at all.
func modifyTransactionReindex(args api.HandleActionArgs, id int64,
	current modifyTransactionMovement, updated modifyTransactionMovement) error {
	moves := []struct {
		schemaName string
		was        int64
		now        int64
	}{
		{config.AccountSchema, current.AccountId, updated.AccountId},
		{config.CategorySchema, current.CategoryId, updated.CategoryId},
	}
	for _, move := range moves {
		if move.was == move.now {
			continue
		}
		instance, err := schemaOf(args, move.schemaName)
		if err != nil {
			return err
		}
		if previous, found := instance.FindById(move.was); found {
			if err := unlinkTransaction(previous, id); err != nil {
				return err
			}
		}
		owner, found := instance.FindById(move.now)
		if !found {
			return errors.New("the " + move.schemaName +
				" the transaction was moved to is no longer there")
		}
		if err := linkTransaction(owner, id); err != nil {
			return err
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
	// AccountId is the id of the account record the movement points at, and
	// the record its id is indexed in.
	AccountId int64
	// CategoryId is the id of the category record it points at, and the
	// record its id is indexed in.
	CategoryId int64
}

// modifyTransactionApply overlays the fields the task was given onto the
// stored movement, checking each one the way AddTransaction checks it.
func modifyTransactionApply(args api.HandleActionArgs,
	current modifyTransactionMovement) (modifyTransactionMovement, error) {
	updated := current
	accounts, err := schemaOf(args, config.AccountSchema)
	if err != nil {
		return updated, err
	}
	categories, err := schemaOf(args, config.CategorySchema)
	if err != nil {
		return updated, err
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
		account, found := accounts.FindByKey(config.NameField, accountName)
		if !found {
			return updated, errors.New("account not found: " + accountName)
		}
		updated.Account = accountName
		updated.AccountId = account.Id
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
		if strings.Contains(description, utils.Separator) || strings.Contains(description, "\n") || strings.Contains(description, "\r") {
			return updated, errors.New(DescriptionField + " may not contain line breaks or " + utils.Separator)
		}
		updated.Description = description
	}
	if entries.Present(args.Entries, AmountField) {
		cents, err := entries.Amount(args.Entries, AmountField)
		if err != nil {
			return updated, err
		}
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
		updated.Date = when
	}
	category, found := categories.FindByKey(config.NameField, updated.Category)
	if !found {
		return updated, errors.New("category not found: " + updated.Category)
	}
	updated.CategoryId = category.Id
	// The corrected movement still has to obey the sign its category accepts.
	revenues := number(category, config.RevenuesField)
	expenses := number(category, config.ExpensesField)
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

// ModifyTransaction returns the task that corrects a movement already in the
// ledger. Every field but `id` is optional: what you give is overwritten,
// what you leave out stays as it is. The id is the one a statement shows
// beside each line.
func ModifyTransaction() api.Task {
	return api.Task{
		Name:         "ModifyTransaction",
		Description:  "Correct a transaction already in the ledger",
		Fields:       modifyTransactionFields(),
		HandleAction: modifyTransactionAction,
	}
}
