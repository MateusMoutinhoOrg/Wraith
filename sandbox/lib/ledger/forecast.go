package ledger

// The forecast: today's position rolled forward through the movements already
// dated ahead of it. It is the only thing in the vault that talks about money
// that has not happened, and it is allowed exactly one ingredient:
// transactions dated after today, which are scheduled facts.
//
// Nothing else is projected. Nothing is inferred from your past and no rule
// is projected on your behalf: a month ahead is worth exactly the movements
// you dated into it. There is no bill to derive and nothing waiting to be
// settled either — every movement moves its whole amount on the day it is
// dated.

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
