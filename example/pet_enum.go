// generated code - do not edit
// github.com/rickb777/enumeration v1.8.0

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

// String returns the string representation of a Pet.
func (i Pet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllPets) {
		return fmt.Sprintf("Pet(%d)", i)
	}
	return petEnumStrings[petEnumIndex[o]:petEnumIndex[o+1]]
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
	return Cat + Dog + Mouse + Elephant + Koala_Bear
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

// Parse parses a string to find the corresponding Pet, accepting either one of the string
// values or an ordinal number.
// The case of s does not matter.
func (v *Pet) Parse(s string) error {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "_", " ")
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllPets) {
		*v = AllPets[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(petEnumIndex); j++ {
		i1 := petEnumIndex[j]
		p := petEnumStrings[i0:i1]
		if s == p {
			*v = AllPets[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Pet")
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
	if PetMarshalJSONUsingString {
		s := []byte(i.String())
		b := make([]byte, len(s)+2)
		b[0] = '"'
		copy(b[1:], s)
		b[len(s)+1] = '"'
		return b, nil
	}
	// else use the ordinal
	s := strconv.Itoa(i.Ordinal())
	return []byte(s), nil
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
