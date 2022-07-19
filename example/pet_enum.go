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

const (
	petEnumStrings = "catdogmouseelephantkoala bear"
)

var (
	petEnumIndex = [...]uint16{0, 3, 6, 11, 19, 29}
)

func (i Pet) toString(concats string, indexes []uint16) string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllPets) {
		return fmt.Sprintf("Pet(%d)", i)
	}
	return concats[indexes[o]:indexes[o+1]]
}

func (v *Pet) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllPets[j-1]
			return true
		}
		i0 = i1
	}
	*v, ok = petAliases[s]
	return ok
}

var petTagsInverse = map[string]Pet{}

func init() {
	for _, id := range AllPets {
		v, exists := petTags[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Pet: %s is missing from petTags\n", id)
		} else {
			k := petTransformInput(v)
			if _, exists := petTagsInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: Pet: %q is duplicated in petTags\n", k)
			}
			petTagsInverse[k] = id
		}
	}

	if len(petTags) != 5 {
		panic(fmt.Sprintf("Pet: petTags has %d items but should have 5", len(petTags)))
	}

	if len(petTags) != len(petTagsInverse) {
		panic(fmt.Sprintf("Pet: petTags has %d items but there are only %d distinct items",
			len(petTags), len(petTagsInverse)))
	}
}

// Tag returns the string representation of a Pet. For invalid values,
// this returns i.String() (see IsValid).
func (i Pet) Tag() string {
	s, ok := petTags[i]
	if ok {
		return s
	}
	return i.String()
}

// String returns the literal string representation of a Pet, which is
// the same as the const identifier but without prefix or suffix.
func (i Pet) String() string {
	return i.toString(petEnumStrings, petEnumIndex[:])
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
// This facilitates polymorphism (see enum.IntEnum).
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

	s := petTransformInput(in)

	if rep == enum.Identifier {
		if v.parseString(s, petEnumStrings, petEnumIndex[:]) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseString(s, petEnumStrings, petEnumIndex[:]) {
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

// parseTag attempts to match an entry in petTagsInverse
func (v *Pet) parseTag(s string) (ok bool) {
	*v, ok = petTagsInverse[s]
	return ok
}

// petTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var petTransformInput = func(in string) string {
	return strings.ToLower(strings.ReplaceAll(in, "_", " "))
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
var petMarshalTextRep = enum.Tag

// MarshalText converts values to a form suitable for transmission via XML etc.
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
	if petMarshalTextRep != enum.Ordinal && i.Ordinal() < 0 {
		return petMarshalNumber(i)
	}

	var bs []byte
	switch rep {
	case enum.Number:
		return petMarshalNumber(i)
	case enum.Ordinal:
		return i.marshalOrdinal()
	case enum.Tag:
		if quoted {
			bs = enum.QuotedString(i.Tag())
		} else {
			bs = []byte(i.Tag())
		}
	default:
		if quoted {
			bs = enum.QuotedString(i.String())
		} else {
			bs = []byte(i.String())
		}
	}
	return bs, nil
}

// petMarshalNumber handles marshaling where a number is required or where
// the value is out of range but petMarshalTextRep != enum.Ordinal.
// This function can be replaced with any bespoke function than matches signature.
var petMarshalNumber = func(i Pet) (text []byte, err error) {
	bs := []byte(strconv.FormatInt(int64(i), 10))
	return bs, nil
}

func (i Pet) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(i.Ordinal()))
	return bs, nil
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
	return i.unmarshalJSON(s)
}

func (i *Pet) unmarshalJSON(s string) error {
	return i.Parse(s)
}

// petStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var petStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Pet) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch v := value.(type) {
	case int64:
		if petStoreRep == enum.Ordinal {
			*i = PetOf(int(v))
		} else {
			*i = Pet(v)
		}
		return nil
	case float64:
		*i = Pet(v)
		return nil
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("%T %+v is not a meaningful pet", value, value)
	}

	return i.parse(s, petStoreRep)
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
