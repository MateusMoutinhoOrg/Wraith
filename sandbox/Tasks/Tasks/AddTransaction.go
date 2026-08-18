package tasks

import (
	"errors"
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/store"
)

// The bounds a purchase may be split over, as the guide states them.
const (
	// MinInstallments is the smallest split that is still a split.
	MinInstallments = 2
	// MaxInstallments is six years of monthly parts, which is longer than
	// any card offers.
	MaxInstallments = 72
)

// AddTransaction returns the task that records a movement: money in, money
// out, or one leg of a transfer between two of your own accounts. It is the
// task you will run most, and the only one that writes to the ledger.
//
// A purchase paid over several months is still one AddTransaction: give it
// `installments: N` and it writes N monthly parts at once, each landing in
// its own month, each adding back up to the total exactly.
func AddTransaction() api.Task {
	return api.Task{
		Name:        "AddTransaction",
		Description: "Record an income, an expense, or one leg of a transfer",
		Fields: []api.Field{
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
		},
		HandleAction: func(args api.HandleActionArgs) error {
			transaction, parts, err := readTransaction(args)
			if err != nil {
				return err
			}
			for index, part := range spread(transaction, parts) {
				key := store.NextKey(args.DataBase, store.TransactionSchema)
				subject := "transaction " + strconv.Itoa(index+1)
				if err := insert(args, store.TransactionSchema, subject,
					store.TransactionFields(key, part)); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

// readTransaction validates every field of the task and hands back the
// movement it describes, together with how many monthly parts it is split
// over.
func readTransaction(args api.HandleActionArgs) (store.Transaction, int64, error) {
	accountName, err := name(args.Entries, AccountField)
	if err != nil {
		return store.Transaction{}, 0, err
	}
	if _, err := requireAccount(args, accountName); err != nil {
		return store.Transaction{}, 0, err
	}
	categoryName, err := name(args.Entries, CategoryField)
	if err != nil {
		return store.Transaction{}, 0, err
	}
	category, err := requireCategory(args, categoryName)
	if err != nil {
		return store.Transaction{}, 0, err
	}
	amount, err := cents(args.Entries, AmountField)
	if err != nil {
		return store.Transaction{}, 0, err
	}
	if amount == 0 {
		return store.Transaction{}, 0, errors.New(AmountField + " may not be zero")
	}
	if err := checkSign(category, amount); err != nil {
		return store.Transaction{}, 0, err
	}
	when, err := date(args.Entries, DateField)
	if err != nil {
		return store.Transaction{}, 0, err
	}
	description, err := optionalText(args.Entries, DescriptionField)
	if err != nil {
		return store.Transaction{}, 0, err
	}
	payment := when
	if entries.Present(args.Entries, PaymentDateField) {
		payment, err = date(args.Entries, PaymentDateField)
		if err != nil {
			return store.Transaction{}, 0, err
		}
	}
	parts := int64(1)
	if entries.Present(args.Entries, InstallmentsField) {
		parts, err = entries.Whole(args.Entries, InstallmentsField)
		if err != nil {
			return store.Transaction{}, 0, err
		}
		if parts < MinInstallments || parts > MaxInstallments {
			return store.Transaction{}, 0, errors.New(InstallmentsField +
				" must be a whole number from " + strconv.Itoa(MinInstallments) +
				" to " + strconv.Itoa(MaxInstallments))
		}
		if entries.Present(args.Entries, PaymentDateField) {
			return store.Transaction{}, 0, errors.New(InstallmentsField + " and " +
				PaymentDateField + " cannot be combined — each part settles on its own date")
		}
	}
	return store.Transaction{
		Account:     accountName,
		Category:    categoryName,
		Description: description,
		Amount:      amount,
		Date:        when,
		PaymentDate: payment,
	}, parts, nil
}

// spread turns one movement into the monthly parts it is split over. Part 1
// falls on the given date; part k falls on the same day of the k-1-th
// following month, clamped to that month's last day. Each part gets the
// amount divided by the count, and any remainder goes to the first part, so
// the parts always add back up to the total exactly.
func spread(transaction store.Transaction, parts int64) []store.Transaction {
	if parts < MinInstallments {
		return []store.Transaction{transaction}
	}
	each := transaction.Amount / parts
	remainder := transaction.Amount - each*parts
	spread := make([]store.Transaction, 0, parts)
	day := store.DayOf(transaction.Date)
	for index := int64(0); index < parts; index++ {
		part := transaction
		part.Amount = each
		if index == 0 {
			part.Amount += remainder
		}
		part.Date = store.DateIn(store.AddMonths(store.MonthOf(transaction.Date), int(index)), day)
		part.PaymentDate = part.Date
		part.Description = label(transaction.Description, index+1, parts)
		spread = append(spread, part)
	}
	return spread
}

// label marks one part of a split purchase, so a statement shows `Laptop
// (3/12)` rather than twelve identical lines.
func label(description string, part int64, parts int64) string {
	suffix := " (" + strconv.FormatInt(part, 10) + "/" + strconv.FormatInt(parts, 10) + ")"
	if description == "" {
		return "installment" + suffix
	}
	return description + suffix
}
