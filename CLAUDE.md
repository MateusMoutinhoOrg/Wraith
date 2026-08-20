# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Before Editing Anything

This repository is spec-driven. Two documents are binding and must be consulted **before** creating or editing a file:

- [docs/References/Specs.md](/docs/References/Specs.md) — the index of every file specification. Find the file you are about to touch in an **Applies To** column and follow the linked `Specs.md` plus its `sample`. Never browse `docs/References/Specs/` directly; the index is the only entry point.
- [docs/References/Structure.md](/docs/References/Structure.md) — the project schema: which kind of file belongs where.

Per-goal workflows live in `docs/Tutorials/`, indexed by theme in `docs/Index/` (`Development.md` is the contributor index). Adding a task, a visualization, a CLI command, a dep, or a doc each has a tutorial that lists the companion files to update in the same change.

## Commands

```bash
go build ./...                              # compile everything
go vet ./...                                # static check (no CI or lint config in the repo)
go run ./cmd/main <command> [args] [flags]  # run the CLI from source
bash ./scripts/mac86.sh                     # build one target into release/ (see scripts/ for the rest)
bash ./scripts/all.sh                       # build every OS/arch target
bash ./localInstal.sh                       # build mac86, install as /usr/local/bin/wraith, tick WraithSample
bash ./examples/cliExamples/BasicVault.sh   # run a CLI example end-to-end
go run ./examples/libraryExamples/<name>/<name>.go
```

There are **no Go tests** in this repository. Behavior is exercised through the runnable examples in `examples/cliExamples/` (shell, against an installed `wraith`) and `examples/libraryExamples/` (Go, wiring an adapter into the lib). Changing behavior means updating or adding one of those, not a `_test.go`.

`release/` and `WraithSample/` are git-ignored scratch output.

## Architecture

Wraith is a state machine over a folder ("a vault"): a task is written into `Task.yaml`, `wraith tick` applies it to the Keep database under `data/` and re-renders every visualization declared in `Visualization.yaml`. The financial brain shipped here (accounts, categories, transactions) is a worked example of the shape, meant to be replaced.

### The sandbox wall

```
adapters/  ──▶  sandbox/  ◀──  cmd/, examples/libraryExamples/
(reaches the OS)  (closed)     (wire the two together)
```

`sandbox/` is closed. A file under it may **not** import `adapters/`, `cmd/`, `examples/`, any third-party module, any OS-bound stdlib (`os`, `net`, `syscall`, …), or `embed`/`//go:embed`. Every effect arrives as a function field on the injected `deps.Deps`: `Now`, `Printf`, `IoLib`, `VerbLib` (argv parsing), `KeepLib` (the database), `EmbedDeps` (the compiled-in `assets/`), `NewRequest`. Adding a new effect means adding a `Deps` field, not an import — see [HandleDependencies.md](/docs/Tutorials/HandleDependencies.md).

Because of this, the **entire CLI lives inside the sandbox** (`sandbox/cli/`, reached through `api.Lib.Sandboxmain`). `cmd/main/main.go` holds no command, flag, or output of its own: it builds the standard adapter, calls `lib.New(deps, "data")`, and exits with `l.Sandboxmain(os.Args[1:])`. Full rationale in [SandboxIsolation.md](/docs/References/SandboxIsolation.md).

### Contracts are structs of function fields

No interfaces. `sandbox/contracts/deps` (what an adapter fills) and `sandbox/contracts/api` (what the lib hands back) are structs whose fields are functions. Each field is filled by a `<Field>Factory(l *api.Lib)` returning a closure that reads the struct's deps; the package constructor assigns every factory's return value. `sandbox/lib/new.go` is that assignment list for `api.Lib` — adding a public function means adding a file in `sandbox/lib/publicfunctions/` and a line there. Types never live in `sandbox/lib/`; they stay in `contracts/`. See [StructContracts.md](/docs/References/StructContracts.md).

Because the sandbox cannot import Keep or Verb, `sandbox/contracts/deps/keepdeps`, `verbdeps`, `embeddeps`, `iodeps`, and `requestdeps` restate those libraries' public APIs field for field; the adapter converts values as they cross.

### Registries: one place each

- **Tasks** — one file per action in `sandbox/Tasks/Tasks/<Name>.go`, returning an `api.Task` (name, declared `[]api.Field`, `HandleAction` closure). Registered in exactly one place: `TaskArray()` in `sandbox/Tasks/run.go`. A task never parses its own input — `sandbox/lib/entries/` validates fields against the declaration and fills defaults. A task writes through the database it is handed, and nothing else. Repeated checks belong in `Tasks/Tasks/shared.go` so every task fails in the same words.
- **Visualizations** — one file per renderer in `sandbox/Visualization/Visualization/<Name>.go`, returning an `api.Visualizer`. Registered in `Catalog()` in that package's `catalog.go`. A visualization **returns files and never writes one**; putting bytes on disk is `sandbox/lib/vault/`. Pages are built with `page.go`'s markdown builder.

The CLI, the tick, and the `TaskList`/`Help` visualizations all read those two registries rather than lists of their own, so a name added in one place is reachable from every direction.

### Where logic belongs

- `sandbox/lib/ledger/` — every derived figure (balances, month results, `State`, `NetWorth`, forecast). Visualizations do layout, not arithmetic.
- `sandbox/lib/utils/` — `Pack`/`Unpack` for the database's packed unique keys, calendar arithmetic, money rendering. Shared by tasks, visualizations, and the ledger so a figure renders identically everywhere.
- `sandbox/lib/vault/` — the outside of a tick: read the task file, disarm it, read and validate the config, write under a `dest`, report failures in `Error.md`.
- `sandbox/lib/yaml/` — the hand-written YAML subset the two driving files use.
- `sandbox/config/` — compile-time constants: registry and field names, CLI text, flag spellings, defaults, version. Names the compiler can catch, so prefer a constant here over a literal. Long-form or editable text goes in `assets/` instead.
- `assets/` — files compiled in via a single `//go:embed all:*`; adding one needs no directive change. Carries `start/Task.yaml` and `start/Visualization.yaml` that `wraith start` writes.

Call the injected libraries (Keep, Verb) directly. Do not add wrapper layers over them; shared helpers belong in `sandbox/lib/utils/`.

## Documentation Rules

Every `.md` file follows the GeneralDoc spec: one H1, topic-driven headings with no skipped levels, `---` between top-level sections, relative links to the most specific anchor, no duplicated content, language-tagged code fences. Adding, renaming, or deleting a doc requires updating the referring index — [HandleDocuments.md](/docs/Tutorials/HandleDocuments.md) lists what to fix so no reference breaks. Creating or removing a spec means updating `Specs.md` in the same commit.
