package enum_test

import (
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v2/enum"
	"github.com/rickb777/enumeration/v2/example"
	"strings"
	"testing"
)

func TestIntEnums_Strings(t *testing.T) {
	RegisterTestingT(t)

	days := example.AllDayEnums.Strings()

	Ω(strings.Join(days, "|")).Should(Equal("Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday"))

	days = enum.Enums{example.Wednesday, example.Friday, example.Sunday}.Strings()

	Ω(strings.Join(days, "|")).Should(Equal("Wednesday|Friday|Sunday"))
}

func TestIntEnums_Tags_fallback(t *testing.T) {
	RegisterTestingT(t)

	days := example.AllDayEnums.Tags()

	Ω(strings.Join(days, "|")).Should(Equal("Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday"))

	days = enum.Enums{example.Wednesday, example.Friday, example.Sunday}.Tags()

	Ω(strings.Join(days, "|")).Should(Equal("Wednesday|Friday|Sunday"))
}

func TestIntEnums_Tags_defined(t *testing.T) {
	RegisterTestingT(t)

	alphabet := example.AllGreekAlphabetEnums.Tags()

	Ω(strings.Join(alphabet, "|")).Should(Equal("alpha|beta|gamma|delta|epsilon|zeta|eta|theta|iota|kappa|" +
		"lambda|mu|nu|xi|omicron|pi|rho|sigma|tau|upsilon|phi|chi|psi|omega"))

	alphabet = enum.Enums{example.Αλφα, example.Δέλτα, example.Ωμέγα}.Tags()

	Ω(strings.Join(alphabet, "|")).Should(Equal("alpha|delta|omega"))
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

func TestFloatEnums_Tags(t *testing.T) {
	RegisterTestingT(t)

	es := example.AllBaseEnums.Tags()

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
