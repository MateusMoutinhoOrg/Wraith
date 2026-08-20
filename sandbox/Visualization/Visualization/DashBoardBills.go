package visualizations

// The bill tree of the DashBoard visualization: one page per credit card
// showing every statement it has produced — what each one charged, what has
// been paid against it, and what is still owed — plus the one page that
// answers "what do I still have to pay?" for the whole vault.
//
// A statement is not a record. It is the ledger read through the card's two
// days, so nothing here is written by a task: paying a bill writes movements,
// and the cycles they land in are computed. The arithmetic of that lives in
// sandbox/lib/ledger/bills.go; these pages are layout.

import (
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// billsPath is where one card's statements page is written, below the entry's
// dest. Every link to it anywhere in the tree is built from this, so the
// folder is named in one place.
func billsPath(card ledger.Account) string {
	return "Bills/" + slug(card.Name) + ".md"
}

// billsPage writes Bills/<card>.md: every statement of one card, oldest cycle
// last, with what it charged, what answered it, and what is left on it.
func billsPage(state ledger.State, card ledger.Account) api.VisualizationRender {
	p := &page{}
	open := state.OpenBill(card)
	p.heading(1, card.Name+" — Bills")
	p.line("> **Credit card** · closes day " +
		strconv.FormatInt(ledger.CardClosingDay(card), 10) + ", due day " +
		strconv.FormatInt(ledger.CardDueDay(card), 10) + " · **Updated:** " +
		utils.PrettyDate(state.Today))
	p.blank()
	p.line("[" + card.Name + "](../" + accountPath(card) + ") · " + navigationAt("../"))
	p.blank()
	p.rule()

	billsPosition(p, state, card, open)
	billsHistory(p, state, card)
	billsOpen(p, state, open)
	billsPaid(p, state, card)
	return p.render(billsPath(card))
}

// billsPosition writes section 1: what the card is asking for right now, and
// what it is still collecting.
func billsPosition(p *page, state ledger.State, card ledger.Account, open ledger.Bill) {
	owed := state.Owed(card)
	p.heading(2, "1. What this card is asking for")
	p.table("Indicator", ">Value", "Where it comes from")
	p.row("**Bills to pay**", "**"+money(state.AmountDue(card))+"**",
		"every closed statement that still carries a remainder")
	p.row("Statement being written", money(open.Remaining()),
		"charged since the card last closed, closing "+utils.PrettyDate(open.Closes))
	p.row("Next due date", utils.PrettyDate(open.Due),
		"day "+strconv.FormatInt(ledger.CardDueDay(card), 10)+" of the month")
	p.row("Outstanding in total", money(owed), "every purchase − every payment")
	p.row("Limit", money(card.Limit), "what the card was declared with")
	p.row("Available", money(card.Limit-owed), "limit − outstanding")
	p.blank()
	p.rule()
}

// billsHistory writes section 2: the statements themselves, newest first.
func billsHistory(p *page, state ledger.State, card ledger.Account) {
	bills := state.Bills(card)
	p.heading(2, "2. Every statement")
	if len(bills) == 0 {
		p.line("This card has no statement yet. Record a purchase on it with the " +
			"`AddTransaction` task and its cycle appears here by itself.")
		p.blank()
		p.rule()
		return
	}
	p.table("Cycle", "Closes", "Due", ">Charges", ">Total", ">Paid", ">Remaining",
		">Purchases", "Status")
	for index := len(bills) - 1; index >= 0; index-- {
		bill := bills[index]
		p.row(utils.PrettyMonth(bill.Cycle), utils.PrettyDate(bill.Closes),
			utils.PrettyDate(bill.Due), signed(bill.Charges()), money(bill.Total()),
			money(bill.Paid()), "**"+money(bill.Remaining())+"**",
			strconv.Itoa(len(bill.Purchases)), billStatus(bill, state.Today))
	}
	p.blank()
	p.line("A purchase counts on the statement of the cycle it is dated in — after the " +
		"previous close, up to and including this one. Money arriving on the card pays the " +
		"oldest statement still asking for something, and what is left of it runs on to the " +
		"next one.")
	if bills[0].Carried != 0 {
		p.blank()
		p.line("The **" + utils.PrettyMonth(bills[0].Cycle) + "** statement also carries " +
			money(-bills[0].Carried) + " — what the card was declared as already owing, " +
			"which is why its total is more than the purchases on it.")
	}
	p.blank()
	p.rule()
}

// billsOpen writes section 3: the purchases the statement still being written
// has collected so far.
func billsOpen(p *page, state ledger.State, open ledger.Bill) {
	p.heading(2, "3. The statement being written now")
	p.line("Cycle **" + utils.PrettyMonth(open.Cycle) + "** — charges dated after " +
		utils.PrettyDate(open.After) + ", closing " + utils.PrettyDate(open.Closes) +
		", due " + utils.PrettyDate(open.Due) + ".")
	p.blank()
	if len(open.Purchases) == 0 {
		p.line("Nothing has been charged to this card since it last closed.")
		p.blank()
		p.rule()
		return
	}
	billMovements(p, open.Purchases)
	p.line("**Charged so far:** " + money(open.Total()) + " · **already paid:** " +
		money(open.Paid()) + " · **it will ask for:** " + money(open.Remaining()))
	p.blank()
	p.rule()
}

// billsPaid writes section 4: what has already been paid on this card, and
// which statement each payment answered.
func billsPaid(p *page, state ledger.State, card ledger.Account) {
	bills := state.Bills(card)
	p.heading(2, "4. What has been paid")
	rows := 0
	p.table("Date", ">Id", "Statement it paid", "Description", ">Applied to it",
		">The movement", "Settles")
	for index := len(bills) - 1; index >= 0; index-- {
		bill := bills[index]
		for _, paid := range bill.Settlements {
			rows++
			p.row(utils.PrettyDate(paid.Transaction.Date), idText(paid.Transaction.Id),
				utils.PrettyMonth(bill.Cycle), dash(paid.Transaction.Description),
				money(paid.Applied), signed(paid.Transaction.Amount),
				settlement(paid.Transaction))
		}
	}
	p.blank()
	if rows == 0 {
		p.line("No payment has been recorded on this card yet. Pay a bill with the " +
			"`PayCreditCardBill` task — it writes both legs at once, the money leaving " +
			"the account and the same amount arriving on the card.")
		return
	}
	p.line("**Payments recorded:** " + strconv.Itoa(rows) + " · **total paid:** " +
		money(state.AccountFlow(card.Name, 0).In))
	p.blank()
	p.line("One payment can appear on two statements: it fills the oldest one still asking " +
		"for something, and what is left of it runs on to the next. Paying a bill is never " +
		"an expense either — the purchases were already counted on the day they happened, " +
		"which is why both legs share a transfer category.")
}

// billMovements writes a list of card movements as the statement shows them.
func billMovements(p *page, movements []ledger.Transaction) {
	p.table("Date", ">Id", "Category", "Description", ">Amount", "Settles")
	for _, transaction := range movements {
		p.row(utils.PrettyDate(transaction.Date), idText(transaction.Id),
			transaction.Category, dash(transaction.Description),
			signed(transaction.Amount), settlement(transaction))
	}
	p.blank()
}

// billStatus renders where a statement stands, with the overdue one marked so
// a table of them can be read in one pass.
func billStatus(bill ledger.Bill, today int64) string {
	if bill.Status(today) == ledger.BillOverdue {
		return "**" + ledger.BillOverdue + "**"
	}
	return bill.Status(today)
}

// pendingPage writes Pending.md: everything the vault knows is still waiting
// to be paid — the closed statements asking for money, the statements still
// collecting, the movements recorded ahead of their settlement, and the
// commitments the rest of the open month declares.
func pendingPage(state ledger.State) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Pending payment")
	p.line("> **Updated:** " + utils.PrettyDate(state.Today) +
		" · everything recorded that has not moved yet")
	p.blank()
	p.line(navigation())
	p.blank()
	p.rule()

	pendingSummary(p, state)
	pendingBills(p, state)
	pendingOpenBills(p, state)
	pendingMovements(p, state)
	pendingCommitments(p, state)
	return p.render("Pending.md")
}

