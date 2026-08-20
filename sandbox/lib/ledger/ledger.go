package ledger

// Every figure the vault shows, derived from the five registries and nothing
// else. If a number cannot be computed from accounts, cards, categories,
// transactions and recurrences, it does not belong on a page — so this
// package is where the arithmetic lives, and the visualizations next door are
// only about layout.
//
// Nothing here writes. A visualization is handed the same database a task
// writes to, and reads it through the functions below.
//
// This package declares the State type and plain functions over it; it is
// neither an object package nor a lib function package, so it carries no
// factories.

import (
	"sort"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// State is the whole registry, read once, plus the day it is being read on.
// Every visualization loads it exactly once and renders its whole tree from
// it, so no two pages of one render can disagree about what the data says.
type State struct {
	// Today is the date the render happens on, taken from Deps.Now.
	Today int64
	// Accounts is every account and card, ordered by name.
	Accounts []Account
	// Categories is every category, ordered by name.
	Categories []Category
	// Transactions is every recorded movement, oldest first.
	Transactions []Transaction
	// Recurrences is every declared commitment, ordered by description.
	Recurrences []Recurrence
}

// Load reads the whole registry into a State.
func Load(d deps.Deps, database keepdeps.KeepDatabase) State {
	return State{
		Today:        utils.DateOf(d.Now()),
		Accounts:     Accounts(database),
		Categories:   Categories(database),
		Transactions: Transactions(database),
		Recurrences:  Recurrences(database),
	}
}

// OpenMonth is the month today falls in — the one the vault is currently
// writing into.
func (s State) OpenMonth() int64 { return utils.MonthOf(s.Today) }

// Cards returns every credit card of the registry.
func (s State) Cards() []Account {
	cards := []Account{}
	for _, account := range s.Accounts {
		if account.IsCard() {
			cards = append(cards, account)
		}
	}
	return cards
}

// PlainAccounts returns every account that is not a credit card — the ones
// money actually sits in.
func (s State) PlainAccounts() []Account {
	accounts := []Account{}
	for _, account := range s.Accounts {
		if !account.IsCard() {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

// BalanceOn returns what an account holds on a given date: what it opened
// with, plus every movement that has actually settled by then. A transaction
// whose payment_date is still ahead counts in its month's result but not yet
// in the balance, which is exactly the difference between what a month cost
// and what has left the account.
func (s State) BalanceOn(account Account, date int64) int64 {
	balance := account.Opening
	for _, transaction := range s.Transactions {
		if transaction.Account != account.Name {
			continue
		}
		if transaction.PaymentDate > date {
			continue
		}
		balance += transaction.Amount
	}
	return balance
}

// Balance returns what an account holds today.
func (s State) Balance(account Account) int64 {
	return s.BalanceOn(account, s.Today)
}

// OwedOn returns what is outstanding on a credit card on a given date, as a
// positive figure — a card's balance is negative when money is owed on it.
func (s State) OwedOn(card Account, date int64) int64 {
	balance := s.BalanceOn(card, date)
	if balance > 0 {
		return 0
	}
	return -balance
}

// Owed returns what is outstanding on a credit card today.
func (s State) Owed(card Account) int64 { return s.OwedOn(card, s.Today) }

// HeldOn is the total of every plain account on a given date.
func (s State) HeldOn(date int64) int64 {
	total := int64(0)
	for _, account := range s.PlainAccounts() {
		total += s.BalanceOn(account, date)
	}
	return total
}

// Held is the total of every plain account — the money you actually hold.
func (s State) Held() int64 { return s.HeldOn(s.Today) }

// TotalOwedOn is the total outstanding across every credit card on a date.
func (s State) TotalOwedOn(date int64) int64 {
	total := int64(0)
	for _, card := range s.Cards() {
		total += s.OwedOn(card, date)
	}
	return total
}

// TotalOwed is the total outstanding across every credit card today.
func (s State) TotalOwed() int64 { return s.TotalOwedOn(s.Today) }

// NetOn is what you hold minus what you owe on a given date.
func (s State) NetOn(date int64) int64 { return s.HeldOn(date) - s.TotalOwedOn(date) }

// Net is what you hold minus what you owe.
func (s State) Net() int64 { return s.NetOn(s.Today) }

// Horizon is the last date the registry settles anything on: the payment date
// of the furthest movement already recorded, or today when nothing is dated
// ahead. It is where a purchase in 12x stops costing you.
func (s State) Horizon() int64 {
	horizon := s.Today
	for _, transaction := range s.Transactions {
		if transaction.PaymentDate > horizon {
			horizon = transaction.PaymentDate
		}
	}
	return horizon
}

// NetWorth is everything you have minus everything you owe, once every
// movement already recorded has settled — the last future entry included. It
// differs from Net by exactly what is still pending: Net is where you stand
// today, NetWorth is where what you have already committed to leaves you.
//
// Recurrences are left out of it on purpose. A recurrence is a rule, not a
// movement: it is the forecast that projects it, and this figure only counts
// what the ledger already holds.
func (s State) NetWorth() int64 { return s.NetOn(s.Horizon()) }

// Pending is the sum of movements recorded but not yet settled — everything
// whose payment_date is still ahead of today.
func (s State) Pending() int64 {
	total := int64(0)
	for _, transaction := range s.Transactions {
		if transaction.PaymentDate > s.Today {
			total += transaction.Amount
		}
	}
	return total
}

// PendingIncome is the sum of positive movements not yet settled — money you
// are still waiting to receive.
func (s State) PendingIncome() int64 {
	total := int64(0)
	for _, transaction := range s.Transactions {
		if transaction.PaymentDate > s.Today && transaction.Amount > 0 {
			total += transaction.Amount
		}
	}
	return total
}

// PendingExpenses is the sum of negative movements not yet settled — money
// you are still going to pay. The figure is kept negative.
func (s State) PendingExpenses() int64 {
	total := int64(0)
	for _, transaction := range s.Transactions {
		if transaction.PaymentDate > s.Today && transaction.Amount < 0 {
			total += transaction.Amount
		}
	}
	return total
}

// Result is what one month came to: what came in, what went out, the two
// added together, and how many movements it holds.
type Result struct {
	// Income is everything positive dated in the month.
	Income int64
	// Expenses is everything negative dated in the month, kept negative.
	Expenses int64
	// Count is how many movements the month holds.
	Count int
}

// Total is income and expenses added together — the month's bottom line.
func (r Result) Total() int64 { return r.Income + r.Expenses }

// MonthResult returns what one month came to. Transfers between your own
// accounts are left out on both sides: moving money is neither earning it nor
// spending it.
func (s State) MonthResult(month int64) Result {
	result := Result{}
	for _, transaction := range s.In(month) {
		result.Count++
		if s.IsTransfer(transaction) {
			continue
		}
		if transaction.Amount > 0 {
			result.Income += transaction.Amount
			continue
		}
		result.Expenses += transaction.Amount
	}
	return result
}

// In returns every movement dated in one month.
func (s State) In(month int64) []Transaction {
	found := []Transaction{}
	for _, transaction := range s.Transactions {
		if utils.MonthOf(transaction.Date) == month {
			found = append(found, transaction)
		}
	}
	return found
}

// OfAccount narrows a list of movements down to one account.
func OfAccount(transactions []Transaction, account string) []Transaction {
	found := []Transaction{}
	for _, transaction := range transactions {
		if transaction.Account == account {
			found = append(found, transaction)
		}
	}
	return found
}

// Flow is what a set of movements moved: what came in, what went out, and how
// many movements it took. It is what a Result is to a month, for one account —
// with the difference that a transfer counts on both sides here, because money
// leaving one of your own accounts for another has still left that account.
type Flow struct {
	// In is the total of every positive amount.
	In int64
	// Out is the total of every negative amount, kept negative.
	Out int64
	// Count is how many movements the flow holds.
	Count int
}

// Net is what came in and what went out added together.
func (f Flow) Net() int64 { return f.In + f.Out }

// FlowOf returns what a list of movements moved.
func FlowOf(transactions []Transaction) Flow {
	flow := Flow{}
	for _, transaction := range transactions {
		flow.Count++
		if transaction.Amount > 0 {
			flow.In += transaction.Amount
			continue
		}
		flow.Out += transaction.Amount
	}
	return flow
}

// AccountFlow returns what one account moved in one month, or across the whole
// ledger when month is zero.
func (s State) AccountFlow(name string, month int64) Flow {
	if month == 0 {
		return FlowOf(OfAccount(s.Transactions, name))
	}
	return FlowOf(OfAccount(s.In(month), name))
}

// AccountMonths returns every month holding at least one movement on one
// account, oldest first. Like Months, a month is never created by hand: it
// exists as soon as a movement on that account carries a date inside it.
func (s State) AccountMonths(name string) []int64 {
	seen := map[int64]bool{}
	months := []int64{}
	for _, transaction := range OfAccount(s.Transactions, name) {
		month := utils.MonthOf(transaction.Date)
		if seen[month] {
			continue
		}
		seen[month] = true
		months = append(months, month)
	}
	sort.Slice(months, func(i int, j int) bool { return months[i] < months[j] })
	return months
}

// PendingOf returns what is recorded on one account but has not settled yet —
// everything on it whose payment_date is still ahead of today.
func (s State) PendingOf(account Account) int64 {
	total := int64(0)
	for _, transaction := range OfAccount(s.Transactions, account.Name) {
		if transaction.PaymentDate > s.Today {
			total += transaction.Amount
		}
	}
	return total
}

// IsTransfer reports whether a movement is a leg of a transfer between your
// own accounts, and therefore counts as neither income nor expense.
func (s State) IsTransfer(transaction Transaction) bool {
	category, found := s.Category(transaction.Category)
	return found && category.IsTransfer()
}

// Category returns one category by name.
func (s State) Category(name string) (Category, bool) {
	for _, category := range s.Categories {
		if category.Name == name {
			return category, true
		}
	}
	return Category{}, false
}

// Account returns one account by name.
func (s State) Account(name string) (Account, bool) {
	for _, account := range s.Accounts {
		if account.Name == name {
			return account, true
		}
	}
	return Account{}, false
}

// Months returns every month holding at least one movement, oldest first. A
// month is never created by hand: it exists as soon as a transaction carries
// a date inside it.
func (s State) Months() []int64 {
	seen := map[int64]bool{}
	months := []int64{}
	for _, transaction := range s.Transactions {
		month := utils.MonthOf(transaction.Date)
		if seen[month] {
			continue
		}
		seen[month] = true
		months = append(months, month)
	}
	sort.Slice(months, func(i int, j int) bool { return months[i] < months[j] })
	return months
}

// RenderedMonths returns the months a dashboard writes a folder for: the
// months holding movements, from `prev-months` before the open one onwards.
// Months holding nothing are skipped, and months ahead of the open one are
// kept — a purchase in 12x opens the next eleven.
func (s State) RenderedMonths(previous int) []int64 {
	if previous < 0 {
		previous = 0
	}
	earliest := utils.AddMonths(s.OpenMonth(), -previous)
	months := []int64{}
	for _, month := range s.Months() {
		if month < earliest {
			continue
		}
		months = append(months, month)
	}
	return months
}

// CategoryTotal returns what one category came to in one month, or across the
// whole ledger when month is zero.
func (s State) CategoryTotal(name string, month int64) int64 {
	total := int64(0)
	for _, transaction := range s.Transactions {
		if !s.IsCategoryOrDescendant(transaction.Category, name) {
			continue
		}
		if month != 0 && utils.MonthOf(transaction.Date) != month {
			continue
		}
		total += transaction.Amount
	}
	return total
}

// CategoryCount returns how many movements one category classifies.
func (s State) CategoryCount(name string) int {
	count := 0
	for _, transaction := range s.Transactions {
		if s.IsCategoryOrDescendant(transaction.Category, name) {
			count++
		}
	}
	return count
}

// IsCategoryOrDescendant reports whether category 'child' is the same as 'parent' or a descendant of it.
func (s State) IsCategoryOrDescendant(child string, parent string) bool {
	if child == parent {
		return true
	}
	for {
		cat, found := s.Category(child)
		if !found || cat.Parent == "" {
			return false
		}
		if cat.Parent == parent {
			return true
		}
		child = cat.Parent
	}
}
