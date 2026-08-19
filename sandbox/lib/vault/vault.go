package vault

// The vault: the two files a tick is driven by, and the tree it writes.
//
// Everything here is the *outside* of a tick — reading `Task.yaml`, resetting
// its `apply` switch, reading `Visualization.yaml`, putting rendered bytes on
// disk, reporting a failure in `Error.md`. The tasks and the visualizations
// themselves never touch any of it: they are handed a database and hand back
// records or bytes, which is what makes them testable and what keeps the
// filesystem in one place.
//
// Every path below is reached through the injected filesystem library
// (Deps.IoLib), so the sandbox stays closed.

import (
	"errors"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/yaml"
)

// The paths of the defaults inside the embedded asset tree. A vault created
// by `wraith start` is a copy of these two files, so what a new brain starts
// life with is editable content rather than a Go constant.
const (
	// StartTaskAsset is the default task file.
	StartTaskAsset = "start/Task.yaml"
	// StartVisualizationAsset is the default visualization config.
	StartVisualizationAsset = "start/Visualization.yaml"
)

// The keys a visualization entry is written with.
const (
	// NameKey is the visualization asked for.
	NameKey = "name"
	// DestKey is where it writes.
	DestKey = "dest"
	// ArgsKey opens the per-entry options.
	ArgsKey = "args"
	// EnabledKey silences an entry without deleting it.
	EnabledKey = "enabled"
)

// ErrNoTask reports a task file that is not there. It is not a failure: a
// vault with no task file has simply not been started yet, and a tick over it
// does nothing.
var ErrNoTask = errors.New("no task file")

// Task is one pending action, as the task file declares it.
type Task struct {
	// Name is the task asked for.
	Name string
	// Apply reports whether the tick should run it.
	Apply bool
	// Values are the fields it was declared with, the task's own `name` and
	// `apply` included.
	Values map[string]any
}

// ReadTask reads the pending action out of the task file. A file that is not
// there is reported as ErrNoTask; a file that cannot be parsed is a failure
// like any other.
func ReadTask(d deps.Deps, path string) (Task, error) {
	if !d.IoLib.IsFile(path) {
		return Task{}, ErrNoTask
	}
	content, err := d.IoLib.ReadFile(path)
	if err != nil {
		return Task{}, errors.New("could not read " + path + ": " + err.Error())
	}
	values, err := yaml.DecodeMap(content)
	if err != nil {
		return Task{}, errors.New(path + " is not a task file: it must be a list of " +
			"`field: value` lines")
	}
	name, err := entries.Text(values, entries.NameKey)
	if err != nil {
		return Task{}, err
	}
	apply, err := entries.Bool(values, entries.ApplyKey)
	if err != nil {
		return Task{}, err
	}
	return Task{Name: name, Apply: apply, Values: values}, nil
}

// Fields returns the task's fields with the two keys every task carries taken
// out, which is what a task's own validation is run against.
func (t Task) Fields() map[string]any {
	fields := map[string]any{}
	for key, value := range t.Values {
		if key == entries.NameKey || key == entries.ApplyKey {
			continue
		}
		fields[key] = value
	}
	return fields
}

// ResetApply writes the task file back with `apply: false`, leaving every
// other line as it was. It is what stops a tick from running the same action
// again on the next one — and it runs whether the task succeeded or failed,
// so a broken task is not retried forever.
func ResetApply(d deps.Deps, path string, task Task) error {
	content, err := d.IoLib.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		trimmed := strings.TrimLeft(line, " \t")
		if strings.HasPrefix(trimmed, "apply:") {
			colon := strings.Index(line, ":")
			if colon != -1 {
				before := line[:colon+1]
				after := line[colon+1:]
				replaced := false
				for _, truth := range []string{"true", "yes", "on", "1", "True", "TRUE"} {
					wordStart := -1
					for j, c := range after {
						if c != ' ' && c != '\t' {
							wordStart = j
							break
						}
					}
					if wordStart != -1 && strings.HasPrefix(after[wordStart:], truth) {
						after = after[:wordStart] + "false" + after[wordStart+len(truth):]
						lines[i] = before + after
						replaced = true
						break
					}
				}
				if replaced {
					break
				}
			}
		}
	}
	return d.IoLib.WriteFile(path, []byte(strings.Join(lines, "\n")))
}

// ReadEntries reads the visualization config. A file that is not there is
// created from the embedded default and then read, so a vault always renders
// something rather than reporting a missing file.
func ReadEntries(d deps.Deps, path string, known []api.Visualizer) ([]api.VisualizationEntry, error) {
	if !d.IoLib.IsFile(path) {
		if err := writeAsset(d, StartVisualizationAsset, path); err != nil {
			return nil, err
		}
	}
	content, err := d.IoLib.ReadFile(path)
	if err != nil {
		return nil, errors.New("could not read " + path + ": " + err.Error())
	}
	items, err := yaml.DecodeList(content)
	if err != nil {
		return nil, errors.New(path + " is not a visualization config: it must be a list of " +
			"`- name:` entries")
	}
	parsed := []api.VisualizationEntry{}
	for _, item := range items {
		entry, err := readEntry(item, path)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, entry)
	}
	if err := check(parsed, known, path); err != nil {
		return nil, err
	}
	return parsed, nil
}

