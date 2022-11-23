package enum

//go:generate enumeration -type Representation -ic

type Representation int

const (
	None       Representation = iota // disables the feature (new in v3)
	Identifier                       // uses the main identifier of the corresponding constant
	Number                           // the value of the enumerant as a decimal number
)
