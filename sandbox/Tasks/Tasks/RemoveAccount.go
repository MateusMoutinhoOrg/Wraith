package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
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
	accounts, reachable := args.DataBase.GetSchema(config.AccountSchema)
	if !reachable {
		return errors.New("the " + config.AccountSchema + " registry is unreachable")
	}
	account, found := accounts.FindByKey(config.NameField, accountName)
	if !found {
		return errors.New("account not found: " + accountName)
	}
	// A ledger whose movements point at an account that no longer exists
	// cannot be rendered, so an account anything still names is refused.
	for _, transaction := range removeAccountRecords(args, config.TransactionSchema) {
		detail := removeAccountText(transaction, config.DetailField)
		if utils.Part(detail, config.TransactionParts, config.TransactionAccount) == accountName {
			return errors.New(accountName + " still holds transactions — " +
				"remove them before removing the account")
		}
	}
	if failure := account.Remove(); failure != nil {
		return errors.New("account could not be removed: " + failure.Message)
	}
	return nil
}

// removeAccountRecords returns every record of one registry. An unreachable
// registry reads as no records: a task that cannot see one has nothing to
// refuse over.
func removeAccountRecords(args api.HandleActionArgs, schemaName string) []keepdeps.SchemaItem {
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

// removeAccountText reads a text field of a stored record, as "" when the
// field is absent or holds something else.
func removeAccountText(record keepdeps.SchemaItem, field string) string {
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
