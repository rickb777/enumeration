package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestTransformers_Converter_non_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	xf := Transformers{toUpper, xUnsnake, toLower}
	fn := xf.TransformFunc()
	r := fn("in")
	g.Expect(r).To(Equal(`strings.ToLower(strings.ReplaceAll(strings.ToUpper(in), "_", " "))`))
}

func TestTransformers_Converter_empty(t *testing.T) {
	g := NewGomegaWithT(t)
	xf := Transformers{}
	fn := xf.TransformFunc()
	r := fn("in")
	g.Expect(r).To(Equal(`in`))
}
