// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.1

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
	"strings"
)

// AllSeason_Uc_Tns lists all 4 values in order.
var AllSeason_Uc_Tns = []Season_Uc_Tn{
	Spring_Uc_Tn, Summer_Uc_Tn, Autumn_Uc_Tn, Winter_Uc_Tn,
}

// AllSeason_Uc_TnEnums lists all 4 values in order.
var AllSeason_Uc_TnEnums = enum.IntEnums{
	Spring_Uc_Tn, Summer_Uc_Tn, Autumn_Uc_Tn, Winter_Uc_Tn,
}

const (
	season_uc_tnEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_tnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Tn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Tn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_tnEnumStrings, season_uc_tnEnumIndex[:])
}

func (v Season_Uc_Tn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tns) {
		return fmt.Sprintf("Season_Uc_Tn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Uc_Tn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Tn) Ordinal() int {
	switch v {
	case Spring_Uc_Tn:
		return 0
	case Summer_Uc_Tn:
		return 1
	case Autumn_Uc_Tn:
		return 2
	case Winter_Uc_Tn:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Uc_Tn is one of the defined constants.
func (v Season_Uc_Tn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Tn) Int() int {
	return int(v)
}

// Season_Uc_TnOf returns a Season_Uc_Tn based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Tn is returned.
func Season_Uc_TnOf(v int) Season_Uc_Tn {
	if 0 <= v && v < len(AllSeason_Uc_Tns) {
		return AllSeason_Uc_Tns[v]
	}
	// an invalid result
	return Spring_Uc_Tn + Summer_Uc_Tn + Autumn_Uc_Tn + Winter_Uc_Tn + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Tn) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Tn(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Uc_Tn, accepting one of the string values or
// a number. The input representation is determined by Number. It is used by AsSeason_Uc_Tn.
//
// Usage Example
//
//    v := new(Season_Uc_Tn)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Season_Uc_Tn) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_tnTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Uc_Tn) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_tnEnumStrings, season_uc_tnEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_tn")
}

// season_uc_tnTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_tnTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

func (v *Season_Uc_Tn) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Tns[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Uc_Tn parses a string to find the corresponding Season_Uc_Tn, accepting either one of the string values or
// a number. The input representation is determined by season_uc_tnMarshalTextRep. It wraps Parse.
func AsSeason_Uc_Tn(s string) (Season_Uc_Tn, error) {
	var v = new(Season_Uc_Tn)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Tn is similar to AsSeason_Uc_Tn except that it panics on error.
func MustParseSeason_Uc_Tn(s string) Season_Uc_Tn {
	v, err := AsSeason_Uc_Tn(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Uc_Tn) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The number representation is chosen according to -marshaltext.
func (v Season_Uc_Tn) marshalText() (string, error) {
	if !v.IsValid() {
		return v.marshalNumberStringOrError()
	}

	return season_uc_tnMarshalNumber(v), nil
}

// season_uc_tnMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_uc_tnMarshalNumber = func(v Season_Uc_Tn) string {
	return strconv.FormatInt(int64(v), 10)
}

func (v Season_Uc_Tn) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Tn) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Tn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_tn", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Uc_Tn) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}


func (v *Season_Uc_Tn) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_tnTransformInput(in)

	return v.parseFallback(in, s)
}
