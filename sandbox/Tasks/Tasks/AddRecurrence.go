package tasks

import (
	"errors"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// addRecurrenceFields declares what the task accepts.
func addRecurrenceFields() []api.Field {
	return []api.Field{
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
	}
}

// addRecurrenceAction runs the task against the database it is handed.
func addRecurrenceAction(args api.HandleActionArgs) error {
	recurrence, err := readRecurrence(args)
	if err != nil {
		return err
	}
	return insert(args, config.RecurrenceSchema,
		"recurrence "+recurrence.Description, map[string]any{
			config.NameField: recurrence.Description,
			config.DetailField: utils.Pack(recurrence.Description, recurrence.Account,
				recurrence.ToAccount, recurrence.Category),
			config.AmountField: recurrence.Amount,
			config.DayField:    recurrence.Day,
			config.StartField:  recurrence.Start,
			config.EndField:    recurrence.End,
		})
}

// AddRecurrence returns the task that declares a commitment repeating every
// month — a salary, rent, a subscription, a standing transfer into savings.
//
// A recurrence is the only thing in the brain that describes the future. It
// moves no money and writes no transaction: it is a rule the forecast reads.
// When the day arrives you still record what actually happened with
// AddTransaction, with the real amount.
func AddRecurrence() api.Task {
	return api.Task{
		Name:         "AddRecurrence",
		Description:  "Declare a commitment that repeats every month",
		Fields:       addRecurrenceFields(),
		HandleAction: addRecurrenceAction,
	}
}

// commitment is one recurrence as the task reads it, before it is stored.
type commitment struct {
	// Description identifies the recurrence, and is how RemoveRecurrence
	// addresses it.
	Description string
	// Account is the account the money leaves from or arrives in.
	Account string
	// ToAccount is the destination account of a recurring transfer, or "".
	ToAccount string
	// Category is the category it is classified under.
	Category string
	// Amount is the value per occurrence, in cents.
	Amount int64
	// Day is the day of the month it falls on.
	Day int64
	// Start is the first month it applies, as yyyymm.
	Start int64
	// End is the last month it applies, as yyyymm, or 0 when open-ended.
	End int64
}

// readRecurrence validates every field of the task and hands back the
// commitment it declares.
func readRecurrence(args api.HandleActionArgs) (commitment, error) {
	description, err := name(args.Entries, DescriptionField)
	if err != nil {
		return commitment{}, err
	}
	accountName, err := name(args.Entries, AccountField)
	if err != nil {
		return commitment{}, err
	}
	if _, err := requireAccount(args, accountName); err != nil {
		return commitment{}, err
	}
	categoryName, err := name(args.Entries, CategoryField)
	if err != nil {
		return commitment{}, err
	}
	category, err := requireCategory(args, categoryName)
	if err != nil {
		return commitment{}, err
	}
	amount, err := cents(args.Entries, AmountField)
	if err != nil {
		return commitment{}, err
	}
	if amount == 0 {
		return commitment{}, errors.New(AmountField + " may not be zero")
	}
	if err := checkSign(category, amount); err != nil {
		return commitment{}, err
	}
	day, err := dayOfMonth(args.Entries, DayField)
	if err != nil {
		return commitment{}, err
	}
	start, err := month(args.Entries, StartField)
	if err != nil {
		return commitment{}, err
	}
	end := int64(0)
	if entries.Present(args.Entries, EndField) {
		end, err = month(args.Entries, EndField)
		if err != nil {
			return commitment{}, err
		}
		if end < start {
			return commitment{}, errors.New(EndField + " comes before " + StartField)
		}
	}
	toAccount, err := optionalText(args.Entries, ToAccountField)
	if err != nil {
		return commitment{}, err
	}
	if toAccount != "" {
		if !isTransfer(category) {
			return commitment{}, errors.New(ToAccountField +
				" only belongs on a transfer category — one with revenues: false and expenses: false")
		}
		if _, err := requireAccount(args, toAccount); err != nil {
			return commitment{}, err
		}
		if toAccount == accountName {
			return commitment{}, errors.New(ToAccountField +
				" is the same account the money leaves from")
		}
	}
	return commitment{
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
