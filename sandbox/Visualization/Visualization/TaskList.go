package visualizations

// The Task-List visualization: one reference page per task the binary
// carries, generated from the task registry itself.
//
// That is the whole reason it exists. A guide written by hand drifts from the
// code the day someone adds a field; a guide generated from the declaration
// cannot describe a task that does not exist, or miss one that does.

import (
	"strconv"
	"strings"

	task "github.com/MateusMoutinhoOrg/Wraith/sandbox/Tasks"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
)

// TaskList returns the visualization that writes one page per task, plus the
// index that links them.
func TaskList() api.Visualizer {
	return api.Visualizer{
		Name:        "Task-List",
		Description: "One reference page per task the binary carries",
		Folder:      true,
		HandleVisualizer: func(args api.HandleVisualizationArgs) ([]api.VisualizationRender, error) {
			renders := []api.VisualizationRender{taskIndex()}
			for _, declared := range task.TaskArray() {
				renders = append(renders, taskPage(declared))
			}
			return renders, nil
		},
	}
}

// taskIndex writes README.md: every task, what it does, and where its page
// is.
func taskIndex() api.VisualizationRender {
	p := &page{}
	p.heading(1, "Tasks")
	p.line("Every action this brain can perform. A task is declared in `Task.yaml` and picked " +
		"up on the next tick, or run straight from the command line with `wraith run`.")
	p.blank()
	p.rule()
	p.table("Task", "What it does", "Page")
	for _, declared := range task.TaskArray() {
		p.row("`"+declared.Name+"`", declared.Description,
			"["+declared.Name+".md]("+declared.Name+".md)")
	}
	p.blank()
	p.rule()
	p.heading(2, "How to run one")
	p.line("1. Open the task's page and copy its **Sample** into `Task.yaml`.")
	p.line("2. Fill in your values, and make sure `apply: true` is set.")
	p.line("3. Run `wraith tick` — or leave `wraith watch --time 1s` running and just save the file.")
	p.blank()
	p.line("On success every visualization is re-rendered and `apply` goes back to `false`. On " +
		"failure an `Error.md` is written and nothing is changed.")
	return p.render("README.md")
}

// taskPage writes one task's reference page: its fields, and a sample ready
// to be copied into Task.yaml.
func taskPage(declared api.Task) api.VisualizationRender {
	p := &page{}
	p.heading(1, declared.Name)
	p.line(declared.Description + ".")
	p.blank()
	p.line("[All tasks](README.md)")
	p.blank()
	p.rule()

	p.heading(2, "Fields")
	p.table("Field", "Required", "Type", "Description")
	p.row("`"+entries.NameKey+"`", "yes", "text", "Must be `"+declared.Name+"`")
	for _, field := range declared.Fields {
		p.row("`"+field.Name+"`", required(field), typeName(field.Type),
			field.Description+defaulted(field))
	}
	p.row("`"+entries.ApplyKey+"`", "yes", "true or false",
		"Set `true` to execute on the next tick")
	p.blank()
	p.rule()

	p.heading(2, "Sample")
	p.line("```yaml")
	p.line(entries.NameKey + ": " + declared.Name)
	for _, field := range declared.Fields {
		p.line(field.Name + ": " + sample(field))
	}
	p.line(entries.ApplyKey + ": true")
	p.line("```")
	p.blank()
	p.rule()

	p.heading(2, "From the command line")
	p.line("The same task, without touching `Task.yaml` — one flag per field:")
	p.blank()
	p.line("```bash")
	p.line("wraith run " + declared.Name + commandFlags(declared))
	p.line("```")
	return p.render(declared.Name + ".md")
}

// required renders whether a field must be given.
func required(field api.Field) string {
	if field.Required {
		return "yes"
	}
	return "no"
}

// defaulted appends a field's default to its description, when it declares
// one.
func defaulted(field api.Field) string {
	if field.Required || field.Default == nil {
		return ""
	}
	return ". Defaults to `" + scalar(field.Default) + "`"
}

// typeName renders a field's type in the words a task file uses.
func typeName(kind int) string {
	switch kind {
	case api.NumberField:
		return "number"
	case api.BoolField:
		return "true or false"
	}
	return "text"
}

// sample renders a plausible value for a field, so a copied sample is a
// working task once the words are replaced.
func sample(field api.Field) string {
	if field.Default != nil {
		return scalar(field.Default)
	}
	if !field.Required {
		return "null"
	}
	switch field.Type {
	case api.NumberField:
		return "0"
	case api.BoolField:
		return "false"
	}
	return "your " + strings.ReplaceAll(field.Name, "_", " ")
}

// scalar renders a declared default the way a task file writes it.
func scalar(value any) string {
	switch typed := value.(type) {
	case bool:
		return strconv.FormatBool(typed)
	case int64:
		return strconv.FormatInt(typed, 10)
	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64)
	case string:
		return typed
	}
	return "null"
}

// commandFlags renders the required fields of a task as the flags a command
// line passes them with.
func commandFlags(declared api.Task) string {
	flags := strings.Builder{}
	for _, field := range declared.Fields {
		if !field.Required {
			continue
		}
		flags.WriteString(" --" + field.Name + " " + quoted(sample(field)))
	}
	return flags.String()
}

// quoted wraps a sample value in quotes when it carries a space, so the
// rendered command line can be pasted as it is.
func quoted(value string) string {
	if strings.ContainsAny(value, " \t") {
		return "\"" + value + "\""
	}
	return value
}
