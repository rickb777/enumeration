package model

import (
	"fmt"
	"go/types"
	"io"
	"strings"
	"text/template"

	"github.com/rickb777/enumeration/v3/enum"
	"github.com/rickb777/enumeration/v3/internal/collection"
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
<<- if .Imports.Database >>
	"database/sql/driver"
<<- end >>
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
<<- if .Imports.Strings >>
	"strings"
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

const toStringMethod = `
func (v <<.MainType>>) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}
`

func (m Model) writeToStringMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.toString", toStringMethod)
}

//-------------------------------------------------------------------------------------------------

const parseStringMethod = `
func (v *<<.MainType>>) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = All<<.Plural>>[j-1]
			return true
		}
		i0 = i1
	}
<<- if .AliasTable>>
	*v, ok = <<.AliasTable>>[s]
	return ok
<<- else>>
	return false
<<- end>>
}
`

func (m Model) writeParseStringMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.parseString", parseStringMethod)
}

//-------------------------------------------------------------------------------------------------

const stringMethod = `
// String returns the literal string representation of a <<.MainType>>, which is
// the same as the const identifier but without prefix or suffix.
func (v <<.MainType>>) String() string {
	o := v.Ordinal()
	return v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
}
`

func (m Model) writeStringMethod(w DualWriter) {
	m.execTemplate(w, stringMethod)
	m.writeToStringMethod(w)
}

//-------------------------------------------------------------------------------------------------

const ordinalMethod = `
// Ordinal returns the ordinal number of a <<.MainType>>. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v <<.MainType>>) Ordinal() int {
	switch v {
	<<- range $i, $v := .Values>>
	case <<$v.Identifier>>:
		return <<$i>>
	<<- end>>
	}
	return -1
}
`

func (m Model) writeOrdinalMethod(w DualWriter) {
	m.execTemplate(w, ordinalMethod)
}

//-------------------------------------------------------------------------------------------------

const baseMethod = `<<if .IsFloat>>
// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (v <<.MainType>>) Float() float64 {
	return float64(v)
}
<<- else>>
// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v <<.MainType>>) Int() int {
	return int(v)
}
<<- end>>
`

func (m Model) writeBaseMethod(w DualWriter) {
	m.execTemplate(w, baseMethod)
}

//-------------------------------------------------------------------------------------------------

const ofMethod = `
// <<.MainType>>Of returns a <<.MainType>> based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid <<.MainType>> is returned.
func <<.MainType>>Of(v int) <<.MainType>> {
	if 0 <= v && v < len(All<<.Plural>>) {
		return All<<.Plural>>[v]
	}
	// an invalid result
	return <<.ValuesJoined 0 " + ">> + 1
}
`

func (m Model) writeOfMethod(w DualWriter) {
	m.execTemplate(w, ofMethod)
}

//-------------------------------------------------------------------------------------------------

const isValidMethod = `
// IsValid determines whether a <<.MainType>> is one of the defined constants.
func (v <<.MainType>>) IsValid() bool {
	return v.Ordinal() >= 0
}
`

func (m Model) writeIsValidMethod(w DualWriter) {
	m.execTemplate(w, isValidMethod)
}

//-------------------------------------------------------------------------------------------------

const parse_body = `
<< if .Extra.Doc ->>
// Parse parses a string to find the corresponding <<.MainType>>, accepting one of the string values or
// a number. The input representation is determined by <<.MarshalTextRep>>. It is used by As<<.MainType>>.
<<- if .IgnoreCase>>
// The input case does not matter.
<<- end>>
//
// Usage Example
//
//    v := new(<<.MainType>>)
//    err := v.Parse(s)
//    ...  etc
//
<< end ->>
func (v *<<.MainType>>) <<.Extra.Method>>(in string) error {
	if v.<<.Extra.parseNumber>>(in) {
		return nil
	}

	s := <<.LcType>>TransformInput(in)
<<- if .Extra.Enum>>
<<- if .Asymmetric>>

	if v.parseString(s, <<.LcType>><<.Extra.Table>>Inputs, <<.LcType>><<.Extra.Table>>Index[:]) {
		return nil
	}
<<- else >>

	if v.parseString(s, <<.LcType>><<.Extra.Table>>Strings, <<.LcType>><<.Extra.Table>>Index[:]) {
		return nil
	}
<<- end >>
<<- end >>

	return v.parseFallback(in, s)
}
`

