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

// AllSeason_Uc_Sos lists all 4 values in order.
var AllSeason_Uc_Sos = []Season_Uc_So{
	Spring_Uc_So, Summer_Uc_So, Autumn_Uc_So, Winter_Uc_So,
}

// AllSeason_Uc_SoEnums lists all 4 values in order.
var AllSeason_Uc_SoEnums = enum.IntEnums{
	Spring_Uc_So, Summer_Uc_So, Autumn_Uc_So, Winter_Uc_So,
}

const (
	season_uc_soEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_soEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_So, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_So) String() string {
	return v.toString(season_uc_soEnumStrings, season_uc_soEnumIndex[:])
}

func (v Season_Uc_So) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllSeason_Uc_Sos) {
		return fmt.Sprintf("Season_Uc_So(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Uc_So. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_So) Ordinal() int {
	switch v {
	case Spring_Uc_So:
		return 0
	case Summer_Uc_So:
		return 1
	case Autumn_Uc_So:
		return 2
	case Winter_Uc_So:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Uc_So is one of the defined constants.
func (v Season_Uc_So) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_So) Int() int {
	return int(v)
}

// Season_Uc_SoOf returns a Season_Uc_So based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_So is returned.
func Season_Uc_SoOf(v int) Season_Uc_So {
	if 0 <= v && v < len(AllSeason_Uc_Sos) {
		return AllSeason_Uc_Sos[v]
	}
	// an invalid result
	return Spring_Uc_So + Summer_Uc_So + Autumn_Uc_So + Winter_Uc_So + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_So) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_So(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Uc_So, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Uc_So.
//
// Usage Example
//
//	v := new(Season_Uc_So)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_So) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_soTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Uc_So) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_soEnumStrings, season_uc_soEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_so")
}

// season_uc_soTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_soTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

func (v *Season_Uc_So) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Sos[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Uc_So parses a string to find the corresponding Season_Uc_So, accepting either one of the string values or
// a number. The input representation is determined by season_uc_soMarshalTextRep. It wraps Parse.
func AsSeason_Uc_So(s string) (Season_Uc_So, error) {
	var v = new(Season_Uc_So)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_So is similar to AsSeason_Uc_So except that it panics on error.
func MustParseSeason_Uc_So(s string) Season_Uc_So {
	v, err := AsSeason_Uc_So(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Uc_So) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Uc_SoOf(int(x))
		return v.errorIfInvalid()
	case float64:
		*v = Season_Uc_SoOf(int(x))
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_uc_so", value, value)
	}

	return v.scanParse(s)
}

// parseOrdinal attempts to convert an ordinal value.
func (v *Season_Uc_So) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSeason_Uc_Sos) {
		*v = AllSeason_Uc_Sos[ord]
		return true
	}
	return false
}

func (v *Season_Uc_So) scanParse(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := season_uc_soTransformInput(in)

	return v.parseFallback(in, s)
}

func (v Season_Uc_So) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Uc_So) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_so", v)
}

// Value converts the Season_Uc_So to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Uc_So) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return int64(v.Ordinal()), nil
}
