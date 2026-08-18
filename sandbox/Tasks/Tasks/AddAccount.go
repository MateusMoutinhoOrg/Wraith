package tasks

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// AddAccount returns the task that adds an account to the registry — a bank,
// a wallet, a savings pot. An account is where money sits; every transaction
// names one. A credit card is added with AddCreditCard instead, because it
// carries a limit and a billing cycle an account does not.
func AddAccount() api.Task {
	return api.Task{
		Name:        "AddAccount",
		Description: "Add an account — somewhere money sits",
		Fields: []api.Field{
			{Name: AccountField, Type: api.TextField, Required: true,
				Description: "Display name, e.g. Vacation savings"},
			{Name: OpeningField, Type: api.NumberField, Required: true,
				Description: "Opening balance, the amount the account holds today"},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			accountName, err := name(args.Entries, AccountField)
			if err != nil {
				return err
			}
			opening, err := cents(args.Entries, OpeningField)
			if err != nil {
				return err
			}
			return insert(args, store.AccountSchema, "account "+accountName, map[string]any{
				store.NameField:    accountName,
				store.KindField:    int64(store.KindAccount),
				store.OpeningField: opening,
			})
		},
	}
}
