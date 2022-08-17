package enum

type Representation int

const (
	None Representation = iota // disables the feature (new in v3)
	Identifier
	Tag    // deprecated (v2 only)
	Number // the value of the enumerant as a decimal number
	Ordinal
)
