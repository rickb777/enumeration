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
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
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
	return v.parse(in, sweetMarshalTextRep)
}

func (v *Sweet) parse(in string, rep enum.Representation) error {
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
	return v.parse(in, sweetMarshalTextRep)
}

func (v *Sweet) parse(in string, rep enum.Representation) error {
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
	s = strings.ToLower(s)

	if rep == enum.Identifier {
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
// sweetMarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var sweetMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalText() (text []byte, err error) {
	return i.marshalText(sweetMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalJSON() ([]byte, error) {
	return i.marshalText(sweetMarshalTextRep, true)
}

func (i Sweet) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	var bs []byte
	switch rep {
	case enum.Number:
		bs = []byte(strconv.FormatInt(int64(i), 10))
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

func (i Sweet) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}
`

const e12lc = `
// sweetMarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var sweetMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalText() (text []byte, err error) {
	return i.marshalText(sweetMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalJSON() ([]byte, error) {
	return i.marshalText(sweetMarshalTextRep, true)
}

func (i Sweet) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	var bs []byte
	switch rep {
	case enum.Number:
		bs = []byte(strconv.FormatFloat(float64(i), 'g', 7, 64))
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

func (i Sweet) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
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

const e14 = `
// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

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

func TestWriteUnmarshalText(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeUnmarshalText(buf)
	Ω(buf.String()).Should(Equal(e14), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeUnmarshalText(buf)
	Ω(buf.String()).Should(Equal(e14), buf.String())
}

//-------------------------------------------------------------------------------------------------

const e15nc = `
// sweetStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var sweetStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Sweet) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if sweetStoreRep == enum.Ordinal {
			*i = SweetOf(int(v))
		} else {
			*i = Sweet(v)
		}
	case float64:
		*i = Sweet(v)
	case []byte:
		err = i.parse(string(v), sweetStoreRep)
	case string:
		err = i.parse(v, sweetStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful sweet", value, value)
	}

	return err
}

// Value converts the Sweet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Sweet) Value() (driver.Value, error) {
	switch sweetStoreRep {
	case enum.Number:
		return int64(i), nil
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.String(), nil
	}
}
`

const e15lc = `
// sweetStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var sweetStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Sweet) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if sweetStoreRep == enum.Ordinal {
			*i = SweetOf(int(v))
		} else {
			*i = Sweet(v)
		}
	case float64:
		*i = Sweet(v)
	case []byte:
		err = i.parse(string(v), sweetStoreRep)
	case string:
		err = i.parse(v, sweetStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful sweet", value, value)
	}

	return err
}

// Value converts the Sweet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Sweet) Value() (driver.Value, error) {
	switch sweetStoreRep {
	case enum.Number:
		return float64(i), nil
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.String(), nil
	}
}
`

func TestWriteScanValue(t *testing.T) {
	RegisterTestingT(t)
	buf := &strings.Builder{}
	modelNoChange.writeScanValue(buf)
	Ω(buf.String()).Should(Equal(e15nc), buf.String())

	buf.Reset()
	modelLowerWithLookupTable.writeScanValue(buf)
	Ω(buf.String()).Should(Equal(e15lc), buf.String())
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
