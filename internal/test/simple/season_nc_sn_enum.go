// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.1

package simple

import (
	"database/sql/driver"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
)

// AllSeason_Nc_Sns lists all 4 values in order.
var AllSeason_Nc_Sns = []Season_Nc_Sn{
	Spring_Nc_Sn, Summer_Nc_Sn, Autumn_Nc_Sn, Winter_Nc_Sn,
}

// AllSeason_Nc_SnEnums lists all 4 values in order.
var AllSeason_Nc_SnEnums = enum.IntEnums{
	Spring_Nc_Sn, Summer_Nc_Sn, Autumn_Nc_Sn, Winter_Nc_Sn,
}

const (
	season_nc_snEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_snEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Sn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Sn) Ordinal() int {
	switch v {
	case Spring_Nc_Sn:
		return 0
	case Summer_Nc_Sn:
		return 1
	case Autumn_Nc_Sn:
		return 2
	case Winter_Nc_Sn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Sn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Sn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_snEnumStrings, season_nc_snEnumIndex[:])
}

func (v Season_Nc_Sn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Sns) {
		return fmt.Sprintf("Season_Nc_Sn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Sn is one of the defined constants.
func (v Season_Nc_Sn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Sn) Int() int {
	return int(v)
}

// Value converts the Season_Nc_Sn to a number (based on '-store number').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Nc_Sn) Value() (driver.Value, error) {
	return int64(v), nil
}
