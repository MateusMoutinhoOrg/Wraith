# Add a Task

## Description
Adds a new action to a brain — one file under `sandbox/Tasks/Tasks/`, one line in the registry. Adding a renderer instead is [HandleVisualizations.md](/docs/Tutorials/HandleVisualizations.md); turning this repository into a brain of your own is [ForkTemplate.md](/docs/Tutorials/ForkTemplate.md).

### Rules
- One task per file, named after the task: `sandbox/Tasks/Tasks/<Name>.go` returning an `api.Task`.
- A task declares its fields. It never parses a command line, never reads a file, and never prints — validation, flags and documentation are all generated from that declaration.
- A task writes through `args.DataBase` and nothing else. That is what makes the same task safe from a tick, from the command line, and from a test.
- A task that fails must fail **before** writing anything. Validate first, write last.

---

## Workflow

1. Create the file, named after the task.

```bash
touch sandbox/Tasks/Tasks/AddNote.go
```

2. Return an `api.Task` from a function with the same name. The `Fields` list is the whole interface of the task: it drives validation, the `--flags` the command line accepts, and the page the `Task-List` visualization writes.

```go
package tasks

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// AddNote returns the task that files one note under a title.
func AddNote() api.Task {
	return api.Task{
		Name:        "AddNote",
		Description: "File a note under a title",
		Fields: []api.Field{
			{Name: "title", Type: api.TextField, Required: true,
				Description: "What the note is called"},
			{Name: "body", Type: api.TextField,
				Description: "The note itself"},
		},
		HandleAction: func(args api.HandleActionArgs) error {
			title, err := name(args.Entries, "title")
			if err != nil {
				return err
			}
			body, err := optionalText(args.Entries, "body")
			if err != nil {
				return err
			}
			return insert(args, config.NoteSchema, "note "+title, map[string]any{
				config.NameField:   title,
				config.DetailField: utils.Pack(title, body),
			})
		},
	}
}
```

The helpers — `name`, `optionalText`, `schema`, `find`, `records`, `insert`, `remove`, `requireAccount` — live in [`shared.go`](/sandbox/Tasks/Tasks/shared.go) beside the tasks, so every task reaches the injected database the same way and reports a missing field or a name already taken in the same words.

3. If the task writes something the registries do not hold yet, add its schema in [`database.go`](/sandbox/config/database.go) — a name field, plus whatever it carries. Keep the injected database in mind: it holds unique string keys and whole numbers, so money goes in as cents, dates go in as `20260818`, and free text travels packed into one key with `utils.Pack`.

4. Register the task in [`sandbox/Tasks/run.go`](/sandbox/Tasks/run.go). This is the only registration there is:

```go
func TaskArray() []api.Task {
	return []api.Task{
		// …
		tasks.AddNote(),
	}
}
```

5. Build and check it is there:

```bash
go build ./...
go run ./cmd/main tasks
```

6. Run it — from the command line, and from a task file:

```bash
wraith run AddNote --title Groceries --body "milk, bread"
```

```yaml
name: AddNote
title: Groceries
body: milk, bread
apply: true
```

7. Run a tick. The `Task-List` visualization writes `Tasks/AddNote.md` by itself, with the fields, the sample and the command line — you do not write that page.

```bash
wraith tick
```
