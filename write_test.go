package main

import (
	"bytes"
	. "github.com/onsi/gomega"
	"testing"
)

const e0 = `// generated code - do not edit
// bitbucket.org/rickb777/enumeration `

const e1 = `

package confectionary

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

`

const e3 = `

var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24}

// AllSweets lists all 4 values in order.
var AllSweets = []Sweet{Mars, Bounty, Snickers, Kitkat}
`

const e4 = `
// String returns the string representation of a Sweet.
func (i Sweet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[o]:sweetEnumIndex[o+1]]
}
`

const e5 = `
// Ordinal returns the ordinal number of a Sweet.
func (i Sweet) Ordinal() int {
	switch i {
	case Mars:
		return 0
	case Bounty:
		return 1
	case Snickers:
		return 2
	case Kitkat:
		return 3
	}
	return -1
}
`

const e6 = `
// SweetOf returns a Sweet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Sweet is returned.
func SweetOf(i int) Sweet {
	if 0 <= i && i < len(AllSweets) {
		return AllSweets[i]
	}
	// an invalid result
	return Mars + Bounty + Snickers + Kitkat
}
`

const e7 = `
// IsValid determines whether a Sweet is one of the defined constants.
func (i Sweet) IsValid() bool {
	switch i {
	case Mars, Bounty, Snickers, Kitkat:
		return true
	}
	return false
}
`

const e8nc = `
// Parse parses a string to find the corresponding Sweet, accepting either one of the string
// values or an ordinal number.
func (v *Sweet) Parse(s string) error {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(sweetEnumIndex); j++ {
		i1 := sweetEnumIndex[j]
		p := sweetEnumStrings[i0:i1]
		if s == p {
			*v = AllSweets[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Sweet")
}
`

const e8lc = `
// Parse parses a string to find the corresponding Sweet, accepting either one of the string
// values or an ordinal number.
// The case of s does not matter.
func (v *Sweet) Parse(s string) error {
	s = strings.ToLower(s)
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(sweetEnumIndex); j++ {
		i1 := sweetEnumIndex[j]
		p := sweetEnumStrings[i0:i1]
		if s == p {
			*v = AllSweets[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Sweet")
}
`

const e9nc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string
// values or an ordinal number.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const e9lc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string
// values or an ordinal number.
// The case of s does not matter.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const e10 = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Sweet) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

const e11 = `
// SweetMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var SweetMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// SweetMarshalJSONUsingString is true.
func (i Sweet) MarshalJSON() ([]byte, error) {
	if SweetMarshalJSONUsingString {
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
func (i *Sweet) UnmarshalJSON(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
`

func TestWriteFuncString(t *testing.T) {
	RegisterTestingT(t)
	buf := &bytes.Buffer{}
	m := model{
		mainType: "Sweet",
		baseType: "int",
		plural:   "Sweets",
	}
	m.writeFuncString(&printer{w: buf}, "sweetEnumStrings", "sweetEnumIndex")
	got := buf.String()
	Ω(got).Should(Equal(e4), got)
}

func TestWriteFuncOrdinal(t *testing.T) {
	RegisterTestingT(t)
	buf := &bytes.Buffer{}
	m := model{
		mainType: "Sweet",
		baseType: "int",
		plural:   "Sweets",
		values:   []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		xf:       nil,
	}
	m.writeFuncOrdinal(&printer{w: buf})
	got := buf.String()
	Ω(got).Should(Equal(e5), got)
}

func TestWriteNoChange(t *testing.T) {
	RegisterTestingT(t)
	buf := &bytes.Buffer{}
	m := model{
		mainType: "Sweet",
		baseType: "int",
		plural:   "Sweets",
		pkg:      "confectionary",
		values:   []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		xf:       nil,
	}
	err := m.write(buf)
	got := buf.String()
	Ω(err).Should(Not(HaveOccurred()))
	Ω(got).Should(Equal(e0+version+e1+`const sweetEnumStrings = "MarsBountySnickersKitkat"`+e3+e4+e5+e6+e7+e8nc+e9nc+e10+e11), got)
}

func TestWriteLower(t *testing.T) {
	RegisterTestingT(t)
	buf := &bytes.Buffer{}
	m := model{
		mainType: "Sweet",
		baseType: "int",
		plural:   "Sweets",
		pkg:      "confectionary",
		values:   []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		xf:       []Transform{ToLower},
	}
	err := m.write(buf)
	got := buf.String()
	Ω(err).Should(Not(HaveOccurred()))
	Ω(got).Should(Equal(e0+version+e1+`const sweetEnumStrings = "marsbountysnickerskitkat"`+e3+e4+e5+e6+e7+e8lc+e9lc+e10+e11), got)
}
