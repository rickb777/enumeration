package main

import (
	"github.com/rickb777/enumeration/v2/transform"
	"strings"
	"testing"
)

const e0 = `// generated code - do not edit
// github.com/rickb777/enumeration/v2 `

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
	buf := &strings.Builder{}
	modelNoChange.writeHead(buf)
	compare(t, buf.String(), e0+version+e1)
}

//-------------------------------------------------------------------------------------------------

const e2nc = `
const sweetEnumStrings = "MarsBountySnickersKitkatFerrero Rocher"

var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 38}
`

const e2lc = `
const sweetEnumStrings = "marsbountysnickerskitkatferrero rocher"

var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 38}
`

const e2ic = `
const (
	sweetEnumStrings = "MarsBountySnickersKitkatFerrero_Rocher"
	sweetEnumInputs  = "marsbountysnickerskitkatferrero_rocher"
)

var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 38}
`

func TestWriteJoinedStringAndIndexes(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeJoinedStringAndIndexes(buf)
	compare(t, buf.String(), e2nc)

	buf.Reset()
	modelLowerWithLookupTable.writeJoinedStringAndIndexes(buf)
	compare(t, buf.String(), e2lc)

	buf.Reset()
	modelIgnoreCase.writeJoinedStringAndIndexes(buf)
	compare(t, buf.String(), e2ic)
}

//-------------------------------------------------------------------------------------------------

const e3 = `
// AllSweets lists all 5 values in order.
var AllSweets = []Sweet{
	Mars, Bounty, Snickers, Kitkat,
	Ferrero_Rocher,
}

// AllSweetEnums lists all 5 values in order.
var AllSweetEnums = enum.IntEnums{
	Mars, Bounty, Snickers, Kitkat,
	Ferrero_Rocher,
}
`

func TestWriteAllItems(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeAllItems(buf)
	compare(t, buf.String(), e3)
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
	buf := &strings.Builder{}
	modelNoChange.writeStringMethod(buf)
	compare(t, buf.String(), e4nc)

	buf.Reset()
	modelLowerWithLookupTable.writeStringMethod(buf)
	compare(t, buf.String(), e4lc)
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
	for _, id := range AllSweets {
		v, exists := sweetNames[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Sweet: %s is missing from sweetNames\n", id)
		} else {
			k := strings.ToLower(strings.ReplaceAll(v, "_", " "))
			if _, exists := sweetNamesInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: Sweet: %q is duplicated in sweetNames\n", k)
			}
			sweetNamesInverse[k] = id
		}
	}

	if len(sweetNames) != 5 {
		panic(fmt.Sprintf("Sweet: sweetNames has %d items but should have 5", len(sweetNames)))
	}

	if len(sweetNames) != len(sweetNamesInverse) {
		panic(fmt.Sprintf("Sweet: sweetNames has %d items but there are only %d distinct items",
			len(sweetNames), len(sweetNamesInverse)))
	}
}

// Tag returns the string representation of a Sweet. For invalid values,
// this returns i.String() (see IsValid).
func (i Sweet) Tag() string {
	s, ok := sweetNames[i]
	if ok {
		return s
	}
	return i.String()
}
`

func TestWriteTagMethod(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeTagMethod(buf)
	compare(t, buf.String(), e5nc)

	buf.Reset()
	modelLowerWithLookupTable.writeTagMethod(buf)
	compare(t, buf.String(), e5lc)
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
	case Ferrero_Rocher:
		return 4
	}
	return -1
}
`

