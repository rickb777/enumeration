package test

//go:generate enumeration -v -f -i season.go -type Season1 -suffix 1

type Season1 uint

const (
	_ Season1 = iota
	Spring1
	Summer1
	Autumn1
	Winter1
)

//-------------------------------------------------------------------------------------------------
//go:generate enumeration -v -f -i season.go -type Season2 -suffix 2 -alias season2Alias

type Season2 uint

const (
	_ Season2 = iota
	Spring2
	Summer2
	Autumn2
	Winter2
)

var season2Alias = map[string]Season2{
	"Fall2": Autumn2,
}
