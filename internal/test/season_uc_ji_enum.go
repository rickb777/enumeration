// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.0.2

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
	"strings"
)

// AllSeason_Uc_Jis lists all 4 values in order.
var AllSeason_Uc_Jis = []Season_Uc_Ji{
	Spring_Uc_Ji, Summer_Uc_Ji, Autumn_Uc_Ji, Winter_Uc_Ji,
}

// AllSeason_Uc_JiEnums lists all 4 values in order.
var AllSeason_Uc_JiEnums = enum.IntEnums{
	Spring_Uc_Ji, Summer_Uc_Ji, Autumn_Uc_Ji, Winter_Uc_Ji,
}

const (
	season_uc_jiEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_jiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Ji, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ji) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:])
}

func (v Season_Uc_Ji) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jis) {
		return fmt.Sprintf("Season_Uc_Ji(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Uc_Ji. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ji) Ordinal() int {
	switch v {
	case Spring_Uc_Ji:
		return 0
	case Summer_Uc_Ji:
		return 1
	case Autumn_Uc_Ji:
		return 2
	case Winter_Uc_Ji:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Uc_Ji is one of the defined constants.
func (v Season_Uc_Ji) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ji) Int() int {
	return int(v)
}

// Season_Uc_JiOf returns a Season_Uc_Ji based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Ji is returned.
func Season_Uc_JiOf(v int) Season_Uc_Ji {
	if 0 <= v && v < len(AllSeason_Uc_Jis) {
		return AllSeason_Uc_Jis[v]
	}
	// an invalid result
	return Spring_Uc_Ji + Summer_Uc_Ji + Autumn_Uc_Ji + Winter_Uc_Ji + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Ji) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Ji(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Uc_Ji, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Uc_Ji.
//
// Usage Example
//
//	v := new(Season_Uc_Ji)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Ji) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_jiTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Uc_Ji) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_ji")
}

// season_uc_jiTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_jiTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

func (v *Season_Uc_Ji) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Jis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Uc_Ji parses a string to find the corresponding Season_Uc_Ji, accepting either one of the string values or
// a number. The input representation is determined by season_uc_jiMarshalTextRep. It wraps Parse.
func AsSeason_Uc_Ji(s string) (Season_Uc_Ji, error) {
	var v = new(Season_Uc_Ji)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Ji is similar to AsSeason_Uc_Ji except that it panics on error.
func MustParseSeason_Uc_Ji(s string) Season_Uc_Ji {
	v, err := AsSeason_Uc_Ji(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Uc_Ji) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:])
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v Season_Uc_Ji) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:])
	return enum.QuotedString(s), nil
}

func (v Season_Uc_Ji) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Ji) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Ji) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_ji", v)
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Uc_Ji) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Uc_Ji) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_jiTransformInput(in)

	if v.parseString(s, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_ji")
}
