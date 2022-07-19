// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.14.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"os"
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
)

var (
	methodEnumIndex = [...]uint16{0, 4, 7, 10, 14, 19, 25}
)

func (v Method) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllMethods) {
		return fmt.Sprintf("Method(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
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

var methodTagsInverse = map[string]Method{}

func init() {
	for _, id := range AllMethods {
		v, exists := methodTags[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Method: %s is missing from methodTags\n", id)
		} else {
			k := methodTransformInput(v)
			if _, exists := methodTagsInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: Method: %q is duplicated in methodTags\n", k)
			}
			methodTagsInverse[k] = id
		}
	}

	if len(methodTags) != 6 {
		panic(fmt.Sprintf("Method: methodTags has %d items but should have 6", len(methodTags)))
	}

	if len(methodTags) != len(methodTagsInverse) {
		panic(fmt.Sprintf("Method: methodTags has %d items but there are only %d distinct items",
			len(methodTags), len(methodTagsInverse)))
	}
}

// Tag returns the string representation of a Method. For invalid values,
// this returns v.String() (see IsValid).
func (v Method) Tag() string {
	s, ok := methodTags[v]
	if ok {
		return s
	}
	return v.String()
}

// String returns the literal string representation of a Method, which is
// the same as the const identifier but without prefix or suffix.
func (v Method) String() string {
	return v.toString(methodEnumStrings, methodEnumIndex[:])
}

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

// IsValid determines whether a Method is one of the defined constants.
func (v Method) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Method) Int() int {
	return int(v)
}

// MethodOf returns a Method based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Method is returned.
func MethodOf(v int) Method {
	if 0 <= v && v < len(AllMethods) {
		return AllMethods[v]
	}
	// an invalid result
	return HEAD + GET + PUT + POST + PATCH + DELETE + 1
}

// Parse parses a string to find the corresponding Method, accepting one of the string values or
// a number. The input representation is determined by methodMarshalTextRep. It is used by AsMethod.
// The input case does not matter.
//
// Usage Example
//
//    v := new(Method)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Method) Parse(s string) error {
	return v.parse(s, methodMarshalTextRep)
}

func (v *Method) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := methodTransformInput(in)

	if rep == enum.Identifier {
		if v.parseString(s, methodEnumInputs, methodEnumIndex[:]) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseString(s, methodEnumInputs, methodEnumIndex[:]) {
			return nil
		}
	}

	return errors.New(in + ": unrecognised method")
}

// parseNumber attempts to convert a decimal value
func (v *Method) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Method(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Method) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllMethods) {
		*v = AllMethods[ord]
		return true
	}
	return false
}

// parseTag attempts to match an entry in methodTagsInverse
func (v *Method) parseTag(s string) (ok bool) {
	*v, ok = methodTagsInverse[s]
	return ok
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

// methodMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
var methodMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to methodMarshalTextRep.
func (v Method) MarshalText() (text []byte, err error) {
	return v.marshalText(methodMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to methodMarshalTextRep.
func (v Method) MarshalJSON() ([]byte, error) {
	return v.marshalText(methodMarshalTextRep, true)
}

func (v Method) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	if rep != enum.Ordinal && !v.IsValid() {
		return methodMarshalNumber(v)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return methodMarshalNumber(v)
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

// methodMarshalNumber handles marshaling where a number is required or where
// the value is out of range but methodMarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var methodMarshalNumber = func(v Method) (text []byte, err error) {
	bs := []byte(strconv.FormatInt(int64(v), 10))
	return bs, nil
}

func (v Method) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(v.Ordinal()))
	return bs, nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *Method) UnmarshalText(text []byte) error {
	return v.Parse(string(text))
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

func (v *Method) unmarshalJSON(s string) error {
	return v.Parse(s)
}

// methodStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var methodStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *Method) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		if methodStoreRep == enum.Ordinal {
			*v = MethodOf(int(x))
		} else {
			*v = Method(x)
		}
		return nil
	case float64:
		*v = Method(x)
		return nil
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful method", value, value)
	}

	return v.parse(s, methodStoreRep)
}

// Value converts the Method to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Method) Value() (driver.Value, error) {
	switch methodStoreRep {
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