func (m Model) writeParseHelperMethod(w DualWriter, method, table string, parseNumberAsOrdinal bool) {
	if !done.Contains("v." + method) {
		done.Add("v." + method)

		m.Extra["Method"] = method
		m.Extra["Table"] = table
		m.Extra["Doc"] = method == "Parse"

		if table != "Enum" {
			m.Extra["Enum"] = true
		}

		if parseNumberAsOrdinal {
			m.Extra["parseNumber"] = "parseOrdinal"
			m.writeParseOrdinalMethod(w)
		} else {
			m.Extra["parseNumber"] = "parseNumber"
			m.writeParseNumberMethod(w)
		}
		m.execTemplate(w, parse_body)

		m.writeParseFallback(w)
		m.writeTransformInputFunction(w)
		m.writeParseStringMethod(w)
	}
}

func (m Model) writeParseMethod(w DualWriter) {
	m.writeParseHelperMethod(w, "Parse", "Enum", m.ParseNumberAsOrdinal)
}

//-------------------------------------------------------------------------------------------------

const parseFallbackMethod = `
func (v *<<.MainType>>) parseFallback(in, s string) error {
	<<- if .Asymmetric>>
	if v.parseString(s, <<.LcType>>EnumInputs, <<.LcType>>EnumIndex[:]) {
		return nil
	}
	<<- else >>
	if v.parseString(s, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]) {
		return nil
	}
	<<- end >>
	<<- if .AliasTable>>

	var ok bool
	*v, ok = <<.AliasTable>>[s]
	if ok {
		return nil
	}
	<<- end>>

	return errors.New(in + ": unrecognised <<.LcType>>")
}
`

func (m Model) writeParseFallback(w DualWriter) {
	m.writeUnexportedFunc(w, "v.parseFallback", parseFallbackMethod)
}

//-------------------------------------------------------------------------------------------------

const parseNumberMethod = `
// parseNumber attempts to convert a decimal value.
// << if .Lenient >>Any number is allowed, even if the result is invalid.<< else ->>
Only numbers that correspond to the enumeration are valid.<< end >>
func (v *<<.MainType>>) parseNumber(s string) (ok bool) {
<<- if .IsFloat>>
	num, err := strconv.ParseFloat(s, 64)
<<- else>>
	num, err := strconv.ParseInt(s, 10, 64)
<<- end>>
	if err == nil {
		*v = <<.MainType>>(num)
		return << if .Lenient >>true<< else >>v.IsValid()<< end >>
	}
	return false
}
`

func (m Model) writeParseNumberMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.parseNumber", parseNumberMethod)
}

//-------------------------------------------------------------------------------------------------

const parseOrdinalMethod = `
// parseOrdinal attempts to convert an ordinal value.
func (v *<<.MainType>>) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(All<<.Plural>>) {
		*v = All<<.Plural>>[ord]
		return true
	}
	return false
}
`

func (m Model) writeParseOrdinalMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.parseOrdinal", parseOrdinalMethod)
}

//-------------------------------------------------------------------------------------------------

const asMethod = `
// As<<.MainType>> parses a string to find the corresponding <<.MainType>>, accepting either one of the string values or
// a number. The input representation is determined by <<.LcType>>MarshalTextRep. It wraps Parse.
<<- if .IgnoreCase>>
// The input case does not matter.
<<- end>>
func As<<.MainType>>(s string) (<<.MainType>>, error) {
	var v = new(<<.MainType>>)
	err := v.Parse(s)
	return *v, err
}
`

