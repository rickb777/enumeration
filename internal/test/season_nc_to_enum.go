// generated code - do not edit
// github.com/rickb777/enumeration/v3 v2.14.0

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
)

// AllSeason_Nc_Tos lists all 4 values in order.
var AllSeason_Nc_Tos = []Season_Nc_To{
	Spring_Nc_To, Summer_Nc_To, Autumn_Nc_To, Winter_Nc_To,
}

// AllSeason_Nc_ToEnums lists all 4 values in order.
var AllSeason_Nc_ToEnums = enum.IntEnums{
	Spring_Nc_To, Summer_Nc_To, Autumn_Nc_To, Winter_Nc_To,
}

const (
	season_nc_toEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_toEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Nc_To, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_To) String() string {
	return v.toString(season_nc_toEnumStrings, season_nc_toEnumIndex[:])
}

func (v Season_Nc_To) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllSeason_Nc_Tos) {
		return fmt.Sprintf("Season_Nc_To(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Nc_To. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_To) Ordinal() int {
	switch v {
	case Spring_Nc_To:
		return 0
	case Summer_Nc_To:
		return 1
	case Autumn_Nc_To:
		return 2
	case Winter_Nc_To:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Nc_To is one of the defined constants.
func (v Season_Nc_To) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_To) Int() int {
	return int(v)
}

// Season_Nc_ToOf returns a Season_Nc_To based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_To is returned.
func Season_Nc_ToOf(v int) Season_Nc_To {
	if 0 <= v && v < len(AllSeason_Nc_Tos) {
		return AllSeason_Nc_Tos[v]
	}
	// an invalid result
	return Spring_Nc_To + Summer_Nc_To + Autumn_Nc_To + Winter_Nc_To + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_To) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_To(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Nc_To, accepting one of the string values or
// a number. The input representation is determined by Ordinal. It is used by AsSeason_Nc_To.
//
// Usage Example
//
//	v := new(Season_Nc_To)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_To) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_toTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Nc_To) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_toEnumStrings, season_nc_toEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_to")
}

// season_nc_toTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_toTransformInput = func(in string) string {
	return in
}

func (v *Season_Nc_To) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Tos[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Nc_To parses a string to find the corresponding Season_Nc_To, accepting either one of the string values or
// a number. The input representation is determined by season_nc_toMarshalTextRep. It wraps Parse.
func AsSeason_Nc_To(s string) (Season_Nc_To, error) {
	var v = new(Season_Nc_To)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_To is similar to AsSeason_Nc_To except that it panics on error.
func MustParseSeason_Nc_To(s string) Season_Nc_To {
	v, err := AsSeason_Nc_To(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to a form suitable for transmission via XML, JSON etc.
// The ordinal representation is chosen according to -marshaltext.
func (v Season_Nc_To) MarshalText() (text []byte, err error) {
	if !v.IsValid() {
		return nil, v.invalidError()
	}

	return v.marshalOrdinal()
}

func (v Season_Nc_To) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_to", v)
}

func (v Season_Nc_To) marshalOrdinal() (text []byte, err error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, v.invalidError()
	}
	return []byte(strconv.Itoa(o)), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Nc_To) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

// parseOrdinal attempts to convert an ordinal value.
func (v *Season_Nc_To) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSeason_Nc_Tos) {
		*v = AllSeason_Nc_Tos[ord]
		return true
	}
	return false
}

func (v *Season_Nc_To) unmarshalText(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := season_nc_toTransformInput(in)

	return v.parseFallback(in, s)
}
