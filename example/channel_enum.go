// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-dirty

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
	"slices"
	"strconv"
	"strings"
)

// AllSalesChannels lists all 3 values in order.
var AllSalesChannels = []SalesChannel{
	OnlineSales, InstoreSales, TelephoneSales,
}

// AllSalesChannelEnums lists all 3 values in order.
var AllSalesChannelEnums = enum.IntEnums{
	OnlineSales, InstoreSales, TelephoneSales,
}

const (
	saleschannelEnumStrings = "onlineinstoretelephone"
	saleschannelJSONStrings = "webshopstorephone"
	saleschannelSQLStrings  = "ost"
)

var (
	saleschannelEnumIndex = [...]uint16{0, 6, 13, 22}
	saleschannelJSONIndex = [...]uint16{0, 7, 12, 17}
	saleschannelSQLIndex  = [...]uint16{0, 1, 2, 3}
)

// String returns the literal string representation of a SalesChannel, which is
// the same as the const identifier but without prefix or suffix.
func (v SalesChannel) String() string {
	o := v.Ordinal()
	return v.toString(o, saleschannelEnumStrings, saleschannelEnumIndex[:])
}

// Ordinal returns the ordinal number of a SalesChannel. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v SalesChannel) Ordinal() int {
	switch v {
	case OnlineSales:
		return 0
	case InstoreSales:
		return 1
	case TelephoneSales:
		return 2
	}
	return -1
}

func (v SalesChannel) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSalesChannels) {
		return fmt.Sprintf("SalesChannel(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a SalesChannel is one of the defined constants.
func (v SalesChannel) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v SalesChannel) Int() int {
	return int(v)
}

var invalidSalesChannelValue = func() SalesChannel {
	var v SalesChannel
	for {
		if !slices.Contains(AllSalesChannels, v) {
			return v
		}
		v++
	} // AllSalesChannels is a finite set so loop will terminate eventually
}()

// SalesChannelOf returns a SalesChannel based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid SalesChannel is returned.
func SalesChannelOf(v int) SalesChannel {
	if 0 <= v && v < len(AllSalesChannels) {
		return AllSalesChannels[v]
	}
	return invalidSalesChannelValue
}

// Parse parses a string to find the corresponding SalesChannel, accepting one of the string values or
// a number. It is used by AsSalesChannel.
//
// Usage Example
//
//	v := new(SalesChannel)
//	err := v.Parse(s)
//	...  etc
func (v *SalesChannel) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := saleschannelTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *SalesChannel) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = SalesChannel(num)
		return v.IsValid()
	}
	return false
}

func (v *SalesChannel) parseFallback(in, s string) error {
	if v.parseString(s, saleschannelEnumStrings, saleschannelEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised saleschannel")
}

func (v *SalesChannel) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSalesChannels[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// saleschannelTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var saleschannelTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsSalesChannel parses a string to find the corresponding SalesChannel, accepting either one of the string values or
// a number. It wraps Parse.
func AsSalesChannel(s string) (SalesChannel, error) {
	var v = new(SalesChannel)
	err := v.Parse(s)
	return *v, err
}

// MustParseSalesChannel is similar to AsSalesChannel except that it panics on error.
func MustParseSalesChannel(s string) SalesChannel {
	v, err := AsSalesChannel(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v SalesChannel) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, saleschannelJSONStrings, saleschannelJSONIndex[:])
}

func (v SalesChannel) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v SalesChannel) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v SalesChannel) invalidError() error {
	return fmt.Errorf("%d is not a valid saleschannel", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v SalesChannel) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, saleschannelJSONStrings, saleschannelJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *SalesChannel) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *SalesChannel) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := saleschannelTransformInput(in)

	if v.parseString(s, saleschannelJSONStrings, saleschannelJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, saleschannelEnumStrings, saleschannelEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised saleschannel")
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *SalesChannel) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = SalesChannel(x)
		return v.errorIfInvalid()
	case float64:
		*v = SalesChannel(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful saleschannel", value, value)
	}

	return v.scanParse(s)
}

func (v SalesChannel) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v *SalesChannel) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := saleschannelTransformInput(in)

	if v.parseString(s, saleschannelSQLStrings, saleschannelSQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// Value converts the SalesChannel to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v SalesChannel) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, saleschannelSQLStrings, saleschannelSQLIndex[:]), nil
}
