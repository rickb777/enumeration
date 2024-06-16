// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.3.0

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

// AllSeason_Ic_Sss lists all 4 values in order.
var AllSeason_Ic_Sss = []Season_Ic_Ss{
	Spring_Ic_Ss, Summer_Ic_Ss, Autumn_Ic_Ss, Winter_Ic_Ss,
}

// AllSeason_Ic_SsEnums lists all 4 values in order.
var AllSeason_Ic_SsEnums = enum.IntEnums{
	Spring_Ic_Ss, Summer_Ic_Ss, Autumn_Ic_Ss, Winter_Ic_Ss,
}

const (
	season_ic_ssEnumStrings = "SpringSummerAutumnWinter"
	season_ic_ssEnumInputs  = "springsummerautumnwinter"
	season_ic_ssSQLStrings  = "SprgSumrAutmWint"
	season_ic_ssSQLInputs   = "sprgsumrautmwint"
)

var (
	season_ic_ssEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_ic_ssSQLIndex  = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Ic_Ss. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Ic_Ss) Ordinal() int {
	switch v {
	case Spring_Ic_Ss:
		return 0
	case Summer_Ic_Ss:
		return 1
	case Autumn_Ic_Ss:
		return 2
	case Winter_Ic_Ss:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Ic_Ss, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Ic_Ss) String() string {
	o := v.Ordinal()
	return v.toString(o, season_ic_ssEnumStrings, season_ic_ssEnumIndex[:])
}

func (v Season_Ic_Ss) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Ic_Sss) {
		return fmt.Sprintf("Season_Ic_Ss(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Ic_Ss is one of the defined constants.
func (v Season_Ic_Ss) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Ic_Ss) Int() int {
	return int(v)
}

var invalidSeason_Ic_SsValue = func() Season_Ic_Ss {
	var v Season_Ic_Ss
	for {
		if !slices.Contains(AllSeason_Ic_Sss, v) {
			return v
		}
		v++
	} // AllSeason_Ic_Sss is a finite set so loop will terminate eventually
}()

// Season_Ic_SsOf returns a Season_Ic_Ss based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Season_Ic_Ss is returned.
func Season_Ic_SsOf(v int) Season_Ic_Ss {
	if 0 <= v && v < len(AllSeason_Ic_Sss) {
		return AllSeason_Ic_Sss[v]
	}
	return invalidSeason_Ic_SsValue
}

// Parse parses a string to find the corresponding Season_Ic_Ss, accepting one of the string values or
// a number. It is used by AsSeason_Ic_Ss.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Season_Ic_Ss)
//	err := v.Parse(s)
//	...  etc
func (v *Season_Ic_Ss) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_ssTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Season_Ic_Ss) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Season_Ic_Ss(num)
		return v.IsValid()
	}
	return false
}

func (v *Season_Ic_Ss) parseFallback(in, s string) error {
	if v.parseString(s, season_ic_ssEnumInputs, season_ic_ssEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised season_ic_ss")
}

func (v *Season_Ic_Ss) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSeason_Ic_Sss[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// season_ic_ssTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var season_ic_ssTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsSeason_Ic_Ss parses a string to find the corresponding Season_Ic_Ss, accepting either one of the string values or
// a number. It wraps Parse.
// The input case does not matter.
func AsSeason_Ic_Ss(s string) (Season_Ic_Ss, error) {
	var v = new(Season_Ic_Ss)
	err := v.Parse(s)
	return *v, err
}

// MustParseSeason_Ic_Ss is similar to AsSeason_Ic_Ss except that it panics on error.
// The input case does not matter.
func MustParseSeason_Ic_Ss(s string) Season_Ic_Ss {
	v, err := AsSeason_Ic_Ss(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Season_Ic_Ss) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Season_Ic_Ss(x)
		return v.errorIfInvalid()
	case float64:
		*v = Season_Ic_Ss(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful season_ic_ss", value, value)
	}

	return v.scanParse(s)
}

func (v Season_Ic_Ss) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v Season_Ic_Ss) invalidError() error {
	return fmt.Errorf("%d is not a valid season_ic_ss", v)
}

func (v *Season_Ic_Ss) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := season_ic_ssTransformInput(in)

	if v.parseString(s, season_ic_ssSQLInputs, season_ic_ssSQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// Value converts the Season_Ic_Ss to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Ic_Ss) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, season_ic_ssSQLStrings, season_ic_ssSQLIndex[:]), nil
}
