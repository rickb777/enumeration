package parse

import (
	"bytes"
	"flag"
	"go/types"
	"os"
	"testing"

	. "github.com/benmoss/matchers"
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v3/internal/collection"
	"github.com/rickb777/enumeration/v3/internal/model"
	"github.com/rickb777/enumeration/v3/internal/transform"
	"github.com/rickb777/enumeration/v3/internal/util"
)

const enumBlock1 = `
type Sweet int
const (
	Mars Sweet = iota + 1
	Bounty
	Snickers
	Kitkat
)
`

func TestConvertBlock1(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumBlock1)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: false,
			Unsnake:    false,
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "int",
		BaseKind: types.Int,
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]interface{}),
		Imports:  collection.NewStringSet(basicImports...),
	}))
}

const enumBlock2 = `
/* inline comments are allowed, also var declarations are ignored */
var x = 100 *
		100
type Sweet int // <-- buried here
const (
	_ Sweet = iota
	Mars
	Bounty    // coconuts and more
	Snickers  // I need this

	// yum yum
	Kitkat
)
// as are blank lines
`

func TestConvertBlock2(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumBlock2)
	m, err := Convert(s, "filename.go", transform.Stet,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		},
		LcType:   "sweet",
		BaseType: "int",
		BaseKind: types.Int,
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Extra:    make(map[string]interface{}),
		Imports:  collection.NewStringSet(basicImports...),
	}))
}

const enumBlock3 = `
type Sweet int

const (
	Mars     Sweet = 0 // zero -> default
	Bounty   Sweet = 1
	Snickers Sweet = 2
	Kitkat   Sweet = 3
	Ignore         = "nothing"
)
`

func TestConvertBlock3(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumBlock3)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "int",
		BaseKind: types.Int,
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]interface{}),
		Imports:  collection.NewStringSet(basicImports...),
	}))
}

const enumBlock4 = `
type IgnoreMe int
var s = "123"
type Sweet int // <-- buried here
type Transformer struct {
	V int
}
type Bar interface {
	X()
}
var x = 0
const (
	Jam IgnoreMe = iota
	Toast
	Butter
)
var y = 1
const (
	_ Sweet = iota
	ignored  //
	Mars     // json:"mmm"
	ignored2 //
	Bounty   // json:"bbb"
	Snickers // json:"sss"
	ignored4 //
	Kitkat   // json:"kkk"
	numberOfSweets = int(Kitkat) // this is not exported
)
var sweetStrings = map[Sweet]string{
	Mars:     "peanut",
	Bounty:   "coconut",
	Snickers: "toffee",
	Kitkat:   "biscuit",
}
`

func TestConvertBlock4(t *testing.T) {
	RegisterTestingT(t)
	AliasTable = "sweetAliases"
	defer func() { AliasTable = "" }()

	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumBlock4)
	m, err := Convert(s, "filename.go", transform.Stet,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		})
	Ω(err).Should(BeNil())

	values := model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat")
	values[0].JSON = "mmm"
	values[1].JSON = "bbb"
	values[2].JSON = "sss"
	values[3].JSON = "kkk"

	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		},
		LcType:     "sweet",
		BaseType:   "int",
		BaseKind:   types.Int,
		Version:    util.Version,
		Values:     values,
		AliasTable: "sweetAliases",
		Extra:      make(map[string]interface{}),
		Imports:    collection.NewStringSet(basicImports...),
	}))
}

//-------------------------------------------------------------------------------------------------

const enumBlockMultiple = `
type (
	Sweet int
)
const (
	Mars, Bounty, Snickers, Kitkat Sweet = 1, 2, 3, 4
)
`

func TestConvertBlockMultiple(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumBlockMultiple)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: false,
			Unsnake:    false,
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "int",
		BaseKind: types.Int,
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]interface{}),
		Imports:  collection.NewStringSet(basicImports...),
	}))
}

