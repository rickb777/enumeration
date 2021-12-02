package example

// This example has a cross-mapping table specified via the '-using' option. So the parser
// recognises inputs from both sets of strings. The '-ic' option means the parser ignores
// the case of its inputs.

//go:generate enumeration -v -type Method -ic -using methodStrings

type Method uint

const (
	HEAD Method = iota
	GET
	PUT
	POST
	PATCH
	DELETE
)

var methodStrings = map[Method]string{
	HEAD:   "HE",
	GET:    "GE",
	PUT:    "PU",
	POST:   "PO",
	PATCH:  "PA",
	DELETE: "DE",
}