func (m Model) writeAsMethod(w DualWriter) {
	m.execTemplate(w, asMethod)
}

//-------------------------------------------------------------------------------------------------

const mustParseMethod = `
// MustParse<<.MainType>> is similar to As<<.MainType>> except that it panics on error.
<<- if .IgnoreCase>>
// The input case does not matter.
<<- end>>
func MustParse<<.MainType>>(s string) <<.MainType>> {
	v, err := As<<.MainType>>(s)
	if err != nil {
		panic(err)
	}
	return v
}
`

func (m Model) writeMustParseMethod(w DualWriter) {
	m.execTemplate(w, mustParseMethod)
}

//-------------------------------------------------------------------------------------------------

const marshalText_Main = `
// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v <<.MainType>>) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}
`

const marshalText_struct_tags = `
// Text returns the representation used for transmission via XML, JSON etc.
func (v <<.MainType>>) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v <<.MainType>>) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, <<.LcType>>TextStrings, <<.LcType>>TextIndex[:]), nil
}
`

const marshalText_identifier = `
// Text returns the representation used for transmission via XML, JSON etc.
func (v <<.MainType>>) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v <<.MainType>>) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]), nil
}
`

const marshalText_number = `
// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The number representation is chosen according to -marshaltext.
func (v <<.MainType>>) marshalText() (string, error) {
	if !v.IsValid() {
		return v.marshalNumberStringOrError()
	}

	return <<.LcType>>MarshalNumber(v), nil
}
`

const marshalText_ordinal = `
// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The ordinal representation is chosen according to -marshaltext.
func (v <<.MainType>>) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return "", v.invalidError()
	}

	return strconv.Itoa(o), nil
}
`

func (m Model) writeMarshalText(w DualWriter) {
	if m.HasTextTags() {
		m.execTemplate(w, marshalText_Main)
		m.execTemplate(w, marshalText_struct_tags)
		m.writeMarshalNumberOrErrorMethod(w)
		m.writeToStringMethod(w)
		return
	}

	switch m.MarshalTextRep {
	case enum.Identifier:
		m.execTemplate(w, marshalText_Main)
		m.execTemplate(w, marshalText_identifier)
		m.writeMarshalNumberOrErrorMethod(w)
		m.writeToStringMethod(w)
	case enum.Number:
		m.execTemplate(w, marshalText_Main)
		m.execTemplate(w, marshalText_number)
		m.writeMarshalNumberVarFunc(w)
		m.writeMarshalNumberOrErrorMethod(w)
	case enum.Ordinal:
		m.execTemplate(w, marshalText_Main)
		m.execTemplate(w, marshalText_ordinal)
		m.writeInvalidErrorMethod(w)
	}
}

//-------------------------------------------------------------------------------------------------

const marshalJSON_struct_tags = `
// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v <<.MainType>>) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, <<.LcType>>JSONStrings, <<.LcType>>JSONIndex[:])
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, <<.LcType>>JSONStrings, <<.LcType>>JSONIndex[:])
	return enum.QuotedString(s), nil
}
`

const marshalJSON_identifier = `
// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v <<.MainType>>) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
	return enum.QuotedString(s), nil
}
`

const marshalJSON_number = `
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The number representation is chosen according to -marshaljson.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	s := <<.LcType>>MarshalNumber(v)
	return []byte(s), nil
}
`

const marshalJSON_ordinal = `
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The ordinal representation is chosen according to -marshaljson.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, v.invalidError()
	}

	s := strconv.Itoa(o)
	return []byte(s), nil
}
`

