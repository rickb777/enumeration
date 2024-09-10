// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.0

package simple

import (
	"database/sql/driver"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
)

// AllSeason_Uc_Sns lists all 4 values in order.
var AllSeason_Uc_Sns = []Season_Uc_Sn{
	Spring_Uc_Sn, Summer_Uc_Sn, Autumn_Uc_Sn, Winter_Uc_Sn,
}

// AllSeason_Uc_SnEnums lists all 4 values in order.
var AllSeason_Uc_SnEnums = enum.IntEnums{
	Spring_Uc_Sn, Summer_Uc_Sn, Autumn_Uc_Sn, Winter_Uc_Sn,
}

const (
	season_uc_snEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_snEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Sn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Sn) Ordinal() int {
	switch v {
	case Spring_Uc_Sn:
		return 0
	case Summer_Uc_Sn:
		return 1
	case Autumn_Uc_Sn:
		return 2
	case Winter_Uc_Sn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Sn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Sn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_snEnumStrings, season_uc_snEnumIndex[:])
}

func (v Season_Uc_Sn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Sns) {
		return fmt.Sprintf("Season_Uc_Sn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Sn is one of the defined constants.
func (v Season_Uc_Sn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Sn) Int() int {
	return int(v)
}

// Value converts the Season_Uc_Sn to a number (based on '-store number').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Uc_Sn) Value() (driver.Value, error) {
	return int64(v), nil
}
