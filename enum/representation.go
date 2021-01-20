package enum

type Representation int

const (
	Identifier Representation = iota
	Tag
	Number // the value of the enumerant as a decimal number
	Ordinal
)