// pendingSummary writes section 1: the figures the rest of the page breaks
// down.
func pendingSummary(p *page, state ledger.State) {
	p.heading(2, "1. What is still to pay")
	p.table("Line", ">Value", "What it is")
	p.row("**Card bills to pay**", "**"+money(state.TotalAmountDue())+"**",
		"closed statements still carrying a remainder")
	p.row("— of which overdue", money(state.TotalOverdue()), "past their due date")
	p.row("Statements still open", money(state.TotalOpenBills()),
		"charged since the cards last closed, not billed yet")
	p.row("Expenses not settled", signed(state.PendingAccountExpenses()),
		"movements on your accounts with a payment date still ahead")
	p.row("Card charges not settled", signed(state.PendingCardCharges()),
		"installments already recorded on the cards, part of a statement to come")
	p.row("Income not settled", signed(state.PendingAccountIncome()),
		"money you are still waiting to receive")
	p.row("**Leaving your accounts**", "**"+money(state.DueFromAccounts())+"**",
		"the bills to pay + the expenses not settled yet")
	p.blank()
	p.line("Card charges are left out of the last line on purpose: what a card has charged " +
		"leaves your accounts as part of a statement, so counting both would count it twice.")
	p.blank()
	p.rule()
}

// pendingBills writes section 2: the bills you actually have to pay, soonest
// due first.
func pendingBills(p *page, state ledger.State) {
	unpaid := state.UnpaidBills()
	p.heading(2, "2. Card bills to pay")
	if len(unpaid) == 0 {
		p.line("No closed statement is waiting for money. Every bill the cards have " +
			"issued is paid.")
		p.blank()
		p.rule()
		return
	}
	p.table("Due", "Card", "Cycle", "Closed", ">Total", ">Paid", ">Remaining", "Status")
	for _, bill := range unpaid {
		card, _ := state.Account(bill.Card)
		if card.Name == "" {
			card.Name = bill.Card
		}
		p.row(utils.PrettyDate(bill.Due), "["+bill.Card+"]("+billsPath(card)+")",
			utils.PrettyMonth(bill.Cycle), utils.PrettyDate(bill.Closes),
			money(bill.Total()), money(bill.Paid()), "**"+money(bill.Remaining())+"**",
			billStatus(bill, state.Today))
	}
	p.blank()
	p.line("Pay one with the `PayCreditCardBill` task. Given no `amount` it pays exactly " +
		"what the card's closed statements still ask for.")
	p.blank()
	p.rule()
}

