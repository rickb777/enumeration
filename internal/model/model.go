package model

import (
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"go/types"
	"reflect"
	"strings"
	"text/template"
)

var Prefix, Suffix string

// Config contains the model parameters obtained from command line options
// (either directly or computed).
type Config struct {
	MainType       string
	Plural, Pkg    string
	MarshalTextRep enum.Representation
	StoreRep       enum.Representation
	IgnoreCase     bool
	Unsnake        bool
}

type Value struct {
	Identifier string
	Shortened  string
	JSON, SQL  string
}

type Values []Value

func ValuesOf(ss ...string) Values {
	vs := make(Values, len(ss))
	for i, s := range ss {
		vs[i] = Value{
			Identifier: s,
			Shortened:  shortenIdentifier(s, Prefix, Suffix),
		}
	}
	return vs
}

func (vs Values) Append(s string, tag reflect.StructTag) Values {
	vs = append(vs, Value{
		Identifier: s,
		Shortened:  shortenIdentifier(s, Prefix, Suffix),
		JSON:       tag.Get("json"),
		SQL:        tag.Get("sql"),
	})
	return vs
}

func shortenIdentifier(id, prefix, suffix string) string {
	if prefix == "" && suffix == "" {
		return id
	}
	short := id
	if prefix != "" && strings.HasPrefix(short, prefix) {
		short = short[len(prefix):]
	}
	if suffix != "" && strings.HasSuffix(short, suffix) {
		short = short[:len(short)-len(suffix)]
	}
	if short == "" {
		panic(id + ": cannot strip prefix/suffix when the identifier matches exactly")
	}
	return short
}

func (vs Values) Identifiers() []string {
	ids := make([]string, len(vs))
	for i, v := range vs {
		ids[i] = v.Identifier
	}
	return ids
}

func (vs Values) Shortened() []string {
	short := make([]string, len(vs))
	for i, v := range vs {
		short[i] = v.Shortened
	}
	return short
}

//-------------------------------------------------------------------------------------------------

// Model holds the information available during template evaluation.
type Model struct {
	Config
	LcType, BaseType string
	Version          string
	Values           Values
	Case             transform.Case
	TagTable         string
	AliasTable       string
	Extra            map[string]string
}

func (m Model) CheckBadPrefixSuffix() error {
	if Prefix == "" && Suffix == "" {
		return nil
	}

	for _, v := range m.Values {
		s := shortenIdentifier(v.Identifier, Prefix, Suffix)
		if s == "" {
			return fmt.Errorf("%s %s: cannot strip prefix/suffix when the identifier matches exactly", m.MainType, v.Identifier)
		}
	}

	if Prefix != "" {
		any := false
		for _, v := range m.Values {
			if strings.HasPrefix(v.Identifier, Prefix) {
				any = true
				break
			}
		}
		if any {
			for _, v := range m.Values {
				if !strings.HasPrefix(v.Identifier, Prefix) {
					return fmt.Errorf("%s %s: all identifiers must have the prefix %s (or none)", m.MainType, v, Prefix)
				}
			}
		}
	}

	if Suffix != "" {
		any := false
		for _, v := range m.Values {
			if strings.HasSuffix(v.Identifier, Suffix) {
				any = true
				break
			}
		}
		if any {
			for _, v := range m.Values {
				if !strings.HasSuffix(v.Identifier, Suffix) {
					return fmt.Errorf("%s %s: all identifiers must have the suffix %s (or none)", m.MainType, v, Suffix)
				}
			}
		}
	}

	return nil
}

func (m Model) CheckBadTags() error {
	jsonCount := 0
	for _, v := range m.Values {
		if v.Identifier != "" && v.JSON != "" {
			jsonCount++
		}
	}

	sqlCount := 0
	for _, v := range m.Values {
		if v.Identifier != "" && v.SQL != "" {
			sqlCount++
		}
	}

	if 0 < jsonCount && jsonCount < len(m.Values) {
		return fmt.Errorf("%s: some identifiers don't have the `json` tag", m.MainType)
	}

	if 0 < sqlCount && sqlCount < len(m.Values) {
		return fmt.Errorf("%s: some identifiers don't have the `sql` tag", m.MainType)
	}

	return nil
}

func (m Model) HasJSONTags() bool {
	if len(m.Values) == 0 {
		return false
	}
	return m.Values[0].JSON != ""
}

func (m Model) HasSQLTags() bool {
	if len(m.Values) == 0 {
		return false
	}
	return m.Values[0].SQL != ""
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
	fns["concat"] = func(ss []string) string {
		return strings.Join(ss, "")
	}
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
	return strings.Join(m.Values[from:].Identifiers(), separator)
}
