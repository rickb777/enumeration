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

// AllSeason_Uc_Jos lists all 4 values in order.
var AllSeason_Uc_Jos = []Season_Uc_Jo{
	Spring_Uc_Jo, Summer_Uc_Jo, Autumn_Uc_Jo, Winter_Uc_Jo,
}

// AllSeason_Uc_JoEnums lists all 4 values in order.
var AllSeason_Uc_JoEnums = enum.IntEnums{
	Spring_Uc_Jo, Summer_Uc_Jo, Autumn_Uc_Jo, Winter_Uc_Jo,
}

const (
	season_uc_joEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_joEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Jo, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Jo) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_joEnumStrings, season_uc_joEnumIndex[:])
}

func (v Season_Uc_Jo) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jos) {
		return fmt.Sprintf("Season_Uc_Jo(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Uc_Jo. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Jo) Ordinal() int {
	switch v {
	case Spring_Uc_Jo:
		return 0
	case Summer_Uc_Jo:
		return 1
	case Autumn_Uc_Jo:
		return 2
	case Winter_Uc_Jo:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Uc_Jo is one of the defined constants.
func (v Season_Uc_Jo) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Jo) Int() int {
	return int(v)
}

// Season_Uc_JoOf returns a Season_Uc_Jo based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Jo is returned.
func Season_Uc_JoOf(v int) Season_Uc_Jo {
	if 0 <= v && v < len(AllSeason_Uc_Jos) {
		return AllSeason_Uc_Jos[v]
	}
	// an invalid result
	return Spring_Uc_Jo + Summer_Uc_Jo + Autumn_Uc_Jo + Winter_Uc_Jo + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Jo) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Jo(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Uc_Jo, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Uc_Jo.
//
// Usage Example
//
//	v := new(Season_Uc_Jo)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Jo) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_joTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Uc_Jo) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_joEnumStrings, season_uc_joEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_jo")
}

// season_uc_joTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_joTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

func (v *Season_Uc_Jo) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Jos[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Uc_Jo parses a string to find the corresponding Season_Uc_Jo, accepting either one of the string values or
// a number. The input representation is determined by season_uc_joMarshalTextRep. It wraps Parse.
func AsSeason_Uc_Jo(s string) (Season_Uc_Jo, error) {
	var v = new(Season_Uc_Jo)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Jo is similar to AsSeason_Uc_Jo except that it panics on error.
func MustParseSeason_Uc_Jo(s string) Season_Uc_Jo {
	v, err := AsSeason_Uc_Jo(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The ordinal representation is chosen according to -marshaljson.
func (v Season_Uc_Jo) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, v.invalidError()
	}

	s := strconv.Itoa(o)
	return []byte(s), nil
}

func (v Season_Uc_Jo) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_jo", v)
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Uc_Jo) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

// parseOrdinal attempts to convert an ordinal value.
func (v *Season_Uc_Jo) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSeason_Uc_Jos) {
		*v = AllSeason_Uc_Jos[ord]
		return true
	}
	return false
}

func (v *Season_Uc_Jo) unmarshalJSON(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := season_uc_joTransformInput(in)

	return v.parseFallback(in, s)
}
