// generated code - do not edit
// github.com/rickb777/enumeration/v4 v4.0.0-dirty

package simple

import (
	"fmt"
	"github.com/rickb777/enumeration/v4/enum"
	"strconv"
)

// AllSeason_Nc_Tns lists all 4 values in order.
var AllSeason_Nc_Tns = []Season_Nc_Tn{
	Spring_Nc_Tn, Summer_Nc_Tn, Autumn_Nc_Tn, Winter_Nc_Tn,
}

// AllSeason_Nc_TnEnums lists all 4 values in order.
var AllSeason_Nc_TnEnums = enum.IntEnums{
	Spring_Nc_Tn, Summer_Nc_Tn, Autumn_Nc_Tn, Winter_Nc_Tn,
}

const (
	season_nc_tnEnumStrings = "SpringSummerAutumnWinter"
)

var (
	season_nc_tnEnumIndex = [...]uint16{0, 6, 12, 18, 24}
)

// Ordinal returns the ordinal number of a Season_Nc_Tn. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v Season_Nc_Tn) Ordinal() int {
	switch v {
	case Spring_Nc_Tn:
		return 0
	case Summer_Nc_Tn:
		return 1
	case Autumn_Nc_Tn:
		return 2
	case Winter_Nc_Tn:
		return 3
	}
	return -1
}

// String returns the literal string representation of a Season_Nc_Tn, which is
// the same as the const identifier but without prefix or suffix.
func (v Season_Nc_Tn) String() string {
	o := v.Ordinal()
	return v.toString(o, season_nc_tnEnumStrings, season_nc_tnEnumIndex[:])
}

func (v Season_Nc_Tn) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllSeason_Nc_Tns) {
		return fmt.Sprintf("Season_Nc_Tn(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// IsValid determines whether a Season_Nc_Tn is one of the defined constants.
func (v Season_Nc_Tn) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v Season_Nc_Tn) Int() int {
	return int(v)
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v Season_Nc_Tn) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The number representation is chosen according to -marshaltext.
func (v Season_Nc_Tn) marshalText() (string, error) {
	if !v.IsValid() {
		return v.marshalNumberStringOrError()
	}

	return season_nc_tnMarshalNumber(v), nil
}

func (v Season_Nc_Tn) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v Season_Nc_Tn) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v Season_Nc_Tn) invalidError() error {
	return fmt.Errorf("%d is not a valid season_nc_tn", v)
}

// season_nc_tnMarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var season_nc_tnMarshalNumber = func(v Season_Nc_Tn) string {
	return strconv.FormatInt(int64(v), 10)
}
