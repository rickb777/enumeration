// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.0-dirty

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
	"slices"
	"strconv"
)

// AllSeason_Nc_Tis lists all 4 values in order.
var AllSeason_Nc_Tis = []Season_Nc_Ti{
	Spring_Nc_Ti, Summer_Nc_Ti, Autumn_Nc_Ti, Winter_Nc_Ti,
}

// AllSeason_Nc_TiEnums lists all 4 values in order.
var AllSeason_Nc_TiEnums = enum.IntEnums{
	Spring_Nc_Ti, Summer_Nc_Ti, Autumn_Nc_Ti, Winter_Nc_Ti,
}

const (
	season_nc_tiEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_tiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Ti. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ti) Ordinal() int {
	switch v {
	case Spring_Nc_Ti:
		return 0
	case Summer_Nc_Ti:
		return 1
	case Autumn_Nc_Ti:
		return 2
	case Winter_Nc_Ti:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Ti, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ti) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_tiEnumStrings, season_nc_tiEnumIndex[:])
}

func (v Season_Nc_Ti) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Tis) {
		return fmt.Sprintf("Season_Nc_Ti(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Ti is one of the defined constants.
func (v Season_Nc_Ti) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ti) Int() int {
	return int(v)
}

var invalidSeason_Nc_TiValue = func() Season_Nc_Ti {
	var v Season_Nc_Ti
	for {
		if !slices.Contains(AllSeason_Nc_Tis, v) {
			return v
		}
		v++
	} // AllSeason_Nc_Tis is a finite set so loop will terminate eventually
}()

// Season_Nc_TiOf returns a Season_Nc_Ti based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Ti is returned.
func Season_Nc_TiOf(v int) Season_Nc_Ti {
	if 0 <= v && v < len(AllSeason_Nc_Tis) {
		return AllSeason_Nc_Tis[v]
	}
	return invalidSeason_Nc_TiValue
}

// Parse parses a string to find the corresponding Season_Nc_Ti, accepting one of the string values or
// a number. The input representation is determined by Identifier. It is used by AsSeason_Nc_Ti.
//
// Usage Example
//
//	v := new(Season_Nc_Ti)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Ti) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_tiTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Ti) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Ti(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Nc_Ti) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_tiEnumStrings, season_nc_tiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_ti")
}

func (v *Season_Nc_Ti) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Tis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_nc_tiTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_tiTransformInput = func(in string) string {
	return in
}

// AsSeason_Nc_Ti parses a string to find the corresponding Season_Nc_Ti, accepting either one of the string values or
// a number. The input representation is determined by season_nc_tiMarshalTextRep. It wraps Parse.
func AsSeason_Nc_Ti(s string) (Season_Nc_Ti, error) {
	var v = new(Season_Nc_Ti)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Ti is similar to AsSeason_Nc_Ti except that it panics on error.
func MustParseSeason_Nc_Ti(s string) Season_Nc_Ti {
	v, err := AsSeason_Nc_Ti(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Nc_Ti) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Nc_Ti) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v Season_Nc_Ti) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_nc_tiEnumStrings, season_nc_tiEnumIndex[:]), nil
}

func (v Season_Nc_Ti) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Ti) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Ti) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ti", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Nc_Ti) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Nc_Ti) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_tiTransformInput(in)

	return v.parseFallback(in, s)
}
