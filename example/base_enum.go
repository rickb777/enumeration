// generated code - do not edit
// bitbucket.org/rickb777/enumeration v1.1.0

package example

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const baseEnumStrings = "ACGT"

var baseEnumIndex = [...]uint16{0, 1, 2, 3, 4}

// AllBases lists all 4 values in order.
var AllBases = []Base{A, C, G, T}

// String returns the string representation of a Base.
func (i Base) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllBases) {
		return fmt.Sprintf("Base(%g)", i)
	}
	return baseEnumStrings[baseEnumIndex[o]:baseEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Base.
func (i Base) Ordinal() int {
	switch i {
	case A:
		return 0
	case C:
		return 1
	case G:
		return 2
	case T:
		return 3
	}
	return -1
}

// Parse parses a string to find the corresponding Base, accepting either one of the string
// values or an ordinal number.
func (v *Base) Parse(s string) error {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllBases) {
		*v = AllBases[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(baseEnumIndex); j++ {
		i1 := baseEnumIndex[j]
		p := baseEnumStrings[i0:i1]
		if s == p {
			*v = AllBases[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Base")
}

// AsBase parses a string to find the corresponding Base, accepting either one of the string
// values or an ordinal number.
func AsBase(s string) (Base, error) {
	var i = new(Base)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Base) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Base) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// BaseMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var BaseMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// BaseMarshalJSONUsingString is true.
func (i Base) MarshalJSON() ([]byte, error) {
	if BaseMarshalJSONUsingString {
		s := []byte(i.String())
		b := make([]byte, len(s)+2)
		b[0] = '"'
		copy(b[1:], s)
		b[len(s)+1] = '"'
		return b, nil
	}
	// else use the ordinal
	s := strconv.Itoa(i.Ordinal())
	return []byte(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Base) UnmarshalJSON(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
