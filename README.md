# Enumerations for Go

* Make simple Go enumerations work well for you using this easy-to-use code generator.

You could use stringer [here](https://github.com/golang/tools) or a related alterative
[here](https://github.com/clipperhouse/stringer). But they both depend on your code compiling
cleanly.

Sometimes, the code generation gets in the way because you can't compile without the generated
code and you can't generate the code without compiling.

**That's where this tool steps in**. It has a basic Go source code parser that finds idiomatic
interfaces, as shown below. It runs on any source code that's "clean enough", even if the
source code is still incomplete. Provided your `type` is ok and yoyur `const` is complete, you'll be fine.

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

## Next, Run The Tool

For example:

```
enumeration -i example/sample1.go -type Base -package example
```

Options are:

 * `-type <name>`
    - the name of the primary Go type for which code generation is being used.

 * `-plural <name>`
    - the plural equivalent for the name of the primary Go type. This is optional and the default is
        to use the type name and append letter 's'.

 * `-i <name>` or `-input <name>`
    - the name of the input Go file containing the `type` and `const` declarations.

 * `-o <name>` or `-output <name>`
    - the name of the output Go file to be written. If omitted, `<type>_enum.go` is used.

 * `-package <name>`
    - the name of the Go package. If omitted, the directory of the output will be used (which becomes mandatory
        therefore).

 * `-f`
    - force output generation; if this is not set the output file is only produced when it is older than the
      dependencies

 * `-v`
    - verbose info messages

The option parser will also infer the template and output file names, so it is also permitted to use

```
enumeration -type Base
```

when this matches your needs.

## Generated Go

The generated code complements your `type` and `const` definitions as follows. Let's assume that you wrote
the `Day` type above. You will get:

 * `func (d Day) String() string`
    - Converts Day values to strings and satisfies the well-known `Stringer` interface.

 * `func (d Day) Ordinal() int`
    - Converts Day values into their ordinal numbers, i.e. the indexes indicating the order in which you declared
        the constants, starting from zero. These may happen to be the same as the values you chose, but need not be.

 * `func AsDay(s string) (Day, error)`
    - Converts a string representation to a Day value, if it can. The name of this function depends on the name
        of your type.

 * `var AllDays = []string{ ... }`
    - Provides all the Day values in a single slice. This is useful if you need to iterate, for example. The
        name depends on the name of your type, although it can be overridden using `-plural`.
