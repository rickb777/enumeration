// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.5.0

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
)

// AllSeason_Nc_Jns lists all 4 values in order.
var AllSeason_Nc_Jns = []Season_Nc_Jn{
	Spring_Nc_Jn, Summer_Nc_Jn, Autumn_Nc_Jn, Winter_Nc_Jn,
}

// AllSeason_Nc_JnEnums lists all 4 values in order.
var AllSeason_Nc_JnEnums = enum.IntEnums{
	Spring_Nc_Jn, Summer_Nc_Jn, Autumn_Nc_Jn, Winter_Nc_Jn,
}

const (
	season_nc_jnEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_jnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Jn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Jn) Ordinal() int {
	switch v {
	case Spring_Nc_Jn:
		return 0
	case Summer_Nc_Jn:
		return 1
	case Autumn_Nc_Jn:
		return 2
	case Winter_Nc_Jn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Jn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Jn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_jnEnumStrings, season_nc_jnEnumIndex[:])
}

func (v Season_Nc_Jn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Jns) {
		return fmt.Sprintf("Season_Nc_Jn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Jn is one of the defined constants.
func (v Season_Nc_Jn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Jn) Int() int {
	return int(v)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The number representation is chosen according to -marshaljson.
func (v Season_Nc_Jn) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	s := season_nc_jnMarshalNumber(v)
	return []byte(s), nil
}

func (v Season_Nc_Jn) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Jn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_jn", v)
}

// season_nc_jnMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_nc_jnMarshalNumber = func(v Season_Nc_Jn) string {
	return strconv.FormatInt(int64(v), 10)
}

func (v Season_Nc_Jn) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}
