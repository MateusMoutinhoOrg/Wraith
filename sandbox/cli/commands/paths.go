package commands

import (
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
)

// Paths reads the three flags that point a command at a vault, overriding the
// library's defaults for this invocation only. Every command takes them, and
// every command takes them the same way — which is what lets one binary drive
// as many vaults as you have folders.
//
// They are read before any positional argument: the Verb parser marks a
// matched flag used, so draining what is left afterwards hands back only the
// command's own words.
func Paths(l *api.Lib) {
	verb := l.Deps.VerbLib
	if value, err := verb.GetStringOption([]string{config.TaskFlag}, 0); err == nil && value != "" {
		l.TaskPath = value
	}
	if value, err := verb.GetStringOption([]string{config.VisualizationFlag}, 0); err == nil && value != "" {
		l.VisualizationPath = value
	}
	if value, err := verb.GetStringOption([]string{config.DatabaseFlag}, 0); err == nil && value != "" {
		l.DatabasePath = value
	}
}

// Flags reads one `--<name> <value>` pair for every field a task or a
// visualization declares, and returns them as the entry map that task or
// visualization is run with. A flag that was not given is simply absent, so
// the declared defaults still apply.
//
// Values arrive as text, because a command line has no other kind. The same
// validator that checks a task file turns them into numbers and switches, so
// `--amount -32.90` and `amount: -32.90` reach a task as the same value.
func Flags(l *api.Lib, fields []api.Field) map[string]any {
	verb := l.Deps.VerbLib
	entries := map[string]any{}
	for _, field := range fields {
		value, err := verb.GetStringOption([]string{"--" + field.Name}, 0)
		if err != nil {
			continue
		}
		entries[field.Name] = value
	}
	return entries
}
