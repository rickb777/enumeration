// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.9.0

package transform

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"strconv"
	"strings"
)

const caseEnumStrings = "StetUpperLower"

var caseEnumIndex = [...]uint16{0, 4, 9, 14}

// AllCases lists all 3 values in order.
var AllCases = []Case{
	Stet, Upper, Lower,
}

// AllCaseEnums lists all 3 values in order.
var AllCaseEnums = enum.IntEnums{
	Stet, Upper, Lower,
}

// String returns the literal string representation of a Case, which is
// the same as the const identifier.
func (i Case) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllCases) {
		return fmt.Sprintf("Case(%d)", i)
	}
	return caseEnumStrings[caseEnumIndex[o]:caseEnumIndex[o+1]]
}

// Tag returns the string representation of a Case. This is an alias for String.
func (i Case) Tag() string {
	return i.String()
}

// Ordinal returns the ordinal number of a Case.
func (i Case) Ordinal() int {
	switch i {
	case Stet:
		return 0
	case Upper:
		return 1
	case Lower:
		return 2
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Case) Int() int {
	return int(i)
}

// CaseOf returns a Case based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Case is returned.
func CaseOf(i int) Case {
	if 0 <= i && i < len(AllCases) {
		return AllCases[i]
	}
	// an invalid result
	return Stet + Upper + Lower + 1
}

// IsValid determines whether a Case is one of the defined constants.
func (i Case) IsValid() bool {
	return i.Ordinal() >= 0
}

// Parse parses a string to find the corresponding Case, accepting one of the string values or
// a number. The input representation is determined by caseMarshalTextRep. It is used by AsCase.
//
// Usage Example
//
//    v := new(Case)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Case) Parse(s string) error {
	return v.parse(s, caseMarshalTextRep)
}

func (v *Case) parse(in string, rep enum.Representation) error {
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

	return errors.New(in + ": unrecognised case")
}

// parseNumber attempts to convert a decimal value
func (v *Case) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Case(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Case) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllCases) {
		*v = AllCases[ord]
		return true
	}
	return false
}

// parseIdentifier attempts to match an identifier.
func (v *Case) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(caseEnumIndex); j++ {
		i1 := caseEnumIndex[j]
		p := caseEnumStrings[i0:i1]
		if s == p {
			*v = AllCases[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsCase parses a string to find the corresponding Case, accepting either one of the string values or
// a number. The input representation is determined by caseMarshalTextRep. It wraps Parse.
func AsCase(s string) (Case, error) {
	var i = new(Case)
	err := i.Parse(s)
	return *i, err
}

// MustParseCase is similar to AsCase except that it panics on error.
func MustParseCase(s string) Case {
	i, err := AsCase(s)
	if err != nil {
		panic(err)
	}
	return i
}

// caseMarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var caseMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to caseMarshalTextRep.
func (i Case) MarshalText() (text []byte, err error) {
	return i.marshalText(caseMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to caseMarshalTextRep.
func (i Case) MarshalJSON() ([]byte, error) {
	return i.marshalText(caseMarshalTextRep, true)
}

func (i Case) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
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

func (i Case) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Case) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Case) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}

// caseStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var caseStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Case) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if caseStoreRep == enum.Ordinal {
			*i = CaseOf(int(v))
		} else {
			*i = Case(v)
		}
	case float64:
		*i = Case(v)
	case []byte:
		err = i.parse(string(v), caseStoreRep)
	case string:
		err = i.parse(v, caseStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful case", value, value)
	}

	return err
}

// Value converts the Case to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Case) Value() (driver.Value, error) {
	switch caseStoreRep {
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