func (m Model) writeMarshalJSON(w DualWriter) {
	if m.HasJSONTags() {
		m.execTemplate(w, marshalJSON_struct_tags)
		m.writeMarshalNumberOrErrorMethod(w)

	} else {
		switch m.MarshalJSONRep {
		case enum.Identifier:
			m.execTemplate(w, marshalJSON_identifier)
			m.writeMarshalNumberOrErrorMethod(w)
		case enum.Number:
			m.execTemplate(w, marshalJSON_number)
			m.writeMarshalNumberOrErrorMethod(w)
			m.writeMarshalNumberVarFunc(w)
		case enum.Ordinal:
			m.execTemplate(w, marshalJSON_ordinal)
			m.writeInvalidErrorMethod(w)
		}
	}
}

//-------------------------------------------------------------------------------------------------

const marshalNumberVarFunc = `
// <<.LcType>>MarshalNumber handles marshaling where a number is required or where
// the value is out of range but <<.LcType>>MarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var <<.LcType>>MarshalNumber = func(v <<.MainType>>) string {
<<- if .IsFloat>>
	return strconv.FormatFloat(float64(v), 'g', 7, 64)
<<- else>>
	return strconv.FormatInt(int64(v), 10)
<<- end>>
}
`

func (m Model) writeMarshalNumberVarFunc(w DualWriter) {
	m.writeUnexportedFunc(w, "marshalNumberVarFunc", marshalNumberVarFunc)
}

//-------------------------------------------------------------------------------------------------

const marshalNumberOrError = `
func (v <<.MainType>>) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v <<.MainType>>) marshalNumberOrError() ([]byte, error) {
<<- if and .Lenient (ne .MarshalTextRep.String "Ordinal") >>
	// allow lenient marshaling
	return []byte(<<.LcType>>MarshalNumber(v)), nil
<<- else >>
	// disallow lenient marshaling
	return nil, v.invalidError()
<<- end >>
}
`

func (m Model) writeMarshalNumberOrErrorMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.marshalNumberOrError", marshalNumberOrError)
	if m.Lenient && m.MarshalTextRep != enum.Ordinal {
		m.writeMarshalNumberVarFunc(w)
	} else {
		m.writeInvalidErrorMethod(w)
	}
}

//-------------------------------------------------------------------------------------------------

const invalidError = `
func (v <<.MainType>>) invalidError() error {
<<- if .IsFloat>>
	return fmt.Errorf("%g is not a valid <<.LcType>>", v)
<<- else>>
	return fmt.Errorf("%d is not a valid <<.LcType>>", v)
<<- end>>
}
`

func (m Model) writeInvalidErrorMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.invalidError", invalidError)
}

//-------------------------------------------------------------------------------------------------

const errorIfInvalid = `
func (v <<.MainType>>) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}
`

func (m Model) writeErrorIfInvalid(w DualWriter) {
	m.writeUnexportedFunc(w, "v.errorIfInvalid", errorIfInvalid)
	m.writeInvalidErrorMethod(w)
}

//-------------------------------------------------------------------------------------------------

const unmarshalText = `
// UnmarshalText converts transmitted values to ordinary values.
func (v *<<.MainType>>) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

`

func (m Model) writeUnmarshalText(w DualWriter) {
	if m.MarshalTextRep > 0 || m.HasTextTags() {
		m.execTemplate(w, unmarshalText)
		if m.HasTextTags() {
			m.writeParseHelperMethod(w, "unmarshalText", "Text", m.MarshalTextRep == enum.Ordinal)
		} else {
			m.writeParseHelperMethod(w, "unmarshalText", "Enum", m.MarshalTextRep == enum.Ordinal)
		}
	}
}

//-------------------------------------------------------------------------------------------------

const transformFunction = `
// <<.LcType>>TransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var <<.LcType>>TransformInput = func(in string) string {
	return << transform "in" >>
}
`

func (m Model) writeTransformInputFunction(w DualWriter) {
	m.writeUnexportedFunc(w, "xTransformInput", transformFunction)
}

//-------------------------------------------------------------------------------------------------

const unmarshalJSON_plain = `
// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *<<.MainType>>) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}
`

