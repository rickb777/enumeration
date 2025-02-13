// generated code - do not edit
// github.com/rickb777/enumeration/v4 6fbd6b0a14258861d58a8efc36602c830d2f5fce-dirty

package test

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
	"slices"
	"strconv"
)

// AllSeason_Nc_Sss lists all 4 values in order.
var AllSeason_Nc_Sss = []Season_Nc_Ss{
	Spring_Nc_Ss, Summer_Nc_Ss, Autumn_Nc_Ss, Winter_Nc_Ss,
}

// AllSeason_Nc_SsEnums lists all 4 values in order.
var AllSeason_Nc_SsEnums = enum.IntEnums{
	Spring_Nc_Ss, Summer_Nc_Ss, Autumn_Nc_Ss, Winter_Nc_Ss,
}

const (
	season_nc_ssEnumStrings = "SpringSummerAutumnWinter"
	season_nc_ssSQLStrings  = "SprgSumrAutmWint"
)

var (
	season_nc_ssEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_nc_ssSQLIndex  = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Nc_Ss. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ss) Ordinal() int {
	switch v {
	case Spring_Nc_Ss:
		return 0
	case Summer_Nc_Ss:
		return 1
	case Autumn_Nc_Ss:
		return 2
	case Winter_Nc_Ss:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Ss, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ss) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_ssEnumStrings, season_nc_ssEnumIndex[:])
}

func (v Season_Nc_Ss) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Sss) {
		return fmt.Sprintf("Season_Nc_Ss(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Ss is one of the defined constants.
func (v Season_Nc_Ss) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ss) Int() int {
	return int(v)
}

var invalidSeason_Nc_SsValue = func() Season_Nc_Ss {
	var v Season_Nc_Ss
	for {
		if !slices.Contains(AllSeason_Nc_Sss, v) {
			return v
		}
		v++
	} // AllSeason_Nc_Sss is a finite set so loop will terminate eventually
}()

// Season_Nc_SsOf returns a Season_Nc_Ss based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Nc_Ss is returned.
func Season_Nc_SsOf(v int) Season_Nc_Ss {
	if 0 <= v && v < len(AllSeason_Nc_Sss) {
		return AllSeason_Nc_Sss[v]
	}
	return invalidSeason_Nc_SsValue
}

// Parse parses a string to find the corresponding Season_Nc_Ss, accepting one of the string values or
// a number. It is used by AsSeason_Nc_Ss.
//
// Usage Example
//
//	v := new(Season_Nc_Ss)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Nc_Ss) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_ssTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Nc_Ss) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Nc_Ss(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Nc_Ss) parseFallback(in, s string) error {
	if v.parseString(s, season_nc_ssEnumStrings, season_nc_ssEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_nc_ss")
}

func (v *Season_Nc_Ss) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Nc_Sss[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_nc_ssTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_nc_ssTransformInput = func(in string) string {
	return in
}

// AsSeason_Nc_Ss parses a string to find the corresponding Season_Nc_Ss, accepting either one of the string values or
// a number. It wraps Parse.
func AsSeason_Nc_Ss(s string) (Season_Nc_Ss, error) {
	var v = new(Season_Nc_Ss)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Nc_Ss is similar to AsSeason_Nc_Ss except that it panics on error.
func MustParseSeason_Nc_Ss(s string) Season_Nc_Ss {
	v, err := AsSeason_Nc_Ss(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Nc_Ss) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Nc_Ss(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Nc_Ss(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_nc_ss", value, value)
	}

	return v.scanParse(s)
}

func (v Season_Nc_Ss) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Nc_Ss) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ss", v)
}

func (v *Season_Nc_Ss) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_nc_ssTransformInput(in)

	if v.parseString(s, season_nc_ssSQLStrings, season_nc_ssSQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// Value converts the Season_Nc_Ss to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Nc_Ss) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, season_nc_ssSQLStrings, season_nc_ssSQLIndex[:]), nil
}
