package model

import (
	"github.com/rickb777/enumeration/v2/internal/transform"
	"github.com/rickb777/enumeration/v2/internal/util"
	"io"
	"strings"
	"testing"
)

const head0 = `// generated code - do not edit
// github.com/rickb777/enumeration/v2 `

const head1 = `

package confectionary

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
	"strconv"
	"strings"
)
`

func TestWriteHead(t *testing.T) {
	expected := head0 + util.Version + head1
	testStage(t, unsnake(basicModel()).writeHead, expected)
}

//-------------------------------------------------------------------------------------------------

const AllSweets_simple = `
// AllSweets lists all 5 values in order.
var AllSweets = []Sweet{
	Mars, Bounty, Snickers, Kitkat,
	Dairy_Milk,
}

// AllSweetEnums lists all 5 values in order.
var AllSweetEnums = enum.IntEnums{
	Mars, Bounty, Snickers, Kitkat,
	Dairy_Milk,
}
`

const AllSweets_prefixed = `
// AllSweets lists all 5 values in order.
var AllSweets = []Sweet{
	AMarsBar, ABountyBar, ASnickersBar, AKitkatBar,
	ADairy_MilkBar,
}

// AllSweetEnums lists all 5 values in order.
var AllSweetEnums = enum.FloatEnums{
	AMarsBar, ABountyBar, ASnickersBar, AKitkatBar,
	ADairy_MilkBar,
}
`

func TestWriteAllItems(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeAllItems, AllSweets_simple)
	testStage(t, unsnake(basicModel()).writeAllItems, AllSweets_simple)
	testStage(t, floatModelWithPrefixes().writeAllItems, AllSweets_prefixed)
}

//-------------------------------------------------------------------------------------------------

const enumStrings_nc = `
const (
	sweetEnumStrings = "MarsBountySnickersKitkatDairy_Milk"
)

var (
	sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 34}
)
`

const enumStrings_nc_unsnake = `
const (
	sweetEnumStrings = "MarsBountySnickersKitkatDairy Milk"
)

var (
	sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 34}
)
`

const enumStrings_lc = `
const (
	sweetEnumStrings = "marsbountysnickerskitkatdairy_milk"
)

var (
	sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 34}
)
`

const enumStrings_lc_unsnake = `
const (
	sweetEnumStrings = "marsbountysnickerskitkatdairy milk"
)

var (
	sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 34}
)
`

const enumStrings_ic = `
const (
	sweetEnumStrings = "MarsBountySnickersKitkatDairy_Milk"
	sweetEnumInputs  = "marsbountysnickerskitkatdairy_milk"
)

var (
	sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 34}
)
`

const enumStrings_tagged_nc = `
const (
	sweetEnumStrings = "MarsBountySnickersKitkatDairy_Milk"
	sweetEnumInputs  = "marsbountysnickerskitkatdairy_milk"
	sweetJSONStrings = "mmmbbbssskkkddd"
	sweetJSONInputs  = "mmmbbbssskkkddd"
	sweetSQLStrings  = "mbskd"
	sweetSQLInputs   = "mbskd"
)

var (
	sweetEnumIndex = [...]uint16{0, 4, 10, 18, 24, 34}
	sweetJSONIndex = [...]uint16{0, 3, 6, 9, 12, 15}
	sweetSQLIndex = [...]uint16{0, 1, 2, 3, 4, 5}
)
`

func TestWriteJoinedStringAndIndexes(t *testing.T) {
	testStage(t, basicModel().writeJoinedStringAndIndexes, enumStrings_nc)
	testStage(t, unsnake(basicModel()).writeJoinedStringAndIndexes, enumStrings_nc_unsnake)
	testStage(t, lowerCase(floatModelWithPrefixes()).writeJoinedStringAndIndexes, enumStrings_lc)
	testStage(t, lowerCase(unsnake(floatModelWithPrefixes())).writeJoinedStringAndIndexes, enumStrings_lc_unsnake)
	testStage(t, ignoreCase(basicModel()).writeJoinedStringAndIndexes, enumStrings_ic)
	testStage(t, ignoreCase(modelWithStructTags()).writeJoinedStringAndIndexes, enumStrings_tagged_nc)
}

