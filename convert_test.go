package main

import (
	"bufio"
	"bytes"
	"fmt"
	. "github.com/rickb777/terst"
	"strings"
	"testing"
)

const e1 = `// generated code - do not edit

package confectionary

import (
	"errors"
	"fmt"
)

`
const e2 = `

var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24}

var AllSweets = []Sweet{Mars, Bounty, Snickers, Kitkat}

// String returns the string representation of a Sweet
func (i Sweet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%v)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[o]:sweetEnumIndex[o+1]]
}

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

// Parse parses a string to find the corresponding Sweet
func (v *Sweet) Parse(s string) error {
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

// AsSweet parses a string to find the corresponding Sweet
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i Sweet) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

func TestWriteUpper(t *testing.T) {
	Terst(t)
	buf := &bytes.Buffer{}
	write(buf, "Sweet", "int", "Sweets", "confectionary", []string{"Mars", "Bounty", "Snickers", "Kitkat"}, noop)
	got := buf.String()
	strEq(t, got, e1+`const sweetEnumStrings = "MarsBountySnickersKitkat"`+e2)
}

func TestWriteLower(t *testing.T) {
	Terst(t)
	buf := &bytes.Buffer{}
	write(buf, "Sweet", "int", "Sweets", "confectionary", []string{"Mars", "Bounty", "Snickers", "Kitkat"}, strings.ToLower)
	got := buf.String()
	strEq(t, got, e1+`const sweetEnumStrings = "marsbountysnickerskitkat"`+e2)
}

const enum1 = `
// inline comments are allowed
const (
	Mars Sweet = iota
	Bounty    // coconuts and more
	Snickers  // I need this

	// yum yum
	Kitkat
)
// as are blank lines
`

func TestScanValuesHappy(t *testing.T) {
	Terst(t)
	s := bufio.NewScanner(bytes.NewBufferString(enum1))
	values := scanValues(s, "Sweet")
	Is(values, []string{"Mars", "Bounty", "Snickers", "Kitkat"})
}

const enum2 = `
const (
	Mars Irrelevant = iota
	Bounty
	Snickers
	Kitkat
)
`

func TestScanValuesIrrelevant(t *testing.T) {
	Terst(t)
	s := bufio.NewScanner(bytes.NewBufferString(enum2))
	values := scanValues(s, "Sweet")
	Is(values, nil)
}

func strEq(t *testing.T, want, got string) {
	if want != got {
		wl := len(want)
		gl := len(got)
		ll := wl
		if ll > gl {
			ll = gl
		}
		same := true
		i := 0
		for ; i < ll; i++ {
			if want[i] != got[i] && want[i] != ' ' {
				if same {
					same = false
					fmt.Printf("<<[")
				}
			} else {
				if !same {
					same = true
					fmt.Printf("]>>")
				}
			}
			fmt.Printf("%c", want[i])
		}
		for ; i < wl; i++ {
			fmt.Printf("<<[")
			fmt.Printf("%c", want[i])
			fmt.Printf("]>>")
		}
		for ; i < gl; i++ {
			fmt.Printf("<<#")
			fmt.Printf("%c", got[i])
			fmt.Printf("#>>")
		}
		fmt.Println("")
		t.Fail()
	}
}

// deeper edge-case testing for line processing
func TestRemoveCommentsAndSplitWords(t *testing.T) {
	Terst(t)
	cases := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"one", []string{"one"}},
		{"one two", []string{"one", "two"}},
		{"one // two", []string{"one"}},
		{"\tNotice\t          // announcements to all // old comment", []string{"Notice"}},
	}
	for _, c := range cases {
		values := removeCommentsAndSplitWords(c.input)
		Is(values, c.expected)
	}
}

const enum3 = `
type IgnoreMe int
var s = "123"
type Sweet int // <-- buried here
type Foo struct {
	V int
}
type Bar interface {
	X()
}
var x = 0
const (
	Jam IgnoreMe = iota
	Toast
	Butter
)
var y = 1
const (
	Mars Sweet = iota
	Bounty
	Snickers
	Kitkat
)
`

func TestConvertHappy1(t *testing.T) {
	Terst(t)
	w := &bytes.Buffer{}
	s := bytes.NewBufferString(enum3)
	err := convert(w, s, "filename.go", "Sweet", "Sweets", "confectionary", noop)
	Is(err, nil)
	strEq(t, w.String(), e1+`const sweetEnumStrings = "MarsBountySnickersKitkat"`+e2)
}

const enum4 = `
type Sweet int
const (
	Mars, Bounty, Snickers, Kitkat Sweet = 1, 2, 3, 4
)
`

func TestConvertHappy2(t *testing.T) {
	Terst(t)
	w := &bytes.Buffer{}
	s := bytes.NewBufferString(enum4)
	err := convert(w, s, "filename.go", "Sweet", "Sweets", "confectionary", noop)
	Is(err, nil)
	strEq(t, w.String(), e1+`const sweetEnumStrings = "MarsBountySnickersKitkat"`+e2)
}

const enum5 = `
type IgnoreMe int
const (
	Mars Sweet = iota
	Bounty
	Snickers
	Kitkat
)
const (
	Jam IgnoreMe = iota
	Toast
	Butter
)
`

func TestConvertError(t *testing.T) {
	Terst(t)
	w := &bytes.Buffer{}
	s := bytes.NewBufferString(enum5)
	err := convert(w, s, "filename.go", "Sweet", "Sweets", "confectionary", noop)
	Is(err.Error(), "Failed to find Sweet in filename.go")
}
