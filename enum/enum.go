// package enum provides an API for manipulating enumerations.
package enum

// Enum is a generic contract for all generated enums.
type Enum interface {
	Ordinal() int
	String() string
	IsValid() bool
}
