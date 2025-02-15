// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.0-dirty

package simple

import (
	"fmt"
)

// AllSeason_Uc_Jns lists all 4 values in order.
var AllSeason_Uc_Jns = []Season_Uc_Jn{
	Spring_Uc_Jn, Summer_Uc_Jn, Autumn_Uc_Jn, Winter_Uc_Jn,
}

const (
	season_uc_jnEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_jnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Uc_Jn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Jn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jnEnumStrings, season_uc_jnEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Uc_Jn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Jn) Ordinal() int {
	switch v {
	case Spring_Uc_Jn:
		return 0
	case Summer_Uc_Jn:
		return 1
	case Autumn_Uc_Jn:
		return 2
	case Winter_Uc_Jn:
		return 3
	}
	return -1
}

func (v Season_Uc_Jn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jns) {
		return fmt.Sprintf("Season_Uc_Jn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Jn is one of the defined constants.
func (v Season_Uc_Jn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Jn) Int() int {
	return int(v)
}
