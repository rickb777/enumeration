package model

import (
	"fmt"
	"github.com/rickb777/enumeration/v3/internal/collection"
	"go/types"
	"io"
	"strings"
)

type DualWriter interface {
	io.Writer
	io.StringWriter
}

var done = collection.NewStringSet()

//-------------------------------------------------------------------------------------------------

const head = `// generated code - do not edit
// github.com/rickb777/enumeration/v3 <<.Version>>

package <<.Pkg>>

import (
<<- range .Imports.Sorted >>
	"<< . >>"
<<- end >>
)
`

func (m Model) writeHead(w DualWriter) {
	m.execTemplate(w, head)
}

//-------------------------------------------------------------------------------------------------

const allItems = `
// All<<.Plural>> lists all <<len .Values>> values in order.
var All<<.Plural>> = []<<.MainType>>{
	<<.ValuesWithWrapping 1>>,
}

// All<<.MainType>>Enums lists all <<len .Values>> values in order.
var All<<.MainType>>Enums = <<.AllItemsSlice>>{
	<<.ValuesWithWrapping 1>>,
}
`

func (m Model) AllItemsSlice() string {
	switch m.BaseKind() {
	case types.Int:
		return "enum.IntEnums"
	case types.Float64:
		return "enum.FloatEnums"
	}
	panic("undefined")
}

func (m Model) writeAllItems(w DualWriter) {
	m.execTemplate(w, allItems)
}

//-------------------------------------------------------------------------------------------------

func (m Model) TransformedInputValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = m.inputTransform(v.Shortened)
	}
	return vs
}

func (m Model) TransformedOutputValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = m.outputTransform(v.Shortened)
	}
	return vs
}

func (m Model) Indexes() string {
	buf := &strings.Builder{}
	buf.WriteString("0")
	n := 0
	for _, v := range m.Values {
		n += len(v.Shortened)
		fmt.Fprintf(buf, ", %d", n)
	}
	return buf.String()
}

//-------------------------------------------------------------------------------------------------

func (m Model) InputTextValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = m.InputCase().Transform(v.Text)
	}
	return vs
}

func (m Model) OutputTextValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = v.Text
	}
	return vs
}

func (m Model) TextIndexes() string {
	buf := &strings.Builder{}
	buf.WriteString("0")
	n := 0
	for _, v := range m.Values {
		n += len(v.Text)
		fmt.Fprintf(buf, ", %d", n)
	}
	return buf.String()
}

//-------------------------------------------------------------------------------------------------

func (m Model) InputJSONValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = m.InputCase().Transform(v.JSON)
	}
	return vs
}

func (m Model) OutputJSONValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = v.JSON
	}
	return vs
}

func (m Model) JSONIndexes() string {
	buf := &strings.Builder{}
	buf.WriteString("0")
	n := 0
	for _, v := range m.Values {
		n += len(v.JSON)
		fmt.Fprintf(buf, ", %d", n)
	}
	return buf.String()
}

//-------------------------------------------------------------------------------------------------

func (m Model) InputSQLValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = m.InputCase().Transform(v.SQL)
	}
	return vs
}

func (m Model) OutputSQLValues() []string {
	vs := make([]string, len(m.Values))
	for i, v := range m.Values {
		vs[i] = v.SQL
	}
	return vs
}

func (m Model) SQLIndexes() string {
	buf := &strings.Builder{}
	buf.WriteString("0")
	n := 0
	for _, v := range m.Values {
		n += len(v.SQL)
		fmt.Fprintf(buf, ", %d", n)
	}
	return buf.String()
}

//-------------------------------------------------------------------------------------------------

func (m Model) ValuesWithWrapping(nTabs int) string {
	tabs := "\t\t"[:nTabs]
	buf := &strings.Builder{}
	nl := 5
	for i, v := range m.Values {
		if i > 0 {
			buf.WriteString(",")
		}
		nl--
		if nl == 0 {
			buf.WriteString("\n")
			buf.WriteString(tabs)
			nl = 5
		} else if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(v.Identifier)
	}
	return buf.String()
}

//-------------------------------------------------------------------------------------------------

func (m Model) writeOneJoinedString(w DualWriter, table string, ov, iv []string) {
	fmt.Fprintf(w, "\t%s%sStrings = \"%s\"\n", m.LcType, table, strings.Join(ov, ""))
	if m.Asymmetric() {
		fmt.Fprintf(w, "\t%s%sInputs = \"%s\"\n", m.LcType, table, strings.Join(iv, ""))
	}
}

func (m Model) writeJoinedStringAndIndexes(w DualWriter) {
	w.WriteString("\nconst (\n")

	m.writeOneJoinedString(w, "Enum", m.TransformedOutputValues(), m.TransformedInputValues())
	if m.HasTextTags() {
		m.writeOneJoinedString(w, "Text", m.OutputTextValues(), m.InputTextValues())
	}
	if m.HasJSONTags() {
		m.writeOneJoinedString(w, "JSON", m.OutputJSONValues(), m.OutputJSONValues())
	}
	if m.HasSQLTags() {
		m.writeOneJoinedString(w, "SQL", m.OutputSQLValues(), m.InputSQLValues())
	}

	w.WriteString(")\n\nvar (\n")

	fmt.Fprintf(w, "\t%sEnumIndex = [...]uint16{%s}\n", m.LcType, m.Indexes())
	if m.HasTextTags() {
		fmt.Fprintf(w, "\t%sTextIndex = [...]uint16{%s}\n", m.LcType, m.TextIndexes())
	}
	if m.HasJSONTags() {
		fmt.Fprintf(w, "\t%sJSONIndex = [...]uint16{%s}\n", m.LcType, m.JSONIndexes())
	}
	if m.HasSQLTags() {
		fmt.Fprintf(w, "\t%sSQLIndex = [...]uint16{%s}\n", m.LcType, m.SQLIndexes())
	}

	w.WriteString(")\n")
}

//-------------------------------------------------------------------------------------------------

//func (m Model) buildUnits() {
//	units := make(Units)
//	buildHead(units)
//	buildAllItems(units)
//	buildJoinedStringAndIndexes(units)
//	buildStringMethod(units)
//	buildOrdinalMethod(units)
//	buildIsValidMethod(units)
//	buildBaseMethod(units, m)
//	buildOfMethod(units)
//	buildParseMethod(units)
//	buildAsMethod(units)
//	buildMustParseMethod(units)
//	buildMarshalText(units, m)
//	buildMarshalJSON(units, m)
//	buildUnmarshalText(units, m)
//	buildUnmarshalJSON(units, m)
//	buildScanMethod(units, m)
//	buildValueMethod(units, m)
//}

func (m Model) WriteGo(w DualWriter) {
	m.writeHead(w)
	m.writeAllItems(w)
	m.writeJoinedStringAndIndexes(w)
	m.writeStringMethod(w)
	m.writeOrdinalMethod(w)
	m.writeIsValidMethod(w)
	m.writeBaseMethod(w)
	m.writeOfMethod(w)
	m.writeParseMethod(w)
	m.writeAsMethod(w)
	m.writeMustParseMethod(w)
	m.writeMarshalText(w)
	m.writeMarshalJSON(w)
	m.writeUnmarshalText(w)
	m.writeUnmarshalJSON(w)
	m.writeScanMethod(w)
	m.writeValueMethod(w)

	if c, ok := w.(io.Closer); ok {
		checkErr(c.Close())
	}
}

//-------------------------------------------------------------------------------------------------

func (m Model) writeUnexportedFunc(w DualWriter, method, template string) {
	if !done.Contains(method) {
		done.Add(method)
		m.execTemplate(w, template)
	}
}
