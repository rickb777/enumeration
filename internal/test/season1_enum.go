// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.0.2

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
)

// AllSeason1s lists all 4 values in order.
var AllSeason1s = []Season1{
	Spring1, Summer1, Autumn1, Winter1,
}

// AllSeason1Enums lists all 4 values in order.
var AllSeason1Enums = enum.IntEnums{
	Spring1, Summer1, Autumn1, Winter1,
}

const (
	season1EnumStrings = "SpringSummerAutumnWinter"
)

var (
	season1EnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season1, which is
// the same as the const identifier but without prefix or suffix.
func (v Season1) String() string {
	o := v.Ordinal()
	return v.toString(o, season1EnumStrings, season1EnumIndex[:])
}

func (v Season1) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason1s) {
		return fmt.Sprintf("Season1(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season1. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season1) Ordinal() int {
	switch v {
	case Spring1:
		return 0
	case Summer1:
		return 1
	case Autumn1:
		return 2
	case Winter1:
		return 3
	}
	return -1
}

// IsValid determines whether a Season1 is one of the defined constants.
func (v Season1) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season1) Int() int {
	return int(v)
}

// Season1Of returns a Season1 based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season1 is returned.
func Season1Of(v int) Season1 {
	if 0 <= v && v < len(AllSeason1s) {
		return AllSeason1s[v]
	}
	// an invalid result
	return Spring1 + Summer1 + Autumn1 + Winter1 + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season1) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season1(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season1, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason1.
//
// Usage Example
//
//	v := new(Season1)
//	err := v.Parse(s)
//	...  etc
func (v *Season1) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season1TransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season1) parseFallback(in, s string) error {
	if v.parseString(s, season1EnumStrings, season1EnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season1")
}

// season1TransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season1TransformInput = func(in string) string {
	return in
}

func (v *Season1) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason1s[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason1 parses a string to find the corresponding Season1, accepting either one of the string values or
// a number. The input representation is determined by season1MarshalTextRep. It wraps Parse.
func AsSeason1(s string) (Season1, error) {
	var v = new(Season1)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason1 is similar to AsSeason1 except that it panics on error.
func MustParseSeason1(s string) Season1 {
	v, err := AsSeason1(s)
	if err != nil {
		panic(err)
	}
	return v
}