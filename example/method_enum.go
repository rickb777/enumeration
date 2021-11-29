// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.8.0

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

const (
	methodEnumStrings = "HEADGETPUTPOSTPATCHDELETE"
	methodEnumInputs  = "headgetputpostpatchdelete"
)

var methodEnumIndex = [...]uint16{0, 4, 7, 10, 14, 19, 25}

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

// String returns the literal string representation of a Method, which is
// the same as the const identifier.
func (i Method) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllMethods) {
		return fmt.Sprintf("Method(%d)", i)
	}
	return methodEnumStrings[methodEnumIndex[o]:methodEnumIndex[o+1]]
}

var methodStringsInverse = map[string]Method{}

func init() {
	for _, id := range AllMethods {
		v, exists := methodStrings[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Method: %s is missing from methodStrings\n", id)
		} else {
			k := strings.ToLower(v)
			if _, exists := methodStringsInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: Method: %q is duplicated in methodStrings\n", k)
			}
			methodStringsInverse[k] = id
		}
	}

	if len(methodStrings) != 6 {
		panic(fmt.Sprintf("Method: methodStrings has %d items but should have 6", len(methodStrings)))
	}

	if len(methodStrings) != len(methodStringsInverse) {
		panic(fmt.Sprintf("Method: methodStrings has %d items but there are only %d distinct items",
			len(methodStrings), len(methodStringsInverse)))
	}
}

// Tag returns the string representation of a Method. For invalid values,
// this returns i.String() (see IsValid).
func (i Method) Tag() string {
	s, ok := methodStrings[i]
	if ok {
		return s
	}
	return i.String()
}

// Ordinal returns the ordinal number of a Method.
func (i Method) Ordinal() int {
	switch i {
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

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Method) Int() int {
	return int(i)
}

// MethodOf returns a Method based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Method is returned.
func MethodOf(i int) Method {
	if 0 <= i && i < len(AllMethods) {
		return AllMethods[i]
	}
	// an invalid result
	return HEAD + GET + PUT + POST + PATCH + DELETE + 1
}

// IsValid determines whether a Method is one of the defined constants.
func (i Method) IsValid() bool {
	return i.Ordinal() >= 0
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

	s := strings.ToLower(in)

	if rep == enum.Identifier {
		if v.parseIdentifier(s) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseIdentifier(s) {
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

// parseTag attempts to match an entry in methodStringsInverse
func (v *Method) parseTag(s string) (ok bool) {
	*v, ok = methodStringsInverse[s]
	return ok
}

// parseIdentifier attempts to match an identifier.
func (v *Method) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(methodEnumIndex); j++ {
		i1 := methodEnumIndex[j]
		p := methodEnumInputs[i0:i1]
		if s == p {
			*v = AllMethods[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsMethod parses a string to find the corresponding Method, accepting either one of the string values or
// a number. The input representation is determined by methodMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsMethod(s string) (Method, error) {
	var i = new(Method)
	err := i.Parse(s)
	return *i, err
}

// MustParseMethod is similar to AsMethod except that it panics on error.
// The input case does not matter.
func MustParseMethod(s string) Method {
	i, err := AsMethod(s)
	if err != nil {
		panic(err)
	}
	return i
}

// methodMarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var methodMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to methodMarshalTextRep.
func (i Method) MarshalText() (text []byte, err error) {
	return i.marshalText(methodMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to methodMarshalTextRep.
func (i Method) MarshalJSON() ([]byte, error) {
	return i.marshalText(methodMarshalTextRep, true)
}

func (i Method) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
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

func (i Method) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Method) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Method) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}

// methodStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var methodStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Method) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if methodStoreRep == enum.Ordinal {
			*i = MethodOf(int(v))
		} else {
			*i = Method(v)
		}
	case float64:
		*i = Method(v)
	case []byte:
		err = i.parse(string(v), methodStoreRep)
	case string:
		err = i.parse(v, methodStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful method", value, value)
	}

	return err
}

// Value converts the Method to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Method) Value() (driver.Value, error) {
	switch methodStoreRep {
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
