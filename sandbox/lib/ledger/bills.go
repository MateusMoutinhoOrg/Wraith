package ledger

// Credit card statements: the bill one billing cycle produces, what has been
// paid against it, and what is still pending. Everything here is derived from
// the account and transaction registries and nothing else — a bill is not a
// record, it is a reading of the ledger through the card's two days.
//
// A card carries a closing day and a due day, and those two days cut the
// ledger into cycles. The **charges** of a cycle are the purchases dated
// after the previous close, up to and including this one; the cycle's bill
// falls due on the due day that follows it.
//
// Money arriving on a card pays its **oldest unsettled statement first**, and
// what is left over runs on to the next one. That is the one rule worth
// knowing here: a payment is not matched to a statement by its date, because
// one payment routinely settles two bills at once and a person paying "what
// the card is asking for" is paying the oldest thing first. What overflows
// every statement a card has issued sits as a credit on the one it is
// currently charging to.
//
// The balance a card was declared with — what it already owed before it had a
// statement here — is counted into its earliest cycle, so the remainders of a
// card's statements always add back up to what the card owes once everything
// recorded on it has settled.

import (
	"sort"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// The states a statement can be in, as the pages word them.
const (
	// BillAhead is a cycle that has not begun collecting charges yet: it
	// exists because something already recorded is dated inside it, and it
	// asks for nothing until it closes.
	BillAhead = "ahead"
	// BillOpen is the cycle a card is charging to now. It has not closed yet,
	// so nothing is owed on it until it does.
	BillOpen = "open"
	// BillDue is a closed statement still carrying a remainder, whose due
	// date has not passed.
	BillDue = "to pay"
	// BillOverdue is a closed statement still carrying a remainder past its
	// due date.
	BillOverdue = "overdue"
	// BillPaid is a closed statement nothing is left on.
	BillPaid = "paid"
)

// Settlement is one payment applied to one statement: the movement that
// arrived on the card, and how much of it this statement took. A payment
// large enough to settle two bills appears once on each, carrying the part
// that landed there.
type Settlement struct {
	// Transaction is the movement that arrived on the card.
	Transaction Transaction
	// Applied is how much of it this statement took, in cents, as a positive
	// figure.
	Applied int64
}

// Bill is one statement of one credit card: the window it covers, the
// purchases it holds, and the payments applied to it.
type Bill struct {
	// Card is the name of the card the statement belongs to.
	Card string
	// Cycle is the month the statement closes in, as yyyymm.
	Cycle int64
	// After is the previous cycle's closing date: a purchase counts on this
	// statement when it is dated after it.
	After int64
	// Closes is the date the statement closes on, as yyyymmdd.
	Closes int64
	// Due is the date the statement falls due on, as yyyymmdd.
	Due int64
	// Carried is what the card already owed before its first statement, in
	// cents and therefore negative. It is zero on every statement but the
	// earliest one.
	Carried int64
	// Purchases are the charges the cycle holds, oldest first.
	Purchases []Transaction
	// Settlements are the payments applied to the statement, oldest first.
	Settlements []Settlement
}

// Charges is the total of the purchases the statement holds, kept negative.
func (b Bill) Charges() int64 {
	total := int64(0)
	for _, purchase := range b.Purchases {
		total += purchase.Amount
	}
	return total
}

// Paid is the total of the payments applied to the statement, kept positive.
func (b Bill) Paid() int64 {
	total := int64(0)
	for _, settlement := range b.Settlements {
		total += settlement.Applied
	}
	return total
}

// Total is what the statement asks for, as a positive figure.
func (b Bill) Total() int64 { return -(b.Carried + b.Charges()) }

// Remaining is what the statement still asks for once the payments applied to
// it are taken off. It is negative when more has arrived on the card than
// every statement it has issued asks for — a credit, not a debt.
func (b Bill) Remaining() int64 { return b.Total() - b.Paid() }

// IsClosed reports whether the cycle has closed by the given date, and is
// therefore a bill rather than a statement still being written.
func (b Bill) IsClosed(today int64) bool { return today > b.Closes }

// IsSettled reports whether nothing is left on the statement.
func (b Bill) IsSettled() bool { return b.Remaining() <= 0 }

// IsEmpty reports whether the cycle holds nothing at all — no purchase, no
// payment and no balance carried in.
func (b Bill) IsEmpty() bool {
	return len(b.Purchases) == 0 && len(b.Settlements) == 0 && b.Carried == 0
}

// Status reports where the statement stands on a given day, in the words the
// pages show it with.
func (b Bill) Status(today int64) string {
	if today <= b.After {
		return BillAhead
	}
	if !b.IsClosed(today) {
		return BillOpen
	}
	if b.IsSettled() {
		return BillPaid
	}
	if today > b.Due {
		return BillOverdue
	}
	return BillDue
}

// CardClosingDay is the day of the month a card's cycle closes on, never
// outside the calendar. A card is declared with a day from 1 to 31, so the
// clamp only ever matters for a record written before the field existed.
func CardClosingDay(card Account) int64 { return billDay(card.ClosingDay) }

// CardDueDay is the day of the month a card's bill falls due on.
func CardDueDay(card Account) int64 { return billDay(card.DueDay) }

// billDay keeps a declared day of the month inside the calendar.
func billDay(day int64) int64 {
	if day < 1 {
		return 1
	}
	if day > 31 {
		return 31
	}
	return day
}

// CardCloses returns the date one cycle of a card closes on, clamped to the
// month's last day when the closing day does not exist in it.
func CardCloses(card Account, cycle int64) int64 {
	return utils.DateIn(cycle, CardClosingDay(card))
}

// CardDue returns the date one cycle's bill falls due on: the due day of the
// cycle's own month, or of the next one when the card is due before it
// closes.
func CardDue(card Account, cycle int64) int64 {
	month := cycle
	if CardDueDay(card) < CardClosingDay(card) {
		month = utils.AddMonths(cycle, 1)
	}
	return utils.DateIn(month, CardDueDay(card))
}

// CardCycle returns the cycle a date is charged to on a card: its own month
// when it falls on or before that month's close, the next one when it falls
// after it.
func CardCycle(card Account, date int64) int64 {
	month := utils.MonthOf(date)
	if date <= CardCloses(card, month) {
		return month
	}
	return utils.AddMonths(month, 1)
}

// OpenCycle is the cycle a card is currently charging to — the one today
// falls in.
func (s State) OpenCycle(card Account) int64 { return CardCycle(card, s.Today) }

// CardCycles returns every cycle a card has a statement for, oldest first:
// the cycle of every movement on it, and the one being charged to today. A
// cycle is never created by hand — it exists because something on the card is
// dated inside it, which is why a purchase split into twelve parts opens the
// next eleven statements.
func (s State) CardCycles(card Account) []int64 {
	seen := map[int64]bool{s.OpenCycle(card): true}
	for _, transaction := range OfAccount(s.Transactions, card.Name) {
		seen[CardCycle(card, transaction.Date)] = true
	}
	cycles := []int64{}
	for cycle := range seen {
		cycles = append(cycles, cycle)
	}
	sort.Slice(cycles, func(i int, j int) bool { return cycles[i] < cycles[j] })
	return cycles
}

// Bills returns the statements of one card, oldest first. A cycle holding
// nothing is left out, except the one being charged to today: an empty open
// statement is still where the next purchase lands.
func (s State) Bills(card Account) []Bill {
	open := s.OpenCycle(card)
	bills := []Bill{}
	for _, bill := range s.allBills(card) {
		if bill.IsEmpty() && bill.Cycle != open {
			continue
		}
		bills = append(bills, bill)
	}
	return bills
}

// Bill returns one statement of one card, empty when the cycle holds nothing.
func (s State) Bill(card Account, cycle int64) Bill {
	for _, bill := range s.allBills(card) {
		if bill.Cycle == cycle {
			return bill
		}
	}
	return s.emptyBill(card, cycle)
}

// OpenBill returns the statement a card is still charging to.
func (s State) OpenBill(card Account) Bill { return s.Bill(card, s.OpenCycle(card)) }

// allBills returns every cycle of a card as a statement, oldest first, with
// the purchases charged to it and the payments applied to it. It is the one
// place a card's statements are computed: the cycles are filled with their
// charges first, and the money that arrived on the card is then run through
// them oldest first, which is what makes one payment able to settle two
// bills.
func (s State) allBills(card Account) []Bill {
	movements := OfAccount(s.Transactions, card.Name)
	bills := []Bill{}
	for _, cycle := range s.CardCycles(card) {
		bill := s.emptyBill(card, cycle)
		if len(bills) == 0 {
			bill.Carried = card.Opening
		}
		for _, transaction := range movements {
			if transaction.Amount >= 0 {
				continue
			}
			if transaction.Date <= bill.After || transaction.Date > bill.Closes {
				continue
			}
			bill.Purchases = append(bill.Purchases, transaction)
		}
		bills = append(bills, bill)
	}
	return applyPayments(bills, movements)
}

// applyPayments runs the money that arrived on a card through its statements,
// oldest first: each payment fills what the oldest statement still asks for,
// and what is left of it runs on to the next one. A statement whose cycle had
// not begun on the day of the payment is stepped over, and anything left once
// every statement is full sits as a credit on the last one — so the
// remainders always add back up to what the card owes.
func applyPayments(bills []Bill, movements []Transaction) []Bill {
	if len(bills) == 0 {
		return bills
	}
	for _, payment := range movements {
		if payment.Amount <= 0 {
			continue
		}
		left := payment.Amount
		for index := range bills {
			if left == 0 {
				break
			}
			if payment.Date <= bills[index].After {
				break
			}
			room := bills[index].Remaining()
			if room <= 0 {
				continue
			}
			applied := room
			if left < applied {
				applied = left
			}
			bills[index].Settlements = append(bills[index].Settlements,
				Settlement{Transaction: payment, Applied: applied})
			left -= applied
		}
		if left == 0 {
			continue
		}
		last := len(bills) - 1
		bills[last].Settlements = append(bills[last].Settlements,
			Settlement{Transaction: payment, Applied: left})
	}
	return bills
}

// emptyBill is one cycle of a card with its dates filled in and nothing on it
// yet.
func (s State) emptyBill(card Account, cycle int64) Bill {
	return Bill{
		Card:   card.Name,
		Cycle:  cycle,
		After:  CardCloses(card, utils.AddMonths(cycle, -1)),
		Closes: CardCloses(card, cycle),
		Due:    CardDue(card, cycle),
	}
}

// UnpaidBillsOf returns the closed statements of one card that still carry a
// remainder, oldest first.
func (s State) UnpaidBillsOf(card Account) []Bill {
	unpaid := []Bill{}
	for _, bill := range s.Bills(card) {
		if bill.IsClosed(s.Today) && !bill.IsSettled() {
			unpaid = append(unpaid, bill)
		}
	}
	return unpaid
}

// UnpaidBills returns every closed statement still carrying a remainder,
// across every card, soonest due first — the bills you actually have to pay.
func (s State) UnpaidBills() []Bill {
	unpaid := []Bill{}
	for _, card := range s.Cards() {
		unpaid = append(unpaid, s.UnpaidBillsOf(card)...)
	}
	sort.Slice(unpaid, func(i int, j int) bool {
		if unpaid[i].Due != unpaid[j].Due {
			return unpaid[i].Due < unpaid[j].Due
		}
		return unpaid[i].Card < unpaid[j].Card
	})
	return unpaid
}

// AmountDue is what one card's closed statements still ask for. It is what
// the PayCreditCardBill task settles when it is not given an amount.
func (s State) AmountDue(card Account) int64 {
	total := int64(0)
	for _, bill := range s.UnpaidBillsOf(card) {
		total += bill.Remaining()
	}
	return total
}

// TotalAmountDue is what every card's closed statements still ask for.
func (s State) TotalAmountDue() int64 {
	total := int64(0)
	for _, card := range s.Cards() {
		total += s.AmountDue(card)
	}
	return total
}

// TotalOverdue is what the statements past their due date still ask for.
func (s State) TotalOverdue() int64 {
	total := int64(0)
	for _, bill := range s.UnpaidBills() {
		if bill.Status(s.Today) != BillOverdue {
			continue
		}
		total += bill.Remaining()
	}
	return total
}

// TotalOpenBills is what the statements still being written add up to — what
// the cards have collected since they last closed, and will ask for when they
// do.
func (s State) TotalOpenBills() int64 {
	total := int64(0)
	for _, card := range s.Cards() {
		remaining := s.OpenBill(card).Remaining()
		if remaining <= 0 {
			continue
		}
		total += remaining
	}
	return total
}

// Unsettled returns every movement already recorded whose payment date is
// still ahead of today, soonest first — everything the ledger knows it is
// going to move and has not moved yet.
func (s State) Unsettled() []Transaction {
	pending := []Transaction{}
	for _, transaction := range s.Transactions {
		if transaction.PaymentDate > s.Today {
			pending = append(pending, transaction)
		}
	}
	sort.Slice(pending, func(i int, j int) bool {
		if pending[i].PaymentDate != pending[j].PaymentDate {
			return pending[i].PaymentDate < pending[j].PaymentDate
		}
		return pending[i].Key < pending[j].Key
	})
	return pending
}

// PendingAccountExpenses is the sum of unsettled movements out of your plain
// accounts, kept negative. It leaves the cards out: what a card has charged
// leaves your accounts as part of a statement, so counting both would count
// it twice.
func (s State) PendingAccountExpenses() int64 {
	total := int64(0)
	for _, transaction := range s.Unsettled() {
		if transaction.Amount >= 0 || s.isCardMovement(transaction) {
			continue
		}
		total += transaction.Amount
	}
	return total
}

// PendingAccountIncome is the sum of unsettled movements into your plain
// accounts, kept positive.
func (s State) PendingAccountIncome() int64 {
	total := int64(0)
	for _, transaction := range s.Unsettled() {
		if transaction.Amount <= 0 || s.isCardMovement(transaction) {
			continue
		}
		total += transaction.Amount
	}
	return total
}

// PendingCardCharges is the sum of unsettled charges recorded on the cards,
// kept negative — the installments already written into statements that have
// not closed yet.
func (s State) PendingCardCharges() int64 {
	total := int64(0)
	for _, transaction := range s.Unsettled() {
		if transaction.Amount >= 0 || !s.isCardMovement(transaction) {
			continue
		}
		total += transaction.Amount
	}
	return total
}

// DueFromAccounts is what is already known to leave your accounts, as a
// positive figure: the remainder of every closed statement, plus every
// unsettled expense recorded on an account.
func (s State) DueFromAccounts() int64 {
	return s.TotalAmountDue() - s.PendingAccountExpenses()
}

// isCardMovement reports whether a movement was recorded on a credit card
// rather than on an account money sits in.
func (s State) isCardMovement(transaction Transaction) bool {
	account, found := s.Account(transaction.Account)
	return found && account.IsCard()
}
