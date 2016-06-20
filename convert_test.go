package main

import (
	. "github.com/rickb777/terst"
	"testing"
	"bytes"
	"bufio"
)

const e1 = `// generated code - do not edit

package confectionary

import "fmt"

const sweetEnumStrings = "MarsSnickersKitkat"

var sweetEnumIndex = [...]uint16{0, 4, 12, 18}

func (i Sweet) String() string {
	if i < 0 || i >= Sweet(len(sweetEnumIndex)-1) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[i]:sweetEnumIndex[i+1]]
}

func AsSweet(s string) (Sweet, error) {
	i0 := 0
	for j := 1; j < len(sweetEnumIndex); j++ {
		i1 := sweetEnumIndex[j]
		p := sweetEnumStrings[i0:i1]
		if s == p {
			return Sweet(j-1), nil
		}
		i0 = i1
	}
	return Mars, errors.New(s + ": unrecognised Sweet")
}

`

func TestWrite(t *testing.T) {
	Terst(t)
	buf := &bytes.Buffer{}
	write(buf, "Sweet", "confectionary", []string{"Mars", "Snickers", "Kitkat"})
	got := buf.String()
	Is(got, e1)
}

const enum1 = `
const (
	Mars Sweet = iota
	Snickers
	Kitkat
)
`

const enum2 = `
const (
	Mars Irrelevant = iota
	Snickers
	Kitkat
)
`

func TestScanValuesHappy(t *testing.T) {
	Terst(t)
	s := bufio.NewScanner(bytes.NewBufferString(enum1))
	values := scanValues(s, "Sweet")
	Is(values, []string{"Mars", "Snickers", "Kitkat"})
}

func TestScanValuesIrrelevant(t *testing.T) {
	Terst(t)
	s := bufio.NewScanner(bytes.NewBufferString(enum2))
	values := scanValues(s, "Sweet")
	Is(values, nil)
}


