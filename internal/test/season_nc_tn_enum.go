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

// AllSeason_Nc_Tns lists all 4 values in order.
var AllSeason_Nc_Tns = []Season_Nc_Tn{
	Spring_Nc_Tn, Summer_Nc_Tn, Autumn_Nc_Tn, Winter_Nc_Tn,
}

// AllSeason_Nc_TnEnums lists all 4 values in order.
var AllSeason_Nc_TnEnums = enum.IntEnums{
	Spring_Nc_Tn, Summer_Nc_Tn, Autumn_Nc_Tn, Winter_Nc_Tn,
}

const (
	season_nc_tnEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_tnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Tn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Tn) Ordinal() int {
	switch v {
	case Spring_Nc_Tn:
		return 0
	case Summer_Nc_Tn:
		return 1
	case Autumn_Nc_Tn:
		return 2
	case Winter_Nc_Tn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Tn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Tn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_tnEnumStrings, season_nc_tnEnumIndex[:])
}

func (v Season_Nc_Tn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Tns) {
		return fmt.Sprintf("Season_Nc_Tn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Tn is one of the defined constants.
func (v Season_Nc_Tn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Tn) Int() int {
	return int(v)
}

var invalidSeason_Nc_TnValue = func() Season_Nc_Tn {
	var v Season_Nc_Tn
	for {
		if !slices.Contains(AllSeason_Nc_Tns, v) {
			return v
		}
		v++
	} // AllSeason_Nc_Tns is a finite set so loop will terminate eventually
}()

// Season_Nc_TnOf returns a Season_Nc_Tn based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Tn is returned.
func Season_Nc_TnOf(v int) Season_Nc_Tn {
	if 0 <= v && v < len(AllSeason_Nc_Tns) {
		return AllSeason_Nc_Tns[v]
	}
	return invalidSeason_Nc_TnValue
}

// Parse parses a string to find the corresponding Season_Nc_Tn, accepting one of the string values or
// a number. The input representation is determined by Number. It is used by AsSeason_Nc_Tn.
//
// Usage Example
//
//	v := new(Season_Nc_Tn)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Tn) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_tnTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Tn) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Tn(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Nc_Tn) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_tnEnumStrings, season_nc_tnEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_tn")
}

func (v *Season_Nc_Tn) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Tns[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_nc_tnTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_tnTransformInput = func(in string) string {
	return in
}

// AsSeason_Nc_Tn parses a string to find the corresponding Season_Nc_Tn, accepting either one of the string values or
// a number. The input representation is determined by season_nc_tnMarshalTextRep. It wraps Parse.
func AsSeason_Nc_Tn(s string) (Season_Nc_Tn, error) {
	var v = new(Season_Nc_Tn)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Tn is similar to AsSeason_Nc_Tn except that it panics on error.
func MustParseSeason_Nc_Tn(s string) Season_Nc_Tn {
	v, err := AsSeason_Nc_Tn(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Nc_Tn) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The number representation is chosen according to -marshaltext.
func (v Season_Nc_Tn) marshalText() (string, error) {
	if !v.IsValid() {
		return v.marshalNumberStringOrError()
	}

	return season_nc_tnMarshalNumber(v), nil
}

func (v Season_Nc_Tn) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Tn) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Tn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_tn", v)
}

// season_nc_tnMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_nc_tnMarshalNumber = func(v Season_Nc_Tn) string {
	return strconv.FormatInt(int64(v), 10)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Nc_Tn) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Nc_Tn) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_tnTransformInput(in)

	return v.parseFallback(in, s)
}
