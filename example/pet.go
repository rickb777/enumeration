package example

// This example has a cross-mapping table specified via the '-using' option. So the parser
// recognises inputs from both sets of strings. Also:
//
//   * The '-lc' option means the parser expects lowercase inputs and the String method
//     gives lowercase values.
//
//  * The -prefix option means each identifer, MyCat etc, is stored in the enumeration
//     as just "cat" without the "My" prefix.
//
//  * Because of '-unsnake', underscores are replaced with spaces so "MyKoala_Bear" is
//     treated as "koala bear".

//go:generate enumeration -v -type Pet -prefix My -unsnake -lc -using petTags -marshaltext tag -alias petAliases -lenient

type Pet uint16

// These all have prefix "My", which is stripped from the String representation.
const (
	MyCat Pet = iota
	MyDog
	MyMouse
	MyElephant
	MyKoala_Bear
)

// petTags is used for the Tag method.
var petTags = map[Pet]string{
	MyCat:        "Felis Catus",
	MyDog:        "Canis Lupus",
	MyMouse:      "Mus Musculus",
	MyElephant:   "Loxodonta Africana",
	MyKoala_Bear: "Phascolarctos Cinereus",
}

// petAliases provide more strings that are recognised during parsing.
// Although the map keys must be unique, the values do not need to be.
// Note that -lc means the keys here mus also be lowercase.
var petAliases = map[string]Pet{
	"sid":      MyCat,
	"diego":    MyCat,
	"pooch":    MyDog,
	"whiskers": MyMouse,
	"faithful": MyElephant,
	"cuddly":   MyKoala_Bear,
}
