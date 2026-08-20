package visualizations

// The Help visualization: the three guides, generated from the registries the
// binary carries. Like Task-List, no guide here can describe a command, a
// task or a visualization that does not exist.

import (
	task "github.com/MateusMoutinhoOrg/Wraith/sandbox/Tasks"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Help returns the visualization that writes the guide tree:
//
//	Help/
//	├── Usage.md            every command and the tick workflow
//	├── Task.md             how to run a task, and every task there is
//	└── Visualization.md    how to choose what gets rendered
func Help() api.Visualizer {
	return api.Visualizer{
		Name:        "Help",
		Description: "The guides: the commands, how tasks work, how visualizations work",
		Folder:      true,
		HandleVisualizer: func(args api.HandleVisualizationArgs) ([]api.VisualizationRender, error) {
			return []api.VisualizationRender{
				usageDocument("Usage.md"),
				taskGuide(),
				visualizationGuide(),
			}, nil
		},
	}
}

// taskGuide writes Help/Task.md: what a task is, how to run one, and the
// catalog of every one the binary carries.
func taskGuide() api.VisualizationRender {
	p := &page{}
	p.heading(1, "Task Guide")
	p.line("Everything this brain does, it does because a task told it to. This page is how to " +
		"run one, and what there is to run.")
	p.blank()
	p.line("[Usage](Usage.md) · [Visualization](Visualization.md)")
	p.blank()
	p.rule()

	p.heading(2, "1. How to run a task")
	p.line("1. Pick a task from the table below and open its guide — every guide ends with a " +
		"**Sample**.")
	p.line("2. Copy the sample into `Task.yaml` and fill in your values.")
	p.line("3. Make sure `apply: true` is set. With `apply: false` the task is ignored, and " +
		"that is not an error — it is a task waiting to be armed.")
	p.line("4. Run a tick:")
	p.blank()
	p.line("```bash")
	p.line("wraith tick              # run once")
	p.line("wraith watch --time 1s   # or keep watching, and just save the file")
	p.line("```")
	p.blank()
	p.line("On success every visualization declared in `Visualization.yaml` is re-rendered and " +
		"`apply` goes back to `false`. On failure an `Error.md` is written and **nothing** is " +
		"changed — a failed task never leaves the data half-written.")
	p.blank()
	p.line("A task can also be run without editing any file, which is how a script drives this " +
		"brain:")
	p.blank()
	p.line("```bash")
	p.line("wraith run AddTransaction --account Bank --category Food --amount -32.90 " +
		"--date 2026-08-18")
	p.line("```")
	p.blank()
	p.rule()

	p.heading(2, "2. Every task")
	p.table("Task", "What it does", "Fields", "Guide")
	for _, declared := range task.TaskArray() {
		p.row("`"+declared.Name+"`", declared.Description,
			count(len(declared.Fields), "field", "fields"),
			"[open](../Tasks/"+declared.Name+".md)")
	}
	p.blank()
	p.line("The guides live wherever you pointed the `Task-List` visualization; the links above " +
		"assume the default `Tasks/`.")
	p.blank()
	p.rule()

	p.heading(2, "3. Rules to remember")
	p.line("- One task per tick — `Task.yaml` holds a single action.")
	p.line("- Rendered pages are **generated**. A hand edit is overwritten on the next tick, so " +
		"everything you want to keep goes into a task.")
	p.line("- One category and one account per transaction — there are no splits.")
	p.line("- Every movement settles on the day it is dated, in full. There is nothing to " +
		"record as owed and nothing to pay later.")
	p.line("- A positive amount needs a category with `revenues: true`; a negative one needs " +
		"`expenses: true`. A transfer category — both `false` — accepts either, because it " +
		"counts as neither.")
	p.line("- Moving money between your own accounts is two transactions sharing a transfer " +
		"category. They net to zero, so a transfer never shows up as an expense.")
	p.line("- A recurrence describes the future and moves nothing. Nothing is recorded until " +
		"you record it.")
	return p.render("Task.md")
}

// visualizationGuide writes Help/Visualization.md: what a visualization is,
// how `Visualization.yaml` is shaped, and the catalog of every renderer.
func visualizationGuide() api.VisualizationRender {
	p := &page{}
	p.heading(1, "Visualization Guide")
	p.line("A task changes the data. A **visualization** renders it. Which pages this brain " +
		"writes, and where, is not built in — you declare it in `Visualization.yaml`.")
	p.blank()
	p.line("[Usage](Usage.md) · [Task](Task.md)")
	p.blank()
	p.rule()

	p.heading(2, "1. How `Visualization.yaml` works")
	p.line("The file is a **list**. Every entry asks for one visualization, rendered to one " +
		"destination:")
	p.blank()
	p.line("```yaml")
	p.line("- name: DashBoard")
	p.line("  args:")
	p.line("    prev-months: 3")
	p.line("    future-months: 8")
	p.line("  dest: DashBoard")
	p.line("")
	p.line("- name: Task-List")
	p.line("  dest: Tasks")
	p.line("```")
	p.blank()
	p.table("Field", "Required", "Description")
	p.row("`name`", "yes", "A visualization from the catalog below. An unknown name is an error.")
	p.row("`dest`", "yes", "Where to write it, relative to the vault root. Missing folders are created.")
	p.row("`args`", "no", "Per-visualization options. Anything omitted falls back to the default.")
	p.row("`enabled`", "no", "`false` silences the entry without deleting it. Defaults to `true`.")
	p.blank()
	p.line("What is not in `Visualization.yaml` is not rendered. There are no implicit pages.")
	p.blank()
	p.rule()

	p.heading(2, "2. What `dest` means")
	p.table("Kind", "`dest` is", "What gets written")
	p.row("file", "the path of the file, extension included", "that one file")
	p.row("folder", "the path of a folder", "every page the visualization produces, in the tree it defines")
	p.blank()
	p.line("A folder visualization owns its internal layout: rename its `dest` and the same " +
		"tree appears somewhere else. A folder `dest` is written **into**, never emptied — a " +
		"page that stops being produced sits there frozen at its last render until you delete it.")
	p.blank()
	p.rule()

	p.heading(2, "3. The catalog")
	p.table("Visualization", "Kind", "Args", "What it renders")
	for _, declared := range Catalog() {
		p.row("`"+declared.Name+"`", kindOf(declared), argList(declared), declared.Description)
	}
	p.blank()
	for _, declared := range Catalog() {
		if len(declared.Args) == 0 {
			continue
		}
		p.heading(3, "`"+declared.Name+"` args")
		p.table("Arg", ">Default", "Description")
		for _, arg := range declared.Args {
			p.row("`"+arg.Name+"`", "`"+scalar(arg.Default)+"`", arg.Description)
		}
		p.blank()
	}
	p.rule()

	p.heading(2, "4. Rendering one entry")
	p.line("```bash")
	p.line("wraith render DashBoard --future-months 24")
	p.line("```")
	p.blank()
	p.line("It writes that one entry and nothing else: `Task.yaml` is not executed and no other " +
		"entry is re-rendered. The destination and the defaults come from the matching entry — " +
		"`--dest` overrides where it goes, and is required when the name is declared nowhere.")
	return p.render("Visualization.md")
}

// kindOf renders whether a visualization writes a file or a folder.
func kindOf(declared api.Visualizer) string {
	if declared.Folder {
		return "folder"
	}
	return "file"
}

// argList renders the args a visualization declares, as one cell.
func argList(declared api.Visualizer) string {
	if len(declared.Args) == 0 {
		return "—"
	}
	names := ""
	for index, arg := range declared.Args {
		if index > 0 {
			names += ", "
		}
		names += "`" + arg.Name + "`"
	}
	return names
}
