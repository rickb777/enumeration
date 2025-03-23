package enum_test

import (
	"github.com/rickb777/expect"
	"strings"
	"testing"

	"github.com/rickb777/enumeration/v4/enum"
	"github.com/rickb777/enumeration/v4/example"
)

func TestIntEnums_Strings(t *testing.T) {
	methods := example.AllMethodEnums.Strings()

	expect.String(strings.Join(methods, "|")).ToBe(t, "HEAD|GET|PUT|POST|PATCH|DELETE")

	methods = enum.Enums{example.HEAD, example.PUT, example.PATCH}.Strings()

	expect.String(strings.Join(methods, "|")).ToBe(t, "HEAD|PUT|PATCH")
}

func TestIntEnums_Ordinals(t *testing.T) {
	days := enum.Enums{example.Wednesday, example.Friday, example.Sunday}.Ordinals()

	expect.Slice(days).ToBe(t, 3, 5, 0)

	days = enum.IntEnums{example.Wednesday, example.Friday, example.Sunday}.Ordinals()

	expect.Slice(days).ToBe(t, 3, 5, 0)
}

func TestIntEnums_Ints(t *testing.T) {
	days := enum.IntEnums{example.Wednesday, example.Friday, example.Sunday}.Ints()

	expect.Slice(days).ToBe(t, 4, 6, 1)
}

//-------------------------------------------------------------------------------------------------

func TestFloatEnums_Strings(t *testing.T) {
	es := example.AllBaseEnums.Strings()

	expect.Slice(es).ToBe(t, "a", "c", "g", "t")
}

func TestFloatEnums_Ordinals(t *testing.T) {
	es := enum.Enums{example.C, example.T}.Ordinals()

	expect.Slice(es).ToBe(t, 1, 3)

	es = enum.FloatEnums{example.C, example.T}.Ordinals()

	expect.Slice(es).ToBe(t, 1, 3)
}

func TestFloatEnums_Floats(t *testing.T) {
	es := example.AllBaseEnums.Floats()

	expect.Slice(es).ToBe(t, float64(example.A), float64(example.C), float64(example.G), float64(example.T))
}

func TestQuotedString(t *testing.T) {
	qs1 := enum.QuotedString("")
	expect.String(qs1).ToBe(t, []byte{'"', '"'})

	qs2 := enum.QuotedString("XYZ")
	expect.String(qs2).ToBe(t, []byte{'"', 'X', 'Y', 'Z', '"'})
}

func TestRepresentation(t *testing.T) {
	r1 := enum.MustParseRepresentation("number")
	expect.Number(r1).ToBe(t, enum.Number)

	r2 := enum.MustParseRepresentation("2")
	expect.Number(r2).ToBe(t, enum.Number)

	_, err := enum.AsRepresentation("foobar")
	expect.Error(err).ToHaveOccurred(t)

	num := enum.Number.String()
	expect.String(num).ToBe(t, "Number")

	o1 := enum.RepresentationOf(1)
	expect.Number(o1).ToBe(t, enum.Identifier)
	expect.Number(o1.Int()).ToBe(t, 1)
	expect.Bool(o1.IsValid()).ToBeTrue(t)

	o3 := enum.RepresentationOf(3)
	expect.Number(o1.Int()).ToBe(t, 1)
	expect.Bool(o3.IsValid()).ToBeFalse(t)

	expect.Bool(enum.None.IsValid()).ToBeTrue(t)
	expect.Bool(enum.Representation(4).IsValid()).ToBeFalse(t)
	expect.String(enum.Representation(4).String()).ToBe(t, "Representation(4)")
}
