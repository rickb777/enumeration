package example

type Day uint

const (
	_ Day = iota // throw away zeroth so that Sunday is 1
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays = int(Saturday) // this is not exported
)
