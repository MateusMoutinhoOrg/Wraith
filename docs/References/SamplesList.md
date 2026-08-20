# CLI Samples List

## Description
Every shell example shipped in [`examples/cliExamples/`](/examples/cliExamples/). Each one builds the binary and drives it in a temporary vault of its own, so it never touches a brain of yours. Running one is [RunCliSample.md](/docs/Tutorials/RunCliSample.md); adding one is [HandleCliExamples.md](/docs/Tutorials/HandleCliExamples.md).

---

## Examples

| Sample | Description |
| --- | --- |
| [BasicVault.sh](/examples/cliExamples/BasicVault.sh) | Goes from an empty folder to a rendered vault: the registries, an income, a transfer and a card purchase. |
| [DriveItWithTaskFile.sh](/examples/cliExamples/DriveItWithTaskFile.sh) | Drives the state machine through `Task.yaml`, including what an armed action that fails looks like. |
| [CreditCardBill.sh](/examples/cliExamples/CreditCardBill.sh) | Puts a card in the vault, spends on it, splits a purchase into installments, and reads the bill. |
| [MonthlyBudget.sh](/examples/cliExamples/MonthlyBudget.sh) | Declares the commitments that repeat every month and reads the forecast they produce. |
| [FreelanceIncome.sh](/examples/cliExamples/FreelanceIncome.sh) | An irregular income across several accounts: a category tree, an invoice settling later, money set aside. |
| [CorrectTheLedger.sh](/examples/cliExamples/CorrectTheLedger.sh) | Fixes what was recorded wrong: correcting by id, removing a movement, clearing a registry entry. |
| [ChooseWhatIsRendered.sh](/examples/cliExamples/ChooseWhatIsRendered.sh) | The rendering half of the interface: listings, `render`, arg and `dest` overrides, disabling an entry. |
