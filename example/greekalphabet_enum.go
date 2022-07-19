// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.14.0

package example

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"os"
	"strconv"
	"strings"
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
)

var (
	greekalphabetEnumIndex = [...]uint16{0, 8, 16, 26, 36, 48, 56, 62, 70, 78, 88, 100, 104, 108, 112, 126, 130, 134, 144, 150, 162, 166, 170, 174, 184}
)

func (v GreekAlphabet) toString(concats string, indexes []uint16) string {
	o := v.Ordinal()
	if o < 0 || o >= len(AllGreekAlphabets) {
		return fmt.Sprintf("GreekAlphabet(%d)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
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

var greekTagsInverse = map[string]GreekAlphabet{}

func init() {
	for _, id := range AllGreekAlphabets {
		v, exists := greekTags[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: GreekAlphabet: %s is missing from greekTags\n", id)
		} else {
			k := greekalphabetTransformInput(v)
			if _, exists := greekTagsInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: GreekAlphabet: %q is duplicated in greekTags\n", k)
			}
			greekTagsInverse[k] = id
		}
	}

	if len(greekTags) != 24 {
		panic(fmt.Sprintf("GreekAlphabet: greekTags has %d items but should have 24", len(greekTags)))
	}

	if len(greekTags) != len(greekTagsInverse) {
		panic(fmt.Sprintf("GreekAlphabet: greekTags has %d items but there are only %d distinct items",
			len(greekTags), len(greekTagsInverse)))
	}
}

// Tag returns the string representation of a GreekAlphabet. For invalid values,
// this returns v.String() (see IsValid).
func (v GreekAlphabet) Tag() string {
	s, ok := greekTags[v]
	if ok {
		return s
	}
	return v.String()
}

// String returns the literal string representation of a GreekAlphabet, which is
// the same as the const identifier but without prefix or suffix.
func (v GreekAlphabet) String() string {
	return v.toString(greekalphabetEnumStrings, greekalphabetEnumIndex[:])
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

// Parse parses a string to find the corresponding GreekAlphabet, accepting one of the string values or
// a number. The input representation is determined by greekalphabetMarshalTextRep. It is used by AsGreekAlphabet.
//
// Usage Example
//
//    v := new(GreekAlphabet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *GreekAlphabet) Parse(s string) error {
	return v.parse(s, greekalphabetMarshalTextRep)
}

func (v *GreekAlphabet) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := greekalphabetTransformInput(in)

	if rep == enum.Identifier {
		if v.parseString(s, greekalphabetEnumStrings, greekalphabetEnumIndex[:]) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseString(s, greekalphabetEnumStrings, greekalphabetEnumIndex[:]) {
			return nil
		}
	}

	return errors.New(in + ": unrecognised greekalphabet")
}

// parseNumber attempts to convert a decimal value
func (v *GreekAlphabet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = GreekAlphabet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *GreekAlphabet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllGreekAlphabets) {
		*v = AllGreekAlphabets[ord]
		return true
	}
	return false
}

// parseTag attempts to match an entry in greekTagsInverse
func (v *GreekAlphabet) parseTag(s string) (ok bool) {
	*v, ok = greekTagsInverse[s]
	return ok
}

// greekalphabetTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var greekalphabetTransformInput = func(in string) string {
	return in
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

// greekalphabetMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
// By default, it is enum.Tag.
// The initial value is set using the -marshaltext command line parameter.
var greekalphabetMarshalTextRep = enum.Tag

// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to greekalphabetMarshalTextRep.
func (v GreekAlphabet) MarshalText() (text []byte, err error) {
	return v.marshalText(greekalphabetMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to greekalphabetMarshalTextRep.
func (v GreekAlphabet) MarshalJSON() ([]byte, error) {
	return v.marshalText(greekalphabetMarshalTextRep, true)
}

func (v GreekAlphabet) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	if rep != enum.Ordinal && !v.IsValid() {
		return greekalphabetMarshalNumber(v)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return greekalphabetMarshalNumber(v)
	case enum.Ordinal:
		return v.marshalOrdinal()
	case enum.Tag:
		if quoted {
			bs = enum.QuotedString(v.Tag())
		} else {
			bs = []byte(v.Tag())
		}
	default:
		if quoted {
			bs = enum.QuotedString(v.String())
		} else {
			bs = []byte(v.String())
		}
	}
	return bs, nil
}

// greekalphabetMarshalNumber handles marshaling where a number is required or where
// the value is out of range but greekalphabetMarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var greekalphabetMarshalNumber = func(v GreekAlphabet) (text []byte, err error) {
	bs := []byte(strconv.FormatInt(int64(v), 10))
	return bs, nil
}

func (v GreekAlphabet) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(v.Ordinal()))
	return bs, nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (v *GreekAlphabet) UnmarshalText(text []byte) error {
	return v.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *GreekAlphabet) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *GreekAlphabet) unmarshalJSON(s string) error {
	return v.Parse(s)
}

// greekalphabetStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier.
// The initial value is set using the -store command line parameter.
var greekalphabetStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *GreekAlphabet) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		if greekalphabetStoreRep == enum.Ordinal {
			*v = GreekAlphabetOf(int(x))
		} else {
			*v = GreekAlphabet(x)
		}
		return nil
	case float64:
		*v = GreekAlphabet(x)
		return nil
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful greekalphabet", value, value)
	}

	return v.parse(s, greekalphabetStoreRep)
}

// Value converts the GreekAlphabet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v GreekAlphabet) Value() (driver.Value, error) {
	switch greekalphabetStoreRep {
	case enum.Number:
		return int64(v), nil
	case enum.Ordinal:
		return int64(v.Ordinal()), nil
	case enum.Tag:
		return v.Tag(), nil
	default:
		return v.String(), nil
	}
}
