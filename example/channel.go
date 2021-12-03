package example

// The example demonstrates the removing of a suffix string from the identifiers
// when their string equivalent is accessed.

//go:generate enumeration -v -i channel.go -o channel_enum.go -lc -type SalesChannel -suffix Sales

type SalesChannel int

const (
	_              SalesChannel = iota
	OnlineSales                 // represented as "online"
	InstoreSales                // represented as "instore"
	TelephoneSales              // represented as "telephone"
)
