package main

import (
	"fmt"
	"strings"
)

type Monad func(string) string

type Transformer struct {
	Fn   Monad  // this applies the monadic function
	Str  Monad  // this describes the monadic function
	Info string // this documents the function
	Next *Transformer
}

func (t *Transformer) Exists() bool {
	return t != nil
}

func (t *Transformer) Apply(s string) string {
	if t == nil {
		return s
	}

	return t.Next.Apply(t.Fn(s))
}

func (t *Transformer) Format(s string) string {
	if t == nil {
		return s
	}

	return t.Next.Format(t.Str(s))
}

func (t *Transformer) Describe() string {
	if t == nil {
		return ""
	}

	i := fmt.Sprintf("\n// %s", t.Info)
	return i + t.Next.Describe()
}

func (t *Transformer) Then(u *Transformer) *Transformer {
	t.Next = u
	return t
}

func noChange() *Transformer {
	return &Transformer{
		Fn:   noop,
		Str:  func(string) string { return "" },
		Info: "", // must be blank
	}
}

func toUpper() *Transformer {
	return &Transformer{
		Fn:   strings.ToUpper,
		Str:  describe("strings.ToUpper(%s)"),
		Info: "The case of s does not matter.",
	}
}

func toLower() *Transformer {
	return &Transformer{
		Fn:   strings.ToLower,
		Str:  describe("strings.ToLower(%s)"),
		Info: "The case of s does not matter.",
	}
}

func xUnsnake() *Transformer {
	return &Transformer{
		Fn: func(s string) string {
			return strings.ReplaceAll(s, "_", " ")
		},
		Str:  describe(`strings.ReplaceAll(%s, "_", " ")`),
		Info: "All underscores are replaced with space.",
	}
}

func noop(s string) string {
	return s
}

func describe(format string) func(string) string {
	return func(s string) string {
		return fmt.Sprintf(format, s)
	}
}
