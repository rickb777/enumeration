// generated code - do not edit

package example

import (
	"errors"
	"fmt"
)

const dayEnumStrings = "SundayMondayTuesdayWednesdayThursdayFridayPartydaynumberOfDays"

var dayEnumIndex = [...]uint16{0, 6, 12, 19, 28, 36, 42, 50, 62}

var AllDays = []Day{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Partyday, numberOfDays}

// String returns the string representation of a Day
func (i Day) String() string {
	if i < 0 || i >= Day(len(dayEnumIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return dayEnumStrings[dayEnumIndex[i]:dayEnumIndex[i+1]]
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
	case Partyday:
		return 6
	case numberOfDays:
		return 7
	}
	panic(fmt.Errorf("%d: unknown Day", i))
}

// Parse parses a string to find the corresponding Day
func (v *Day) Parse(s string) error {
	var i0 uint16 = 0
	for j := 1; j < len(dayEnumIndex); j++ {
		i1 := dayEnumIndex[j]
		p := dayEnumStrings[i0:i1]
		if s == p {
			*v = Day(j - 1)
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Day")
}

// AsDay parses a string to find the corresponding Day
func AsDay(s string) (Day, error) {
	var i = new(Day)
	err := i.Parse(s)
	return *i, err
}
