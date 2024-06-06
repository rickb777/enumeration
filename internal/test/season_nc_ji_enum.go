// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.2.0

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Nc_Jis lists all 4 values in order.
var AllSeason_Nc_Jis = []Season_Nc_Ji{
	Spring_Nc_Ji, Summer_Nc_Ji, Autumn_Nc_Ji, Winter_Nc_Ji,
}

// AllSeason_Nc_JiEnums lists all 4 values in order.
var AllSeason_Nc_JiEnums = enum.IntEnums{
	Spring_Nc_Ji, Summer_Nc_Ji, Autumn_Nc_Ji, Winter_Nc_Ji,
}

const (
	season_nc_jiEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_jiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Ji. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ji) Ordinal() int {
	switch v {
	case Spring_Nc_Ji:
		return 0
	case Summer_Nc_Ji:
		return 1
	case Autumn_Nc_Ji:
		return 2
	case Winter_Nc_Ji:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Ji, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ji) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:])
}

func (v Season_Nc_Ji) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Jis) {
		return fmt.Sprintf("Season_Nc_Ji(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Ji is one of the defined constants.
func (v Season_Nc_Ji) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ji) Int() int {
	return int(v)
}

var invalidSeason_Nc_JiValue = func() Season_Nc_Ji {
	var v Season_Nc_Ji
	for {
		if !slices.Contains(AllSeason_Nc_Jis, v) {
			return v
		}
		v++
	} // AllSeason_Nc_Jis is a finite set so loop will terminate eventually
}()

// Season_Nc_JiOf returns a Season_Nc_Ji based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Ji is returned.
func Season_Nc_JiOf(v int) Season_Nc_Ji {
	if 0 <= v && v < len(AllSeason_Nc_Jis) {
		return AllSeason_Nc_Jis[v]
	}
	return invalidSeason_Nc_JiValue
}

// Parse parses a string to find the corresponding Season_Nc_Ji, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Nc_Ji.
//
// Usage Example
//
//	v := new(Season_Nc_Ji)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Ji) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_jiTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Ji) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Ji(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Nc_Ji) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_ji")
}

func (v *Season_Nc_Ji) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Jis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_nc_jiTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_jiTransformInput = func(in string) string {
	return in
}

// AsSeason_Nc_Ji parses a string to find the corresponding Season_Nc_Ji, accepting either one of the string values or
// a number. The input representation is determined by season_nc_jiMarshalTextRep. It wraps Parse.
func AsSeason_Nc_Ji(s string) (Season_Nc_Ji, error) {
	var v = new(Season_Nc_Ji)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Ji is similar to AsSeason_Nc_Ji except that it panics on error.
func MustParseSeason_Nc_Ji(s string) Season_Nc_Ji {
	v, err := AsSeason_Nc_Ji(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Nc_Ji) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:])
}

func (v Season_Nc_Ji) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Ji) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Ji) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ji", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v Season_Nc_Ji) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Nc_Ji) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Nc_Ji) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_jiTransformInput(in)

	if v.parseString(s, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_ji")
}

// season_nc_jiMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_nc_jiMarshalNumber = func(v Season_Nc_Ji) string {
	return strconv.FormatInt(int64(v), 10)
}