//-------------------------------------------------------------------------------------------------

const toString_int = `
func (i Sweet) toString(concats string, indexes []uint16) string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%d)", i)
	}
	return concats[indexes[o]:indexes[o+1]]
}
`

const toString_float = `
func (i Sweet) toString(concats string, indexes []uint16) string {
	o := i.Ordinal()
	if o < 0 || o >= len(AllSweets) {
		return fmt.Sprintf("Sweet(%g)", i)
	}
	return concats[indexes[o]:indexes[o+1]]
}
`

func TestWriteLiteralMethod(t *testing.T) {
	testStage(t, basicModel().writeToStringMethod, toString_int)
	testStage(t, floatModelWithPrefixes().writeToStringMethod, toString_float)
}

//-------------------------------------------------------------------------------------------------

const String_all = `
// String returns the literal string representation of a Sweet, which is
// the same as the const identifier but without prefix or suffix.
func (i Sweet) String() string {
	return i.toString(sweetEnumStrings, sweetEnumIndex[:])
}
`

func TestWriteStringMethod(t *testing.T) {
	testStage(t, basicModel().writeStringMethod, String_all)
}

//-------------------------------------------------------------------------------------------------

const Tag_no_table = `
// Tag returns the string representation of a Sweet. This is an alias for String.
func (i Sweet) Tag() string {
	return i.String()
}
`

const Tag_as_JSON = `
// Tag returns the JSON representation of a Sweet.
func (i Sweet) Tag() string {
	return i.toString(sweetJSONStrings, sweetJSONIndex[:])
}
`

const Tag_with_table = `
var sweetNamesInverse = map[string]Sweet{}

func init() {
	for _, id := range AllSweets {
		v, exists := sweetNames[id]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Sweet: %s is missing from sweetNames\n", id)
		} else {
			k := sweetTransformInput(v)
			if _, exists := sweetNamesInverse[k]; exists {
				fmt.Fprintf(os.Stderr, "Warning: Sweet: %q is duplicated in sweetNames\n", k)
			}
			sweetNamesInverse[k] = id
		}
	}

	if len(sweetNames) != 5 {
		panic(fmt.Sprintf("Sweet: sweetNames has %d items but should have 5", len(sweetNames)))
	}

	if len(sweetNames) != len(sweetNamesInverse) {
		panic(fmt.Sprintf("Sweet: sweetNames has %d items but there are only %d distinct items",
			len(sweetNames), len(sweetNamesInverse)))
	}
}

// Tag returns the string representation of a Sweet. For invalid values,
// this returns i.String() (see IsValid).
func (i Sweet) Tag() string {
	s, ok := sweetNames[i]
	if ok {
		return s
	}
	return i.String()
}
`

func TestWriteTagMethod(t *testing.T) {
	testStage(t, basicModel().writeTagMethod, Tag_no_table)
	testStage(t, modelWithStructTags().writeTagMethod, Tag_as_JSON)
	testStage(t, lookupTables(basicModel()).writeTagMethod, Tag_with_table)
}

//-------------------------------------------------------------------------------------------------

const ordinal_no_prefix = `
// Ordinal returns the ordinal number of a Sweet. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (i Sweet) Ordinal() int {
	switch i {
	case Mars:
		return 0
	case Bounty:
		return 1
	case Snickers:
		return 2
	case Kitkat:
		return 3
	case Dairy_Milk:
		return 4
	}
	return -1
}
`

const ordinal_with_prefix = `
// Ordinal returns the ordinal number of a Sweet. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (i Sweet) Ordinal() int {
	switch i {
	case AMarsBar:
		return 0
	case ABountyBar:
		return 1
	case ASnickersBar:
		return 2
	case AKitkatBar:
		return 3
	case ADairy_MilkBar:
		return 4
	}
	return -1
}
`

