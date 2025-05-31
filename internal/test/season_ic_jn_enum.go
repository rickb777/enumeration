// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.7

package test

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Ic_Jns lists all 4 values in order.
var AllSeason_Ic_Jns = []Season_Ic_Jn{
	Spring_Ic_Jn, Summer_Ic_Jn, Autumn_Ic_Jn, Winter_Ic_Jn,
}

const (
	season_ic_jnEnumStrings = "SpringSummerAutumnWinter"
	season_ic_jnEnumInputs  = "springsummerautumnwinter"
)

var (
	season_ic_jnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Ic_Jn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Jn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_jnEnumStrings, season_ic_jnEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Ic_Jn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Jn) Ordinal() int {
	switch v {
	case Spring_Ic_Jn:
		return 0
	case Summer_Ic_Jn:
		return 1
	case Autumn_Ic_Jn:
		return 2
	case Winter_Ic_Jn:
		return 3
	}
	return -1
}

func (v Season_Ic_Jn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Jns) {
		return fmt.Sprintf("Season_Ic_Jn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Ic_Jn is one of the defined constants.
func (v Season_Ic_Jn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Jn) Int() int {
	return int(v)
}

var invalidSeason_Ic_JnValue = func() Season_Ic_Jn {
	var v Season_Ic_Jn
	for {
		if !slices.Contains(AllSeason_Ic_Jns, v) {
			return v
		}
		v++
	} // AllSeason_Ic_Jns is a finite set so loop will terminate eventually
}()

// Season_Ic_JnOf returns a Season_Ic_Jn based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Jn is returned.
func Season_Ic_JnOf(v int) Season_Ic_Jn {
	if 0 <= v && v < len(AllSeason_Ic_Jns) {
		return AllSeason_Ic_Jns[v]
	}
	return invalidSeason_Ic_JnValue
}

// Parse parses a string to find the corresponding Season_Ic_Jn, accepting one of the string values or
// a number. It is used by AsSeason_Ic_Jn.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_Jn)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_Jn) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_jnTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Jn) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Jn(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Ic_Jn) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_jnEnumInputs, season_ic_jnEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_jn")
}

func (v *Season_Ic_Jn) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Jns[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_ic_jnTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_jnTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsSeason_Ic_Jn parses a string to find the corresponding Season_Ic_Jn, accepting either one of the string values or
// a number. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Jn(s string) (Season_Ic_Jn, error) {
	var v = new(Season_Ic_Jn)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Jn is similar to AsSeason_Ic_Jn except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Jn(s string) Season_Ic_Jn {
	v, err := AsSeason_Ic_Jn(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The number representation is chosen according to -marshaljson.
func (v Season_Ic_Jn) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	s := season_ic_jnMarshalNumber(v)
	return []byte(s), nil
}

func (v Season_Ic_Jn) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Ic_Jn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_jn", v)
}

// season_ic_jnMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_ic_jnMarshalNumber = func(v Season_Ic_Jn) string {
	return strconv.FormatInt(int64(v), 10)
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Ic_Jn) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Ic_Jn) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_jnTransformInput(in)

	if v.parseString(s, season_ic_jnEnumInputs, season_ic_jnEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_jn")
}
