package main

import "strings"

type Transform int

const (
	NoChange Transform = iota
	ToUpper
	ToLower
)

func (t Transform) Func() func(string) string {
	switch t {
	case NoChange:
		return noop
	case ToUpper:
		return strings.ToUpper
	case ToLower:
		return strings.ToLower
	}
	panic(t)
}

func (t Transform) String() string {
	switch t {
	case NoChange:
		return ""
	case ToUpper:
		return "strings.ToUpper"
	case ToLower:
		return "strings.ToLower"
	}
	panic(t)
}

func noop(s string) string {
	return s
}
