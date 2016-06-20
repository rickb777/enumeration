package main

import (
	. "github.com/rickb777/terst"
	"testing"
	"bytes"
)

const e1 = `// generated code - do not edit

package confectionary

import "fmt"

const sweetEnumStrings = "MarsSnickerKitkat"

var sweetEnumIndex = [...]uint16{0, 4, 11, 17}

func (i Sweet) String() string {
	if i < 0 || i >= Sweet(len(sweetEnumIndex)-1) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return sweetEnumStrings[sweetEnumIndex[i]:sweetEnumIndex[i+1]]
}

func AsSweet(s string) (Sweet, error) {
	i0 := sweetEnumIndex[0]
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
	write(buf, "Sweet", "confectionary", []string{"Mars", "Snicker", "Kitkat"})
	got := buf.String()
	Is(got, e1)
}


