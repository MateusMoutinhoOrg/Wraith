package config

// The shape of the brain's database, as compile-time constants: the four
// registries every task writes and every visualization reads, the fields each
// one carries, and the Props the injected Keep library is handed to open
// them.
//
// It lives here for the same reason the interface's words do — a registry or
// a field is spelled once, under the compiler's eye, so a renamed constant is
// a build failure rather than a record that silently stops being found.
//
// Keep stores unique string keys and whole numbers, and nothing else. Three
// consequences shape every schema below, and they are the whole trick worth
// knowing before adding a registry of your own:
//
//   - Money is held in cents, as a whole number. utils.Money renders it back.
//   - Dates are held as whole numbers: 2026-08-18 is 20260818, 2026-08 is
//     202608. They sort correctly as numbers, which is what a month index is
//     built out of.
//   - Free text that is not unique — a description, an account's name on a
//     transaction — cannot be a key of its own, so it travels packed into one
//     key alongside a value that *is* unique. utils.Pack writes it,
//     utils.Unpack reads it back.

import (
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
)

// The registries the brain keeps, one Keep schema each.
const (
	// AccountSchema holds the accounts money sits in.
	AccountSchema = "account"
	// CategorySchema holds the categories transactions are classified under.
	CategorySchema = "category"
	// TransactionSchema holds every recorded movement.
	TransactionSchema = "transaction"
	// RecurrenceSchema holds the commitments the forecast projects.
	RecurrenceSchema = "recurrence"
)

// The fields the registries carry. A field named Detail is a packed key; see
// utils.Pack.
const (
	// NameField is a record's unique name.
	NameField = "name"
	// DetailField is the packed key carrying a record's free text.
	DetailField = "detail"
	// RevenuesField is 1 when a category accepts positive amounts.
	RevenuesField = "revenues"
	// ExpensesField is 1 when a category accepts negative amounts.
	ExpensesField = "expenses"
	// AmountField is a value in cents, signed.
	AmountField = "amount"
	// DateField is a date as yyyymmdd.
	DateField = "date"
	// DayField is the day of the month a recurrence falls on.
	DayField = "day"
	// StartField is the first month a recurrence applies, as yyyymm.
	StartField = "start"
	// EndField is the last month a recurrence applies, as yyyymm, or 0 for
	// open-ended.
	EndField = "end"
)

// The parts a packed DetailField is composed of, one pair per registry: how
// many parts the key holds, and which position each of them is at. The first
// part is always the record's own unique name, which is what makes the whole
// key unique.
const (
	// CategoryParts is `name|parent|description`.
	CategoryParts = 3
	// CategoryParent is the position of a category's parent.
	CategoryParent = 1
	// CategoryDescription is the position of a category's description.
	CategoryDescription = 2

	// TransactionParts is `key|account|category|description`.
	TransactionParts = 4
	// TransactionAccount is the position of a movement's account.
	TransactionAccount = 1
	// TransactionCategory is the position of a movement's category.
	TransactionCategory = 2
	// TransactionDescription is the position of a movement's description.
	TransactionDescription = 3

	// RecurrenceParts is `description|account|toAccount|category`.
	RecurrenceParts = 4
	// RecurrenceAccount is the position of a commitment's account.
	RecurrenceAccount = 1
	// RecurrenceToAccount is the position of a commitment's destination.
	RecurrenceToAccount = 2
	// RecurrenceCategory is the position of a commitment's category.
	RecurrenceCategory = 3
)

// DatabaseProps describes the brain's database. path is the folder the
// registries live in, relative to the vault root: it is api.Lib.DatabasePath,
// which the interface's `--database` flag overrides, so pointing the same
// binary at a second vault changes nothing but this prefix.
func DatabaseProps(path string) keepdeps.Props {
	return keepdeps.Props{
		Path: strings.TrimSuffix(path, "/") + "/",
		Schemas: []keepdeps.Schema{
			{
				Name: AccountSchema,
				Itens: []keepdeps.Item{
					{Name: NameField, Type: keepdeps.Key, Required: true},
				},
			},
			{
				Name: CategorySchema,
				Itens: []keepdeps.Item{
					{Name: NameField, Type: keepdeps.Key, Required: true},
					{Name: DetailField, Type: keepdeps.Key, Required: true},
					{Name: RevenuesField, Type: keepdeps.Int, Required: true},
					{Name: ExpensesField, Type: keepdeps.Int, Required: true},
				},
			},
			{
				Name: TransactionSchema,
				Itens: []keepdeps.Item{
					{Name: NameField, Type: keepdeps.Key, Required: true},
					{Name: DetailField, Type: keepdeps.Key, Required: true},
					{Name: AmountField, Type: keepdeps.Int, Required: true},
					{Name: DateField, Type: keepdeps.Int, Required: true},
				},
			},
			{
				Name: RecurrenceSchema,
				Itens: []keepdeps.Item{
					{Name: NameField, Type: keepdeps.Key, Required: true},
					{Name: DetailField, Type: keepdeps.Key, Required: true},
					{Name: AmountField, Type: keepdeps.Int, Required: true},
					{Name: DayField, Type: keepdeps.Int, Required: true},
					{Name: StartField, Type: keepdeps.Int, Required: true},
					{Name: EndField, Type: keepdeps.Int, Required: true},
				},
			},
		},
	}
}
