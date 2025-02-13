package codegen

import (
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v4/internal/collection"
	"testing"
)

func TestUnits(t *testing.T) {
	RegisterTestingT(t)
	us1 := New()
	u1 := Unit{
		Declares: "aaa",
		Requires: []string{"bbb"},
		Extra:    map[string]any{"eee": 1},
		Imports:  collection.NewSet[string]("iii"),
		Template: "ttt",
	}
	us1.Add(u1)
	Ω(len(us1.m)).Should(Equal(1))
	Ω(len(us1.l)).Should(Equal(1))

	s := us1.Slice()
	Ω(s).Should(HaveLen(1))
	Ω(s).Should(ConsistOf(u1))

	q1, found := us1.Take("aaa")
	Ω(found).Should(BeTrue())
	Ω(q1).Should(Equal(u1))
	//Ω(len(us1.m)).Should(Equal(0))
	Ω(len(us1.l)).Should(Equal(0))
}
