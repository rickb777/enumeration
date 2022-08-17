package enum_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v3/enum"
	"github.com/rickb777/enumeration/v3/example"
)

func TestIntEnums_Strings(t *testing.T) {
	RegisterTestingT(t)

	days := example.AllDayEnums.Strings()

	Ω(strings.Join(days, "|")).Should(Equal("Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday"))

	days = enum.Enums{example.Wednesday, example.Friday, example.Sunday}.Strings()

	Ω(strings.Join(days, "|")).Should(Equal("Wednesday|Friday|Sunday"))
}

func TestIntEnums_Ordinals(t *testing.T) {
	RegisterTestingT(t)

	days := enum.Enums{example.Wednesday, example.Friday, example.Sunday}.Ordinals()

	Ω(days).Should(Equal([]int{3, 5, 0}))

	days = enum.IntEnums{example.Wednesday, example.Friday, example.Sunday}.Ordinals()

	Ω(days).Should(Equal([]int{3, 5, 0}))
}

func TestIntEnums_Ints(t *testing.T) {
	RegisterTestingT(t)

	days := enum.IntEnums{example.Wednesday, example.Friday, example.Sunday}.Ints()

	Ω(days).Should(Equal([]int{4, 6, 1}))
}

//-------------------------------------------------------------------------------------------------

func TestFloatEnums_Strings(t *testing.T) {
	RegisterTestingT(t)

	es := example.AllBaseEnums.Strings()

	Ω(es).Should(HaveLen(4))
	Ω(strings.Join(es, "|")).Should(Equal("a|c|g|t"))
}

func TestFloatEnums_Ordinals(t *testing.T) {
	RegisterTestingT(t)

	es := enum.Enums{example.C, example.T}.Ordinals()

	Ω(es).Should(Equal([]int{1, 3}))

	es = enum.FloatEnums{example.C, example.T}.Ordinals()

	Ω(es).Should(Equal([]int{1, 3}))
}

func TestFloatEnums_Floats(t *testing.T) {
	RegisterTestingT(t)

	es := example.AllBaseEnums.Floats()

	Ω(es).Should(Equal([]float64{float64(example.A), float64(example.C), float64(example.G), float64(example.T)}))
}
