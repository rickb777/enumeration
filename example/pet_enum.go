// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.1

package example

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllPets lists all 5 values in order.
var AllPets = []Pet{
	MyCat, MyDog, MyMouse, MyElephant,
	MyKoala_Bear,
}

// AllPetEnums lists all 5 values in order.
var AllPetEnums = enum.IntEnums{
	MyCat, MyDog, MyMouse, MyElephant,
	MyKoala_Bear,
}

const (
	petEnumStrings = "catdogmouseelephantkoala bear"
	petTextStrings = "Felis CatusCanis LupusMus MusculusLoxodonta AfricanaPhascolarctos Cinereus"
)

var (
	petEnumIndex = [...]uint16{0, 3, 6, 11, 19, 29}
	petTextIndex = [...]uint16{0, 11, 22, 34, 52, 74}
)

// Ordinal returns the ordinal number of a Pet. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Pet) Ordinal() int {
	switch v {
	case MyCat:
		return 0
	case MyDog:
		return 1
	case MyMouse:
		return 2
	case MyElephant:
		return 3
	case MyKoala_Bear:
		return 4
	}
	return -1
}

// String returns the literal string representation of a Pet, which is
// the same as the const identifier but without prefix or suffix.
func (v Pet) String() string {
	o := v.Ordinal()
	return v.toString(o, petEnumStrings, petEnumIndex[:])
}

func (v Pet) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllPets) {
		return fmt.Sprintf("Pet(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Pet is one of the defined constants.
func (v Pet) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Pet) Int() int {
	return int(v)
}

var invalidPetValue = func() Pet {
	var v Pet
	for {
		if !slices.Contains(AllPets, v) {
			return v
		}
		v++
	} // AllPets is a finite set so loop will terminate eventually
}()

// PetOf returns a Pet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Pet is returned.
func PetOf(v int) Pet {
	if 0 <= v && v < len(AllPets) {
		return AllPets[v]
	}
	return invalidPetValue
}

// Parse parses a string to find the corresponding Pet, accepting one of the string values or
// a number. It is used by AsPet.
//
// Usage Example
//
//	v := new(Pet)
//	err := v.Parse(s)
//	...  etc
func (v *Pet) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := petTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Any number is allowed, even if the result is invalid.
func (v *Pet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Pet(num)
		return true
	}
	return false
}

func (v *Pet) parseFallback(in, s string) error {
	if v.parseString(s, petEnumStrings, petEnumIndex[:]) {
		return nil
	}

	var ok bool
	*v, ok = petAliases[s]
	if ok {
		return nil
	}

	return errors.New(in + ": unrecognised pet")
}

func (v *Pet) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllPets[j-1]
			return true
		}
		i0 = i1
	}
	*v, ok = petAliases[s]
	return ok
}

// petTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var petTransformInput = func(in string) string {
	return strings.ToLower(strings.ReplaceAll(in, "_", " "))
}

// AsPet parses a string to find the corresponding Pet, accepting either one of the string values or
// a number. It wraps Parse.
func AsPet(s string) (Pet, error) {
	var v = new(Pet)
	err := v.Parse(s)
	return *v, err
}

// MustParsePet is similar to AsPet except that it panics on error.
func MustParsePet(s string) Pet {
	v, err := AsPet(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Pet) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Pet) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Pet) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, petTextStrings, petTextIndex[:]), nil
}

func (v Pet) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Pet) marshalNumberOrError() ([]byte, error) {
	// allow lenient marshaling
	return []byte(petMarshalNumber(v)), nil
}

// petMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var petMarshalNumber = func(v Pet) string {
	return strconv.FormatInt(int64(v), 10)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Pet) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Pet) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := petTransformInput(in)

	if v.parseString(s, petTextStrings, petTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}
