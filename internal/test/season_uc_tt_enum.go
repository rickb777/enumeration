// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.2.1

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Uc_Tts lists all 4 values in order.
var AllSeason_Uc_Tts = []Season_Uc_Tt{
	Spring_Uc_Tt, Summer_Uc_Tt, Autumn_Uc_Tt, Winter_Uc_Tt,
}

// AllSeason_Uc_TtEnums lists all 4 values in order.
var AllSeason_Uc_TtEnums = enum.IntEnums{
	Spring_Uc_Tt, Summer_Uc_Tt, Autumn_Uc_Tt, Winter_Uc_Tt,
}

const (
	season_uc_ttEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
	season_uc_ttTextStrings = "SprgSumrAutmWint"
)

var (
	season_uc_ttEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_uc_ttTextIndex = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Uc_Tt. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Tt) Ordinal() int {
	switch v {
	case Spring_Uc_Tt:
		return 0
	case Summer_Uc_Tt:
		return 1
	case Autumn_Uc_Tt:
		return 2
	case Winter_Uc_Tt:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Tt, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Tt) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_ttEnumStrings, season_uc_ttEnumIndex[:])
}

func (v Season_Uc_Tt) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tts) {
		return fmt.Sprintf("Season_Uc_Tt(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Tt is one of the defined constants.
func (v Season_Uc_Tt) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Tt) Int() int {
	return int(v)
}

var invalidSeason_Uc_TtValue = func() Season_Uc_Tt {
	var v Season_Uc_Tt
	for {
		if !slices.Contains(AllSeason_Uc_Tts, v) {
			return v
		}
		v++
	} // AllSeason_Uc_Tts is a finite set so loop will terminate eventually
}()

// Season_Uc_TtOf returns a Season_Uc_Tt based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Tt is returned.
func Season_Uc_TtOf(v int) Season_Uc_Tt {
	if 0 <= v && v < len(AllSeason_Uc_Tts) {
		return AllSeason_Uc_Tts[v]
	}
	return invalidSeason_Uc_TtValue
}

// Parse parses a string to find the corresponding Season_Uc_Tt, accepting one of the string values or
// a number. It is used by AsSeason_Uc_Tt.
//
// Usage Example
//
//	v := new(Season_Uc_Tt)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Tt) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_ttTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Tt) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Tt(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Uc_Tt) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_ttEnumStrings, season_uc_ttEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_tt")
}

func (v *Season_Uc_Tt) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Tts[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_uc_ttTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_ttTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

// AsSeason_Uc_Tt parses a string to find the corresponding Season_Uc_Tt, accepting either one of the string values or
// a number. It wraps Parse.
func AsSeason_Uc_Tt(s string) (Season_Uc_Tt, error) {
	var v = new(Season_Uc_Tt)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Tt is similar to AsSeason_Uc_Tt except that it panics on error.
func MustParseSeason_Uc_Tt(s string) Season_Uc_Tt {
	v, err := AsSeason_Uc_Tt(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Uc_Tt) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Uc_Tt) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Season_Uc_Tt) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_uc_ttTextStrings, season_uc_ttTextIndex[:]), nil
}

func (v Season_Uc_Tt) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Tt) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Tt) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_tt", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Uc_Tt) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Uc_Tt) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_ttTransformInput(in)

	if v.parseString(s, season_uc_ttTextStrings, season_uc_ttTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}
