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
	"github.com/rickb777/enumeration/enum"
	"strconv"
	"strings"
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
// String returns the literal string representation of a Sweet, which is
// the same as the const identifier.
func (i Sweet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[o]:sweetEnumIndex[o+1]]
}
`

const e4lc = `
// String returns the literal string representation of a Sweet, which is
// the same as the const identifier.
func (i Sweet) String() string {
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
	modelNoChange.writeStringMethod(buf)
	Ω(buf.String()).Should(Equal(e4nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeStringMethod(buf)
	Ω(buf.String()).Should(Equal(e4lc), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e5nc = `
// Tag returns the string representation of a Sweet. This is an alias for String.
func (i Sweet) Tag() string {
	return i.String()
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

// Tag returns the string representation of a Sweet.
func (i Sweet) Tag() string {
	s, ok := sweetNames[i]
	if ok {
		return s
	}
	return fmt.Sprintf("%02d", i)
}
`

func TestWriteTagMethod(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeTagMethod(buf)
	Ω(buf.String()).Should(Equal(e5nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeTagMethod(buf)
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
// values or a number.
func (v *Sweet) Parse(in string) error {
	if sweetMarshalTextUsing == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := in

	if v.parseIdentifier(s) {
		return nil
	}

	return errors.New(in + ": unrecognised sweet")
}

// parseNumber attempts to convert a decimal value
func (v *Sweet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Sweet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
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
// values or a number.
// The case of s does not matter.
func (v *Sweet) Parse(in string) error {
	if sweetMarshalTextUsing == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := in
	s = strings.ToLower(s)

	if sweetMarshalTextUsing == enum.Identifier {
		if v.parseIdentifier(s) || v.parseTag(in) {
			return nil
		}
	} else {
		if v.parseTag(in) || v.parseIdentifier(s) {
			return nil
		}
	}

	return errors.New(in + ": unrecognised sweet")
}

// parseNumber attempts to convert a decimal value
func (v *Sweet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*v = Sweet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Sweet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return true
	}
	return false
}

// parseTag attempts to match an entry in sweetNamesInverse
func (v *Sweet) parseTag(s string) (ok bool) {
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
// sweetMarshalTextUsingLiteral controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var sweetMarshalTextUsing = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to SweetMarshalTextUsing. 
func (i Sweet) MarshalText() (text []byte, err error) {
	var s string
	switch sweetMarshalTextUsing {
	case enum.Number:
		s = strconv.FormatInt(int64(i), 10)
	case enum.Ordinal:
		s = strconv.Itoa(i.Ordinal())
	case enum.Tag:
		s = i.Tag()
	default:
		s = i.String()
	}
	return []byte(s), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

const e12lc = `
// sweetMarshalTextUsingLiteral controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var sweetMarshalTextUsing = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to SweetMarshalTextUsing. 
func (i Sweet) MarshalText() (text []byte, err error) {
	var s string
	switch sweetMarshalTextUsing {
	case enum.Number:
		s = strconv.FormatFloat(float64(i), 'g', 7, 64)
	case enum.Ordinal:
		s = strconv.Itoa(i.Ordinal())
	case enum.Tag:
		s = i.Tag()
	default:
		s = i.String()
	}
	return []byte(s), nil
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
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to SweetMarshalTextUsing. 
func (i Sweet) MarshalJSON() ([]byte, error) {
	var s []byte
	switch sweetMarshalTextUsing {
	case enum.Number:
		s = []byte(strconv.FormatInt(int64(i), 10))
	case enum.Ordinal:
		s = []byte(strconv.Itoa(i.Ordinal()))
	case enum.Tag:
		s = i.quotedString(i.Tag())
	default:
		s = i.quotedString(i.String())
	}
	return s, nil
}

func (i Sweet) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}
`

const e13lc = `
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to SweetMarshalTextUsing. 
func (i Sweet) MarshalJSON() ([]byte, error) {
	var s []byte
	switch sweetMarshalTextUsing {
	case enum.Number:
		s = []byte(strconv.FormatFloat(float64(i), 'g', 7, 64))
	case enum.Ordinal:
		s = []byte(strconv.Itoa(i.Ordinal()))
	case enum.Tag:
		s = i.quotedString(i.Tag())
	default:
		s = i.quotedString(i.String())
	}
	return s, nil
}

func (i Sweet) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
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
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
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
// Value converts the Sweet to a string.
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
