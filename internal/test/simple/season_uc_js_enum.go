// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.0-dirty

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
)

// AllSeason_Uc_Jss lists all 4 values in order.
var AllSeason_Uc_Jss = []Season_Uc_Js{
	Spring_Uc_Js, Summer_Uc_Js, Autumn_Uc_Js, Winter_Uc_Js,
}

// AllSeason_Uc_JsEnums lists all 4 values in order.
var AllSeason_Uc_JsEnums = enum.IntEnums{
	Spring_Uc_Js, Summer_Uc_Js, Autumn_Uc_Js, Winter_Uc_Js,
}

const (
	season_uc_jsEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
	season_uc_jsJSONStrings = "SprgSumrAutmWint"
)

var (
	season_uc_jsEnumIndex = [...]uint16{0, 6, 12, 18, 24}
	season_uc_jsJSONIndex = [...]uint16{0, 4, 8, 12, 16}
)

// Ordinal returns the ordinal number of a Season_Uc_Js. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Js) Ordinal() int {
	switch v {
	case Spring_Uc_Js:
		return 0
	case Summer_Uc_Js:
		return 1
	case Autumn_Uc_Js:
		return 2
	case Winter_Uc_Js:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Js, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Js) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jsEnumStrings, season_uc_jsEnumIndex[:])
}

func (v Season_Uc_Js) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jss) {
		return fmt.Sprintf("Season_Uc_Js(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Js is one of the defined constants.
func (v Season_Uc_Js) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Js) Int() int {
	return int(v)
}

// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v Season_Uc_Js) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, season_uc_jsJSONStrings, season_uc_jsJSONIndex[:])
}

func (v Season_Uc_Js) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Uc_Js) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Js) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_js", v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v Season_Uc_Js) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, season_uc_jsJSONStrings, season_uc_jsJSONIndex[:])
	return enum.QuotedString(s), nil
}
