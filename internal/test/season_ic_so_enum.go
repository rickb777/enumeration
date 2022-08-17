// generated code - do not edit
// github.com/rickb777/enumeration/v3 v2.14.0

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
	"strings"
)

// AllSeason_Ic_Sos lists all 4 values in order.
var AllSeason_Ic_Sos = []Season_Ic_So{
	Spring_Ic_So, Summer_Ic_So, Autumn_Ic_So, Winter_Ic_So,
}

// AllSeason_Ic_SoEnums lists all 4 values in order.
var AllSeason_Ic_SoEnums = enum.IntEnums{
	Spring_Ic_So, Summer_Ic_So, Autumn_Ic_So, Winter_Ic_So,
}

const (
	season_ic_soEnumStrings = "SpringSummerAutumnWinter"
	season_ic_soEnumInputs  = "springsummerautumnwinter"
)

var (
	season_ic_soEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Ic_So, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_So) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_soEnumStrings, season_ic_soEnumIndex[:])
}

func (v Season_Ic_So) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Sos) {
		return fmt.Sprintf("Season_Ic_So(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Ic_So. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_So) Ordinal() int {
	switch v {
	case Spring_Ic_So:
		return 0
	case Summer_Ic_So:
		return 1
	case Autumn_Ic_So:
		return 2
	case Winter_Ic_So:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Ic_So is one of the defined constants.
func (v Season_Ic_So) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_So) Int() int {
	return int(v)
}

// Season_Ic_SoOf returns a Season_Ic_So based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_So is returned.
func Season_Ic_SoOf(v int) Season_Ic_So {
	if 0 <= v && v < len(AllSeason_Ic_Sos) {
		return AllSeason_Ic_Sos[v]
	}
	// an invalid result
	return Spring_Ic_So + Summer_Ic_So + Autumn_Ic_So + Winter_Ic_So + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_So) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_So(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Ic_So, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Ic_So.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_So)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_So) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_soTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Ic_So) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_soEnumInputs, season_ic_soEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_so")
}

// season_ic_soTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_soTransformInput = func(in string) string {
	return strings.ToLower(in)
}

func (v *Season_Ic_So) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Sos[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Ic_So parses a string to find the corresponding Season_Ic_So, accepting either one of the string values or
// a number. The input representation is determined by season_ic_soMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_So(s string) (Season_Ic_So, error) {
	var v = new(Season_Ic_So)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_So is similar to AsSeason_Ic_So except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_So(s string) Season_Ic_So {
	v, err := AsSeason_Ic_So(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Ic_So) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Ic_SoOf(int(x))
		return v.errorIfInvalid()
	case float64:
		*v = Season_Ic_SoOf(int(x))
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_ic_so", value, value)
	}

	return v.scanParse(s)
}

// parseOrdinal attempts to convert an ordinal value.
func (v *Season_Ic_So) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSeason_Ic_Sos) {
		*v = AllSeason_Ic_Sos[ord]
		return true
	}
	return false
}

func (v *Season_Ic_So) scanParse(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := season_ic_soTransformInput(in)

	return v.parseFallback(in, s)
}

func (v Season_Ic_So) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Ic_So) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_so", v)
}

// Value converts the Season_Ic_So to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Ic_So) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return int64(v.Ordinal()), nil
}
