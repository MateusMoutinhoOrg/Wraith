package commands

import (
	visualization "github.com/MateusMoutinhoOrg/Wraith/sandbox/Visualization"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/config"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/vault"
)

// Render writes one visualization and nothing else: the task file is not
// executed and no other entry is re-rendered. It is to the visualization
// config what `run` is to the task file — the whole tick narrowed down to one
// thing, driven from the command line.
//
// Where it writes and what its args default to come from the matching entry
// in the config; any flag given here overrides that entry for this invocation
// only, and the file itself is never edited. `enabled: false` does not block
// it: asking for an entry by name is the decision to render it.
func Render(l *api.Lib, quiet bool) int {
	name, err := l.Deps.VerbLib.GetNextStringArg()
	if err != nil || name == "" {
		return UsageError(l, config.MissingVisualizationName)
	}
	declared, found := visualization.Find(name)
	if !found {
		return UsageError(l, config.UnknownVisualization, name)
	}

	dest, args := declaredEntry(l, declared.Name)
	if override, err := l.Deps.VerbLib.GetStringOption([]string{config.DestFlag}, 0); err == nil && override != "" {
		dest = override
	}
	if dest == "" {
		return UsageError(l, config.MissingDest, declared.Name, l.VisualizationPath)
	}
	if !vault.Inside(dest) {
		return Failure(l, errOutside(dest))
	}
	for key, value := range Flags(l, declared.Args) {
		args[key] = value
	}

	renders, err := l.PerformVisualization(declared.Name, args)
	if err != nil {
		return Failure(l, err)
	}
	if err := vault.Write(l.Deps, dest, declared.Folder, renders); err != nil {
		return Failure(l, err)
	}
	if !quiet {
		l.Deps.Printf(config.Rendered, declared.Name, dest)
	}
	return api.ExitOk
}

// declaredEntry looks a visualization up in the config to find where it
// writes and what its args were declared as. A name the config never mentions
// has no destination to fall back on, which is what makes `--dest` required
// for it.
func declaredEntry(l *api.Lib, name string) (string, map[string]any) {
	args := map[string]any{}
	entries, err := vault.ReadEntries(l.Deps, l.VisualizationPath, l.Visualizations)
	if err != nil {
		return "", args
	}
	for _, entry := range entries {
		if entry.Name != name {
			continue
		}
		for key, value := range entry.Args {
			args[key] = value
		}
		return entry.Dest, args
	}
	return "", args
}
