// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.0

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Uc_Tas lists all 4 values in order.
var AllSeason_Uc_Tas = []Season_Uc_Ta{
	Spring_Uc_Ta, Summer_Uc_Ta, Autumn_Uc_Ta, Winter_Uc_Ta,
}

// AllSeason_Uc_TaEnums lists all 4 values in order.
var AllSeason_Uc_TaEnums = enum.IntEnums{
	Spring_Uc_Ta, Summer_Uc_Ta, Autumn_Uc_Ta, Winter_Uc_Ta,
}

const (
	season_uc_taEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
	season_uc_taTextStrings = "SprgSumrAutmWint"
	season_uc_taJSONStrings = "SprgSumrAutmWint"
	season_uc_taSQLStrings  = "SprgSumrAutmWint"
)

var (
	season_uc_taEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_uc_taTextIndex = [...]uint16{0, 4, 8, 12, 16}
	season_uc_taJSONIndex = [...]uint16{0, 4, 8, 12, 16}
	season_uc_taSQLIndex  = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Uc_Ta. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ta) Ordinal() int {
	switch v {
	case Spring_Uc_Ta:
		return 0
	case Summer_Uc_Ta:
		return 1
	case Autumn_Uc_Ta:
		return 2
	case Winter_Uc_Ta:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Ta, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ta) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_taEnumStrings, season_uc_taEnumIndex[:])
}

func (v Season_Uc_Ta) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tas) {
		return fmt.Sprintf("Season_Uc_Ta(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Ta is one of the defined constants.
func (v Season_Uc_Ta) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ta) Int() int {
	return int(v)
}

var invalidSeason_Uc_TaValue = func() Season_Uc_Ta {
	var v Season_Uc_Ta
	for {
		if !slices.Contains(AllSeason_Uc_Tas, v) {
			return v
		}
		v++
	} // AllSeason_Uc_Tas is a finite set so loop will terminate eventually
}()

// Season_Uc_TaOf returns a Season_Uc_Ta based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Ta is returned.
func Season_Uc_TaOf(v int) Season_Uc_Ta {
	if 0 <= v && v < len(AllSeason_Uc_Tas) {
		return AllSeason_Uc_Tas[v]
	}
	return invalidSeason_Uc_TaValue
}

// Parse parses a string to find the corresponding Season_Uc_Ta, accepting one of the string values or
// a number. It is used by AsSeason_Uc_Ta.
//
// Usage Example
//
//	v := new(Season_Uc_Ta)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Ta) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_taTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Ta) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Ta(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Uc_Ta) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_taEnumStrings, season_uc_taEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_ta")
}

func (v *Season_Uc_Ta) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Tas[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_uc_taTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_taTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

// AsSeason_Uc_Ta parses a string to find the corresponding Season_Uc_Ta, accepting either one of the string values or
// a number. It wraps Parse.
func AsSeason_Uc_Ta(s string) (Season_Uc_Ta, error) {
	var v = new(Season_Uc_Ta)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Ta is similar to AsSeason_Uc_Ta except that it panics on error.
func MustParseSeason_Uc_Ta(s string) Season_Uc_Ta {
	v, err := AsSeason_Uc_Ta(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Uc_Ta) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Uc_Ta) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Season_Uc_Ta) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_uc_taTextStrings, season_uc_taTextIndex[:]), nil
}

func (v Season_Uc_Ta) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Ta) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Ta) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_ta", v)
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Uc_Ta) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_uc_taJSONStrings, season_uc_taJSONIndex[:])
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Season_Uc_Ta) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_uc_taJSONStrings, season_uc_taJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Season_Uc_Ta) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Season_Uc_Ta) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_taTransformInput(in)

	if v.parseString(s, season_uc_taTextStrings, season_uc_taTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Season_Uc_Ta) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Season_Uc_Ta) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_taTransformInput(in)

	if v.parseString(s, season_uc_taJSONStrings, season_uc_taJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, season_uc_taEnumStrings, season_uc_taEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_ta")
}

// season_uc_taMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_uc_taMarshalNumber = func(v Season_Uc_Ta) string {
	return strconv.FormatInt(int64(v), 10)
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Uc_Ta) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Uc_Ta(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Uc_Ta(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_uc_ta", value, value)
	}

	return v.scanParse(s)
}

func (v Season_Uc_Ta) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v *Season_Uc_Ta) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_taTransformInput(in)

	if v.parseString(s, season_uc_taSQLStrings, season_uc_taSQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// Value converts the Season_Uc_Ta to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Uc_Ta) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, season_uc_taSQLStrings, season_uc_taSQLIndex[:]), nil
}
