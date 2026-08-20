package ledger

// The three registries, read back as plain Go values. A task writes records
// through the Keep schemas config declares; everything that reads them goes
// through the views here, so the packing rules are unpacked in exactly one
// place.
//
// Reading is done straight through the injected Keep library: a schema is
// asked of the database, the schema lists its records, and a record hands
// back one field at a time. The two helpers at the bottom are only the type
// assertions that reading a field costs.

import (
	"sort"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// Account is one account of the registry — somewhere money sits. It carries
// nothing but its name: what it holds is the sum of the movements dated on
// it, and nothing else.
type Account struct {
	// Name is the account's unique display name.
	Name string
}

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
// is neither revenue nor expense, and is therefore how a movement between two
// of your own accounts is recorded.
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
	// Date is the date it counts on, and the date it settles on, as
	// yyyymmdd. A movement moves its whole amount that day.
	Date int64
	// AccountId is the id of the account record the movement points at — the
	// account whose nested registry indexes this movement.
	AccountId int64
	// CategoryId is the id of the category record it points at, indexing it
	// the same way.
	CategoryId int64
}

// Accounts returns every account of the registry, ordered by name.
func Accounts(database keepdeps.KeepDatabase) []Account {
	accounts := []Account{}
	for _, record := range all(database, config.AccountSchema) {
		accounts = append(accounts, Account{
			Name: text(record, config.NameField),
		})
	}
	sort.Slice(accounts, func(i int, j int) bool { return accounts[i].Name < accounts[j].Name })
	return accounts
}

// Categories returns every category of the registry, ordered by name. Its
// packed detail is `name|parent|description`.
func Categories(database keepdeps.KeepDatabase) []Category {
	categories := []Category{}
	for _, record := range all(database, config.CategorySchema) {
		detail := text(record, config.DetailField)
		categories = append(categories, Category{
			Name:        text(record, config.NameField),
			Parent:      utils.Part(detail, config.CategoryParts, config.CategoryParent),
			Description: utils.Part(detail, config.CategoryParts, config.CategoryDescription),
			Revenues:    number(record, config.RevenuesField) == 1,
			Expenses:    number(record, config.ExpensesField) == 1,
		})
	}
	sort.Slice(categories, func(i int, j int) bool { return categories[i].Name < categories[j].Name })
	return categories
}

// Transactions returns every recorded movement, oldest first, ties broken by
// the storage key so the order never depends on how the records came back.
func Transactions(database keepdeps.KeepDatabase) []Transaction {
	transactions := []Transaction{}
	for _, record := range all(database, config.TransactionSchema) {
		transactions = append(transactions, transactionOf(record))
	}
	return sorted(transactions)
}

// AccountTransactions returns every movement dated on one account, oldest
// first. It is read through the account's own nested registry — the ids that
// were written onto it as each movement was recorded — so it costs one record
// plus the movements it names, rather than a walk through the whole ledger.
// An account that is not there reads as no movements.
func AccountTransactions(database keepdeps.KeepDatabase, name string) []Transaction {
	return indexed(database, config.AccountSchema, name)
}

// CategoryTransactions returns every movement classified under one category,
// oldest first, read through the category's own nested registry the way
// AccountTransactions reads an account's. Only the movements naming the
// category itself are returned: a child category indexes its own.
func CategoryTransactions(database keepdeps.KeepDatabase, name string) []Transaction {
	return indexed(database, config.CategorySchema, name)
}

// indexed returns the movements one record of a registry holds in its nested
// TransactionsDB. A link pointing at a movement that is no longer there is
// skipped rather than reported: the listing is what the ledger still holds.
func indexed(database keepdeps.KeepDatabase, schemaName string, name string) []Transaction {
	transactions := []Transaction{}
	instance, ok := database.GetSchema(schemaName)
	if !ok {
		return transactions
	}
	owner, found := instance.FindByKey(config.NameField, name)
	if !found {
		return transactions
	}
	ledger, ok := database.GetSchema(config.TransactionSchema)
	if !ok {
		return transactions
	}
	for _, link := range owner.ListAll(config.TransactionsDB) {
		record, found := ledger.FindById(number(link, config.TransactionId))
		if !found {
			continue
		}
		transactions = append(transactions, transactionOf(record))
	}
	return sorted(transactions)
}

// transactionOf reads one stored movement back as a value. Its packed detail
// is `key|account|category|description`: the names are carried there so a
// listing renders without a lookup per line, while the ids beside them are
// what actually ties the movement to its account and its category.
func transactionOf(record keepdeps.SchemaItem) Transaction {
	detail := text(record, config.DetailField)
	return Transaction{
		Id:          record.Id,
		Key:         text(record, config.NameField),
		Account:     utils.Part(detail, config.TransactionParts, config.TransactionAccount),
		Category:    utils.Part(detail, config.TransactionParts, config.TransactionCategory),
		Description: utils.Part(detail, config.TransactionParts, config.TransactionDescription),
		Amount:      number(record, config.AmountField),
		Date:        number(record, config.DateField),
		AccountId:   number(record, config.AccountID),
		CategoryId:  number(record, config.CategoryID),
	}
}

// sorted orders movements oldest first, ties broken by the storage key.
func sorted(transactions []Transaction) []Transaction {
	sort.Slice(transactions, func(i int, j int) bool {
		if transactions[i].Date != transactions[j].Date {
			return transactions[i].Date < transactions[j].Date
		}
		return transactions[i].Key < transactions[j].Key
	})
	return transactions
}

// all returns every record of one registry, or an empty slice when the
// registry is empty or unreachable. A missing registry reads as no records
// rather than as a failure: a brand-new vault has written nothing yet, and a
// visualization of an empty vault is a page of zeroes, not an error.
func all(database keepdeps.KeepDatabase, schema string) []keepdeps.SchemaItem {
	instance, ok := database.GetSchema(schema)
	if !ok {
		return nil
	}
	records, err := instance.ListAll()
	if err != nil {
		return nil
	}
	return records
}

// text reads a Key field of a record, returning "" when the field is absent
// or does not hold a string.
func text(record keepdeps.SchemaItem, field string) string {
	value, err := record.Get(field)
	if err != nil {
		return ""
	}
	stored, ok := value.(string)
	if !ok {
		return ""
	}
	return stored
}

// number reads an Int field of a record, returning 0 when the field is absent
// or does not hold a whole number.
func number(record keepdeps.SchemaItem, field string) int64 {
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
