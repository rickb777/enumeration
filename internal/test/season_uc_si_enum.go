// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.0-dirty

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
	"slices"
	"strconv"
	"strings"
)

// AllSeason_Uc_Sis lists all 4 values in order.
var AllSeason_Uc_Sis = []Season_Uc_Si{
	Spring_Uc_Si, Summer_Uc_Si, Autumn_Uc_Si, Winter_Uc_Si,
}

// AllSeason_Uc_SiEnums lists all 4 values in order.
var AllSeason_Uc_SiEnums = enum.IntEnums{
	Spring_Uc_Si, Summer_Uc_Si, Autumn_Uc_Si, Winter_Uc_Si,
}

const (
	season_uc_siEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_siEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Si. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Si) Ordinal() int {
	switch v {
	case Spring_Uc_Si:
		return 0
	case Summer_Uc_Si:
		return 1
	case Autumn_Uc_Si:
		return 2
	case Winter_Uc_Si:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Si, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Si) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_siEnumStrings, season_uc_siEnumIndex[:])
}

func (v Season_Uc_Si) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Sis) {
		return fmt.Sprintf("Season_Uc_Si(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Si is one of the defined constants.
func (v Season_Uc_Si) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Si) Int() int {
	return int(v)
}

var invalidSeason_Uc_SiValue = func() Season_Uc_Si {
	var v Season_Uc_Si
	for {
		if !slices.Contains(AllSeason_Uc_Sis, v) {
			return v
		}
		v++
	} // AllSeason_Uc_Sis is a finite set so loop will terminate eventually
}()

// Season_Uc_SiOf returns a Season_Uc_Si based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Uc_Si is returned.
func Season_Uc_SiOf(v int) Season_Uc_Si {
	if 0 <= v && v < len(AllSeason_Uc_Sis) {
		return AllSeason_Uc_Sis[v]
	}
	return invalidSeason_Uc_SiValue
}

// Parse parses a string to find the corresponding Season_Uc_Si, accepting one of the string values or
// a number. It is used by AsSeason_Uc_Si.
//
// Usage Example
//
//	v := new(Season_Uc_Si)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Uc_Si) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_siTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Uc_Si) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Uc_Si(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Uc_Si) parseFallback(in, s string) error {
	if v.parseString(s, season_uc_siEnumStrings, season_uc_siEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_uc_si")
}

func (v *Season_Uc_Si) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Uc_Sis[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_uc_siTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_uc_siTransformInput = func(in string) string {
	return strings.ToUpper(in)
}

// AsSeason_Uc_Si parses a string to find the corresponding Season_Uc_Si, accepting either one of the string values or
// a number. It wraps Parse.
func AsSeason_Uc_Si(s string) (Season_Uc_Si, error) {
	var v = new(Season_Uc_Si)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Uc_Si is similar to AsSeason_Uc_Si except that it panics on error.
func MustParseSeason_Uc_Si(s string) Season_Uc_Si {
	v, err := AsSeason_Uc_Si(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Uc_Si) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Uc_Si(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Uc_Si(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_uc_si", value, value)
	}

	return v.scanParse(s)
}

func (v Season_Uc_Si) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Uc_Si) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_si", v)
}

func (v *Season_Uc_Si) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_uc_siTransformInput(in)

	return v.parseFallback(in, s)
}

// Value converts the Season_Uc_Si to a string  (based on '-store identifier').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Uc_Si) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.String(), nil
}
