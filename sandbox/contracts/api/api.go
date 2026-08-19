// Package api defines the output contracts of the sandbox, representing
// the objects and functions the library exposes to callers.
package api

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps/keepdeps"
)

// Exit codes api.Lib.Sandboxmain returns, and the process exits with.
const (
	// ExitOk reports a command that did what it was asked to.
	ExitOk = 0
	// ExitError reports a command that failed — an unknown task, an invalid
	// field, a file that could not be written.
	ExitError = 1
	// ExitUsage reports a command line that could not be understood, and is
	// answered with the usage screen.
	ExitUsage = 2
)

// HandleActionArgs is everything a task is handed when it runs: the injected
// deps, the registries it may write, and the fields it was called with.
type HandleActionArgs struct {
	// Deps is the dependency set the library was built with — the clock a
	// record is stamped with, and nothing else a task needs.
	Deps deps.Deps

	// DataBase is the schema database the task writes its registries in. A
	// task reaches storage through this field and no other, which is what
	// keeps a task from touching a file, a socket, or the terminal.
	DataBase keepdeps.KeepDatabase

	// Entries are the task's fields, keyed by the name they carry in
	// Task.yaml — `account`, `amount`, `date`. Values are strings, int64,
	// float64, bool, or nil, exactly as the file declared them.
	Entries map[string]any
}

// Task represents an executable action within the library.
// Its behavior is filled by the factories in sandbox/Tasks/Tasks/.
type Task struct {
	// Name is the identifier of the task — the value `Task.yaml.name`
	// carries, and the word `wraith run` takes.
	Name string

	// Description explains what the task does. It is what the Task-List
	// visualization renders, so no task can be documented without existing.
	Description string

	// Fields are the entries the task declares, in the order they are
	// documented. Validation of a task's input is driven from this list.
	Fields []Field

	// make in these way, ensure handleAction only allows to write in db
	HandleAction func(args HandleActionArgs) error
}

// Field is one entry a task or a visualization declares: its name, whether
// it is required, and what it holds.
type Field struct {
	// Name is the key the field carries in Task.yaml, and the `--flag` the
	// command line passes it as.
	Name string

	// Type is one of TextField, NumberField or BoolField.
	Type int

	// Required reports whether the task refuses to run without it.
	Required bool

	// Description explains what the field means.
	Description string

	// Default is the value used when the field is omitted. It is only read
	// for a field that is not required.
	Default any
}

// Field types, reported by Field.Type.
const (
	// TextField holds a string — a name, a date, a description.
	TextField = iota
	// NumberField holds a number, whole or decimal.
	NumberField
	// BoolField holds true or false.
	BoolField
)

// HandleVisualizationArgs is everything a visualization is handed when it
// renders: the injected deps, the registries it may read, and the args it was
// declared with.
type HandleVisualizationArgs struct {
	// Deps is the dependency set the library was built with. A visualization
	// reads the clock through it to know which month is the open one.
	Deps deps.Deps

	// DataBase is the schema database the visualization reads. It is handed
	// the same database a task writes, and never writes to it.
	DataBase keepdeps.KeepDatabase

	// Entries are the visualization's args, keyed by the name they carry
	// under `args:` in Visualization.yaml.
	Entries map[string]any
}

// VisualizationRender is one file a visualization produced: where it goes,
// relative to the entry's `dest`, and what it holds.
type VisualizationRender struct {
	// Path is the file's path below `dest`, slash-separated. A file
	// visualization returns a single render with an empty Path — `dest` is
	// the file itself.
	Path string

	// Content is the bytes to write.
	Content []byte
}

// Visualizer represents a component that generates visual output.
// Its behavior is filled by the factories in
// sandbox/Visualization/Visualization/.
type Visualizer struct {
	// Name is the identifier of the visualizer — the value a
	// Visualization.yaml entry's `name` carries.
	Name string

	// Description explains what the visualizer displays.
	Description string

	// Folder reports whether the visualizer renders a whole tree below
	// `dest` rather than one file at `dest`.
	Folder bool

	// Args are the options the visualizer declares, with their defaults.
	Args []Field

	// HandleVisualizer produces the visualization output — one render per
	// file it writes. It is filled by the factories in
	// sandbox/Visualization/Visualization/.
	HandleVisualizer func(args HandleVisualizationArgs) ([]VisualizationRender, error)
}

// VisualizationEntry is one line of Visualization.yaml: a visualizer, where
// it writes, and the args it writes with.
type VisualizationEntry struct {
	// Name is the visualizer asked for.
	Name string

	// Dest is where it writes, relative to the vault root.
	Dest string

	// Args are the per-entry options, overriding the catalog defaults.
	Args map[string]any

	// Enabled reports whether a tick renders the entry. It defaults to true.
	Enabled bool
}

// Lib is the entry point handed back by lib.New. Every object it
// creates carries the same deps it was built with.
type Lib struct {
	// Deps is the dependency set injected by lib.New, carried here so
	// every factory-built function field can reach it.
	Deps deps.Deps

	// DatabasePath is the folder the registries live in, relative to the
	// vault root. It is required on construction — lib.New takes it — and
	// the interface's `--database` flag overrides it for one invocation.
	DatabasePath string

	// TaskPath is the task file a tick reads and resets. The interface's
	// `--task` flag overrides it for one invocation.
	TaskPath string

	// VisualizationPath is the config file a tick renders from. The
	// interface's `--visualization` flag overrides it for one invocation.
	VisualizationPath string

	// Tasks is a collection of available tasks.
	Tasks []Task

	// Visualizations is a collection of available visualizers.
	Visualizations []Visualizer

	// PerformTask executes a specific task by name.
	// It is filled by lib.PerformTaskFactory.
	PerformTask func(taskName string, entries map[string]any) error

	// PerformVisualization executes a specific visualizer by name.
	// It is filled by lib.PerformVisualizationFactory.
	PerformVisualization func(visualizerName string, entries map[string]any) ([]VisualizationRender, error)

	// PerformTaskTick runs the periodic task tick operations.
	// It reads the task.yaml and runs the tasks based on the schedule.
	PerformTaskTick func() (string, error)

	// PerformVisualizationTick runs the periodic visualization tick operations.
	// It reads the visualization.yaml and runs the visualizers based on the schedule.
	PerformVisualizationTick func() error

	// PerformFullTick runs both task and visualization tick operations.
	PerformFullTick func() (string, error)

	// Start writes a default Task.yaml and Visualization.yaml, creating a
	// vault where there was none, and immediately runs a tick to render it.
	// It is filled by lib.StartFactory.
	Start func() error

	// Sandboxmain is the CLI entry point.
	// It is filled by lib.SandboxmainFactory.
	Sandboxmain func(args []string) int
}
