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

// AllSeason_Ic_Jos lists all 4 values in order.
var AllSeason_Ic_Jos = []Season_Ic_Jo{
	Spring_Ic_Jo, Summer_Ic_Jo, Autumn_Ic_Jo, Winter_Ic_Jo,
}

// AllSeason_Ic_JoEnums lists all 4 values in order.
var AllSeason_Ic_JoEnums = enum.IntEnums{
	Spring_Ic_Jo, Summer_Ic_Jo, Autumn_Ic_Jo, Winter_Ic_Jo,
}

const (
	season_ic_joEnumStrings = "SpringSummerAutumnWinter"
	season_ic_joEnumInputs  = "springsummerautumnwinter"
)

var (
	season_ic_joEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Ic_Jo, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Jo) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_joEnumStrings, season_ic_joEnumIndex[:])
}

func (v Season_Ic_Jo) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Jos) {
		return fmt.Sprintf("Season_Ic_Jo(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Ic_Jo. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Jo) Ordinal() int {
	switch v {
	case Spring_Ic_Jo:
		return 0
	case Summer_Ic_Jo:
		return 1
	case Autumn_Ic_Jo:
		return 2
	case Winter_Ic_Jo:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Ic_Jo is one of the defined constants.
func (v Season_Ic_Jo) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Jo) Int() int {
	return int(v)
}

// Season_Ic_JoOf returns a Season_Ic_Jo based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Jo is returned.
func Season_Ic_JoOf(v int) Season_Ic_Jo {
	if 0 <= v && v < len(AllSeason_Ic_Jos) {
		return AllSeason_Ic_Jos[v]
	}
	// an invalid result
	return Spring_Ic_Jo + Summer_Ic_Jo + Autumn_Ic_Jo + Winter_Ic_Jo + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Jo) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Jo(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Ic_Jo, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Ic_Jo.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_Jo)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_Jo) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_joTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Ic_Jo) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_joEnumInputs, season_ic_joEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_jo")
}

// season_ic_joTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_joTransformInput = func(in string) string {
	return strings.ToLower(in)
}

func (v *Season_Ic_Jo) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Jos[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Ic_Jo parses a string to find the corresponding Season_Ic_Jo, accepting either one of the string values or
// a number. The input representation is determined by season_ic_joMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Jo(s string) (Season_Ic_Jo, error) {
	var v = new(Season_Ic_Jo)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Jo is similar to AsSeason_Ic_Jo except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Jo(s string) Season_Ic_Jo {
	v, err := AsSeason_Ic_Jo(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The ordinal representation is chosen according to -marshaljson.
func (v Season_Ic_Jo) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, v.invalidError()
	}

	s := strconv.Itoa(o)
	return []byte(s), nil
}

func (v Season_Ic_Jo) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_jo", v)
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Ic_Jo) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

// parseOrdinal attempts to convert an ordinal value.
func (v *Season_Ic_Jo) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSeason_Ic_Jos) {
		*v = AllSeason_Ic_Jos[ord]
		return true
	}
	return false
}

func (v *Season_Ic_Jo) unmarshalJSON(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := season_ic_joTransformInput(in)

	return v.parseFallback(in, s)
}
