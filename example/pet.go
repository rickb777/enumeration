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

//go:generate enumeration -v -type Pet -prefix My -unsnake -lc -using petStrings -json

type Pet uint16

const (
	MyCat Pet = iota
	MyDog
	MyMouse
	MyElephant
	MyKoala_Bear
)

var petStrings = map[Pet]string{
	MyCat:        "Felis Catus",
	MyDog:        "Canis Lupus",
	MyMouse:      "Mus Musculus",
	MyElephant:   "Loxodonta Africana",
	MyKoala_Bear: "Phascolarctos Cinereus",
}