func TestWriteOrdinalMethod(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeOrdinalMethod, ordinal_no_prefix)
	testStage(t, floatModelWithPrefixes().writeOrdinalMethod, ordinal_with_prefix)
}

//-------------------------------------------------------------------------------------------------

const int_nc = `
// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (i Sweet) Int() int {
	return int(i)
}
`

const float_lc = `
// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (i Sweet) Float() float64 {
	return float64(i)
}
`

func TestWriteBaseMethod(t *testing.T) {
	testStage(t, basicModel().writeBaseMethod, int_nc)
	testStage(t, floatModelWithPrefixes().writeBaseMethod, float_lc)
}

//-------------------------------------------------------------------------------------------------

const SweetOf_no_prefix = `
// SweetOf returns a Sweet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Sweet is returned.
func SweetOf(i int) Sweet {
	if 0 <= i && i < len(AllSweets) {
		return AllSweets[i]
	}
	// an invalid result
	return Mars + Bounty + Snickers + Kitkat + Dairy_Milk + 1
}
`

const SweetOf_with_prefix = `
// SweetOf returns a Sweet based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid Sweet is returned.
func SweetOf(i int) Sweet {
	if 0 <= i && i < len(AllSweets) {
		return AllSweets[i]
	}
	// an invalid result
	return AMarsBar + ABountyBar + ASnickersBar + AKitkatBar + ADairy_MilkBar + 1
}
`

func TestWriteOfMethod(t *testing.T) {
	testStage(t, basicModel().writeOfMethod, SweetOf_no_prefix)
	testStage(t, floatModelWithPrefixes().writeOfMethod, SweetOf_with_prefix)
}

//-------------------------------------------------------------------------------------------------

const IsValid_all = `
// IsValid determines whether a Sweet is one of the defined constants.
func (i Sweet) IsValid() bool {
	return i.Ordinal() >= 0
}
`

func TestWriteIsValid(t *testing.T) {
	testStage(t, basicModel().writeIsValidMethod, IsValid_all)
	testStage(t, floatModelWithPrefixes().writeIsValidMethod, IsValid_all)
}

//-------------------------------------------------------------------------------------------------

const Parse_nc = `
// Parse parses a string to find the corresponding Sweet, accepting one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It is used by AsSweet.
//
// Usage Example
//
//    v := new(Sweet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Sweet) Parse(s string) error {
	return v.parse(s, sweetMarshalTextRep)
}

func (v *Sweet) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := sweetTransformInput(in)

	if v.parseString(s, sweetEnumStrings, sweetEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised sweet")
}
`

const Parse_lc = `
// Parse parses a string to find the corresponding Sweet, accepting one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It is used by AsSweet.
//
// Usage Example
//
//    v := new(Sweet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Sweet) Parse(s string) error {
	return v.parse(s, sweetMarshalTextRep)
}

func (v *Sweet) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := sweetTransformInput(in)

	if rep == enum.Identifier {
		if v.parseString(s, sweetEnumStrings, sweetEnumIndex[:]) || v.parseTag(s) {
			return nil
		}
	} else {
		if v.parseTag(s) || v.parseString(s, sweetEnumStrings, sweetEnumIndex[:]) {
			return nil
		}
	}

	return errors.New(in + ": unrecognised sweet")
}
`

const Parse_ic = `
// Parse parses a string to find the corresponding Sweet, accepting one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It is used by AsSweet.
// The input case does not matter.
//
// Usage Example
//
//    v := new(Sweet)
//    err := v.Parse(s)
//    ...  etc
//
func (v *Sweet) Parse(s string) error {
	return v.parse(s, sweetMarshalTextRep)
}

func (v *Sweet) parse(in string, rep enum.Representation) error {
	if rep == enum.Ordinal {
		if v.parseOrdinal(in) {
			return nil
		}
	} else {
		if v.parseNumber(in) {
			return nil
		}
	}

	s := sweetTransformInput(in)

	if v.parseString(s, sweetEnumInputs, sweetEnumIndex[:]) {
		return nil
	}

	return errors.New(in + ": unrecognised sweet")
}
`

