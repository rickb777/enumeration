package collection

import (
	"testing"

	"github.com/rickb777/expect"
)

func TestSetUnits(t *testing.T) {
	s1 := NewSet[string]("a", "b", "c")
	expect.Number(len(s1)).ToBe(t, 3)
	expect.Bool(s1.Contains("b")).ToBeTrue(t)
	expect.Bool(s1.Contains("x")).ToBeFalse(t)

	s1.AddAll("a", "x", "y")
	expect.Number(len(s1)).ToBe(t, 5)
	expect.Bool(s1.Contains("x")).ToBeTrue(t)

	s1.Add("z")
	expect.Number(len(s1)).ToBe(t, 6)
	expect.Bool(s1.Contains("z")).ToBeTrue(t)

	s1.Union(NewSet("a", "1", "b", "2"))
	expect.Number(len(s1)).ToBe(t, 8)
	expect.Bool(s1.Contains("2")).ToBeTrue(t)

	ss := s1.Sorted()
	expect.Slice(ss).ToBe(t, "1", "2", "a", "b", "c", "x", "y", "z")
}
