package main

import (
	"fmt"
	"go/types"
	"io"
	"strings"
	"text/template"
)

//-------------------------------------------------------------------------------------------------

const head = `// generated code - do not edit
// github.com/rickb777/enumeration <<.Version>>

package <<.Pkg>>

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/enum"
	"strconv"
	"strings"
)
`

func (m model) writeHead(w io.Writer) {
	m.execTemplate(w, head)
}

//-------------------------------------------------------------------------------------------------

const joinedStringAndIndexes = `
const <<.LcType>>EnumStrings = "<<.TransformedValues>>"

var <<.LcType>>EnumIndex = [...]uint16{<<.Indexes>>}
`

func (m model) TransformedValues() string {
	buf := &strings.Builder{}
	for _, s := range m.Values {
		for _, f := range m.XF {
			s = f.Fn(s)
		}
		fmt.Fprintf(buf, s)
	}
	return buf.String()
}

func (m model) Indexes() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "0")
	n := 0
	for _, s := range m.Values {
		n += len(s)
		fmt.Fprintf(buf, ", %d", n)
	}
	return buf.String()
}

func (m model) writeJoinedStringAndIndexes(w io.Writer) {
	m.execTemplate(w, joinedStringAndIndexes)
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

func (m model) AllItemsSlice() string {
	switch m.BaseKind() {
	case types.Int:
		return "enum.IntEnums"
	case types.Float64:
		return "enum.FloatEnums"
	}
	panic("undefined")
}

func (m model) writeAllItems(w io.Writer) {
	m.execTemplate(w, allItems)
}

//-------------------------------------------------------------------------------------------------

const literalMethod = `
// Literal returns the literal string representation of a <<.MainType>>, which is
// the same as the const identifier.
func (i <<.MainType>>) Literal() string {
	o := i.Ordinal()
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", i)
	}
	return <<.LcType>>EnumStrings[<<.LcType>>EnumIndex[o]:<<.LcType>>EnumIndex[o+1]]
}
`

func (m model) writeLiteralMethod(w io.Writer) {
	m.execTemplate(w, literalMethod)
}

//-------------------------------------------------------------------------------------------------

const stringMethod = `<<if .LookupTable>>
var <<.LookupTable>>Inverse = map[string]<<.MainType>>{}

func init() {
	if len(<<.LookupTable>>) != <<len .Values>> {
		panic(fmt.Sprintf("<<.LookupTable>> has %d items but should have <<len .Values>>", len(<<.LookupTable>>)))
	}

	for k, v := range <<.LookupTable>> {
		<<.LookupTable>>Inverse[v] = k
	}

	if len(<<.LookupTable>>) != len(<<.LookupTable>>Inverse) {
		panic(fmt.Sprintf("<<.LookupTable>> has %d items but they are not distinct", len(<<.LookupTable>>)))
	}
}

// String returns the string representation of a <<.MainType>>.
func (i <<.MainType>>) String() string {
	s, ok := <<.LookupTable>>[i]
	if ok {
		return s
	}
	return fmt.Sprintf("%02d", i)
}
<<- else>>
// String returns the string representation of a <<.MainType>>. This uses Literal.
func (i <<.MainType>>) String() string {
	return i.Literal()
}
<<- end>>
`

func (m model) writeStringMethod(w io.Writer) {
	m.execTemplate(w, stringMethod)
}

//-------------------------------------------------------------------------------------------------

const ordinalMethod = `
// Ordinal returns the ordinal number of a <<.MainType>>.
func (i <<.MainType>>) Ordinal() int {
	switch i {
	<<- range $i, $v := .Values>>
	case <<$v>>:
		return <<$i>>
	<<- end>>
	}
	return -1
}
`

func (m model) writeOrdinalMethod(w io.Writer) {
	m.execTemplate(w, ordinalMethod)
}

//-------------------------------------------------------------------------------------------------

const baseMethod = `<<if .IsFloat>>
// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (i <<.MainType>>) Float() float64 {
	return float64(i)
}
<<- else>>
// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i <<.MainType>>) Int() int {
	return int(i)
}
<<- end>>
`

func (m model) writeBaseMethod(w io.Writer) {
	m.execTemplate(w, baseMethod)
}

//-------------------------------------------------------------------------------------------------

const ofMethod = `
// <<.MainType>>Of returns a <<.MainType>> based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid <<.MainType>> is returned.
func <<.MainType>>Of(i int) <<.MainType>> {
	if 0 <= i && i < len(All<<.Plural>>) {
		return All<<.Plural>>[i]
	}
	// an invalid result
	return <<.ValuesJoined 0 " + ">> + 1
}
`

func (m model) writeOfMethod(w io.Writer) {
	m.execTemplate(w, ofMethod)
}

//-------------------------------------------------------------------------------------------------

const isValidMethod = `
// IsValid determines whether a <<.MainType>> is one of the defined constants.
func (i <<.MainType>>) IsValid() bool {
	switch i {
	case <<.ValuesWithWrapping 2>>:
		return true
	}
	return false
}
`

func (m model) ValuesWithWrapping(nTabs int) string {
	tabs := "\t\t"[:nTabs]
	buf := &strings.Builder{}
	nl := 5
	for i, s := range m.Values {
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
		buf.WriteString(s)
	}
	return buf.String()
}

func (m model) writeIsValidMethod(w io.Writer) {
	m.execTemplate(w, isValidMethod)
}

//-------------------------------------------------------------------------------------------------

const parseMethod = `
// Parse parses a string to find the corresponding <<.MainType>>, accepting one of the string
// values or an ordinal number.
<<- range .XF>><<if ne .Info "">>
// <<.Info>>
<<- end>>
<<- end>>
func (v *<<.MainType>>) Parse(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := in
<<- range .XF>><<if ne .Str "">>
	s = <<.Str>>
<<- end>>
<<- end>>
<<- if .LookupTable>>

	if <<.LcType>>MarshalTextUsingLiteral {
		if v.parseIdentifier(s) || v.parseString(in) {
			return nil
		}
	} else {
		if v.parseString(in) || v.parseIdentifier(s) {
			return nil
		}
	}
<<- else>>

	if v.parseIdentifier(s) {
		return nil
	}
<<- end>>

	return errors.New(in + ": unrecognised <<.MainType>>")
}

// parseOrdinal attempts to convert ordinal value
func (v *<<.MainType>>) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(All<<.Plural>>) {
		*v = All<<.Plural>>[ord]
		return true
	}
	return false
}
<<- if .LookupTable>>

// parseString attempts to match an entry in <<.LookupTable>>Inverse
func (v *<<.MainType>>) parseString(s string) (ok bool) {
	*v, ok = <<.LookupTable>>Inverse[s]
	return ok
}
<<- end>>

// parseIdentifier attempts to match an identifier.
func (v *<<.MainType>>) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0
	for j := 1; j < len(<<.LcType>>EnumIndex); j++ {
		i1 := <<.LcType>>EnumIndex[j]
		p := <<.LcType>>EnumStrings[i0:i1]
		if s == p {
			*v = All<<.Plural>>[j-1]
			return true
		}
		i0 = i1
	}
	return false
}
`

func (m model) writeParseMethod(w io.Writer) {
	m.execTemplate(w, parseMethod)
}

//-------------------------------------------------------------------------------------------------

const asMethod = `
// As<<.MainType>> parses a string to find the corresponding <<.MainType>>, accepting either one of the string
// values or an ordinal number.
<<- range .XF>><<if ne .Info "">>
// <<.Info>>
<<- end>>
<<- end>>
func As<<.MainType>>(s string) (<<.MainType>>, error) {
	var i = new(<<.MainType>>)
	err := i.Parse(s)
	return *i, err
}
`

func (m model) writeAsMethod(w io.Writer) {
	m.execTemplate(w, asMethod)
}

//-------------------------------------------------------------------------------------------------

const marshalText = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i <<.MainType>>) MarshalText() (text []byte, err error) {
<<- if .LookupTable>>
	if <<.LcType>>MarshalTextUsingLiteral {
		return []byte(i.Literal()), nil
	}
<<- end>>
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *<<.MainType>>) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

