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

// addRecurrenceAction reads the commitment the task declares and writes it
// into the recurrence registry.
func addRecurrenceAction(args api.HandleActionArgs) error {
	recurrence, err := addRecurrenceRead(args)
	if err != nil {
		return err
	}
	recurrences, reachable := args.DataBase.GetSchema(config.RecurrenceSchema)
	if !reachable {
		return errors.New("the " + config.RecurrenceSchema + " registry is unreachable")
	}
	_, failure := recurrences.NewItem(map[string]any{
		config.NameField: recurrence.Description,
		config.DetailField: utils.Pack(recurrence.Description, recurrence.Account,
			recurrence.ToAccount, recurrence.Category),
		config.AmountField: recurrence.Amount,
		config.DayField:    recurrence.Day,
		config.StartField:  recurrence.Start,
		config.EndField:    recurrence.End,
	})
	if failure == nil {
		return nil
	}
	if failure.Type == keepdeps.KeyConflict {
		return errors.New("recurrence " + recurrence.Description + " already exists")
	}
	return errors.New("recurrence " + recurrence.Description +
		" could not be stored: " + failure.Message)
}

// addRecurrenceCommitment is one recurrence as the task reads it, before it
// is stored.
type addRecurrenceCommitment struct {
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

// addRecurrenceRead validates every field of the task against the registries
// it names, and hands back the commitment it declares.
func addRecurrenceRead(args api.HandleActionArgs) (addRecurrenceCommitment, error) {
	empty := addRecurrenceCommitment{}
	description, err := entries.Text(args.Entries, DescriptionField)
	if err != nil {
		return empty, err
	}
	if description == "" {
		return empty, errors.New(DescriptionField + " is required")
	}
	if strings.Contains(description, utils.Separator) {
		return empty, errors.New(DescriptionField + " may not contain " + utils.Separator)
	}
	accounts, reachable := args.DataBase.GetSchema(config.AccountSchema)
	if !reachable {
		return empty, errors.New("the " + config.AccountSchema + " registry is unreachable")
	}
	accountName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return empty, err
	}
	if accountName == "" {
		return empty, errors.New(AccountField + " is required")
	}
	if _, found := accounts.FindByKey(config.NameField, accountName); !found {
		return empty, errors.New("account not found: " + accountName)
	}
	categories, reachable := args.DataBase.GetSchema(config.CategorySchema)
	if !reachable {
		return empty, errors.New("the " + config.CategorySchema + " registry is unreachable")
	}
	categoryName, err := entries.Text(args.Entries, CategoryField)
	if err != nil {
		return empty, err
	}
	if categoryName == "" {
		return empty, errors.New(CategoryField + " is required")
	}
	category, found := categories.FindByKey(config.NameField, categoryName)
	if !found {
		return empty, errors.New("category not found: " + categoryName)
	}
	// A category is a transfer when it is neither revenue nor expense, and a
	// transfer accepts an amount of either sign because it counts as neither.
	revenues := addRecurrenceNumber(category, config.RevenuesField)
	expenses := addRecurrenceNumber(category, config.ExpensesField)
	transfer := revenues != 1 && expenses != 1
	amount, err := entries.Number(args.Entries, AmountField)
	if err != nil {
		return empty, err
	}
	cents := utils.Cents(amount)
	if cents == 0 {
		return empty, errors.New(AmountField + " may not be zero")
	}
	if !transfer && cents > 0 && revenues != 1 {
		return empty, errors.New("a positive amount needs a category with revenues: true — " +
			categoryName + " does not accept income")
	}
	if !transfer && cents < 0 && expenses != 1 {
		return empty, errors.New("a negative amount needs a category with expenses: true — " +
			categoryName + " does not accept expenses")
	}
	day, err := entries.Whole(args.Entries, DayField)
	if err != nil {
		return empty, err
	}
	if day < 1 || day > 31 {
		return empty, errors.New(DayField + " must be a day of the month, from 1 to 31")
	}
	startText, err := entries.Text(args.Entries, StartField)
	if err != nil {
		return empty, err
	}
	start, parseErr := utils.ParseMonth(startText)
	if parseErr != nil {
		return empty, errors.New(StartField +
			" must be a month written as YYYY-MM, not " + startText)
	}
	end := int64(0)
	if entries.Present(args.Entries, EndField) {
		endText, textErr := entries.Text(args.Entries, EndField)
		if textErr != nil {
			return empty, textErr
		}
		end, parseErr = utils.ParseMonth(endText)
		if parseErr != nil {
			return empty, errors.New(EndField +
				" must be a month written as YYYY-MM, not " + endText)
		}
		if end < start {
			return empty, errors.New(EndField + " comes before " + StartField)
		}
	}
	toAccount := ""
	if entries.Present(args.Entries, ToAccountField) {
		toAccount, err = entries.Text(args.Entries, ToAccountField)
		if err != nil {
			return empty, err
		}
		if strings.Contains(toAccount, utils.Separator) {
			return empty, errors.New(ToAccountField + " may not contain " + utils.Separator)
		}
	}
	if toAccount != "" {
		if !transfer {
			return empty, errors.New(ToAccountField +
				" only belongs on a transfer category — one with revenues: false and expenses: false")
		}
		if _, found := accounts.FindByKey(config.NameField, toAccount); !found {
			return empty, errors.New("account not found: " + toAccount)
		}
		if toAccount == accountName {
			return empty, errors.New(ToAccountField +
				" is the same account the money leaves from")
		}
	}
	return addRecurrenceCommitment{
		Description: description,
		Account:     accountName,
		ToAccount:   toAccount,
		Category:    categoryName,
		Amount:      cents,
		Day:         day,
		Start:       start,
		End:         end,
	}, nil
}

// addRecurrenceNumber reads a whole-number field of a stored record, as 0
// when the field is absent or holds something else.
func addRecurrenceNumber(record keepdeps.SchemaItem, field string) int64 {
	value, err := record.Get(field)
	if err != nil {
		return 0
	}
	switch stored := value.(type) {
	case int64:
		return stored
	case int:
		return int64(stored)
	}
	return 0
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
