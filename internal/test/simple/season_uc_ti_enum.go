// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-2-g23b0e6f-dirty

package simple

import (
	"fmt"
)

// AllSeason_Uc_Tis lists all 4 values in order.
var AllSeason_Uc_Tis = []Season_Uc_Ti{
	Spring_Uc_Ti, Summer_Uc_Ti, Autumn_Uc_Ti, Winter_Uc_Ti,
}

const (
	season_uc_tiEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_tiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Ti, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ti) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_tiEnumStrings, season_uc_tiEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Uc_Ti. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ti) Ordinal() int {
	switch v {
	case Spring_Uc_Ti:
		return 0
	case Summer_Uc_Ti:
		return 1
	case Autumn_Uc_Ti:
		return 2
	case Winter_Uc_Ti:
		return 3
	}
	return -1
}

func (v Season_Uc_Ti) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tis) {
		return fmt.Sprintf("Season_Uc_Ti(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Ti is one of the defined constants.
func (v Season_Uc_Ti) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ti) Int() int {
	return int(v)
}
