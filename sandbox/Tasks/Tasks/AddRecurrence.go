package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// AddRecurrence returns the task that declares a commitment repeating every
// month — a salary, rent, a subscription, a standing transfer into savings.
//
// A recurrence is the only thing in the brain that describes the future. It
// moves no money and writes no transaction: it is a rule the forecast reads.
// When the day arrives you still record what actually happened with
// AddTransaction, with the real amount.
func AddRecurrence() api.Task {
	return api.Task{
		Name:        "AddRecurrence",
		Description: "Declare a commitment that repeats every month",
		Fields: []api.Field{
			{Name: DescriptionField, Type: api.TextField, Required: true,
				Description: "Identifies the recurrence — it is how RemoveRecurrence finds it"},
			{Name: AccountField, Type: api.TextField, Required: true,
				Description: "The account the money leaves from or arrives in"},
			{Name: CategoryField, Type: api.TextField, Required: true,
				Description: "The category it is classified under"},
			{Name: AmountField, Type: api.NumberField, Required: true,
				Description: "Amount per occurrence, positive for income"},
			{Name: DayField, Type: api.NumberField, Required: true,
				Description: "Day of the month it falls on (1-31)"},
			{Name: StartField, Type: api.TextField, Required: true,
				Description: "First month it applies, written as YYYY-MM"},
			{Name: EndField, Type: api.TextField,
				Description: "Last month it applies, as YYYY-MM. Omitted means open-ended"},
			{Name: ToAccountField, Type: api.TextField,
				Description: "Destination account, for a recurring transfer"},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			recurrence, err := readRecurrence(args)
			if err != nil {
				return err
			}
			return insert(args, store.RecurrenceSchema,
				"recurrence "+recurrence.Description, map[string]any{
					store.NameField: recurrence.Description,
					store.DetailField: store.Pack(recurrence.Description, recurrence.Account,
						recurrence.ToAccount, recurrence.Category),
					store.AmountField: recurrence.Amount,
					store.DayField:    recurrence.Day,
					store.StartField:  recurrence.Start,
					store.EndField:    recurrence.End,
				})
		},
	}
}

// readRecurrence validates every field of the task and hands back the
// commitment it declares.
func readRecurrence(args api.HandleActionArgs) (store.Recurrence, error) {
	description, err := name(args.Entries, DescriptionField)
	if err != nil {
		return store.Recurrence{}, err
	}
	accountName, err := name(args.Entries, AccountField)
	if err != nil {
		return store.Recurrence{}, err
	}
	if _, err := requireAccount(args, accountName); err != nil {
		return store.Recurrence{}, err
	}
	categoryName, err := name(args.Entries, CategoryField)
	if err != nil {
		return store.Recurrence{}, err
	}
	category, err := requireCategory(args, categoryName)
	if err != nil {
		return store.Recurrence{}, err
	}
	amount, err := cents(args.Entries, AmountField)
	if err != nil {
		return store.Recurrence{}, err
	}
	if amount == 0 {
		return store.Recurrence{}, errors.New(AmountField + " may not be zero")
	}
	if err := checkSign(category, amount); err != nil {
		return store.Recurrence{}, err
	}
	day, err := dayOfMonth(args.Entries, DayField)
	if err != nil {
		return store.Recurrence{}, err
	}
	start, err := month(args.Entries, StartField)
	if err != nil {
		return store.Recurrence{}, err
	}
	end := int64(0)
	if entries.Present(args.Entries, EndField) {
		end, err = month(args.Entries, EndField)
		if err != nil {
			return store.Recurrence{}, err
		}
		if end < start {
			return store.Recurrence{}, errors.New(EndField + " comes before " + StartField)
		}
	}
	toAccount, err := optionalText(args.Entries, ToAccountField)
	if err != nil {
		return store.Recurrence{}, err
	}
	if toAccount != "" {
		if !category.IsTransfer() {
			return store.Recurrence{}, errors.New(ToAccountField +
				" only belongs on a transfer category — one with revenues: false and expenses: false")
		}
		if _, err := requireAccount(args, toAccount); err != nil {
			return store.Recurrence{}, err
		}
		if toAccount == accountName {
			return store.Recurrence{}, errors.New(ToAccountField +
				" is the same account the money leaves from")
		}
	}
	return store.Recurrence{
		Description: description,
		Account:     accountName,
		ToAccount:   toAccount,
		Category:    categoryName,
		Amount:      amount,
		Day:         day,
		Start:       start,
		End:         end,
	}, nil
}
