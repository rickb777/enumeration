// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.2.1

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
)

// AllSeason_Nc_Tts lists all 4 values in order.
var AllSeason_Nc_Tts = []Season_Nc_Tt{
	Spring_Nc_Tt, Summer_Nc_Tt, Autumn_Nc_Tt, Winter_Nc_Tt,
}

// AllSeason_Nc_TtEnums lists all 4 values in order.
var AllSeason_Nc_TtEnums = enum.IntEnums{
	Spring_Nc_Tt, Summer_Nc_Tt, Autumn_Nc_Tt, Winter_Nc_Tt,
}

const (
	season_nc_ttEnumStrings = "SpringSummerAutumnWinter"
	season_nc_ttTextStrings = "SprgSumrAutmWint"
)

var (
	season_nc_ttEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_nc_ttTextIndex = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Nc_Tt. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Tt) Ordinal() int {
	switch v {
	case Spring_Nc_Tt:
		return 0
	case Summer_Nc_Tt:
		return 1
	case Autumn_Nc_Tt:
		return 2
	case Winter_Nc_Tt:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Tt, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Tt) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_ttEnumStrings, season_nc_ttEnumIndex[:])
}

func (v Season_Nc_Tt) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Tts) {
		return fmt.Sprintf("Season_Nc_Tt(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Tt is one of the defined constants.
func (v Season_Nc_Tt) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Tt) Int() int {
	return int(v)
}

var invalidSeason_Nc_TtValue = func() Season_Nc_Tt {
	var v Season_Nc_Tt
	for {
		if !slices.Contains(AllSeason_Nc_Tts, v) {
			return v
		}
		v++
	} // AllSeason_Nc_Tts is a finite set so loop will terminate eventually
}()

// Season_Nc_TtOf returns a Season_Nc_Tt based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Tt is returned.
func Season_Nc_TtOf(v int) Season_Nc_Tt {
	if 0 <= v && v < len(AllSeason_Nc_Tts) {
		return AllSeason_Nc_Tts[v]
	}
	return invalidSeason_Nc_TtValue
}

// Parse parses a string to find the corresponding Season_Nc_Tt, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Nc_Tt.
//
// Usage Example
//
//	v := new(Season_Nc_Tt)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Tt) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_ttTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Tt) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Tt(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Nc_Tt) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_ttEnumStrings, season_nc_ttEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_tt")
}

func (v *Season_Nc_Tt) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Tts[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_nc_ttTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_ttTransformInput = func(in string) string {
	return in
}

// AsSeason_Nc_Tt parses a string to find the corresponding Season_Nc_Tt, accepting either one of the string values or
// a number. The input representation is determined by season_nc_ttMarshalTextRep. It wraps Parse.
func AsSeason_Nc_Tt(s string) (Season_Nc_Tt, error) {
	var v = new(Season_Nc_Tt)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Tt is similar to AsSeason_Nc_Tt except that it panics on error.
func MustParseSeason_Nc_Tt(s string) Season_Nc_Tt {
	v, err := AsSeason_Nc_Tt(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Nc_Tt) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Nc_Tt) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Season_Nc_Tt) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_nc_ttTextStrings, season_nc_ttTextIndex[:]), nil
}

func (v Season_Nc_Tt) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Tt) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Tt) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_tt", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Nc_Tt) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Nc_Tt) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_ttTransformInput(in)

	if v.parseString(s, season_nc_ttTextStrings, season_nc_ttTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}
