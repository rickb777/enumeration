// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.14.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"strconv"
	"strings"
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

func (v Day) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllDays) {
		return fmt.Sprintf("Day(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
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

// Tag returns the string representation of a Day. This is an alias for String.
func (v Day) Tag() string {
	return v.String()
}

// String returns the literal string representation of a Day, which is
// the same as the const identifier but without prefix or suffix.
func (v Day) String() string {
	return v.toString(dayEnumStrings, dayEnumIndex[:])
}

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

// IsValid determines whether a Day is one of the defined constants.
func (v Day) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Day) Int() int {
	return int(v)
}

// DayOf returns a Day based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Day is returned.
func DayOf(v int) Day {
	if 0 <= v && v < len(AllDays) {
		return AllDays[v]
	}
	// an invalid result
	return Sunday + Monday + Tuesday + Wednesday + Thursday + Friday + Saturday + 1
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

	s := dayTransformInput(in)

	if v.parseString(s, dayEnumStrings, dayEnumIndex[:]) {
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

// dayMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
// By default, it is enum.Identifier.
// The initial value is set using the -marshaltext command line parameter.
var dayMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to dayMarshalTextRep.
func (v Day) MarshalText() (text []byte, err error) {
	return v.marshalText(dayMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to dayMarshalTextRep.
func (v Day) MarshalJSON() ([]byte, error) {
	return v.marshalText(dayMarshalTextRep, true)
}

func (v Day) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	if rep != enum.Ordinal && !v.IsValid() {
		return dayMarshalNumber(v)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return dayMarshalNumber(v)
	case enum.Ordinal:
		return v.marshalOrdinal()
	case enum.Tag:
		if quoted {
			bs = enum.QuotedString(v.Tag())
		} else {
			bs = []byte(v.Tag())
		}
	default:
		if quoted {
			bs = enum.QuotedString(v.String())
		} else {
			bs = []byte(v.String())
		}
	}
	return bs, nil
}

// dayMarshalNumber handles marshaling where a number is required or where
// the value is out of range but dayMarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var dayMarshalNumber = func(v Day) (text []byte, err error) {
	bs := []byte(strconv.FormatInt(int64(v), 10))
	return bs, nil
}

func (v Day) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(v.Ordinal()))
	return bs, nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Day) UnmarshalText(text []byte) error {
	return v.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Day) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Day) unmarshalJSON(s string) error {
	return v.Parse(s)
}

// dayStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier.
// The initial value is set using the -store command line parameter.
var dayStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Day) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		if dayStoreRep == enum.Ordinal {
			*v = DayOf(int(x))
		} else {
			*v = Day(x)
		}
		return nil
	case float64:
		*v = Day(x)
		return nil
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful day", value, value)
	}

	return v.parse(s, dayStoreRep)
}

// Value converts the Day to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Day) Value() (driver.Value, error) {
	switch dayStoreRep {
	case enum.Number:
		return int64(v), nil
	case enum.Ordinal:
		return int64(v.Ordinal()), nil
	case enum.Tag:
		return v.Tag(), nil
	default:
		return v.String(), nil
	}
}
