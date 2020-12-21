package main

import (
	. "github.com/onsi/gomega"
	"strings"
	"testing"
)

const e0 = `// generated code - do not edit
// github.com/rickb777/enumeration `

const e1 = `

package confectionary

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/rickb777/enumeration/enum"
)
`

func TestWriteHead(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeHead(buf)
	Ω(buf.String()).Should(Equal(e0+version+e1), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e2nc = `
const sweetEnumStrings = "MarsBountySnickersKitkat"

var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24}
`

const e2lc = `
const sweetEnumStrings = "marsbountysnickerskitkat"

var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24}
`

func TestWriteJoinedStringAndIndexes(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeJoinedStringAndIndexes(buf)
	Ω(buf.String()).Should(Equal(e2nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeJoinedStringAndIndexes(buf)
	Ω(buf.String()).Should(Equal(e2lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e3 = `
// AllSweets lists all 4 values in order.
var AllSweets = []Sweet{
	Mars, Bounty, Snickers, Kitkat,
}

// AllSweetEnums lists all 4 values in order.
var AllSweetEnums = enum.IntEnums{
	Mars, Bounty, Snickers, Kitkat,
}
`

func TestWriteAllItems(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeAllItems(buf)
	Ω(buf.String()).Should(Equal(e3), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e4nc = `
// Literal returns the literal string representation of a Sweet, which is
// the same as the const identifier.
func (i Sweet) Literal() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[o]:sweetEnumIndex[o+1]]
}
`

const e4lc = `
// Literal returns the literal string representation of a Sweet, which is
// the same as the const identifier.
func (i Sweet) Literal() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%g)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[o]:sweetEnumIndex[o+1]]
}
`

func TestWriteLiteralMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeLiteralMethod(buf)
	Ω(buf.String()).Should(Equal(e4nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeLiteralMethod(buf)
	Ω(buf.String()).Should(Equal(e4lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e5nc = `
// String returns the string representation of a Sweet. This uses Literal.
func (i Sweet) String() string {
	return i.Literal()
}
`

const e5lc = `
var sweetNamesInverse = map[string]Sweet{}

func init() {
	if len(sweetNames) != 4 {
		panic(fmt.Sprintf("sweetNames has %d items but should have 4", len(sweetNames)))
	}

	for k, v := range sweetNames {
		sweetNamesInverse[v] = k
	}

	if len(sweetNames) != len(sweetNamesInverse) {
		panic(fmt.Sprintf("sweetNames has %d items but they are not distinct", len(sweetNames)))
	}
}

// String returns the string representation of a Sweet.
func (i Sweet) String() string {
	s, ok := sweetNames[i]
	if ok {
		return s
	}
	return fmt.Sprintf("%02d", i)
}
`

func TestWriteStringMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeStringMethod(buf)
	Ω(buf.String()).Should(Equal(e5nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeStringMethod(buf)
	Ω(buf.String()).Should(Equal(e5lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e6 = `
// Ordinal returns the ordinal number of a Sweet.
func (i Sweet) Ordinal() int {
	switch i {
	case Mars:
		return 0
	case Bounty:
		return 1
	case Snickers:
		return 2
	case Kitkat:
		return 3
	}
	return -1
}
`

func TestWriteOrdinalMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeOrdinalMethod(buf)
	Ω(buf.String()).Should(Equal(e6), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeOrdinalMethod(buf)
	Ω(buf.String()).Should(Equal(e6), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e7nc = `
// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Sweet) Int() int {
	return int(i)
}
`

const e7lc = `
// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (i Sweet) Float() float64 {
	return float64(i)
}
`

func TestWriteBaseMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeBaseMethod(buf)
	Ω(buf.String()).Should(Equal(e7nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeBaseMethod(buf)
	Ω(buf.String()).Should(Equal(e7lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e8 = `
// SweetOf returns a Sweet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Sweet is returned.
func SweetOf(i int) Sweet {
	if 0 <= i && i < len(AllSweets) {
		return AllSweets[i]
	}
	// an invalid result
	return Mars + Bounty + Snickers + Kitkat + 1
}
`

func TestWriteOfMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeOfMethod(buf)
	Ω(buf.String()).Should(Equal(e8), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeOfMethod(buf)
	Ω(buf.String()).Should(Equal(e8), buf.String())
}

//-------------------------------------------------------------------------------------------------
const e9 = `
// IsValid determines whether a Sweet is one of the defined constants.
func (i Sweet) IsValid() bool {
	switch i {
	case Mars, Bounty, Snickers, Kitkat:
		return true
	}
	return false
}
`

func TestWriteIsValid(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeIsValidMethod(buf)
	Ω(buf.String()).Should(Equal(e9), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeIsValidMethod(buf)
	Ω(buf.String()).Should(Equal(e9), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e10nc = `
// Parse parses a string to find the corresponding Sweet, accepting one of the string
// values or an ordinal number.
func (v *Sweet) Parse(s string) error {
	if v.parseOrdinal(s) {
		return nil
	}

	if v.parseIdentifier(s) {
		return nil
	}

	return errors.New(s + ": unrecognised Sweet")
}

// parseOrdinal attempts to convert ordinal value
func (v *Sweet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return true
	}
	return false
}

// parseIdentifier attempts to match an identifier.
func (v *Sweet) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0
	for j := 1; j < len(sweetEnumIndex); j++ {
		i1 := sweetEnumIndex[j]
		p := sweetEnumStrings[i0:i1]
		if s == p {
			*v = AllSweets[j-1]
			return true
		}
		i0 = i1
	}
	return false
}
`

const e10lc = `
// Parse parses a string to find the corresponding Sweet, accepting one of the string
// values or an ordinal number.
// The case of s does not matter.
func (v *Sweet) Parse(s string) error {
	s = strings.ToLower(s)
	if v.parseOrdinal(s) {
		return nil
	}

	if sweetMarshalTextUsingLiteral {
		if v.parseIdentifier(s) || v.parseString(s) {
			return nil
		}
	} else {
		if v.parseString(s) || v.parseIdentifier(s) {
			return nil
		}
	}

	return errors.New(s + ": unrecognised Sweet")
}

// parseOrdinal attempts to convert ordinal value
func (v *Sweet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return true
	}
	return false
}

// parseString attempts to match an entry in sweetNamesInverse
func (v *Sweet) parseString(s string) (ok bool) {
	*v, ok = sweetNamesInverse[s]
	return ok
}

// parseIdentifier attempts to match an identifier.
func (v *Sweet) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0
	for j := 1; j < len(sweetEnumIndex); j++ {
		i1 := sweetEnumIndex[j]
		p := sweetEnumStrings[i0:i1]
		if s == p {
			*v = AllSweets[j-1]
			return true
		}
		i0 = i1
	}
	return false
}
`

func TestWriteParseMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeParseMethod(buf)
	Ω(buf.String()).Should(Equal(e10nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeParseMethod(buf)
	Ω(buf.String()).Should(Equal(e10lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e11nc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string
// values or an ordinal number.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const e11lc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string
// values or an ordinal number.
// The case of s does not matter.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

func TestWriteAsMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeAsMethod(buf)
	Ω(buf.String()).Should(Equal(e11nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeAsMethod(buf)
	Ω(buf.String()).Should(Equal(e11lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e12nc = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Sweet) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

const e12lc = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Sweet) MarshalText() (text []byte, err error) {
	if sweetMarshalTextUsingLiteral {
		return []byte(i.Literal()), nil
	}
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

func TestWriteMarshalText(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeMarshalText(buf)
	Ω(buf.String()).Should(Equal(e12nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeMarshalText(buf)
	Ω(buf.String()).Should(Equal(e12lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e13nc = `
// SweetMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var SweetMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// SweetMarshalJSONUsingString is true.
func (i Sweet) MarshalJSON() ([]byte, error) {
	if !SweetMarshalJSONUsingString {
		// use the ordinal
		s := strconv.Itoa(i.Ordinal())
		return []byte(s), nil
	}
	return i.quotedString(i.String())
}

func (i Sweet) quotedString(s string) ([]byte, error) {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b, nil
}
`

const e13lc = `
// SweetMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var SweetMarshalJSONUsingString = false

// sweetMarshalTextUsingLiteral controls whether generated XML or JSON uses the String()
// or the Literal() method.
var sweetMarshalTextUsingLiteral = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// SweetMarshalJSONUsingString is true.
func (i Sweet) MarshalJSON() ([]byte, error) {
	if !SweetMarshalJSONUsingString {
		// use the ordinal
		s := strconv.Itoa(i.Ordinal())
		return []byte(s), nil
	}
	if sweetMarshalTextUsingLiteral {
		return i.quotedString(i.Literal())
	}
	return i.quotedString(i.String())
}

func (i Sweet) quotedString(s string) ([]byte, error) {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b, nil
}
`

func TestWriteMarshaJSON(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeMarshalJSON(buf)
	Ω(buf.String()).Should(Equal(e13nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeMarshalJSON(buf)
	Ω(buf.String()).Should(Equal(e13lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e14 = `
// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Sweet) UnmarshalJSON(text []byte) error {
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

func TestWriteUnmarshaJSON(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeUnmarshalJSON(buf)
	Ω(buf.String()).Should(Equal(e14), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeUnmarshalJSON(buf)
	Ω(buf.String()).Should(Equal(e14), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e15 = `
// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Sweet) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		*i = Sweet(v)
	case float64:
		*i = Sweet(v)
	case []byte:
		err = i.Parse(string(v))
	case string:
		err = i.Parse(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful Sweet", value, value)
	}

	return err
}

// -- copy this somewhere and uncomment it if you need DB storage to use strings --
// Value converts the period to a string. 
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i Sweet) Value() (driver.Value, error) {
//    return i.String(), nil
//}
`

func TestWriteScanValue(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeScanValue(buf)
	Ω(buf.String()).Should(Equal(e15), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeScanValue(buf)
	Ω(buf.String()).Should(Equal(e15), buf.String())
}

//-------------------------------------------------------------------------------------------------

var modelNoChange = model{
	MainType: "Sweet",
	LcType:   "sweet",
	BaseType: "int",
	Plural:   "Sweets",
	Pkg:      "confectionary",
	Version:  version,
	Values:   []string{"Mars", "Bounty", "Snickers", "Kitkat"},
	XF:       nil,
}

var modelLowerWithLookupTable = model{
	MainType:    "Sweet",
	LcType:      "sweet",
	BaseType:    "float64",
	Plural:      "Sweets",
	Pkg:         "confectionary",
	Version:     version,
	Values:      []string{"Mars", "Bounty", "Snickers", "Kitkat"},
	XF:          []Transformer{toLower},
	LookupTable: "sweetNames",
}
