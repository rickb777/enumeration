// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.4.0

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Uc_Tis lists all 4 values in order.
var AllSeason_Uc_Tis = []Season_Uc_Ti{
	Spring_Uc_Ti, Summer_Uc_Ti, Autumn_Uc_Ti, Winter_Uc_Ti,
}

// AllSeason_Uc_TiEnums lists all 4 values in order.
var AllSeason_Uc_TiEnums = enum.IntEnums{
	Spring_Uc_Ti, Summer_Uc_Ti, Autumn_Uc_Ti, Winter_Uc_Ti,
}

const (
	season_uc_tiEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_tiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Ti. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ti) Ordinal() int {
	switch v {
	case Spring_Uc_Ti:
		return 0
	case Summer_Uc_Ti:
		return 1
	case Autumn_Uc_Ti:
		return 2
	case Winter_Uc_Ti:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Ti, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ti) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_tiEnumStrings, season_uc_tiEnumIndex[:])
}

func (v Season_Uc_Ti) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tis) {
		return fmt.Sprintf("Season_Uc_Ti(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Ti is one of the defined constants.
func (v Season_Uc_Ti) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ti) Int() int {
	return int(v)
}

var invalidSeason_Uc_TiValue = func() Season_Uc_Ti {
	var v Season_Uc_Ti
	for {
		if !slices.Contains(AllSeason_Uc_Tis, v) {
			return v
		}
		v++
	} // AllSeason_Uc_Tis is a finite set so loop will terminate eventually
}()

// Season_Uc_TiOf returns a Season_Uc_Ti based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Ti is returned.
func Season_Uc_TiOf(v int) Season_Uc_Ti {
	if 0 <= v && v < len(AllSeason_Uc_Tis) {
		return AllSeason_Uc_Tis[v]
	}
	return invalidSeason_Uc_TiValue
}

// Parse parses a string to find the corresponding Season_Uc_Ti, accepting one of the string values or
// a number. The input representation is determined by Identifier. It is used by AsSeason_Uc_Ti.
//
// Usage Example
//
//	v := new(Season_Uc_Ti)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Ti) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_tiTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Ti) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Ti(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Uc_Ti) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_tiEnumStrings, season_uc_tiEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_ti")
}

func (v *Season_Uc_Ti) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Tis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_uc_tiTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_tiTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

// AsSeason_Uc_Ti parses a string to find the corresponding Season_Uc_Ti, accepting either one of the string values or
// a number. The input representation is determined by season_uc_tiMarshalTextRep. It wraps Parse.
func AsSeason_Uc_Ti(s string) (Season_Uc_Ti, error) {
	var v = new(Season_Uc_Ti)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Ti is similar to AsSeason_Uc_Ti except that it panics on error.
func MustParseSeason_Uc_Ti(s string) Season_Uc_Ti {
	v, err := AsSeason_Uc_Ti(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Uc_Ti) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Uc_Ti) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v Season_Uc_Ti) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_uc_tiEnumStrings, season_uc_tiEnumIndex[:]), nil
}

func (v Season_Uc_Ti) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Ti) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Ti) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_ti", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Uc_Ti) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Uc_Ti) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_tiTransformInput(in)

	return v.parseFallback(in, s)
}
