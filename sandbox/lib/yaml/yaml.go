package yaml

// The subset of YAML the brain reads and writes: `Task.yaml` — a flat
// mapping of fields — and `Visualization.yaml` — a sequence of mappings, one
// of which may nest an `args` mapping. Nothing wider is supported, and
// nothing wider is needed: both files are configuration a beginner types by
// hand, so the grammar is deliberately the smallest one that covers the
// samples in the guides.
//
// Parsing lives inside the sandbox rather than behind a dep because it is
// pure computation over bytes — no clock, no file, no socket — and because a
// brain that could not read its own task file without an adapter would not be
// a closed sandbox at all.
//
// This package declares the Value type and plain functions over it; it is
// neither an object package nor a lib function package, so it carries no
// factories.

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// Supported scalars, as they arrive in a decoded mapping:
//
//	string   — a bare or quoted word
//	int64    — a whole number
//	float64  — a number with a decimal point
//	bool     — true or false
//	nil      — null, ~, or an empty value
//
// Every consumer coerces from there; see sandbox/lib/entries.

// ErrShape reports a document whose shape is not the one asked for — a
// sequence where a mapping was expected, or the other way round.
var ErrShape = errors.New("unexpected document shape")

// line is one significant line of a document: how deep it is indented, and
// what it says once the comment and the surrounding spaces are gone.
type line struct {
	indent int
	text   string
}

// scan splits a document into its significant lines, dropping blank lines and
// whole-line comments and recording each remaining line's indentation. Tabs
// are counted as one column, which is enough: a YAML document may not indent
// with tabs, so a file that does is malformed however it is measured.
func scan(content []byte) []line {
	lines := []line{}
	for _, raw := range strings.Split(string(content), "\n") {
		raw = strings.TrimRight(raw, " \t\r")
		trimmed := strings.TrimLeft(raw, " \t")
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		lines = append(lines, line{indent: len(raw) - len(trimmed), text: trimmed})
	}
	return lines
}

// stripComment removes a trailing `# comment` from a scalar, leaving a `#`
// that sits inside quotes alone. The samples in the task guides annotate
// values this way — `opening: 0   # opening balance` — so a decoder that kept
// the comment would hand the task a field it cannot parse.
func stripComment(text string) string {
	quote := byte(0)
	for i := 0; i < len(text); i++ {
		c := text[i]
		switch {
		case quote != 0:
			if c == quote {
				quote = 0
			}
		case c == '\'' || c == '"':
			quote = c
		case c == '#' && (i == 0 || text[i-1] == ' ' || text[i-1] == '\t'):
			return strings.TrimRight(text[:i], " \t")
		}
	}
	return text
}

// Scalar decodes one value as it is written in a document: a quoted string
// verbatim, `true`/`false` as a bool, `null`/`~`/empty as nil, a number as
// int64 or float64, and anything else as the string it reads as.
func Scalar(text string) any {
	text = strings.TrimSpace(stripComment(text))
	if text == "" || text == "null" || text == "~" || text == "Null" || text == "NULL" {
		return nil
	}
	if len(text) >= 2 && (text[0] == '"' || text[0] == '\'') && text[len(text)-1] == text[0] {
		return text[1 : len(text)-1]
	}
	switch text {
	case "true", "True", "TRUE", "yes", "on":
		return true
	case "false", "False", "FALSE", "no", "off":
		return false
	}
	if whole, err := strconv.ParseInt(text, 10, 64); err == nil {
		return whole
	}
	if decimal, err := strconv.ParseFloat(text, 64); err == nil {
		return decimal
	}
	return text
}

// splitKey cuts a `key: value` line into its key and the rest of the line,
// which is empty when the value opens a nested block. found is false when the
// line carries no key at all.
func splitKey(text string) (key string, rest string, found bool) {
	key, rest, found = strings.Cut(text, ":")
	if !found {
		return "", "", false
	}
	return strings.TrimSpace(key), strings.TrimSpace(rest), true
}

// mapping decodes the block of lines at the given indent into a map, reading
// a deeper block below a keyless key as a nested mapping. It returns the map
// and how many lines it consumed.
func mapping(lines []line, start int, indent int) (map[string]any, int) {
	decoded := map[string]any{}
	i := start
	for i < len(lines) {
		current := lines[i]
		if current.indent < indent {
			break
		}
		key, rest, found := splitKey(current.text)
		if !found {
			break
		}
		i++
		if rest != "" {
			decoded[key] = Scalar(rest)
			continue
		}
		if i < len(lines) && lines[i].indent > current.indent && !strings.HasPrefix(lines[i].text, "- ") {
			nested, consumed := mapping(lines, i, lines[i].indent)
			decoded[key] = nested
			i = consumed
			continue
		}
		decoded[key] = nil
	}
	return decoded, i
}

