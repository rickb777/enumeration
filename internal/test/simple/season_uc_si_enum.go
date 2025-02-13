// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.0-dirty

package simple

import (
	"database/sql/driver"
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
)

// AllSeason_Uc_Sis lists all 4 values in order.
var AllSeason_Uc_Sis = []Season_Uc_Si{
	Spring_Uc_Si, Summer_Uc_Si, Autumn_Uc_Si, Winter_Uc_Si,
}

// AllSeason_Uc_SiEnums lists all 4 values in order.
var AllSeason_Uc_SiEnums = enum.IntEnums{
	Spring_Uc_Si, Summer_Uc_Si, Autumn_Uc_Si, Winter_Uc_Si,
}

const (
	season_uc_siEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_siEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Si. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Si) Ordinal() int {
	switch v {
	case Spring_Uc_Si:
		return 0
	case Summer_Uc_Si:
		return 1
	case Autumn_Uc_Si:
		return 2
	case Winter_Uc_Si:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Si, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Si) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_siEnumStrings, season_uc_siEnumIndex[:])
}

func (v Season_Uc_Si) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Sis) {
		return fmt.Sprintf("Season_Uc_Si(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Si is one of the defined constants.
func (v Season_Uc_Si) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Si) Int() int {
	return int(v)
}

// Value converts the Season_Uc_Si to a string  (based on '-store identifier').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v Season_Uc_Si) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.String(), nil
}
