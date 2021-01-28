package main

import (
	"fmt"
	"strings"
)

type Transformers []Transformer

func (xf Transformers) Apply(s string) string {
	for _, f := range xf {
		s = f.Fn(s)
	}
	return s
}

func (xf Transformers) FormatString() string {
	if len(xf) == 0 {
		return "%s"
	}

	s := xf[0].Str("@@@@@")
	for _, f := range xf[1:] {
		s = fmt.Sprintf(f.Str(s))
	}
	s = strings.Replace(s, "@@@@@", "%s", 1)
	return s
}

func (xf Transformers) TransformFunc() func(string) string {
	format := xf.FormatString()
	return func(s string) string {
		return fmt.Sprintf(format, s)
	}
}

func (xf Transformers) Exist() bool {
	return len(xf) > 0
}

//-------------------------------------------------------------------------------------------------

type Transformer struct {
	Fn   func(string) string // this applies the monadic function
	Str  func(string) string // this describes the monadic function
	Info string
}

var noChange = Transformer{
	Fn:   noop,
	Str:  func(string) string { return "" },
	Info: "", // must be blank
}

var toUpper = Transformer{
	Fn:   strings.ToUpper,
	Str:  describe("strings.ToUpper(%s)"),
	Info: "The case of s does not matter.",
}

var toLower = Transformer{
	Fn:   strings.ToLower,
	Str:  describe("strings.ToLower(%s)"),
	Info: "The case of s does not matter.",
}

var xUnsnake = Transformer{
	Fn: func(s string) string {
		return strings.ReplaceAll(s, "_", " ")
	},
	Str:  describe(`strings.ReplaceAll(%s, "_", " ")`),
	Info: "All underscores are replaced with space.",
}

func noop(s string) string {
	return s
}

func describe(format string) func(string) string {
	return func(s string) string {
		return fmt.Sprintf(format, s)
	}
}
