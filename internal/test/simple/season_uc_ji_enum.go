// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-dirty

package simple

import (
	"fmt"
)

// AllSeason_Uc_Jis lists all 4 values in order.
var AllSeason_Uc_Jis = []Season_Uc_Ji{
	Spring_Uc_Ji, Summer_Uc_Ji, Autumn_Uc_Ji, Winter_Uc_Ji,
}

const (
	season_uc_jiEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_jiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Ji, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ji) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Uc_Ji. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ji) Ordinal() int {
	switch v {
	case Spring_Uc_Ji:
		return 0
	case Summer_Uc_Ji:
		return 1
	case Autumn_Uc_Ji:
		return 2
	case Winter_Uc_Ji:
		return 3
	}
	return -1
}

func (v Season_Uc_Ji) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jis) {
		return fmt.Sprintf("Season_Uc_Ji(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Ji is one of the defined constants.
func (v Season_Uc_Ji) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ji) Int() int {
	return int(v)
}
