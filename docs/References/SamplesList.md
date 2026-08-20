# CLI Samples List

## Description
Every shell example shipped in [`examples/cliExamples/`](/examples/cliExamples/). Each one builds the binary and drives it in a temporary vault of its own, so it never touches a brain of yours. Running one is [RunCliSample.md](/docs/Tutorials/RunCliSample.md); adding one is [HandleCliExamples.md](/docs/Tutorials/HandleCliExamples.md).

---

## Examples

| Sample | Description |
| --- | --- |
| [BasicVault.sh](/examples/cliExamples/BasicVault.sh) | Goes from an empty folder to a rendered vault: the registries, an income, a transfer and two expenses. |
| [DriveItWithTaskFile.sh](/examples/cliExamples/DriveItWithTaskFile.sh) | Drives the state machine through `Task.yaml`, including what an armed action that fails looks like. |
| [MonthlyBudget.sh](/examples/cliExamples/MonthlyBudget.sh) | Dates a month's fixed lines forward and reads the forecast they produce on the month index. |
| [FreelanceIncome.sh](/examples/cliExamples/FreelanceIncome.sh) | An irregular income across several accounts: a category tree, invoices landing as they land, money set aside. |
| [CorrectTheLedger.sh](/examples/cliExamples/CorrectTheLedger.sh) | Fixes what was recorded wrong: correcting by id, removing a movement, clearing a registry entry. |
| [ChooseWhatIsRendered.sh](/examples/cliExamples/ChooseWhatIsRendered.sh) | The rendering half of the interface: listings, `render`, arg and `dest` overrides, disabling an entry. |
