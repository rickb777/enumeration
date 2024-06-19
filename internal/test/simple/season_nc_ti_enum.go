// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.4.0

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
)

// AllSeason_Nc_Tis lists all 4 values in order.
var AllSeason_Nc_Tis = []Season_Nc_Ti{
	Spring_Nc_Ti, Summer_Nc_Ti, Autumn_Nc_Ti, Winter_Nc_Ti,
}

// AllSeason_Nc_TiEnums lists all 4 values in order.
var AllSeason_Nc_TiEnums = enum.IntEnums{
	Spring_Nc_Ti, Summer_Nc_Ti, Autumn_Nc_Ti, Winter_Nc_Ti,
}

const (
	season_nc_tiEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_tiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Ti. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ti) Ordinal() int {
	switch v {
	case Spring_Nc_Ti:
		return 0
	case Summer_Nc_Ti:
		return 1
	case Autumn_Nc_Ti:
		return 2
	case Winter_Nc_Ti:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Ti, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ti) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_tiEnumStrings, season_nc_tiEnumIndex[:])
}

func (v Season_Nc_Ti) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Tis) {
		return fmt.Sprintf("Season_Nc_Ti(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Ti is one of the defined constants.
func (v Season_Nc_Ti) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ti) Int() int {
	return int(v)
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Nc_Ti) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Nc_Ti) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v Season_Nc_Ti) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_nc_tiEnumStrings, season_nc_tiEnumIndex[:]), nil
}

func (v Season_Nc_Ti) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Ti) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Ti) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ti", v)
}
