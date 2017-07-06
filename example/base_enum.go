// generated code - do not edit

package example

import (
	"errors"
	"fmt"
)

const baseEnumStrings = "ACGT"

var baseEnumIndex = [...]uint16{0, 1, 2, 3, 4}

var AllBases = []Base{A, C, G, T}

// String returns the string representation of a Base
func (i Base) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllBases) {
		return fmt.Sprintf("Base(%v)", i)
	}
	return baseEnumStrings[baseEnumIndex[o]:baseEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Base
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

// Parse parses a string to find the corresponding Base
func (v *Base) Parse(s string) error {
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

// AsBase parses a string to find the corresponding Base
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
