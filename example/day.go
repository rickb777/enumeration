package example

type Day uint

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // this constant is not exported
)
