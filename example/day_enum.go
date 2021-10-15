// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.6.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
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
	return i.Ordinal() >= 0
}

// Parse parses a string to find the corresponding Day, accepting one of the string values or
// a number. The input representation is determined by dayMarshalTextRep. It is used by AsDay.
//
// Usage Example
//
//    v := new(Day)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Day) Parse(s string) error {
	return v.parse(s, dayMarshalTextRep)
}

func (v *Day) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
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

// AsDay parses a string to find the corresponding Day, accepting either one of the string values or
// a number. The input representation is determined by dayMarshalTextRep. It wraps Parse.
func AsDay(s string) (Day, error) {
	var i = new(Day)
	err := i.Parse(s)
	return *i, err
}

// dayMarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var dayMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to dayMarshalTextRep.
func (i Day) MarshalText() (text []byte, err error) {
	return i.marshalText(dayMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to dayMarshalTextRep.
func (i Day) MarshalJSON() ([]byte, error) {
	return i.marshalText(dayMarshalTextRep, true)
}

func (i Day) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	var bs []byte
	switch rep {
	case enum.Number:
		bs = []byte(strconv.FormatInt(int64(i), 10))
	case enum.Ordinal:
		bs = []byte(strconv.Itoa(i.Ordinal()))
	case enum.Tag:
		if quoted {
			bs = i.quotedString(i.Tag())
		} else {
			bs = []byte(i.Tag())
		}
	default:
		if quoted {
			bs = []byte(i.quotedString(i.String()))
		} else {
			bs = []byte(i.String())
		}
	}
	return bs, nil
}

func (i Day) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Day) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
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

// dayStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var dayStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Day) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if dayStoreRep == enum.Ordinal {
			*i = DayOf(int(v))
		} else {
			*i = Day(v)
		}
	case float64:
		*i = Day(v)
	case []byte:
		err = i.parse(string(v), dayStoreRep)
	case string:
		err = i.parse(v, dayStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful day", value, value)
	}

	return err
}

// Value converts the Day to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Day) Value() (driver.Value, error) {
	switch dayStoreRep {
	case enum.Number:
		return int64(i), nil
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.String(), nil
	}
}
