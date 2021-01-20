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

const monthEnumStrings = "JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember"

var monthEnumIndex = [...]uint16{0, 7, 15, 20, 25, 28, 32, 36, 42, 51, 58, 66, 74}

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

// String returns the literal string representation of a Month, which is
// the same as the const identifier.
func (i Month) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllMonths) {
		return fmt.Sprintf("Month(%d)", i)
	}
	return monthEnumStrings[monthEnumIndex[o]:monthEnumIndex[o+1]]
}

// Tag returns the string representation of a Month. This is an alias for String.
func (i Month) Tag() string {
	return i.String()
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

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Month) Int() int {
	return int(i)
}

// MonthOf returns a Month based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Month is returned.
func MonthOf(i int) Month {
	if 0 <= i && i < len(AllMonths) {
		return AllMonths[i]
	}
	// an invalid result
	return January + February + March + April + May + June + July + August + September + October + November + December + 1
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

// Parse parses a string to find the corresponding Month, accepting one of the string
// values or a number.
func (v *Month) Parse(in string) error {
	if monthMarshalTextUsing == enum.Ordinal {
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

// parseIdentifier attempts to match an identifier.
func (v *Month) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0
	for j := 1; j < len(monthEnumIndex); j++ {
		i1 := monthEnumIndex[j]
		p := monthEnumStrings[i0:i1]
		if s == p {
			*v = AllMonths[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsMonth parses a string to find the corresponding Month, accepting either one of the string
// values or an ordinal number.
func AsMonth(s string) (Month, error) {
	var i = new(Month)
	err := i.Parse(s)
	return *i, err
}

// monthMarshalTextUsingLiteral controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var monthMarshalTextUsing = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to MonthMarshalTextUsing.
func (i Month) MarshalText() (text []byte, err error) {
	var s string
	switch monthMarshalTextUsing {
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
func (i *Month) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to MonthMarshalTextUsing.
func (i Month) MarshalJSON() ([]byte, error) {
	var s []byte
	switch monthMarshalTextUsing {
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

func (i Month) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Month) UnmarshalJSON(text []byte) error {
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
func (i *Month) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		*i = Month(v)
	case float64:
		*i = Month(v)
	case []byte:
		err = i.Parse(string(v))
	case string:
		err = i.Parse(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful Month", value, value)
	}

	return err
}

// -- copy this somewhere and uncomment it if you need DB storage to use strings --
// Value converts the Month to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i Month) Value() (driver.Value, error) {
//    return i.String(), nil
//}
