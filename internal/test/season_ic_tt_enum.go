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

// AllSeason_Ic_Tts lists all 4 values in order.
var AllSeason_Ic_Tts = []Season_Ic_Tt{
	Spring_Ic_Tt, Summer_Ic_Tt, Autumn_Ic_Tt, Winter_Ic_Tt,
}

// AllSeason_Ic_TtEnums lists all 4 values in order.
var AllSeason_Ic_TtEnums = enum.IntEnums{
	Spring_Ic_Tt, Summer_Ic_Tt, Autumn_Ic_Tt, Winter_Ic_Tt,
}

const (
	season_ic_ttEnumStrings = "SpringSummerAutumnWinter"
	season_ic_ttEnumInputs  = "springsummerautumnwinter"
	season_ic_ttTextStrings = "SprgSumrAutmWint"
	season_ic_ttTextInputs  = "sprgsumrautmwint"
)

var (
	season_ic_ttEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_ic_ttTextIndex = [...]uint16{0, 4, 8, 12, 16}
)

// String returns the literal string representation of a Season_Ic_Tt, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Tt) String() string {
	return v.toString(season_ic_ttEnumStrings, season_ic_ttEnumIndex[:])
}

func (v Season_Ic_Tt) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllSeason_Ic_Tts) {
		return fmt.Sprintf("Season_Ic_Tt(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Ic_Tt. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Tt) Ordinal() int {
	switch v {
	case Spring_Ic_Tt:
		return 0
	case Summer_Ic_Tt:
		return 1
	case Autumn_Ic_Tt:
		return 2
	case Winter_Ic_Tt:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Ic_Tt is one of the defined constants.
func (v Season_Ic_Tt) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Tt) Int() int {
	return int(v)
}

// Season_Ic_TtOf returns a Season_Ic_Tt based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Tt is returned.
func Season_Ic_TtOf(v int) Season_Ic_Tt {
	if 0 <= v && v < len(AllSeason_Ic_Tts) {
		return AllSeason_Ic_Tts[v]
	}
	// an invalid result
	return Spring_Ic_Tt + Summer_Ic_Tt + Autumn_Ic_Tt + Winter_Ic_Tt + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Tt) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Tt(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Ic_Tt, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Ic_Tt.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_Tt)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_Tt) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_ttTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Ic_Tt) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_ttEnumInputs, season_ic_ttEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_tt")
}

// season_ic_ttTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_ttTransformInput = func(in string) string {
	return strings.ToLower(in)
}

func (v *Season_Ic_Tt) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Tts[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Ic_Tt parses a string to find the corresponding Season_Ic_Tt, accepting either one of the string values or
// a number. The input representation is determined by season_ic_ttMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Tt(s string) (Season_Ic_Tt, error) {
	var v = new(Season_Ic_Tt)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Tt is similar to AsSeason_Ic_Tt except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Tt(s string) Season_Ic_Tt {
	v, err := AsSeason_Ic_Tt(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Season_Ic_Tt) MarshalText() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}
	s := season_ic_ttTextStrings[season_ic_ttTextIndex[o]:season_ic_ttTextIndex[o+1]]
	return []byte(s), nil
}

func (v Season_Ic_Tt) marshalNumberOrError() ([]byte, error) {
	return nil, v.invalidError()
}

func (v Season_Ic_Tt) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_tt", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Ic_Tt) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Ic_Tt) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_ttTransformInput(in)

	if v.parseString(s, season_ic_ttTextInputs, season_ic_ttTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}