func TestWriteOrdinalMethod(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeOrdinalMethod(buf)
	compare(t, buf.String(), e6)

	buf.Reset()
	modelLowerWithLookupTable.writeOrdinalMethod(buf)
	compare(t, buf.String(), e6)
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
	buf := &strings.Builder{}
	modelNoChange.writeBaseMethod(buf)
	compare(t, buf.String(), e7nc)

	buf.Reset()
	modelLowerWithLookupTable.writeBaseMethod(buf)
	compare(t, buf.String(), e7lc)
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
	return Mars + Bounty + Snickers + Kitkat + Ferrero_Rocher + 1
}
`

func TestWriteOfMethod(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeOfMethod(buf)
	compare(t, buf.String(), e8)

	buf.Reset()
	modelLowerWithLookupTable.writeOfMethod(buf)
	compare(t, buf.String(), e8)
}

//-------------------------------------------------------------------------------------------------
const e9 = `
// IsValid determines whether a Sweet is one of the defined constants.
func (i Sweet) IsValid() bool {
	return i.Ordinal() >= 0
}
`

func TestWriteIsValid(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeIsValidMethod(buf)
	compare(t, buf.String(), e9)

	buf.Reset()
	modelLowerWithLookupTable.writeIsValidMethod(buf)
	compare(t, buf.String(), e9)
}

//-------------------------------------------------------------------------------------------------

const e10nc = `
// Parse parses a string to find the corresponding Sweet, accepting one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It is used by AsSweet.
//
// Usage Example
//
//    v := new(Sweet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Sweet) Parse(s string) error {
	return v.parse(s, sweetMarshalTextRep)
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

	s := strings.ReplaceAll(in, "_", " ")

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
// Parse parses a string to find the corresponding Sweet, accepting one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It is used by AsSweet.
//
// Usage Example
//
//    v := new(Sweet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Sweet) Parse(s string) error {
	return v.parse(s, sweetMarshalTextRep)
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

	s := strings.ToLower(strings.ReplaceAll(in, "_", " "))

	if rep == enum.Identifier {
		if v.parseIdentifier(s) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseIdentifier(s) {
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

const e10ic = `
// Parse parses a string to find the corresponding Sweet, accepting one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It is used by AsSweet.
// The input case does not matter.
//
// Usage Example
//
//    v := new(Sweet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Sweet) Parse(s string) error {
	return v.parse(s, sweetMarshalTextRep)
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

	s := strings.ToLower(in)

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
		p := sweetEnumInputs[i0:i1]
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
	buf := &strings.Builder{}
	modelNoChange.writeParseMethod(buf)
	compare(t, buf.String(), e10nc)

	buf.Reset()
	modelLowerWithLookupTable.writeParseMethod(buf)
	compare(t, buf.String(), e10lc)

	buf.Reset()
	modelIgnoreCase.writeParseMethod(buf)
	compare(t, buf.String(), e10ic)
}

//-------------------------------------------------------------------------------------------------

const e11nc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It wraps Parse.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const e11lc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It wraps Parse.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const e11ic = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

func TestWriteAsMethod(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeAsMethod(buf)
	compare(t, buf.String(), e11nc)

	buf.Reset()
	modelLowerWithLookupTable.writeAsMethod(buf)
	compare(t, buf.String(), e11lc)

	buf.Reset()
	modelIgnoreCase.writeAsMethod(buf)
	compare(t, buf.String(), e11ic)
}

//-------------------------------------------------------------------------------------------------

const e12lc = `
// MustParseSweet is similar to AsSweet except that it panics on error.
func MustParseSweet(s string) Sweet {
	i, err := AsSweet(s)
	if err != nil {
		panic(err)
	}
	return i
}
`

const e12ic = `
// MustParseSweet is similar to AsSweet except that it panics on error.
// The input case does not matter.
func MustParseSweet(s string) Sweet {
	i, err := AsSweet(s)
	if err != nil {
		panic(err)
	}
	return i
}
`

func TestWriteMustParseMethod(t *testing.T) {
	buf := &strings.Builder{}
	modelNoChange.writeMustParseMethod(buf)
	compare(t, buf.String(), e12lc)

	buf.Reset()
	modelIgnoreCase.writeMustParseMethod(buf)
	compare(t, buf.String(), e12ic)
}

//-------------------------------------------------------------------------------------------------

const e13nc = `
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

const e13lc = `
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
	buf := &strings.Builder{}
	modelNoChange.writeMarshalText(buf)
	compare(t, buf.String(), e13nc)

	buf.Reset()
	modelLowerWithLookupTable.writeMarshalText(buf)
	compare(t, buf.String(), e13lc)
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
	buf := &strings.Builder{}
	modelNoChange.writeUnmarshalText(buf)
	compare(t, buf.String(), e14)

	buf.Reset()
	modelLowerWithLookupTable.writeUnmarshalText(buf)
	compare(t, buf.String(), e14)
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
	buf := &strings.Builder{}
	modelNoChange.writeScanValue(buf)
	compare(t, buf.String(), e15nc)

	buf.Reset()
	modelLowerWithLookupTable.writeScanValue(buf)
	compare(t, buf.String(), e15lc)
}

//-------------------------------------------------------------------------------------------------

func compare(t *testing.T, actual, expected string) {
	t.Helper()
	a := strings.Split(actual, "\n")
	b := strings.Split(expected, "\n")
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	ok := true
	for i := 0; i < n; i++ {
		if ok && a[i] != b[i] {
			ap, an, bp, bn := "", "", "", ""
			if i > 0 {
				ap = a[i-1]
				bp = b[i-1]
			}
			if i < n-1 {
				an = a[i+1]
				bn = b[i+1]
			}
			t.Errorf("Line %d\n--\n  %s\na:%s\n  %s\n--\n  %s\ne:%s\n  %s\n--", i+1, ap, a[i], an, bp, b[i], bn)
			ok = false
		}
	}
	if len(a) > len(b) {
		t.Errorf("Actual has %d more lines than expected.", len(a)-len(b))
		ok = false
	} else if len(a) < len(b) {
		t.Errorf("Actual has %d fewer lines than expected.", len(b)-len(a))
		ok = false
	}
	if !ok {
		t.Logf("%s\n", actual)
	}
}

//-------------------------------------------------------------------------------------------------

var modelNoChange = model{
	MainType: "Sweet",
	LcType:   "sweet",
	BaseType: "int",
	Plural:   "Sweets",
	Pkg:      "confectionary",
	Version:  version,
	Values:   []string{"Mars", "Bounty", "Snickers", "Kitkat", "Ferrero_Rocher"},
	Unsnake:  true,
}

var modelIgnoreCase = model{
	MainType:   "Sweet",
	LcType:     "sweet",
	BaseType:   "int",
	Plural:     "Sweets",
	Pkg:        "confectionary",
	Version:    version,
	Values:     []string{"Mars", "Bounty", "Snickers", "Kitkat", "Ferrero_Rocher"},
	IgnoreCase: true,
}

var modelLowerWithLookupTable = model{
	MainType:    "Sweet",
	LcType:      "sweet",
	BaseType:    "float64",
	Plural:      "Sweets",
	Pkg:         "confectionary",
	Version:     version,
	Values:      []string{"Mars", "Bounty", "Snickers", "Kitkat", "Ferrero_Rocher"},
	IgnoreCase:  false,
	Unsnake:     true,
	Case:        transform.Lower,
	LookupTable: "sweetNames",
}
