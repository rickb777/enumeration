package example

//go:generate enumeration -v -i channel.go -o channel_enum.go -lc -type SalesChannel -suffix Sales

// preamble const declarations are ignored
const IgnoreThisItem, AndThis = 741, "quack"

// this is ignored too
const (
	One = 1
)

// SalesChannel: The example demonstrates the removing of a suffix string from the identifiers
// when their string equivalent is accessed.
//
// The `json` tags in comments control values used for JSON marshalling.
// The `sql` tags in comments control values used for SQL storage.
type SalesChannel int

const (
	OnlineSales    SalesChannel = 3 // json:"webshop" sql:"o" -- String() is "online"
	InstoreSales   SalesChannel = 5 // json:"store"   sql:"s" -- String() is "instore"
	TelephoneSales SalesChannel = 7 // json:"phone"   sql:"t" -- String() is "telephone"
)
