// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.1

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
)

// AllSeason_Uc_Jis lists all 4 values in order.
var AllSeason_Uc_Jis = []Season_Uc_Ji{
	Spring_Uc_Ji, Summer_Uc_Ji, Autumn_Uc_Ji, Winter_Uc_Ji,
}

// AllSeason_Uc_JiEnums lists all 4 values in order.
var AllSeason_Uc_JiEnums = enum.IntEnums{
	Spring_Uc_Ji, Summer_Uc_Ji, Autumn_Uc_Ji, Winter_Uc_Ji,
}

const (
	season_uc_jiEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_jiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Ji. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Ji) Ordinal() int {
	switch v {
	case Spring_Uc_Ji:
		return 0
	case Summer_Uc_Ji:
		return 1
	case Autumn_Uc_Ji:
		return 2
	case Winter_Uc_Ji:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Ji, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Ji) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:])
}

func (v Season_Uc_Ji) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jis) {
		return fmt.Sprintf("Season_Uc_Ji(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Ji is one of the defined constants.
func (v Season_Uc_Ji) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Ji) Int() int {
	return int(v)
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Uc_Ji) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:])
}

func (v Season_Uc_Ji) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Ji) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Ji) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_ji", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v Season_Uc_Ji) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_uc_jiEnumStrings, season_uc_jiEnumIndex[:])
	return enum.QuotedString(s), nil
}
