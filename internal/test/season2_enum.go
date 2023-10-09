// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.2

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
)

// AllSeason2s lists all 4 values in order.
var AllSeason2s = []Season2{
	Spring2, Summer2, Autumn2, Winter2,
}

// AllSeason2Enums lists all 4 values in order.
var AllSeason2Enums = enum.IntEnums{
	Spring2, Summer2, Autumn2, Winter2,
}

const (
	season2EnumStrings = "SpringSummerAutumnWinter"
)

var (
	season2EnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season2. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season2) Ordinal() int {
	switch v {
	case Spring2:
		return 0
	case Summer2:
		return 1
	case Autumn2:
		return 2
	case Winter2:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season2, which is
// the same as the const identifier but without prefix or suffix.
func (v Season2) String() string {
	o := v.Ordinal()
	return v.toString(o, season2EnumStrings, season2EnumIndex[:])
}

func (v Season2) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason2s) {
		return fmt.Sprintf("Season2(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season2 is one of the defined constants.
func (v Season2) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season2) Int() int {
	return int(v)
}

// Season2Of returns a Season2 based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season2 is returned.
func Season2Of(v int) Season2 {
	if 0 <= v && v < len(AllSeason2s) {
		return AllSeason2s[v]
	}
	// an invalid result
	return Spring2 + Summer2 + Autumn2 + Winter2 + 1
}

// Parse parses a string to find the corresponding Season2, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason2.
//
// Usage Example
//
//    v := new(Season2)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Season2) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season2TransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season2) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season2(num)
		return v.IsValid()
	}
	return false
}

func (v *Season2) parseFallback(in, s string) error {
	if v.parseString(s, season2EnumStrings, season2EnumIndex[:]) {
		return nil
	}

	var ok bool
	*v, ok = season2Alias[s]
	if ok {
		return nil
	}

	return errors.New(in + ": unrecognised season2")
}

func (v *Season2) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason2s[j-1]
			return true
		}
		i0 = i1
	}
	*v, ok = season2Alias[s]
	return ok
}

// season2TransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season2TransformInput = func(in string) string {
	return in
}

// AsSeason2 parses a string to find the corresponding Season2, accepting either one of the string values or
// a number. The input representation is determined by season2MarshalTextRep. It wraps Parse.
func AsSeason2(s string) (Season2, error) {
	var v = new(Season2)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason2 is similar to AsSeason2 except that it panics on error.
func MustParseSeason2(s string) Season2 {
	v, err := AsSeason2(s)
	if err != nil {
		panic(err)
	}
	return v
}
