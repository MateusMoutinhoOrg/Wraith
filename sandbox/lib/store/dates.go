package store

// Dates, as the registries hold them and as the pages show them. A date is a
// whole number — 2026-08-18 is 20260818 — and a month is a whole number too —
// 2026-08 is 202608. Both sort correctly as numbers, which is what the month
// index and the forecast are built out of.
//
// `time` is the one standard package the sandbox may reach for here: it
// computes, it does not reach the operating system. The current time itself
// still arrives injected, through Deps.Now.

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ErrDate reports a date that is not written as YYYY-MM-DD, or a month that
// is not written as YYYY-MM.
var ErrDate = errors.New("invalid date")

// monthNames are the three-letter month names every rendered page uses.
var monthNames = [...]string{
	"jan", "feb", "mar", "apr", "may", "jun",
	"jul", "aug", "sep", "oct", "nov", "dec",
}

// ParseDate reads a date written as YYYY-MM-DD into the whole number the
// registries hold it as. A day that does not exist in its month — 2026-02-30
// — is rejected rather than rolled forward.
func ParseDate(text string) (int64, error) {
	parsed, err := time.Parse("2006-01-02", strings.TrimSpace(text))
	if err != nil {
		return 0, ErrDate
	}
	return DateOf(parsed), nil
}

// ParseMonth reads a month written as YYYY-MM into the whole number the
// registries hold it as.
func ParseMonth(text string) (int64, error) {
	parsed, err := time.Parse("2006-01", strings.TrimSpace(text))
	if err != nil {
		return 0, ErrDate
	}
	return MonthOf(DateOf(parsed)), nil
}

// DateOf converts a time into the stored date it falls on.
func DateOf(moment time.Time) int64 {
	year, month, day := moment.Date()
	return int64(year)*10000 + int64(month)*100 + int64(day)
}

// MonthOf returns the month a stored date belongs to.
func MonthOf(date int64) int64 {
	return date / 100
}

// DayOf returns the day of the month of a stored date.
func DayOf(date int64) int64 {
	return date % 100
}

// Split breaks a stored date into its year, month and day.
func Split(date int64) (year int, month int, day int) {
	return int(date / 10000), int((date / 100) % 100), int(date % 100)
}

// DateText renders a stored date as YYYY-MM-DD, the way a task file writes
// it.
func DateText(date int64) string {
	year, month, day := Split(date)
	return pad(int64(year), 4) + "-" + pad(int64(month), 2) + "-" + pad(int64(day), 2)
}

// PrettyDate renders a stored date the way the pages show it — `18-aug-2026`.
func PrettyDate(date int64) string {
	year, month, day := Split(date)
	return pad(int64(day), 2) + "-" + MonthName(int64(month)) + "-" + strconv.Itoa(year)
}

// MonthText renders a stored month as YYYY-MM, the way a task file writes it
// and the way a month folder is named.
func MonthText(month int64) string {
	return pad(month/100, 4) + "-" + pad(month%100, 2)
}

// PrettyMonth renders a stored month the way the pages show it —
// `aug-2026`.
func PrettyMonth(month int64) string {
	return MonthName(month%100) + "-" + strconv.FormatInt(month/100, 10)
}

// MonthName returns the three-letter name of a month number.
func MonthName(month int64) string {
	if month < 1 || month > 12 {
		return "???"
	}
	return monthNames[month-1]
}

// AddMonths returns the month that many months after the given one, counting
// backwards for a negative count.
func AddMonths(month int64, count int) int64 {
	total := (month/100)*12 + (month % 100) - 1 + int64(count)
	return (total/12)*100 + total%12 + 1
}

// LastDay returns the last day of a month — 28, 29, 30 or 31.
func LastDay(month int64) int64 {
	year, number, _ := Split(month * 100)
	first := time.Date(year, time.Month(number), 1, 0, 0, 0, 0, time.UTC)
	return int64(first.AddDate(0, 1, -1).Day())
}

// DateIn returns the date the given day of the month falls on, clamped to the
// month's last day when the day does not exist in it: day 31 in April is
// 30-apr, and day 30 in February is the 28th or the 29th. It is never skipped
// and never spills into the next month.
func DateIn(month int64, day int64) int64 {
	last := LastDay(month)
	if day > last {
		day = last
	}
	if day < 1 {
		day = 1
	}
	return month*100 + day
}

// pad renders a number with leading zeroes to the given width.
func pad(value int64, width int) string {
	digits := strconv.FormatInt(value, 10)
	for len(digits) < width {
		digits = "0" + digits
	}
	return digits
}
