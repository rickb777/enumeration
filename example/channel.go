package example

//go:generate enumeration -v -i channel.go -o channel_enum.go -lc -type SalesChannel -suffix Sales -poly

// preamble const declarations are ignored by the enumeration tool
const IgnoreThisItem, AndThis = 741, "quack"

const (
	// One is ignored by the enumeration tool too
	One = 1
)

// SalesChannel example demonstrates the removing of a suffix string from the identifiers
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
