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

func (v SalesChannel) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllSalesChannels) {
		return fmt.Sprintf("SalesChannel(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
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

// Tag returns the JSON representation of a SalesChannel.
func (v SalesChannel) Tag() string {
	return v.toString(saleschannelJSONStrings, saleschannelJSONIndex[:])
}

// String returns the literal string representation of a SalesChannel, which is
// the same as the const identifier but without prefix or suffix.
func (v SalesChannel) String() string {
	return v.toString(saleschannelEnumStrings, saleschannelEnumIndex[:])
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

// IsValid determines whether a SalesChannel is one of the defined constants.
func (v SalesChannel) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v SalesChannel) Int() int {
	return int(v)
}

// SalesChannelOf returns a SalesChannel based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid SalesChannel is returned.
func SalesChannelOf(v int) SalesChannel {
	if 0 <= v && v < len(AllSalesChannels) {
		return AllSalesChannels[v]
	}
	// an invalid result
	return OnlineSales + InstoreSales + TelephoneSales + 1
}

// Parse parses a string to find the corresponding SalesChannel, accepting one of the string values or
// a number. The input representation is determined by saleschannelMarshalTextRep. It is used by AsSalesChannel.
//
// Usage Example
//
//    v := new(SalesChannel)
//    err := v.Parse(s)
//    ...  etc
//
func (v *SalesChannel) Parse(s string) error {
	return v.parse(s, saleschannelMarshalTextRep)
}

func (v *SalesChannel) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := saleschannelTransformInput(in)

	if v.parseString(s, saleschannelEnumStrings, saleschannelEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised saleschannel")
}

// parseNumber attempts to convert a decimal value
func (v *SalesChannel) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = SalesChannel(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *SalesChannel) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSalesChannels) {
		*v = AllSalesChannels[ord]
		return true
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
// a number. The input representation is determined by saleschannelMarshalTextRep. It wraps Parse.
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

// saleschannelMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
// By default, it is enum.Identifier.
// The initial value is set using the -marshaltext command line parameter.
var saleschannelMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to saleschannelMarshalTextRep.
func (v SalesChannel) MarshalText() (text []byte, err error) {
	return v.marshalText(saleschannelMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to saleschannelMarshalTextRep.
func (v SalesChannel) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		if saleschannelMarshalTextRep == enum.Ordinal {
			return nil, fmt.Errorf("%v is out of range", v)
		}
		return saleschannelMarshalNumber(v)
	}
	s := saleschannelJSONStrings[saleschannelJSONIndex[o]:saleschannelJSONIndex[o+1]]
	return enum.QuotedString(s), nil
}

func (v SalesChannel) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	if rep != enum.Ordinal && !v.IsValid() {
		return saleschannelMarshalNumber(v)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return saleschannelMarshalNumber(v)
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

// saleschannelMarshalNumber handles marshaling where a number is required or where
// the value is out of range but saleschannelMarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var saleschannelMarshalNumber = func(v SalesChannel) (text []byte, err error) {
	bs := []byte(strconv.FormatInt(int64(v), 10))
	return bs, nil
}

func (v SalesChannel) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(v.Ordinal()))
	return bs, nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *SalesChannel) UnmarshalText(text []byte) error {
	return v.Parse(string(text))
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

	return errors.New(in + ": unrecognised saleschannel")
}

// saleschannelStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier.
// The initial value is set using the -store command line parameter.
var saleschannelStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *SalesChannel) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		if saleschannelStoreRep == enum.Ordinal {
			*v = SalesChannelOf(int(x))
		} else {
			*v = SalesChannel(x)
		}
		return nil
	case float64:
		*v = SalesChannel(x)
		return nil
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful saleschannel", value, value)
	}

	if v.parseString(s, saleschannelSQLStrings, saleschannelSQLIndex[:]) {
		return nil
	}

	return errors.New(s + ": unrecognised saleschannel")
}

// Value converts the SalesChannel to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v SalesChannel) Value() (driver.Value, error) {
	switch saleschannelStoreRep {
	case enum.Number:
		return int64(v), nil
	case enum.Ordinal:
		return int64(v.Ordinal()), nil
	case enum.Tag:
		return v.Tag(), nil
	default:
		return v.toString(saleschannelSQLStrings, saleschannelSQLIndex[:]), nil
	}
}
