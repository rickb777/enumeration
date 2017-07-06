// generated code - do not edit

package example

import (
	"errors"
	"fmt"
)

const monthEnumStrings = "JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember"

var monthEnumIndex = [...]uint16{0, 7, 15, 20, 25, 28, 32, 36, 42, 51, 58, 66, 74}

var AllMonths = []Month{January, February, March, April, May, June, July, August, September, October, November, December}

// String returns the string representation of a Month
func (i Month) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(monthEnumIndex)-1 {
		return fmt.Sprintf("Month(%v)", i)
	}
	return monthEnumStrings[monthEnumIndex[o]:monthEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Month
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

// Parse parses a string to find the corresponding Month
func (v *Month) Parse(s string) error {
	var i0 uint16 = 0
	for j := 1; j < len(monthEnumIndex); j++ {
		i1 := monthEnumIndex[j]
		p := monthEnumStrings[i0:i1]
		if s == p {
			*v = Month(j - 1)
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Month")
}

// AsMonth parses a string to find the corresponding Month
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
