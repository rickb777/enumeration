// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.7

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Uc_Sns lists all 4 values in order.
var AllSeason_Uc_Sns = []Season_Uc_Sn{
	Spring_Uc_Sn, Summer_Uc_Sn, Autumn_Uc_Sn, Winter_Uc_Sn,
}

const (
	season_uc_snEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_snEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Sn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Sn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_snEnumStrings, season_uc_snEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Uc_Sn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Sn) Ordinal() int {
	switch v {
	case Spring_Uc_Sn:
		return 0
	case Summer_Uc_Sn:
		return 1
	case Autumn_Uc_Sn:
		return 2
	case Winter_Uc_Sn:
		return 3
	}
	return -1
}

func (v Season_Uc_Sn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Sns) {
		return fmt.Sprintf("Season_Uc_Sn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Sn is one of the defined constants.
func (v Season_Uc_Sn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Sn) Int() int {
	return int(v)
}

var invalidSeason_Uc_SnValue = func() Season_Uc_Sn {
	var v Season_Uc_Sn
	for {
		if !slices.Contains(AllSeason_Uc_Sns, v) {
			return v
		}
		v++
	} // AllSeason_Uc_Sns is a finite set so loop will terminate eventually
}()

// Season_Uc_SnOf returns a Season_Uc_Sn based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Sn is returned.
func Season_Uc_SnOf(v int) Season_Uc_Sn {
	if 0 <= v && v < len(AllSeason_Uc_Sns) {
		return AllSeason_Uc_Sns[v]
	}
	return invalidSeason_Uc_SnValue
}

// Parse parses a string to find the corresponding Season_Uc_Sn, accepting one of the string values or
// a number. It is used by AsSeason_Uc_Sn.
//
// Usage Example
//
//	v := new(Season_Uc_Sn)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Sn) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_snTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Sn) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Sn(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Uc_Sn) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_snEnumStrings, season_uc_snEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_sn")
}

func (v *Season_Uc_Sn) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Sns[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_uc_snTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_snTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

// AsSeason_Uc_Sn parses a string to find the corresponding Season_Uc_Sn, accepting either one of the string values or
// a number. It wraps Parse.
func AsSeason_Uc_Sn(s string) (Season_Uc_Sn, error) {
	var v = new(Season_Uc_Sn)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Sn is similar to AsSeason_Uc_Sn except that it panics on error.
func MustParseSeason_Uc_Sn(s string) Season_Uc_Sn {
	v, err := AsSeason_Uc_Sn(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Uc_Sn) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Uc_Sn(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Uc_Sn(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_uc_sn", value, value)
	}

	return v.scanParse(s)
}

func (v Season_Uc_Sn) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Uc_Sn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_sn", v)
}

func (v *Season_Uc_Sn) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_snTransformInput(in)

	return v.parseFallback(in, s)
}

// Value converts the Season_Uc_Sn to a number (based on '-store number').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Uc_Sn) Value() (driver.Value, error) {
	return int64(v), nil
}
