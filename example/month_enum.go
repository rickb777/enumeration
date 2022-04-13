// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.10.1

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"strconv"
	"strings"
)

const (
	monthEnumStrings = "JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember"
	monthEnumInputs  = "januaryfebruarymarchaprilmayjunejulyaugustseptemberoctobernovemberdecember"
)

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

// Ordinal returns the ordinal number of a Month. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
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
	return i.Ordinal() >= 0
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

	s := strings.ToLower(in)

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
		p := monthEnumInputs[i0:i1]
		if s == p {
			*v = AllMonths[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsMonth parses a string to find the corresponding Month, accepting either one of the string values or
// a number. The input representation is determined by monthMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsMonth(s string) (Month, error) {
	var i = new(Month)
	err := i.Parse(s)
	return *i, err
}

// MustParseMonth is similar to AsMonth except that it panics on error.
// The input case does not matter.
func MustParseMonth(s string) Month {
	i, err := AsMonth(s)
	if err != nil {
		panic(err)
	}
	return i
}

// monthMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
var monthMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to monthMarshalTextRep.
func (i Month) MarshalText() (text []byte, err error) {
	return i.marshalText(monthMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to monthMarshalTextRep.
func (i Month) MarshalJSON() ([]byte, error) {
	return i.marshalText(monthMarshalTextRep, true)
}

func (i Month) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
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

func (i Month) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Month) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
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

// monthStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var monthStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Month) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if monthStoreRep == enum.Ordinal {
			*i = MonthOf(int(v))
		} else {
			*i = Month(v)
		}
	case float64:
		*i = Month(v)
	case []byte:
		err = i.parse(string(v), monthStoreRep)
	case string:
		err = i.parse(v, monthStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful month", value, value)
	}

	return err
}

// Value converts the Month to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Month) Value() (driver.Value, error) {
	switch monthStoreRep {
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
