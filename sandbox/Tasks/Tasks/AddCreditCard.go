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

// addCreditCardFields declares what the task accepts.
func addCreditCardFields() []api.Field {
	return []api.Field{
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "Display name, e.g. Nubank Card"},
		{Name: LimitField, Type: api.NumberField, Required: true,
			Description: "Total credit limit"},
		{Name: ClosingDayField, Type: api.NumberField, Required: true,
			Description: "Day of the month the bill closes (1-31)"},
		{Name: DueDayField, Type: api.NumberField, Required: true,
			Description: "Day of the month the bill is due (1-31)"},
		{Name: OpeningField, Type: api.NumberField,
			Description: "Amount already owed, written as a negative number",
			Default:     float64(0)},
	}
}

// addCreditCardAction reads every field, checks the limit and the billing
// days make sense, and writes one record into the account registry.
func addCreditCardAction(args api.HandleActionArgs) error {
	cardName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return err
	}
	if cardName == "" {
		return errors.New(AccountField + " is required")
	}
	if strings.Contains(cardName, utils.Separator) {
		return errors.New(AccountField + " may not contain " + utils.Separator)
	}
	limitAmount, err := entries.Number(args.Entries, LimitField)
	if err != nil {
		return err
	}
	limit := utils.Cents(limitAmount)
	if limit < 0 {
		return errors.New(LimitField + " is a credit limit, so it may not be negative")
	}
	closingDay, err := entries.Whole(args.Entries, ClosingDayField)
	if err != nil {
		return err
	}
	if closingDay < 1 || closingDay > 31 {
		return errors.New(ClosingDayField + " must be a day of the month, from 1 to 31")
	}
	dueDay, err := entries.Whole(args.Entries, DueDayField)
	if err != nil {
		return err
	}
	if dueDay < 1 || dueDay > 31 {
		return errors.New(DueDayField + " must be a day of the month, from 1 to 31")
	}
	openingAmount, err := entries.Number(args.Entries, OpeningField)
	if err != nil {
		return err
	}
	opening := utils.Cents(openingAmount)
	if opening > 0 {
		return errors.New(OpeningField + " on a card is what you already owe, " +
			"so it is written as a negative number")
	}
	accounts, reachable := args.DataBase.GetSchema(config.AccountSchema)
	if !reachable {
		return errors.New("the " + config.AccountSchema + " registry is unreachable")
	}
	_, failure := accounts.NewItem(map[string]any{
		config.NameField:       cardName,
		config.KindField:       int64(config.KindCard),
		config.OpeningField:    opening,
		config.LimitField:      limit,
		config.ClosingDayField: closingDay,
		config.DueDayField:     dueDay,
	})
	if failure == nil {
		return nil
	}
	if failure.Type == keepdeps.KeyConflict {
		return errors.New("credit card " + cardName + " already exists")
	}
	return errors.New("credit card " + cardName + " could not be stored: " + failure.Message)
}

// AddCreditCard returns the task that adds a credit card to the registry. A
// card is an account with a limit and a billing cycle: a purchase counts on
// the day it happens, and the money only leaves your bank when you record the
// bill payment as a transfer between the two.
func AddCreditCard() api.Task {
	return api.Task{
		Name:         "AddCreditCard",
		Description:  "Add a credit card — an account with a limit and a bill",
		Fields:       addCreditCardFields(),
		HandleAction: addCreditCardAction,
	}
}
