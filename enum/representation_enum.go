// generated code - do not edit

package enum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	representationEnumStrings = "IdentifierTagNumberOrdinal"
	representationEnumInputs  = "identifiertagnumberordinal"
)

var representationEnumIndex = [...]uint16{0, 10, 13, 19, 26}

// AllRepresentations lists all 4 values in order.
var AllRepresentations = []Representation{
	Identifier, Tag, Number, Ordinal,
}

// String returns the literal string representation of a Representation, which is
// the same as the const identifier.
func (i Representation) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllRepresentations) {
		return fmt.Sprintf("Representation(%d)", i)
	}
	return representationEnumStrings[representationEnumIndex[o]:representationEnumIndex[o+1]]
}

// Tag returns the string representation of a Representation. This is an alias for String.
func (i Representation) Tag() string {
	return i.String()
}

// Ordinal returns the ordinal number of a Representation. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (i Representation) Ordinal() int {
	switch i {
	case Identifier:
		return 0
	case Tag:
		return 1
	case Number:
		return 2
	case Ordinal:
		return 3
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Representation) Int() int {
	return int(i)
}

// RepresentationOf returns a Representation based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Representation is returned.
func RepresentationOf(i int) Representation {
	if 0 <= i && i < len(AllRepresentations) {
		return AllRepresentations[i]
	}
	// an invalid result
	return Identifier + Tag + Number + Ordinal + 1
}

// IsValid determines whether a Representation is one of the defined constants.
func (i Representation) IsValid() bool {
	return i.Ordinal() >= 0
}

// Parse parses a string to find the corresponding Representation, accepting one of the string values or
// a number. The input representation is determined by representationMarshalTextRep. It is used by AsRepresentation.
// The input case does not matter.
//
// Usage Example
//
//    v := new(Representation)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Representation) Parse(s string) error {
	return v.parse(s, Identifier)
}

func (v *Representation) parse(in string, rep Representation) error {
	if rep == Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := strings.ToLower(in)

	if v.parseString(s) {
		return nil
	}

	return errors.New(in + ": unrecognised representation")
}

// parseNumber attempts to convert a decimal value
func (v *Representation) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Representation(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Representation) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllRepresentations) {
		*v = AllRepresentations[ord]
		return true
	}
	return false
}

// parseString attempts to match an identifier.
func (v *Representation) parseString(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(representationEnumIndex); j++ {
		i1 := representationEnumIndex[j]
		p := representationEnumInputs[i0:i1]
		if s == p {
			*v = AllRepresentations[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsRepresentation parses a string to find the corresponding Representation, accepting either one of the string values or
// a number. The input representation is determined by representationMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsRepresentation(s string) (Representation, error) {
	var i = new(Representation)
	err := i.Parse(s)
	return *i, err
}

// MustParseRepresentation is similar to AsRepresentation except that it panics on error.
// The input case does not matter.
func MustParseRepresentation(s string) Representation {
	i, err := AsRepresentation(s)
	if err != nil {
		panic(err)
	}
	return i
}
