package main

import (
	"bufio"
	"bytes"
	. "github.com/onsi/gomega"
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
	RegisterTestingT(t)
	w := &bytes.Buffer{}
	s := bytes.NewBufferString(enum3)
	err := convert(w, s, "filename.go", "Sweet", "Sweets", "confectionary", noop)
	Ω(err).Should(BeNil())
	str := w.String()
	Ω(str).Should(ContainSubstring(`const sweetEnumStrings = "MarsBountySnickersKitkat"`), str)
	Ω(str).Should(ContainSubstring(`var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24}`), str)
	Ω(str).Should(ContainSubstring(`var AllSweets = []Sweet{Mars, Bounty, Snickers, Kitkat}`), str)
}

const enum4 = `
type Sweet int
const (
	Mars, Bounty, Snickers, Kitkat Sweet = 1, 2, 3, 4
)
`

func TestConvertHappy2(t *testing.T) {
	RegisterTestingT(t)
	w := &bytes.Buffer{}
	s := bytes.NewBufferString(enum4)
	err := convert(w, s, "filename.go", "Sweet", "Sweets", "confectionary", noop)
	Ω(err).Should(BeNil())
	str := w.String()
	Ω(str).Should(ContainSubstring(`const sweetEnumStrings = "MarsBountySnickersKitkat"`), str)
	Ω(str).Should(ContainSubstring(`var sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24}`), str)
	Ω(str).Should(ContainSubstring(`var AllSweets = []Sweet{Mars, Bounty, Snickers, Kitkat}`), str)
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
	w := &bytes.Buffer{}
	s := bytes.NewBufferString(enum5)
	err := convert(w, s, "filename.go", "Sweet", "Sweets", "confectionary", noop)
	Ω(err.Error()).Should(Equal("Failed to find Sweet in filename.go"))
}
