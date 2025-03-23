// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-1-ga50534c

package simple

import (
	"fmt"
)

// AllSeason_Nc_Tss lists all 4 values in order.
var AllSeason_Nc_Tss = []Season_Nc_Ts{
	Spring_Nc_Tt, Summer_Nc_Tt, Autumn_Nc_Tt, Winter_Nc_Tt,
}

const (
	season_nc_tsEnumStrings = "Spring_Nc_TtSummer_Nc_TtAutumn_Nc_TtWinter_Nc_Tt"
	season_nc_tsTextStrings = "SprgSumrAutmWint"
)

var (
	season_nc_tsEnumIndex = [...]uint16{0, 12, 24, 36, 48}
	season_nc_tsTextIndex = [...]uint16{0, 4, 8, 12, 16}
)

// String returns the literal string representation of a Season_Nc_Ts, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ts) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_tsEnumStrings, season_nc_tsEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Nc_Ts. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ts) Ordinal() int {
	switch v {
	case Spring_Nc_Tt:
		return 0
	case Summer_Nc_Tt:
		return 1
	case Autumn_Nc_Tt:
		return 2
	case Winter_Nc_Tt:
		return 3
	}
	return -1
}

func (v Season_Nc_Ts) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Tss) {
		return fmt.Sprintf("Season_Nc_Ts(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Ts is one of the defined constants.
func (v Season_Nc_Ts) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ts) Int() int {
	return int(v)
}
