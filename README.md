# Enumerations for Go

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rickb777/enumeration/enum)
[![Issues](https://img.shields.io/github/issues/rickb777/enumeration.svg)](https://github.com/rickb777/enumeration/issues)

* Make simple Go enumerations work well for you using this easy-to-use code generator.

You could use stringer [here](https://github.com/golang/tools) or a related alterative
[here](https://github.com/clipperhouse/stringer). But they both depend on your code compiling
cleanly.

Sometimes, the code generation gets in the way because you can't compile without the generated
code and you can't generate the code without compiling.

**That's where this tool steps in**. It has a basic Go source code parser that finds idiomatic
enumerations, as shown below. It runs on any source code that's "clean enough", even if the
source code is still incomplete. Provided your `type` is ok and your `const` is complete, you'll be fine.

It will not handle C-style comments though, so these must not be present. The normal double-slash is OK.

## First, Install

```
go get github.com/rickb777/enumeration/v3
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

Another example, which also includes the optional `go:generate` line,

```Go
//go:generate enumeration -type Month

type Month uint

const (
	January, February, March    Month = 1, 2, 3
	April, May, June            Month = 4, 5, 6
	July, August, September     Month = 7, 8, 9
	October, November, December Month = 10, 11, 12
)
```

Note that the full Go language specification is flexible, allowing several alternative equivalent syntaxes for declaring constants. The `enumeration` tool only handles the most common two cases, illustrated above.

Above, `Month` is based on `uint`; the base type does not have to be an `int`; any integer or float base type can be used, for example

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

There is one other restriction: the type declaration must be *before* the constants, as it does in all the examples above. White-space should be canonically formatted using `gofmt` before using the tool (the results are unpredictable otherwise). (This is because of the limited parser used, for the reasons given earlier.)

Although the default behaviour is to generate strings for each enumeration value which match the constants you declared, you can take full control and override this with your own lookup table. The `-using` argument needed:

```Go
//go:generate enumeration -type Day -using shortDayNames

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

var shortDayNames = map[Day]string{
    Sunday:    "Su",
    Monday:    "Mo",
    Tuesday:   "Tu",
    Wednesday: "We",
    Thursday:  "Th",
    Friday:    "Fr",
}
```

Just occasionally, you might have several enumerations in the same package and you need to avoid their identifiers clashing. This is easy to control using the `-prefix` and `-suffix` options. Here's an example using a suffix; prefixes are similar and you can have both if you want.

```Go
//go:generate enumeration -lc -type SalesChannel -suffix Sales

type SalesChannel int

const (
    _              SalesChannel = iota
    OnlineSales                 // represented as "online"
    InstoreSales                // represented as "instore"
    TelephoneSales              // represented as "telephone"
)
```

If you want more control of the strings used for JSON and for SQL marshaling, structured comments can be used. In this example, the `String()` method, the JSON string and the SQL value will all be different values.

```Go
//go:generate enumeration -lc -type SalesChannel

type SalesChannel int

const (
    _              SalesChannel = iota
    Online                  // json:"webshop" sql:"o" -- String() is "online"
    Instore                 // json:"store"   sql:"s" -- String() is "instore"
    Telephone               // json:"phone"   sql:"t" -- String() is "telephone"
)
```

Structured comments are deliberately similar to Go `struct` tags. The supported tags are `text` (for MarshalText/UnmarshalText), `json` (for MarshalJSON/UnmarshalJSON) and `sql` (for Value/Scan); these can be used in any combination; also `all` sets a value for all of them. When present, these tags override the `-marshaltext`, `-marshaljson` and/or `-store` options described below. They also override the `-uc`, `-ic` and `-lc` case modifiers on the corresponding marshal methods.

## Next, Run The Tool

For example:

```
enumeration -i example/sample1.go -type Base -package example
```

Options are:

 * `-type <name>`
    - the name of the primary Go type for which code generation is being used.

 * `-plural <name>`
    - the plural equivalent for the name of the primary Go type. This is optional and the default is to use the type name and append letter 's', as is common for many English nouns. For example, with `-type Party`, you might usefully specify `-plural Parties`.

 * `-prefix <prefix>`
   - a prefix at the start of every identifier in the enumeration that is not included in the String representation.

 * `-suffix <suffix>`
   - a suffix at the end of every identifier in the enumeration that is not included in the String representation.

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

 * `-ic`
    - ignore the case when parsing, but don't convert the case of the string representations of the enumeration values.

 * `-unsnake`
    - convert underscores in identifiers to spaces (e.g. `Hello_world` becomes "Hello world"). When parsing, each underscore is treated as a space.

 * `-alias <map-name>`
    - declare a `var <map-name> = map[string]Type{ ... }` that gives aliases to be recognised during parsing. Each value of `Type` can have as many aliases as you need.

 * `-marshaltext <as>`
    - changes the way that text is marshaled (in JSON or XML) to be one of `identifier`, `number` or `ordinal` (can be overridden by the `text` struct tab, above).

 * `-marshaljson <as>`
    - changes the way that JSON is marshaled (not XML) to be one of `identifier`, `number` or `ordinal` (can be overridden by the `json` struct tab, above).

 * `-store <as>`
    - changes the way that values are stored in a DB to be one of `identifier`, `number` or `ordinal` (can be overridden by the `sql` struct tab, above).

 * `-lenient`
    - when the Parse method is given a number, this allows parsing to yield invalid values (normally parsing an unrecognised number will yield an error). Using this with `-marshaltext number` means the enumeration is an open set of which some values have names.

 * `-f`
    - force output generation; if this is not set, the output file is only produced when it is is absent or when it is older than the input file.

 * `-v`
    - verbose info messages

 * `-version`
    - print the version and exit.

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
    - Converts Day values into their ordinal numbers, i.e. the indexes indicating the order in which you declared the constants, starting from zero. These may happen to be the same as the values you chose, but need not be. For invalid Day values, `Ordinal()` returns -1.

 * `func DayOf(o int) Day`
    - Converts an ordinal to a Day value, if it can. The name of this function depends on the name of your type (`DayOf` in this example). The related type conversion `Day(i)` should be used when converting a *value* instead of an *ordinal*.

 * `func (d Day) IsValid() bool`
    - Tests whether a given value is one of the defined `Day` constants. Type conversion allows possibly out-of range values to be created; these can be tested with this method. `IsValid()` is related to `Ordinal()` because all valid values have an ordinal >= 0. 

 * `func (d Day) Int() int`
    - Converts Day values into their int values, i.e. just the value of the constant int. This is merely a type conversion to `int`, but conveniently matches the `enum.IntEnum` interface, allowing polymorphism. This method is only present when the base type is any integer type.

 * `func (d Day) Float() float64`
    - Converts Day values into their float values, i.e. just the value of the constant float. This is merely a type conversion to `float64`, but conveniently matches the `enum.FloatEnum` interface, allowing polymorphism. This method is only present when the base type is `float32` or `float64`.

 * `func (d *Day) Parse(s string) error`
    - Converts a string representation to a Day value, if it can, then assigns it to `d`. If `s` holds an integer, it is treated as a number (or possibly as an ordinal) and will result in the corresponding value. Numbers must be within the valid Day range unless `-lenient` was specified. Ordinals are not normally used but will be expected when the `-marshaltext ordinal` or `-store ordinal` options are specified.

 * `func AsDay(s string) (Day, error)`, `func MustParseDay(s string) Day`
    - Converts a string representation to a Day value, if it can. The function name depends on the name of your type (`AsDay` and `MustParseDay` in this example). These functions are a convenient wrapper for the `Parse` method.

 * `var AllDays = []Day{ ... }`
    - Provides all the `Day` values in a single slice. This is particularly useful if you need to iterate over them. Usually, the identifier name depends on the name of your type, but it can be overridden using `-plural`.

 * `var AllDayEnums = enum.IntEnums{ ... }`
    - Provides all the `Day` values in a single slice, held using an interface for polymorphism. The slice type would instead be `enum.FloatEnums` if the base type is `float32` or `float64`.

If you used the `-marshaltext` option or `text` struct tags, you will also get:

 * `encoding.TextMarshaler`, `encoding.TextUnmarshaler`
    - Provides methods to satisfy these interfaces so that your enumeration can be easily used by XML and other codecs in the standard Go library.

If you used the `-marshaljson` option or `json` struct tags, you will also get:

 * `json.Marshaler`, `json.Unmarshaler`
    - Provides methods to satisfy these interfaces so that your enumeration can be easily used by JSON codecs in the standard Go library.

If you used the `-store` option or `sql` struct tags, you will also get:

 * `sql.Scanner`, `driver.Valuer`
    - Provides methods to satisfy these two interfaces so that your enumeration can be easily used by SQL drivers in the standard Go library. Note that `driver.Valuer` is provided as a template for you to copy if you need it; otherwise the SQL driver will automatically make use of the numeric values of enumerations.

Other items generated:

 * `var dayMarshalNumber`
    - This unexported `var` is a function that converts values to strings using `strconv.FormatInt` or `strconv.FormatFloat`. Within the same package, you can replace this function with your own.

 * `func (d Day) Text()` and `func (d Day) JSON()`
    - These are present whenever `MarshalText` / `MarshalJSON` methods are generated and are using the identifier or a struct tag. They return the corresponding values as string (with silent errors). Note that `JSON()` does not quote strings, unlike `MarshalJSON`.

## Other Use Options

This tool is compatible with `go generate` - [more](https://blog.golang.org/generate). However, `go generate` may not always work if the code is still incomplete and doesn't yet compile, in which case you can just run `enumeration` directly on the command line.

## Credits

Thanks are due to others for their earlier work. Notable is the earlier stringer
[here](https://github.com/golang/tools) by Rob Pike or a related alterative
[here](https://github.com/clipperhouse/stringer) by Matt Sherman.
