// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-1-ga50534c

package simple

import (
	"fmt"
)

// AllSeason_Nc_Sis lists all 4 values in order.
var AllSeason_Nc_Sis = []Season_Nc_Si{
	Spring_Nc_Si, Summer_Nc_Si, Autumn_Nc_Si, Winter_Nc_Si,
}

const (
	season_nc_siEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_siEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Nc_Si, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Si) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_siEnumStrings, season_nc_siEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Nc_Si. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Si) Ordinal() int {
	switch v {
	case Spring_Nc_Si:
		return 0
	case Summer_Nc_Si:
		return 1
	case Autumn_Nc_Si:
		return 2
	case Winter_Nc_Si:
		return 3
	}
	return -1
}

func (v Season_Nc_Si) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Sis) {
		return fmt.Sprintf("Season_Nc_Si(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Si is one of the defined constants.
func (v Season_Nc_Si) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Si) Int() int {
	return int(v)
}
