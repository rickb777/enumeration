// generated code - do not edit
// github.com/rickb777/enumeration v1.10.0

package example

import (
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/enum"
	"strconv"
	"strings"
)

const methodEnumStrings = "HEADGETPUTPOSTPATCHDELETE"

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

// Literal returns the literal string representation of a Method, which is
// the same as the const identifier.
func (i Method) Literal() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllMethods) {
		return fmt.Sprintf("Method(%d)", i)
	}
	return methodEnumStrings[methodEnumIndex[o]:methodEnumIndex[o+1]]
}

var methodStringsInverse = map[string]Method{}

func init() {
	if len(methodStrings) != 6 {
		panic(fmt.Sprintf("methodStrings has %d items but should have 6", len(methodStrings)))
	}

	for k, v := range methodStrings {
		methodStringsInverse[v] = k
	}

	if len(methodStrings) != len(methodStringsInverse) {
		panic(fmt.Sprintf("methodStrings has %d items but they are not distinct", len(methodStrings)))
	}
}

// String returns the string representation of a Method.
func (i Method) String() string {
	s, ok := methodStrings[i]
	if ok {
		return s
	}
	return fmt.Sprintf("%02d", i)
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
	switch i {
	case HEAD, GET, PUT, POST,
		PATCH, DELETE:
		return true
	}
	return false
}

// Parse parses a string to find the corresponding Method, accepting one of the string
// values or an ordinal number.
func (v *Method) Parse(in string) error {
	if v.parseOrdinal(in) {
		return nil
	}

	s := in

	if methodMarshalTextUsingLiteral {
		if v.parseIdentifier(s) || v.parseString(in) {
			return nil
		}
	} else {
		if v.parseString(in) || v.parseIdentifier(s) {
			return nil
		}
	}

	return errors.New(in + ": unrecognised Method")
}

// parseOrdinal attempts to convert ordinal value
func (v *Method) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllMethods) {
		*v = AllMethods[ord]
		return true
	}
	return false
}

// parseString attempts to match an entry in methodStringsInverse
func (v *Method) parseString(s string) (ok bool) {
	*v, ok = methodStringsInverse[s]
	return ok
}

// parseIdentifier attempts to match an identifier.
func (v *Method) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0
	for j := 1; j < len(methodEnumIndex); j++ {
		i1 := methodEnumIndex[j]
		p := methodEnumStrings[i0:i1]
		if s == p {
			*v = AllMethods[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsMethod parses a string to find the corresponding Method, accepting either one of the string
// values or an ordinal number.
func AsMethod(s string) (Method, error) {
	var i = new(Method)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Method) MarshalText() (text []byte, err error) {
	if methodMarshalTextUsingLiteral {
		return []byte(i.Literal()), nil
	}
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Method) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// MethodMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var MethodMarshalJSONUsingString = false

// methodMarshalTextUsingLiteral controls whether generated XML or JSON uses the String()
// or the Literal() method.
var methodMarshalTextUsingLiteral = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// MethodMarshalJSONUsingString is true.
func (i Method) MarshalJSON() ([]byte, error) {
	if !MethodMarshalJSONUsingString {
		// use the ordinal
		s := strconv.Itoa(i.Ordinal())
		return []byte(s), nil
	}
	if methodMarshalTextUsingLiteral {
		return i.quotedString(i.Literal())
	}
	return i.quotedString(i.String())
}

func (i Method) quotedString(s string) ([]byte, error) {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b, nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Method) UnmarshalJSON(text []byte) error {
	if len(text) >= 2 && text[0] == '"' && text[len(text)-1] == '"' {
		s := string(text[1 : len(text)-1])
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
func (i *Method) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		*i = Method(v)
	case float64:
		*i = Method(v)
	case []byte:
		err = i.Parse(string(v))
	case string:
		err = i.Parse(v)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful Method", value, value)
	}

	return err
}

// -- copy this somewhere and uncomment it if you need DB storage to use strings --
// Value converts the period to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
//func (i Method) Value() (driver.Value, error) {
//    return i.String(), nil
//}
