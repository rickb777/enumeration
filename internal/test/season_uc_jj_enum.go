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

// AllSeason_Uc_Jjs lists all 4 values in order.
var AllSeason_Uc_Jjs = []Season_Uc_Jj{
	Spring_Uc_Jj, Summer_Uc_Jj, Autumn_Uc_Jj, Winter_Uc_Jj,
}

// AllSeason_Uc_JjEnums lists all 4 values in order.
var AllSeason_Uc_JjEnums = enum.IntEnums{
	Spring_Uc_Jj, Summer_Uc_Jj, Autumn_Uc_Jj, Winter_Uc_Jj,
}

const (
	season_uc_jjEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
	season_uc_jjJSONStrings = "SprgSumrAutmWint"
)

var (
	season_uc_jjEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_uc_jjJSONIndex = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Uc_Jj. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Jj) Ordinal() int {
	switch v {
	case Spring_Uc_Jj:
		return 0
	case Summer_Uc_Jj:
		return 1
	case Autumn_Uc_Jj:
		return 2
	case Winter_Uc_Jj:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Jj, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Jj) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jjEnumStrings, season_uc_jjEnumIndex[:])
}

func (v Season_Uc_Jj) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jjs) {
		return fmt.Sprintf("Season_Uc_Jj(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Jj is one of the defined constants.
func (v Season_Uc_Jj) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Jj) Int() int {
	return int(v)
}

var invalidSeason_Uc_JjValue = func() Season_Uc_Jj {
	var v Season_Uc_Jj
	for {
		if !slices.Contains(AllSeason_Uc_Jjs, v) {
			return v
		}
		v++
	} // AllSeason_Uc_Jjs is a finite set so loop will terminate eventually
}()

// Season_Uc_JjOf returns a Season_Uc_Jj based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Jj is returned.
func Season_Uc_JjOf(v int) Season_Uc_Jj {
	if 0 <= v && v < len(AllSeason_Uc_Jjs) {
		return AllSeason_Uc_Jjs[v]
	}
	return invalidSeason_Uc_JjValue
}

// Parse parses a string to find the corresponding Season_Uc_Jj, accepting one of the string values or
// a number. It is used by AsSeason_Uc_Jj.
//
// Usage Example
//
//	v := new(Season_Uc_Jj)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Jj) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_jjTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Jj) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Jj(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Uc_Jj) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_jjEnumStrings, season_uc_jjEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_jj")
}

func (v *Season_Uc_Jj) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Jjs[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_uc_jjTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_jjTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

// AsSeason_Uc_Jj parses a string to find the corresponding Season_Uc_Jj, accepting either one of the string values or
// a number. It wraps Parse.
func AsSeason_Uc_Jj(s string) (Season_Uc_Jj, error) {
	var v = new(Season_Uc_Jj)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Jj is similar to AsSeason_Uc_Jj except that it panics on error.
func MustParseSeason_Uc_Jj(s string) Season_Uc_Jj {
	v, err := AsSeason_Uc_Jj(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Uc_Jj) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_uc_jjJSONStrings, season_uc_jjJSONIndex[:])
}

func (v Season_Uc_Jj) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Jj) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Jj) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_jj", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Season_Uc_Jj) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_uc_jjJSONStrings, season_uc_jjJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Uc_Jj) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Uc_Jj) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_jjTransformInput(in)

	if v.parseString(s, season_uc_jjJSONStrings, season_uc_jjJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, season_uc_jjEnumStrings, season_uc_jjEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_jj")
}

// season_uc_jjMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_uc_jjMarshalNumber = func(v Season_Uc_Jj) string {
	return strconv.FormatInt(int64(v), 10)
}
