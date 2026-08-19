# Project Structure

This document maps the project **schema** — the kinds of files the project is built from — not every concrete file. A slot with a **Spec** name is governed by a specification; resolve the name through [Specs.md](/docs/References/Specs.md) to get its description and sample.

The project is a **brain**: a state machine over a folder, driven by two files, whose command-line interface lives inside the library. It is split into three top-level trees, and the dependency flow between them is one-way:

```
adapters/  ──▶  sandbox/  ◀──  cmd/, examples/libraryExamples/
(reaches the OS)  (closed)     (wire the two together)
```

- **`/sandbox/`** is a **closed sandbox**: the pure library, and the command-line interface with it. Nothing inside it may import `adapters/`, `cmd/`, `examples/libraryExamples/`, a third-party module, or any OS-bound standard-library package. Every effect it needs arrives through the injected `Deps`. See [SandboxIsolation.md](/docs/References/SandboxIsolation.md).
- **`/adapters/`** sits outside the sandbox and is the only place OS-bound and third-party code is allowed. Each adapter imports `sandbox/contracts/deps` and nothing else from the sandbox.
- **`/cmd/`** and **`/examples/libraryExamples/`** sit outside the sandbox too, and are the only places where an adapter and the sandbox meet — `cmd/` for the installable binary, `examples/libraryExamples/` for the runnable Go samples.
- **`/assets/`** sits outside the sandbox as well: files compiled into the binary and reached only through the injected `Deps.EmbedDeps` contract, so the sandbox holds no asset files of its own. It carries the defaults `wraith start` writes — see [`/assets/`](#assets).

Because the interface is `api.Lib.Sandboxmain` — one field of the library like any other — the binary in `cmd/main/` holds no command, no flag, and no output of its own: it wires, runs, and exits.

## Root

| File | Description | Spec |
|------|-------------|------|
| `README.md` | Project overview and the Doc Index pointing at each theme index under `docs/Index/` | Readme |
| `LICENSE` | License terms for the project | |
| `go.mod` | Go module definition and dependencies | |
| `.gitignore` | Intentionally untracked files to ignore | |

---

## `/scripts/`
The cross-platform build scripts. One shell script per OS/architecture target, each a thin wrapper over `go build` with the target's `GOOS`/`GOARCH` set — the Go runtime cross-compiles on its own, so nothing here needs a container runtime or a cross-compiler. Every script resolves the repository root from its own path and writes its artifact to `release/`, which is git-ignored. Building is [Build.md](/docs/Tutorials/Build.md).

| File | Description | Spec |
|------|-------------|------|
| `all.sh` | Runs every target script below, in order | |
| `linux86.sh` | Builds for Linux amd64, producing `release/linux86.out` | |
| `linuxarm64.sh` | Builds for Linux arm64, producing `release/linuxarm64.out` | |
| `linuxi32.sh` | Builds for Linux 386, producing `release/linuxi32.out` | |
| `windows86.sh` | Builds for Windows amd64, producing `release/windows86.exe` | |
| `windowsi32.sh` | Builds for Windows 386, producing `release/windowsi32.exe` | |
| `mac86.sh` | Builds for macOS Intel (amd64), producing `release/mac86.bin` | |
| `macarm64.sh` | Builds for macOS Apple Silicon (arm64), producing `release/macarm64.bin` | |

**Build one target:**
```sh
bash ./scripts/linux86.sh
```

**Build every target:**
```sh
bash ./scripts/all.sh
```

---

## `/sandbox/`
The closed sandbox — the pure library. It holds its own entry point, the contracts everything is wired through, the configuration constants, and the internal implementation. It reaches nothing outside itself: every OS-bound or third-party effect arrives through the injected `Deps`. Its package is named `lib`, so consumers import it as `lib "…/sandbox"` and call `lib.New`.

| File | Description | Spec |
|------|-------------|------|
| `new.go` | The `New(d deps.Deps, databasePath string) api.Lib` constructor, storing `Deps` and the database path on `api.Lib` and delegating to the internal lib constructor | |

### `/sandbox/contracts/`
The structs the rest of the project is wired through — the only part of the sandbox anything outside it may import. Contracts hold the project's **public types** and are structs of function fields, never interfaces; see [StructContracts.md](/docs/References/StructContracts.md). Contracts import nothing from `adapters/` or `sandbox/`.

#### `/sandbox/contracts/deps/`
The contract every adapter must fill.

| File | Description | Spec |
|------|-------------|------|
| `deps.go` | The `Deps` struct, one function field per injectable behavior, plus one plain field per embedded library | Deps |

##### `/sandbox/contracts/deps/verbdeps/`
The sandbox's copy of the embedded [Verb](https://github.com/MateusMoutinhoOrg/Verb) argv-parser library's public api. The sandbox may not import Verb — that would be a third-party import — so it restates the shape it needs, field for field; the adapter, outside the sandbox, is what fills it.

| File | Description | Spec |
|------|-------------|------|
| `verbdeps.go` | Copy of the embedded Verb library's `api.Lib` struct, injected whole as the `Deps.VerbLib` field | |

##### `/sandbox/contracts/deps/keepdeps/`
The sandbox's copy of the embedded [Keep](https://github.com/MateusMoutinhoOrg/Keep) schema-database library's public api, for the same reason `verbdeps/` exists. Keep's api is a tree of structs — `Lib` hands back a `KeepDatabase`, which hands back a `SchemaInstance`, which hands back `SchemaItem`s — so the copy restates the whole tree, and the adapter converts each returned struct as it crosses.

| File | Description | Spec |
|------|-------------|------|
| `keepdeps.go` | Copy of the embedded Keep library's api structs and constants, injected whole as the `Deps.KeepLib` field | |

##### `/sandbox/contracts/deps/embeddeps/`
The sandbox's copy of an embedded-asset library's public api, for the same reason the two above exist: reading a file is an OS-bound effect, and compiling one into a binary needs the `//go:embed` directive, so the sandbox may declare neither. It is how the library reaches the assets under [`/assets/`](#assets).

| File | Description | Spec |
|------|-------------|------|
| `embeddeps.go` | Copy of the asset-reading api — `ReadFile`, `ListFiles`, `ListFilesRecursively` — injected whole as the `Deps.EmbedDeps` field | |

##### `/sandbox/contracts/deps/iodeps/`
The sandbox's copy of a filesystem IO library's public api. The sandbox may not import `os` or any OS-bound package, so it declares the shape it needs — file read/write, directory creation, existence checks, and listing — and the adapter, outside the sandbox, fills it.

| File | Description | Spec |
|------|-------------|------|
| `iodeps.go` | The `Lib` struct declaring every filesystem operation the sandbox requires, injected whole as the `Deps.IoLib` field | |

##### `/sandbox/contracts/deps/requestdeps/`
The sandbox's copy of an HTTP client library's public api. The sandbox may not import `net/http`, so it declares the request/response shape it needs, and the adapter fills it.

| File | Description | Spec |
|------|-------------|------|
| `requestdeps.go` | The `Request` and `Response` structs for HTTP communication, used by the `Deps.NewRequest` function field | |

#### `/sandbox/contracts/api/`
The structs the library hands back to callers.

| File | Description | Spec |
|------|-------------|------|
| `api.go` | The `Lib` entry-point struct, the `Task` and `Visualizer` declarations, the `Field` they are described with, and the argument structs a task and a visualization are handed | Outputs |

### `/sandbox/config/`
Static configuration the sandbox reads at compile time: the text every command prints, the flag spellings the interface understands, the version string, and the shape of the database — the registries every task writes and every visualization reads, and the fields each one carries. Holding them as Go constants rather than as files keeps every reference under the compiler's eye — a renamed constant is a build failure rather than a blank line at runtime — and costs no read at all. Text long enough to be edited as a document, or shaped as a template, belongs in [`/assets/`](#assets) instead. Nothing outside the sandbox imports this package.

| File | Description | Spec |
|------|-------------|------|
| `cli.go` | The `Usages` help screen, every message constant the interface prints, the default `Task.yaml` / `Visualization.yaml` / `data` paths, and the flag spellings | |
| `database.go` | The registries and their fields as constants, and the `DatabaseProps` the injected Keep library opens them from | |
| `version.go` | The `Version` constant reported by `wraith version` and `--version` | |

### `/sandbox/lib/`
The entry-point implementation and every internal package the library is built from. Each package here holds the functions that take a pointer to an [`api`](#sandboxcontractsapi) struct and return closures reading that struct's `Deps`, which the package's constructor assigns into the matching function fields. Types never live here; they stay in `contracts/`.

| File | Description | Spec |
|------|-------------|------|
| `new.go` | The `New(d deps.Deps, databasePath string) api.Lib` constructor: it reads the task and visualization registries once, then assigns every factory's return value | LibFunctions |

#### `/sandbox/lib/publicfunctions/`
One file per public function field of `api.Lib`. Each file holds its `<Field>Factory(l *api.Lib)` that returns a closure. The `New` constructor in `sandbox/lib/new.go` calls every factory here to fill `api.Lib`.

| File | Description | Spec |
|------|-------------|------|
| `<Function>.go` | One file per lib function, holding its `<Field>Factory(l *api.Lib)` that returns a closure | LibFunctions |

#### `/sandbox/lib/utils/`
The small computations the rest of the sandbox is written with, none of which knows what a registry is: rendering an amount held in cents, doing arithmetic on a calendar, and packing several strings into one storage key. A task, a visualization and the ledger all reach for the same one, which is what makes a figure render the same way wherever it appears. It declares no types and no factories, so no specification governs it.

| File | Description | Spec |
|------|-------------|------|
| `utils.go` | `Pack`, `Unpack` and `Part` — the packed keys the injected database's unique-key rule forces | |
| `dates.go` | Dates and months as whole numbers, and the calendar arithmetic every month page and the forecast are built from | |
| `money.go` | Rendering an amount held in cents, and the bars and percentages the dashboards show a share with | |

#### `/sandbox/lib/ledger/`
The registries read back as typed values, and every figure the vault shows derived from them and nothing else. Keeping the arithmetic here is what lets a visualization be about layout. Records are read straight through `deps.Deps.KeepLib` — a schema is asked of the database, the schema lists its records, and a record hands back one field at a time. Declares no factories.

| File | Description | Spec |
|------|-------------|------|
| `records.go` | One typed view per registry, and the ordered listings that read them off the injected database | |
| `ledger.go` | The `State` read once per render, and the balances, month results and totals over it | |
| `forecast.go` | Today's position rolled forward through the declared commitments | |

#### `/sandbox/lib/vault/`
The outside of a tick: reading the task file, disarming it, reading and validating the visualization config, writing rendered bytes under a `dest`, and reporting a failure in `Error.md`. Every path is reached through `deps.Deps.IoLib`. Declares no factories.

| File | Description | Spec |
|------|-------------|------|
| `vault.go` | The task and config readers, the destination rules, and the writers a tick puts files on disk with | |

#### `/sandbox/lib/yaml/`
The subset of YAML the two driving files are written in — a flat mapping, and a list of mappings with one nested block. Pure computation over bytes, so it needs no dependency. Declares no factories.

| File | Description | Spec |
|------|-------------|------|
| `yaml.go` | The decoders, the scalar coercion, and the encoder a disarmed task file is rewritten with | |

#### `/sandbox/lib/entries/`
The fields a task or a visualization was called with, read back as typed values and checked against what it declared. One validator serves every task, which is why a task never parses its own input. Declares no factories.

| File | Description | Spec |
|------|-------------|------|
| `entries.go` | The typed readers, `Validate` against a `[]api.Field`, and the filling-in of declared defaults | |

### `/sandbox/Tasks/`
Every action the brain can perform. The switcher here is the seam between "what a brain can do" and everything that asks it to do something — a tick, the command line, a Go caller. Adding an action is [HandleTasks.md](/docs/Tutorials/HandleTasks.md).

| File | Description | Spec |
|------|-------------|------|
| `run.go` | `TaskArray` — the one place a task is registered — plus `Find`, `Names`, and the `Run` switcher that validates a task's fields and calls its `HandleAction` | |

#### `/sandbox/Tasks/Tasks/`
One task per file, each returning an `api.Task` carrying its name, its declared fields, and the closure that runs it. A task writes through the database it is handed and nothing else.

| File | Description | Spec |
|------|-------------|------|
| `<Task>.go` | One task, named after it: its declaration and its `HandleAction` closure | |
| `shared.go` | The checks every task repeats — required names, dates, days of the month, amounts, insert and remove — so all of them report a failure in the same words | |

### `/sandbox/Visualization/`
Every renderer the brain carries. The mirror of `/sandbox/Tasks/`: a tick reads names out of the config, the command line takes one as an argument, and both arrive here. Adding a renderer is [HandleVisualizations.md](/docs/Tutorials/HandleVisualizations.md).

| File | Description | Spec |
|------|-------------|------|
| `run.go` | `VisualizationArray` — the catalog, read from the package below — plus `Find`, `Names`, and the `Run` switcher that validates a visualization's args and calls its `HandleVisualizer` | |

#### `/sandbox/Visualization/Visualization/`
One visualization per file, each returning an `api.Visualizer`. A visualization **returns files** and never writes one; putting bytes on disk happens in `/sandbox/lib/vault/`.

| File | Description | Spec |
|------|-------------|------|
| `<Visualization>.go` | One renderer, named after it: its declaration and its `HandleVisualizer` closure | |
| `catalog.go` | `Catalog` — the one place a visualization is registered | |
| `page.go` | The markdown builder every page is written with, and the naming rules its links follow | |

### `/sandbox/cli/`
The command-line interface itself: the command dispatch `Sandboxmain` delegates to. It reads the command line through `deps.Deps.VerbLib`, takes the text it prints from `sandbox/config`, and writes every line through `deps.Deps.Printf`, so the whole interface stays inside the closed sandbox. Like `utils/`, it is neither an object nor the entry point, so no specification governs it, and it declares **no types and no factories**.

| File | Description | Spec |
|------|-------------|------|
| `run.go` | The `Run(l *api.Lib, args []string) int` dispatch: flag handling, command matching, and delegation to `commands/` | |

#### `/sandbox/cli/commands/`
One file per command or command group the interface supports. Each file holds the handler function the `Run` dispatch delegates to. Like `cli/` itself, this package declares **no types and no factories**.

| File | Description | Spec |
|------|-------------|------|
| `<command>.go` | One file per command or command group, holding the handler function that `Run` dispatches to | |

---

## `/adapters/`
Outside the sandbox. Opinionated implementations of the [`Deps`](#sandboxcontractsdeps) contract, each providing a distinct concrete behavior. This is where OS-bound and third-party code lives; an adapter imports `sandbox/contracts/deps` and nothing else from `sandbox/`. An adapter fills its contract with the same **factories** [`sandbox/`](#sandboxlib) uses — the carrier is the adapter struct, which declares the `Deps` field the factories' return values are assigned into.

### `/adapters/<name>/`

| File | Description | Spec |
|------|-------------|------|
| `<name>.go` | A struct carrying a `Deps` field, one `<Field>Factory` per `Deps` field returning a closure, plus the `New(...) deps.Deps` constructor that assigns every factory's return value and runs them all | Adapters |
| `<field>.go` | One factory, split out of `<name>.go` when it carries conversion helpers of its own — `embed.go` wraps the compiled-in [`/assets/`](#assets) into the `Deps.EmbedDeps` contract, `io.go` fills `Deps.IoLib` with `os`/`filepath` calls, `server.go` fills `Deps.NewRequest` with `net/http` calls. `New` still calls each one | Factories |

---

## `/assets/`
Outside the sandbox. The files the library serves through the injected `Deps.EmbedDeps` contract: compiled into the binary, so an installed `wraith` carries them with no files beside it, and reached only through the injected contract — never imported by the sandbox. An asset is a payload better kept as a file than as a Go constant: a template, a long-form document, an image. Adding one is [HandleAssets.md](/docs/Tutorials/HandleAssets.md).

The tree carries what a new vault starts life with: `start/Task.yaml` and `start/Visualization.yaml`, the two files [`wraith start`](/docs/References/Commands.md#start) copies into an empty folder. They are assets rather than Go constants precisely because they are meant to be edited — changing what a forked brain begins with is editing those files and nothing else. The words the interface itself says are the opposite case, and live in [`/sandbox/config/`](#sandboxconfig) as compile-time constants.

This directory is a Go package for one reason: a `//go:embed` directive can only reach files inside its own package directory, so the directive has to sit next to the assets. That single directive is `//go:embed all:*`, which takes **every** file in the tree, so a new asset needs no change to it — put the file here and it exists at runtime.

| File | Description | Spec |
|------|-------------|------|
| `asset.go` | Package `assets`: the `//go:embed all:*` directive and the `Files` embedded filesystem the standard adapter serves | |
| `start/Task.yaml` | The default task file `wraith start` writes | |
| `start/Visualization.yaml` | The default visualization config `wraith start` writes | |

---

## `/cmd/`
Outside the sandbox. The executables the project ships — for this project, the single binary a user installs.

### `/cmd/main/`

| File | Description | Spec |
|------|-------------|------|
| `main.go` | Self-contained `package main` that wires an adapter into the lib, calls `api.Lib.Sandboxmain(os.Args[1:])`, and exits with its return | CliMain |

**Run the CLI from source:**
```sh
go run ./cmd/main <command> [arguments]
```

**Install it globally:**
```sh
go install github.com/MateusMoutinhoOrg/Wraith/cmd/main@latest
```

---

## `/examples/cliExamples/`
Outside the sandbox. Shell scripts driving the built binary the way a user would from a terminal. Each one builds the binary into a scratch directory and runs it in a temporary vault of its own, so nothing a script does touches a brain of yours.

| File | Description | Spec |
|------|-------------|------|
| `<Name>.sh` | Self-contained shell script demonstrating one goal against the built CLI | CliExamples |

**Run a CLI example:**
```sh
bash ./examples/cliExamples/StartAVault.sh
```

---

## `/examples/libraryExamples/`
Outside the sandbox. Runnable Go examples demonstrating how to use the library from code, when the CLI is not what the caller wants.

### `/examples/libraryExamples/<example>/`

| File | Description | Spec |
|------|-------------|------|
| `<example>.go` | Self-contained `package main` wiring an adapter into the lib | LibraryExamples |

**Run an example:**
```sh
go run ./examples/libraryExamples/<example>/<example>.go
```

---


## `/docs/`
Documentation of the project, split by **kind of page**: `Index/` holds one entry point per theme, `Tutorials/` holds every workflow, `References/` holds every lookup and explanation. A **theme** — what the reader wants to accomplish — is not a directory: it is the index that lists a page. `Tutorials/` and `References/` are flat, so a page's file name must be unique inside the directory it lands in; when two themes need the same topic, the name carries the subject. The [README](/README.md) links to the four indexes and to nothing else inside `docs/`.

| Directory | Description |
|-----------|-------------|
| `Index/` | One entry point per theme, each listing the pages of its theme |
| `Tutorials/` | Every workflow page of the project, whatever theme it belongs to |
| `References/` | Every lookup and explanation page, whatever theme it belongs to |

### `/docs/Index/`
One page per theme. The four themes are `Brain-Usage` — installing the binary, driving a vault from a terminal, and the command surface; `Brain-Config` — forking this repository into a brain of your own, and adding tasks and visualizations to it; `LibUsage` — consuming the same behavior as a Go library, and the public API; and `Development` — contributing: the mechanics, the workflows, the specifications.

| File | Description | Spec |
|------|-------------|------|
| `<Theme>.md` | The theme's entry point: its Tutorials and its References, each entry listing that page's sections | Index |

### `/docs/Tutorials/`
One page per workflow, its title phrased as the action it performs. A page can belong to one or more themes, and the theme indexes in [`/docs/Index/`](#docsindex) are what say which.

| File | Description | Spec |
|------|-------------|------|
| `<Goal>.md` | One page per workflow, written as numbered steps a reader follows to the end | TutorialDocs |

### `/docs/References/`
One page per lookup table or explained mechanic, plus the two directories the project's biggest listings live in.

| File | Description | Spec |
|------|-------------|------|
| `<Name>.md` | One page per lookup table or explained mechanic | ReferenceDocs / ExplanationDocs |
| `Commands.md` | Every command, flag, value format and exit code of the interface, and the tick workflow | ReferenceDocs |
| `SamplesList.md` | Every example under `examples/cliExamples/` | ReferenceDocs |
| `PublicApi.md` | Index of all public-facing components, linking to their detail pages | ReferenceDocs |
| `Adapters.md` | Lists every shipped adapter and when to use each one | AdaptersDoc |
| `ApiSamplesList.md` | Every example under `examples/libraryExamples/` | ReferenceDocs |
| `Structure.md` | The project's schema and the purpose of each component | Structure |
| `Specs.md` | Index of every specification and the files each one governs | |
| `SandboxIsolation.md` | What the sandbox may not import, and why every effect is a dep | ExplanationDocs |
| `StructContracts.md` | Why contracts are structs of function fields, and how factories fill them | ExplanationDocs |
| `TemplateFileActions.md` | The action each template file takes when forking or adapting | ReferenceDocs |

#### `/docs/References/PublicApi/`
One detail page per public-facing component, indexed by `PublicApi.md`. Reach a page through that index rather than by browsing the directory.

| File | Description | Spec |
|------|-------------|------|
| `<pkg>.<Symbol>.md` | One detail page per public struct, function, or field, named after the package the symbol is declared in | ReferenceDocs |

#### `/docs/References/Specs/`
The specifications describing how each kind of file in the project must be shaped. Never browse this directory — locate a specification by reading `Specs.md`.

| File | Description | Spec |
|------|-------------|------|
| `<Spec>/Specs.md` | The required shape of the artifact the specification governs | |
| `<Spec>/sample.<ext>` | Concrete reference implementation of the specification | |
