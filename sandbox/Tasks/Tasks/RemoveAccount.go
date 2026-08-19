package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// removeAccountFields declares what the task accepts.
func removeAccountFields() []api.Field {
	return []api.Field{
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "Display name of the account to remove"},
	}
}

// removeAccountAction runs the task against the database it is handed.
func removeAccountAction(args api.HandleActionArgs) error {
	accountName, err := name(args.Entries, AccountField)
	if err != nil {
		return err
	}
	account, found := find(args, config.AccountSchema, accountName)
	if found && isCard(account) {
		return errors.New(accountName + " is a credit card — remove it with RemoveCreditCard")
	}
	if err := refuseWhenInUse(args, accountName); err != nil {
		return err
	}
	return remove(args, config.AccountSchema, "account", accountName)
}

// RemoveAccount returns the task that removes an account from the registry.
// It refuses an account transactions still name: a ledger whose movements
// point at an account that no longer exists cannot be rendered, and silently
// deleting those movements would change history. Remove them first, or leave
// the account where it is.
func RemoveAccount() api.Task {
	return api.Task{
		Name:         "RemoveAccount",
		Description:  "Remove an account the registry no longer needs",
		Fields:       removeAccountFields(),
		HandleAction: removeAccountAction,
	}
}

// refuseWhenInUse reports the transactions and recurrences still naming an
// account, so removing one never orphans a movement.
func refuseWhenInUse(args api.HandleActionArgs, accountName string) error {
	for _, transaction := range records(args, config.TransactionSchema) {
		if detail(transaction, config.TransactionParts, config.TransactionAccount) == accountName {
			return errors.New(accountName + " still holds transactions — " +
				"remove them before removing the account")
		}
	}
	for _, recurrence := range records(args, config.RecurrenceSchema) {
		named := detail(recurrence, config.RecurrenceParts, config.RecurrenceAccount) == accountName ||
			detail(recurrence, config.RecurrenceParts, config.RecurrenceToAccount) == accountName
		if named {
			return errors.New(accountName + " is still named by the recurrence " +
				text(recurrence, config.NameField) + " — remove it first")
		}
	}
	return nil
}
