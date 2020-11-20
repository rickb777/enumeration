// generated code - do not edit
// github.com/rickb777/enumeration v1.6.0

package example

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/enum"
	"strconv"
	"strings"
)

const monthEnumStrings = "JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember"

var monthEnumIndex = [...]uint16{0, 7, 15, 20, 25, 28, 32, 36, 42, 51, 58, 66, 74}

// AllMonths lists all 12 values in order.
var AllMonths = []Month{January, February, March, April, May, June, July, August, September, October, November, December}

// AllMonthEnums lists all 12 values in order.
var AllMonthEnums = enum.Enums{January, February, March, April, May, June, July, August, September, October, November, December}

// String returns the string representation of a Month.
func (i Month) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllMonths) {
		return fmt.Sprintf("Month(%d)", i)
	}
	return monthEnumStrings[monthEnumIndex[o]:monthEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Month.
func (i Month) Ordinal() int {
	switch i {
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

// MonthOf returns a Month based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Month is returned.
func MonthOf(i int) Month {
	if 0 <= i && i < len(AllMonths) {
		return AllMonths[i]
	}
	// an invalid result
	return January + February + March + April + May + June + July + August + September + October + November + December
}

// IsValid determines whether a Month is one of the defined constants.
func (i Month) IsValid() bool {
	switch i {
	case January, February, March, April,
		May, June, July, August, September,
		October, November, December:
		return true
	}
	return false
}

// Parse parses a string to find the corresponding Month, accepting either one of the string
// values or an ordinal number.
func (v *Month) Parse(s string) error {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllMonths) {
		*v = AllMonths[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(monthEnumIndex); j++ {
		i1 := monthEnumIndex[j]
		p := monthEnumStrings[i0:i1]
		if s == p {
			*v = AllMonths[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Month")
}

// AsMonth parses a string to find the corresponding Month, accepting either one of the string
// values or an ordinal number.
func AsMonth(s string) (Month, error) {
	var i = new(Month)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Month) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Month) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// MonthMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var MonthMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// MonthMarshalJSONUsingString is true.
func (i Month) MarshalJSON() ([]byte, error) {
	if MonthMarshalJSONUsingString {
		s := []byte(i.String())
		b := make([]byte, len(s)+2)
		b[0] = '"'
		copy(b[1:], s)
		b[len(s)+1] = '"'
		return b, nil
	}
	// else use the ordinal
	s := strconv.Itoa(i.Ordinal())
	return []byte(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Month) UnmarshalJSON(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
