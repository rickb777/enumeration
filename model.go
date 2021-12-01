package main

import (
	"fmt"
	"github.com/rickb777/enumeration/v2/transform"
	"go/types"
	"strings"
	"text/template"
)

type model struct {
	MainType, LcType, BaseType string
	Plural, Pkg, Version       string
	Values                     []string
	IgnoreCase                 bool
	Unsnake                    bool
	Case                       transform.Case
	S1, S2                     string
	LookupTable                string
}

func (m model) ShortenedValues() []string {
	ss := make([]string, len(m.Values))
	for i, v := range m.Values {
		ss[i] = shortenIdentifier(v)
	}
	return ss
}

func (m model) Asymmetric() bool {
	return m.IgnoreCase
}

func (m model) InputCase() transform.Case {
	c := m.Case
	if m.IgnoreCase && c == transform.Stet {
		c = transform.Lower
	}
	return c
}

func (m model) inputTransform(s string) string {
	if m.Unsnake {
		s = strings.ReplaceAll(s, "_", " ")
	}
	return m.InputCase().Transform(s)
}

func (m model) outputTransform(s string) string {
	if m.Unsnake {
		s = strings.ReplaceAll(s, "_", " ")
	}
	return m.Case.Transform(s)
}

func (m model) expression(s string) string {
	if m.Unsnake {
		s = fmt.Sprintf(`strings.ReplaceAll(%s, "_", " ")`, s)
	}
	return m.InputCase().Expression(s)
}

func (m model) FnMap() template.FuncMap {
	fns := make(template.FuncMap)
	fns["transform"] = m.expression
	return fns
}

func (m model) IsFloat() bool {
	return m.BaseKind() == types.Float64
}

func (m model) BaseKind() types.BasicKind {
	var kind types.BasicKind
	switch m.BaseType {
	case "int", "uint",
		"int8", "uint8",
		"int16", "uint16",
		"int32", "uint32",
		"int64", "uint64":
		kind = types.Int
	case "float32", "float64":
		kind = types.Float64
	}
	return kind
}

func (m model) Placeholder() string {
	switch m.BaseKind() {
	case types.Int:
		return "%d"
	case types.Float64:
		return "%g"
	}
	return "%s"
}

func (m model) ValuesJoined(from int, separator string) string {
	return strings.Join(m.Values[from:], separator)
}
