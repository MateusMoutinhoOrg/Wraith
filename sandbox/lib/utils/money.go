package utils

// Rendering money the way every page of the vault shows it. Amounts are held
// in cents as whole numbers, so nothing here rounds and nothing drifts: the
// formatting is the only place a value becomes text.

import (
	"errors"
	"strconv"
	"strings"
)

// Money renders an amount held in cents the way the vault shows it —
// `R$ 4,694`, `-R$ 350`, `R$ 1,234.56`. The cents are printed only when they
// are not zero, which is what keeps a table of round figures readable.
func Money(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		if cents == -9223372036854775808 { // math.MinInt64
			// Cannot negate MinInt64, so handle manually or since we bounded to 10^14 it's safe.
			// However, a good formatter should be robust.
			return "-R$ 92,233,720,368,547,758.08"
		}
		cents = -cents
	}
	return sign + "R$ " + grouped(cents/100) + fraction(cents%100)
}

// Signed renders an amount the way a change is shown — with an explicit `+`
// in front of a positive figure, so a column of movements reads at a glance.
func Signed(cents int64) string {
	if cents > 0 {
		return "+" + Money(cents)
	}
	return Money(cents)
}

// fraction renders the cents of an amount, as an empty string when there are
// none.
func fraction(cents int64) string {
	if cents == 0 {
		return ""
	}
	if cents < 10 {
		return ".0" + strconv.FormatInt(cents, 10)
	}
	return "." + strconv.FormatInt(cents, 10)
}

// grouped renders a whole number with a comma every three digits.
func grouped(whole int64) string {
	digits := strconv.FormatInt(whole, 10)
	if len(digits) <= 3 {
		return digits
	}
	out := make([]byte, 0, len(digits)+len(digits)/3)
	lead := len(digits) % 3
	if lead > 0 {
		out = append(out, digits[:lead]...)
	}
	for i := lead; i < len(digits); i += 3 {
		if len(out) > 0 {
			out = append(out, ',')
		}
		out = append(out, digits[i:i+3]...)
	}
	return string(out)
}

// Cents converts an amount written as a number in a task file into the whole
// number of cents it stands for. It rounds to the nearest cent, away from
// zero, so `-75.005` becomes -7501 rather than a value that depends on how
// the float was stored.
func Cents(amount float64) int64 {
	if amount < 0 {
		return -int64(-amount*100 + 0.5)
	}
	return int64(amount*100 + 0.5)
}

// ParseCents converts a string representing a monetary amount exactly into cents.
// It refuses amounts with more than two decimal places, preventing silent
// precision loss, and enforces a +/- 10^14 range to prevent overflow.
func ParseCents(text string) (int64, error) {
	if text == "" {
		return 0, nil
	}
	sign := int64(1)
	if text[0] == '-' {
		sign = -1
		text = text[1:]
	} else if text[0] == '+' {
		text = text[1:]
	}
	
	parts := strings.Split(text, ".")
	if len(parts) > 2 {
		return 0, errors.New("invalid number format")
	}
	
	var whole int64
	var err error
	if parts[0] != "" {
		whole, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return 0, errors.New("invalid whole number")
		}
	}
	
	var frac int64
	if len(parts) == 2 {
		if len(parts[1]) > 2 {
			return 0, errors.New("precision below the cent is discarded")
		}
		if len(parts[1]) == 1 {
			parts[1] += "0"
		}
		if len(parts[1]) == 2 {
			frac, err = strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return 0, errors.New("invalid decimal fraction")
			}
		}
	}
	
	cents := whole*100 + frac
	if cents > 100000000000000 || cents < -100000000000000 {
		return 0, errors.New("amount is outside the allowed range")
	}
	return cents * sign, nil
}

// Bar renders a share of a total as the twenty-cell bar the dashboards use —
// `███████████░░░░░░░░░`.
func Bar(part int64, total int64) string {
	const cells = 20
	filled := 0
	if total > 0 && part > 0 {
		filled = int((part*cells + total/2) / total)
	}
	if filled > cells {
		filled = cells
	}
	bar := ""
	for i := 0; i < cells; i++ {
		if i < filled {
			bar += "█"
			continue
		}
		bar += "░"
	}
	return bar
}

// Percent renders a share of a total as a whole percentage, `0` when the
// total is zero.
func Percent(part int64, total int64) string {
	if total == 0 {
		return "0%"
	}
	return strconv.FormatInt((part*100+total/2)/total, 10) + "%"
}
