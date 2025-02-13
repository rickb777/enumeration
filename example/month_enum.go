// generated code - do not edit
// github.com/rickb777/enumeration/v4 6fbd6b0a14258861d58a8efc36602c830d2f5fce-dirty

package example

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
	"slices"
	"strconv"
	"strings"
)

// AllMonths lists all 12 values in order.
var AllMonths = []Month{
	January, February, March, April,
	May, June, July, August, September,
	October, November, December,
}

// AllMonthEnums lists all 12 values in order.
var AllMonthEnums = enum.IntEnums{
	January, February, March, April,
	May, June, July, August, September,
	October, November, December,
}

const (
	monthEnumStrings = "JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember"
	monthEnumInputs  = "januaryfebruarymarchaprilmayjunejulyaugustseptemberoctobernovemberdecember"
)

var (
	monthEnumIndex = [...]uint16{0, 7, 15, 20, 25, 28, 32, 36, 42, 51, 58, 66, 74}
)

// Ordinal returns the ordinal number of a Month. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Month) Ordinal() int {
	switch v {
	case January:
		return 0
	case February:
		return 1
	case March:
		return 2
	case April:
		return 3
	case May:
		return 4
	case June:
		return 5
	case July:
		return 6
	case August:
		return 7
	case September:
		return 8
	case October:
		return 9
	case November:
		return 10
	case December:
		return 11
	}
	return -1
}

// String returns the literal string representation of a Month, which is
// the same as the const identifier but without prefix or suffix.
func (v Month) String() string {
	o := v.Ordinal()
	return v.toString(o, monthEnumStrings, monthEnumIndex[:])
}

func (v Month) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllMonths) {
		return fmt.Sprintf("Month(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Month is one of the defined constants.
func (v Month) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Month) Int() int {
	return int(v)
}

var invalidMonthValue = func() Month {
	var v Month
	for {
		if !slices.Contains(AllMonths, v) {
			return v
		}
		v++
	} // AllMonths is a finite set so loop will terminate eventually
}()

// MonthOf returns a Month based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Month is returned.
func MonthOf(v int) Month {
	if 0 <= v && v < len(AllMonths) {
		return AllMonths[v]
	}
	return invalidMonthValue
}

// Parse parses a string to find the corresponding Month, accepting one of the string values or
// a number. The input representation is determined by Identifier. It is used by AsMonth.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Month)
//	err := v.Parse(s)
//	...  etc
func (v *Month) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := monthTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Month) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Month(num)
		return v.IsValid()
	}
	return false
}

func (v *Month) parseFallback(in, s string) error {
	if v.parseString(s, monthEnumInputs, monthEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised month")
}

func (v *Month) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllMonths[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// monthTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var monthTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsMonth parses a string to find the corresponding Month, accepting either one of the string values or
// a number. The input representation is determined by monthMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsMonth(s string) (Month, error) {
	var v = new(Month)
	err := v.Parse(s)
	return *v, err
}

// MustParseMonth is similar to AsMonth except that it panics on error.
// The input case does not matter.
func MustParseMonth(s string) Month {
	v, err := AsMonth(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Month) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Month) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v Month) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, monthEnumStrings, monthEnumIndex[:]), nil
}

func (v Month) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Month) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Month) invalidError() error {
	return fmt.Errorf("%d is not a valid month", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Month) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *Month) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := monthTransformInput(in)

	return v.parseFallback(in, s)
}
