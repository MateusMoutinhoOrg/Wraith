package tasks

import (
	"errors"
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
)

// ModifyTransaction returns the task that corrects a movement already in the
// ledger. Every field but `id` is optional: what you give is overwritten,
// what you leave out stays as it is. The id is the one a statement shows
// beside each line.
//
// One part of a split purchase is an ordinary transaction like any other, so
// this is how a single part is corrected — the other parts are untouched.
func ModifyTransaction() api.Task {
	return api.Task{
		Name:        "ModifyTransaction",
		Description: "Correct a transaction already in the ledger",
		Fields: []api.Field{
			{Name: IdField, Type: api.NumberField, Required: true,
				Description: "The id of the transaction, as the statement shows it"},
			{Name: AccountField, Type: api.TextField,
				Description: "Move it to another account"},
			{Name: CategoryField, Type: api.TextField,
				Description: "Classify it under another category"},
			{Name: DescriptionField, Type: api.TextField,
				Description: "Rewrite what it was"},
			{Name: AmountField, Type: api.NumberField,
				Description: "Correct the amount"},
			{Name: DateField, Type: api.TextField,
				Description: "Correct the date it counts on, as YYYY-MM-DD"},
			{Name: PaymentDateField, Type: api.TextField,
				Description: "Correct when the money actually moves, as YYYY-MM-DD"},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			id, err := entries.Whole(args.Entries, IdField)
			if err != nil {
				return err
			}
			record, found := findTransaction(args, id)
			if !found {
				return errors.New("transaction not found: " + strconv.FormatInt(id, 10))
			}
			current := readMovement(record)
			updated, err := applyChanges(args, current)
			if err != nil {
				return err
			}
			for field, value := range movementFields(current.Key, updated) {
				if err := record.Update(field, value); err != nil {
					return errors.New("transaction " + strconv.FormatInt(id, 10) +
						" could not be updated: " + err.Message)
				}
			}
			return nil
		},
	}
}

// applyChanges overlays the fields the task was given onto the stored
// movement, checking each one the same way AddTransaction checks it.
func applyChanges(args api.HandleActionArgs, current movement) (movement, error) {
	updated := current
	if entries.Present(args.Entries, AccountField) {
		accountName, err := name(args.Entries, AccountField)
		if err != nil {
			return updated, err
		}
		if _, err := requireAccount(args, accountName); err != nil {
			return updated, err
		}
		updated.Account = accountName
	}
	if entries.Present(args.Entries, CategoryField) {
		categoryName, err := name(args.Entries, CategoryField)
		if err != nil {
			return updated, err
		}
		if _, err := requireCategory(args, categoryName); err != nil {
			return updated, err
		}
		updated.Category = categoryName
	}
	if entries.Present(args.Entries, DescriptionField) {
		description, err := optionalText(args.Entries, DescriptionField)
		if err != nil {
			return updated, err
		}
		updated.Description = description
	}
	if entries.Present(args.Entries, AmountField) {
		amount, err := cents(args.Entries, AmountField)
		if err != nil {
			return updated, err
		}
		if amount == 0 {
			return updated, errors.New(AmountField + " may not be zero")
		}
		updated.Amount = amount
	}
	if entries.Present(args.Entries, DateField) {
		when, err := date(args.Entries, DateField)
		if err != nil {
			return updated, err
		}
		if updated.PaymentDate == current.Date {
			updated.PaymentDate = when
		}
		updated.Date = when
	}
	if entries.Present(args.Entries, PaymentDateField) {
		payment, err := date(args.Entries, PaymentDateField)
		if err != nil {
			return updated, err
		}
		updated.PaymentDate = payment
	}
	category, err := requireCategory(args, updated.Category)
	if err != nil {
		return updated, err
	}
	if err := checkSign(category, updated.Amount); err != nil {
		return updated, err
	}
	return updated, nil
}

// findTransaction returns the stored record of one movement by its permanent
// identifier — the id a statement shows beside each line. The record itself is
// what the task needs, because correcting a movement is writing back to it.
func findTransaction(args api.HandleActionArgs, id int64) (keepdeps.SchemaItem, bool) {
	for _, record := range records(args, config.TransactionSchema) {
		if record.Id == id {
			return record, true
		}
	}
	return keepdeps.SchemaItem{}, false
}
