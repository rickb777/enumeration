// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.4

package test

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
	"strings"
)

// AllSeason_Uc_Jns lists all 4 values in order.
var AllSeason_Uc_Jns = []Season_Uc_Jn{
	Spring_Uc_Jn, Summer_Uc_Jn, Autumn_Uc_Jn, Winter_Uc_Jn,
}

// AllSeason_Uc_JnEnums lists all 4 values in order.
var AllSeason_Uc_JnEnums = enum.IntEnums{
	Spring_Uc_Jn, Summer_Uc_Jn, Autumn_Uc_Jn, Winter_Uc_Jn,
}

const (
	season_uc_jnEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_jnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Jn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Jn) Ordinal() int {
	switch v {
	case Spring_Uc_Jn:
		return 0
	case Summer_Uc_Jn:
		return 1
	case Autumn_Uc_Jn:
		return 2
	case Winter_Uc_Jn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Jn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Jn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jnEnumStrings, season_uc_jnEnumIndex[:])
}

func (v Season_Uc_Jn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jns) {
		return fmt.Sprintf("Season_Uc_Jn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Jn is one of the defined constants.
func (v Season_Uc_Jn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Jn) Int() int {
	return int(v)
}

// Season_Uc_JnOf returns a Season_Uc_Jn based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Jn is returned.
func Season_Uc_JnOf(v int) Season_Uc_Jn {
	if 0 <= v && v < len(AllSeason_Uc_Jns) {
		return AllSeason_Uc_Jns[v]
	}
	// an invalid result
	return Spring_Uc_Jn + Summer_Uc_Jn + Autumn_Uc_Jn + Winter_Uc_Jn + 1
}

// Parse parses a string to find the corresponding Season_Uc_Jn, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Uc_Jn.
//
// Usage Example
//
//	v := new(Season_Uc_Jn)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Jn) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_jnTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Jn) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Jn(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Uc_Jn) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_jnEnumStrings, season_uc_jnEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_jn")
}

func (v *Season_Uc_Jn) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Jns[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_uc_jnTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_jnTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

// AsSeason_Uc_Jn parses a string to find the corresponding Season_Uc_Jn, accepting either one of the string values or
// a number. The input representation is determined by season_uc_jnMarshalTextRep. It wraps Parse.
func AsSeason_Uc_Jn(s string) (Season_Uc_Jn, error) {
	var v = new(Season_Uc_Jn)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Jn is similar to AsSeason_Uc_Jn except that it panics on error.
func MustParseSeason_Uc_Jn(s string) Season_Uc_Jn {
	v, err := AsSeason_Uc_Jn(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The number representation is chosen according to -marshaljson.
func (v Season_Uc_Jn) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	s := season_uc_jnMarshalNumber(v)
	return []byte(s), nil
}

func (v Season_Uc_Jn) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Jn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_jn", v)
}

// season_uc_jnMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_uc_jnMarshalNumber = func(v Season_Uc_Jn) string {
	return strconv.FormatInt(int64(v), 10)
}

func (v Season_Uc_Jn) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Uc_Jn) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Uc_Jn) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_jnTransformInput(in)

	if v.parseString(s, season_uc_jnEnumStrings, season_uc_jnEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_jn")
}
