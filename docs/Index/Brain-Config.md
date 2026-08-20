# Brain Config

## Description
Index of the documentation for people **building** a brain: forking this repository into one of their own, renaming it, and replacing its actions and pages with theirs. Driving a brain that already exists is indexed by [Brain-Usage.md](/docs/Index/Brain-Usage.md); contributing to this repository itself is indexed by [Development.md](/docs/Index/Development.md).

This repository is a template meant to be forked. The financial brain inside it — accounts, categories, transactions, recurrences — is a worked example of the shape, not the point of it: swap the tasks and the visualizations and you have a brain for something else entirely.

---

## Tutorials

- [ForkTemplate.md](/docs/Tutorials/ForkTemplate.md)
  - **description:** **Start here**: fork this repository and turn it into a brain of your own
  - [Phase 1 — Create the repository](/docs/Tutorials/ForkTemplate.md#phase-1--create-the-repository)
  - [Phase 2 — Decide what your brain holds](/docs/Tutorials/ForkTemplate.md#phase-2--decide-what-your-brain-holds)
  - [Phase 3 — Write the actions and the pages](/docs/Tutorials/ForkTemplate.md#phase-3--write-the-actions-and-the-pages)
  - [Phase 4 — Adjust the frame, only where it moved](/docs/Tutorials/ForkTemplate.md#phase-4--adjust-the-frame-only-where-it-moved)
  - [Phase 5 — Rewrite the documentation](/docs/Tutorials/ForkTemplate.md#phase-5--rewrite-the-documentation)
  - [Phase 6 — Verify](/docs/Tutorials/ForkTemplate.md#phase-6--verify)
- [RenameModule.md](/docs/Tutorials/RenameModule.md)
  - **description:** Rename the Go module path to your GitHub name, and fix every import
- [HandleTasks.md](/docs/Tutorials/HandleTasks.md)
  - **description:** Add an action: one file, one registry line, and the guides write themselves
- [HandleVisualizations.md](/docs/Tutorials/HandleVisualizations.md)
  - **description:** Add a renderer: one file, one catalog line, one entry in your config
- [HandleDependencies.md](/docs/Tutorials/HandleDependencies.md)
  - **description:** Give your tasks an effect the contract does not carry yet
  - [Find Dependencies Functions you can use](/docs/Tutorials/HandleDependencies.md#find-dependencies-functions-you-can-use)
  - [Add New Dependencie](/docs/Tutorials/HandleDependencies.md#add-new-dependencie)
  - [Overwrinting a adapter function](/docs/Tutorials/HandleDependencies.md#overwrinting-a-adapter-function)
  - [Creating a adapter in repo](/docs/Tutorials/HandleDependencies.md#creating-a-adapter-in-repo)
  - [Creating a adapter in your project](/docs/Tutorials/HandleDependencies.md#creating-a-adapter-in-your-project)
- [HandleCliCommands.md](/docs/Tutorials/HandleCliCommands.md)
  - **description:** Add a command to the interface, when your brain needs a new verb
  - [Add CLI Command](/docs/Tutorials/HandleCliCommands.md#add-cli-command)
  - [Remove CLI Command](/docs/Tutorials/HandleCliCommands.md#remove-cli-command)
- [HandleAssets.md](/docs/Tutorials/HandleAssets.md)
  - **description:** Change the defaults a new vault starts with, or ship any file in the binary
  - [Add Asset](/docs/Tutorials/HandleAssets.md#add-asset)
  - [ListAssets in Runtime](/docs/Tutorials/HandleAssets.md#listassets-in-runtime)
  - [Retrieve Asset in Runtime](/docs/Tutorials/HandleAssets.md#retrieve-asset-in-runtime)
- [AdaptExistingLib.md](/docs/Tutorials/AdaptExistingLib.md)
  - **description:** Convert a library you already have to this dependency-injection structure
  - [Phase 1 — Bring the structure in](/docs/Tutorials/AdaptExistingLib.md#phase-1--bring-the-structure-in)
  - [Phase 2 — Rewrite the contracts](/docs/Tutorials/AdaptExistingLib.md#phase-2--rewrite-the-contracts)
  - [Phase 3 — Move the code into the sandbox](/docs/Tutorials/AdaptExistingLib.md#phase-3--move-the-code-into-the-sandbox)
  - [Phase 4 — Rewrite the documentation](/docs/Tutorials/AdaptExistingLib.md#phase-4--rewrite-the-documentation)
  - [Phase 5 — Verify](/docs/Tutorials/AdaptExistingLib.md#phase-5--verify)
- [Build.md](/docs/Tutorials/Build.md)
  - **description:** Cross-compile your brain into a binary for each supported OS and architecture
  - [Build a single target](/docs/Tutorials/Build.md#build-a-single-target)
  - [Build every target at once](/docs/Tutorials/Build.md#build-every-target-at-once)
  - [Add a new target](/docs/Tutorials/Build.md#add-a-new-target)

---

## References

- [TemplateFileActions.md](/docs/References/TemplateFileActions.md)
  - **description:** The per-file action a fork follows: copy, create, rewrite, or delete
  - [Copy](/docs/References/TemplateFileActions.md#copy)
  - [Rewrite](/docs/References/TemplateFileActions.md#rewrite)
  - [Create](/docs/References/TemplateFileActions.md#create)
  - [Delete](/docs/References/TemplateFileActions.md#delete)
- [Structure.md](/docs/References/Structure.md)
  - **description:** The project's schema: which kind of file lives where, and its spec
  - [Root](/docs/References/Structure.md#root)
  - [`/scripts/`](/docs/References/Structure.md#scripts)
  - [`/sandbox/`](/docs/References/Structure.md#sandbox)
  - [`/adapters/`](/docs/References/Structure.md#adapters)
  - [`/assets/`](/docs/References/Structure.md#assets)
  - [`/cmd/`](/docs/References/Structure.md#cmd)
  - [`/examples/cliExamples/`](/docs/References/Structure.md#examplescliexamples)
  - [`/examples/libraryExamples/`](/docs/References/Structure.md#exampleslibraryexamples)
  - [`/docs/`](/docs/References/Structure.md#docs)
- [StructContracts.md](/docs/References/StructContracts.md)
  - **description:** Why every contract is a struct of function fields, and how factories fill them
  - [The Shape](/docs/References/StructContracts.md#the-shape)
  - [Factories Fill the Fields](/docs/References/StructContracts.md#factories-fill-the-fields)
  - [Adapters Fill Their Contract the Same Way](/docs/References/StructContracts.md#adapters-fill-their-contract-the-same-way)
  - [Replacing One Behavior](/docs/References/StructContracts.md#replacing-one-behavior)
  - [Consuming a Library That Uses This Pattern](/docs/References/StructContracts.md#consuming-a-library-that-uses-this-pattern)
  - [What It Costs](/docs/References/StructContracts.md#what-it-costs)
