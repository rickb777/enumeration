// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.9.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"strconv"
	"strings"
)

const saleschannelEnumStrings = "onlineinstoretelephone"

var saleschannelEnumIndex = [...]uint16{0, 6, 13, 22}

// AllSalesChannels lists all 3 values in order.
var AllSalesChannels = []SalesChannel{
	OnlineSales, InstoreSales, TelephoneSales,
}

// AllSalesChannelEnums lists all 3 values in order.
var AllSalesChannelEnums = enum.IntEnums{
	OnlineSales, InstoreSales, TelephoneSales,
}

// String returns the literal string representation of a SalesChannel, which is
// the same as the const identifier.
func (i SalesChannel) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSalesChannels) {
		return fmt.Sprintf("SalesChannel(%d)", i)
	}
	return saleschannelEnumStrings[saleschannelEnumIndex[o]:saleschannelEnumIndex[o+1]]
}

// Tag returns the string representation of a SalesChannel. This is an alias for String.
func (i SalesChannel) Tag() string {
	return i.String()
}

// Ordinal returns the ordinal number of a SalesChannel. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (i SalesChannel) Ordinal() int {
	switch i {
	case OnlineSales:
		return 0
	case InstoreSales:
		return 1
	case TelephoneSales:
		return 2
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i SalesChannel) Int() int {
	return int(i)
}

// SalesChannelOf returns a SalesChannel based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid SalesChannel is returned.
func SalesChannelOf(i int) SalesChannel {
	if 0 <= i && i < len(AllSalesChannels) {
		return AllSalesChannels[i]
	}
	// an invalid result
	return OnlineSales + InstoreSales + TelephoneSales + 1
}

// IsValid determines whether a SalesChannel is one of the defined constants.
func (i SalesChannel) IsValid() bool {
	return i.Ordinal() >= 0
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

	s := strings.ToLower(in)

	if v.parseIdentifier(s) {
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

// parseIdentifier attempts to match an identifier.
func (v *SalesChannel) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(saleschannelEnumIndex); j++ {
		i1 := saleschannelEnumIndex[j]
		p := saleschannelEnumStrings[i0:i1]
		if s == p {
			*v = AllSalesChannels[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsSalesChannel parses a string to find the corresponding SalesChannel, accepting either one of the string values or
// a number. The input representation is determined by saleschannelMarshalTextRep. It wraps Parse.
func AsSalesChannel(s string) (SalesChannel, error) {
	var i = new(SalesChannel)
	err := i.Parse(s)
	return *i, err
}

// MustParseSalesChannel is similar to AsSalesChannel except that it panics on error.
func MustParseSalesChannel(s string) SalesChannel {
	i, err := AsSalesChannel(s)
	if err != nil {
		panic(err)
	}
	return i
}

// saleschannelMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
var saleschannelMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to saleschannelMarshalTextRep.
func (i SalesChannel) MarshalText() (text []byte, err error) {
	return i.marshalText(saleschannelMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to saleschannelMarshalTextRep.
func (i SalesChannel) MarshalJSON() ([]byte, error) {
	return i.marshalText(saleschannelMarshalTextRep, true)
}

func (i SalesChannel) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
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

func (i SalesChannel) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *SalesChannel) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *SalesChannel) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}

// saleschannelStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var saleschannelStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *SalesChannel) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if saleschannelStoreRep == enum.Ordinal {
			*i = SalesChannelOf(int(v))
		} else {
			*i = SalesChannel(v)
		}
	case float64:
		*i = SalesChannel(v)
	case []byte:
		err = i.parse(string(v), saleschannelStoreRep)
	case string:
		err = i.parse(v, saleschannelStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful saleschannel", value, value)
	}

	return err
}

// Value converts the SalesChannel to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i SalesChannel) Value() (driver.Value, error) {
	switch saleschannelStoreRep {
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