//-------------------------------------------------------------------------------------------------

const enumSeparate1 = `
type Sweet int

const Mars     Sweet = 1
const Bounty   Sweet = 2
const Other    int   = 3
const Snickers Sweet = 4
const Kitkat   Sweet = 5
`

func TestConvertSeparate1(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumSeparate1)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "int",
		BaseKind: types.Int,
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]interface{}),
		Imports:  collection.NewStringSet(basicImports...),
	}))
}

const enumSeparate2 = `
type Sweet int

const Mars     Sweet = 0 // json:"toffee" zero -> default
const Bounty   Sweet = 1 // json:"coconut"
const Other    int   = 5 // json:"hazlenut"
const Snickers Sweet = 2 // json:"peanut"
const Kitkat   Sweet = 3 // json:"biscuit"
const Ignore         = 6 // json:"orange"
`

func TestConvertSeparate2(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumSeparate2)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err).Should(BeNil())

	var expected model.Values
	expected = expected.Append("Mars", `json:"toffee"`)
	expected = expected.Append("Bounty", `json:"coconut"`)
	expected = expected.Append("Snickers", `json:"peanut"`)
	expected = expected.Append("Kitkat", `json:"biscuit"`)

	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "int",
		BaseKind: types.Int,
		Version:  util.Version,
		Values:   expected,
		Case:     transform.Upper,
		Extra:    make(map[string]interface{}),
		Imports:  collection.NewStringSet(basicImports...),
	}))
}

const enumSeparateMultiple = `
type Sweet float64
const Mars, Bounty, Snickers, Kitkat Sweet = 1, 2, 3, 4
`

func TestConvertSeparateMultiple(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enumSeparateMultiple)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(DeepEqual(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:   "sweet",
		BaseType: "float64",
		BaseKind: types.Float64,
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]interface{}),
		Imports:  collection.NewStringSet(basicImports...),
	}))
}

//-------------------------------------------------------------------------------------------------

var enumErrors = map[string]string{
	// type Sweet is missing
	`type Sweet uint
	type Sweet int
	const (
		Mars Sweet = iota
	)
	`: "filename.go:2:7: found multiple type Sweet declarations",

	// type Sweet is missing
	`type IgnoreMe int
	const (
		Mars Sweet = iota
	)
	const (
		Jam IgnoreMe = iota
	)
	`: "filename.go: failed to find type Sweet",

	// type Sweet is not numeric - simple
	`type Sweet string
	const (
		Mars Sweet = iota
	)
	const (
		Jam IgnoreMe = iota
	)
	`: "filename.go:1:12: enumeration type Sweet must be an integer or float type",

	// type Sweet is not numeric - block
	`type (
		Sweet string
	)
	const (
		Mars Sweet = iota
	)
	`: "filename.go:2:9: enumeration type Sweet must be an integer or float type",

	// type Sweet is a type alias
	`type Sweet = Alias
	const (
		Mars Sweet = iota
	)
	`: "filename.go:1:12: type Sweet is a type alias (not supported)",

	// type Sweet is a type alias - block
	`type (
		Sweet = Alias
	)
	const (
		Mars Sweet = iota
	)
	`: "filename.go:2:9: type Sweet is a type alias (not supported)",

	// type Sweet is a type alias - block
	`type (
		Sweet
	)
	`: "filename.go:2:8: syntax error in type Sweet declaration",
}

func TestConvertErrors(t *testing.T) {
	RegisterTestingT(t)
	for src, msg := range enumErrors {
		s := bytes.NewBufferString(src)
		_, err := Convert(s, "filename.go", transform.Stet,
			model.Config{
				MainType: "Sweet",
				Plural:   "Sweets",
				Pkg:      "confectionary",
			})
		Ω(err.Error()).Should(Equal(msg), msg)
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	util.Verbose = testing.Verbose()
	util.Dbg = testing.Verbose()
	os.Exit(m.Run())
}
