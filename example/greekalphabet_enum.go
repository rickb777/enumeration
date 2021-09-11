// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.5.2

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

const greekalphabetEnumStrings = "ΑλφαΒήταΓάμμαΔέλταΕψιλονΖήταΗταΘήταΙώταΚάππαΛάμβδαΜυΝυΞιΟμικρονΠιΡώΣίγμαΤαυΥψιλονΦιΧιΨιΩμέγα"

var greekalphabetEnumIndex = [...]uint16{0, 8, 16, 26, 36, 48, 56, 62, 70, 78, 88, 100, 104, 108, 112, 126, 130, 134, 144, 150, 162, 166, 170, 174, 184}

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

// String returns the literal string representation of a GreekAlphabet, which is
// the same as the const identifier.
func (i GreekAlphabet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllGreekAlphabets) {
		return fmt.Sprintf("GreekAlphabet(%d)", i)
	}
	return greekalphabetEnumStrings[greekalphabetEnumIndex[o]:greekalphabetEnumIndex[o+1]]
}

var greekStringsInverse = map[string]GreekAlphabet{}

func init() {
	for _, id := range AllGreekAlphabets {
		v, exists := greekStrings[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: GreekAlphabet: %s is missing from greekStrings\n", id)
		} else {
			k := v
			if _, exists := greekStringsInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: GreekAlphabet: %q is duplicated in greekStrings\n", k)
			}
			greekStringsInverse[k] = id
		}
	}

	if len(greekStrings) != 24 {
		panic(fmt.Sprintf("GreekAlphabet: greekStrings has %d items but should have 24", len(greekStrings)))
	}

	if len(greekStrings) != len(greekStringsInverse) {
		panic(fmt.Sprintf("GreekAlphabet: greekStrings has %d items but there are only %d distinct items",
			len(greekStrings), len(greekStringsInverse)))
	}
}

// Tag returns the string representation of a GreekAlphabet.
func (i GreekAlphabet) Tag() string {
	s, ok := greekStrings[i]
	if ok {
		return s
	}
	return fmt.Sprintf("%02d", i)
}

// Ordinal returns the ordinal number of a GreekAlphabet.
func (i GreekAlphabet) Ordinal() int {
	switch i {
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

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i GreekAlphabet) Int() int {
	return int(i)
}

// GreekAlphabetOf returns a GreekAlphabet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid GreekAlphabet is returned.
func GreekAlphabetOf(i int) GreekAlphabet {
	if 0 <= i && i < len(AllGreekAlphabets) {
		return AllGreekAlphabets[i]
	}
	// an invalid result
	return Αλφα + Βήτα + Γάμμα + Δέλτα + Εψιλον + Ζήτα + Ητα + Θήτα + Ιώτα + Κάππα + Λάμβδα + Μυ + Νυ + Ξι + Ομικρον + Πι + Ρώ + Σίγμα + Ταυ + Υψιλον + Φι + Χι + Ψι + Ωμέγα + 1
}

// IsValid determines whether a GreekAlphabet is one of the defined constants.
func (i GreekAlphabet) IsValid() bool {
	switch i {
	case Αλφα, Βήτα, Γάμμα, Δέλτα,
		Εψιλον, Ζήτα, Ητα, Θήτα, Ιώτα,
		Κάππα, Λάμβδα, Μυ, Νυ, Ξι,
		Ομικρον, Πι, Ρώ, Σίγμα, Ταυ,
		Υψιλον, Φι, Χι, Ψι, Ωμέγα:
		return true
	}
	return false
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

	s := in

	if rep == enum.Identifier {
		if v.parseIdentifier(s) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseIdentifier(s) {
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

// parseTag attempts to match an entry in greekStringsInverse
func (v *GreekAlphabet) parseTag(s string) (ok bool) {
	*v, ok = greekStringsInverse[s]
	return ok
}

// parseIdentifier attempts to match an identifier.
func (v *GreekAlphabet) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(greekalphabetEnumIndex); j++ {
		i1 := greekalphabetEnumIndex[j]
		p := greekalphabetEnumStrings[i0:i1]
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
	var i = new(GreekAlphabet)
	err := i.Parse(s)
	return *i, err
}

// greekalphabetMarshalTextRep controls representation used for XML and other text encodings.
// By default, it is enum.Identifier and quoted strings are used.
var greekalphabetMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to greekalphabetMarshalTextRep.
func (i GreekAlphabet) MarshalText() (text []byte, err error) {
	return i.marshalText(greekalphabetMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to greekalphabetMarshalTextRep.
func (i GreekAlphabet) MarshalJSON() ([]byte, error) {
	return i.marshalText(greekalphabetMarshalTextRep, true)
}

func (i GreekAlphabet) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	var bs []byte
	switch rep {
	case enum.Number:
		bs = []byte(strconv.FormatInt(int64(i), 10))
	case enum.Ordinal:
		bs = []byte(strconv.Itoa(i.Ordinal()))
	case enum.Tag:
		if quoted {
			bs = i.quotedString(i.Tag())
		} else {
			bs = []byte(i.Tag())
		}
	default:
		if quoted {
			bs = []byte(i.quotedString(i.String()))
		} else {
			bs = []byte(i.String())
		}
	}
	return bs, nil
}

func (i GreekAlphabet) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *GreekAlphabet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *GreekAlphabet) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}

// greekalphabetStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var greekalphabetStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *GreekAlphabet) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if greekalphabetStoreRep == enum.Ordinal {
			*i = GreekAlphabetOf(int(v))
		} else {
			*i = GreekAlphabet(v)
		}
	case float64:
		*i = GreekAlphabet(v)
	case []byte:
		err = i.parse(string(v), greekalphabetStoreRep)
	case string:
		err = i.parse(v, greekalphabetStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful greekalphabet", value, value)
	}

	return err
}

// Value converts the GreekAlphabet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i GreekAlphabet) Value() (driver.Value, error) {
	switch greekalphabetStoreRep {
	case enum.Number:
		return int64(i), nil
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.String(), nil
	}
}
