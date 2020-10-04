package strftime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Fmt formats the time t according to the spec s. The spec s is a string which
// may contain format sequences which begin with a '%' character and are
// terminated by some other character. Formatting functions are registered to
// special format sequences so that when a special format sequence is found it
// is replaced with the result of the formatting function.
func Fmt(s string, t time.Time) string {
	b := []byte(s)
	o := []string{}
	for i := 0; i < len(b); i++ {
		switch b[i] {
		// assume all '%' strings are format sequences --- strftime.get will do
		// the right thing and handle orinary sequences by returning the bytes
		// back unmodified
		case '%':
			f := string(b[i : i+2])
			o = append(o, strftime.get(f, t))
			i += 1
		default:
			o = append(o, string(b[i]))
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

// get safely returns a formatted value or returns the value if it can't be formatted
func (r *reg) get(f string, t time.Time) string {
	if _, ok := r.fns[f]; ok {
		return r.fns[f](t)
	}
	return f
}

func init() {
	strftime = &reg{
		fns: map[string]func(time.Time) string{
			"%a": func(t time.Time) string { return t.Format("Mon") },
			"%A": func(t time.Time) string { return t.Format("Monday") },
			"%b": func(t time.Time) string { return t.Format("Feb") },
			"%B": func(t time.Time) string { return t.Format("February") },
			"%c": func(t time.Time) string { return Fmt("%F %T", t) },
			"%C": func(t time.Time) string {
				y, err := strconv.Atoi(Fmt("%Y", t))
				if err != nil {
					y = 0
				}
				return fmt.Sprintf("%d", y/100)
			},
			"%d": func(t time.Time) string { return t.Format("02") },
			"%D": func(t time.Time) string { return Fmt("%m/%d/%y", t) },
			"%e": func(t time.Time) string { return t.Format("_2") },
			"%F": func(t time.Time) string { return Fmt("%Y-%m-%d", t) },
			"%g": func(t time.Time) string { return Fmt("%y", t) },
			"%G": func(t time.Time) string { return Fmt("%Y", t) },
			"%h": func(t time.Time) string { return Fmt("%b", t) },
			"%H": func(t time.Time) string { return t.Format("15") },
			"%I": func(t time.Time) string { return t.Format("03") },
			"%l": func(t time.Time) string { return t.Format("_3") },
			"%m": func(t time.Time) string { return t.Format("01") },
			"%M": func(t time.Time) string { return t.Format("04") },
			"%n": func(t time.Time) string { return "\n" },
			"%R": func(t time.Time) string { return t.Format("15:04") },
			"%s": func(t time.Time) string { return fmt.Sprintf("%d", t.Unix()) },
			"%S": func(t time.Time) string { return t.Format("05") },
			"%t": func(t time.Time) string { return "\t" },
			"%T": func(t time.Time) string { return Fmt("%H:%M:%S", t) },
			"%u": func(t time.Time) string {
				if t.Weekday() == 0 {
					return "7"
				}
				return string(t.Weekday())
			},
			"%w": func(t time.Time) string { return string(t.Weekday()) },
			"%x": func(t time.Time) string { return Fmt("%F", t) },
			"%X": func(t time.Time) string { return Fmt("%T", t) },

			"%Y": func(t time.Time) string { return t.Format("2006") },
			"%y": func(t time.Time) string { return t.Format("06") },
			"%%": func(t time.Time) string { return "%" },
		},
	}
}
