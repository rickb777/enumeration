package util

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

var Verbose, Dbg bool
var Stdout = os.Stdout

func Version() string {
	version, dirty := "unknown", ""

	bi, _ := debug.ReadBuildInfo()
	for _, s := range bi.Settings {
		switch s.Key {
		case "vcs.revision":
			version = s.Value
		case "vcs.modified":
			if strings.EqualFold(s.Value, "true") {
				dirty = "-dirty"
			}
		}
	}
	return version + dirty
}

func Must(err error, args ...interface{}) {
	if err != nil {
		a2 := make([]interface{}, 0, len(args)+1)
		a2 = append(a2, err.Error())
		a2 = append(a2, args...)
		Fail(a2...)
	}
}

func Fail(args ...interface{}) {
	fmt.Fprint(os.Stderr, "Error: ")
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func Info(msg string, args ...interface{}) {
	if Verbose {
		fmt.Fprintf(Stdout, msg, args...)
	}
}

func Debug(msg string, args ...interface{}) {
	if Dbg {
		fmt.Fprintf(Stdout, msg, args...)
	}
}

func removeAfterS(s, sep string) string {
	p := strings.Index(s, sep)
	if p < 0 {
		return s
	}
	return s[:p]
}

func listIndexOf(words []string, target string) int {
	for i, w := range words {
		if w == target {
			return i
		}
	}
	return -1
}
