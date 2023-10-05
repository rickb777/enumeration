// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.2

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
	"strings"
)

// AllSeason_Ic_Tas lists all 4 values in order.
var AllSeason_Ic_Tas = []Season_Ic_Ta{
	Spring_Ic_Ta, Summer_Ic_Ta, Autumn_Ic_Ta, Winter_Ic_Ta,
}

// AllSeason_Ic_TaEnums lists all 4 values in order.
var AllSeason_Ic_TaEnums = enum.IntEnums{
	Spring_Ic_Ta, Summer_Ic_Ta, Autumn_Ic_Ta, Winter_Ic_Ta,
}

const (
	season_ic_taEnumStrings = "SpringSummerAutumnWinter"
	season_ic_taEnumInputs  = "springsummerautumnwinter"
	season_ic_taTextStrings = "SprgSumrAutmWint"
	season_ic_taTextInputs  = "sprgsumrautmwint"
	season_ic_taJSONStrings = "SprgSumrAutmWint"
	season_ic_taJSONInputs  = "SprgSumrAutmWint"
	season_ic_taSQLStrings  = "SprgSumrAutmWint"
	season_ic_taSQLInputs   = "sprgsumrautmwint"
)

var (
	season_ic_taEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_ic_taTextIndex = [...]uint16{0, 4, 8, 12, 16}
	season_ic_taJSONIndex = [...]uint16{0, 4, 8, 12, 16}
	season_ic_taSQLIndex  = [...]uint16{0, 4, 8, 12, 16}
)

// String returns the literal string representation of a Season_Ic_Ta, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Ta) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_taEnumStrings, season_ic_taEnumIndex[:])
}

func (v Season_Ic_Ta) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Tas) {
		return fmt.Sprintf("Season_Ic_Ta(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Ic_Ta. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Ta) Ordinal() int {
	switch v {
	case Spring_Ic_Ta:
		return 0
	case Summer_Ic_Ta:
		return 1
	case Autumn_Ic_Ta:
		return 2
	case Winter_Ic_Ta:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Ic_Ta is one of the defined constants.
func (v Season_Ic_Ta) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Ta) Int() int {
	return int(v)
}

// Season_Ic_TaOf returns a Season_Ic_Ta based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Ta is returned.
func Season_Ic_TaOf(v int) Season_Ic_Ta {
	if 0 <= v && v < len(AllSeason_Ic_Tas) {
		return AllSeason_Ic_Tas[v]
	}
	// an invalid result
	return Spring_Ic_Ta + Summer_Ic_Ta + Autumn_Ic_Ta + Winter_Ic_Ta + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Ta) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Ta(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Ic_Ta, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Ic_Ta.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_Ta)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_Ta) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_taTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Ic_Ta) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_taEnumInputs, season_ic_taEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_ta")
}

// season_ic_taTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_taTransformInput = func(in string) string {
	return strings.ToLower(in)
}

func (v *Season_Ic_Ta) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Tas[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Ic_Ta parses a string to find the corresponding Season_Ic_Ta, accepting either one of the string values or
// a number. The input representation is determined by season_ic_taMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Ta(s string) (Season_Ic_Ta, error) {
	var v = new(Season_Ic_Ta)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Ta is similar to AsSeason_Ic_Ta except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Ta(s string) Season_Ic_Ta {
	v, err := AsSeason_Ic_Ta(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Ic_Ta) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Ic_Ta) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Season_Ic_Ta) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_ic_taTextStrings, season_ic_taTextIndex[:]), nil
}

func (v Season_Ic_Ta) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Ic_Ta) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Ic_Ta) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_ta", v)
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Ic_Ta) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_ic_taJSONStrings, season_ic_taJSONIndex[:])
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Season_Ic_Ta) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_ic_taJSONStrings, season_ic_taJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Ic_Ta) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Ic_Ta) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_taTransformInput(in)

	if v.parseString(s, season_ic_taTextInputs, season_ic_taTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Ic_Ta) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Ic_Ta) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_taTransformInput(in)

	if v.parseString(s, season_ic_taJSONInputs, season_ic_taJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, season_ic_taEnumInputs, season_ic_taEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_ta")
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Ic_Ta) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Ic_Ta(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Ic_Ta(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_ic_ta", value, value)
	}

	return v.scanParse(s)
}

func (v *Season_Ic_Ta) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_taTransformInput(in)

	if v.parseString(s, season_ic_taSQLInputs, season_ic_taSQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

func (v Season_Ic_Ta) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

// Value converts the Season_Ic_Ta to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Ic_Ta) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, season_ic_taSQLStrings, season_ic_taSQLIndex[:]), nil
}
