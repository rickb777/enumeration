// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.2.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"slices"
	"strconv"
	"strings"
)

// AllMethods lists all 6 values in order.
var AllMethods = []Method{
	HEAD, GET, PUT, POST,
	PATCH, DELETE,
}

// AllMethodEnums lists all 6 values in order.
var AllMethodEnums = enum.IntEnums{
	HEAD, GET, PUT, POST,
	PATCH, DELETE,
}

const (
	methodEnumStrings = "HEADGETPUTPOSTPATCHDELETE"
	methodEnumInputs  = "headgetputpostpatchdelete"
	methodJSONStrings = "HEGEPUPOPADE"
	methodJSONInputs  = "HEGEPUPOPADE"
)

var (
	methodEnumIndex = [...]uint16{0, 4, 7, 10, 14, 19, 25}
	methodJSONIndex = [...]uint16{0, 2, 4, 6, 8, 10, 12}
)

// Ordinal returns the ordinal number of a Method. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Method) Ordinal() int {
	switch v {
	case HEAD:
		return 0
	case GET:
		return 1
	case PUT:
		return 2
	case POST:
		return 3
	case PATCH:
		return 4
	case DELETE:
		return 5
	}
	return -1
}

// String returns the literal string representation of a Method, which is
// the same as the const identifier but without prefix or suffix.
func (v Method) String() string {
	o := v.Ordinal()
	return v.toString(o, methodEnumStrings, methodEnumIndex[:])
}

func (v Method) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllMethods) {
		return fmt.Sprintf("Method(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Method is one of the defined constants.
func (v Method) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Method) Int() int {
	return int(v)
}

var invalidMethodValue = func() Method {
	var v Method
	for {
		if !slices.Contains(AllMethods, v) {
			return v
		}
		v++
	} // AllMethods is a finite set so loop will terminate eventually
}()

// MethodOf returns a Method based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Method is returned.
func MethodOf(v int) Method {
	if 0 <= v && v < len(AllMethods) {
		return AllMethods[v]
	}
	return invalidMethodValue
}

// Parse parses a string to find the corresponding Method, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsMethod.
// The input case does not matter.
//
// Usage Example
//
//	v := new(Method)
//	err := v.Parse(s)
//	...  etc
func (v *Method) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := methodTransformInput(in)

	return v.parseFallback(in, s)
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *Method) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Method(num)
		return v.IsValid()
	}
	return false
}

func (v *Method) parseFallback(in, s string) error {
	if v.parseString(s, methodEnumInputs, methodEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised method")
}

func (v *Method) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllMethods[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// methodTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var methodTransformInput = func(in string) string {
	return strings.ToLower(in)
}

// AsMethod parses a string to find the corresponding Method, accepting either one of the string values or
// a number. The input representation is determined by methodMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsMethod(s string) (Method, error) {
	var v = new(Method)
	err := v.Parse(s)
	return *v, err
}

// MustParseMethod is similar to AsMethod except that it panics on error.
// The input case does not matter.
func MustParseMethod(s string) Method {
	v, err := AsMethod(s)
	if err != nil {
		panic(err)
	}
	return v
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Method) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, methodJSONStrings, methodJSONIndex[:])
}

func (v Method) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Method) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Method) invalidError() error {
	return fmt.Errorf("%d is not a valid method", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Method) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, methodJSONStrings, methodJSONIndex[:])
	return enum.QuotedString(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *Method) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *Method) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := methodTransformInput(in)

	if v.parseString(s, methodJSONInputs, methodJSONIndex[:]) {
		return nil
	}

	if v.parseString(s, methodEnumInputs, methodEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised method")
}

// methodMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var methodMarshalNumber = func(v Method) string {
	return strconv.FormatInt(int64(v), 10)
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Method) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = Method(x)
		return v.errorIfInvalid()
	case float64:
		*v = Method(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful method", value, value)
	}

	return v.scanParse(s)
}

func (v Method) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

func (v *Method) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := methodTransformInput(in)

	return v.parseFallback(in, s)
}

// Value converts the Method to a number (based on '-store number').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Method) Value() (driver.Value, error) {
	return int64(v), nil
}
