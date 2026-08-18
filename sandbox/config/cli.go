package config

// Every word the interface says, as compile-time constants. No display text
// is written anywhere else in the sandbox: sandbox/cli addresses these by
// name, so a renamed constant is a build failure rather than a blank line at
// runtime.

// The defaults a command falls back to when its path flags are not given.
// They are the same three names the guides use, so a vault created by
// `wraith start` works with no flags at all.
const (
	// DefaultTaskPath is the task file a tick reads and resets.
	DefaultTaskPath = "Task.yaml"
	// DefaultVisualizationPath is the config file a tick renders from.
	DefaultVisualizationPath = "Visualization.yaml"
	// DefaultDatabasePath is the folder the registries are persisted in.
	DefaultDatabasePath = "data"
	// ErrorPath is the file a failed tick writes its report to.
	ErrorPath = "Error.md"
)

// The flags every command accepts, in the form the injected Verb parser
// matches them.
var (
	// HelpFlags print the usage screen and exit.
	HelpFlags = []string{"-h", "--help"}
	// VersionFlags print the interface version and exit.
	VersionFlags = []string{"-v", "--version"}
	// QuietFlags silence everything but listings and errors.
	QuietFlags = []string{"-q", "--quiet"}
)

// The named flags read as `--flag value`.
const (
	// TaskFlag points a command at another task file.
	TaskFlag = "--task"
	// VisualizationFlag points a command at another visualization config.
	VisualizationFlag = "--visualization"
	// DatabaseFlag points a command at another database folder.
	DatabaseFlag = "--database"
	// DestFlag overrides where a single render is written.
	DestFlag = "--dest"
	// TimeFlag is the interval `watch` sleeps between ticks.
	TimeFlag = "--time"
)

const (
	// Usages is the screen printed by `wraith help`, by `--help`, and by an
	// empty command line.
	Usages = `wraith — a second brain you drive with two files

Usage:
  wraith <command> [arguments] [flags]

Commands:
  start                              create Task.yaml and Visualization.yaml
  tick                               run the pending task and render everything
  watch --time <interval>            run a tick on an interval, until stopped
  run <task> [--<field> <value>]     run one task from the command line
  render <visualization> [--<arg>]   render one visualization to its dest
  tasks                              list every task the binary carries
  visualizations                     list every visualization it can render
  help                               print this screen
  version                            print the interface version

Flags:
  --task <path>                      task file to read        (Task.yaml)
  --visualization <path>             config to render from    (Visualization.yaml)
  --database <path>                  folder the data lives in (data)
  --dest <path>                      where a single render is written
  --time <interval>                  how long watch sleeps between ticks
  -h, --help                         print this screen and exit
  -v, --version                      print the interface version and exit
  -q, --quiet                        print only listings and errors

A task changes the data; a visualization renders it. Everything wraith writes
is one entry you declared in Visualization.yaml.
`

	// NoCommand answers a command line carrying no command word.
	NoCommand = "no command given\n"
	// UnknownCommand answers a command word the interface does not know.
	UnknownCommand = "unknown command: %s\n"
	// UnknownTask answers a task name the binary does not carry.
	UnknownTask = "unknown task: %s\n"
	// UnknownVisualization answers a visualization the binary cannot render.
	UnknownVisualization = "unknown visualization: %s\n"
	// MissingTaskName answers `run` called without a task name.
	MissingTaskName = "run needs a task name\n"
	// MissingVisualizationName answers `render` called without a name.
	MissingVisualizationName = "render needs a visualization name\n"
	// MissingInterval answers `watch` called without --time.
	MissingInterval = "watch needs --time, such as --time 1s\n"
	// InvalidInterval answers a --time that is not a duration.
	InvalidInterval = "not an interval: %s\n"
	// MissingDest answers a render of an entry that is declared nowhere and
	// was given no destination either.
	MissingDest = "%s is not declared in %s — give it a --dest\n"
	// Failed reports an error the command could not recover from.
	Failed = "error: %s\n"

	// Started confirms a vault was created.
	Started = "created %s and %s\n"
	// AlreadyStarted reports a vault that was already there, and was left
	// exactly as it was.
	AlreadyStarted = "%s already exists — nothing was changed\n"
	// Ticked confirms one full tick.
	Ticked = "tick done\n"
	// Ran confirms one task run from the command line.
	Ran = "%s done\n"
	// Rendered confirms one visualization written to disk.
	Rendered = "rendered %s to %s\n"
	// Watching announces the watch loop and the interval it runs on.
	Watching = "watching every %s — press ctrl-c to stop\n"

	// TaskListHeader titles the `tasks` listing.
	TaskListHeader = "Tasks:\n"
	// VisualizationListHeader titles the `visualizations` listing.
	VisualizationListHeader = "Visualizations:\n"
	// ListEntry is one line of either listing.
	ListEntry = "  %-20s %s\n"
)
