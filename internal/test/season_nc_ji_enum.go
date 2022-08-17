// generated code - do not edit
// github.com/rickb777/enumeration/v3 v2.14.0

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
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

// String returns the literal string representation of a Season_Nc_Ji, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ji) String() string {
	return v.toString(season_nc_jiEnumStrings, season_nc_jiEnumIndex[:])
}

func (v Season_Nc_Ji) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllSeason_Nc_Jis) {
		return fmt.Sprintf("Season_Nc_Ji(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

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

// IsValid determines whether a Season_Nc_Ji is one of the defined constants.
func (v Season_Nc_Ji) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ji) Int() int {
	return int(v)
}

// Season_Nc_JiOf returns a Season_Nc_Ji based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Ji is returned.
func Season_Nc_JiOf(v int) Season_Nc_Ji {
	if 0 <= v && v < len(AllSeason_Nc_Jis) {
		return AllSeason_Nc_Jis[v]
	}
	// an invalid result
	return Spring_Nc_Ji + Summer_Nc_Ji + Autumn_Nc_Ji + Winter_Nc_Ji + 1
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

func (v *Season_Nc_Ji) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_ji")
}

// season_nc_jiTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_jiTransformInput = func(in string) string {
	return in
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

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v Season_Nc_Ji) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	return enum.QuotedString(v.String()), nil
}

func (v Season_Nc_Ji) marshalNumberOrError() ([]byte, error) {
	return nil, v.invalidError()
}

func (v Season_Nc_Ji) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ji", v)
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

	return v.parseFallback(in, s)
}
