package example

//go:generate enumeration -v -type Pet -prefix My -unsnake -lc -alias petAliases -lenient

// Pet: This example has a cross-mapping table specified via the 'text' tags. So the parser
// recognises inputs from both sets of strings. Also:
//
//   - The '-lc' option means the parser expects lowercase inputs and the String method
//     gives lowercase values.
//
//   - The -prefix option means each identifer, MyCat etc, is stored in the enumeration
//     as just "cat" without the "My" prefix.
//
//   - Because of '-unsnake', underscores are replaced with spaces so "MyKoala_Bear" is
//     treated as "koala bear".
//
//   - Because of the '-lenient' option, the parser will allow numbers outside the valid
//     range 0 to 4.
type Pet uint16

// These all have prefix "My", which is stripped from the String representation.
const (
	MyCat        Pet = iota // text:"Felis Catus"
	MyDog                   // text:"Canis Lupus"
	MyMouse                 // text:"Mus Musculus"
	MyElephant              // text:"Loxodonta Africana"
	MyKoala_Bear            // text:"Phascolarctos Cinereus"
)

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
