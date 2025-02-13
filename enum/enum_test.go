package enum_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v4/enum"
	"github.com/rickb777/enumeration/v4/example"
)

func TestIntEnums_Strings(t *testing.T) {
	RegisterTestingT(t)

	methods := example.AllMethodEnums.Strings()

	Ω(strings.Join(methods, "|")).Should(Equal("HEAD|GET|PUT|POST|PATCH|DELETE"))

	methods = enum.Enums{example.HEAD, example.PUT, example.PATCH}.Strings()

	Ω(strings.Join(methods, "|")).Should(Equal("HEAD|PUT|PATCH"))
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

func TestQuotedString(t *testing.T) {
	RegisterTestingT(t)

	qs1 := enum.QuotedString("")
	Ω(qs1).Should(Equal([]byte{'"', '"'}))

	qs2 := enum.QuotedString("XYZ")
	Ω(qs2).Should(Equal([]byte{'"', 'X', 'Y', 'Z', '"'}))
}

func TestRepresentation(t *testing.T) {
	RegisterTestingT(t)

	r1 := enum.MustParseRepresentation("number")
	Ω(r1).Should(Equal(enum.Number))

	r2 := enum.MustParseRepresentation("2")
	Ω(r2).Should(Equal(enum.Number))

	_, err := enum.AsRepresentation("foobar")
	Ω(err).Should(HaveOccurred())

	num := enum.Number.String()
	Ω(num).Should(Equal("Number"))

	o1 := enum.RepresentationOf(1)
	Ω(o1).Should(Equal(enum.Identifier))
	Ω(o1.Int()).Should(Equal(1))
	Ω(o1.IsValid()).Should(BeTrue())

	o3 := enum.RepresentationOf(3)
	Ω(o1.Int()).Should(Equal(1))
	Ω(o3.IsValid()).Should(BeFalse())

	Ω(enum.None.IsValid()).Should(BeTrue())
	Ω(enum.Representation(4).IsValid()).Should(BeFalse())
	Ω(enum.Representation(4).String()).Should(Equal("Representation(4)"))
}
