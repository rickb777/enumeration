package transform

import (
	"fmt"
	"strings"
)

type Transform interface {
	NoOp() bool
	Transform(s string) string
	Expression(s string) string
	Imports() []string
}

type Transforms []Transform

func ListOf(ts ...Transform) (list Transforms) {
	for _, t := range ts {
		if !t.NoOp() {
			list = append(list, t)
		}
	}
	return list
}

func (ts Transforms) Transform(s string) string {
	for _, t := range ts {
		s = t.Transform(s)
	}
	return s
}

func (ts Transforms) Expression(s string) string {
	for _, t := range ts {
		s = t.Expression(s)
	}
	return s
}

func (ts Transforms) Imports() (list []string) {
	for _, t := range ts {
		list = append(list, t.Imports()...)
	}
	return list // may contain duplicates
}

//-------------------------------------------------------------------------------------------------

type Unsnake bool

func (u Unsnake) NoOp() bool {
	return !bool(u)
}

func (u Unsnake) Transform(s string) string {
	if u {
		return strings.ReplaceAll(s, "_", " ")
	}
	return s
}

func (u Unsnake) Expression(s string) string {
	if u {
		return fmt.Sprintf(`strings.ReplaceAll(%s, "_", " ")`, s)
	}
	return s
}

func (u Unsnake) Imports() []string {
	if u {
		return []string{"strings"}
	}
	return nil
}
