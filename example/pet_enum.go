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

const petEnumStrings = "catdogmouseelephantkoala bear"

var petEnumIndex = [...]uint16{0, 3, 6, 11, 19, 29}

// AllPets lists all 5 values in order.
var AllPets = []Pet{
	Cat, Dog, Mouse, Elephant,
	Koala_Bear,
}

// AllPetEnums lists all 5 values in order.
var AllPetEnums = enum.IntEnums{
	Cat, Dog, Mouse, Elephant,
	Koala_Bear,
}

// Literal returns the literal string representation of a Pet, which is
// the same as the const identifier.
func (i Pet) Literal() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllPets) {
		return fmt.Sprintf("Pet(%d)", i)
	}
	return petEnumStrings[petEnumIndex[o]:petEnumIndex[o+1]]
}

// String returns the string representation of a Pet. This uses Literal.
func (i Pet) String() string {
	return i.Literal()
}

// Ordinal returns the ordinal number of a Pet.
func (i Pet) Ordinal() int {
	switch i {
	case Cat:
		return 0
	case Dog:
		return 1
	case Mouse:
		return 2
	case Elephant:
		return 3
	case Koala_Bear:
		return 4
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Pet) Int() int {
	return int(i)
}

// PetOf returns a Pet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Pet is returned.
func PetOf(i int) Pet {
	if 0 <= i && i < len(AllPets) {
		return AllPets[i]
	}
	// an invalid result
	return Cat + Dog + Mouse + Elephant + Koala_Bear + 1
}

// IsValid determines whether a Pet is one of the defined constants.
func (i Pet) IsValid() bool {
	switch i {
	case Cat, Dog, Mouse, Elephant,
		Koala_Bear:
		return true
	}
	return false
}

// Parse parses a string to find the corresponding Pet, accepting one of the string
// values or an ordinal number.
// The case of s does not matter.
func (v *Pet) Parse(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := in
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "_", " ")

	if v.parseIdentifier(s) {
		return nil
	}

	return errors.New(in + ": unrecognised Pet")
}

// parseOrdinal attempts to convert ordinal value
func (v *Pet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllPets) {
		*v = AllPets[ord]
		return true
	}
	return false
}

// parseIdentifier attempts to match an identifier.
func (v *Pet) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0
	for j := 1; j < len(petEnumIndex); j++ {
		i1 := petEnumIndex[j]
		p := petEnumStrings[i0:i1]
		if s == p {
			*v = AllPets[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsPet parses a string to find the corresponding Pet, accepting either one of the string
// values or an ordinal number.
// The case of s does not matter.
func AsPet(s string) (Pet, error) {
	var i = new(Pet)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Pet) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Pet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// PetMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var PetMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// PetMarshalJSONUsingString is true.
func (i Pet) MarshalJSON() ([]byte, error) {
	if !PetMarshalJSONUsingString {
		// use the ordinal
		s := strconv.Itoa(i.Ordinal())
		return []byte(s), nil
	}
	return i.quotedString(i.String())
}

func (i Pet) quotedString(s string) ([]byte, error) {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b, nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Pet) UnmarshalJSON(text []byte) error {
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
func (i *Pet) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		*i = Pet(v)
	case float64:
		*i = Pet(v)
	case []byte:
		err = i.Parse(string(v))
	case string:
		err = i.Parse(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful Pet", value, value)
	}

	return err
}

// -- copy this somewhere and uncomment it if you need DB storage to use strings --
// Value converts the Pet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i Pet) Value() (driver.Value, error) {
//    return i.String(), nil
//}
