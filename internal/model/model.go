package model

import (
	"fmt"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"go/types"
	"strings"
	"text/template"
)

type Model struct {
	MainType, LcType, BaseType string
	Plural, Pkg, Version       string
	Prefix, Suffix             string
	Values                     []string
	IgnoreCase                 bool
	Unsnake                    bool
	Case                       transform.Case
	S1, S2                     string
	LookupTable                string
}

func shortenIdentifier(id, prefix, suffix string) string {
	if prefix != "" && strings.HasPrefix(id, prefix) {
		id = id[len(prefix):]
		if strings.HasPrefix(id, "_") {
			id = id[1:]
		}
	}
	if suffix != "" && strings.HasSuffix(id, suffix) {
		id = id[:len(id)-len(suffix)]
		if strings.HasSuffix(id, "_") {
			id = id[:len(id)-1]
		}
	}
	return id
}

func (m Model) ShortenedValues() []string {
	ss := make([]string, len(m.Values))
	for i, v := range m.Values {
		ss[i] = shortenIdentifier(v, m.Prefix, m.Suffix)
	}
	return ss
}

func (m Model) Asymmetric() bool {
	return m.IgnoreCase
}

func (m Model) InputCase() transform.Case {
	c := m.Case
	if m.IgnoreCase && c == transform.Stet {
		c = transform.Lower
	}
	return c
}

func (m Model) inputTransform(s string) string {
	if m.Unsnake {
		s = strings.ReplaceAll(s, "_", " ")
	}
	return m.InputCase().Transform(s)
}

func (m Model) outputTransform(s string) string {
	if m.Unsnake {
		s = strings.ReplaceAll(s, "_", " ")
	}
	return m.Case.Transform(s)
}

func (m Model) expression(s string) string {
	if m.Unsnake {
		s = fmt.Sprintf(`strings.ReplaceAll(%s, "_", " ")`, s)
	}
	return m.InputCase().Expression(s)
}

func (m Model) FnMap() template.FuncMap {
	fns := make(template.FuncMap)
	fns["transform"] = m.expression
	return fns
}

func (m Model) IsFloat() bool {
	return m.BaseKind() == types.Float64
}

func (m Model) BaseKind() types.BasicKind {
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

func (m Model) Placeholder() string {
	switch m.BaseKind() {
	case types.Int:
		return "%d"
	case types.Float64:
		return "%g"
	}
	return "%s"
}

func (m Model) ValuesJoined(from int, separator string) string {
	return strings.Join(m.Values[from:], separator)
}
