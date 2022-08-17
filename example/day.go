package example

//go:generate enumeration -v -type Day

// Day: This simple example also shows that more than one 'const' block can be used
// provided that the integer values are all distinct.
type Day uint

const (
	_ Day = iota // throw away zeroth so that Sunday is 1
	Sunday
	Monday
	Tuesday
	Wednesday
)

// It is allowable to break the constant blocks (although it's a bit
// unnecessary here). However, the numeric values *must* all be distinct,
// otherwise the generated enum will not compile.
const (
	Thursday Day = iota + 5
	Friday
	Saturday
	numberOfDays = int(Saturday) // this is not exported
)
