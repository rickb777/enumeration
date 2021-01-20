// generated code - do not edit
// github.com/rickb777/enumeration v2.0.0

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

// String returns the literal string representation of a Pet, which is
// the same as the const identifier.
func (i Pet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllPets) {
		return fmt.Sprintf("Pet(%d)", i)
	}
	return petEnumStrings[petEnumIndex[o]:petEnumIndex[o+1]]
}

// Tag returns the string representation of a Pet. This is an alias for String.
func (i Pet) Tag() string {
	return i.String()
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
// values or a number.
// The case of s does not matter.
func (v *Pet) Parse(in string) error {
	if petMarshalTextUsing == enum.Ordinal {
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
	s = strings.ReplaceAll(s, "_", " ")

	if v.parseIdentifier(s) {
		return nil
	}

	return errors.New(in + ": unrecognised pet")
}

// parseNumber attempts to convert a decimal value
func (v *Pet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Pet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
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

// petMarshalTextUsingLiteral controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var petMarshalTextUsing = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to PetMarshalTextUsing.
func (i Pet) MarshalText() (text []byte, err error) {
	var s string
	switch petMarshalTextUsing {
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
func (i *Pet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to PetMarshalTextUsing.
func (i Pet) MarshalJSON() ([]byte, error) {
	var s []byte
	switch petMarshalTextUsing {
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

func (i Pet) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Pet) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
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
