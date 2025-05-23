// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-2-g23b0e6f-dirty

package simple

import (
	"fmt"
)

// AllSeason_Uc_Tss lists all 4 values in order.
var AllSeason_Uc_Tss = []Season_Uc_Ts{
	Spring_Uc_Tt, Summer_Uc_Tt, Autumn_Uc_Tt, Winter_Uc_Tt,
}

const (
	season_uc_tsEnumStrings = "SPRING_UC_TTSUMMER_UC_TTAUTUMN_UC_TTWINTER_UC_TT"
	season_uc_tsTextStrings = "SprgSumrAutmWint"
)

var (
	season_uc_tsEnumIndex = [...]uint16{0, 12, 24, 36, 48}
	season_uc_tsTextIndex = [...]uint16{0, 4, 8, 12, 16}
)

// String returns the literal string representation of a Season_Uc_Ts, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ts) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_tsEnumStrings, season_uc_tsEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Uc_Ts. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ts) Ordinal() int {
	switch v {
	case Spring_Uc_Tt:
		return 0
	case Summer_Uc_Tt:
		return 1
	case Autumn_Uc_Tt:
		return 2
	case Winter_Uc_Tt:
		return 3
	}
	return -1
}

func (v Season_Uc_Ts) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tss) {
		return fmt.Sprintf("Season_Uc_Ts(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Ts is one of the defined constants.
func (v Season_Uc_Ts) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ts) Int() int {
	return int(v)
}
