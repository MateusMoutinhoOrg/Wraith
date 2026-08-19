package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// RemoveCreditCard returns the task that removes a credit card from the
// registry. Like RemoveAccount, it refuses a card that still carries
// purchases: the ledger is history, and history is not deleted by removing
// the card it happened on.
func RemoveCreditCard() api.Task {
	return api.Task{
		Name:        "RemoveCreditCard",
		Description: "Remove a credit card the registry no longer needs",
		Fields: []api.Field{
			{Name: AccountField, Type: api.TextField, Required: true,
				Description: "Display name of the card to remove"},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			cardName, err := name(args.Entries, AccountField)
			if err != nil {
				return err
			}
			card, found := find(args, config.AccountSchema, cardName)
			if !found {
				return errors.New("credit card not found: " + cardName)
			}
			if !isCard(card) {
				return errors.New(cardName + " is an account, not a credit card — " +
					"remove it with RemoveAccount")
			}
			if err := refuseWhenInUse(args, cardName); err != nil {
				return err
			}
			return remove(args, config.AccountSchema, "credit card", cardName)
		},
	}
}
