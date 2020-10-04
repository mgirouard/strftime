package strftime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Fmt formats a string with strftime
func Fmt(s string, t time.Time) string {
	b := []byte(s)
	o := []string{}
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '%':
			f := string(b[i : i+2])
			o = append(o, strftime.get(f, t))
			i += 1
		default:
			o = append(o, string(b[i:i+1]))
		}
	}
	return strings.Join(o, "")
}

// strftime is a global registry of formatter funcs
var strftime *reg

// reg is a registry of format funcs
type reg struct {
	fns map[string]func(time.Time) string
}

// get returns a formatted value or returns the value if it can't be formatted
func (r *reg) get(f string, t time.Time) string {
	if _, ok := r.fns[f]; ok {
		return r.fns[f](t)
	}
	return f
}

func init() {
	strftime = &reg{
		fns: map[string]func(time.Time) string{
			// The abbreviated name of the day of the week
			"%a": func(t time.Time) string { return t.Format("Mon") },

			// The full name of the day of the week
			"%A": func(t time.Time) string { return t.Format("Monday") },

			// The abbreviated month name
			"%b": func(t time.Time) string { return t.Format("Feb") },

			// The full month name
			"%B": func(t time.Time) string { return t.Format("February") },

			// The preferred date and time representation
			"%c": func(t time.Time) string { return Fmt("%a %b %e %H:%M:%S %Y", t) },

			// The century number (year/100) as a 2-digit integer
			"%C": func(t time.Time) string {
				y, err := strconv.Atoi(Fmt("%Y", t))
				if err != nil {
					y = 0
				}
				return fmt.Sprintf("%d", y/100)
			},
			"%d": func(t time.Time) string { return t.Format("02") },
			"%D": func(t time.Time) string { return Fmt("%m/%d/%y", t) },
			"%e": func(t time.Time) string { return Fmt(" 2", t) },

			"%m": func(t time.Time) string { return t.Format("01") },
			"%Y": func(t time.Time) string { return t.Format("2006") },
			"%y": func(t time.Time) string { return t.Format("06") },

			// A literal '%' character.
			"%%": func(t time.Time) string { return "%" },
		},
	}
}
