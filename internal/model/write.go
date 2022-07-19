package model

import (
	"fmt"
	"go/types"
	"io"
	"strings"
	"text/template"
)

//-------------------------------------------------------------------------------------------------

const head = `// generated code - do not edit
// github.com/rickb777/enumeration/v2 <<.Version>>

package <<.Pkg>>

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
<< if .TagTable >>	"os"
<< end >>	"strconv"
	"strings"
)
`

func (m Model) writeHead(w io.Writer) {
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

func (m Model) writeAllItems(w io.Writer) {
	m.execTemplate(w, allItems)
}

//-------------------------------------------------------------------------------------------------

const joinedStringAndIndexes = `
const (
	<<.LcType>>EnumStrings = "<<concat .TransformedOutputValues>>"
<<- if .Asymmetric>>
	<<.LcType>>EnumInputs  = "<<concat .TransformedInputValues>>"
<<- end>>
<<- if .HasJSONTags>>
	<<.LcType>>JSONStrings = "<<concat .OutputJSONValues>>"
<<- if .Asymmetric>>
	<<.LcType>>JSONInputs  = "<<concat .InputJSONValues>>"
<<- end>>
<<- end>>
<<- if .HasSQLTags>>
	<<.LcType>>SQLStrings  = "<<concat .OutputSQLValues>>"
<<- if .Asymmetric>>
	<<.LcType>>SQLInputs   = "<<concat .InputSQLValues>>"
<<- end>>
<<- end>>
)

var (
	<<.LcType>>EnumIndex = [...]uint16{<<.Indexes>>}
<<- if .HasJSONTags>>
	<<.LcType>>JSONIndex = [...]uint16{<<.JSONIndexes>>}
<<- end>>
<<- if .HasSQLTags>>
	<<.LcType>>SQLIndex = [...]uint16{<<.SQLIndexes>>}
<<- end>>
)
`

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

func (m Model) writeJoinedStringAndIndexes(w io.Writer) {
	m.execTemplate(w, joinedStringAndIndexes)
}

//-------------------------------------------------------------------------------------------------

const toStringMethod = `
func (v <<.MainType>>) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}
`

func (m Model) writeToStringMethod(w io.Writer) {
	m.execTemplate(w, toStringMethod)
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

func (m Model) writeParseIdentifierMethod(w io.Writer) {
	m.execTemplate(w, parseStringMethod)
}

//-------------------------------------------------------------------------------------------------

const tagMethod = `<<if .TagTable>>
var <<.TagTable>>Inverse = map[string]<<.MainType>>{}

func init() {
	for _, id := range All<<.Plural>> {
		v, exists := <<.TagTable>>[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: <<.MainType>>: %s is missing from <<.TagTable>>\n", id)
		} else {
			k := <<.LcType>>TransformInput(v)
			if _, exists := <<.TagTable>>Inverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: <<.MainType>>: %q is duplicated in <<.TagTable>>\n", k)
			}
			<<.TagTable>>Inverse[k] = id
		}
	}

	if len(<<.TagTable>>) != <<len .Values>> {
		panic(fmt.Sprintf("<<.MainType>>: <<.TagTable>> has %d items but should have <<len .Values>>", len(<<.TagTable>>)))
	}

	if len(<<.TagTable>>) != len(<<.TagTable>>Inverse) {
		panic(fmt.Sprintf("<<.MainType>>: <<.TagTable>> has %d items but there are only %d distinct items",
			len(<<.TagTable>>), len(<<.TagTable>>Inverse)))
	}
}

// Tag returns the string representation of a <<.MainType>>. For invalid values,
// this returns v.String() (see IsValid).
func (v <<.MainType>>) Tag() string {
	s, ok := <<.TagTable>>[v]
	if ok {
		return s
	}
	return v.String()
}
<<- else if .HasJSONTags>>
// Tag returns the JSON representation of a <<.MainType>>.
func (v <<.MainType>>) Tag() string {
	return v.toString(<<.LcType>>JSONStrings, <<.LcType>>JSONIndex[:])
}
<<- else>>
// Tag returns the string representation of a <<.MainType>>. This is an alias for String.
func (v <<.MainType>>) Tag() string {
	return v.String()
}
<<- end>>
`

func (m Model) writeTagMethod(w io.Writer) {
	m.execTemplate(w, tagMethod)
}

//-------------------------------------------------------------------------------------------------

const stringMethod = `
// String returns the literal string representation of a <<.MainType>>, which is
// the same as the const identifier but without prefix or suffix.
func (v <<.MainType>>) String() string {
	return v.toString(<<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
}
`

func (m Model) writeStringMethod(w io.Writer) {
	m.execTemplate(w, stringMethod)
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

func (m Model) writeOrdinalMethod(w io.Writer) {
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

func (m Model) writeBaseMethod(w io.Writer) {
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

func (m Model) writeOfMethod(w io.Writer) {
	m.execTemplate(w, ofMethod)
}

//-------------------------------------------------------------------------------------------------

const isValidMethod = `
// IsValid determines whether a <<.MainType>> is one of the defined constants.
func (v <<.MainType>>) IsValid() bool {
	return v.Ordinal() >= 0
}
`

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

func (m Model) writeIsValidMethod(w io.Writer) {
	m.execTemplate(w, isValidMethod)
}

//-------------------------------------------------------------------------------------------------

const parseMethod = `
// Parse parses a string to find the corresponding <<.MainType>>, accepting one of the string values or
// a number. The input representation is determined by <<.LcType>>MarshalTextRep. It is used by As<<.MainType>>.
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
func (v *<<.MainType>>) Parse(s string) error {
	return v.parse(s, <<.LcType>>MarshalTextRep)
}

func (v *<<.MainType>>) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := <<.LcType>>TransformInput(in)
<<- if .TagTable>>

	if rep == enum.Identifier {
<<- if .Asymmetric>>
		if v.parseString(s, <<.LcType>>EnumInputs, <<.LcType>>EnumIndex[:]) || v.parseTag(s) {
<<- else >>
		if v.parseString(s, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]) || v.parseTag(s) {
<<- end >>
			return nil
		}
	} else {
<<- if .Asymmetric>>
		if v.parseTag(s) || v.parseString(s, <<.LcType>>EnumInputs, <<.LcType>>EnumIndex[:]) {
<<- else >>
		if v.parseTag(s) || v.parseString(s, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]) {
<<- end >>
			return nil
		}
	}
<<- else>>
<<- if .Asymmetric>>

	if v.parseString(s, <<.LcType>>EnumInputs, <<.LcType>>EnumIndex[:]) {
<<- else >>

	if v.parseString(s, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]) {
<<- end >>
		return nil
	}
<<- end>>

	return errors.New(in + ": unrecognised <<.LcType>>")
}
`

func (m Model) writeParseMethod(w io.Writer) {
	m.execTemplate(w, parseMethod)
}

//-------------------------------------------------------------------------------------------------

const parseHelperMethods = `
// parseNumber attempts to convert a decimal value
func (v *<<.MainType>>) parseNumber(s string) (ok bool) {
<<- if .IsFloat>>
	num, err := strconv.ParseFloat(s, 64)
<<- else>>
	num, err := strconv.ParseInt(s, 10, 64)
<<- end>>
	if err == nil {
		*v = <<.MainType>>(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *<<.MainType>>) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(All<<.Plural>>) {
		*v = All<<.Plural>>[ord]
		return true
	}
	return false
}
<<- if .TagTable>>

// parseTag attempts to match an entry in <<.TagTable>>Inverse
func (v *<<.MainType>>) parseTag(s string) (ok bool) {
	*v, ok = <<.TagTable>>Inverse[s]
	return ok
}
<<- end>>
`

func (m Model) writeParseHelperMethods(w io.Writer) {
	m.execTemplate(w, parseHelperMethods)
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

func (m Model) writeAsMethod(w io.Writer) {
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

func (m Model) writeMustParseMethod(w io.Writer) {
	m.execTemplate(w, mustParseMethod)
}

//-------------------------------------------------------------------------------------------------

const marshalText = `
// <<.LcType>>MarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
// By default, it is enum.<<.MarshalTextRep>>.
// The initial value is set using the -marshaltext command line parameter.
var <<.LcType>>MarshalTextRep = enum.<<.MarshalTextRep>>

// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to <<.LcType>>MarshalTextRep.
func (v <<.MainType>>) MarshalText() (text []byte, err error) {
	return v.marshalText(<<.LcType>>MarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to <<.LcType>>MarshalTextRep.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
<<- if .HasJSONTags>>
	o := v.Ordinal()
	if o < 0 {
		if <<.LcType>>MarshalTextRep == enum.Ordinal {
			return nil, fmt.Errorf("%v is out of range", v)
		}
		return <<.LcType>>MarshalNumber(v)
	}
	s := <<.LcType>>JSONStrings[<<.LcType>>JSONIndex[o]:<<.LcType>>JSONIndex[o+1]]
	return enum.QuotedString(s), nil
<<- else >>
	return v.marshalText(<<.LcType>>MarshalTextRep, true)
<<- end >>
}

func (v <<.MainType>>) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	if rep != enum.Ordinal && !v.IsValid() {
		return <<.LcType>>MarshalNumber(v)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return <<.LcType>>MarshalNumber(v)
	case enum.Ordinal:
		return v.marshalOrdinal()
	case enum.Tag:
		if quoted {
			bs = enum.QuotedString(v.Tag())
		} else {
			bs = []byte(v.Tag())
		}
	default:
		if quoted {
			bs = enum.QuotedString(v.String())
		} else {
			bs = []byte(v.String())
		}
	}
	return bs, nil
}

// <<.LcType>>MarshalNumber handles marshaling where a number is required or where
// the value is out of range but <<.LcType>>MarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var <<.LcType>>MarshalNumber = func(v <<.MainType>>) (text []byte, err error) {
<<- if .IsFloat>>
	bs := []byte(strconv.FormatFloat(float64(v), 'g', 7, 64))
<<- else>>
	bs := []byte(strconv.FormatInt(int64(v), 10))
<<- end>>
	return bs, nil
}

func (v <<.MainType>>) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(v.Ordinal()))
	return bs, nil
}
`

func (m Model) writeMarshalText(w io.Writer) {
	m.execTemplate(w, marshalText)
}

//-------------------------------------------------------------------------------------------------

const unmarshalText = `
// UnmarshalText converts transmitted values to ordinary values.
func (v *<<.MainType>>) UnmarshalText(text []byte) error {
	return v.Parse(string(text))
}

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

func (m Model) writeUnmarshalText(w io.Writer) {
	m.execTemplate(w, unmarshalText)
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

func (m Model) writeTransformInputFunction(w io.Writer) {
	m.execTemplate(w, transformFunction)
}

//-------------------------------------------------------------------------------------------------

const unmarshalJSONUsingParse = `
func (v *<<.MainType>>) unmarshalJSON(s string) error {
	return v.Parse(s)
}
`

const unmarshalJSONUsingStructTags = `
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

func (m Model) writeUnmarshalJSON(w io.Writer) {
	if m.HasJSONTags() {
		m.execTemplate(w, unmarshalJSONUsingStructTags)
	} else {
		m.execTemplate(w, unmarshalJSONUsingParse)
	}
}

//-------------------------------------------------------------------------------------------------

const scan_all = `
// <<.LcType>>StoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.<<.StoreRep>>.
// The initial value is set using the -store command line parameter.
var <<.LcType>>StoreRep = enum.<<.StoreRep>>

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *<<.MainType>>) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		if <<.LcType>>StoreRep == enum.Ordinal {
			*v = <<.MainType>>Of(int(x))
		} else {
			*v = <<.MainType>>(x)
		}
		return nil
	case float64:
		*v = <<.MainType>>(x)
		return nil
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful <<.LcType>>", value, value)
	}
<<- if .HasSQLTags>>

	if v.parseString(s, <<.LcType>>SQLStrings, <<.LcType>>SQLIndex[:]) {
		return nil
	}

	return errors.New(s + ": unrecognised <<.LcType>>")
<<- else >>

	return v.parse(s, <<.LcType>>StoreRep)
<<- end >>
}
`

func (m Model) writeScanMethod(w io.Writer) {
	m.execTemplate(w, scan_all)
}

//-------------------------------------------------------------------------------------------------

const value_all = `
// Value converts the <<.MainType>> to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
	switch <<.LcType>>StoreRep {
	case enum.Number:
<<- if .IsFloat>>
		return float64(v), nil
<<- else>>
		return int64(v), nil
<<- end>>
	case enum.Ordinal:
		return int64(v.Ordinal()), nil
	case enum.Tag:
		return v.Tag(), nil
	default:
<<- if .HasSQLTags>>
		return v.toString(<<.LcType>>SQLStrings, <<.LcType>>SQLIndex[:]), nil
<<- else >>
		return v.String(), nil
<<- end >>
	}
}
`

func (m Model) writeValueMethod(w io.Writer) {
	m.execTemplate(w, value_all)
}

//-------------------------------------------------------------------------------------------------

func (m Model) WriteGo(w io.Writer) {
	m.writeHead(w)
	m.writeAllItems(w)
	m.writeJoinedStringAndIndexes(w)
	m.writeToStringMethod(w)
	m.writeParseIdentifierMethod(w)
	m.writeTagMethod(w)
	m.writeStringMethod(w)
	m.writeOrdinalMethod(w)
	m.writeIsValidMethod(w)
	m.writeBaseMethod(w)
	m.writeOfMethod(w)
	m.writeParseMethod(w)
	m.writeParseHelperMethods(w)
	m.writeTransformInputFunction(w)
	m.writeAsMethod(w)
	m.writeMustParseMethod(w)
	m.writeMarshalText(w)
	m.writeUnmarshalText(w)
	m.writeUnmarshalJSON(w)
	m.writeScanMethod(w)
	m.writeValueMethod(w)

	if c, ok := w.(io.Closer); ok {
		checkErr(c.Close())
	}
}

//-------------------------------------------------------------------------------------------------

func (m Model) execTemplate(w io.Writer, tpl string) {
	tmpl, err := template.New("t").Funcs(m.FnMap()).Delims("<<", ">>").Parse(tpl)
	checkErr(err)
	checkErr(tmpl.Execute(w, m))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
