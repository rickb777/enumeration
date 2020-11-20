// package enum provides an API for manipulating enumerations.
package enum

// Enum is a generic contract for all generated enums.
type Enum interface {
	Ordinal() int
	String() string
	IsValid() bool
}

// Enums is a slice of Enum.
type Enums []Enum

// Strings gets the string values of the enums in the same order.
func (es Enums) Strings() []string {
	ss := make([]string, len(es))
	for i, e := range es {
		ss[i] = e.String()
	}
	return ss
}

// Strings gets the ordinal values of the enums in the same order.
func (es Enums) Ordinals() []int {
	ss := make([]int, len(es))
	for i, e := range es {
		ss[i] = e.Ordinal()
	}
	return ss
}