const unmarshalJSON_struct_tags = `
func (v *<<.MainType>>) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := <<.LcType>>TransformInput(in)
<<- if .Asymmetric>>

	if v.parseString(s, <<.LcType>>JSONInputs, <<.LcType>>JSONIndex[:]) {
<<- else >>

	if v.parseString(s, <<.LcType>>JSONStrings, <<.LcType>>JSONIndex[:]) {
<<- end >>
		return nil
	}
<<- if .Asymmetric>>

	if v.parseString(s, <<.LcType>>EnumInputs, <<.LcType>>EnumIndex[:]) {
<<- else >>

	if v.parseString(s, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]) {
<<- end >>
		return nil
	}
<<- if .AliasTable>>

	var ok bool
	*v, ok = <<.AliasTable>>[s]
	if ok {
		return nil
	}
<<- end>>

	return errors.New(in + ": unrecognised <<.LcType>>")
}
`

func (m Model) writeUnmarshalJSON(w DualWriter) {
	if m.HasJSONTags() {
		m.execTemplate(w, unmarshalJSON_struct_tags)
	} else if m.MarshalJSONRep > 0 {
		m.execTemplate(w, unmarshalJSON_plain)
		m.writeParseHelperMethod(w, "unmarshalJSON", "Enum", m.MarshalJSONRep == enum.Ordinal)
		m.writeParseStringMethod(w)
	}
}

//-------------------------------------------------------------------------------------------------

const scan_all = `
// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *<<.MainType>>) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
<<- if eq .StoreRep.String "Ordinal" >>
		*v = <<.MainType>>Of(int(x))
<<- else >>
		*v = <<.MainType>>(x)
<<- end >>
		return v.errorIfInvalid()
	case float64:
<<- if eq .StoreRep.String "Ordinal" >>
		*v = <<.MainType>>Of(int(x))
<<- else >>
		*v = <<.MainType>>(x)
<<- end >>
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful <<.LcType>>", value, value)
	}

	return v.scanParse(s)
}
`

func (m Model) writeScanMethod(w DualWriter) {
	if m.StoreRep > 0 || m.HasSQLTags() {
		m.execTemplate(w, scan_all)
		if m.HasSQLTags() {
			m.writeParseHelperMethod(w, "scanParse", "SQL", m.StoreRep == enum.Ordinal)
		} else {
			m.writeParseHelperMethod(w, "scanParse", "Enum", m.StoreRep == enum.Ordinal)
		}
		m.writeErrorIfInvalid(w)
	}
}

//-------------------------------------------------------------------------------------------------

const value_identifier = `
// Value converts the <<.MainType>> to a string  (based on '-store identifier').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.String(), nil
}
`

const value_number = `
// Value converts the <<.MainType>> to a number (based on '-store number').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
<<- if .IsFloat>>
	return float64(v), nil
<<- else>>
	return int64(v), nil
<<- end>>
}
`

const value_ordinal = `
// Value converts the <<.MainType>> to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return int64(v.Ordinal()), nil
}
`

const value_struct_tags = `
// Value converts the <<.MainType>> to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, <<.LcType>>SQLStrings, <<.LcType>>SQLIndex[:]), nil
}
`

func (m Model) writeValueMethod(w DualWriter) {
	if m.HasSQLTags() {
		m.execTemplate(w, value_struct_tags)
		return
	}

	switch m.StoreRep {
	case enum.Identifier:
		m.execTemplate(w, value_identifier)
	case enum.Number:
		m.execTemplate(w, value_number)
	case enum.Ordinal:
		m.execTemplate(w, value_ordinal)
		m.writeToStringMethod(w)
	}
}

//-------------------------------------------------------------------------------------------------

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

func (m Model) execTemplate(w DualWriter, tpl string) {
	tmpl, err := template.New("t").Funcs(m.FnMap()).Delims("<<", ">>").Parse(tpl)
	checkErr(err)
	checkErr(tmpl.Execute(w, m))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
