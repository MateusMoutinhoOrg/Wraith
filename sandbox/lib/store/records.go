package store

// The five registries, read back as plain Go values. A task writes records
// through the Keep schemas above; everything that reads them — every other
// task's validation, and every visualization — goes through the views here,
// so the packing rules live in exactly one place.

import (
	"sort"
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
)

// Account is one account or credit card of the registry.
type Account struct {
	// Name is the account's unique display name.
	Name string
	// Kind is KindAccount or KindCard.
	Kind int64
	// Opening is the balance the account started at, in cents. On a card it
	// is the amount already owed, and therefore negative.
	Opening int64
	// Limit is a card's total credit limit, in cents. It is zero on a plain
	// account.
	Limit int64
	// ClosingDay is the day of the month a card's bill closes.
	ClosingDay int64
	// DueDay is the day of the month a card's bill is due.
	DueDay int64
}

// IsCard reports whether the account is a credit card.
func (a Account) IsCard() bool { return a.Kind == KindCard }

// Category is one category of the registry.
type Category struct {
	// Name is the category's unique display name.
	Name string
	// Parent is the category this one hangs under, or "" when it is a root.
	Parent string
	// Description explains what the category classifies.
	Description string
	// Revenues reports whether the category accepts positive amounts.
	Revenues bool
	// Expenses reports whether the category accepts negative amounts.
	Expenses bool
}

// IsTransfer reports whether the category is a transfer category — one that
// is neither revenue nor expense, and is therefore how a movement between
// two of your own accounts is recorded.
func (c Category) IsTransfer() bool { return !c.Revenues && !c.Expenses }

