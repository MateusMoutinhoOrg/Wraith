package store

// The database layout of the brain: the five registries every task writes
// and every visualization reads, plus the helpers both sides reach them
// with. Everything here is expressed against the embedded Keep library
// injected as Deps.KeepLib, so the sandbox stays closed — no adapter, no
// third-party module, no OS-bound package.
//
// Keep stores unique string keys and whole numbers, and nothing else. Three
// consequences shape every schema below, and they are the whole trick worth
// knowing before adding a registry of your own:
//
//   - Money is held in cents, as a whole number. Money renders it back.
//   - Dates are held as whole numbers: 2026-08-18 is 20260818, 2026-08 is
//     202608. They sort correctly as numbers, which is what a month index
//     is built out of.
//   - Free text that is not unique — a description, an account's name on a
//     transaction — cannot be a key of its own, so it travels packed into
//     one key alongside a value that *is* unique. Pack writes it, Unpack
//     reads it back.
//
// This package declares no types; the tasks in sandbox/Tasks/Tasks and the
// visualizations in sandbox/Visualization/Visualization call into it.

import (
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
)

// The registries the brain keeps, one Keep schema each.
const (
	// AccountSchema holds accounts and credit cards alike — a card is an
	// account whose Kind is KindCard.
	AccountSchema = "account"
	// CategorySchema holds the categories transactions are classified under.
	CategorySchema = "category"
	// TransactionSchema holds every recorded movement.
	TransactionSchema = "transaction"
	// RecurrenceSchema holds the commitments the forecast projects.
	RecurrenceSchema = "recurrence"
)

// The fields the registries carry. A field named Detail is a packed key; see
// Pack.
const (
	// NameField is a record's unique name.
	NameField = "name"
	// DetailField is the packed key carrying a record's free text.
	DetailField = "detail"
	// KindField is an account's kind, KindAccount or KindCard.
	KindField = "kind"
	// OpeningField is an account's opening balance, in cents.
	OpeningField = "opening"
	// LimitField is a credit card's total limit, in cents.
	LimitField = "limit"
	// ClosingDayField is the day of the month a card's bill closes.
	ClosingDayField = "closingDay"
	// DueDayField is the day of the month a card's bill is due.
	DueDayField = "dueDay"
	// RevenuesField is 1 when a category accepts positive amounts.
	RevenuesField = "revenues"
	// ExpensesField is 1 when a category accepts negative amounts.
	ExpensesField = "expenses"
	// AmountField is a value in cents, signed.
	AmountField = "amount"
	// DateField is a date as yyyymmdd.
	DateField = "date"
	// PaymentDateField is the date the money actually moves, as yyyymmdd.
	PaymentDateField = "paymentDate"
	// DayField is the day of the month a recurrence falls on.
	DayField = "day"
	// StartField is the first month a recurrence applies, as yyyymm.
	StartField = "start"
	// EndField is the last month a recurrence applies, as yyyymm, or 0 for
	// open-ended.
	EndField = "end"
)

// Account kinds, held in KindField.
const (
	// KindAccount is a plain account — cash, a bank, a savings pot.
	KindAccount = 0
	// KindCard is a credit card, which carries a limit and a billing cycle.
	KindCard = 1
)

// Separator splits the parts of a packed key. It is a character no name,
// date or amount can hold, so unpacking is never ambiguous — a description
// carrying one is rejected when the task validates its fields.
const Separator = "|"

// Props describes the brain's database. path is the folder the registries
// live in, relative to the vault root: it is api.Lib.DatabasePath, which the
// interface's `--database` flag overrides, so pointing the same binary at a
// second vault changes nothing but this prefix.
func Props(path string) keepdeps.Props {
	return keepdeps.Props{
		Path: strings.TrimSuffix(path, "/") + "/",
		Schemas: []keepdeps.Schema{
			{
				Name: AccountSchema,
				Itens: []keepdeps.Item{
					{Name: NameField, Type: keepdeps.Key, Required: true},
					{Name: KindField, Type: keepdeps.Int, Required: true},
					{Name: OpeningField, Type: keepdeps.Int, Required: true},
					{Name: LimitField, Type: keepdeps.Int},
					{Name: ClosingDayField, Type: keepdeps.Int},
					{Name: DueDayField, Type: keepdeps.Int},
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
					{Name: PaymentDateField, Type: keepdeps.Int, Required: true},
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

// Database opens the brain's database over the injected Keep library.
func Database(d deps.Deps, path string) keepdeps.KeepDatabase {
	return d.KeepLib.NewDatabase(Props(path))
}

// Schema returns one registry of a database. ok is false when the injected
// library hands back no such schema, which can only happen if Props and the
// name asked for have drifted apart.
func Schema(database keepdeps.KeepDatabase, name string) (keepdeps.SchemaInstance, bool) {
	return database.GetSchema(name)
}

// All returns every record of one registry, or an empty slice when the
// registry is empty or unreachable. A missing registry reads as no records
// rather than as a failure: a brand-new vault has written nothing yet, and a
// visualization of an empty vault is a page of zeroes, not an error.
func All(database keepdeps.KeepDatabase, schema string) []keepdeps.SchemaItem {
	instance, ok := Schema(database, schema)
	if !ok {
		return nil
	}
	records, err := instance.ListAll()
	if err != nil {
		return nil
	}
	return records
}

// Find returns the record of one registry carrying the given name. ok is
// false when no record does.
func Find(database keepdeps.KeepDatabase, schema string, name string) (keepdeps.SchemaItem, bool) {
	instance, ok := Schema(database, schema)
	if !ok {
		return keepdeps.SchemaItem{}, false
	}
	return instance.FindByKey(NameField, name)
}

// Exists reports whether one registry already carries a record under the
// given name.
func Exists(database keepdeps.KeepDatabase, schema string, name string) bool {
	_, found := Find(database, schema, name)
	return found
}

// Text reads a Key field of a record, returning "" when the field is absent
// or does not hold a string.
func Text(record keepdeps.SchemaItem, field string) string {
	value, err := record.Get(field)
	if err != nil {
		return ""
	}
	text, ok := value.(string)
	if !ok {
		return ""
	}
	return text
}

// Number reads an Int field of a record, returning 0 when the field is
// absent or does not hold a whole number.
func Number(record keepdeps.SchemaItem, field string) int64 {
	value, err := record.Get(field)
	if err != nil {
		return 0
	}
	switch typed := value.(type) {
	case int64:
		return typed
	case int:
		return int64(typed)
	}
	return 0
}

// Pack composes one key out of several parts, so free text that is not
// unique can be stored beside a part that is. The first part is what makes
// the whole key unique, and the last part is the only one allowed to contain
// the separator — Unpack hands it back whole.
func Pack(parts ...string) string {
	return strings.Join(parts, Separator)
}

// Unpack splits a key composed by Pack back into count parts, padding with
// empty strings when the stored key carries fewer. The last part keeps
// everything that is left, separators included.
func Unpack(key string, count int) []string {
	parts := strings.SplitN(key, Separator, count)
	for len(parts) < count {
		parts = append(parts, "")
	}
	return parts
}

// Detail reads one part of a record's packed DetailField.
func Detail(record keepdeps.SchemaItem, count int, index int) string {
	parts := Unpack(Text(record, DetailField), count)
	if index < 0 || index >= len(parts) {
		return ""
	}
	return parts[index]
}
