// generated code - do not edit
// github.com/rickb777/enumeration v1.10.0

package example

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/rickb777/enumeration/enum"
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

// Literal returns the literal string representation of a Month, which is
// the same as the const identifier.
func (i Month) Literal() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllMonths) {
		return fmt.Sprintf("Month(%d)", i)
	}
	return monthEnumStrings[monthEnumIndex[o]:monthEnumIndex[o+1]]
}

// String returns the string representation of a Month. This uses Literal.
func (i Month) String() string {
	return i.Literal()
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
// values or an ordinal number.
func (v *Month) Parse(s string) error {
	// attempt to convert ordinal value
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllMonths) {
		*v = AllMonths[ord]
		return nil
	}

	// attempt to match an identifier
	var i0 uint16 = 0
	for j := 1; j < len(monthEnumIndex); j++ {
		i1 := monthEnumIndex[j]
		p := monthEnumStrings[i0:i1]
		if s == p {
			*v = AllMonths[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Month")
}

// AsMonth parses a string to find the corresponding Month, accepting either one of the string
// values or an ordinal number.
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

// MonthMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var MonthMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// MonthMarshalJSONUsingString is true.
func (i Month) MarshalJSON() ([]byte, error) {
	if !MonthMarshalJSONUsingString {
		// use the ordinal
		s := strconv.Itoa(i.Ordinal())
		return []byte(s), nil
	}
	return i.quotedString(i.String())
}

func (i Month) quotedString(s string) ([]byte, error) {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b, nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Month) UnmarshalJSON(text []byte) error {
	if len(text) >= 2 && text[0] == '"' && text[len(text)-1] == '"' {
		s := string(text[1:len(text)-1])
		return i.Parse(s)
	}

	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
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
// Value converts the period to a string. 
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i Month) Value() (driver.Value, error) {
//    return i.String(), nil
//}