// pendingOpenBills writes section 3: the statements still collecting charges,
// which will ask for money when their cycle closes.
func pendingOpenBills(p *page, state ledger.State) {
	p.heading(2, "3. Statements still open")
	if len(state.Cards()) == 0 {
		p.line("No credit card yet. Add one with the `AddCreditCard` task.")
		p.blank()
		p.rule()
		return
	}
	p.table("Card", "Cycle", "Closes", "Due", ">Charged so far", ">Purchases", "Statements")
	for _, card := range state.Cards() {
		open := state.OpenBill(card)
		p.row("["+card.Name+"]("+billsPath(card)+")", utils.PrettyMonth(open.Cycle),
			utils.PrettyDate(open.Closes), utils.PrettyDate(open.Due),
			money(open.Remaining()), strconv.Itoa(len(open.Purchases)),
			"[every statement]("+billsPath(card)+")")
	}
	p.blank()
	p.line("A statement asks for nothing until its cycle closes. What is charged to it now " +
		"is what it will ask for then.")
	p.blank()
	p.rule()
}

// pendingMovements writes section 4: every movement already recorded whose
// payment date is still ahead, soonest first.
func pendingMovements(p *page, state ledger.State) {
	unsettled := state.Unsettled()
	p.heading(2, "4. Movements not settled yet")
	if len(unsettled) == 0 {
		p.line("Every movement recorded has already settled — nothing is dated ahead.")
		p.blank()
		p.rule()
		return
	}
	p.table("Settles", ">Id", "Account", "Category", "Description", ">Amount", "Dated")
	for _, transaction := range unsettled {
		p.row(utils.PrettyDate(transaction.PaymentDate), idText(transaction.Id),
			"["+transaction.Account+"]("+accountPathOf(state, transaction.Account)+")",
			transaction.Category, dash(transaction.Description),
			signed(transaction.Amount), utils.PrettyDate(transaction.Date))
	}
	p.blank()
	p.line("**Movements:** " + strconv.Itoa(len(unsettled)) + " · **net when they all " +
		"settle:** " + signed(state.Pending()))
	p.blank()
	p.line("These are recorded facts, not projections: they already count in their month's " +
		"result, and they reach a balance on the day they settle.")
	p.blank()
	p.rule()
}

// pendingCommitments writes section 5: what the rest of the open month
// declares, which is the one part of this page that is not a recorded
// movement.
func pendingCommitments(p *page, state ledger.State) {
	open := state.OpenMonth()
	p.heading(2, "5. Commitments left this month")
	ahead := []ledger.Due{}
	for _, entry := range state.DueIn(open) {
		if entry.Date < state.Today {
			continue
		}
		ahead = append(ahead, entry)
	}
	if len(ahead) == 0 {
		p.line("Nothing else is declared for " + utils.PrettyMonth(open) + ". " +
			"The whole projection is on [Forecast.md](Forecast.md).")
		return
	}
	p.table("Date", "Commitment", "Account", ">Amount")
	for _, entry := range ahead {
		p.row(utils.PrettyDate(entry.Date), entry.Description,
			"["+entry.Account+"]("+accountPathOf(state, entry.Account)+")",
			signed(entry.Amount))
	}
	p.blank()
	p.line("A recurrence moves no money on its own. When the day arrives you record what " +
		"actually happened with `AddTransaction` — the whole projection is on " +
		"[Forecast.md](Forecast.md).")
}
