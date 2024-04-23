// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.5

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
	"strings"
)

// AllSeason_Ic_Jis lists all 4 values in order.
var AllSeason_Ic_Jis = []Season_Ic_Ji{
	Spring_Ic_Ji, Summer_Ic_Ji, Autumn_Ic_Ji, Winter_Ic_Ji,
}

// AllSeason_Ic_JiEnums lists all 4 values in order.
var AllSeason_Ic_JiEnums = enum.IntEnums{
	Spring_Ic_Ji, Summer_Ic_Ji, Autumn_Ic_Ji, Winter_Ic_Ji,
}

const (
	season_ic_jiEnumStrings = "SpringSummerAutumnWinter"
	season_ic_jiEnumInputs  = "springsummerautumnwinter"
)

var (
	season_ic_jiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Ic_Ji. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Ji) Ordinal() int {
	switch v {
	case Spring_Ic_Ji:
		return 0
	case Summer_Ic_Ji:
		return 1
	case Autumn_Ic_Ji:
		return 2
	case Winter_Ic_Ji:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Ic_Ji, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Ji) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_jiEnumStrings, season_ic_jiEnumIndex[:])
}

func (v Season_Ic_Ji) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Jis) {
		return fmt.Sprintf("Season_Ic_Ji(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Ic_Ji is one of the defined constants.
func (v Season_Ic_Ji) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Ji) Int() int {
	return int(v)
}

// Season_Ic_JiOf returns a Season_Ic_Ji based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Ji is returned.
func Season_Ic_JiOf(v int) Season_Ic_Ji {
	if 0 <= v && v < len(AllSeason_Ic_Jis) {
		return AllSeason_Ic_Jis[v]
	}
	// an invalid result
	return Spring_Ic_Ji + Summer_Ic_Ji + Autumn_Ic_Ji + Winter_Ic_Ji + 1
}

// Parse parses a string to find the corresponding Season_Ic_Ji, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Ic_Ji.
// The input case does not matter.
//
// Usage Example
//
//    v := new(Season_Ic_Ji)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Season_Ic_Ji) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_jiTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Ji) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Ji(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Ic_Ji) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_jiEnumInputs, season_ic_jiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_ji")
}

func (v *Season_Ic_Ji) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Jis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_ic_jiTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_jiTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsSeason_Ic_Ji parses a string to find the corresponding Season_Ic_Ji, accepting either one of the string values or
// a number. The input representation is determined by season_ic_jiMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Ji(s string) (Season_Ic_Ji, error) {
	var v = new(Season_Ic_Ji)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Ji is similar to AsSeason_Ic_Ji except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Ji(s string) Season_Ic_Ji {
	v, err := AsSeason_Ic_Ji(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Ic_Ji) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_ic_jiEnumStrings, season_ic_jiEnumIndex[:])
}

func (v Season_Ic_Ji) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Ic_Ji) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Ic_Ji) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_ji", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v Season_Ic_Ji) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_ic_jiEnumStrings, season_ic_jiEnumIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Ic_Ji) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Ic_Ji) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_jiTransformInput(in)

	if v.parseString(s, season_ic_jiEnumInputs, season_ic_jiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_ji")
}

// season_ic_jiMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_ic_jiMarshalNumber = func(v Season_Ic_Ji) string {
	return strconv.FormatInt(int64(v), 10)
}
