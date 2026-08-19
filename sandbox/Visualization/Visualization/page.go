// Package visualizations holds every renderer the brain can write, one
// visualization per file.
//
// A visualization is a value, not a function: it carries its name, what it
// shows, whether it writes one file or a whole folder, the args it accepts,
// and the closure that renders it. It is handed the database and hands back
// bytes — it never writes a file itself, so the same renderer serves a tick,
// a single `wraith render`, and a test that only looks at what came out.
//
// Adding a visualization is adding one file here that returns an
// api.Visualizer, and one line in Catalog. See
// docs/Tutorials/HandleVisualizations.md.
package visualizations

// The small markdown builder every page in this package is written with, and
// the naming rules its links follow.

import (
	"strconv"
	"strings"

	"github.com/MateusMoutinhoOrg/Wraith/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/entries"
	"github.com/MateusMoutinhoOrg/Wraith/sandbox/lib/utils"
)

// page accumulates one rendered markdown file.
type page struct {
	// out is the text written so far.
	out strings.Builder
}

// line writes one line of text.
func (p *page) line(text string) {
	p.out.WriteString(text)
	p.out.WriteString("\n")
}

// blank writes an empty line.
func (p *page) blank() { p.out.WriteString("\n") }

// heading writes a heading of the given level.
func (p *page) heading(level int, text string) {
	p.line(strings.Repeat("#", level) + " " + text)
	p.blank()
}

// rule writes a horizontal rule between sections.
func (p *page) rule() {
	p.line("---")
	p.blank()
}

// table writes a header row and its alignment row. Columns whose name is
// prefixed with `>` are right-aligned, which is what every column of money
// in the vault is.
func (p *page) table(columns ...string) {
	header := []string{}
	divider := []string{}
	for _, column := range columns {
		if strings.HasPrefix(column, ">") {
			header = append(header, strings.TrimPrefix(column, ">"))
			divider = append(divider, "---:")
			continue
		}
		header = append(header, column)
		divider = append(divider, "---")
	}
	p.line("| " + strings.Join(header, " | ") + " |")
	p.line("| " + strings.Join(divider, " | ") + " |")
}

// row writes one row of the table opened by table.
func (p *page) row(cells ...string) {
	p.line("| " + strings.Join(cells, " | ") + " |")
}

// render closes the page into the file it becomes, at the given path below
// the entry's dest.
func (p *page) render(path string) api.VisualizationRender {
	return api.VisualizationRender{Path: path, Content: []byte(p.out.String())}
}

// slug turns a display name into the file name it is written under, so
// `Emergency Savings` becomes `Emergency-Savings.md` and a link to it never
// has to be escaped.
func slug(name string) string {
	safe := strings.Builder{}
	for _, letter := range name {
		switch {
		case letter == ' ' || letter == '/' || letter == '\\':
			safe.WriteRune('-')
		case letter == '.' || letter == ':' || letter == '#' || letter == '?':
			continue
		default:
			safe.WriteRune(letter)
		}
	}
	trimmed := strings.Trim(safe.String(), "-")
	if trimmed == "" {
		return "unnamed"
	}
	return trimmed
}

// count renders a whole number with the word it counts, in the singular when
// there is one of it.
func count(number int, singular string, plural string) string {
	if number == 1 {
		return "1 " + singular
	}
	return strconv.Itoa(number) + " " + plural
}

// dash renders empty text as an em dash, so a table cell is never blank.
func dash(text string) string {
	if strings.TrimSpace(text) == "" {
		return "—"
	}
	return text
}

// wholeArg reads one of a visualization's args as a whole number, falling
// back to the declared default when it is missing or unreadable. A
// visualization never fails over an arg: the config was validated before it
// was called.
func wholeArg(values map[string]any, key string, fallback int) int {
	if !entries.Present(values, key) {
		return fallback
	}
	number, err := entries.Whole(values, key)
	if err != nil || number < 0 {
		return fallback
	}
	return int(number)
}

// money renders an amount the way every page shows it.
func money(cents int64) string { return utils.Money(cents) }

// signed renders an amount with an explicit sign.
func signed(cents int64) string { return utils.Signed(cents) }

// idText renders a transaction's permanent identifier, which is what a
// ModifyTransaction task addresses it by.
func idText(id int64) string { return strconv.FormatInt(id, 10) }
