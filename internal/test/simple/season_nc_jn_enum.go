// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.2-2-g23b0e6f-dirty

package simple

import (
	"fmt"
)

// AllSeason_Nc_Jns lists all 4 values in order.
var AllSeason_Nc_Jns = []Season_Nc_Jn{
	Spring_Nc_Jn, Summer_Nc_Jn, Autumn_Nc_Jn, Winter_Nc_Jn,
}

const (
	season_nc_jnEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_jnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// String returns the literal string representation of a Season_Nc_Jn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Jn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_jnEnumStrings, season_nc_jnEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Nc_Jn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Jn) Ordinal() int {
	switch v {
	case Spring_Nc_Jn:
		return 0
	case Summer_Nc_Jn:
		return 1
	case Autumn_Nc_Jn:
		return 2
	case Winter_Nc_Jn:
		return 3
	}
	return -1
}

func (v Season_Nc_Jn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Jns) {
		return fmt.Sprintf("Season_Nc_Jn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Jn is one of the defined constants.
func (v Season_Nc_Jn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Jn) Int() int {
	return int(v)
}
