// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.1

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
)

// AllSeason_Nc_Tss lists all 4 values in order.
var AllSeason_Nc_Tss = []Season_Nc_Ts{
	Spring_Nc_Tt, Summer_Nc_Tt, Autumn_Nc_Tt, Winter_Nc_Tt,
}

// AllSeason_Nc_TsEnums lists all 4 values in order.
var AllSeason_Nc_TsEnums = enum.IntEnums{
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

// String returns the literal string representation of a Season_Nc_Ts, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ts) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_tsEnumStrings, season_nc_tsEnumIndex[:])
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

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Nc_Ts) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Nc_Ts) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v Season_Nc_Ts) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_nc_tsTextStrings, season_nc_tsTextIndex[:]), nil
}

func (v Season_Nc_Ts) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Ts) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Ts) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ts", v)
}
