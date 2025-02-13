package example

//go:generate enumeration -v -type Month -ic -marshaltext identifier

// Month example shows that several comma-separated enumeration constants can
// be on each line, and they can have explicit values. The '-ic' option means
// the parser ignores the case of its inputs.
type Month uint

const (
	January, February, March    Month = 1, 2, 3
	April, May, June            Month = 4, 5, 6
	July, August, September     Month = 7, 8, 9
	October, November, December Month = 10, 11, 12
)
