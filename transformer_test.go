package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestTransformers_Format_non_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	xf := toUpper().Then(xUnsnake().Then(toLower()))
	r := xf.Format("in")
	g.Expect(r).To(Equal(`strings.ToLower(strings.ReplaceAll(strings.ToUpper(in), "_", " "))`))
}

func TestTransformers_Apply_non_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	xf := toUpper().Then(xUnsnake().Then(toLower()))
	r := xf.Apply("AbC_DeF")
	g.Expect(r).To(Equal(`abc def`))
}

func TestTransformers_Describe_non_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	xf := toUpper().Then(xUnsnake().Then(toLower()))
	r := xf.Describe()
	g.Expect(r).To(Equal("\n// The case of s does not matter.\n// All underscores are replaced with space.\n// The case of s does not matter."))
}

func TestTransformers_Format_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	var xf *Transformer
	r := xf.Format("in")
	g.Expect(r).To(Equal(`in`))
}

func TestTransformers_Apply_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	var xf *Transformer
	r := xf.Apply("in")
	g.Expect(r).To(Equal(`in`))
}

func TestTransformers_Describe_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	var xf *Transformer
	r := xf.Describe()
	g.Expect(r).To(Equal(""))
}
