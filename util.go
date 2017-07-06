package main

import (
	"fmt"
	"os"
	"strings"
)

var stdout = os.Stdout

func fail(args ...interface{}) {
	fmt.Fprint(os.Stderr, "Error: ")
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func info(msg string, args ...interface{}) {
	if *verbose {
		fmt.Fprintf(stdout, msg, args...)
	}
}

func debug(msg string, args ...interface{}) {
	if *dbg {
		fmt.Fprintf(stdout, msg, args...)
	}
}

func divideC(s string, c byte) (string, string) {
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s, ""
	}
	return s[:p], s[p+1:]
}

func divideS(s, sep string) (string, string) {
	p := strings.LastIndex(s, sep)
	if p < 0 {
		return s, ""
	}
	x := len(sep)
	return s[:p], s[p+x:]
}

func removeBeforeC(s string, c byte) string {
	p := strings.LastIndexByte(s, c)
	if p < 0 {
		return s
	}
	return s[p+1:]
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
