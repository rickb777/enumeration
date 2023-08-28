// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.1.1

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
)

// AllSeason_Nc_Sis lists all 4 values in order.
var AllSeason_Nc_Sis = []Season_Nc_Si{
	Spring_Nc_Si, Summer_Nc_Si, Autumn_Nc_Si, Winter_Nc_Si,
}

// AllSeason_Nc_SiEnums lists all 4 values in order.
var AllSeason_Nc_SiEnums = enum.IntEnums{
	Spring_Nc_Si, Summer_Nc_Si, Autumn_Nc_Si, Winter_Nc_Si,
}

const (
	season_nc_siEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_siEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Nc_Si, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Si) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_siEnumStrings, season_nc_siEnumIndex[:])
}

func (v Season_Nc_Si) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Sis) {
		return fmt.Sprintf("Season_Nc_Si(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a Season_Nc_Si. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Si) Ordinal() int {
	switch v {
	case Spring_Nc_Si:
		return 0
	case Summer_Nc_Si:
		return 1
	case Autumn_Nc_Si:
		return 2
	case Winter_Nc_Si:
		return 3
	}
	return -1
}

// IsValid determines whether a Season_Nc_Si is one of the defined constants.
func (v Season_Nc_Si) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Si) Int() int {
	return int(v)
}

// Season_Nc_SiOf returns a Season_Nc_Si based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Si is returned.
func Season_Nc_SiOf(v int) Season_Nc_Si {
	if 0 <= v && v < len(AllSeason_Nc_Sis) {
		return AllSeason_Nc_Sis[v]
	}
	// an invalid result
	return Spring_Nc_Si + Summer_Nc_Si + Autumn_Nc_Si + Winter_Nc_Si + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Si) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Si(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding Season_Nc_Si, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsSeason_Nc_Si.
//
// Usage Example
//
//    v := new(Season_Nc_Si)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Season_Nc_Si) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_siTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *Season_Nc_Si) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_siEnumStrings, season_nc_siEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_si")
}

// season_nc_siTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_siTransformInput = func(in string) string {
	return in
}

func (v *Season_Nc_Si) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Sis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSeason_Nc_Si parses a string to find the corresponding Season_Nc_Si, accepting either one of the string values or
// a number. The input representation is determined by season_nc_siMarshalTextRep. It wraps Parse.
func AsSeason_Nc_Si(s string) (Season_Nc_Si, error) {
	var v = new(Season_Nc_Si)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Si is similar to AsSeason_Nc_Si except that it panics on error.
func MustParseSeason_Nc_Si(s string) Season_Nc_Si {
	v, err := AsSeason_Nc_Si(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Nc_Si) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Nc_Si(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Nc_Si(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_nc_si", value, value)
	}

	return v.scanParse(s)
}

func (v *Season_Nc_Si) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_siTransformInput(in)

	return v.parseFallback(in, s)
}

func (v Season_Nc_Si) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Nc_Si) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_si", v)
}

// Value converts the Season_Nc_Si to a string  (based on '-store identifier').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Nc_Si) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.String(), nil
}
