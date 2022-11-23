// generated code - do not edit
// github.com/rickb777/enumeration/v3 v3.0.2

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"strconv"
)

// AllGreekAlphabets lists all 24 values in order.
var AllGreekAlphabets = []GreekAlphabet{
	Αλφα, Βήτα, Γάμμα, Δέλτα,
	Εψιλον, Ζήτα, Ητα, Θήτα, Ιώτα,
	Κάππα, Λάμβδα, Μυ, Νυ, Ξι,
	Ομικρον, Πι, Ρώ, Σίγμα, Ταυ,
	Υψιλον, Φι, Χι, Ψι, Ωμέγα,
}

// AllGreekAlphabetEnums lists all 24 values in order.
var AllGreekAlphabetEnums = enum.IntEnums{
	Αλφα, Βήτα, Γάμμα, Δέλτα,
	Εψιλον, Ζήτα, Ητα, Θήτα, Ιώτα,
	Κάππα, Λάμβδα, Μυ, Νυ, Ξι,
	Ομικρον, Πι, Ρώ, Σίγμα, Ταυ,
	Υψιλον, Φι, Χι, Ψι, Ωμέγα,
}

const (
	greekalphabetEnumStrings = "ΑλφαΒήταΓάμμαΔέλταΕψιλονΖήταΗταΘήταΙώταΚάππαΛάμβδαΜυΝυΞιΟμικρονΠιΡώΣίγμαΤαυΥψιλονΦιΧιΨιΩμέγα"
	greekalphabetTextStrings = "alphabetagammadeltaepsilonzetaetathetaiotakappalambdamunuxiomicronpirhosigmatauupsilonphichipsiomega"
	greekalphabetSQLStrings  = "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"
)

var (
	greekalphabetEnumIndex = [...]uint16{0, 8, 16, 26, 36, 48, 56, 62, 70, 78, 88, 100, 104, 108, 112, 126, 130, 134, 144, 150, 162, 166, 170, 174, 184}
	greekalphabetTextIndex = [...]uint16{0, 5, 9, 14, 19, 26, 30, 33, 38, 42, 47, 53, 55, 57, 59, 66, 68, 71, 76, 79, 86, 89, 92, 95, 100}
	greekalphabetSQLIndex  = [...]uint16{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48}
)

// String returns the literal string representation of a GreekAlphabet, which is
// the same as the const identifier but without prefix or suffix.
func (v GreekAlphabet) String() string {
	o := v.Ordinal()
	return v.toString(o, greekalphabetEnumStrings, greekalphabetEnumIndex[:])
}

func (v GreekAlphabet) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(AllGreekAlphabets) {
		return fmt.Sprintf("GreekAlphabet(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}

// Ordinal returns the ordinal number of a GreekAlphabet. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v GreekAlphabet) Ordinal() int {
	switch v {
	case Αλφα:
		return 0
	case Βήτα:
		return 1
	case Γάμμα:
		return 2
	case Δέλτα:
		return 3
	case Εψιλον:
		return 4
	case Ζήτα:
		return 5
	case Ητα:
		return 6
	case Θήτα:
		return 7
	case Ιώτα:
		return 8
	case Κάππα:
		return 9
	case Λάμβδα:
		return 10
	case Μυ:
		return 11
	case Νυ:
		return 12
	case Ξι:
		return 13
	case Ομικρον:
		return 14
	case Πι:
		return 15
	case Ρώ:
		return 16
	case Σίγμα:
		return 17
	case Ταυ:
		return 18
	case Υψιλον:
		return 19
	case Φι:
		return 20
	case Χι:
		return 21
	case Ψι:
		return 22
	case Ωμέγα:
		return 23
	}
	return -1
}

// IsValid determines whether a GreekAlphabet is one of the defined constants.
func (v GreekAlphabet) IsValid() bool {
	return v.Ordinal() >= 0
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v GreekAlphabet) Int() int {
	return int(v)
}

// GreekAlphabetOf returns a GreekAlphabet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid GreekAlphabet is returned.
func GreekAlphabetOf(v int) GreekAlphabet {
	if 0 <= v && v < len(AllGreekAlphabets) {
		return AllGreekAlphabets[v]
	}
	// an invalid result
	return Αλφα + Βήτα + Γάμμα + Δέλτα + Εψιλον + Ζήτα + Ητα + Θήτα + Ιώτα + Κάππα + Λάμβδα + Μυ + Νυ + Ξι + Ομικρον + Πι + Ρώ + Σίγμα + Ταυ + Υψιλον + Φι + Χι + Ψι + Ωμέγα + 1
}

// parseNumber attempts to convert a decimal value.
// Only numbers that correspond to the enumeration are valid.
func (v *GreekAlphabet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = GreekAlphabet(num)
		return v.IsValid()
	}
	return false
}

// Parse parses a string to find the corresponding GreekAlphabet, accepting one of the string values or
// a number. The input representation is determined by None. It is used by AsGreekAlphabet.
//
// Usage Example
//
//	v := new(GreekAlphabet)
//	err := v.Parse(s)
//	...  etc
func (v *GreekAlphabet) Parse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := greekalphabetTransformInput(in)

	return v.parseFallback(in, s)
}

func (v *GreekAlphabet) parseFallback(in, s string) error {
	if v.parseString(s, greekalphabetEnumStrings, greekalphabetEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised greekalphabet")
}

// greekalphabetTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var greekalphabetTransformInput = func(in string) string {
	return in
}

func (v *GreekAlphabet) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllGreekAlphabets[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsGreekAlphabet parses a string to find the corresponding GreekAlphabet, accepting either one of the string values or
// a number. The input representation is determined by greekalphabetMarshalTextRep. It wraps Parse.
func AsGreekAlphabet(s string) (GreekAlphabet, error) {
	var v = new(GreekAlphabet)
	err := v.Parse(s)
	return *v, err
}

// MustParseGreekAlphabet is similar to AsGreekAlphabet except that it panics on error.
func MustParseGreekAlphabet(s string) GreekAlphabet {
	v, err := AsGreekAlphabet(s)
	if err != nil {
		panic(err)
	}
	return v
}

// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v GreekAlphabet) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}

// Text returns the representation used for transmission via XML, JSON etc.
func (v GreekAlphabet) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v GreekAlphabet) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, greekalphabetTextStrings, greekalphabetTextIndex[:]), nil
}

func (v GreekAlphabet) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}

func (v GreekAlphabet) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}

func (v GreekAlphabet) invalidError() error {
	return fmt.Errorf("%d is not a valid greekalphabet", v)
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *GreekAlphabet) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}

func (v *GreekAlphabet) unmarshalText(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := greekalphabetTransformInput(in)

	if v.parseString(s, greekalphabetTextStrings, greekalphabetTextIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *GreekAlphabet) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = GreekAlphabet(x)
		return v.errorIfInvalid()
	case float64:
		*v = GreekAlphabet(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful greekalphabet", value, value)
	}

	return v.scanParse(s)
}

func (v *GreekAlphabet) scanParse(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := greekalphabetTransformInput(in)

	if v.parseString(s, greekalphabetSQLStrings, greekalphabetSQLIndex[:]) {
		return nil
	}

	return v.parseFallback(in, s)
}

func (v GreekAlphabet) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}

// Value converts the GreekAlphabet to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v GreekAlphabet) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, greekalphabetSQLStrings, greekalphabetSQLIndex[:]), nil
}
