// generated code - do not edit
// github.com/rickb777/enumeration/v4 6fbd6b0a14258861d58a8efc36602c830d2f5fce-dirty

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
)

// AllSeason_Uc_Tis lists all 4 values in order.
var AllSeason_Uc_Tis = []Season_Uc_Ti{
	Spring_Uc_Ti, Summer_Uc_Ti, Autumn_Uc_Ti, Winter_Uc_Ti,
}

// AllSeason_Uc_TiEnums lists all 4 values in order.
var AllSeason_Uc_TiEnums = enum.IntEnums{
	Spring_Uc_Ti, Summer_Uc_Ti, Autumn_Uc_Ti, Winter_Uc_Ti,
}

const (
	season_uc_tiEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_tiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Ti. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ti) Ordinal() int {
	switch v {
	case Spring_Uc_Ti:
		return 0
	case Summer_Uc_Ti:
		return 1
	case Autumn_Uc_Ti:
		return 2
	case Winter_Uc_Ti:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Ti, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ti) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_tiEnumStrings, season_uc_tiEnumIndex[:])
}

func (v Season_Uc_Ti) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Tis) {
		return fmt.Sprintf("Season_Uc_Ti(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Ti is one of the defined constants.
func (v Season_Uc_Ti) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ti) Int() int {
	return int(v)
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Uc_Ti) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v Season_Uc_Ti) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v Season_Uc_Ti) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, season_uc_tiEnumStrings, season_uc_tiEnumIndex[:]), nil
}

func (v Season_Uc_Ti) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Ti) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Ti) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_ti", v)
}
