# Enumerations for Go

* Make simple Go enumerations work well for you using this easy-to-use code generator.

You could use stringer [here](https://github.com/golang/tools) or a related alterative
[here](https://github.com/clipperhouse/stringer). But they both depend on your code compiling
cleanly.

Sometimes, the code generation gets in the way because you can't compile without the generated
code and you can't generate the code without compiling. That's where this tool steps in.

## First, write some Go

For example,

```Go
type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
```

this example comes from the [Go Language Reference](https://golang.org/ref/spec#Constant_declarations).

There is [plenty more](http://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go)
on Stackoverflow.

