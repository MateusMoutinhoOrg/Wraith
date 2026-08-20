package entries

// The fields a task or a visualization was called with, read back as typed
// values. The same map arrives from three places — a `Task.yaml` a person
// typed, a `--flag` on the command line, and an `args:` block of
// `Visualization.yaml` — so every coercion and every error message a beginner
// will read lives here rather than in each of the twelve tasks.
//
// Reading is forgiving in one direction only: a number written as text is
// accepted, because a command line has no way to say otherwise, but a word
// where a number belongs is a reported error rather than a zero. Nothing here
// guesses.
//
// This package declares no types and carries no factories; the tasks in
// sandbox/Tasks/Tasks and the visualizations in
// sandbox/Visualization/Visualization call into it.

import (
	"errors"
	"strconv"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// The keys every task carries whatever it declares: the task's own name, and
// the switch a tick reads before running anything.
const (
	// NameKey is the task asked for.
	NameKey = "name"
	// ApplyKey is the switch a tick honours — `apply: false` is not an error,
	// it is a task waiting to be armed.
	ApplyKey = "apply"
)

// Present reports whether the map carries a usable value under the key. A
// key written with no value — `parent: null` — counts as absent, which is
// what lets a sample be copied whole and filled in line by line.
func Present(values map[string]any, key string) bool {
	value, found := values[key]
	return found && value != nil
}

// Text reads a value as a string. Numbers and bools are rendered rather than
// refused, so a date typed without quotes and a `--flag` carrying a name are
// read the same way.
func Text(values map[string]any, key string) (string, error) {
	value, found := values[key]
	if !found || value == nil {
		return "", nil
	}
	switch typed := value.(type) {
	case string:
		return strings.TrimSpace(typed), nil
	case bool:
		return strconv.FormatBool(typed), nil
	case int64:
		return strconv.FormatInt(typed, 10), nil
	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64), nil
	}
	return "", errors.New(key + " must be text")
}

// Number reads a value as a number, accepting one written as text.
func Number(values map[string]any, key string) (float64, error) {
	value, found := values[key]
	if !found || value == nil {
		return 0, nil
	}
	switch typed := value.(type) {
	case int64:
		return float64(typed), nil
	case int:
		return float64(typed), nil
	case float64:
		return typed, nil
	case string:
		parsed, err := strconv.ParseFloat(strings.TrimSpace(typed), 64)
		if err != nil {
			return 0, errors.New(key + " must be a number, not " + strconv.Quote(typed))
		}
		return parsed, nil
	}
	return 0, errors.New(key + " must be a number")
}

// Whole reads a value as a whole number, refusing one that carries a
// fraction — a day of the month and a record's id are counts, not
// measurements.
func Whole(values map[string]any, key string) (int64, error) {
	number, err := Number(values, key)
	if err != nil {
		return 0, err
	}
	whole := int64(number)
	if float64(whole) != number {
		return 0, errors.New(key + " must be a whole number")
	}
	return whole, nil
}

// Amount reads a value as a monetary amount, accepting text or a whole number,
// and returning the exact whole number of cents it stands for. It rejects
// amounts outside the +/- 10^14 cents range to prevent overflow, and refuses
// any value with precision below the cent.
func Amount(values map[string]any, key string) (int64, error) {
	value, found := values[key]
	if !found || value == nil {
		return 0, nil
	}
	var text string
	switch typed := value.(type) {
	case int64:
		text = strconv.FormatInt(typed, 10)
	case int:
		text = strconv.Itoa(typed)
	case float64:
		text = strconv.FormatFloat(typed, 'f', -1, 64)
	case string:
		text = strings.TrimSpace(typed)
	default:
		return 0, errors.New(key + " must be a number")
	}
	
	// We use utils.ParseCents to parse the text safely.
	return utils.ParseCents(text)
}

// Bool reads a value as true or false, accepting the words a person types
// into a task file.
func Bool(values map[string]any, key string) (bool, error) {
	value, found := values[key]
	if !found || value == nil {
		return false, nil
	}
	switch typed := value.(type) {
	case bool:
		return typed, nil
	case string:
		switch strings.ToLower(strings.TrimSpace(typed)) {
		case "true", "yes", "on", "1":
			return true, nil
		case "false", "no", "off", "0", "":
			return false, nil
		}
	}
	return false, errors.New(key + " must be true or false")
}

// Validate checks a map against the fields a task or a visualization
// declares: every required field is present, every field carries the type it
// declared, and no field the declaration does not know about was given. It is
// what makes a typo in a task file a reported error instead of a silently
// ignored line.
func Validate(fields []api.Field, values map[string]any) error {
	known := map[string]api.Field{}
	for _, field := range fields {
		known[field.Name] = field
	}
	for key := range values {
		if key == NameKey || key == ApplyKey {
			continue
		}
		if _, found := known[key]; !found {
			return errors.New("unknown field: " + key + " — " + accepted(fields))
		}
	}
	for _, field := range fields {
		if !Present(values, field.Name) {
			if field.Required {
				return errors.New("missing required field: " + field.Name + " — " + field.Description)
			}
			continue
		}
		if err := check(field, values); err != nil {
			return err
		}
	}
	return nil
}

// check reports whether one given field carries the type it declared.
func check(field api.Field, values map[string]any) error {
	switch field.Type {
	case api.NumberField:
		_, err := Number(values, field.Name)
		return err
	case api.BoolField:
		_, err := Bool(values, field.Name)
		return err
	}
	_, err := Text(values, field.Name)
	return err
}

// accepted lists the fields a declaration knows about, for the message a
// misspelled field is answered with.
func accepted(fields []api.Field) string {
	names := []string{}
	for _, field := range fields {
		names = append(names, field.Name)
	}
	if len(names) == 0 {
		return "this one takes no fields"
	}
	return "accepted fields are " + strings.Join(names, ", ")
}

// WithDefaults returns the values a task or visualization actually runs
// with: what it was given, plus the declared default of every optional field
// it was not.
func WithDefaults(fields []api.Field, values map[string]any) map[string]any {
	filled := map[string]any{}
	for key, value := range values {
		filled[key] = value
	}
	for _, field := range fields {
		if field.Default == nil || Present(filled, field.Name) {
			continue
		}
		filled[field.Name] = field.Default
	}
	return filled
}
