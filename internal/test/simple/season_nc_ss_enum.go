// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-1-ga50534c

package simple

import (
	"fmt"
)

// AllSeason_Nc_Sss lists all 4 values in order.
var AllSeason_Nc_Sss = []Season_Nc_Ss{
	Spring_Nc_Ss, Summer_Nc_Ss, Autumn_Nc_Ss, Winter_Nc_Ss,
}

const (
	season_nc_ssEnumStrings = "SpringSummerAutumnWinter"
	season_nc_ssSQLStrings  = "SprgSumrAutmWint"
)

var (
	season_nc_ssEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_nc_ssSQLIndex  = [...]uint16{0, 4, 8, 12, 16}
)

// String returns the literal string representation of a Season_Nc_Ss, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ss) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_ssEnumStrings, season_nc_ssEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Nc_Ss. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ss) Ordinal() int {
	switch v {
	case Spring_Nc_Ss:
		return 0
	case Summer_Nc_Ss:
		return 1
	case Autumn_Nc_Ss:
		return 2
	case Winter_Nc_Ss:
		return 3
	}
	return -1
}

func (v Season_Nc_Ss) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Sss) {
		return fmt.Sprintf("Season_Nc_Ss(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Ss is one of the defined constants.
func (v Season_Nc_Ss) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ss) Int() int {
	return int(v)
}
