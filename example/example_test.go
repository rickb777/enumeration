package example

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v2/enum"
	"testing"
)

func TestString(t *testing.T) {
	g := gomega.NewWithT(t)
	g.Expect(Sunday.String()).Should(gomega.Equal("Sunday"))
	g.Expect(Monday.String()).Should(gomega.Equal("Monday"))
}

func TestOrdinal(t *testing.T) {
	g := gomega.NewWithT(t)
	g.Expect(int(Sunday)).Should(gomega.Equal(1))
	g.Expect(Sunday.Ordinal()).Should(gomega.Equal(0))
	g.Expect(int(Monday)).Should(gomega.Equal(2))
	g.Expect(Monday.Ordinal()).Should(gomega.Equal(1))
	g.Expect(Friday.Ordinal()).Should(gomega.Equal(5))
	g.Expect(DayOf(3)).Should(gomega.Equal(Wednesday))
	g.Expect(numberOfDays).Should(gomega.Equal(7))
}

func TestValue(t *testing.T) {
	g := gomega.NewWithT(t)
	g.Expect(G.Float()).Should(gomega.Equal(347.20001220703125))
	g.Expect(Sunday.Int()).Should(gomega.Equal(1))
	g.Expect(Wednesday.Int()).Should(gomega.Equal(4))
	g.Expect(POST.Int()).Should(gomega.Equal(3))
	g.Expect(November.Int()).Should(gomega.Equal(11))
	g.Expect(Koala_Bear.Int()).Should(gomega.Equal(4))
}

func TestAllDays(t *testing.T) {
	g := gomega.NewWithT(t)
	g.Expect(AllDays[0]).Should(gomega.Equal(Sunday))
	g.Expect(AllDays[5]).Should(gomega.Equal(Friday))
}

func TestAsDay(t *testing.T) {
	g := gomega.NewWithT(t)
	v, err := AsDay("Tuesday")
	g.Expect(v, err).Should(gomega.Equal(Tuesday))
	_, err = AsDay("Nosuchday")
	g.Expect(err).Should(gomega.HaveOccurred())
}

func TestAsMethod(t *testing.T) {
	g := gomega.NewWithT(t)
	methodMarshalTextRep = enum.Identifier
	g.Expect(AsMethod("POST")).Should(gomega.Equal(POST))
	g.Expect(AsMethod("PO")).Should(gomega.Equal(POST))
	g.Expect(AsMethod("3")).Should(gomega.Equal(POST))

	g.Expect(AsMethod("PUT")).Should(gomega.Equal(PUT))
	g.Expect(AsMethod("PU")).Should(gomega.Equal(PUT))
	g.Expect(AsMethod("2")).Should(gomega.Equal(PUT))
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
	g := gomega.NewWithT(t)

	baseMarshalTextRep = enum.Number
	dayMarshalTextRep = enum.Number
	methodMarshalTextRep = enum.Number
	monthMarshalTextRep = enum.Number
	petMarshalTextRep = enum.Number

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":347.2,"D":3,"X":3,"M":11,"P":4}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>347.2</B><D>3</D><X>3</X><M>11</M><P>4</P></Group>`), string(x))
}

func TestMarshalUsingOrdinal(t *testing.T) {
	g := gomega.NewWithT(t)

	baseMarshalTextRep = enum.Ordinal
	dayMarshalTextRep = enum.Ordinal
	methodMarshalTextRep = enum.Ordinal
	monthMarshalTextRep = enum.Ordinal
	petMarshalTextRep = enum.Ordinal

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":2,"D":2,"X":3,"M":10,"P":4}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>2</B><D>2</D><X>3</X><M>10</M><P>4</P></Group>`), string(x))
}

func TestMarshalUsingIdentifier(t *testing.T) {
	g := gomega.NewWithT(t)

	baseMarshalTextRep = enum.Identifier
	dayMarshalTextRep = enum.Identifier
	methodMarshalTextRep = enum.Identifier
	monthMarshalTextRep = enum.Identifier
	petMarshalTextRep = enum.Identifier

	g.Expect(G.MarshalJSON()).Should(gomega.Equal([]byte{'"', 'g', '"'}))

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":"g","D":"Tuesday","X":"POST","M":"November","P":"koala bear"}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>g</B><D>Tuesday</D><X>POST</X><M>November</M><P>koala bear</P></Group>`), string(x))
}

