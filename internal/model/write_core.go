package model

import (
	"fmt"
	"go/types"
	"io"
	"strings"

	"github.com/rickb777/enumeration/v4/internal/codegen"
	"github.com/rickb777/enumeration/v4/internal/collection"
)

type DualWriter interface {
	io.Writer
	io.StringWriter
}

var done = collection.NewSet[string]()

//-------------------------------------------------------------------------------------------------

const head = `// generated code - do not edit
// github.com/rickb777/enumeration/v4 <<.Version>>

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

const (
	allItems = `
// All<<.Plural>> lists all <<len .Values>> values in order.
var All<<.Plural>> = []<<.MainType>>{
	<<.ValuesWithWrapping 1>>,
}
`
)

func (m Model) AllItemsSlice() string {
	switch m.BaseKind {
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
		vs[i] = m.InTrans.Transform(v.Text)
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
		vs[i] = m.InTrans.Transform(v.JSON)
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
		vs[i] = m.InTrans.Transform(v.SQL)
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
	spaces := strings.Repeat(" ", 4-len(table))
	fmt.Fprintf(w, "\t%s%sStrings%s = \"%s\"\n", m.LcType, table, spaces, strings.Join(ov, ""))
	if m.Asymmetric() {
		fmt.Fprintf(w, "\t%s%sInputs%s  = \"%s\"\n", m.LcType, table, spaces, strings.Join(iv, ""))
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
		fmt.Fprintf(w, "\t%sSQLIndex  = [...]uint16{%s}\n", m.LcType, m.SQLIndexes())
	}

	w.WriteString(")\n")
}

//-------------------------------------------------------------------------------------------------

// BuildUnits concatenates the units required to provide the expected methods,
// omitting any units not needed for this configuration.
func (m Model) BuildUnits() *codegen.Units {
	units := codegen.New()
	if m.Polymorphic {
		buildAllTypeEnums(units)
	}
	buildStringMethod(units)
	buildOrdinalMethod(units)
	buildIsValidMethod(units)
	buildNumberMethod(units, m)
	if !m.SimpleOnly {
		buildOfMethod(units)
		buildParseHelperMethod(units, "Parse", "Enum")
		buildAsMethod(units, m)
		buildMustParseMethod(units, m)
		buildMarshalText(units, m)
		buildMarshalJSON(units, m)
		buildUnmarshalText(units, m)
		buildUnmarshalJSON(units, m)
		buildScanMethod(units, m)
		buildValueMethod(units, m)
	}
	return units
}

func WriteGo(units *codegen.Units, m Model, w DualWriter) {
	for _, u := range units.Slice() {
		m.Imports = m.Imports.Union(u.Imports)

		if u.Transforms {
			if m.SimpleOnly {
				m.Imports.AddAll(m.OutTrans.Imports()...)
			} else {
				m.Imports.AddAll(m.InTrans.Imports()...)
			}
		}
	}

	done := collection.NewSet[string]()

	m.writeHead(w)
	m.writeAllItems(w)
	if m.Config.Polymorphic {
		writeUnit(w, units, done, m, "allTypeEnums", "root")
	}
	m.writeJoinedStringAndIndexes(w)

	for _, u := range units.Slice() {
		if u.Exported() {
			writeUnit(w, units, done, m, u.Declares, "root")
		}
	}

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
