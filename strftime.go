package strftime

import (
	"strings"
	"time"
)

// strftime is a global registry of formatter funcs
var strftime *reg

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

func init() {
	strftime = &reg{
		fns: map[string]func(time.Time) string{
			"%%": func(t time.Time) string { return "%" },
			"%a": func(t time.Time) string { return t.Format("Mon") },
			"%A": func(t time.Time) string { return t.Format("Monday") },
			"%b": func(t time.Time) string { return t.Format("Feb") },
			"%B": func(t time.Time) string { return t.Format("February") },
			"%c": func(t time.Time) string { return Fmt("%a %b %e %H:%M:%S %Y", t) },
			"%C": func(t time.Time) string { return t.Format("06") },
			"%d": func(t time.Time) string { return t.Format("02") },
			"%D": func(t time.Time) string { return Fmt("%m/%d/%y", t) },
			"%e": func(t time.Time) string { return Fmt(" 2", t) },

			"%m": func(t time.Time) string { return t.Format("01") },
			"%Y": func(t time.Time) string { return t.Format("2006") },
			"%y": func(t time.Time) string { return t.Format("06") },
		},
	}
}

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
