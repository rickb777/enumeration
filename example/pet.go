package example

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