func TestWriteParseMethod(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeParseMethod, Parse_nc)
	testStage(t, lookupTables(floatModelWithPrefixes()).writeParseMethod, Parse_lc)
	testStage(t, ignoreCase(unsnake(basicModel())).writeParseMethod, Parse_ic)
}

//-------------------------------------------------------------------------------------------------

const parseHelpers_nc = `
// parseNumber attempts to convert a decimal value
func (v *Sweet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Sweet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Sweet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return true
	}
	return false
}
`

const parseHelpers_lc = `
// parseNumber attempts to convert a decimal value
func (v *Sweet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseFloat(s, 64)
	if err == nil {
		*v = Sweet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Sweet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return true
	}
	return false
}

// parseTag attempts to match an entry in sweetNamesInverse
func (v *Sweet) parseTag(s string) (ok bool) {
	*v, ok = sweetNamesInverse[s]
	return ok
}
`

const parseHelpers_ic = `
// parseNumber attempts to convert a decimal value
func (v *Sweet) parseNumber(s string) (ok bool) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		*v = Sweet(num)
		return v.IsValid()
	}
	return false
}

// parseOrdinal attempts to convert an ordinal value
func (v *Sweet) parseOrdinal(s string) (ok bool) {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(AllSweets) {
		*v = AllSweets[ord]
		return true
	}
	return false
}
`

func TestWriteParseHelperMethods(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeParseHelperMethods, parseHelpers_nc)
	testStage(t, lookupTables(floatModelWithPrefixes()).writeParseHelperMethods, parseHelpers_lc)
	testStage(t, ignoreCase(unsnake(basicModel())).writeParseHelperMethods, parseHelpers_ic)
}

//-------------------------------------------------------------------------------------------------

const transformInput_nc = `
// sweetTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var sweetTransformInput = func(in string) string {
	return in
}
`

const transformInput_lc = `
// sweetTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var sweetTransformInput = func(in string) string {
	return strings.ToLower(in)
}
`

const transformInput_uc_unsnake = `
// sweetTransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var sweetTransformInput = func(in string) string {
	return strings.ToUpper(strings.ReplaceAll(in, "_", " "))
}
`

func TestWriteTransformInputMethod(t *testing.T) {
	testStage(t, basicModel().writeTransformInputFunction, transformInput_nc)
	testStage(t, ignoreCase(basicModel()).writeTransformInputFunction, transformInput_lc)
	testStage(t, lowerCase(basicModel()).writeTransformInputFunction, transformInput_lc)
	testStage(t, upperCase(unsnake(basicModel())).writeTransformInputFunction, transformInput_uc_unsnake)
}

//-------------------------------------------------------------------------------------------------

const parseString_nc = `
func (v *Sweet) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSweets[j-1]
			return true
		}
		i0 = i1
	}
	return false
}
`

const parseString_aliasTable = `
func (v *Sweet) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = AllSweets[j-1]
			return true
		}
		i0 = i1
	}
	*v, ok = sweetAlias[s]
	return ok
}
`

func TestWriteParseIdentifierMethod(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeParseIdentifierMethod, parseString_nc)
	testStage(t, modelWithStructTags().writeParseIdentifierMethod, parseString_aliasTable)
	testStage(t, ignoreCase(modelWithStructTags()).writeParseIdentifierMethod, parseString_aliasTable)
}

//-------------------------------------------------------------------------------------------------

const AsSweet_nc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It wraps Parse.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const AsSweet_lc = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It wraps Parse.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

const AsSweet_ic = `
// AsSweet parses a string to find the corresponding Sweet, accepting either one of the string values or
// a number. The input representation is determined by sweetMarshalTextRep. It wraps Parse.
// The input case does not matter.
func AsSweet(s string) (Sweet, error) {
	var i = new(Sweet)
	err := i.Parse(s)
	return *i, err
}
`

func TestWriteAsMethod(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeAsMethod, AsSweet_nc)
	testStage(t, floatModelWithPrefixes().writeAsMethod, AsSweet_lc)
	testStage(t, ignoreCase(unsnake(basicModel())).writeAsMethod, AsSweet_ic)
}

