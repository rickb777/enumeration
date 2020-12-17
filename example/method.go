package example

// This example shows an enumeration for which the lookup table is implemented here, instead
// of relying on the automatically-generated one. This allows arbitrary strings to represent
// their corresponding values.

type Method uint

const (
	_ Method = iota
	HEAD
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
