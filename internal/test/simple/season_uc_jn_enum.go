// generated code - do not edit
// github.com/rickb777/enumeration/v4 6fbd6b0a14258861d58a8efc36602c830d2f5fce-dirty

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
	"strconv"
)

// AllSeason_Uc_Jns lists all 4 values in order.
var AllSeason_Uc_Jns = []Season_Uc_Jn{
	Spring_Uc_Jn, Summer_Uc_Jn, Autumn_Uc_Jn, Winter_Uc_Jn,
}

// AllSeason_Uc_JnEnums lists all 4 values in order.
var AllSeason_Uc_JnEnums = enum.IntEnums{
	Spring_Uc_Jn, Summer_Uc_Jn, Autumn_Uc_Jn, Winter_Uc_Jn,
}

const (
	season_uc_jnEnumStrings = "SPRINGSUMMERAUTUMNWINTER"
)

var (
	season_uc_jnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Uc_Jn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Uc_Jn) Ordinal() int {
	switch v {
	case Spring_Uc_Jn:
		return 0
	case Summer_Uc_Jn:
		return 1
	case Autumn_Uc_Jn:
		return 2
	case Winter_Uc_Jn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Uc_Jn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Uc_Jn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_uc_jnEnumStrings, season_uc_jnEnumIndex[:])
}

func (v Season_Uc_Jn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Uc_Jns) {
		return fmt.Sprintf("Season_Uc_Jn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Uc_Jn is one of the defined constants.
func (v Season_Uc_Jn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Uc_Jn) Int() int {
	return int(v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The number representation is chosen according to -marshaljson.
func (v Season_Uc_Jn) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	s := season_uc_jnMarshalNumber(v)
	return []byte(s), nil
}

func (v Season_Uc_Jn) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Uc_Jn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_uc_jn", v)
}

// season_uc_jnMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_uc_jnMarshalNumber = func(v Season_Uc_Jn) string {
	return strconv.FormatInt(int64(v), 10)
}

func (v Season_Uc_Jn) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}
