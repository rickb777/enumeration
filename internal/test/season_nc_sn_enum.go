// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.1

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
)

// AllSeason_Nc_Sns lists all 4 values in order.
var AllSeason_Nc_Sns = []Season_Nc_Sn{
	Spring_Nc_Sn, Summer_Nc_Sn, Autumn_Nc_Sn, Winter_Nc_Sn,
}

// AllSeason_Nc_SnEnums lists all 4 values in order.
var AllSeason_Nc_SnEnums = enum.IntEnums{
	Spring_Nc_Sn, Summer_Nc_Sn, Autumn_Nc_Sn, Winter_Nc_Sn,
}

const (
	season_nc_snEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_snEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Sn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Sn) Ordinal() int {
	switch v {
	case Spring_Nc_Sn:
		return 0
	case Summer_Nc_Sn:
		return 1
	case Autumn_Nc_Sn:
		return 2
	case Winter_Nc_Sn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Sn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Sn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_snEnumStrings, season_nc_snEnumIndex[:])
}

func (v Season_Nc_Sn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Sns) {
		return fmt.Sprintf("Season_Nc_Sn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Sn is one of the defined constants.
func (v Season_Nc_Sn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Sn) Int() int {
	return int(v)
}

var invalidSeason_Nc_SnValue = func() Season_Nc_Sn {
	var v Season_Nc_Sn
	for {
		if !slices.Contains(AllSeason_Nc_Sns, v) {
			return v
		}
		v++
	} // AllSeason_Nc_Sns is a finite set so loop will terminate eventually
}()

// Season_Nc_SnOf returns a Season_Nc_Sn based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Sn is returned.
func Season_Nc_SnOf(v int) Season_Nc_Sn {
	if 0 <= v && v < len(AllSeason_Nc_Sns) {
		return AllSeason_Nc_Sns[v]
	}
	return invalidSeason_Nc_SnValue
}

// Parse parses a string to find the corresponding Season_Nc_Sn, accepting one of the string values or
// a number. It is used by AsSeason_Nc_Sn.
//
// Usage Example
//
//	v := new(Season_Nc_Sn)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Sn) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_snTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Sn) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Sn(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Nc_Sn) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_snEnumStrings, season_nc_snEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_sn")
}

func (v *Season_Nc_Sn) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Sns[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_nc_snTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_snTransformInput = func(in string) string {
	return in
}

// AsSeason_Nc_Sn parses a string to find the corresponding Season_Nc_Sn, accepting either one of the string values or
// a number. It wraps Parse.
func AsSeason_Nc_Sn(s string) (Season_Nc_Sn, error) {
	var v = new(Season_Nc_Sn)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Sn is similar to AsSeason_Nc_Sn except that it panics on error.
func MustParseSeason_Nc_Sn(s string) Season_Nc_Sn {
	v, err := AsSeason_Nc_Sn(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Nc_Sn) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Nc_Sn(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Nc_Sn(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_nc_sn", value, value)
	}

	return v.scanParse(s)
}

func (v Season_Nc_Sn) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Nc_Sn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_sn", v)
}

func (v *Season_Nc_Sn) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_snTransformInput(in)

	return v.parseFallback(in, s)
}

// Value converts the Season_Nc_Sn to a number (based on '-store number').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Nc_Sn) Value() (driver.Value, error) {
	return int64(v), nil
}
