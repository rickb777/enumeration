package collection

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestSetUnits(t *testing.T) {
	RegisterTestingT(t)

	s1 := NewSet[string]("a", "b", "c")
	Ω(len(s1)).Should(Equal(3))
	Ω(s1.Contains("b")).Should(BeTrue())
	Ω(s1.Contains("x")).Should(BeFalse())

	s1.AddAll("a", "x", "y")
	Ω(len(s1)).Should(Equal(5))
	Ω(s1.Contains("x")).Should(BeTrue())

	s1.Add("z")
	Ω(len(s1)).Should(Equal(6))
	Ω(s1.Contains("z")).Should(BeTrue())

	s1.Union(NewSet("a", "1", "b", "2"))
	Ω(len(s1)).Should(Equal(8))
	Ω(s1.Contains("2")).Should(BeTrue())

	ss := s1.Sorted()
	Ω(ss).Should(ConsistOf("1", "2", "a", "b", "c", "x", "y", "z"))
}
