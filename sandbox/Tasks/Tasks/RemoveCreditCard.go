package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// removeCreditCardFields declares what the task accepts.
func removeCreditCardFields() []api.Field {
	return []api.Field{
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "Display name of the card to remove"},
	}
}

// removeCreditCardAction refuses a plain account, refuses a card anything
// still names, and otherwise deletes the record.
func removeCreditCardAction(args api.HandleActionArgs) error {
	cardName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return err
	}
	if cardName == "" {
		return errors.New(AccountField + " is required")
	}
	accounts, reachable := args.DataBase.GetSchema(config.AccountSchema)
	if !reachable {
		return errors.New("the " + config.AccountSchema + " registry is unreachable")
	}
	card, found := accounts.FindByKey(config.NameField, cardName)
	if !found {
		return errors.New("credit card not found: " + cardName)
	}
	if removeCreditCardNumber(card, config.KindField) != config.KindCard {
		return errors.New(cardName + " is an account, not a credit card — " +
			"remove it with RemoveAccount")
	}
	// The ledger is history, and history is not deleted by removing the card
	// it happened on.
	for _, transaction := range removeCreditCardRecords(args, config.TransactionSchema) {
		detail := removeCreditCardText(transaction, config.DetailField)
		if utils.Part(detail, config.TransactionParts, config.TransactionAccount) == cardName {
			return errors.New(cardName + " still holds transactions — " +
				"remove them before removing the account")
		}
	}
	for _, recurrence := range removeCreditCardRecords(args, config.RecurrenceSchema) {
		detail := removeCreditCardText(recurrence, config.DetailField)
		named := utils.Part(detail, config.RecurrenceParts, config.RecurrenceAccount) == cardName ||
			utils.Part(detail, config.RecurrenceParts, config.RecurrenceToAccount) == cardName
		if named {
			return errors.New(cardName + " is still named by the recurrence " +
				removeCreditCardText(recurrence, config.NameField) + " — remove it first")
		}
	}
	if failure := card.Remove(); failure != nil {
		return errors.New("credit card could not be removed: " + failure.Message)
	}
	return nil
}

// removeCreditCardRecords returns every record of one registry. An
// unreachable registry reads as no records: a task that cannot see one has
// nothing to refuse over.
func removeCreditCardRecords(args api.HandleActionArgs, schemaName string) []keepdeps.SchemaItem {
	instance, reachable := args.DataBase.GetSchema(schemaName)
	if !reachable {
		return nil
	}
	stored, failure := instance.ListAll()
	if failure != nil {
		return nil
	}
	return stored
}

// removeCreditCardText reads a text field of a stored record, as "" when the
// field is absent or holds something else.
func removeCreditCardText(record keepdeps.SchemaItem, field string) string {
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

// removeCreditCardNumber reads a whole-number field of a stored record, as 0
// when the field is absent or holds something else.
func removeCreditCardNumber(record keepdeps.SchemaItem, field string) int64 {
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

// RemoveCreditCard returns the task that removes a credit card from the
// registry. Like RemoveAccount, it refuses a card that still carries
// purchases: the ledger is history, and history is not deleted by removing
// the card it happened on.
func RemoveCreditCard() api.Task {
	return api.Task{
		Name:         "RemoveCreditCard",
		Description:  "Remove a credit card the registry no longer needs",
		Fields:       removeCreditCardFields(),
		HandleAction: removeCreditCardAction,
	}
}
