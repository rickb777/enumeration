package codegen

import (
	"github.com/rickb777/enumeration/v4/internal/collection"
	"github.com/rickb777/expect"
	"testing"
)

func TestUnits(t *testing.T) {
	us1 := New()
	u1 := Unit{
		Declares: "aaa",
		Requires: []string{"bbb"},
		Extra:    map[string]any{"eee": 1},
		Imports:  collection.NewSet[string]("iii"),
		Template: "ttt",
	}
	us1.Add(u1)
	expect.Number(len(us1.m)).ToBe(t, 1)
	expect.Number(len(us1.l)).ToBe(t, 1)

	s := us1.Slice()
	expect.Number(len(s)).ToBe(t, 1)
	expect.Any(s).ToBe(t, []Unit{u1})

	q1, found := us1.Take("aaa")
	expect.Bool(found).ToBeTrue(t)
	expect.Any(q1).ToBe(t, u1)
	//Ω(len(us1.m)).ToBe(0,t)
	expect.Slice(us1.l).ToHaveLength(t, 0)
}
