// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.0.2

package enum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// AllRepresentations lists all 3 values in order.
var AllRepresentations = []Representation{
	None, Identifier, Number,
}

const (
	representationEnumStrings = "NoneIdentifierNumber"
	representationEnumInputs  = "noneidentifiernumber"
)

var (
	representationEnumIndex = [...]uint16{0, 4, 14, 20}
)

// String returns the literal string representation of a Representation, which is
// the same as the const identifier but without prefix or suffix.
func (v Representation) String() string {
	o := v.Ordinal()
	return v.toString(o, representationEnumStrings, representationEnumIndex[:])
}

func (v Representation) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllRepresentations) {
		return fmt.Sprintf("Representation(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Representation. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Representation) Ordinal() int {
	switch v {
	case None:
		return 0
	case Identifier:
		return 1
	case Number:
		return 2
	}
	return -1
}

// IsValid determines whether a Representation is one of the defined constants.
func (v Representation) IsValid() bool {
	return v.Ordinal() >= 0
}

// RepresentationOf returns a Representation based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Representation is returned.
func RepresentationOf(v int) Representation {
	if 0 <= v && v < len(AllRepresentations) {
		return AllRepresentations[v]
	}
	// an invalid result
	return None + Identifier + Number + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Representation) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Representation(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Representation, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsRepresentation.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Representation)
//	err := v.Parse(s)
//	...  etc
func (v *Representation) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := representationTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Representation) parseFallback(in, s string) error {
	if v.parseString(s, representationEnumInputs, representationEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised representation")
}

// representationTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var representationTransformInput = func(in string) string {
	return strings.ToLower(in)
}

func (v *Representation) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
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
	var v = new(Representation)
	err := v.Parse(s)
	return *v, err
}

// MustParseRepresentation is similar to AsRepresentation except that it panics on error.
// The input case does not matter.
func MustParseRepresentation(s string) Representation {
	v, err := AsRepresentation(s)
	if err != nil {
		panic(err)
	}
	return v
}
