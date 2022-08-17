// generated code - do not edit
// github.com/rickb777/enumeration/v3 v2.14.0

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
)

// AllSeason_Nc_Tas lists all 4 values in order.
var AllSeason_Nc_Tas = []Season_Nc_Ta{
	Spring_Nc_Ta, Summer_Nc_Ta, Autumn_Nc_Ta, Winter_Nc_Ta,
}

// AllSeason_Nc_TaEnums lists all 4 values in order.
var AllSeason_Nc_TaEnums = enum.IntEnums{
	Spring_Nc_Ta, Summer_Nc_Ta, Autumn_Nc_Ta, Winter_Nc_Ta,
}

const (
	season_nc_taEnumStrings = "SpringSummerAutumnWinter"
	season_nc_taTextStrings = "SprgSumrAutmWint"
	season_nc_taJSONStrings = "SprgSumrAutmWint"
	season_nc_taSQLStrings  = "SprgSumrAutmWint"
)

var (
	season_nc_taEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_nc_taTextIndex = [...]uint16{0, 4, 8, 12, 16}
	season_nc_taJSONIndex = [...]uint16{0, 4, 8, 12, 16}
	season_nc_taSQLIndex  = [...]uint16{0, 4, 8, 12, 16}
)

// String returns the literal string representation of a Season_Nc_Ta, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ta) String() string {
	return v.toString(season_nc_taEnumStrings, season_nc_taEnumIndex[:])
}

func (v Season_Nc_Ta) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllSeason_Nc_Tas) {
		return fmt.Sprintf("Season_Nc_Ta(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Nc_Ta. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ta) Ordinal() int {
	switch v {
	case Spring_Nc_Ta:
		return 0
	case Summer_Nc_Ta:
		return 1
	case Autumn_Nc_Ta:
		return 2
	case Winter_Nc_Ta:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Nc_Ta is one of the defined constants.
func (v Season_Nc_Ta) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ta) Int() int {
	return int(v)
}

// Season_Nc_TaOf returns a Season_Nc_Ta based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Ta is returned.
func Season_Nc_TaOf(v int) Season_Nc_Ta {
	if 0 <= v && v < len(AllSeason_Nc_Tas) {
		return AllSeason_Nc_Tas[v]
	}
	// an invalid result
	return Spring_Nc_Ta + Summer_Nc_Ta + Autumn_Nc_Ta + Winter_Nc_Ta + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Ta) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Ta(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Nc_Ta, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Nc_Ta.
//
// Usage Example
//
//	v := new(Season_Nc_Ta)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Ta) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_taTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Nc_Ta) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_taEnumStrings, season_nc_taEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_ta")
}

// season_nc_taTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_taTransformInput = func(in string) string {
	return in
}

func (v *Season_Nc_Ta) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Tas[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Nc_Ta parses a string to find the corresponding Season_Nc_Ta, accepting either one of the string values or
// a number. The input representation is determined by season_nc_taMarshalTextRep. It wraps Parse.
func AsSeason_Nc_Ta(s string) (Season_Nc_Ta, error) {
	var v = new(Season_Nc_Ta)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Ta is similar to AsSeason_Nc_Ta except that it panics on error.
func MustParseSeason_Nc_Ta(s string) Season_Nc_Ta {
	v, err := AsSeason_Nc_Ta(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Season_Nc_Ta) MarshalText() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}
	s := season_nc_taTextStrings[season_nc_taTextIndex[o]:season_nc_taTextIndex[o+1]]
	return []byte(s), nil
}

func (v Season_Nc_Ta) marshalNumberOrError() ([]byte, error) {
	return nil, v.invalidError()
}

func (v Season_Nc_Ta) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ta", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Season_Nc_Ta) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}
	s := season_nc_taJSONStrings[season_nc_taJSONIndex[o]:season_nc_taJSONIndex[o+1]]
	return enum.QuotedString(s), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Nc_Ta) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Nc_Ta) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_taTransformInput(in)

	if v.parseString(s, season_nc_taTextStrings, season_nc_taTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

func (v *Season_Nc_Ta) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_taTransformInput(in)

	if v.parseString(s, season_nc_taJSONStrings, season_nc_taJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, season_nc_taEnumStrings, season_nc_taEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_ta")
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Nc_Ta) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Nc_Ta(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Nc_Ta(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_nc_ta", value, value)
	}

	return v.scanParse(s)
}

func (v *Season_Nc_Ta) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_taTransformInput(in)

	if v.parseString(s, season_nc_taSQLStrings, season_nc_taSQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

func (v Season_Nc_Ta) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

// Value converts the Season_Nc_Ta to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Nc_Ta) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(season_nc_taSQLStrings, season_nc_taSQLIndex[:]), nil
}
