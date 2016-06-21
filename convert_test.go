package main

import (
	"bufio"
	"bytes"
	"fmt"
	. "github.com/rickb777/terst"
	"testing"
)

const e1 = `// generated code - do not edit

package confectionary

import (
	"errors"
	"fmt"
)

const sweetEnumStrings = "MarsSnickersKitkat"

var sweetEnumIndex = [...]uint16{0, 4, 12, 18}

var AllSweets = []Sweet{Mars, Snickers, Kitkat}

// String returns the string representation of a Sweet
func (i Sweet) String() string {
	if i < 0 || i >= Sweet(len(sweetEnumIndex)-1) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[i]:sweetEnumIndex[i+1]]
}

// Ordinal returns the ordinal number of a Sweet
func (i Sweet) Ordinal() int {
	switch i {
	case Mars:
		return 0
	case Snickers:
		return 1
	case Kitkat:
		return 2
	}
	panic(fmt.Errorf("%d: unknown Sweet", i))
}

// AsSweet parses a string to find the corresponding Sweet
func AsSweet(s string) (Sweet, error) {
	var i0 uint16 = 0
	for j := 1; j < len(sweetEnumIndex); j++ {
		i1 := sweetEnumIndex[j]
		p := sweetEnumStrings[i0:i1]
		if s == p {
			return Sweet(j - 1), nil
		}
		i0 = i1
	}
	return Mars, errors.New(s + ": unrecognised Sweet")
}
`

func TestWrite(t *testing.T) {
	Terst(t)
	buf := &bytes.Buffer{}
	write(buf, "Sweet", "Sweets", "confectionary", []string{"Mars", "Snickers", "Kitkat"})
	got := buf.String()
	strEq(t, got, e1)
}

const enum1 = `
// inline comments are allowed
const (
	Mars Sweet = iota
	Snickers // I need this

	// yum yum
	Kitkat
)
// as are blank lines
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
