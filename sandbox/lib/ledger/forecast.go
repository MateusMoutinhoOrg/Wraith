package ledger

// The forecast: today's position rolled forward through the commitments you
// declared. It is the only thing in the vault that talks about money that has
// not happened, and it is allowed exactly two ingredients:
//
//   - recurrences, which are rules you wrote down,
//   - transactions already dated ahead of today, which are scheduled facts.
//
// Nothing else is projected. There is no bill to derive and nothing waiting
// to be settled: every movement moves its whole amount on the day it is
// dated, so a month ahead is only what you declared for it.

import "github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"

// Projection is one future month of the forecast.
type Projection struct {
	// Month is the month projected, as yyyymm.
	Month int64
	// Held is what the accounts are expected to hold at its end.
	Held int64
	// Income is what is expected to come in during the month.
	Income int64
	// Expenses is what is expected to go out, kept negative.
	Expenses int64
}

// Forecast rolls today's position forward over the given number of months.
// Every month starts where the previous one ended, so the horizon is one walk
// through the calendar rather than a series of independent guesses.
func (s State) Forecast(months int) []Projection {
	held := s.Held()

	projections := []Projection{}
	for step := 1; step <= months; step++ {
		month := utils.AddMonths(s.OpenMonth(), step)
		projection := Projection{Month: month}

		// Movements already recorded and dated into this month settle first:
		// they are facts, not projections.
		for _, transaction := range s.In(month) {
			if transaction.Date <= s.Today {
				continue
			}
			held += transaction.Amount
			record(&projection, transaction.Amount, s.IsTransfer(transaction))
		}

		// Then the commitments in force this month.
		for _, recurrence := range s.Recurrences {
			if !recurrence.AppliesIn(month) {
				continue
			}
			held += recurrence.Amount
			if recurrence.ToAccount != "" {
				held -= recurrence.Amount
			}
			transfer := recurrence.ToAccount != ""
			if category, found := s.Category(recurrence.Category); found && category.IsTransfer() {
				transfer = true
			}
			record(&projection, recurrence.Amount, transfer)
		}

		projection.Held = held
		projections = append(projections, projection)
	}
	return projections
}

// record books one movement into the month's income or expenses, leaving
// transfers out of both — moving money between your own accounts is neither.
func record(projection *Projection, amount int64, transfer bool) {
	if transfer {
		return
	}
	if amount > 0 {
		projection.Income += amount
		return
	}
	projection.Expenses += amount
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
