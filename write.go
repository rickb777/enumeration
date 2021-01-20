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
	"database/sql/driver"
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

const stringMethod = `
// String returns the literal string representation of a <<.MainType>>, which is
// the same as the const identifier.
func (i <<.MainType>>) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", i)
	}
	return <<.LcType>>EnumStrings[<<.LcType>>EnumIndex[o]:<<.LcType>>EnumIndex[o+1]]
}
`

func (m model) writeStringMethod(w io.Writer) {
	m.execTemplate(w, stringMethod)
}

//-------------------------------------------------------------------------------------------------

const tagMethod = `<<if .LookupTable>>
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

// Tag returns the string representation of a <<.MainType>>.
func (i <<.MainType>>) Tag() string {
	s, ok := <<.LookupTable>>[i]
	if ok {
		return s
	}
	return fmt.Sprintf("%02d", i)
}
<<- else>>
// Tag returns the string representation of a <<.MainType>>. This is an alias for String.
func (i <<.MainType>>) Tag() string {
	return i.String()
}
<<- end>>
`

func (m model) writeTagMethod(w io.Writer) {
	m.execTemplate(w, tagMethod)
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
// values or a number.
<<- range .XF>><<if ne .Info "">>
// <<.Info>>
<<- end>>
<<- end>>
func (v *<<.MainType>>) Parse(in string) error {
	return v.parse(in, <<.LcType>>MarshalTextRep)
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

	s := in
<<- range .XF>><<if ne .Str "">>
	s = <<.Str>>
<<- end>>
<<- end>>
<<- if .LookupTable>>

	if rep == enum.Identifier {
		if v.parseIdentifier(s) || v.parseTag(in) {
			return nil
		}
	} else {
		if v.parseTag(in) || v.parseIdentifier(s) {
			return nil
		}
	}
<<- else>>

	if v.parseIdentifier(s) {
		return nil
	}
<<- end>>

	return errors.New(in + ": unrecognised <<.LcType>>")
}

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
<<- if .LookupTable>>

// parseTag attempts to match an entry in <<.LookupTable>>Inverse
func (v *<<.MainType>>) parseTag(s string) (ok bool) {
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
// <<.LcType>>MarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var <<.LcType>>MarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to <<.MainType>>MarshalTextRep. 
func (i <<.MainType>>) MarshalText() (text []byte, err error) {
	return i.marshalText(<<.LcType>>MarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to <<.MainType>>MarshalTextRep. 
func (i <<.MainType>>) MarshalJSON() ([]byte, error) {
	return i.marshalText(<<.LcType>>MarshalTextRep, true)
}

func (i <<.MainType>>) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	var bs []byte
	switch rep {
	case enum.Number:
<<- if .IsFloat>>
		bs = []byte(strconv.FormatFloat(float64(i), 'g', 7, 64))
<<- else>>
		bs = []byte(strconv.FormatInt(int64(i), 10))
<<- end>>
	case enum.Ordinal:
		bs = []byte(strconv.Itoa(i.Ordinal()))
	case enum.Tag:
		if quoted {
			bs = i.quotedString(i.Tag())
		} else {
			bs = []byte(i.Tag())
		}
	default:
		if quoted {
			bs = []byte(i.quotedString(i.String()))
		} else {
			bs = []byte(i.String())
		}
	}
	return bs, nil
}

func (i <<.MainType>>) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}
`

func (m model) writeMarshalText(w io.Writer) {
	m.execTemplate(w, marshalText)
}

//-------------------------------------------------------------------------------------------------

const unmarshalText = `
// UnmarshalText converts transmitted values to ordinary values.
func (i *<<.MainType>>) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *<<.MainType>>) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}
`

func (m model) writeUnmarshalText(w io.Writer) {
	m.execTemplate(w, unmarshalText)
}

//-------------------------------------------------------------------------------------------------

const scanValue = `
// <<.LcType>>StoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var <<.LcType>>StoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *<<.MainType>>) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if <<.LcType>>StoreRep == enum.Ordinal {
			*i = <<.MainType>>Of(int(v))
		} else {
			*i = <<.MainType>>(v)
		}
	case float64:
		*i = <<.MainType>>(v)
	case []byte:
		err = i.parse(string(v), <<.LcType>>StoreRep)
	case string:
		err = i.parse(v, <<.LcType>>StoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful <<.LcType>>", value, value)
	}

	return err
}

// Value converts the <<.MainType>> to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i <<.MainType>>) Value() (driver.Value, error) {
	switch <<.LcType>>StoreRep {
	case enum.Number:
<<- if .IsFloat>>
		return float64(i), nil
<<- else>>
		return int64(i), nil
<<- end>>
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.String(), nil
	}
}
`

func (m model) writeScanValue(w io.Writer) {
	m.execTemplate(w, scanValue)
}

//-------------------------------------------------------------------------------------------------

func (m model) write(w io.Writer) {
	m.writeHead(w)
	m.writeJoinedStringAndIndexes(w)
	m.writeAllItems(w)
	m.writeStringMethod(w)
	m.writeTagMethod(w)
	m.writeOrdinalMethod(w)
	m.writeBaseMethod(w)
	m.writeOfMethod(w)
	m.writeIsValidMethod(w)
	m.writeParseMethod(w)
	m.writeAsMethod(w)
	m.writeMarshalText(w)
	m.writeUnmarshalText(w)
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
