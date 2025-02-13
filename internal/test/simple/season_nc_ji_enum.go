// generated code - do not edit
// github.com/rickb777/enumeration/v4 6fbd6b0a14258861d58a8efc36602c830d2f5fce-dirty

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
)

// AllSeason_Nc_Jis lists all 4 values in order.
var AllSeason_Nc_Jis = []Season_Nc_Ji{
	Spring_Nc_Ji, Summer_Nc_Ji, Autumn_Nc_Ji, Winter_Nc_Ji,
}

// AllSeason_Nc_JiEnums lists all 4 values in order.
var AllSeason_Nc_JiEnums = enum.IntEnums{
	Spring_Nc_Ji, Summer_Nc_Ji, Autumn_Nc_Ji, Winter_Nc_Ji,
}

const (
	season_nc_jiEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_jiEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Ji. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Ji) Ordinal() int {
	switch v {
	case Spring_Nc_Ji:
		return 0
	case Summer_Nc_Ji:
		return 1
	case Autumn_Nc_Ji:
		return 2
	case Winter_Nc_Ji:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Ji, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Ji) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:])
}

func (v Season_Nc_Ji) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Jis) {
		return fmt.Sprintf("Season_Nc_Ji(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Ji is one of the defined constants.
func (v Season_Nc_Ji) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Ji) Int() int {
	return int(v)
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Nc_Ji) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:])
}

func (v Season_Nc_Ji) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Ji) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Ji) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_ji", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v Season_Nc_Ji) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_nc_jiEnumStrings, season_nc_jiEnumIndex[:])
	return enum.QuotedString(s), nil
}
