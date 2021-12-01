// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.9.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"strconv"
	"strings"
)

const baseEnumStrings = "acgt"

var baseEnumIndex = [...]uint16{0, 1, 2, 3, 4}

// AllBases lists all 4 values in order.
var AllBases = []Base{
	A, C, G, T,
}

// AllBaseEnums lists all 4 values in order.
var AllBaseEnums = enum.FloatEnums{
	A, C, G, T,
}

// String returns the literal string representation of a Base, which is
// the same as the const identifier.
func (i Base) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllBases) {
		return fmt.Sprintf("Base(%g)", i)
	}
	return baseEnumStrings[baseEnumIndex[o]:baseEnumIndex[o+1]]
}

// Tag returns the string representation of a Base. This is an alias for String.
func (i Base) Tag() string {
	return i.String()
}

// Ordinal returns the ordinal number of a Base.
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

	s := strings.ToLower(in)

	if v.parseIdentifier(s) {
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

// parseIdentifier attempts to match an identifier.
func (v *Base) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(baseEnumIndex); j++ {
		i1 := baseEnumIndex[j]
		p := baseEnumStrings[i0:i1]
		if s == p {
			*v = AllBases[j-1]
			return true
		}
		i0 = i1
	}
	return false
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
// By default, it is enum.Identifier and quoted strings are used.
var baseMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
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

func (i Base) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
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
	return i.Parse(s)
}

// baseStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var baseStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Base) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if baseStoreRep == enum.Ordinal {
			*i = BaseOf(int(v))
		} else {
			*i = Base(v)
		}
	case float64:
		*i = Base(v)
	case []byte:
		err = i.parse(string(v), baseStoreRep)
	case string:
		err = i.parse(v, baseStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful base", value, value)
	}

	return err
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
