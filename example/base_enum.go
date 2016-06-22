// generated code - do not edit

package example

import (
	"errors"
	"fmt"
)

const baseEnumStrings = "ACTG"

var baseEnumIndex = [...]uint16{0, 1, 2, 3, 4}

var AllBases = []Base{A, C, T, G}

// String returns the string representation of a Base
func (i Base) String() string {
	if i < 0 || i >= Base(len(baseEnumIndex)-1) {
		return fmt.Sprintf("Base(%d)", i)
	}
	return baseEnumStrings[baseEnumIndex[i]:baseEnumIndex[i+1]]
}

// Ordinal returns the ordinal number of a Base
func (i Base) Ordinal() int {
	switch i {
	case A:
		return 0
	case C:
		return 1
	case T:
		return 2
	case G:
		return 3
	}
	panic(fmt.Errorf("%d: unknown Base", i))
}

// Parse parses a string to find the corresponding Base
func (v *Base) Parse(s string) error {
	var i0 uint16 = 0
	for j := 1; j < len(baseEnumIndex); j++ {
		i1 := baseEnumIndex[j]
		p := baseEnumStrings[i0:i1]
		if s == p {
			*v = Base(j - 1)
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised Base")
}

// AsBase parses a string to find the corresponding Base
func AsBase(s string) (Base, error) {
	var i = new(Base)
	err := i.Parse(s)
	return *i, err
}
