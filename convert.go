package main

import (
	"bufio"
	"fmt"
	"github.com/rickb777/enumeration/v2/transform"
	"go/types"
	"io"
	"strings"
	"text/template"
)

func removeComments(line string) string {
	return removeAfterS(line, "//")
}

func removeMatches(words []string, unwanted string) []string {
	cp := make([]string, 0, len(words))
	for _, w := range words {
		if w != unwanted {
			cp = append(cp, w)
		}
	}
	return cp
}

func removeBlanks(words []string) []string {
	return removeMatches(words, "")
}

func removePlaceholders(words []string) []string {
	return removeMatches(words, "_")
}

func removeCommentsAndSplitWords(line string) []string {
	content := strings.TrimSpace(removeComments(line))
	return removeBlanks(strings.Split(content, " "))
}

func scanValues(s *bufio.Scanner, mainType string) (result []string) {
	debug("scanValues\n")
	found := false
	for s.Scan() {
		words := removeCommentsAndSplitWords(s.Text())
		debug("%#v\n", words)

		if len(words) == 1 && words[0] == ")" {
			if found {
				return
			}
		}

		eq := listIndexOf(words, "=")
		if eq >= 2 && len(words) >= 3 && words[eq-1] == mainType {
			found = true
			for i := 0; i < eq-1; i++ {
				names := removePlaceholders(removeBlanks(strings.Split(words[i], ",")))
				debug("started with %s\n", names)
				result = append(result, names...)
			}
		} else if found && eq < 0 && len(words) >= 1 {
			if words[0] != "_" {
				debug("added %s\n", words[0])
				result = append(result, words[0])
			}
		}
	}

	return
}

func convert(w io.Writer, in io.Reader, input, mainType, plural, pkg string, xCase transform.Case, ignoreCase, unsnake bool) error {
	foundMainType := false
	baseType := "int"
	s := bufio.NewScanner(in)

	for s.Scan() {
		words := removeCommentsAndSplitWords(s.Text())
		debug("%#v\n", words)

		if len(words) == 3 && words[0] == "type" && words[1] == mainType {
			foundMainType = true
			baseType = words[2]
			debug("type %s %s\n", mainType, baseType)

		} else if foundMainType && len(words) == 2 && words[0] == "const" && words[1] == "(" {
			values := scanValues(s, mainType)
			if values != nil {
				m := model{
					MainType:    mainType,
					LcType:      strings.ToLower(mainType),
					BaseType:    baseType,
					Plural:      plural,
					Pkg:         pkg,
					Version:     version,
					Values:      values,
					IgnoreCase:  ignoreCase,
					Unsnake:     unsnake,
					Case:        xCase,
					LookupTable: *usingTable,
				}
				m.write(w)
				return nil
			}
		}
	}

	return fmt.Errorf("Failed to find %s in %s", mainType, input)
}

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

func (m model) InputTransform(s string) string {
	if m.Unsnake {
		s = strings.ReplaceAll(s, "_", " ")
	}
	return m.InputCase().Transform(s)
}

func (m model) OutputTransform(s string) string {
	if m.Unsnake {
		s = strings.ReplaceAll(s, "_", " ")
	}
	return m.Case.Transform(s)
}

func (m model) Expression(s string) string {
	if m.Unsnake {
		s = fmt.Sprintf(`strings.ReplaceAll(%s, "_", " ")`, s)
	}
	return m.InputCase().Expression(s)
}

func (m model) FnMap() template.FuncMap {
	fns := make(template.FuncMap)
	fns["transform"] = m.Expression
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
