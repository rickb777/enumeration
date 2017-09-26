// generated code - do not edit

package example

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const dayEnumStrings = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var dayEnumIndex = [...]uint16{0, 6, 12, 19, 28, 36, 42, 50}

var AllDays = []Day{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// String returns the string representation of a Day
func (i Day) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllDays) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return dayEnumStrings[dayEnumIndex[o]:dayEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Day
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

// MarshalJSON converts values to ordinals suitable for transmission via JSON.
func (i Day) MarshalJSON() ([]byte, error) {
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
