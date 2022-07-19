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

func (v Month) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllMonths) {
		return fmt.Sprintf("Month(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
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

// Tag returns the string representation of a Month. This is an alias for String.
func (v Month) Tag() string {
	return v.String()
}

// String returns the literal string representation of a Month, which is
// the same as the const identifier but without prefix or suffix.
func (v Month) String() string {
	return v.toString(monthEnumStrings, monthEnumIndex[:])
}

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

// IsValid determines whether a Month is one of the defined constants.
func (v Month) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Month) Int() int {
	return int(v)
}

// MonthOf returns a Month based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Month is returned.
func MonthOf(v int) Month {
	if 0 <= v && v < len(AllMonths) {
		return AllMonths[v]
	}
	// an invalid result
	return January + February + March + April + May + June + July + August + September + October + November + December + 1
}

// Parse parses a string to find the corresponding Month, accepting one of the string values or
// a number. The input representation is determined by monthMarshalTextRep. It is used by AsMonth.
// The input case does not matter.
//
// Usage Example
//
//    v := new(Month)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Month) Parse(s string) error {
	return v.parse(s, monthMarshalTextRep)
}

func (v *Month) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := monthTransformInput(in)

	if v.parseString(s, monthEnumInputs, monthEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised month")
}

// parseNumber attempts to convert a decimal value
func (v *Month) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Month(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Month) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllMonths) {
		*v = AllMonths[ord]
		return true
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

// monthMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
var monthMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to monthMarshalTextRep.
func (v Month) MarshalText() (text []byte, err error) {
	return v.marshalText(monthMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to monthMarshalTextRep.
func (v Month) MarshalJSON() ([]byte, error) {
	return v.marshalText(monthMarshalTextRep, true)
}

func (v Month) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	if rep != enum.Ordinal && !v.IsValid() {
		return monthMarshalNumber(v)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return monthMarshalNumber(v)
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

// monthMarshalNumber handles marshaling where a number is required or where
// the value is out of range but monthMarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var monthMarshalNumber = func(v Month) (text []byte, err error) {
	bs := []byte(strconv.FormatInt(int64(v), 10))
	return bs, nil
}

func (v Month) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(v.Ordinal()))
	return bs, nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Month) UnmarshalText(text []byte) error {
	return v.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Month) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Month) unmarshalJSON(s string) error {
	return v.Parse(s)
}

// monthStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var monthStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Month) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		if monthStoreRep == enum.Ordinal {
			*v = MonthOf(int(x))
		} else {
			*v = Month(x)
		}
		return nil
	case float64:
		*v = Month(x)
		return nil
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful month", value, value)
	}

	return v.parse(s, monthStoreRep)
}

// Value converts the Month to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Month) Value() (driver.Value, error) {
	switch monthStoreRep {
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
