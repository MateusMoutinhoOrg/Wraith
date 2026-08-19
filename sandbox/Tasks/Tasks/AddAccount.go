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

// addAccountFields declares what the task accepts.
func addAccountFields() []api.Field {
	return []api.Field{
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "Display name, e.g. Vacation savings"},
		{Name: OpeningField, Type: api.NumberField, Required: true,
			Description: "Opening balance, the amount the account holds today"},
	}
}

// addAccountAction reads the two fields, then writes one record into the
// account registry of the database it was handed.
func addAccountAction(args api.HandleActionArgs) error {
	accountName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return err
	}
	if accountName == "" {
		return errors.New(AccountField + " is required")
	}
	// The storage keys are packed with this one character, so a name may not
	// carry it — a name can then always be read back the way it was written.
	if strings.Contains(accountName, utils.Separator) {
		return errors.New(AccountField + " may not contain " + utils.Separator)
	}
	amount, err := entries.Number(args.Entries, OpeningField)
	if err != nil {
		return err
	}
	accounts, reachable := args.DataBase.GetSchema(config.AccountSchema)
	if !reachable {
		return errors.New("the " + config.AccountSchema + " registry is unreachable")
	}
	_, failure := accounts.NewItem(map[string]any{
		config.NameField:    accountName,
		config.KindField:    int64(config.KindAccount),
		config.OpeningField: utils.Cents(amount),
	})
	if failure == nil {
		return nil
	}
	if failure.Type == keepdeps.KeyConflict {
		return errors.New("account " + accountName + " already exists")
	}
	return errors.New("account " + accountName + " could not be stored: " + failure.Message)
}

// AddAccount returns the task that adds an account to the registry — a bank,
// a wallet, a savings pot. An account is where money sits; every transaction
// names one. A credit card is added with AddCreditCard instead, because it
// carries a limit and a billing cycle an account does not.
func AddAccount() api.Task {
	return api.Task{
		Name:         "AddAccount",
		Description:  "Add an account — somewhere money sits",
		Fields:       addAccountFields(),
		HandleAction: addAccountAction,
	}
}
