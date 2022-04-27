package example

// The example demonstrates the removing of a suffix string from the identifiers
// when their string equivalent is accessed.
//
// The `json` tags in comments control values used for JSON marshalling.
// The `sql` tags in comments control values used for SQL storage.

//go:generate enumeration -v -i channel.go -o channel_enum.go -lc -type SalesChannel -suffix Sales

type SalesChannel int

const (
	_              SalesChannel = iota
	OnlineSales                 // json:"webshop" sql:"o" -- String() is "online"
	InstoreSales                // json:"store"   sql:"s" -- String() is "instore"
	TelephoneSales              // json:"phone"   sql:"t" -- String() is "telephone"
)
