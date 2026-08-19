package ledger

// The forecast: today's position rolled forward through the commitments you
// declared. It is the only thing in the vault that talks about money that has
// not happened, and it is allowed exactly three ingredients:
//
//   - recurrences, which are rules you wrote down,
//   - transactions already dated ahead of today, which are scheduled facts,
//   - credit card bills, which are derived from what the card already owes.
//
// A card bill is never declared as a recurrence, because its amount changes
// every month. It is derived instead: what a card owes when its cycle closes
// is what leaves your accounts on its due day. That is a projection, and the
// pages that show it say so.

import "github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"

// Projection is one future month of the forecast.
type Projection struct {
	// Month is the month projected, as yyyymm.
	Month int64
	// Held is what the plain accounts are expected to hold at its end.
	Held int64
	// Owed is what the cards are expected to owe at its end.
	Owed int64
	// Income is what is expected to come in during the month.
	Income int64
	// Expenses is what is expected to go out, kept negative.
	Expenses int64
	// Bills is what card bills are expected to take out of the accounts.
	Bills int64
}

// Net is what the month is expected to leave you worth.
func (p Projection) Net() int64 { return p.Held - p.Owed }

// Forecast rolls today's position forward over the given number of months.
// Every month starts where the previous one ended, so the horizon is one walk
// through the calendar rather than a series of independent guesses.
func (s State) Forecast(months int) []Projection {
	held := s.Held()
	owed := map[string]int64{}
	for _, card := range s.Cards() {
		owed[card.Name] = s.Owed(card)
	}

	projections := []Projection{}
	for step := 1; step <= months; step++ {
		month := utils.AddMonths(s.OpenMonth(), step)
		projection := Projection{Month: month}

		// Movements already recorded and dated into this month settle first:
		// they are facts, not projections.
		for _, transaction := range s.Transactions {
			if utils.MonthOf(transaction.PaymentDate) != month {
				continue
			}
			if transaction.PaymentDate <= s.Today {
				continue
			}
			held, owed = settle(s, held, owed, transaction.Account, transaction.Amount)
			held, owed = record(&projection, s, held, owed, transaction.Amount,
				s.IsTransfer(transaction))
		}

		// Then the commitments in force this month.
		for _, recurrence := range s.Recurrences {
			if !recurrence.AppliesIn(month) {
				continue
			}
			held, owed = settle(s, held, owed, recurrence.Account, recurrence.Amount)
			if recurrence.ToAccount != "" {
				held, owed = settle(s, held, owed, recurrence.ToAccount, -recurrence.Amount)
			}
			transfer := recurrence.ToAccount != ""
			if category, found := s.Category(recurrence.Category); found && category.IsTransfer() {
				transfer = true
			}
			held, owed = record(&projection, s, held, owed, recurrence.Amount, transfer)
		}

		// Finally the card bills: what each card owes at the close of its
		// cycle leaves the accounts on its due day.
		for _, card := range s.Cards() {
			bill := owed[card.Name]
			if bill <= 0 {
				continue
			}
			owed[card.Name] = 0
			held -= bill
			projection.Bills += bill
		}

		projection.Held = held
		for _, card := range s.Cards() {
			projection.Owed += owed[card.Name]
		}
		projections = append(projections, projection)
	}
	return projections
}

// settle applies one movement to the running position: money on a plain
// account changes what you hold, money on a card changes what you owe.
func settle(s State, held int64, owed map[string]int64, account string, amount int64) (int64, map[string]int64) {
	stored, found := s.Account(account)
	if found && stored.IsCard() {
		owed[account] -= amount
		return held, owed
	}
	return held + amount, owed
}

// record books one movement into the month's income or expenses, leaving
// transfers out of both — moving money between your own accounts is neither.
func record(projection *Projection, s State, held int64, owed map[string]int64, amount int64, transfer bool) (int64, map[string]int64) {
	if transfer {
		return held, owed
	}
	if amount > 0 {
		projection.Income += amount
		return held, owed
	}
	projection.Expenses += amount
	return held, owed
}

// DueIn returns the commitments falling in one month, as the dated lines a
// month page lists.
type Due struct {
	// Date is the day it falls on, as yyyymmdd.
	Date int64
	// Description identifies it.
	Description string
	// Account is where it lands.
	Account string
	// Amount is what it is worth, in cents.
	Amount int64
}

// DueIn returns every recurrence falling in one month, as dated lines, oldest
// first.
func (s State) DueIn(month int64) []Due {
	due := []Due{}
	for _, recurrence := range s.Recurrences {
		if !recurrence.AppliesIn(month) {
			continue
		}
		due = append(due, Due{
			Date:        utils.DateIn(month, recurrence.Day),
			Description: recurrence.Description,
			Account:     recurrence.Account,
			Amount:      recurrence.Amount,
		})
	}
	for index := 1; index < len(due); index++ {
		for back := index; back > 0 && due[back].Date < due[back-1].Date; back-- {
			due[back], due[back-1] = due[back-1], due[back]
		}
	}
	return due
}
