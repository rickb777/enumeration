// generated code - do not edit
// github.com/rickb777/enumeration/v4 v3.0.2

package transform

import (
	"errors"
	"fmt"
	"strconv"
)

// AllCases lists all 3 values in order.
var AllCases = []Case{
	Stet, Upper, Lower,
}

const (
	caseEnumStrings = "StetUpperLower"
)

var (
	caseEnumIndex = [...]uint16{0, 4, 9, 14}
)

// String returns the literal string representation of a Case, which is
// the same as the const identifier but without prefix or suffix.
func (v Case) String() string {
	o := v.Ordinal()
	return v.toString(o, caseEnumStrings, caseEnumIndex[:])
}

func (v Case) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllCases) {
		return fmt.Sprintf("Case(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Case. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Case) Ordinal() int {
	switch v {
	case Stet:
		return 0
	case Upper:
		return 1
	case Lower:
		return 2
	}
	return -1
}

// IsValid determines whether a Case is one of the defined constants.
func (v Case) IsValid() bool {
	return v.Ordinal() >= 0
}

// CaseOf returns a Case based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Case is returned.
func CaseOf(v int) Case {
	if 0 <= v && v < len(AllCases) {
		return AllCases[v]
	}
	// an invalid result
	return Stet + Upper + Lower + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Case) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Case(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Case, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsCase.
//
// Usage Example
//
//	v := new(Case)
//	err := v.Parse(s)
//	...  etc
func (v *Case) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := caseTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Case) parseFallback(in, s string) error {
	if v.parseString(s, caseEnumStrings, caseEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised case")
}

// caseTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var caseTransformInput = func(in string) string {
	return in
}

func (v *Case) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllCases[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsCase parses a string to find the corresponding Case, accepting either one of the string values or
// a number. The input representation is determined by caseMarshalTextRep. It wraps Parse.
func AsCase(s string) (Case, error) {
	var v = new(Case)
	err := v.Parse(s)
	return *v, err
}

// MustParseCase is similar to AsCase except that it panics on error.
func MustParseCase(s string) Case {
	v, err := AsCase(s)
	if err != nil {
		panic(err)
	}
	return v
}
