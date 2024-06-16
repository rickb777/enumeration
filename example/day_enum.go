// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.2.1

package example

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
)

// AllDays lists all 7 values in order.
var AllDays = []Day{
	Sunday, Monday, Tuesday, Wednesday,
	Thursday, Friday, Saturday,
}

// AllDayEnums lists all 7 values in order.
var AllDayEnums = enum.IntEnums{
	Sunday, Monday, Tuesday, Wednesday,
	Thursday, Friday, Saturday,
}

const (
	dayEnumStrings = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"
)

var (
	dayEnumIndex = [...]uint16{0, 6, 12, 19, 28, 36, 42, 50}
)

// Ordinal returns the ordinal number of a Day. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Day) Ordinal() int {
	switch v {
	case Sunday:
		return 0
	case Monday:
		return 1
	case Tuesday:
		return 2
	case Wednesday:
		return 3
	case Thursday:
		return 4
	case Friday:
		return 5
	case Saturday:
		return 6
	}
	return -1
}

// String returns the literal string representation of a Day, which is
// the same as the const identifier but without prefix or suffix.
func (v Day) String() string {
	o := v.Ordinal()
	return v.toString(o, dayEnumStrings, dayEnumIndex[:])
}

func (v Day) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllDays) {
		return fmt.Sprintf("Day(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Day is one of the defined constants.
func (v Day) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Day) Int() int {
	return int(v)
}

var invalidDayValue = func() Day {
	var v Day
	for {
		if !slices.Contains(AllDays, v) {
			return v
		}
		v++
	} // AllDays is a finite set so loop will terminate eventually
}()

// DayOf returns a Day based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Day is returned.
func DayOf(v int) Day {
	if 0 <= v && v < len(AllDays) {
		return AllDays[v]
	}
	return invalidDayValue
}

// Parse parses a string to find the corresponding Day, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsDay.
//
// Usage Example
//
//	v := new(Day)
//	err := v.Parse(s)
//	...  etc
func (v *Day) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := dayTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Day) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Day(num)
		return v.IsValid()
	}
	return false
}

func (v *Day) parseFallback(in, s string) error {
	if v.parseString(s, dayEnumStrings, dayEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised day")
}

func (v *Day) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllDays[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// dayTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var dayTransformInput = func(in string) string {
	return in
}

// AsDay parses a string to find the corresponding Day, accepting either one of the string values or
// a number. The input representation is determined by dayMarshalTextRep. It wraps Parse.
func AsDay(s string) (Day, error) {
	var v = new(Day)
	err := v.Parse(s)
	return *v, err
}

// MustParseDay is similar to AsDay except that it panics on error.
func MustParseDay(s string) Day {
	v, err := AsDay(s)
	if err != nil {
		panic(err)
	}
	return v
}
