package main

import "strings"

type Transform int

const (
	NoChange Transform = iota
	ToUpper
	ToLower
	Unsnake
)

func (t Transform) Func() func(string) string {
	switch t {
	case NoChange:
		return noop
	case ToUpper:
		return strings.ToUpper
	case ToLower:
		return strings.ToLower
	case Unsnake:
		return stringUnsnake
	}
	panic(t)
}

func (t Transform) String() string {
	switch t {
	case NoChange:
		return ""
	case ToUpper:
		return "strings.ToUpper(s)"
	case ToLower:
		return "strings.ToLower(s)"
	case Unsnake:
		return `strings.ReplaceAll(s, "_", " ")`
	}
	panic(t)
}

func (t Transform) Info() string {
	switch t {
	case NoChange:
		return ""
	case ToUpper:
		return "The case of s does not matter."
	case ToLower:
		return "The case of s does not matter."
	case Unsnake:
		return ""
	}
	panic(t)
}

func noop(s string) string {
	return s
}

func stringUnsnake(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}
