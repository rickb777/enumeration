package main

import (
	"bufio"
	"bytes"
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v2/transform"
	"testing"
)

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
	RegisterTestingT(t)
	s := bufio.NewScanner(bytes.NewBufferString(enum1))
	values := scanValues(s, "Sweet")
	Ω(values).Should(Equal([]string{"Mars", "Bounty", "Snickers", "Kitkat"}))
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
	RegisterTestingT(t)
	s := bufio.NewScanner(bytes.NewBufferString(enum2))
	values := scanValues(s, "Sweet")
	Ω(values).Should(BeNil())
}

// deeper edge-case testing for line processing
func TestRemoveCommentsAndSplitWords(t *testing.T) {
	RegisterTestingT(t)
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
		Ω(values).Should(Equal(c.expected))
	}
}

const enum3 = `
type IgnoreMe int
var s = "123"
type Sweet int // <-- buried here
type Transformer struct {
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
	RegisterTestingT(t)
	s := bytes.NewBufferString(enum3)
	m, err := convert(s, "filename.go", "Sweet", "Sweets", "confectionary", transform.Stet, true, true)
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model{
		MainType:    "Sweet",
		LcType:      "sweet",
		BaseType:    "int",
		Plural:      "Sweets",
		Pkg:         "confectionary",
		Version:     version,
		Values:      []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		IgnoreCase:  true,
		Unsnake:     true,
		Case:        0,
		S1:          "",
		S2:          "",
		LookupTable: "",
	}))
}

const enum4 = `
type Sweet int
const (
	Mars, Bounty, Snickers, Kitkat Sweet = 1, 2, 3, 4
)
`

func TestConvertHappy2(t *testing.T) {
	RegisterTestingT(t)
	s := bytes.NewBufferString(enum4)
	m, err := convert(s, "filename.go", "Sweet", "Sweets", "confectionary", transform.Upper, false, false)
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model{
		MainType:    "Sweet",
		LcType:      "sweet",
		BaseType:    "int",
		Plural:      "Sweets",
		Pkg:         "confectionary",
		Version:     version,
		Values:      []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		Case:        transform.Upper,
		S1:          "",
		S2:          "",
		LookupTable: "",
	}))
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
	RegisterTestingT(t)
	s := bytes.NewBufferString(enum5)
	_, err := convert(s, "filename.go", "Sweet", "Sweets", "confectionary", transform.Stet, false, false)
	Ω(err.Error()).Should(Equal("Failed to find Sweet in filename.go"))
}
