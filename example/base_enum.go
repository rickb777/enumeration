// generated code - do not edit
// github.com/rickb777/enumeration v1.10.0

package example

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/enum"
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

// Literal returns the literal string representation of a Base, which is
// the same as the const identifier.
func (i Base) Literal() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllBases) {
		return fmt.Sprintf("Base(%g)", i)
	}
	return baseEnumStrings[baseEnumIndex[o]:baseEnumIndex[o+1]]
}

// String returns the string representation of a Base. This uses Literal.
func (i Base) String() string {
	return i.Literal()
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
	switch i {
	case A, C, G, T:
		return true
	}
	return false
}

// Parse parses a string to find the corresponding Base, accepting one of the string
// values or an ordinal number.
// The case of s does not matter.
func (v *Base) Parse(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := in
	s = strings.ToLower(s)

	if v.parseIdentifier(s) {
		return nil
	}

	return errors.New(in + ": unrecognised Base")
}

// parseOrdinal attempts to convert ordinal value
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

// AsBase parses a string to find the corresponding Base, accepting either one of the string
// values or an ordinal number.
// The case of s does not matter.
func AsBase(s string) (Base, error) {
	var i = new(Base)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Base) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Base) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// BaseMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var BaseMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// BaseMarshalJSONUsingString is true.
func (i Base) MarshalJSON() ([]byte, error) {
	if !BaseMarshalJSONUsingString {
		// use the ordinal
		s := strconv.Itoa(i.Ordinal())
		return []byte(s), nil
	}
	return i.quotedString(i.String())
}

func (i Base) quotedString(s string) ([]byte, error) {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b, nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Base) UnmarshalJSON(text []byte) error {
	if len(text) >= 2 && text[0] == '"' && text[len(text)-1] == '"' {
		s := string(text[1 : len(text)-1])
		return i.Parse(s)
	}

	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Base) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		*i = Base(v)
	case float64:
		*i = Base(v)
	case []byte:
		err = i.Parse(string(v))
	case string:
		err = i.Parse(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful Base", value, value)
	}

	return err
}

// -- copy this somewhere and uncomment it if you need DB storage to use strings --
// Value converts the Base to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i Base) Value() (driver.Value, error) {
//    return i.String(), nil
//}
