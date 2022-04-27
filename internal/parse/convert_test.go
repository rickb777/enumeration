package parse

import (
	"bytes"
	. "github.com/benmoss/matchers"
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v2/internal/model"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"github.com/rickb777/enumeration/v2/internal/util"
	"testing"
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
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]string),
	}))
}

const enumBlock2 = `
/* inline comments are allowed */
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
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Extra:    make(map[string]string),
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
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]string),
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
	UsingTable = "sweetStrings"
	defer func() { UsingTable = "" }()

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
		LcType:   "sweet",
		BaseType: "int",
		Version:  util.Version,
		Values:   values,
		TagTable: "sweetStrings",
		Extra:    make(map[string]string),
	}))
}

//-------------------------------------------------------------------------------------------------

const enumBlockMultiple = `
type Sweet int
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
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]string),
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
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]string),
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
		Version:  util.Version,
		Values:   expected,
		Case:     transform.Upper,
		Extra:    make(map[string]string),
	}))
}

const enumSeparateMultiple = `
type Sweet int
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
		BaseType: "int",
		Version:  util.Version,
		Values:   model.ValuesOf("Mars", "Bounty", "Snickers", "Kitkat"),
		Case:     transform.Upper,
		Extra:    make(map[string]string),
	}))
}

//-------------------------------------------------------------------------------------------------

const enumError1 = `
type IgnoreMe int
// type Sweet is missing
const (
	Mars Sweet = iota
	Bounty
	Snickers
	Kitkat
)
const (
	Jam IgnoreMe = iota
	Toast
	Butter
)
`

func TestConvertError1(t *testing.T) {
	RegisterTestingT(t)
	s := bytes.NewBufferString(enumError1)
	_, err := Convert(s, "filename.go", transform.Stet,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err.Error()).Should(Equal("Failed to find Sweet in filename.go"))
}