func TestMarshalUsingTag(t *testing.T) {
	g := gomega.NewWithT(t)

	baseMarshalTextRep = enum.Tag
	dayMarshalTextRep = enum.Tag
	methodMarshalTextRep = enum.Tag
	monthMarshalTextRep = enum.Tag
	petMarshalTextRep = enum.Tag

	g.Expect(G.MarshalJSON()).Should(gomega.Equal([]byte{'"', 'g', '"'}))

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"Phascolarctos Cinereus"}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>g</B><D>Tuesday</D><X>PO</X><M>November</M><P>Phascolarctos Cinereus</P></Group>`), string(x))
}

func TestUnmarshalJSON1(t *testing.T) {
	g := gomega.NewWithT(t)
	cases := []struct {
		input string
		rep   enum.Representation
	}{
		{input: `{"B":347.2,"D":3,"X":3,"M":11,"P":4}`, rep: enum.Identifier},
		{input: `{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"Koala Bear"}`, rep: enum.Identifier},

		{input: `{"B":347.2,"D":3,"X":3,"M":11,"P":4}`, rep: enum.Tag},
		{input: `{"B":"g","D":"Tuesday","X":"PO","M":"November","P":"Koala Bear"}`, rep: enum.Tag},

		{input: `{"B":347.2,"D":3,"X":3,"M":11,"P":4}`, rep: enum.Number},
		{input: `{"B":"G","D":"Tuesday","X":"PO","M":"november","P":"Koala Bear"}`, rep: enum.Number},

		{input: `{"B":2,"D":2,"X":3,"M":10,"P":4}`, rep: enum.Ordinal},
		{input: `{"B":"G","D":"Tuesday","X":"PO","M":"NOVEMBER","P":"Koala Bear"}`, rep: enum.Ordinal},
	}
	for _, c := range cases {
		baseMarshalTextRep = c.rep
		dayMarshalTextRep = c.rep
		methodMarshalTextRep = c.rep
		monthMarshalTextRep = c.rep
		petMarshalTextRep = c.rep
		var v Group
		err := json.Unmarshal([]byte(c.input), &v)
		g.Expect(err).NotTo(gomega.HaveOccurred(), "%s %d", c.input, c.rep)
		g.Expect(v).Should(gomega.Equal(Group{G, Tuesday, POST, November, Koala_Bear}), "%s %d", c.input, c.rep)
	}
}

func TestMethodScan(t *testing.T) {
	g := gomega.NewWithT(t)
	methodStoreRep = enum.Ordinal
	cases := []interface{}{
		int64(3), int64(3), float64(3), "POST", "PO", []byte("POST"), []byte("PO"),
	}
	for i, s := range cases {
		if i > 0 {
			methodStoreRep = enum.Identifier
		}
		var m = new(Method)
		err := m.Scan(s)
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(*m).Should(gomega.Equal(POST))
	}
}

func TestMonthScan(t *testing.T) {
	g := gomega.NewWithT(t)
	monthStoreRep = enum.Ordinal
	cases := []interface{}{
		int64(10), int64(11), float64(11), "november", []byte("NOVEMBER"),
	}
	for i, s := range cases {
		if i > 0 {
			monthStoreRep = enum.Identifier
		}
		var m = new(Month)
		err := m.Scan(s)
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(*m).Should(gomega.Equal(November))
	}
}

func TestPetScan(t *testing.T) {
	g := gomega.NewWithT(t)
	petStoreRep = enum.Ordinal
	cases := []interface{}{
		int64(4), int64(4), float64(4), "Koala Bear", "koala bear", "koala_bear", []byte("Koala Bear"), "Phascolarctos Cinereus",
	}
	for i, s := range cases {
		if i > 0 {
			petStoreRep = enum.Identifier
		}
		var m = new(Pet)
		err := m.Scan(s)
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(*m).Should(gomega.Equal(Koala_Bear))
	}
}

func TestValueUsingNumber(t *testing.T) {
	g := gomega.NewWithT(t)

	baseStoreRep = enum.Number
	dayStoreRep = enum.Number
	methodStoreRep = enum.Number
	monthStoreRep = enum.Number
	petStoreRep = enum.Number

	g.Expect(G.Value()).To(gomega.Equal(float64(347.20001220703125)))
	g.Expect(Tuesday.Value()).To(gomega.Equal(int64(3)))
	g.Expect(POST.Value()).To(gomega.Equal(int64(3)))
	g.Expect(November.Value()).To(gomega.Equal(int64(11)))
	g.Expect(Koala_Bear.Value()).To(gomega.Equal(int64(4)))
}

func TestValueUsingOrdinal(t *testing.T) {
	g := gomega.NewWithT(t)

	baseStoreRep = enum.Ordinal
	dayStoreRep = enum.Ordinal
	methodStoreRep = enum.Ordinal
	monthStoreRep = enum.Ordinal
	petStoreRep = enum.Ordinal

	g.Expect(G.Value()).To(gomega.Equal(int64(2)))
	g.Expect(Tuesday.Value()).To(gomega.Equal(int64(2)))
	g.Expect(POST.Value()).To(gomega.Equal(int64(3)))
	g.Expect(November.Value()).To(gomega.Equal(int64(10)))
	g.Expect(Koala_Bear.Value()).To(gomega.Equal(int64(4)))
}

func TestValueUsingTag(t *testing.T) {
	g := gomega.NewWithT(t)

	baseStoreRep = enum.Tag
	dayStoreRep = enum.Tag
	methodStoreRep = enum.Tag
	monthStoreRep = enum.Tag
	petStoreRep = enum.Tag

	g.Expect(G.Value()).To(gomega.Equal("g"))
	g.Expect(Tuesday.Value()).To(gomega.Equal("Tuesday"))
	g.Expect(POST.Value()).To(gomega.Equal("PO"))
	g.Expect(November.Value()).To(gomega.Equal("November"))
	g.Expect(Koala_Bear.Value()).To(gomega.Equal("Phascolarctos Cinereus"))
}

func TestValueUsingIdentifier(t *testing.T) {
	g := gomega.NewWithT(t)

	baseStoreRep = enum.Identifier
	dayStoreRep = enum.Identifier
	methodStoreRep = enum.Identifier
	monthStoreRep = enum.Identifier
	petStoreRep = enum.Identifier

	g.Expect(G.Value()).To(gomega.Equal("g"))
	g.Expect(Tuesday.Value()).To(gomega.Equal("Tuesday"))
	g.Expect(POST.Value()).To(gomega.Equal("POST"))
	g.Expect(November.Value()).To(gomega.Equal("November"))
	g.Expect(Koala_Bear.Value()).To(gomega.Equal("koala bear"))
}

func TestGobEncodeAndDecode(t *testing.T) {
	g := gomega.NewWithT(t)
	v1 := Group{G, Tuesday, POST, November, Koala_Bear}
	gob.Register(v1)

	// gob-encode
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(v1)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	// gob-decode
	var v2 Group
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&v2)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(v2).Should(gomega.Equal(v1))
}
