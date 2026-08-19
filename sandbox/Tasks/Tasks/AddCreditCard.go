package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
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

// addCreditCardAction runs the task against the database it is handed.
func addCreditCardAction(args api.HandleActionArgs) error {
	cardName, err := name(args.Entries, AccountField)
	if err != nil {
		return err
	}
	limit, err := cents(args.Entries, LimitField)
	if err != nil {
		return err
	}
	if limit < 0 {
		return errors.New(LimitField + " is a credit limit, so it may not be negative")
	}
	closingDay, err := dayOfMonth(args.Entries, ClosingDayField)
	if err != nil {
		return err
	}
	dueDay, err := dayOfMonth(args.Entries, DueDayField)
	if err != nil {
		return err
	}
	opening, err := cents(args.Entries, OpeningField)
	if err != nil {
		return err
	}
	if opening > 0 {
		return errors.New(OpeningField + " on a card is what you already owe, " +
			"so it is written as a negative number")
	}
	return insert(args, config.AccountSchema, "credit card "+cardName, map[string]any{
		config.NameField:       cardName,
		config.KindField:       int64(config.KindCard),
		config.OpeningField:    opening,
		config.LimitField:      limit,
		config.ClosingDayField: closingDay,
		config.DueDayField:     dueDay,
	})
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
