package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
)

// removeAccountFields declares what the task accepts.
func removeAccountFields() []api.Field {
	return []api.Field{
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "Display name of the account to remove"},
	}
}

// removeAccountAction refuses an account anything still names, and otherwise
// deletes the record.
func removeAccountAction(args api.HandleActionArgs) error {
	accountName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return err
	}
	if accountName == "" {
		return errors.New(AccountField + " is required")
	}
	accounts, err := schemaOf(args, config.AccountSchema)
	if err != nil {
		return err
	}
	account, found := accounts.FindByKey(config.NameField, accountName)
	if !found {
		return errors.New("account not found: " + accountName)
	}
	// A ledger whose movements point at an account that no longer exists
	// cannot be rendered, so an account anything still names is refused. What
	// still names it is read off the account's own nested registry, which is
	// exactly the movements that were indexed onto it.
	if transactionCount(account) > 0 {
		return errors.New(accountName + " still holds transactions — " +
			"remove them before removing the account")
	}
	if failure := account.Remove(); failure != nil {
		return errors.New("account could not be removed: " + failure.Message)
	}
	return nil
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