// readEntry reads one entry of the visualization config.
func readEntry(item map[string]any, path string) (api.VisualizationEntry, error) {
	name, err := entries.Text(item, NameKey)
	if err != nil {
		return api.VisualizationEntry{}, err
	}
	if name == "" {
		return api.VisualizationEntry{}, errors.New(path + ": an entry carries no name")
	}
	dest, err := entries.Text(item, DestKey)
	if err != nil {
		return api.VisualizationEntry{}, err
	}
	if dest == "" {
		return api.VisualizationEntry{}, errors.New(path + ": " + name + " carries no dest")
	}
	enabled := true
	if entries.Present(item, EnabledKey) {
		enabled, err = entries.Bool(item, EnabledKey)
		if err != nil {
			return api.VisualizationEntry{}, err
		}
	}
	args := map[string]any{}
	if nested, ok := item[ArgsKey].(map[string]any); ok {
		args = nested
	}
	return api.VisualizationEntry{Name: name, Dest: dest, Args: args, Enabled: enabled}, nil
}

// check validates the whole config before a single file is written: every
// name is one the binary carries, every dest stays inside the vault, and no
// two destinations collide or nest inside one another. A broken config
// renders nothing rather than half a vault.
func check(parsed []api.VisualizationEntry, known []api.Visualizer, path string) error {
	for _, entry := range parsed {
		found := false
		for _, candidate := range known {
			if candidate.Name == entry.Name {
				found = true
			}
		}
		if !found {
			return errors.New(path + ": unknown visualization: " + entry.Name)
		}
		if !Inside(entry.Dest) {
			return errors.New(path + ": " + entry.Name + " writes outside the vault: " + entry.Dest)
		}
	}
	for outer := range parsed {
		for inner := outer + 1; inner < len(parsed); inner++ {
			if overlaps(parsed[outer].Dest, parsed[inner].Dest) {
				return errors.New(path + ": " + parsed[outer].Name + " and " +
					parsed[inner].Name + " write to the same place — " +
					parsed[outer].Dest + " and " + parsed[inner].Dest)
			}
		}
	}
	return nil
}

// overlaps reports whether two destinations are the same, or one sits inside
// the other — either way the second render would erase part of the first.
func overlaps(first string, second string) bool {
	first = clean(first)
	second = clean(second)
	if first == second {
		return true
	}
	if first == "." || second == "." {
		return true
	}
	return strings.HasPrefix(first, second+"/") || strings.HasPrefix(second, first+"/")
}

// Inside reports whether a destination stays inside the vault. A path
// climbing out of it with `..`, or starting from the root, is refused.
func Inside(dest string) bool {
	if strings.HasPrefix(dest, "/") || strings.Contains(dest, "\\") {
		return false
	}
	for _, part := range strings.Split(clean(dest), "/") {
		if part == ".." {
			return false
		}
	}
	return true
}

// clean normalizes a destination for comparison, so `./DashBoard/` and
// `DashBoard` are recognized as the same place.
func clean(dest string) string {
	trimmed := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(dest), "./"), "/")
	if trimmed == "" || trimmed == "." {
		return "."
	}
	return trimmed
}

// Write puts the files one visualization produced on disk, under the entry's
// dest. A folder visualization writes each render below dest; a file
// visualization writes its single render to dest itself.
//
// A folder dest is written into, never emptied: a file the visualization no
// longer produces is left where it is rather than deleted.
func Write(d deps.Deps, dest string, folder bool, renders []api.VisualizationRender) error {
	for _, render := range renders {
		path := dest
		if folder {
			path = join(dest, render.Path)
		}
		if err := d.IoLib.WriteFile(path, render.Content); err != nil {
			return errors.New("could not write " + path + ": " + err.Error())
		}
	}
	return nil
}

// join puts a rendered file's path below a destination folder.
func join(dest string, path string) string {
	base := clean(dest)
	if base == "." {
		return path
	}
	if path == "" {
		return base
	}
	return base + "/" + path
}

// WriteError reports a failure in the vault itself, as the `Error.md` the
// guides tell a beginner to look for. It is written where the person will see
// it, not to a log they do not know about.
func WriteError(d deps.Deps, path string, when string, failure error) {
	page := strings.Builder{}
	page.WriteString("# Error\n\n")
	page.WriteString("> **When:** " + when + "\n\n")
	page.WriteString("```\n" + failure.Error() + "\n```\n\n")
	page.WriteString("Nothing was changed. Fix the cause above and run another tick — this " +
		"file is overwritten on the next failure, and left behind once the tick succeeds.\n")
	d.IoLib.WriteFile(path, []byte(page.String()))
}

// ClearError removes the report of a failure that has been fixed, so an
// `Error.md` on disk always describes the last tick and not an older one.
func ClearError(d deps.Deps, path string) {
	if !d.IoLib.IsFile(path) {
		return
	}
	d.IoLib.WriteFile(path, []byte("# Error\n\nNothing is wrong. The last tick succeeded.\n"))
}

// WriteAsset copies one embedded default into the vault, without overwriting
// a file that is already there.
func WriteAsset(d deps.Deps, asset string, path string) (bool, error) {
	if d.IoLib.Exist(path) {
		return false, nil
	}
	return true, writeAsset(d, asset, path)
}

// writeAsset copies one embedded default into the vault, overwriting.
func writeAsset(d deps.Deps, asset string, path string) error {
	content, err := d.EmbedDeps.ReadFile(asset)
	if err != nil {
		return errors.New("the default " + asset + " is missing from this binary")
	}
	if err := d.IoLib.WriteFile(path, content); err != nil {
		return errors.New("could not write " + path + ": " + err.Error())
	}
	return nil
}
