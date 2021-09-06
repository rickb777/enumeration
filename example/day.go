package example

type Day uint

const (
	_ Day = iota // throw away zeroth so that Sunday is 1
	Sunday
	Monday
	Tuesday
	Wednesday
)

// It is allowable to break the constant blocks (although it's a bit
// unnecessary here), but the numeric values *must* all be distinct.
const (
	Thursday Day = iota + 5
	Friday
	Saturday
	numberOfDays = int(Saturday) // this is not exported
)
