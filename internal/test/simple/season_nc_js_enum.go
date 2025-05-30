// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.6

package simple

import (
	"fmt"
)

// AllSeason_Nc_Jss lists all 4 values in order.
var AllSeason_Nc_Jss = []Season_Nc_Js{
	Spring_Nc_Js, Summer_Nc_Js, Autumn_Nc_Js, Winter_Nc_Js,
}

const (
	season_nc_jsEnumStrings = "SpringSummerAutumnWinter"
	season_nc_jsJSONStrings = "SprgSumrAutmWint"
)

var (
	season_nc_jsEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_nc_jsJSONIndex = [...]uint16{0, 4, 8, 12, 16}
)

// String returns the literal string representation of a Season_Nc_Js, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Js) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_jsEnumStrings, season_nc_jsEnumIndex[:])
}

// Ordinal returns the ordinal number of a Season_Nc_Js. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Js) Ordinal() int {
	switch v {
	case Spring_Nc_Js:
		return 0
	case Summer_Nc_Js:
		return 1
	case Autumn_Nc_Js:
		return 2
	case Winter_Nc_Js:
		return 3
	}
	return -1
}

func (v Season_Nc_Js) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Jss) {
		return fmt.Sprintf("Season_Nc_Js(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Js is one of the defined constants.
func (v Season_Nc_Js) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Js) Int() int {
	return int(v)
}
