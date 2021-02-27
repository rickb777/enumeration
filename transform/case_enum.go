// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.0.2

package transform

import (
	"fmt"
)

const (
	caseEnumStrings = "StetUpperLower"
)

var caseEnumIndex = [...]uint16{0, 4, 9, 14}

// AllCases lists all 3 values in order.
var AllCases = []Case{
	Stet, Upper, Lower,
}

// String returns the literal string representation of a Case, which is
// the same as the const identifier.
func (i Case) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllCases) {
		return fmt.Sprintf("Case(%d)", i)
	}
	return caseEnumStrings[caseEnumIndex[o]:caseEnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a Case.
func (i Case) Ordinal() int {
	switch i {
	case Stet:
		return 0
	case Upper:
		return 1
	case Lower:
		return 2
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Case) Int() int {
	return int(i)
}
