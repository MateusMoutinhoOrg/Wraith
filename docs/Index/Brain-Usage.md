# Brain Usage

## Description
Index of the documentation for people **driving** a brain: installing the binary, creating a vault, running actions, and choosing what gets rendered. Building a brain of your own is indexed by [Brain-Config.md](/docs/Index/Brain-Config.md); calling the same behavior from Go code is indexed by [LibUsage.md](/docs/Index/LibUsage.md).

A brain is a folder holding two files. `Task.yaml` decides what changes; `Visualization.yaml` decides what you get to see. Everything below is one of those two, or the commands that read them.

---

## Tutorials

- [InstallCli.md](/docs/Tutorials/InstallCli.md)
  - **description:** Install the binary globally, or build and run it from a checkout
  - [macOS](/docs/Tutorials/InstallCli.md#macos)
  - [Linux](/docs/Tutorials/InstallCli.md#linux)
  - [Windows (PowerShell)](/docs/Tutorials/InstallCli.md#windows-powershell)
  - [Verify after reboot](/docs/Tutorials/InstallCli.md#verify-after-reboot)
  - [Troubleshooting](/docs/Tutorials/InstallCli.md#troubleshooting)
  - [Install from a Clone](/docs/Tutorials/InstallCli.md#install-from-a-clone)
- [StartABrain.md](/docs/Tutorials/StartABrain.md)
  - **description:** Turn an empty folder into a working vault, and see what it renders
- [RunTasks.md](/docs/Tutorials/RunTasks.md)
  - **description:** Run one action, from `Task.yaml` or straight from the command line
- [ChooseVisualizations.md](/docs/Tutorials/ChooseVisualizations.md)
  - **description:** Decide which pages your brain writes, and where each one goes
- [TrackTransactions.md](/docs/Tutorials/TrackTransactions.md)
  - **description:** A worked month in the financial brain: accounts, categories, months, forecast
- [RunCliSample.md](/docs/Tutorials/RunCliSample.md)
  - **description:** Run one of the shipped shell examples from the source tree
  - [Run CLI Examples](/docs/Tutorials/RunCliSample.md#run-cli-examples)

---

## References

- [Commands.md](/docs/References/Commands.md)
  - **description:** Every command, flag, value format, and exit code of the interface
  - [Commands](/docs/References/Commands.md#commands)
  - [Flags](/docs/References/Commands.md#flags)
  - [Values](/docs/References/Commands.md#values)
  - [Exit Codes](/docs/References/Commands.md#exit-codes)
  - [The Tick Workflow](/docs/References/Commands.md#the-tick-workflow)
- [SamplesList.md](/docs/References/SamplesList.md)
  - **description:** Every shell example shipped in `examples/cliExamples/`
  - [Examples](/docs/References/SamplesList.md#examples)