func (m model) writeMarshalText(w io.Writer) {
	m.execTemplate(w, marshalText)
}

//-------------------------------------------------------------------------------------------------

const marshalJSON = `
// <<.MainType>>MarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var <<.MainType>>MarshalJSONUsingString = false
<<- if .LookupTable>>

// <<.LcType>>MarshalTextUsingLiteral controls whether generated XML or JSON uses the String()
// or the Literal() method.
var <<.LcType>>MarshalTextUsingLiteral = false
<<- end>>

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// <<.MainType>>MarshalJSONUsingString is true.
func (i <<.MainType>>) MarshalJSON() ([]byte, error) {
	if !<<.MainType>>MarshalJSONUsingString {
		// use the ordinal
		s := strconv.Itoa(i.Ordinal())
		return []byte(s), nil
	}
<<- if .LookupTable>>
	if <<.LcType>>MarshalTextUsingLiteral {
		return i.quotedString(i.Literal())
	}
<<- end>>
	return i.quotedString(i.String())
}

func (i <<.MainType>>) quotedString(s string) ([]byte, error) {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b, nil
}
`

func (m model) writeMarshalJSON(w io.Writer) {
	m.execTemplate(w, marshalJSON)
}

//-------------------------------------------------------------------------------------------------

const unmarshalJSON = `
// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *<<.MainType>>) UnmarshalJSON(text []byte) error {
	if len(text) >= 2 && text[0] == '"' && text[len(text)-1] == '"' {
		s := string(text[1:len(text)-1])
		return i.Parse(s)
	}

	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
`

func (m model) writeUnmarshalJSON(w io.Writer) {
	m.execTemplate(w, unmarshalJSON)
}

//-------------------------------------------------------------------------------------------------

const scanValue = `
// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *<<.MainType>>) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		*i = <<.MainType>>(v)
	case float64:
		*i = <<.MainType>>(v)
	case []byte:
		err = i.Parse(string(v))
	case string:
		err = i.Parse(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful <<.MainType>>", value, value)
	}

	return err
}

// -- copy this somewhere and uncomment it if you need DB storage to use strings --
// Value converts the period to a string. 
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i <<.MainType>>) Value() (driver.Value, error) {
//    return i.String(), nil
//}
`

func (m model) writeScanValue(w io.Writer) {
	m.execTemplate(w, scanValue)
}

//-------------------------------------------------------------------------------------------------

func (m model) write(w io.Writer) {
	m.writeHead(w)
	m.writeJoinedStringAndIndexes(w)
	m.writeAllItems(w)
	m.writeLiteralMethod(w)
	m.writeStringMethod(w)
	m.writeOrdinalMethod(w)
	m.writeBaseMethod(w)
	m.writeOfMethod(w)
	m.writeIsValidMethod(w)
	m.writeParseMethod(w)
	m.writeAsMethod(w)
	m.writeMarshalText(w)
	m.writeMarshalJSON(w)
	m.writeUnmarshalJSON(w)
	m.writeScanValue(w)

	if c, ok := w.(io.Closer); ok {
		checkErr(c.Close())
	}
}

func (m model) execTemplate(w io.Writer, tpl string) {
	tmpl, err := template.New("t").Delims("<<", ">>").Parse(tpl)
	checkErr(err)
	checkErr(tmpl.Execute(w, m))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
