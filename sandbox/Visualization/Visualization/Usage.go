package visualizations

// The Usage visualization: the command reference and the tick workflow, as
// one file. It is the same page the Help tree writes, offered on its own for
// a vault that wants the reference somewhere else — a README at the vault
// root, say — and none of the other guides.

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Usage returns the visualization that writes the command reference to a
// single file.
func Usage() api.Visualizer {
	return api.Visualizer{
		Name:        "Usage",
		Description: "The command reference and the tick workflow, on its own",
		Folder:      false,
		HandleVisualizer: func(args api.HandleVisualizationArgs) ([]api.VisualizationRender, error) {
			return []api.VisualizationRender{usageDocument("")}, nil
		},
	}
}

// usageDocument renders the command reference at the given path. Both the
// Usage visualization and the Help tree call it, which is what keeps the two
// from ever disagreeing: a file visualization writes to its dest, so its path
// is empty, while inside the Help tree the same bytes land on `Usage.md`.
func usageDocument(path string) api.VisualizationRender {
	p := &page{}
	p.heading(1, "Wraith")
	p.line("Wraith applies actions to your data and renders the current state of it through " +
		"whatever dashboards you asked for. It is a small state machine driven by two files:")
	p.blank()
	p.line("1. A task is declared in `Task.yaml` — or passed straight to `wraith run`.")
	p.line("2. Wraith executes it against the data.")
	p.line("3. Every visualization declared in `Visualization.yaml` is re-rendered.")
	p.blank()
	p.line("`Task.yaml` decides what changes; `Visualization.yaml` decides what you get to see.")
	p.blank()
	p.rule()

	p.heading(2, "1. Commands")
	p.line("```")
	p.line("wraith <command> [arguments] [flags]")
	p.line("```")
	p.blank()
	p.table("Command", "Arguments", "Writes", "What it is for")
	p.row("`start`", "—", "`Task.yaml`, `Visualization.yaml`", "Create a vault where there was none")
	p.row("`tick`", "—", "the data, every `dest`", "Run the pending task and re-render everything")
	p.row("`watch`", "—", "the data, every `dest`", "Run a tick on an interval, until stopped")
	p.row("`run`", "`<task>`", "the data, every `dest`", "Run one task from the command line")
	p.row("`render`", "`<visualization>`", "one `dest`", "Render one visualization to disk")
	p.row("`tasks`", "—", "nothing", "List every task the binary carries")
	p.row("`visualizations`", "—", "nothing", "List every visualization it can render")
	p.row("`help`", "—", "nothing", "Print the usage screen")
	p.row("`version`", "—", "nothing", "Print the interface version")
	p.blank()

	p.table("Flag", ">Default", "What it points at")
	p.row("`--task`", "`"+config.DefaultTaskPath+"`", "The task file a tick reads and resets")
	p.row("`--visualization`", "`"+config.DefaultVisualizationPath+"`", "The config a tick renders from")
	p.row("`--database`", "`"+config.DefaultDatabasePath+"`", "The folder the data lives in")
	p.row("`--dest`", "the entry's `dest`", "Where a single render is written")
	p.row("`--time`", "—", "How long `watch` sleeps between ticks. Required by `watch`")
	p.row("`--<field>`", "—", "One flag per field a task or a visualization declares")
	p.blank()
	p.line("The three path flags are what point the same binary at a second vault without " +
		"moving a file. Every `dest` stays relative to the vault root, not to `--database`.")
	p.blank()
	p.rule()

	p.heading(2, "2. The tick workflow")
	p.line("1. Read the task file. If it does not exist, stop — there is nothing to do.")
	p.line("2. Check `apply`. If it is `false`, stop. That is not an error.")
	p.line("3. Check the task's name and its fields. An unknown task, an unknown field or a " +
		"field of the wrong type writes `" + config.ErrorPath + "` and stops.")
	p.line("4. Execute the task. If it fails, write `" + config.ErrorPath + "` and stop — " +
		"nothing is changed.")
	p.line("5. Render every enabled entry of the visualization config.")
	p.line("6. Set `apply` back to `false`, so the same task is not run twice.")
	p.blank()
	p.line("A failed tick leaves the data exactly as it was. There is no half-applied task.")
	p.blank()
	p.rule()

	p.heading(2, "3. Examples")
	p.line("```bash")
	p.line("wraith start                       # create Task.yaml and Visualization.yaml")
	p.line("wraith tick                        # apply Task.yaml, render everything")
	p.line("wraith watch --time 1s             # keep watching, and just save the file")
	p.line("wraith run AddAccount --account Bank --opening 1200")
	p.line("wraith render DashBoard --future-months 24")
	p.line("wraith tick --database vaults/home --task Inbox/Task.yaml")
	p.line("```")
	return p.render(path)
}
