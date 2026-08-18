package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// RemoveAccount returns the task that removes an account from the registry.
// It refuses an account transactions still name: a ledger whose movements
// point at an account that no longer exists cannot be rendered, and silently
// deleting those movements would change history. Remove them first, or leave
// the account where it is.
func RemoveAccount() api.Task {
	return api.Task{
		Name:        "RemoveAccount",
		Description: "Remove an account the registry no longer needs",
		Fields: []api.Field{
			{Name: AccountField, Type: api.TextField, Required: true,
				Description: "Display name of the account to remove"},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			accountName, err := name(args.Entries, AccountField)
			if err != nil {
				return err
			}
			account, found := store.FindAccount(args.DataBase, accountName)
			if found && account.IsCard() {
				return errors.New(accountName + " is a credit card — remove it with RemoveCreditCard")
			}
			if err := refuseWhenInUse(args, accountName); err != nil {
				return err
			}
			return remove(args, store.AccountSchema, "account", accountName)
		},
	}
}

// refuseWhenInUse reports the transactions and recurrences still naming an
// account, so removing one never orphans a movement.
func refuseWhenInUse(args api.HandleActionArgs, accountName string) error {
	for _, transaction := range store.Transactions(args.DataBase) {
		if transaction.Account == accountName {
			return errors.New(accountName + " still holds transactions — " +
				"remove them before removing the account")
		}
	}
	for _, recurrence := range store.Recurrences(args.DataBase) {
		if recurrence.Account == accountName || recurrence.ToAccount == accountName {
			return errors.New(accountName + " is still named by the recurrence " +
				recurrence.Description + " — remove it first")
		}
	}
	return nil
}
