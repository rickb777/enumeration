package example

// This example has a cross-mapping table specified via the '-using' option. So the parser
// recognises inputs from both sets of strings. The '-lc' option means the parser expects
// lowercase inputs and the String method gives lowercase values. Because of '-unsnake',
// underscores are replaced with spaces so "Koala_Bear" is treated as "koala bear".

//go:generate enumeration -v -type Pet -unsnake -lc -using petStrings

type Pet uint16

const (
	Cat Pet = iota
	Dog
	Mouse
	Elephant
	Koala_Bear
)

var petStrings = map[Pet]string{
	Cat:        "Felis Catus",
	Dog:        "Canis Lupus",
	Mouse:      "Mus Musculus",
	Elephant:   "Loxodonta Africana",
	Koala_Bear: "Phascolarctos Cinereus",
}
