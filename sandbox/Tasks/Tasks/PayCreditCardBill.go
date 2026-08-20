package tasks

import (
	"errors"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// payCreditCardBillFields declares what the task accepts.
func payCreditCardBillFields() []api.Field {
	return []api.Field{
		{Name: CardField, Type: api.TextField, Required: true,
			Description: "The credit card whose bill is being paid"},
		{Name: AccountField, Type: api.TextField, Required: true,
			Description: "The account the money leaves from"},
		{Name: CategoryField, Type: api.TextField, Required: true,
			Description: "A transfer category — one with revenues and expenses both false"},
		{Name: DateField, Type: api.TextField, Required: true,
			Description: "The date the bill is paid, written as YYYY-MM-DD"},
		{Name: AmountField, Type: api.NumberField,
			Description: "How much is paid, as a positive number. " +
				"Omitted, it pays everything the closed statements still ask for"},
		{Name: PaymentDateField, Type: api.TextField,
			Description: "When the money actually leaves the account, if not on `date`"},
		{Name: DescriptionField, Type: api.TextField,
			Description: "What the payment was, as both legs will show it"},
	}
}

// payCreditCardBillPayment is the payment the task composes, before it is
// stored.
type payCreditCardBillPayment struct {
	// Card is the credit card the debt sits on.
	Card string
	// Account is the account the money leaves from.
	Account string
	// Category is the transfer category both legs are classified under.
	Category string
	// Description is the free text both legs carry.
	Description string
	// Amount is what is paid, in cents, as a positive figure.
	Amount int64
	// Date is the date the payment counts on, as yyyymmdd.
	Date int64
	// PaymentDate is the date the money actually leaves, as yyyymmdd.
	PaymentDate int64
}

// payCreditCardBillAction validates the two accounts, the category and the
// amount, then writes the two legs of the payment into the ledger.
func payCreditCardBillAction(args api.HandleActionArgs) error {
	payment, err := payCreditCardBillRead(args)
	if err != nil {
		return err
	}
	ledgerSchema, reachable := args.DataBase.GetSchema(config.TransactionSchema)
	if !reachable {
		return errors.New("the " + config.TransactionSchema + " registry is unreachable")
	}
	legs := []struct {
		Account string
		Amount  int64
	}{
		{Account: payment.Account, Amount: -payment.Amount},
		{Account: payment.Card, Amount: payment.Amount},
	}
	for _, leg := range legs {
		key := payCreditCardBillNextKey(ledgerSchema)
		_, failure := ledgerSchema.NewItem(map[string]any{
			config.NameField: key,
			config.DetailField: utils.Pack(key, leg.Account, payment.Category,
				payment.Description),
			config.AmountField:      leg.Amount,
			config.DateField:        payment.Date,
			config.PaymentDateField: payment.PaymentDate,
		})
		if failure == nil {
			continue
		}
		subject := "the " + leg.Account + " leg of the payment"
		if failure.Type == keepdeps.KeyConflict {
			return errors.New(subject + " already exists")
		}
		return errors.New(subject + " could not be stored: " + failure.Message)
	}
	return nil
}

// payCreditCardBillRead validates every field against the registries it names
// and hands back the payment it describes. Nothing is written until it
// returns, so a payment that cannot be made leaves the ledger untouched.
func payCreditCardBillRead(args api.HandleActionArgs) (payCreditCardBillPayment, error) {
	empty := payCreditCardBillPayment{}
	state := ledger.Load(args.Deps, args.DataBase)

	cardName, err := entries.Text(args.Entries, CardField)
	if err != nil {
		return empty, err
	}
	if cardName == "" {
		return empty, errors.New(CardField + " is required")
	}
	card, found := state.Account(cardName)
	if !found {
		return empty, errors.New("credit card not found: " + cardName)
	}
	if !card.IsCard() {
		return empty, errors.New(cardName + " is an account, not a credit card — " +
			"a bill is paid on a card")
	}
	accountName, err := entries.Text(args.Entries, AccountField)
	if err != nil {
		return empty, err
	}
	if accountName == "" {
		return empty, errors.New(AccountField + " is required")
	}
	account, found := state.Account(accountName)
	if !found {
		return empty, errors.New("account not found: " + accountName)
	}
	if account.IsCard() {
		return empty, errors.New(accountName + " is a credit card — a bill is paid " +
			"from an account money actually sits in")
	}
	if accountName == cardName {
		return empty, errors.New(AccountField + " and " + CardField +
			" must name two different records")
	}
	categoryName, err := entries.Text(args.Entries, CategoryField)
	if err != nil {
		return empty, err
	}
	if categoryName == "" {
		return empty, errors.New(CategoryField + " is required")
	}
	category, found := state.Category(categoryName)
	if !found {
		return empty, errors.New("category not found: " + categoryName)
	}
	if !category.IsTransfer() {
		return empty, errors.New("paying a bill moves money between two of your own " +
			"records, so it needs a transfer category — " + categoryName +
			" must have revenues and expenses both false")
	}
	dateText, err := entries.Text(args.Entries, DateField)
	if err != nil {
		return empty, err
	}
	when, parseErr := utils.ParseDate(dateText)
	if parseErr != nil {
		return empty, errors.New(DateField +
			" must be a date written as YYYY-MM-DD, not " + dateText)
	}
	settles := when
	if entries.Present(args.Entries, PaymentDateField) {
		paymentText, textErr := entries.Text(args.Entries, PaymentDateField)
		if textErr != nil {
			return empty, textErr
		}
		settles, parseErr = utils.ParseDate(paymentText)
		if parseErr != nil {
			return empty, errors.New(PaymentDateField +
				" must be a date written as YYYY-MM-DD, not " + paymentText)
		}
	}
	cents, err := payCreditCardBillAmount(args, state, card)
	if err != nil {
		return empty, err
	}
	description := payCreditCardBillLabel(cardName)
	if entries.Present(args.Entries, DescriptionField) {
		description, err = entries.Text(args.Entries, DescriptionField)
		if err != nil {
			return empty, err
		}
		if strings.Contains(description, utils.Separator) ||
			strings.Contains(description, "\n") || strings.Contains(description, "\r") {
			return empty, errors.New(DescriptionField +
				" may not contain line breaks or " + utils.Separator)
		}
		if description == "" {
			description = payCreditCardBillLabel(cardName)
		}
	}
	return payCreditCardBillPayment{
		Card:        cardName,
		Account:     accountName,
		Category:    categoryName,
		Description: description,
		Amount:      cents,
		Date:        when,
		PaymentDate: settles,
	}, nil
}

// payCreditCardBillAmount reads what is being paid: the amount the task was
// given, or everything the card's closed statements still ask for when it was
// given none. A card whose statements ask for nothing is refused rather than
// paid zero, because a movement of zero is not a payment.
func payCreditCardBillAmount(args api.HandleActionArgs, state ledger.State,
	card ledger.Account) (int64, error) {
	if entries.Present(args.Entries, AmountField) {
		cents, err := entries.Amount(args.Entries, AmountField)
		if err != nil {
			return 0, err
		}
		if cents <= 0 {
			return 0, errors.New(AmountField + " is what you are paying, " +
				"so it must be a positive number")
		}
		return cents, nil
	}
	due := state.AmountDue(card)
	if due <= 0 {
		return 0, errors.New(card.Name + " has no closed statement left to pay — " +
			"give " + AmountField + " to pay something else off anyway")
	}
	return due, nil
}

// payCreditCardBillLabel is what a payment is described as when the task was
// given no description of its own.
func payCreditCardBillLabel(cardName string) string {
	return cardName + " bill payment"
}

// payCreditCardBillNextKey returns a storage key no record of the ledger
// carries yet, built from how many records it already holds. Keys are never
// reused, so a key already taken is stepped over rather than filled in.
func payCreditCardBillNextKey(ledgerSchema keepdeps.SchemaInstance) string {
	stored, failure := ledgerSchema.ListAll()
	if failure != nil {
		stored = nil
	}
	candidate := int64(len(stored)) + 1
	for {
		key := utils.Pad(candidate, KeyWidth)
		if _, taken := ledgerSchema.FindByKey(config.NameField, key); !taken {
			return key
		}
		candidate++
	}
}

// PayCreditCardBill returns the task that pays a credit card bill: the one
// movement that actually takes money out of your account for purchases
// already counted on the day they happened.
//
// It writes both legs at once — the money leaving the account, and the same
// amount arriving on the card — under a transfer category, so paying a bill
// never counts as an expense a second time. Given no `amount`, it pays
// exactly what the card's closed statements still ask for.
func PayCreditCardBill() api.Task {
	return api.Task{
		Name:         "PayCreditCardBill",
		Description:  "Pay a credit card bill, moving money from an account to the card",
		Fields:       payCreditCardBillFields(),
		HandleAction: payCreditCardBillAction,
	}
}
