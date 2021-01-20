// generated code - do not edit
// github.com/rickb777/enumeration v2.0.0

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
var AllDays = []Day{
	Sunday, Monday, Tuesday, Wednesday,
	Thursday, Friday, Saturday,
}

// AllDayEnums lists all 7 values in order.
var AllDayEnums = enum.IntEnums{
	Sunday, Monday, Tuesday, Wednesday,
	Thursday, Friday, Saturday,
}

// String returns the literal string representation of a Day, which is
// the same as the const identifier.
func (i Day) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllDays) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return dayEnumStrings[dayEnumIndex[o]:dayEnumIndex[o+1]]
}

// Tag returns the string representation of a Day. This is an alias for String.
func (i Day) Tag() string {
	return i.String()
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

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Day) Int() int {
	return int(i)
}

// DayOf returns a Day based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Day is returned.
func DayOf(i int) Day {
	if 0 <= i && i < len(AllDays) {
		return AllDays[i]
	}
	// an invalid result
	return Sunday + Monday + Tuesday + Wednesday + Thursday + Friday + Saturday + 1
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

// Parse parses a string to find the corresponding Day, accepting one of the string
// values or a number.
func (v *Day) Parse(in string) error {
	if dayMarshalTextUsing == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := in

	if v.parseIdentifier(s) {
		return nil
	}

	return errors.New(in + ": unrecognised day")
}

// parseNumber attempts to convert a decimal value
func (v *Day) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Day(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Day) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllDays) {
		*v = AllDays[ord]
		return true
	}
	return false
}

// parseIdentifier attempts to match an identifier.
func (v *Day) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0
	for j := 1; j < len(dayEnumIndex); j++ {
		i1 := dayEnumIndex[j]
		p := dayEnumStrings[i0:i1]
		if s == p {
			*v = AllDays[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsDay parses a string to find the corresponding Day, accepting either one of the string
// values or an ordinal number.
func AsDay(s string) (Day, error) {
	var i = new(Day)
	err := i.Parse(s)
	return *i, err
}

// dayMarshalTextUsingLiteral controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var dayMarshalTextUsing = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to DayMarshalTextUsing.
func (i Day) MarshalText() (text []byte, err error) {
	var s string
	switch dayMarshalTextUsing {
	case enum.Number:
		s = strconv.FormatInt(int64(i), 10)
	case enum.Ordinal:
		s = strconv.Itoa(i.Ordinal())
	case enum.Tag:
		s = i.Tag()
	default:
		s = i.String()
	}
	return []byte(s), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Day) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to DayMarshalTextUsing.
func (i Day) MarshalJSON() ([]byte, error) {
	var s []byte
	switch dayMarshalTextUsing {
	case enum.Number:
		s = []byte(strconv.FormatInt(int64(i), 10))
	case enum.Ordinal:
		s = []byte(strconv.Itoa(i.Ordinal()))
	case enum.Tag:
		s = i.quotedString(i.Tag())
	default:
		s = i.quotedString(i.String())
	}
	return s, nil
}

func (i Day) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Day) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Day) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		*i = Day(v)
	case float64:
		*i = Day(v)
	case []byte:
		err = i.Parse(string(v))
	case string:
		err = i.Parse(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful Day", value, value)
	}

	return err
}

// -- copy this somewhere and uncomment it if you need DB storage to use strings --
// Value converts the Day to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i Day) Value() (driver.Value, error) {
//    return i.String(), nil
//}
