// generated code - do not edit
// github.com/rickb777/enumeration/v2 v2.9.0

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

const petEnumStrings = "catdogmouseelephantkoala bear"

var petEnumIndex = [...]uint16{0, 3, 6, 11, 19, 29}

// AllPets lists all 5 values in order.
var AllPets = []Pet{
	MyCat, MyDog, MyMouse, MyElephant,
	MyKoala_Bear,
}

// AllPetEnums lists all 5 values in order.
var AllPetEnums = enum.IntEnums{
	MyCat, MyDog, MyMouse, MyElephant,
	MyKoala_Bear,
}

// String returns the literal string representation of a Pet, which is
// the same as the const identifier.
func (i Pet) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllPets) {
		return fmt.Sprintf("Pet(%d)", i)
	}
	return petEnumStrings[petEnumIndex[o]:petEnumIndex[o+1]]
}

var petStringsInverse = map[string]Pet{}

func init() {
	for _, id := range AllPets {
		v, exists := petStrings[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Pet: %s is missing from petStrings\n", id)
		} else {
			k := strings.ToLower(strings.ReplaceAll(v, "_", " "))
			if _, exists := petStringsInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: Pet: %q is duplicated in petStrings\n", k)
			}
			petStringsInverse[k] = id
		}
	}

	if len(petStrings) != 5 {
		panic(fmt.Sprintf("Pet: petStrings has %d items but should have 5", len(petStrings)))
	}

	if len(petStrings) != len(petStringsInverse) {
		panic(fmt.Sprintf("Pet: petStrings has %d items but there are only %d distinct items",
			len(petStrings), len(petStringsInverse)))
	}
}

// Tag returns the string representation of a Pet. For invalid values,
// this returns i.String() (see IsValid).
func (i Pet) Tag() string {
	s, ok := petStrings[i]
	if ok {
		return s
	}
	return i.String()
}

// Ordinal returns the ordinal number of a Pet. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (i Pet) Ordinal() int {
	switch i {
	case MyCat:
		return 0
	case MyDog:
		return 1
	case MyMouse:
		return 2
	case MyElephant:
		return 3
	case MyKoala_Bear:
		return 4
	}
	return -1
}

// Int returns the int value, which is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.IntEnum).
func (i Pet) Int() int {
	return int(i)
}

// PetOf returns a Pet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Pet is returned.
func PetOf(i int) Pet {
	if 0 <= i && i < len(AllPets) {
		return AllPets[i]
	}
	// an invalid result
	return MyCat + MyDog + MyMouse + MyElephant + MyKoala_Bear + 1
}

// IsValid determines whether a Pet is one of the defined constants.
func (i Pet) IsValid() bool {
	return i.Ordinal() >= 0
}

// Parse parses a string to find the corresponding Pet, accepting one of the string values or
// a number. The input representation is determined by petMarshalTextRep. It is used by AsPet.
//
// Usage Example
//
//    v := new(Pet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Pet) Parse(s string) error {
	return v.parse(s, petMarshalTextRep)
}

func (v *Pet) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := strings.ToLower(strings.ReplaceAll(in, "_", " "))

	if rep == enum.Identifier {
		if v.parseIdentifier(s) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseIdentifier(s) {
			return nil
		}
	}

	return errors.New(in + ": unrecognised pet")
}

// parseNumber attempts to convert a decimal value
func (v *Pet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Pet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Pet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllPets) {
		*v = AllPets[ord]
		return true
	}
	return false
}

// parseTag attempts to match an entry in petStringsInverse
func (v *Pet) parseTag(s string) (ok bool) {
	*v, ok = petStringsInverse[s]
	return ok
}

// parseIdentifier attempts to match an identifier.
func (v *Pet) parseIdentifier(s string) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(petEnumIndex); j++ {
		i1 := petEnumIndex[j]
		p := petEnumStrings[i0:i1]
		if s == p {
			*v = AllPets[j-1]
			return true
		}
		i0 = i1
	}
	return false
}

// AsPet parses a string to find the corresponding Pet, accepting either one of the string values or
// a number. The input representation is determined by petMarshalTextRep. It wraps Parse.
func AsPet(s string) (Pet, error) {
	var i = new(Pet)
	err := i.Parse(s)
	return *i, err
}

// MustParsePet is similar to AsPet except that it panics on error.
func MustParsePet(s string) Pet {
	i, err := AsPet(s)
	if err != nil {
		panic(err)
	}
	return i
}

// petMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
var petMarshalTextRep = enum.Identifier

// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
// The representation is chosen according to petMarshalTextRep.
func (i Pet) MarshalText() (text []byte, err error) {
	return i.marshalText(petMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to petMarshalTextRep.
func (i Pet) MarshalJSON() ([]byte, error) {
	return i.marshalText(petMarshalTextRep, true)
}

func (i Pet) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
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

func (i Pet) quotedString(s string) []byte {
	b := make([]byte, len(s)+2)
	b[0] = '"'
	copy(b[1:], s)
	b[len(s)+1] = '"'
	return b
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *Pet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Pet) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.Parse(s)
}

// petStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var petStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Pet) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	err = nil
	switch v := value.(type) {
	case int64:
		if petStoreRep == enum.Ordinal {
			*i = PetOf(int(v))
		} else {
			*i = Pet(v)
		}
	case float64:
		*i = Pet(v)
	case []byte:
		err = i.parse(string(v), petStoreRep)
	case string:
		err = i.parse(v, petStoreRep)
	default:
		err = fmt.Errorf("%T %+v is not a meaningful pet", value, value)
	}

	return err
}

// Value converts the Pet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Pet) Value() (driver.Value, error) {
	switch petStoreRep {
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
