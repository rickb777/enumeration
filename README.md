# Enumerations for Go

* Make simple Go enumerations work well for you using this easy-to-use code generator.

You could use stringer [here](https://github.com/golang/tools) or a related alterative
[here](https://github.com/clipperhouse/stringer). But they both depend on your code compiling
cleanly.

Sometimes, the code generation gets in the way because you can't compile without the generated
code and you can't generate the code without compiling.

**That's where this tool steps in**. It has a basic Go source code parser that finds idiomatic
interfaces, as shown below. It runs on any source code that's "clean enough", even if the
source code is still incomplete. Provided your `type` is ok and your `const` is complete, you'll be fine.

It will not handle C-style comments though, so these must not be present. The normal double-slash is OK.

## First, Install

```
go get bitbucket.org/rickb777/enumeration
```
You should see that the `enumeration` binary is now in the bin folder on your GOPATH. Make sure this is
on your PATH so it can be run.

## Now, write some Go

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
	numberOfDays  // this constant is not exported but is included in the enumeration
)
```

this example comes from the [Go Language Reference](https://golang.org/ref/spec#Constant_declarations). 

Another example,

```Go
type Month uint

const (
	January, February, March    Month = 1, 2, 3
	April, May, June            Month = 4, 5, 6
	July, August, September     Month = 7, 8, 9
	October, November, December Month = 10, 11, 12
)
```

Note that the full language specification is flexible, allowing several alternative equivalent syntaxes for declaring constants. The `enumeration` tool only handles the most common two cases, illustrated above.

The base type does not have to be an integer; other base types can be used, for example

```Go
type Base float32

const (
	// Nucleotide Molecular Weights, g/mol
	A Base = 331.2
	C Base = 307.2
	G Base = 347.2
	T Base = 322.2
)
```

There is one other restriction: the type declaration must be *before* the constants, as it does in all the examples above. White-space should be canonically formatted using `gofmt` before using the tool (the results are unpredictable otherwise).

## Next, Run The Tool

For example:

```
enumeration -i example/sample1.go -type Base -package example
```

Options are:

 * `-type <name>`
    - the name of the primary Go type for which code generation is being used.

 * `-plural <name>`
    - the plural equivalent for the name of the primary Go type. This is optional and the default is to use the type name and append letter 's', as is common for many English nouns. For example, with `-type Party`, you might usefuly specify `-plural Parties`.

 * `-i <name>` or `-input <name>`
    - the name of the input Go source file containing the `type` and `const` declarations.

 * `-o <name>` or `-output <name>`
    - the name of the output Go source file to be written. If omitted, `<type>_enum.go` is used.

 * `-package <name>`
    - the name of the Go package. If omitted, the directory of the output file will be used if specified, or the current directory name otherwise.

 * `-lc`
    - convert to lower case the string representations of the enumeration values.

 * `-uc`
    - convert to upper case the string representations of the enumeration values.

 * `-f`
    - force output generation; if this is not set the output file is only produced when it is is absent or when it is older than the input file.

 * `-v`
    - verbose info messages

The option parser will also infer the source and output file names, so it is also permitted to use

```
enumeration -type Base
```

when this matches your needs. This example would try to read from `base.go` and write to `base_enum.go`.

## Generated Go

The generated code complements your `type` and `const` definitions as follows. Let's assume that you wrote
the `Day` type above. You will get:

 * `func (d Day) String() string`
    - Converts Day values to strings and satisfies the well-known `Stringer` interface. The strings are the human-readable names, as written in the list of constants.

 * `func (d Day) Ordinal() int`
    - Converts Day values into their ordinal numbers, i.e. the indexes indicating the order in which you declared the constants, starting from zero. These may happen to be the same as the values you chose, but need not be.

 * `func AsDay(s string) (Day, error)`
    - Converts a string representation to a Day value, if it can. The name of this function depends on the name of your type (`AsDay` in this example).

 * `var AllDays = []Day{ ... }`
    - Provides all the `Day` values in a single slice. This is particularly useful if you need to iterate over them. Usually, the identifier name depends on the name of your type, but it can be overridden using `-plural`.

## Other Use Options

This tool is compatible with `go generate` - [more](https://blog.golang.org/generate).
