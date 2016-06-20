# Enumerations for Go

* Make simple Go enumerations work well for you using this easy-to-use code generator.

You could use stringer [here](https://github.com/golang/tools} or a related alterative
[here](https://github.com/clipperhouse/stringer). But they both depend on your code compiling
cleanly.

Sometime, the code generation gets in the way because you can't compile without the generated
code and you can't generate the code without compiling.

## First, write some Go

For example,

```Go
type Base int

const (
        A Base = iota
        C
        T
        G
)
```

There is [plenty more](http://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go)
on Stackoverflow.