//-------------------------------------------------------------------------------------------------

const MustParseSweet_nc = `
// MustParseSweet is similar to AsSweet except that it panics on error.
func MustParseSweet(s string) Sweet {
	i, err := AsSweet(s)
	if err != nil {
		panic(err)
	}
	return i
}
`

const MustParseSweet_ic = `
// MustParseSweet is similar to AsSweet except that it panics on error.
// The input case does not matter.
func MustParseSweet(s string) Sweet {
	i, err := AsSweet(s)
	if err != nil {
		panic(err)
	}
	return i
}
`

func TestWriteMustParseMethod(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeMustParseMethod, MustParseSweet_nc)
	testStage(t, ignoreCase(unsnake(basicModel())).writeMustParseMethod, MustParseSweet_ic)
}

//-------------------------------------------------------------------------------------------------

const sweetMarshalTextRep = `
// sweetMarshalTextRep controls representation used for XML and other text encodings.
// When enum.Identifier, quoted strings are used. When enum.Tag the quoted strings will use
// the associated tag map values. When enum.Ordinal, an integer will be used based on the
// Ordinal method. When enum.Number, the number underlying the value will be used.
var sweetMarshalTextRep = enum.Identifier
`

const MarshalJSON_simple = `
// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalText() (text []byte, err error) {
	return i.marshalText(sweetMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalJSON() ([]byte, error) {
	return i.marshalText(sweetMarshalTextRep, true)
}
`

const MarshalJSON_struct_tags = `
// MarshalText converts values to a form suitable for transmission via XML etc.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalText() (text []byte, err error) {
	return i.marshalText(sweetMarshalTextRep, false)
}

// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to sweetMarshalTextRep.
func (i Sweet) MarshalJSON() ([]byte, error) {
	return enum.QuotedString(i.toString(sweetJSONStrings, sweetJSONIndex[:])), nil
}
`

const sweetMarshalText = `
func (i Sweet) marshalText(rep enum.Representation, quoted bool) (text []byte, err error) {
	var bs []byte
	switch rep {
	case enum.Number:
		return sweetMarshalNumber(i)
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
`

const sweetMarshalNumber_float = `
var sweetMarshalNumber = func(i Sweet) (text []byte, err error) {
	bs := []byte(strconv.FormatFloat(float64(i), 'g', 7, 64))
	return bs, nil
}
`

const sweetMarshalNumber_int = `
var sweetMarshalNumber = func(i Sweet) (text []byte, err error) {
	bs := []byte(strconv.FormatInt(int64(i), 10))
	return bs, nil
}
`

const sweetMarshalOrdinal = `
func (i Sweet) marshalOrdinal() (text []byte, err error) {
	bs := []byte(strconv.Itoa(i.Ordinal()))
	return bs, nil
}
`

func TestWriteMarshalText(t *testing.T) {
	testStage(t, basicModel().writeMarshalText, sweetMarshalTextRep+MarshalJSON_simple+
		sweetMarshalText+sweetMarshalNumber_int+sweetMarshalOrdinal)
	testStage(t, floatModelWithPrefixes().writeMarshalText, sweetMarshalTextRep+MarshalJSON_simple+
		sweetMarshalText+sweetMarshalNumber_float+sweetMarshalOrdinal)
	testStage(t, modelWithStructTags().writeMarshalText, sweetMarshalTextRep+MarshalJSON_struct_tags+
		sweetMarshalText+sweetMarshalNumber_int+sweetMarshalOrdinal)
	testStage(t, ignoreCase(modelWithStructTags()).writeMarshalText, sweetMarshalTextRep+MarshalJSON_struct_tags+
		sweetMarshalText+sweetMarshalNumber_int+sweetMarshalOrdinal)
}

//-------------------------------------------------------------------------------------------------

const UnmarshalText_all = `
// UnmarshalText converts transmitted values to ordinary values.
func (i *Sweet) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *Sweet) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return i.unmarshalJSON(s)
}
`

