package main

import (
	"bytes"
	. "github.com/onsi/gomega"
	"strings"
	"testing"
)

const e1 = `// generated code - do not edit

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

var AllSweets = []Sweet{Mars, Bounty, Snickers, Kitkat}
`

const e4 = `
// String returns the string representation of a Sweet
func (i Sweet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[o]:sweetEnumIndex[o+1]]
}
`

const e5 = `
// Ordinal returns the ordinal number of a Sweet
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

const e7 = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string
// values or an ordinal number.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const e8 = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Sweet) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

const e9 = `
// MarshalJSON converts values to ordinals suitable for transmission via JSON.
func (i Sweet) MarshalJSON() ([]byte, error) {
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
	writeFuncString(buf, "Sweet", "Sweets", "sweetEnumStrings", "sweetEnumIndex")
	got := buf.String()
	Ω(got).Should(Equal(e4), got)
}

func TestWriteFuncOrdinal(t *testing.T) {
	RegisterTestingT(t)
	buf := &bytes.Buffer{}
	writeFuncOrdinal(buf, "Sweet", []string{"Mars", "Bounty", "Snickers", "Kitkat"})
	got := buf.String()
	Ω(got).Should(Equal(e5), got)
}

func TestWriteUpper(t *testing.T) {
	RegisterTestingT(t)
	buf := &bytes.Buffer{}
	write(buf, "Sweet", "int", "Sweets", "confectionary", []string{"Mars", "Bounty", "Snickers", "Kitkat"}, noop)
	got := buf.String()
	Ω(got).Should(Equal(e1+`const sweetEnumStrings = "MarsBountySnickersKitkat"`+e3+e4+e5+e6+e7+e8+e9), got)
}

func TestWriteLower(t *testing.T) {
	RegisterTestingT(t)
	buf := &bytes.Buffer{}
	write(buf, "Sweet", "int", "Sweets", "confectionary", []string{"Mars", "Bounty", "Snickers", "Kitkat"}, strings.ToLower)
	got := buf.String()
	Ω(got).Should(Equal(e1+`const sweetEnumStrings = "marsbountysnickerskitkat"`+e3+e4+e5+e6+e7+e8+e9), got)
}
