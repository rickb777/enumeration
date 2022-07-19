// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.14.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"strconv"
	"strings"
)

// AllBases lists all 4 values in order.
var AllBases = []Base{
	A, C, G, T,
}

// AllBaseEnums lists all 4 values in order.
var AllBaseEnums = enum.FloatEnums{
	A, C, G, T,
}

const (
	baseEnumStrings = "acgt"
)

var (
	baseEnumIndex = [...]uint16{0, 1, 2, 3, 4}
)

func (i Base) toString(concats string, indexes []uint16) string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllBases) {
		return fmt.Sprintf("Base(%g)", i)
	}
	return concats[indexes[o]:indexes[o+1]]
}

func (v *Base) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllBases[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// Tag returns the string representation of a Base. This is an alias for String.
func (i Base) Tag() string {
	return i.String()
}

// String returns the literal string representation of a Base, which is
// the same as the const identifier but without prefix or suffix.
func (i Base) String() string {
	return i.toString(baseEnumStrings, baseEnumIndex[:])
}

// Ordinal returns the ordinal number of a Base. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (i Base) Ordinal() int {
	switch i {
	case A:
		return 0
	case C:
		return 1
	case G:
		return 2
	case T:
		return 3
	}
	return -1
}

// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (i Base) Float() float64 {
	return float64(i)
}

// BaseOf returns a Base based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Base is returned.
func BaseOf(i int) Base {
	if 0 <= i && i < len(AllBases) {
		return AllBases[i]
	}
	// an invalid result
	return A + C + G + T + 1
}

// IsValid determines whether a Base is one of the defined constants.
func (i Base) IsValid() bool {
	return i.Ordinal() >= 0
}

// Parse parses a string to find the corresponding Base, accepting one of the string values or
// a number. The input representation is determined by baseMarshalTextRep. It is used by AsBase.
//
// Usage Example
//
//    v := new(Base)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Base) Parse(s string) error {
	return v.parse(s, baseMarshalTextRep)
}

func (v *Base) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := baseTransformInput(in)

	if v.parseString(s, baseEnumStrings, baseEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised base")
}

// parseNumber attempts to convert a decimal value
func (v *Base) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*v = Base(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Base) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllBases) {
		*v = AllBases[ord]
		return true
	}
	return false
}

// baseTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var baseTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsBase parses a string to find the corresponding Base, accepting either one of the string values or
// a number. The input representation is determined by baseMarshalTextRep. It wraps Parse.
func AsBase(s string) (Base, error) {
	var i = new(Base)
	err := i.Parse(s)
	return *i, err
}

// MustParseBase is similar to AsBase except that it panics on error.
func MustParseBase(s string) Base {
	i, err := AsBase(s)
	if err != nil {
		panic(err)
	}
	return i
}

// baseMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
var baseMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to baseMarshalTextRep.
func (i Base) MarshalText() (text []byte, err error) {
	return i.marshalText(baseMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to baseMarshalTextRep.
func (i Base) MarshalJSON() ([]byte, error) {
	return i.marshalText(baseMarshalTextRep, true)
}

func (i Base) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	if baseMarshalTextRep != enum.Ordinal && i.Ordinal() < 0 {
		return baseMarshalNumber(i)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return baseMarshalNumber(i)
	case enum.Ordinal:
		return i.marshalOrdinal()
	case enum.Tag:
		if quoted {
			bs = enum.QuotedString(i.Tag())
		} else {
			bs = []byte(i.Tag())
		}
	default:
		if quoted {
			bs = enum.QuotedString(i.String())
		} else {
			bs = []byte(i.String())
		}
	}
	return bs, nil
}

// baseMarshalNumber handles marshaling where a number is required or where
// the value is out of range but baseMarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var baseMarshalNumber = func(i Base) (text []byte, err error) {
	bs := []byte(strconv.FormatFloat(float64(i), 'g', 7, 64))
	return bs, nil
}

func (i Base) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(i.Ordinal()))
	return bs, nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Base) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Base) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.unmarshalJSON(s)
}

func (i *Base) unmarshalJSON(s string) error {
	return i.Parse(s)
}

// baseStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var baseStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Base) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch v := value.(type) {
	case int64:
		if baseStoreRep == enum.Ordinal {
			*i = BaseOf(int(v))
		} else {
			*i = Base(v)
		}
		return nil
	case float64:
		*i = Base(v)
		return nil
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("%T %+v is not a meaningful base", value, value)
	}

	return i.parse(s, baseStoreRep)
}

// Value converts the Base to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Base) Value() (driver.Value, error) {
	switch baseStoreRep {
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
