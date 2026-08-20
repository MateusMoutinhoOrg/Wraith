# Add a Visualization

## Description
Adds a new renderer to a brain — one file under `sandbox/Visualization/Visualization/`, one line in the catalog. Adding an action instead is [HandleTasks.md](/docs/Tutorials/HandleTasks.md); choosing which renderers a vault actually writes is [ChooseVisualizations.md](/docs/Tutorials/ChooseVisualizations.md).

### Rules
- One visualization per file, named after it: `sandbox/Visualization/Visualization/<Name>.go` returning an `api.Visualizer`.
- A visualization **returns files**; it never writes one. Putting bytes on disk happens in [`sandbox/lib/vault`](/sandbox/lib/vault/vault.go), which is what lets the same renderer serve a tick, a `wraith render`, and a caller that only wants the markdown.
- A visualization reads the database and never writes to it.
- `Folder: true` renders a whole tree below the entry's `dest`, one `VisualizationRender` per file. `Folder: false` renders a single file **at** `dest`, and returns one render with an empty `Path`.

---

## Workflow

1. Create the file, named after the visualization.

```bash
touch sandbox/Visualization/Visualization/Summary.go
```

2. Return an `api.Visualizer`. The `Args` list drives validation, the `--flags` `wraith render` accepts, and the catalog the `Help` guide writes.

```go
package visualizations

import (
	"strconv"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/ledger"
)

// Summary returns the visualization that writes one page: where you stand,
// in as few lines as it can be said.
func Summary() api.Visualizer {
	return api.Visualizer{
		Name:        "Summary",
		Description: "Where you stand, on a single page",
		Folder:      false,
		Args: []api.Field{
			{Name: "top", Type: api.NumberField,
				Description: "How many accounts to list",
				Default:     int64(5)},
		},
		HandleVisualizer: func(args api.HandleVisualizationArgs) ([]api.VisualizationRender, error) {
			state := ledger.Load(args.Deps, args.DataBase)
			p := &page{}
			p.heading(1, "Summary")
			p.table("Account", ">Balance")
			for index, account := range state.Accounts {
				if index >= wholeArg(args.Entries, "top", 5) {
					break
				}
				p.row(account.Name, money(state.Balance(account)))
			}
			p.blank()
			p.line("**Money held:** " + money(state.Held()) + " · " +
				strconv.Itoa(len(state.Transactions)) + " movements")
			return []api.VisualizationRender{p.render("")}, nil
		},
	}
}
```

The markdown builder — `heading`, `table`, `row`, `line`, `render` — and the arg helpers live in [`page.go`](/sandbox/Visualization/Visualization/page.go). Every figure comes from [`sandbox/lib/ledger`](/sandbox/lib/ledger/ledger.go), so a page is about layout and never about arithmetic.

3. Register it in the catalog, [`catalog.go`](/sandbox/Visualization/Visualization/catalog.go):

```go
func Catalog() []api.Visualizer {
	return []api.Visualizer{
		// …
		Summary(),
	}
}
```

4. Build and check it is there:

```bash
go build ./...
go run ./cmd/main visualizations
```

5. Render it. A visualization the config does not declare has no destination to fall back on, so give it one:

```bash
wraith render Summary --dest Summary.md --top 3
```

6. Ask for it on every tick by adding an entry to `Visualization.yaml`:

```yaml
- name: Summary
  args:
    top: 3
  dest: Summary.md
```

7. Run a tick. The `Help` guide lists your visualization and its args by itself — that catalog is generated from the same array you edited in step 3.

```bash
wraith tick
```
