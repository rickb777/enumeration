// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.6

package simple

import (
	"fmt"
)

// AllSeason_Uc_Tns lists all 4 values in order.
var AllSeason_Uc_Tns = []Season_Uc_Tn{
	Spring_Uc_Tn, Summer_Uc_Tn, Autumn_Uc_Tn, Winter_Uc_Tn,
}

const (
	season_uc_tnEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_tnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Tn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Tn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_tnEnumStrings, season_uc_tnEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Uc_Tn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Tn) Ordinal() int {
	switch v {
	case Spring_Uc_Tn:
		return 0
	case Summer_Uc_Tn:
		return 1
	case Autumn_Uc_Tn:
		return 2
	case Winter_Uc_Tn:
		return 3
	}
	return -1
}

func (v Season_Uc_Tn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tns) {
		return fmt.Sprintf("Season_Uc_Tn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Tn is one of the defined constants.
func (v Season_Uc_Tn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Tn) Int() int {
	return int(v)
}
