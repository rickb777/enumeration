// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.1

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Ic_Jjs lists all 4 values in order.
var AllSeason_Ic_Jjs = []Season_Ic_Jj{
	Spring_Ic_Jj, Summer_Ic_Jj, Autumn_Ic_Jj, Winter_Ic_Jj,
}

// AllSeason_Ic_JjEnums lists all 4 values in order.
var AllSeason_Ic_JjEnums = enum.IntEnums{
	Spring_Ic_Jj, Summer_Ic_Jj, Autumn_Ic_Jj, Winter_Ic_Jj,
}

const (
	season_ic_jjEnumStrings = "SpringSummerAutumnWinter"
	season_ic_jjEnumInputs  = "springsummerautumnwinter"
	season_ic_jjJSONStrings = "SprgSumrAutmWint"
	season_ic_jjJSONInputs  = "SprgSumrAutmWint"
)

var (
	season_ic_jjEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_ic_jjJSONIndex = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Ic_Jj. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Jj) Ordinal() int {
	switch v {
	case Spring_Ic_Jj:
		return 0
	case Summer_Ic_Jj:
		return 1
	case Autumn_Ic_Jj:
		return 2
	case Winter_Ic_Jj:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Ic_Jj, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Jj) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_jjEnumStrings, season_ic_jjEnumIndex[:])
}

func (v Season_Ic_Jj) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Jjs) {
		return fmt.Sprintf("Season_Ic_Jj(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Ic_Jj is one of the defined constants.
func (v Season_Ic_Jj) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Jj) Int() int {
	return int(v)
}

var invalidSeason_Ic_JjValue = func() Season_Ic_Jj {
	var v Season_Ic_Jj
	for {
		if !slices.Contains(AllSeason_Ic_Jjs, v) {
			return v
		}
		v++
	} // AllSeason_Ic_Jjs is a finite set so loop will terminate eventually
}()

// Season_Ic_JjOf returns a Season_Ic_Jj based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Jj is returned.
func Season_Ic_JjOf(v int) Season_Ic_Jj {
	if 0 <= v && v < len(AllSeason_Ic_Jjs) {
		return AllSeason_Ic_Jjs[v]
	}
	return invalidSeason_Ic_JjValue
}

// Parse parses a string to find the corresponding Season_Ic_Jj, accepting one of the string values or
// a number. It is used by AsSeason_Ic_Jj.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_Jj)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_Jj) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_jjTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Jj) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Jj(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Ic_Jj) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_jjEnumInputs, season_ic_jjEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_jj")
}

func (v *Season_Ic_Jj) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Jjs[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_ic_jjTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_jjTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsSeason_Ic_Jj parses a string to find the corresponding Season_Ic_Jj, accepting either one of the string values or
// a number. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Jj(s string) (Season_Ic_Jj, error) {
	var v = new(Season_Ic_Jj)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Jj is similar to AsSeason_Ic_Jj except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Jj(s string) Season_Ic_Jj {
	v, err := AsSeason_Ic_Jj(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Ic_Jj) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_ic_jjJSONStrings, season_ic_jjJSONIndex[:])
}

func (v Season_Ic_Jj) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Ic_Jj) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Ic_Jj) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_jj", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Season_Ic_Jj) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_ic_jjJSONStrings, season_ic_jjJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Ic_Jj) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Ic_Jj) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_jjTransformInput(in)

	if v.parseString(s, season_ic_jjJSONInputs, season_ic_jjJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, season_ic_jjEnumInputs, season_ic_jjEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_jj")
}

// season_ic_jjMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_ic_jjMarshalNumber = func(v Season_Ic_Jj) string {
	return strconv.FormatInt(int64(v), 10)
}
