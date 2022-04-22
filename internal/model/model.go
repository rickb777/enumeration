package model

import (
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"go/types"
	"strings"
	"text/template"
)

// Config contains the model parameters obtained from command line options
// (either directly or computed).
type Config struct {
	MainType       string
	Plural, Pkg    string
	Prefix, Suffix string
	MarshalTextRep enum.Representation
	IgnoreCase     bool
	Unsnake        bool
}

// Model holds the information available during template evaluation.
type Model struct {
	Config
	LcType, BaseType string
	Version          string
	Values           []string
	DefaultValue     string
	Tags             map[string]string
	Case             transform.Case
	S1, S2           string
	TagTable         string
	AliasTable       string
}

func shortenIdentifier(id, prefix, suffix string) string {
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

func (m Model) Shortened() []string {
	short := make([]string, len(m.Values))
	for i, id := range m.Values {
		short[i] = shortenIdentifier(id, m.Prefix, m.Suffix)
	}
	return short
}

func (m Model) CheckBadPrefixSuffix() error {
	if m.Prefix == "" && m.Suffix == "" {
		return nil
	}

	for _, id := range m.Values {
		s := shortenIdentifier(id, m.Prefix, m.Suffix)
		if s == "" {
			return fmt.Errorf(id + ": cannot strip prefix/suffix when the identifier matches exactly")
		}
	}

	if m.Prefix != "" {
		any := false
		for _, id := range m.Values {
			if strings.HasPrefix(id, m.Prefix) {
				any = true
				break
			}
		}
		if any {
			for _, id := range m.Values {
				if !strings.HasPrefix(id, m.Prefix) {
					return fmt.Errorf("%s: all identifiers must have the prefix %s (or none)", id, m.Prefix)
				}
			}
		}
	}

	if m.Suffix != "" {
		any := false
		for _, id := range m.Values {
			if strings.HasSuffix(id, m.Suffix) {
				any = true
				break
			}
		}
		if any {
			for _, id := range m.Values {
				if !strings.HasSuffix(id, m.Suffix) {
					return fmt.Errorf("%s: all identifiers must have the suffix %s (or none)", id, m.Suffix)
				}
			}
		}
	}

	return nil
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

//-------------------------------------------------------------------------------------------------

type jsonEnum struct {
	Type string `json:"type,omitempty"`
	//Description string   `json:"description,omitempty"` TODO
	Default string   `json:"default,omitempty"`
	Enum    []string `json:"enum"`
}

func (m Model) toJSON() jsonEnum {
	j := jsonEnum{
		Type: "string",
		Enum: make([]string, len(m.Values)),
	}

	switch m.MarshalTextRep {
	case enum.Identifier:
		for i, id := range m.Values {
			s := shortenIdentifier(id, m.Prefix, m.Suffix)
			v := m.outputTransform(s)
			j.Enum[i] = v
		}
		if m.DefaultValue != "" {
			s := shortenIdentifier(m.DefaultValue, m.Prefix, m.Suffix)
			v := m.outputTransform(s)
			j.Default = v
		}

	case enum.Tag:
		if len(m.Tags) > 0 {
			for i, id := range m.Values {
				v := m.outputTransform(m.Tags[id])
				j.Enum[i] = v
			}
		}
		if m.DefaultValue != "" {
			v := m.outputTransform(m.Tags[m.DefaultValue])
			j.Default = v
		}
	}
	return j
}
