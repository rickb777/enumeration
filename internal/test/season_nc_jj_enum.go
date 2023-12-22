// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.4

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
	"strings"
)

// AllSeason_Nc_Jjs lists all 4 values in order.
var AllSeason_Nc_Jjs = []Season_Nc_Jj{
	Spring_Nc_Jj, Summer_Nc_Jj, Autumn_Nc_Jj, Winter_Nc_Jj,
}

// AllSeason_Nc_JjEnums lists all 4 values in order.
var AllSeason_Nc_JjEnums = enum.IntEnums{
	Spring_Nc_Jj, Summer_Nc_Jj, Autumn_Nc_Jj, Winter_Nc_Jj,
}

const (
	season_nc_jjEnumStrings = "SpringSummerAutumnWinter"
	season_nc_jjJSONStrings = "SprgSumrAutmWint"
)

var (
	season_nc_jjEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_nc_jjJSONIndex = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Nc_Jj. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Jj) Ordinal() int {
	switch v {
	case Spring_Nc_Jj:
		return 0
	case Summer_Nc_Jj:
		return 1
	case Autumn_Nc_Jj:
		return 2
	case Winter_Nc_Jj:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Jj, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Jj) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_jjEnumStrings, season_nc_jjEnumIndex[:])
}

func (v Season_Nc_Jj) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Jjs) {
		return fmt.Sprintf("Season_Nc_Jj(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Jj is one of the defined constants.
func (v Season_Nc_Jj) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Jj) Int() int {
	return int(v)
}

// Season_Nc_JjOf returns a Season_Nc_Jj based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Jj is returned.
func Season_Nc_JjOf(v int) Season_Nc_Jj {
	if 0 <= v && v < len(AllSeason_Nc_Jjs) {
		return AllSeason_Nc_Jjs[v]
	}
	// an invalid result
	return Spring_Nc_Jj + Summer_Nc_Jj + Autumn_Nc_Jj + Winter_Nc_Jj + 1
}

// Parse parses a string to find the corresponding Season_Nc_Jj, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Nc_Jj.
//
// Usage Example
//
//	v := new(Season_Nc_Jj)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Jj) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_jjTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Jj) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Jj(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Nc_Jj) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_jjEnumStrings, season_nc_jjEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_jj")
}

func (v *Season_Nc_Jj) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Jjs[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_nc_jjTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_jjTransformInput = func(in string) string {
	return in
}

// AsSeason_Nc_Jj parses a string to find the corresponding Season_Nc_Jj, accepting either one of the string values or
// a number. The input representation is determined by season_nc_jjMarshalTextRep. It wraps Parse.
func AsSeason_Nc_Jj(s string) (Season_Nc_Jj, error) {
	var v = new(Season_Nc_Jj)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Jj is similar to AsSeason_Nc_Jj except that it panics on error.
func MustParseSeason_Nc_Jj(s string) Season_Nc_Jj {
	v, err := AsSeason_Nc_Jj(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Nc_Jj) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_nc_jjJSONStrings, season_nc_jjJSONIndex[:])
}

func (v Season_Nc_Jj) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Jj) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Jj) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_jj", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Season_Nc_Jj) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_nc_jjJSONStrings, season_nc_jjJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Nc_Jj) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Nc_Jj) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_jjTransformInput(in)

	if v.parseString(s, season_nc_jjJSONStrings, season_nc_jjJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, season_nc_jjEnumStrings, season_nc_jjEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_jj")
}

// season_nc_jjMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_nc_jjMarshalNumber = func(v Season_Nc_Jj) string {
	return strconv.FormatInt(int64(v), 10)
}
