# CLI Samples List

## Description
Every shell example shipped in [`examples/cliExamples/`](/examples/cliExamples/). Each one builds the binary and drives it in a temporary vault of its own, so it never touches a brain of yours. Running one is [RunCliSample.md](/docs/Tutorials/RunCliSample.md); adding one is [HandleCliExamples.md](/docs/Tutorials/HandleCliExamples.md).

---

## Examples

| Sample | Description |
| --- | --- |
| [BasicVault.sh](/examples/cliExamples/BasicVault.sh) | Goes from an empty folder to a rendered vault: the registries, half a year of movements, a transfer and a forecast. |
| [StarterSetup.sh](/examples/cliExamples/StarterSetup.sh) | The setup an ordinary household starts from: everyday categories, a bank account and a wallet, four months of living. |
| [DriveItWithTaskFile.sh](/examples/cliExamples/DriveItWithTaskFile.sh) | Drives the state machine through `Task.yaml`, one armed action per tick, including what a failing one looks like. |
| [MonthlyBudget.sh](/examples/cliExamples/MonthlyBudget.sh) | Dates a household's fixed lines forward and reads the forecast they produce, yearly bill included. |
| [FreelanceIncome.sh](/examples/cliExamples/FreelanceIncome.sh) | An irregular income across several accounts: a category tree, invoices landing as they land, a tax reserve that pays a quarterly bill. |
| [CorrectTheLedger.sh](/examples/cliExamples/CorrectTheLedger.sh) | Fixes what was recorded wrong: correcting by id, moving a movement, removing one, clearing a registry entry. |
| [ChooseWhatIsRendered.sh](/examples/cliExamples/ChooseWhatIsRendered.sh) | The rendering half of the interface: listings, `render`, arg and `dest` overrides, disabling an entry. |

---

## Determinism

Every script above pins its vault to August 2026 with the three dashboard flags of [`wraith start`](/docs/References/Commands.md#start) — `--prev-months`, `--future-months` and `--current-month` — and dates every movement it records. Nothing a script prints is read off the clock, so two runs on different days produce the same transcript, and a diff against a saved one is a real difference.
