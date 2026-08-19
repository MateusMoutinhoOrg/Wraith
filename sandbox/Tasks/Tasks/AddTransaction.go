package tasks

import (
	"errors"
	"strconv"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// The bounds a purchase may be split over, as the guide states them.
const (
	// MinInstallments is the smallest split that is still a split.
	MinInstallments = 2
	// MaxInstallments is six years of monthly parts, which is longer than
	// any card offers.
	MaxInstallments = 72
)

// KeyWidth is how many digits a transaction's storage key is padded to, so
// the keys of one registry sort the way they were written.
const KeyWidth = 6

// addTransactionFields declares what the task accepts.
func addTransactionFields() []api.Field {
	return []api.Field{
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "The account the movement happened on"},
		{Name: CategoryField, Type: api.TextField, Required: true,
			Description: "The category it is classified under"},
		{Name: AmountField, Type: api.NumberField, Required: true,
			Description: "Positive for money in, negative for money out"},
		{Name: DateField, Type: api.TextField, Required: true,
			Description: "The date it counts on, written as YYYY-MM-DD"},
		{Name: DescriptionField, Type: api.TextField,
			Description: "What it was"},
		{Name: PaymentDateField, Type: api.TextField,
			Description: "When the money actually moves, if not on `date`"},
		{Name: InstallmentsField, Type: api.NumberField,
			Description: "Split the amount into this many monthly parts (2-72)"},
	}
}

// addTransactionAction reads the movement the task describes, spreads it over
// the months it is split across, and writes every part into the ledger.
func addTransactionAction(args api.HandleActionArgs) error {
	transaction, parts, err := addTransactionRead(args)
	if err != nil {
		return err
	}
	ledger, reachable := args.DataBase.GetSchema(config.TransactionSchema)
	if !reachable {
		return errors.New("the " + config.TransactionSchema + " registry is unreachable")
	}
	for index, part := range addTransactionSpread(transaction, parts) {
		subject := "transaction " + strconv.Itoa(index+1)
		key := addTransactionNextKey(ledger)
		_, failure := ledger.NewItem(map[string]any{
			config.NameField: key,
			config.DetailField: utils.Pack(key, part.Account, part.Category,
				part.Description),
			config.AmountField:      part.Amount,
			config.DateField:        part.Date,
			config.PaymentDateField: part.PaymentDate,
		})
		if failure == nil {
			continue
		}
		if failure.Type == keepdeps.KeyConflict {
			return errors.New(subject + " already exists")
		}
		return errors.New(subject + " could not be stored: " + failure.Message)
	}
	return nil
}

// addTransactionMovement is one movement as the task composes it, before it
// is stored.
type addTransactionMovement struct {
	// Account is the account the movement happened on.
	Account string
	// Category is the category it is classified under.
	Category string
	// Description is the free text it carries.
	Description string
	// Amount is the value in cents — positive for income, negative for an
	// expense.
	Amount int64
	// Date is the date it counts on, as yyyymmdd.
	Date int64
	// PaymentDate is the date the money actually moves, as yyyymmdd.
	PaymentDate int64
}