// Transaction is one recorded movement.
type Transaction struct {
	// Id is the record's permanent identifier, and what ModifyTransaction
	// addresses it by.
	Id int64
	// Key is the record's unique storage key.
	Key string
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

// Recurrence is one declared commitment the forecast projects.
type Recurrence struct {
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

// AppliesIn reports whether the recurrence is in force in the given month.
func (r Recurrence) AppliesIn(month int64) bool {
	if month < r.Start {
		return false
	}
	return r.End == 0 || month <= r.End
}

// ReadAccount reads one stored record as an Account.
func ReadAccount(record keepdeps.SchemaItem) Account {
	return Account{
		Name:       Text(record, NameField),
		Kind:       Number(record, KindField),
		Opening:    Number(record, OpeningField),
		Limit:      Number(record, LimitField),
		ClosingDay: Number(record, ClosingDayField),
		DueDay:     Number(record, DueDayField),
	}
}

// ReadCategory reads one stored record as a Category. Its packed detail is
// `name|parent|description`.
func ReadCategory(record keepdeps.SchemaItem) Category {
	parts := Unpack(Text(record, DetailField), 3)
	return Category{
		Name:        Text(record, NameField),
		Parent:      parts[1],
		Description: parts[2],
		Revenues:    Number(record, RevenuesField) == 1,
		Expenses:    Number(record, ExpensesField) == 1,
	}
}

// ReadTransaction reads one stored record as a Transaction. Its packed
// detail is `key|account|category|description`.
func ReadTransaction(record keepdeps.SchemaItem) Transaction {
	parts := Unpack(Text(record, DetailField), 4)
	return Transaction{
		Id:          record.Id,
		Key:         Text(record, NameField),
		Account:     parts[1],
		Category:    parts[2],
		Description: parts[3],
		Amount:      Number(record, AmountField),
		Date:        Number(record, DateField),
		PaymentDate: Number(record, PaymentDateField),
	}
}

// ReadRecurrence reads one stored record as a Recurrence. Its packed detail
// is `description|account|toAccount|category`.
func ReadRecurrence(record keepdeps.SchemaItem) Recurrence {
	parts := Unpack(Text(record, DetailField), 4)
	return Recurrence{
		Description: Text(record, NameField),
		Account:     parts[1],
		ToAccount:   parts[2],
		Category:    parts[3],
		Amount:      Number(record, AmountField),
		Day:         Number(record, DayField),
		Start:       Number(record, StartField),
		End:         Number(record, EndField),
	}
}

// Accounts returns every account and card of the registry, ordered by name.
func Accounts(database keepdeps.KeepDatabase) []Account {
	accounts := []Account{}
	for _, record := range All(database, AccountSchema) {
		accounts = append(accounts, ReadAccount(record))
	}
	sort.Slice(accounts, func(i int, j int) bool { return accounts[i].Name < accounts[j].Name })
	return accounts
}

// Categories returns every category of the registry, ordered by name.
func Categories(database keepdeps.KeepDatabase) []Category {
	categories := []Category{}
	for _, record := range All(database, CategorySchema) {
		categories = append(categories, ReadCategory(record))
	}
	sort.Slice(categories, func(i int, j int) bool { return categories[i].Name < categories[j].Name })
	return categories
}

// Transactions returns every recorded movement, oldest first, ties broken by
// the storage key so the order never depends on how the records came back.
func Transactions(database keepdeps.KeepDatabase) []Transaction {
	transactions := []Transaction{}
	for _, record := range All(database, TransactionSchema) {
		transactions = append(transactions, ReadTransaction(record))
	}
	sort.Slice(transactions, func(i int, j int) bool {
		if transactions[i].Date != transactions[j].Date {
			return transactions[i].Date < transactions[j].Date
		}
		return transactions[i].Key < transactions[j].Key
	})
	return transactions
}

// Recurrences returns every declared commitment, ordered by description.
func Recurrences(database keepdeps.KeepDatabase) []Recurrence {
	recurrences := []Recurrence{}
	for _, record := range All(database, RecurrenceSchema) {
		recurrences = append(recurrences, ReadRecurrence(record))
	}
	sort.Slice(recurrences, func(i int, j int) bool {
		return recurrences[i].Description < recurrences[j].Description
	})
	return recurrences
}

// FindAccount returns one account of the registry by name.
func FindAccount(database keepdeps.KeepDatabase, name string) (Account, bool) {
	record, found := Find(database, AccountSchema, name)
	if !found {
		return Account{}, false
	}
	return ReadAccount(record), true
}

// FindCategory returns one category of the registry by name.
func FindCategory(database keepdeps.KeepDatabase, name string) (Category, bool) {
	record, found := Find(database, CategorySchema, name)
	if !found {
		return Category{}, false
	}
	return ReadCategory(record), true
}

// FindTransaction returns the stored record of one transaction by its
// permanent identifier. The record itself is handed back rather than the
// view, because the only caller — ModifyTransaction — has to write to it.
func FindTransaction(database keepdeps.KeepDatabase, id int64) (keepdeps.SchemaItem, bool) {
	for _, record := range All(database, TransactionSchema) {
		if record.Id == id {
			return record, true
		}
	}
	return keepdeps.SchemaItem{}, false
}

// NextKey returns a storage key no record of the registry carries yet, built
// from how many records it already holds. Keys are never reused, so a key
// already taken is stepped over rather than filled in.
func NextKey(database keepdeps.KeepDatabase, schema string) string {
	instance, ok := Schema(database, schema)
	if !ok {
		return "1"
	}
	records, err := instance.ListAll()
	if err != nil {
		records = nil
	}
	candidate := int64(len(records)) + 1
	for {
		key := pad(candidate, 6)
		if _, taken := instance.FindByKey(NameField, key); !taken {
			return key
		}
		candidate++
	}
}

// TransactionFields composes the stored fields of one transaction, packing
// its account, category and description into the detail key beside its own
// unique storage key.
func TransactionFields(key string, t Transaction) map[string]any {
	return map[string]any{
		NameField:        key,
		DetailField:      Pack(key, t.Account, t.Category, t.Description),
		AmountField:      t.Amount,
		DateField:        t.Date,
		PaymentDateField: t.PaymentDate,
	}
}

// IdText renders a transaction's permanent identifier, which is what a
// ModifyTransaction task addresses it by.
func IdText(id int64) string { return strconv.FormatInt(id, 10) }