func TestWriteUnmarshalText(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeUnmarshalText, UnmarshalText_all)
	testStage(t, floatModelWithPrefixes().writeUnmarshalText, UnmarshalText_all)
}

//-------------------------------------------------------------------------------------------------

const unmarshalJSON_short = `
func (i *Sweet) unmarshalJSON(s string) error {
	return i.Parse(s)
}
`

const unmarshalJSON_nc = `
func (v *Sweet) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := sweetTransformInput(in)

	if v.parseString(s, sweetJSONStrings, sweetJSONIndex[:]) {
		return nil
	}

	var ok bool
	*v, ok = sweetAlias[s]
	if ok {
		return nil
	}

	return errors.New(in + ": unrecognised sweet")
}
`

const unmarshalJSON_ic = `
func (v *Sweet) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := sweetTransformInput(in)

	if v.parseString(s, sweetJSONInputs, sweetJSONIndex[:]) {
		return nil
	}

	var ok bool
	*v, ok = sweetAlias[s]
	if ok {
		return nil
	}

	return errors.New(in + ": unrecognised sweet")
}
`

func TestWriteUnmarshalJSON(t *testing.T) {
	testStage(t, unsnake(basicModel()).writeUnmarshalJSON, unmarshalJSON_short)
	testStage(t, modelWithStructTags().writeUnmarshalJSON, unmarshalJSON_nc)
	testStage(t, ignoreCase(modelWithStructTags()).writeUnmarshalJSON, unmarshalJSON_ic)
}

//-------------------------------------------------------------------------------------------------

const Scan_nc = `
// sweetStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var sweetStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Sweet) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch v := value.(type) {
	case int64:
		if sweetStoreRep == enum.Ordinal {
			*i = SweetOf(int(v))
		} else {
			*i = Sweet(v)
		}
		return nil
	case float64:
		*i = Sweet(v)
		return nil
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("%T %+v is not a meaningful sweet", value, value)
	}

	return i.parse(s, sweetStoreRep)
}
`

const Scan_lc = `
// sweetStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var sweetStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Sweet) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch v := value.(type) {
	case int64:
		if sweetStoreRep == enum.Ordinal {
			*i = SweetOf(int(v))
		} else {
			*i = Sweet(v)
		}
		return nil
	case float64:
		*i = Sweet(v)
		return nil
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("%T %+v is not a meaningful sweet", value, value)
	}

	return i.parse(s, sweetStoreRep)
}
`

const Scan_struct_tags = `
// sweetStoreRep controls database storage via the Scan and Value methods.
// By default, it is enum.Identifier and quoted strings are used.
var sweetStoreRep = enum.Identifier

// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (i *Sweet) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch v := value.(type) {
	case int64:
		if sweetStoreRep == enum.Ordinal {
			*i = SweetOf(int(v))
		} else {
			*i = Sweet(v)
		}
		return nil
	case float64:
		*i = Sweet(v)
		return nil
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("%T %+v is not a meaningful sweet", value, value)
	}

	if i.parseString(s, sweetSQLStrings, sweetSQLIndex[:]) {
		return nil
	}

	return errors.New(s + ": unrecognised sweet")
}
`

func TestWriteScanMethod(t *testing.T) {
	testStage(t, basicModel().writeScanMethod, Scan_nc)
	testStage(t, floatModelWithPrefixes().writeScanMethod, Scan_lc)
	testStage(t, modelWithStructTags().writeScanMethod, Scan_struct_tags)
	testStage(t, ignoreCase(modelWithStructTags()).writeScanMethod, Scan_struct_tags)
}

//-------------------------------------------------------------------------------------------------

const Value_int = `
// Value converts the Sweet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Sweet) Value() (driver.Value, error) {
	switch sweetStoreRep {
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
`

const Value_float = `
// Value converts the Sweet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Sweet) Value() (driver.Value, error) {
	switch sweetStoreRep {
	case enum.Number:
		return float64(i), nil
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.String(), nil
	}
}
`

