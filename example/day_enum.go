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

const dayEnumStrings = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var dayEnumIndex = [...]uint16{0, 6, 12, 19, 28, 36, 42, 50}

// AllDays lists all 7 values in order.
var AllDays = []Day{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// AllDayEnums lists all 7 values in order.
var AllDayEnums = []enum.Enum{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// String returns the string representation of a Day.
func (i Day) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllDays) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return dayEnumStrings[dayEnumIndex[o]:dayEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Day.
func (i Day) Ordinal() int {
	switch i {
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

// DayOf returns a Day based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Day is returned.
func DayOf(i int) Day {
	if 0 <= i && i < len(AllDays) {
		return AllDays[i]
	}
	// an invalid result
	return Sunday + Monday + Tuesday + Wednesday + Thursday + Friday + Saturday
}

// IsValid determines whether a Day is one of the defined constants.
func (i Day) IsValid() bool {
	switch i {
	case Sunday, Monday, Tuesday, Wednesday,
		Thursday, Friday, Saturday:
		return true
	}
	return false
}

// Parse parses a string to find the corresponding Day, accepting either one of the string
// values or an ordinal number.
func (v *Day) Parse(s string) error {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllDays) {
		*v = AllDays[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(dayEnumIndex); j++ {
		i1 := dayEnumIndex[j]
		p := dayEnumStrings[i0:i1]
		if s == p {
			*v = AllDays[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Day")
}

// AsDay parses a string to find the corresponding Day, accepting either one of the string
// values or an ordinal number.
func AsDay(s string) (Day, error) {
	var i = new(Day)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Day) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Day) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// DayMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var DayMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// DayMarshalJSONUsingString is true.
func (i Day) MarshalJSON() ([]byte, error) {
	if DayMarshalJSONUsingString {
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
func (i *Day) UnmarshalJSON(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
