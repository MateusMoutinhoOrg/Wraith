# Fork This Repository as a Template

## Description
Covers using this repository as a GitHub template to build a **brain of your own**: the same state machine, driven by the same two files, with your actions and your pages instead of the financial ones. Converting a library that already exists to this structure is [AdaptExistingLib.md](/docs/Tutorials/AdaptExistingLib.md). Every phase takes each file's action — **Copy**, **Create**, **Rewrite**, or **Delete** — from [TemplateFileActions.md](/docs/References/TemplateFileActions.md).

### Rules
- Read [Structure.md](/docs/References/Structure.md) before starting.
- The financial brain in this repository is a **worked example of the shape**, not the point of it. Tasks, visualizations and the registries are yours to replace; the state machine around them is not.
- Keep the separation defined in [Structure.md](/docs/References/Structure.md): contract structs in `sandbox/contracts/`, actions in `sandbox/Tasks/`, renderers in `sandbox/Visualization/`, concrete dependencies in `adapters/`, the installed binary in `cmd/main/`. The command line belongs to the library, as the `Sandboxmain` field of `api.Lib`, never to the binary. Contracts are structs of function fields, never interfaces — see [StructContracts.md](/docs/References/StructContracts.md).
- Every file created or rewritten — code and `.md` alike — must follow its specification, located through [Specs.md](/docs/References/Specs.md).
- The fork is not complete until the final checklist in the last workflow step passes.

---

## Workflow

