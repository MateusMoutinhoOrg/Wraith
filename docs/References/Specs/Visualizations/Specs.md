# Visualizations Specification

## Description
Defines the required shape of a visualization file in `sandbox/Visualization/Visualization/<Name>.go` — one renderer the brain can write. A visualization is a **value**: it declares what it is called, what it shows, whether it owns a folder or a file, and the closure that renders it. The actions that change what it shows are governed by [Tasks](/docs/References/Specs/Tasks/Specs.md) instead; the workflow of adding one is [HandleVisualizations.md](/docs/Tutorials/HandleVisualizations.md).

### Rules
- One visualization per file, in package `visualizations`, named after it: `Summary.go` declaring `func Summary() api.Visualizer`.
- The function takes no arguments and returns a fully populated `api.Visualizer`.
- `Name` matches the file name and the function name. A name carrying a hyphen — `Task-List` — keeps the hyphen in `Name` and drops it in the file and function names.
- `HandleVisualizer` **returns files** and writes none. Putting bytes on disk is `sandbox/lib/vault`'s job, which is what lets the same renderer serve a tick, a single `wraith render`, and a caller that only wants the bytes.
- It reads `args.DataBase` and never writes to it.
- `Folder: true` returns one render per file, each with a `Path` relative to the entry's `dest`. `Folder: false` returns exactly **one** render with an empty `Path`, because `dest` is the file itself.
- Every option is declared in `Args` with a `Default`. A visualization never fails over an arg — the config was validated before it was called — so an unreadable value falls back to the default rather than returning an error.
- No arithmetic. Every figure comes from `sandbox/lib/ledger` or the registry views in `sandbox/lib/store`, so a visualization is about layout and can be changed without touching what the numbers mean.
- Markdown is built with the `page` helper in [page.go](/sandbox/Visualization/Visualization/page.go), never with raw string concatenation of table syntax.
- Adding, renaming, or deleting a visualization requires updating `Catalog` in [catalog.go](/sandbox/Visualization/Visualization/catalog.go) in the same commit. Nothing else has to learn about it: the switcher, the command line and the generated `Help` catalog all read that list.

## Structure
1. **Package clause**: `package visualizations`.
2. **Imports**: `sandbox/contracts/api`, plus the internal packages the pages read — usually `sandbox/lib/ledger` and `sandbox/lib/store`.
3. **Arg constants** *(optional)*: the arg names and their defaults, when the visualization declares any.
4. **Doc comment**: what the visualization shows, and — for a folder — the tree it writes, drawn out.
5. **Constructor**: `func <Name>() api.Visualizer` returning the value, with `HandleVisualizer` last.
6. **Page functions**: one unexported function per page it writes, each returning an `api.VisualizationRender`.

> **Note**: For a concrete example, refer to [sample.go](./sample.go).
