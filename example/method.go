package example

//go:generate enumeration -v -type Method -ic -store number

// Method: This example has json tags that control the JSON representations. So the parser
// recognises inputs from thes and the identifiers too. The '-ic' option means the parser ignores
// the case of its inputs.
//
// See also SalesChannel.
type Method uint

const (
	HEAD   Method = iota // json:"HE"
	GET                  // json:"GE"
	PUT                  // json:"PU"
	POST                 // json:"PO"
	PATCH                // json:"PA"
	DELETE               // json:"DE"
)