### Phase 1 — Create the repository
1. On the GitHub repository page, click **"Use this template"** and create the new repository.
2. Rename the module to the new GitHub path, following [RenameModule.md](/docs/Tutorials/RenameModule.md).
3. Rename the binary: the constants in [sandbox/config/cli.go](/sandbox/config/cli.go) hold every word the interface says, including its own name, and [scripts/](/scripts/) names the built artifacts.
4. Leave every **[Copy](/docs/References/TemplateFileActions.md#copy)** file untouched — they describe the structure, not the brain.

### Phase 2 — Decide what your brain holds
5. Rewrite the registries in [sandbox/config/database.go](/sandbox/config/database.go): one schema per kind of record your brain keeps, replacing accounts, categories, transactions and recurrences. Remember what the injected database offers — unique string keys and whole numbers — so amounts go in scaled to integers, dates go in as `20260818`, and free text travels packed into one key with `utils.Pack`.
6. Rewrite the derived figures in [sandbox/lib/ledger/](/sandbox/lib/ledger/), or delete the package if your brain has no arithmetic. Everything a page shows is computed here, so a visualization stays about layout.

### Phase 3 — Write the actions and the pages
7. Replace the tasks in [sandbox/Tasks/Tasks/](/sandbox/Tasks/Tasks/) with yours, one file each, and register them in [sandbox/Tasks/run.go](/sandbox/Tasks/run.go) — following [HandleTasks.md](/docs/Tutorials/HandleTasks.md).
8. Replace the visualizations in [sandbox/Visualization/Visualization/](/sandbox/Visualization/Visualization/) with yours, one file each, and register them in [catalog.go](/sandbox/Visualization/Visualization/catalog.go) — following [HandleVisualizations.md](/docs/Tutorials/HandleVisualizations.md). Keep `Help` and `Task-List`: both are generated from your own registries, so they document your brain the moment it builds.
9. Rewrite the defaults a new vault starts life with, in [assets/start/](/assets/start/) — the `Task.yaml` and `Visualization.yaml` that `wraith start` writes.

### Phase 4 — Adjust the frame, only where it moved
10. Add a dependency only when a task or a visualization needs an effect the contract does not carry yet: declare it in [sandbox/contracts/deps/deps.go](/sandbox/contracts/deps/deps.go) and fill it in every adapter, following [HandleDependencies.md](/docs/Tutorials/HandleDependencies.md).
11. Add a command only when your brain needs a verb the state machine does not have. The dispatch lives in `sandbox/cli/`, following [HandleCliCommands.md](/docs/Tutorials/HandleCliCommands.md).
12. Replace the samples: the Go programs in [examples/libraryExamples/](/examples/libraryExamples/), following [HandleLibrarySamples.md](/docs/Tutorials/HandleLibrarySamples.md), and the shell scripts in [examples/cliExamples/](/examples/cliExamples/), following [HandleCliExamples.md](/docs/Tutorials/HandleCliExamples.md).

### Phase 5 — Rewrite the documentation
13. Create the new API detail pages (`docs/References/PublicApi/<pkg>.<Symbol>.md`) and rewrite [PublicApi.md](/docs/References/PublicApi.md), following [HandleLibElements.md](/docs/Tutorials/HandleLibElements.md).
14. Rewrite the remaining **[Rewrite](/docs/References/TemplateFileActions.md#rewrite)** docs with your brain's content: [Structure.md](/docs/References/Structure.md), [Commands.md](/docs/References/Commands.md), [Adapters.md](/docs/References/Adapters.md), and the usage guides ([InstallCli.md](/docs/Tutorials/InstallCli.md), [StartABrain.md](/docs/Tutorials/StartABrain.md), [RunTasks.md](/docs/Tutorials/RunTasks.md), [ChooseVisualizations.md](/docs/Tutorials/ChooseVisualizations.md), [LibInitialization.md](/docs/Tutorials/LibInitialization.md), [RunCliSample.md](/docs/Tutorials/RunCliSample.md), [RunApiSample.md](/docs/Tutorials/RunApiSample.md), [SamplesList.md](/docs/References/SamplesList.md), [ApiSamplesList.md](/docs/References/ApiSamplesList.md)).
15. Replace the worked example — [TrackTransactions.md](/docs/Tutorials/TrackTransactions.md) is the financial brain's; yours needs its own one-page walk through a real use, following [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md) and the [TutorialDocs specification](/docs/References/Specs/TutorialDocs/Specs.md).
16. Delete every remaining **[Delete](/docs/References/TemplateFileActions.md#delete)** file — the financial tasks, visualizations, samples and docs your brain replaced. For `.md` files, follow [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md).
17. Rewrite the [README.md](/README.md) — overview, badges, and the Doc Index pointing at each theme index — then rewrite the four theme indexes (`docs/Index/<Theme>.md`) so each lists your brain's pages.

### Phase 6 — Verify
18. Build, then start a brain from nothing:

```bash
go build ./...
mkdir /tmp/verify && cd /tmp/verify
go run <your-module>/cmd/main start
go run <your-module>/cmd/main tick
```

Then confirm every item below — the fork is only done when all pass:
- Every task lives in one file under `sandbox/Tasks/Tasks/` and appears in `TaskArray`; `wraith tasks` lists exactly the ones you meant.
- Every visualization lives in one file under `sandbox/Visualization/Visualization/` and appears in `Catalog`; `wraith visualizations` lists exactly the ones you meant.
- A task writes through `args.DataBase` and nothing else; a visualization returns renders and writes no file.
- No file under `sandbox/` imports `os`, `net`, or a third-party implementation directly — every such call goes through `l.Deps`.
- `sandbox/contracts/deps/deps.go` declares one function field per injected call, and **every** adapter in `adapters/` fills all of them — the compiler does not check this.
- `sandbox/new.go` is the only wiring point, and it imports no adapter.
- `wraith tick` in an empty folder renders a vault with no error, and a deliberately broken `Task.yaml` writes `Error.md` and changes nothing.
- Every created or rewritten file matches its specification from [Specs.md](/docs/References/Specs.md).
- Every theme index in `docs/Index/` lists every page of its theme, the `README.md` Doc Index lists every theme, and the samples lists cover every sample.
- `cmd/main/main.go` wires, calls `Sandboxmain`, and exits — it branches on no command, parses no flag, and prints nothing.