// addTransactionRead validates every field of the task against the registries
// it names, and hands back the movement it describes together with how many
// monthly parts it is split over.
func addTransactionRead(args api.HandleActionArgs) (addTransactionMovement, int64, error) {
	empty := addTransactionMovement{}
	accounts, reachable := args.DataBase.GetSchema(config.AccountSchema)
	if !reachable {
		return empty, 0, errors.New("the " + config.AccountSchema + " registry is unreachable")
	}
	accountName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return empty, 0, err
	}
	if accountName == "" {
		return empty, 0, errors.New(AccountField + " is required")
	}
	if _, found := accounts.FindByKey(config.NameField, accountName); !found {
		return empty, 0, errors.New("account not found: " + accountName)
	}
	categories, reachable := args.DataBase.GetSchema(config.CategorySchema)
	if !reachable {
		return empty, 0, errors.New("the " + config.CategorySchema + " registry is unreachable")
	}
	categoryName, err := entries.Text(args.Entries, CategoryField)
	if err != nil {
		return empty, 0, err
	}
	if categoryName == "" {
		return empty, 0, errors.New(CategoryField + " is required")
	}
	category, found := categories.FindByKey(config.NameField, categoryName)
	if !found {
		return empty, 0, errors.New("category not found: " + categoryName)
	}
	cents, err := entries.Amount(args.Entries, AmountField)
	if err != nil {
		return empty, 0, err
	}
	if cents == 0 {
		return empty, 0, errors.New(AmountField + " may not be zero")
	}
	// Money may only arrive in a category that accepts income, and may only
	// leave through one that accepts expenses. A transfer category — neither
	// revenue nor expense — accepts both, because it counts as neither.
	revenues := addTransactionNumber(category, config.RevenuesField)
	expenses := addTransactionNumber(category, config.ExpensesField)
	transfer := revenues != 1 && expenses != 1
	if !transfer && cents > 0 && revenues != 1 {
		return empty, 0, errors.New("a positive amount needs a category with revenues: true — " +
			categoryName + " does not accept income")
	}
	if !transfer && cents < 0 && expenses != 1 {
		return empty, 0, errors.New("a negative amount needs a category with expenses: true — " +
			categoryName + " does not accept expenses")
	}
	dateText, err := entries.Text(args.Entries, DateField)
	if err != nil {
		return empty, 0, err
	}
	when, parseErr := utils.ParseDate(dateText)
	if parseErr != nil {
		return empty, 0, errors.New(DateField +
			" must be a date written as YYYY-MM-DD, not " + dateText)
	}
	description := ""
	if entries.Present(args.Entries, DescriptionField) {
		description, err = entries.Text(args.Entries, DescriptionField)
		if err != nil {
			return empty, 0, err
		}
		if strings.Contains(description, utils.Separator) || strings.Contains(description, "\n") || strings.Contains(description, "\r") {
			return empty, 0, errors.New(DescriptionField +
				" may not contain line breaks or " + utils.Separator)
		}
	}
	payment := when
	if entries.Present(args.Entries, PaymentDateField) {
		paymentText, textErr := entries.Text(args.Entries, PaymentDateField)
		if textErr != nil {
			return empty, 0, textErr
		}
		payment, parseErr = utils.ParseDate(paymentText)
		if parseErr != nil {
			return empty, 0, errors.New(PaymentDateField +
				" must be a date written as YYYY-MM-DD, not " + paymentText)
		}
	}
	parts := int64(1)
	if entries.Present(args.Entries, InstallmentsField) {
		parts, err = entries.Whole(args.Entries, InstallmentsField)
		if err != nil {
			return empty, 0, err
		}
		if parts < MinInstallments || parts > MaxInstallments {
			return empty, 0, errors.New(InstallmentsField +
				" must be a whole number from " + strconv.Itoa(MinInstallments) +
				" to " + strconv.Itoa(MaxInstallments))
		}
		if entries.Present(args.Entries, PaymentDateField) {
			return empty, 0, errors.New(InstallmentsField + " and " +
				PaymentDateField + " cannot be combined — each part settles on its own date")
		}
	}
	return addTransactionMovement{
		Account:     accountName,
		Category:    categoryName,
		Description: description,
		Amount:      cents,
		Date:        when,
		PaymentDate: payment,
	}, parts, nil
}

// addTransactionSpread turns one movement into the monthly parts it is split
// over. Part 1 falls on the given date; part k falls on the same day of the
// k-1-th following month, clamped to that month's last day. Each part gets
// the amount divided by the count, and any remainder goes to the first part,
// so the parts always add back up to the total exactly.
func addTransactionSpread(transaction addTransactionMovement, parts int64) []addTransactionMovement {
	if parts < MinInstallments {
		return []addTransactionMovement{transaction}
	}
	each := transaction.Amount / parts
	remainder := transaction.Amount - each*parts
	spread := make([]addTransactionMovement, 0, parts)
	day := utils.DayOf(transaction.Date)
	for index := int64(0); index < parts; index++ {
		part := transaction
		part.Amount = each
		if index == 0 {
			part.Amount += remainder
		}
		part.Date = utils.DateIn(utils.AddMonths(utils.MonthOf(transaction.Date), int(index)), day)
		part.PaymentDate = part.Date
		part.Description = addTransactionLabel(transaction.Description, index+1, parts)
		spread = append(spread, part)
	}
	return spread
}

// addTransactionLabel marks one part of a split purchase, so a statement
// shows `Laptop (3/12)` rather than twelve identical lines.
func addTransactionLabel(description string, part int64, parts int64) string {
	suffix := " (" + strconv.FormatInt(part, 10) + "/" + strconv.FormatInt(parts, 10) + ")"
	if description == "" {
		return "installment" + suffix
	}
	return description + suffix
}

// addTransactionNextKey returns a storage key no record of the ledger carries
// yet, built from how many records it already holds. Keys are never reused,
// so a key already taken is stepped over rather than filled in.
func addTransactionNextKey(ledger keepdeps.SchemaInstance) string {
	stored, failure := ledger.ListAll()
	if failure != nil {
		stored = nil
	}
	candidate := int64(len(stored)) + 1
	for {
		key := utils.Pad(candidate, KeyWidth)
		if _, taken := ledger.FindByKey(config.NameField, key); !taken {
			return key
		}
		candidate++
	}
}

// addTransactionNumber reads a whole-number field of a stored record, as 0
// when the field is absent or holds something else.
func addTransactionNumber(record keepdeps.SchemaItem, field string) int64 {
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

// AddTransaction returns the task that records a movement: money in, money
// out, or one leg of a transfer between two of your own accounts. It is the
// task you will run most, and the only one that writes to the ledger.
//
// A purchase paid over several months is still one AddTransaction: give it
// `installments: N` and it writes N monthly parts at once, each landing in
// its own month, each adding back up to the total exactly.
func AddTransaction() api.Task {
	return api.Task{
		Name:         "AddTransaction",
		Description:  "Record an income, an expense, or one leg of a transfer",
		Fields:       addTransactionFields(),
		HandleAction: addTransactionAction,
	}
}
