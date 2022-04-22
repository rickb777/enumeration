package parse

import (
	"bytes"
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v2/internal/model"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"github.com/rickb777/enumeration/v2/internal/util"
	"testing"
)

const enum1 = `
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

func TestScanValuesHappy(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enum1)
	m, err := Convert(s, "filename.go", transform.Stet,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model.Model{
		Config: model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		},
		LcType:       "sweet",
		BaseType:     "int",
		Version:      util.Version,
		Values:       []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		DefaultValue: "",
		Case:         0,
		S1:           "",
		S2:           "",
		TagTable:     "",
	}))
}

const enum3 = `
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
	ignored
	Mars
	ignored2
	Bounty
	Snickers
	ignored4
	Kitkat
	numberOfSweets = int(Kitkat) // this is not exported
)
var sweetStrings = map[Sweet]string{
	Mars:     "peanut",
	Bounty:   "coconut",
	Snickers: "toffee",
	Kitkat:   "biscuit",
}
`

func TestConvertHappy3(t *testing.T) {
	RegisterTestingT(t)
	UsingTable = "sweetStrings"
	defer func() { UsingTable = "" }()

	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enum3)
	m, err := Convert(s, "filename.go", transform.Stet,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model.Model{
		Config: model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: true,
			Unsnake:    true,
		},
		LcType:       "sweet",
		BaseType:     "int",
		Version:      util.Version,
		Values:       []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		DefaultValue: "",
		Tags:         map[string]string{"Mars": "peanut", "Bounty": "coconut", "Snickers": "toffee", "Kitkat": "biscuit"},
		Case:         0,
		S1:           "",
		S2:           "",
		TagTable:     "sweetStrings",
	}))
}

const enum4 = `
type Sweet int
const (
	Mars, Bounty, Snickers, Kitkat Sweet = 1, 2, 3, 4
)
`

func TestConvertHappy4(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enum4)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: false,
			Unsnake:    false,
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:       "sweet",
		BaseType:     "int",
		Version:      util.Version,
		Values:       []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		DefaultValue: "",
		Case:         transform.Upper,
		S1:           "",
		S2:           "",
		TagTable:     "",
	}))
}

const enum5 = `
type Sweet int
const (
	Mars Sweet = iota + 1
	Bounty
	Snickers
	Kitkat
)
`

func TestConvertHappy5(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enum5)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType:   "Sweet",
			Plural:     "Sweets",
			Pkg:        "confectionary",
			IgnoreCase: false,
			Unsnake:    false,
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:       "sweet",
		BaseType:     "int",
		Version:      util.Version,
		Values:       []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		DefaultValue: "",
		Case:         transform.Upper,
		S1:           "",
		S2:           "",
		TagTable:     "",
	}))
}

const enum6 = `
type Sweet int

const Mars     Sweet = 1
const Bounty   Sweet = 2
const Other    int   = 3
const Snickers Sweet = 4
const Kitkat   Sweet = 5
`

func TestConvertHappy6(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enum6)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:       "sweet",
		BaseType:     "int",
		Version:      util.Version,
		Values:       []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		DefaultValue: "",
		Case:         transform.Upper,
		S1:           "",
		S2:           "",
		TagTable:     "",
	}))
}

const enum7 = `
type Sweet int

const Mars     Sweet = 0 // zero -> default
const Bounty   Sweet = 1
const Other    int   = 5
const Snickers Sweet = 2
const Kitkat   Sweet = 3
const Ignore         = 6
`

func TestConvertHappy7(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enum7)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:       "sweet",
		BaseType:     "int",
		Version:      util.Version,
		Values:       []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		DefaultValue: "Mars",
		Case:         transform.Upper,
		S1:           "",
		S2:           "",
		TagTable:     "",
	}))
}

const enum8 = `
type Sweet int

const (
	Mars     Sweet = 0 // zero -> default
	Bounty   Sweet = 1
	Snickers Sweet = 2
	Kitkat   Sweet = 3
	Ignore         = "nothing"
)
`

func TestConvertHappy8(t *testing.T) {
	RegisterTestingT(t)
	util.Dbg = testing.Verbose()
	s := bytes.NewBufferString(enum8)
	m, err := Convert(s, "filename.go", transform.Upper,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err).Should(BeNil())
	Ω(m).Should(Equal(model.Model{
		Config: model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		},
		LcType:       "sweet",
		BaseType:     "int",
		Version:      util.Version,
		Values:       []string{"Mars", "Bounty", "Snickers", "Kitkat"},
		DefaultValue: "Mars",
		Case:         transform.Upper,
		S1:           "",
		S2:           "",
		TagTable:     "",
	}))
}

const enumE1 = `
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
	s := bytes.NewBufferString(enumE1)
	_, err := Convert(s, "filename.go", transform.Stet,
		model.Config{
			MainType: "Sweet",
			Plural:   "Sweets",
			Pkg:      "confectionary",
		})
	Ω(err.Error()).Should(Equal("Failed to find Sweet in filename.go"))
}
