package example

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	. "github.com/onsi/gomega"
	"github.com/rickb777/enumeration/enum"
	"testing"
)

func TestString(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Sunday.String()).Should(Equal("Sunday"))
	g.Expect(Monday.String()).Should(Equal("Monday"))
}

func TestOrdinal(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(int(Sunday)).Should(Equal(1))
	g.Expect(Sunday.Ordinal()).Should(Equal(0))
	g.Expect(int(Monday)).Should(Equal(2))
	g.Expect(Monday.Ordinal()).Should(Equal(1))
	g.Expect(Friday.Ordinal()).Should(Equal(5))
	g.Expect(DayOf(3)).Should(Equal(Wednesday))
	g.Expect(numberOfDays).Should(Equal(7))
}

func TestValue(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Sunday.Int()).Should(Equal(1))
	g.Expect(Wednesday.Int()).Should(Equal(4))
}

func TestAllDays(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(AllDays[0]).Should(Equal(Sunday))
	g.Expect(AllDays[5]).Should(Equal(Friday))
}

func TestAsDay(t *testing.T) {
	g := NewGomegaWithT(t)
	v, err := AsDay("Tuesday")
	g.Expect(v, err).Should(Equal(Tuesday))
	_, err = AsDay("Nosuchday")
	g.Expect(err).Should(HaveOccurred())
}

func TestAsMethod(t *testing.T) {
	g := NewGomegaWithT(t)
	methodMarshalTextRep = enum.Identifier
	g.Expect(AsMethod("POST")).Should(Equal(POST))
	g.Expect(AsMethod("PO")).Should(Equal(POST))
	g.Expect(AsMethod("3")).Should(Equal(POST))

	g.Expect(AsMethod("PUT")).Should(Equal(PUT))
	g.Expect(AsMethod("PU")).Should(Equal(PUT))
	g.Expect(AsMethod("2")).Should(Equal(PUT))
}

//-------------------------------------------------------------------------------------------------

type Group struct {
	B Base
	D Day
	X Method
	M Month
	P Pet
}

func TestMarshalUsingNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	baseMarshalTextRep = enum.Number
	dayMarshalTextRep = enum.Number
	methodMarshalTextRep = enum.Number
	monthMarshalTextRep = enum.Number
	petMarshalTextRep = enum.Number

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":347.2,"D":3,"X":3,"M":11,"P":4}`))
	g.Expect(string(x)).Should(Equal(`<Group><B>347.2</B><D>3</D><X>3</X><M>11</M><P>4</P></Group>`), string(x))
}

func TestMarshalUsingOrdinal(t *testing.T) {
	g := NewGomegaWithT(t)

	baseMarshalTextRep = enum.Ordinal
	dayMarshalTextRep = enum.Ordinal
	methodMarshalTextRep = enum.Ordinal
	monthMarshalTextRep = enum.Ordinal
	petMarshalTextRep = enum.Ordinal

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":2,"D":2,"X":3,"M":10,"P":4}`))
	g.Expect(string(x)).Should(Equal(`<Group><B>2</B><D>2</D><X>3</X><M>10</M><P>4</P></Group>`), string(x))
}

func TestMarshalUsingIdentifier(t *testing.T) {
	g := NewGomegaWithT(t)

	baseMarshalTextRep = enum.Identifier
	dayMarshalTextRep = enum.Identifier
	methodMarshalTextRep = enum.Identifier
	monthMarshalTextRep = enum.Identifier
	petMarshalTextRep = enum.Identifier

	g.Expect(G.MarshalJSON()).Should(Equal([]byte{'"', 'g', '"'}))

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":"g","D":"Tuesday","X":"POST","M":"November","P":"koala bear"}`))
	g.Expect(string(x)).Should(Equal(`<Group><B>g</B><D>Tuesday</D><X>POST</X><M>November</M><P>koala bear</P></Group>`), string(x))
}

func TestMarshalUsingTag(t *testing.T) {
	g := NewGomegaWithT(t)

	baseMarshalTextRep = enum.Tag
	dayMarshalTextRep = enum.Tag
	methodMarshalTextRep = enum.Tag
	monthMarshalTextRep = enum.Tag
	petMarshalTextRep = enum.Tag

	g.Expect(G.MarshalJSON()).Should(Equal([]byte{'"', 'g', '"'}))

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"koala bear"}`))
	g.Expect(string(x)).Should(Equal(`<Group><B>g</B><D>Tuesday</D><X>PO</X><M>November</M><P>koala bear</P></Group>`), string(x))
}

func TestUnmarshalJSON1(t *testing.T) {
	g := NewGomegaWithT(t)
	methodMarshalTextRep = enum.Identifier
	cases := []struct {
		input string
		rep   enum.Representation
	}{
		{input: `{"B":347.2,"D":3,"X":3,"M":11,"P":4}`, rep: enum.Identifier},
		{input: `{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"Koala Bear"}`, rep: enum.Identifier},

		{input: `{"B":347.2,"D":3,"X":3,"M":11,"P":4}`, rep: enum.Tag},
		{input: `{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"Koala Bear"}`, rep: enum.Tag},

		{input: `{"B":347.2,"D":3,"X":3,"M":11,"P":4}`, rep: enum.Number},
		{input: `{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"Koala Bear"}`, rep: enum.Number},

		{input: `{"B":2,"D":2,"X":3,"M":10,"P":4}`, rep: enum.Ordinal},
		{input: `{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"Koala Bear"}`, rep: enum.Ordinal},
	}
	for _, c := range cases {
		baseMarshalTextRep = c.rep
		dayMarshalTextRep = c.rep
		methodMarshalTextRep = c.rep
		monthMarshalTextRep = c.rep
		petMarshalTextRep = c.rep
		var v Group
		err := json.Unmarshal([]byte(c.input), &v)
		g.Expect(err).NotTo(HaveOccurred(), "%s %d", c.input, c.rep)
		g.Expect(v).Should(Equal(Group{G, Tuesday, POST, November, Koala_Bear}), "%s %d", c.input, c.rep)
	}
}

func TestMethodScan(t *testing.T) {
	g := NewGomegaWithT(t)
	methodMarshalTextRep = enum.Identifier
	cases := []interface{}{
		int64(3), "POST", "PO", []byte("POST"), []byte("PO"),
	}
	for _, s := range cases {
		var m = new(Method)
		err := m.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*m).Should(Equal(POST))
	}
}

func TestGobEncodeAndDecode(t *testing.T) {
	g := NewGomegaWithT(t)
	v1 := Group{G, Tuesday, POST, November, Koala_Bear}
	gob.Register(v1)

	// gob-encode
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(v1)
	g.Expect(err).NotTo(HaveOccurred())

	// gob-decode
	var v2 Group
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&v2)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(v2).Should(Equal(v1))
}
