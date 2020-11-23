package main

import "strings"

type Transformer struct {
	Fn   func(string) string
	Str  string
	Info string
}

var noChange = Transformer{
	Fn:   noop,
	Str:  "",
	Info: "",
}

var toUpper = Transformer{
	Fn:   strings.ToUpper,
	Str:  "strings.ToUpper(s)",
	Info: "The case of s does not matter.",
}

var toLower = Transformer{
	Fn:   strings.ToLower,
	Str:  "strings.ToLower(s)",
	Info: "The case of s does not matter.",
}

var xUnsnake = Transformer{
	Fn:   stringUnsnake,
	Str:  `strings.ReplaceAll(s, "_", " ")`,
	Info: "",
}

func noop(s string) string {
	return s
}

func stringUnsnake(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}
