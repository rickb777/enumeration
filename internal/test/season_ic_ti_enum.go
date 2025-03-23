// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-dirty

package test

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Ic_Tis lists all 4 values in order.
var AllSeason_Ic_Tis = []Season_Ic_Ti{
	Spring_Ic_Ti, Summer_Ic_Ti, Autumn_Ic_Ti, Winter_Ic_Ti,
}

const (
	season_ic_tiEnumStrings = "SpringSummerAutumnWinter"
	season_ic_tiEnumInputs  = "springsummerautumnwinter"
)

var (
	season_ic_tiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Ic_Ti, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Ti) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_tiEnumStrings, season_ic_tiEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Ic_Ti. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Ti) Ordinal() int {
	switch v {
	case Spring_Ic_Ti:
		return 0
	case Summer_Ic_Ti:
		return 1
	case Autumn_Ic_Ti:
		return 2
	case Winter_Ic_Ti:
		return 3
	}
	return -1
}

func (v Season_Ic_Ti) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Tis) {
		return fmt.Sprintf("Season_Ic_Ti(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Ic_Ti is one of the defined constants.
func (v Season_Ic_Ti) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Ti) Int() int {
	return int(v)
}

var invalidSeason_Ic_TiValue = func() Season_Ic_Ti {
	var v Season_Ic_Ti
	for {
		if !slices.Contains(AllSeason_Ic_Tis, v) {
			return v
		}
		v++
	} // AllSeason_Ic_Tis is a finite set so loop will terminate eventually
}()

// Season_Ic_TiOf returns a Season_Ic_Ti based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Ti is returned.
func Season_Ic_TiOf(v int) Season_Ic_Ti {
	if 0 <= v && v < len(AllSeason_Ic_Tis) {
		return AllSeason_Ic_Tis[v]
	}
	return invalidSeason_Ic_TiValue
}

// Parse parses a string to find the corresponding Season_Ic_Ti, accepting one of the string values or
// a number. The input representation is determined by Identifier. It is used by AsSeason_Ic_Ti.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_Ti)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_Ti) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_tiTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Ti) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Ti(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Ic_Ti) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_tiEnumInputs, season_ic_tiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_ti")
}

func (v *Season_Ic_Ti) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Tis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_ic_tiTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_tiTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsSeason_Ic_Ti parses a string to find the corresponding Season_Ic_Ti, accepting either one of the string values or
// a number. The input representation is determined by season_ic_tiMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Ti(s string) (Season_Ic_Ti, error) {
	var v = new(Season_Ic_Ti)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Ti is similar to AsSeason_Ic_Ti except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Ti(s string) Season_Ic_Ti {
	v, err := AsSeason_Ic_Ti(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Ic_Ti) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Ic_Ti) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v Season_Ic_Ti) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_ic_tiEnumStrings, season_ic_tiEnumIndex[:]), nil
}

func (v Season_Ic_Ti) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Ic_Ti) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Ic_Ti) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_ti", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Ic_Ti) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Ic_Ti) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_tiTransformInput(in)

	return v.parseFallback(in, s)
}