const Value_struct_tags = `
// Value converts the Sweet to a string.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (i Sweet) Value() (driver.Value, error) {
	switch sweetStoreRep {
	case enum.Number:
		return int64(i), nil
	case enum.Ordinal:
		return int64(i.Ordinal()), nil
	case enum.Tag:
		return i.Tag(), nil
	default:
		return i.toString(sweetSQLStrings, sweetSQLIndex[:]), nil
	}
}
`

func TestWriteValueMethod(t *testing.T) {
	testStage(t, basicModel().writeValueMethod, Value_int)
	testStage(t, floatModelWithPrefixes().writeValueMethod, Value_float)
	testStage(t, modelWithStructTags().writeValueMethod, Value_struct_tags)
	testStage(t, ignoreCase(modelWithStructTags()).writeValueMethod, Value_struct_tags)
}

//-------------------------------------------------------------------------------------------------

func testStage(t *testing.T, fn func(io.Writer), expected string) {
	t.Helper()
	buf := &strings.Builder{}
	fn(buf)
	compare(t, buf.String(), expected)
}

func compare(t *testing.T, actual, expected string) {
	t.Helper()
	a := strings.Split(actual, "\n")
	b := strings.Split(expected, "\n")
	n := len(a)
	if n > len(b) {
		n = len(b)
	}
	ok := true
	for i := 0; i < n; i++ {
		if ok && a[i] != b[i] {
			ap, an, bp, bn := "", "", "", ""
			if i > 0 {
				ap = a[i-1]
				bp = b[i-1]
			}
			if i < n-1 {
				an = a[i+1]
				bn = b[i+1]
			}
			t.Errorf("Line %d\n--\n  %s\na:%s\n  %s\n--\n  %s\ne:%s\n  %s\n--", i+1, ap, a[i], an, bp, b[i], bn)
			ok = false
		}
	}
	if len(a) > len(b) {
		t.Errorf("Actual has %d more lines than expected.", len(a)-len(b))
		ok = false
	} else if len(a) < len(b) {
		t.Errorf("Actual has %d fewer lines than expected.", len(b)-len(a))
		ok = false
	}
	if !ok {
		t.Logf("%s\n", actual)
	}
}

//-------------------------------------------------------------------------------------------------

func basicModel() Model {
	Prefix = ""
	Suffix = ""
	return Model{
		Config: Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "int",
		Version:  util.Version,
		Values:   ValuesOf("Mars", "Bounty", "Snickers", "Kitkat", "Dairy_Milk"),
		Extra:    make(map[string]string),
	}
}

func floatModelWithPrefixes() Model {
	Prefix = "A"
	Suffix = "Bar"
	return Model{
		Config: Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "float64",
		Version:  util.Version,
		Values:   ValuesOf("AMarsBar", "ABountyBar", "ASnickersBar", "AKitkatBar", "ADairy_MilkBar"),
		Extra:    make(map[string]string),
	}
}

func modelWithStructTags() Model {
	Prefix = ""
	Suffix = ""
	var values Values
	values = values.Append("Mars", `json:"mmm" sql:"m"`)
	values = values.Append("Bounty", `json:"bbb" sql:"b"`)
	values = values.Append("Snickers", `json:"sss" sql:"s"`)
	values = values.Append("Kitkat", `json:"kkk" sql:"k"`)
	values = values.Append("Dairy_Milk", `json:"ddd" sql:"d"`)
	return Model{
		Config: Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:     "sweet",
		BaseType:   "int",
		Version:    util.Version,
		Values:     values,
		AliasTable: "sweetAlias",
		Extra:      make(map[string]string),
	}
}

func ignoreCase(m Model) Model {
	m.IgnoreCase = true
	return m
}

func unsnake(m Model) Model {
	m.Unsnake = true
	return m
}

func lowerCase(m Model) Model {
	m.Case = transform.Lower
	return m
}

func upperCase(m Model) Model {
	m.Case = transform.Upper
	return m
}

func lookupTables(m Model) Model {
	m.TagTable = "sweetNames"
	m.AliasTable = "sweetAlias"
	return m
}
