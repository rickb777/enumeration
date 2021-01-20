// package enum provides an API for manipulating enumerations.
package enum

// Enum is a generic contract for all generated enums.
type Enum interface {
	Ordinal() int
	String() string
	Tag() string
	IsValid() bool
}

// IntEnum is a specialisation for those enums that have int or similar as the underlying type.
type IntEnum interface {
	Enum
	Int() int
}

// FloatEnum is a specialisation for those enums that have float32 or float64 as the underlying type.
type FloatEnum interface {
	Enum
	Float() float64
}

//-------------------------------------------------------------------------------------------------

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

// Ordinals gets the ordinal values of the enums in the same order.
func (es Enums) Ordinals() []int {
	os := make([]int, len(es))
	for i, e := range es {
		os[i] = e.Ordinal()
	}
	return os
}

//-------------------------------------------------------------------------------------------------

// IntEnums is a slice of IntEnum.
type IntEnums []IntEnum

// Strings gets the string values of the enums in the same order.
func (es IntEnums) Strings() []string {
	ss := make([]string, len(es))
	for i, e := range es {
		ss[i] = e.String()
	}
	return ss
}

// Ordinals gets the ordinal values of the enums in the same order.
func (es IntEnums) Ordinals() []int {
	os := make([]int, len(es))
	for i, e := range es {
		os[i] = e.Ordinal()
	}
	return os
}

// Ints gets the values of the enums in the same order.
func (es IntEnums) Ints() []int {
	vs := make([]int, len(es))
	for i, e := range es {
		vs[i] = e.Int()
	}
	return vs
}

//-------------------------------------------------------------------------------------------------

// FloatEnums is a slice of FloatEnum.
type FloatEnums []FloatEnum

// Strings gets the string values of the enums in the same order.
func (es FloatEnums) Strings() []string {
	ss := make([]string, len(es))
	for i, e := range es {
		ss[i] = e.String()
	}
	return ss
}

// Ordinals gets the ordinal values of the enums in the same order.
func (es FloatEnums) Ordinals() []int {
	os := make([]int, len(es))
	for i, e := range es {
		os[i] = e.Ordinal()
	}
	return os
}

// Floats gets the values of the enums in the same order.
func (es FloatEnums) Floats() []float64 {
	vs := make([]float64, len(es))
	for i, e := range es {
		vs[i] = e.Float()
	}
	return vs
}