// DecodeMap reads a document that is a single flat mapping — the shape of
// `Task.yaml`. A nested block is decoded as a nested map[string]any. An empty
// document decodes to an empty map rather than an error, since a file with
// every line commented out is still a valid, empty task.
func DecodeMap(content []byte) (map[string]any, error) {
	lines := scan(content)
	if len(lines) > 0 && strings.HasPrefix(lines[0].text, "-") {
		return nil, ErrShape
	}
	decoded, _ := mapping(lines, 0, 0)
	return decoded, nil
}

// DecodeList reads a document that is a sequence of mappings — the shape of
// `Visualization.yaml`. Every item must be a `- key: value` block; a document
// that is a mapping instead is reported as ErrShape.
func DecodeList(content []byte) ([]map[string]any, error) {
	lines := scan(content)
	if len(lines) == 0 {
		return []map[string]any{}, nil
	}
	if !strings.HasPrefix(lines[0].text, "-") {
		return nil, ErrShape
	}
	items := []map[string]any{}
	i := 0
	for i < len(lines) {
		if !strings.HasPrefix(lines[i].text, "-") {
			return nil, ErrShape
		}
		// The first field of an item shares its line with the dash, so the
		// dash is replaced by spaces and the block is decoded as a mapping
		// starting at the column the field actually begins in.
		head := lines[i]
		inner := strings.TrimSpace(strings.TrimPrefix(head.text, "-"))
		block := []line{}
		if inner != "" {
			block = append(block, line{indent: head.indent + 2, text: inner})
		}
		i++
		for i < len(lines) && !strings.HasPrefix(lines[i].text, "-") && lines[i].indent > head.indent {
			block = append(block, lines[i])
			i++
		}
		if len(block) == 0 {
			continue
		}
		decoded, _ := mapping(block, 0, block[0].indent)
		items = append(items, decoded)
	}
	return items, nil
}

// EncodeScalar renders one value the way this package reads it back: a
// missing value as `null`, a bool as `true`/`false`, a number bare, and a
// string quoted only when leaving it bare would change what it decodes to.
func EncodeScalar(value any) string {
	switch typed := value.(type) {
	case nil:
		return "null"
	case bool:
		return strconv.FormatBool(typed)
	case int:
		return strconv.Itoa(typed)
	case int64:
		return strconv.FormatInt(typed, 10)
	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64)
	case string:
		if typed == "" || Scalar(typed) != any(typed) || strings.ContainsAny(typed, ":#\"'") {
			return "\"" + strings.ReplaceAll(typed, "\"", "\\\"") + "\""
		}
		return typed
	}
	return "null"
}

// EncodeMap renders a flat mapping, one `key: value` per line, in the order
// the keys are given. A key the map does not carry is skipped, so the caller
// controls both the order and the presence of every line.
func EncodeMap(values map[string]any, order []string) []byte {
	out := strings.Builder{}
	for _, key := range order {
		value, found := values[key]
		if !found {
			continue
		}
		out.WriteString(key)
		out.WriteString(": ")
		out.WriteString(EncodeScalar(value))
		out.WriteString("\n")
	}
	return []byte(out.String())
}

// EncodeList renders a sequence of mappings — the shape of
// `Visualization.yaml` — as one `- key: value` block per item, each key in
// the order the caller gives and a key the item does not carry skipped. A
// value that is itself a mapping is written as the nested block this subset
// supports, its own keys in alphabetical order so the same config is always
// written the same way. An empty nested mapping is left out rather than
// written as a header with nothing under it.
func EncodeList(items []map[string]any, order []string) []byte {
	out := strings.Builder{}
	for index, item := range items {
		if index > 0 {
			out.WriteString("\n")
		}
		// The first key of an item shares its line with the dash; every key
		// after it is indented to the column that dash opened.
		lead := "- "
		for _, key := range order {
			value, found := item[key]
			if !found {
				continue
			}
			nested, isMapping := value.(map[string]any)
			if !isMapping {
				out.WriteString(lead + key + ": " + EncodeScalar(value) + "\n")
				lead = "  "
				continue
			}
			if len(nested) == 0 {
				continue
			}
			out.WriteString(lead + key + ":\n")
			lead = "  "
			for _, nestedKey := range keys(nested) {
				out.WriteString("    " + nestedKey + ": " + EncodeScalar(nested[nestedKey]) + "\n")
			}
		}
	}
	return []byte(out.String())
}

// keys returns a mapping's keys in alphabetical order.
func keys(values map[string]any) []string {
	ordered := []string{}
	for key := range values {
		ordered = append(ordered, key)
	}
	sort.Strings(ordered)
	return ordered
}
