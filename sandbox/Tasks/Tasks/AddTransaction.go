package tasks

import (
	"errors"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// KeyWidth is how many digits a transaction's storage key is padded to, so
// the keys of one registry sort the way they were written.
const KeyWidth = 6

// addTransactionFields declares what the task accepts.
func addTransactionFields() []api.Field {
	return []api.Field{
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "The account the movement happened on"},
		{Name: CategoryField, Type: api.TextField, Required: true,
			Description: "The category it is classified under"},
		{Name: AmountField, Type: api.NumberField, Required: true,
			Description: "Positive for money in, negative for money out"},
		{Name: DateField, Type: api.TextField, Required: true,
			Description: "The date it counts on, written as YYYY-MM-DD"},
		{Name: DescriptionField, Type: api.TextField,
			Description: "What it was"},
	}
}

// addTransactionAction reads the movement the task describes and writes it
// into the ledger: one record in the transaction registry carrying the ids of
// the account and the category it names, and one line in each of those two
// records' nested transaction registries pointing back at it.
func addTransactionAction(args api.HandleActionArgs) error {
	transaction, err := addTransactionRead(args)
	if err != nil {
		return err
	}
	ledger, err := schemaOf(args, config.TransactionSchema)
	if err != nil {
		return err
	}
	key := addTransactionNextKey(ledger)
	record, failure := ledger.NewItem(map[string]any{
		config.NameField: key,
		config.DetailField: utils.Pack(key, transaction.Account, transaction.Category,
			transaction.Description),
		config.AmountField: transaction.Amount,
		config.DateField:   transaction.Date,
		config.AccountID:   transaction.AccountRecord.Id,
		config.CategoryID:  transaction.CategoryRecord.Id,
	})
	if failure != nil {
		if failure.Type == keepdeps.KeyConflict {
			return errors.New("the transaction already exists")
		}
		return errors.New("the transaction could not be stored: " + failure.Message)
	}
	// The movement exists now, so both sides of its index have to exist too.
	// A link that cannot be written takes the movement back out with it: a
	// half-indexed ledger would hide the movement from the account it was
	// dated on.
	if err := addTransactionIndex(transaction, record.Id); err != nil {
		if failure := record.Remove(); failure != nil {
			return errors.New(err.Error() + ", and the transaction it was written for " +
				"could not be taken back out: " + failure.Message)
		}
		return err
	}
	return nil
}

// addTransactionIndex writes the movement's id into the account and the
// category it names, undoing the first link when the second one fails.
func addTransactionIndex(transaction addTransactionMovement, id int64) error {
	if err := linkTransaction(transaction.AccountRecord, id); err != nil {
		return err
	}
	if err := linkTransaction(transaction.CategoryRecord, id); err != nil {
		if undo := unlinkTransaction(transaction.AccountRecord, id); undo != nil {
			return errors.New(err.Error() + ", and " + undo.Error())
		}
		return err
	}
	return nil
}

// addTransactionMovement is one movement as the task composes it, before it
// is stored.
type addTransactionMovement struct {
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
	// AccountRecord is the stored account the movement happened on, which is
	// what its AccountID points at and where its id is indexed.
	AccountRecord keepdeps.SchemaItem
	// CategoryRecord is the stored category it is classified under, which is
	// what its CategoryID points at and where its id is indexed.
	CategoryRecord keepdeps.SchemaItem
}

// addTransactionRead validates every field of the task against the registries
// it names, and hands back the movement it describes.
func addTransactionRead(args api.HandleActionArgs) (addTransactionMovement, error) {
	empty := addTransactionMovement{}
	accounts, err := schemaOf(args, config.AccountSchema)
	if err != nil {
		return empty, err
	}
	accountName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return empty, err
	}
	if accountName == "" {
		return empty, errors.New(AccountField + " is required")
	}
	account, found := accounts.FindByKey(config.NameField, accountName)
	if !found {
		return empty, errors.New("account not found: " + accountName)
	}
	categories, err := schemaOf(args, config.CategorySchema)
	if err != nil {
		return empty, err
	}
	categoryName, err := entries.Text(args.Entries, CategoryField)
	if err != nil {
		return empty, err
	}
	if categoryName == "" {
		return empty, errors.New(CategoryField + " is required")
	}
	category, found := categories.FindByKey(config.NameField, categoryName)
	if !found {
		return empty, errors.New("category not found: " + categoryName)
	}
	cents, err := entries.Amount(args.Entries, AmountField)
	if err != nil {
		return empty, err
	}
	if cents == 0 {
		return empty, errors.New(AmountField + " may not be zero")
	}
	// Money may only arrive in a category that accepts income, and may only
	// leave through one that accepts expenses. A transfer category — neither
	// revenue nor expense — accepts both, because it counts as neither.
	revenues := number(category, config.RevenuesField)
	expenses := number(category, config.ExpensesField)
	transfer := revenues != 1 && expenses != 1
	if !transfer && cents > 0 && revenues != 1 {
		return empty, errors.New("a positive amount needs a category with revenues: true — " +
			categoryName + " does not accept income")
	}
	if !transfer && cents < 0 && expenses != 1 {
		return empty, errors.New("a negative amount needs a category with expenses: true — " +
			categoryName + " does not accept expenses")
	}
	dateText, err := entries.Text(args.Entries, DateField)
	if err != nil {
		return empty, err
	}
	when, parseErr := utils.ParseDate(dateText)
	if parseErr != nil {
		return empty, errors.New(DateField +
			" must be a date written as YYYY-MM-DD, not " + dateText)
	}
	description := ""
	if entries.Present(args.Entries, DescriptionField) {
		description, err = entries.Text(args.Entries, DescriptionField)
		if err != nil {
			return empty, err
		}
		if strings.Contains(description, utils.Separator) || strings.Contains(description, "\n") || strings.Contains(description, "\r") {
			return empty, errors.New(DescriptionField +
				" may not contain line breaks or " + utils.Separator)
		}
	}
	return addTransactionMovement{
		Account:        accountName,
		Category:       categoryName,
		Description:    description,
		Amount:         cents,
		Date:           when,
		AccountRecord:  account,
		CategoryRecord: category,
	}, nil
}

// addTransactionNextKey returns a storage key no record of the ledger carries
// yet, built from how many records it already holds. Keys are never reused,
// so a key already taken is stepped over rather than filled in.
func addTransactionNextKey(ledger keepdeps.SchemaInstance) string {
	stored, failure := ledger.ListAll()
	if failure != nil {
		stored = nil
	}
	candidate := int64(len(stored)) + 1
	for {
		key := utils.Pad(candidate, KeyWidth)
		if _, taken := ledger.FindByKey(config.NameField, key); !taken {
			return key
		}
		candidate++
	}
}

// AddTransaction returns the task that records a movement: money in, money
// out, or one leg of a transfer between two of your own accounts. It is the
// task you will run most, and the only one that writes to the ledger.
//
// A movement settles on the day it is dated. There is nothing to pay later:
// the amount leaves or reaches the account it names, on its date, in full.
func AddTransaction() api.Task {
	return api.Task{
		Name:         "AddTransaction",
		Description:  "Record an income, an expense, or one leg of a transfer",
		Fields:       addTransactionFields(),
		HandleAction: addTransactionAction,
	}
}
