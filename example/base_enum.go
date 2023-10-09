// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.2

package example

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
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

// Ordinal returns the ordinal number of a Base. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Base) Ordinal() int {
	switch v {
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

// String returns the literal string representation of a Base, which is
// the same as the const identifier but without prefix or suffix.
func (v Base) String() string {
	o := v.Ordinal()
	return v.toString(o, baseEnumStrings, baseEnumIndex[:])
}

func (v Base) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllBases) {
		return fmt.Sprintf("Base(%g)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Base is one of the defined constants.
func (v Base) IsValid() bool {
	return v.Ordinal() >= 0
}

// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (v Base) Float() float64 {
	return float64(v)
}

// BaseOf returns a Base based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Base is returned.
func BaseOf(v int) Base {
	if 0 <= v && v < len(AllBases) {
		return AllBases[v]
	}
	// an invalid result
	return A + C + G + T + 1
}

// Parse parses a string to find the corresponding Base, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsBase.
//
// Usage Example
//
//    v := new(Base)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Base) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := baseTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Base) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*v = Base(num)
		return v.IsValid()
	}
	return false
}

func (v *Base) parseFallback(in, s string) error {
	if v.parseString(s, baseEnumStrings, baseEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised base")
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

// baseTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var baseTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsBase parses a string to find the corresponding Base, accepting either one of the string values or
// a number. The input representation is determined by baseMarshalTextRep. It wraps Parse.
func AsBase(s string) (Base, error) {
	var v = new(Base)
	err := v.Parse(s)
	return *v, err
}

// MustParseBase is similar to AsBase except that it panics on error.
func MustParseBase(s string) Base {
	v, err := AsBase(s)
	if err != nil {
		panic(err)
	}
	return v
}
