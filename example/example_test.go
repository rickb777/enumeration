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
	g.Expect(MyDog.String()).Should(gomega.Equal("dog"))
	g.Expect(OnlineSales.String()).Should(gomega.Equal("online"))
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
	g.Expect(Θήτα.Int()).Should(gomega.Equal(8))
	g.Expect(POST.Int()).Should(gomega.Equal(3))
	g.Expect(November.Int()).Should(gomega.Equal(11))
	g.Expect(MyKoala_Bear.Int()).Should(gomega.Equal(4))
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
	G GreekAlphabet
	X Method
	M Month
	P Pet
	C SalesChannel
}

func TestMarshalUsingNumber(t *testing.T) {
	g := gomega.NewWithT(t)

	setMarshalReps(enum.Number)

	v := Group{G, Tuesday, Θήτα, POST, November, MyKoala_Bear, OnlineSales}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":347.2,"D":3,"G":8,"X":3,"M":11,"P":4,"C":"webshop"}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>347.2</B><D>3</D><G>8</G><X>3</X><M>11</M><P>4</P><C>online</C></Group>`), string(x))
}

func TestMarshalUsingOrdinal(t *testing.T) {
	g := gomega.NewWithT(t)

	setMarshalReps(enum.Ordinal)

	v := Group{G, Tuesday, Θήτα, POST, November, MyKoala_Bear, OnlineSales}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":2,"D":2,"G":7,"X":3,"M":10,"P":4,"C":"webshop"}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>2</B><D>2</D><G>7</G><X>3</X><M>10</M><P>4</P><C>online</C></Group>`), string(x))
}

func TestMarshalUsingIdentifier(t *testing.T) {
	g := gomega.NewWithT(t)

	setMarshalReps(enum.Identifier)

	g.Expect(G.MarshalJSON()).Should(gomega.Equal([]byte{'"', 'g', '"'}))

	v := Group{G, Tuesday, Θήτα, POST, November, MyKoala_Bear, OnlineSales}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":"g","D":"Tuesday","G":"Θήτα","X":"POST","M":"November","P":"koala bear","C":"webshop"}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>g</B><D>Tuesday</D><G>Θήτα</G><X>POST</X><M>November</M><P>koala bear</P><C>online</C></Group>`), string(x))
}

func TestMarshalUsingTag(t *testing.T) {
	g := gomega.NewWithT(t)

	setMarshalReps(enum.Tag)

	g.Expect(G.MarshalJSON()).Should(gomega.Equal([]byte{'"', 'g', '"'}))

	v := Group{G, Tuesday, Θήτα, POST, November, MyElephant, OnlineSales}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":"g","D":"Tuesday","G":"theta","X":"PO","M":"November","P":"Loxodonta Africana","C":"webshop"}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>g</B><D>Tuesday</D><G>theta</G><X>PO</X><M>November</M><P>Loxodonta Africana</P><C>online</C></Group>`), string(x))
}

func TestUnmarshalJSON1(t *testing.T) {
	g := gomega.NewWithT(t)
	cases := []struct {
		input string
		rep   enum.Representation
	}{
		{input: `{"B":347.2,"D":3,"G":8,"X":3,"M":11,"P":4,"C":1}`, rep: enum.Identifier},
		{input: `{"B":"g","D":"Tuesday","G":"Θήτα","X":"post","M":"November","P":"Koala Bear","C":"webshop"}`, rep: enum.Identifier},

		{input: `{"B":347.2,"D":3,"G":8,"X":3,"M":11,"P":4,"C":1}`, rep: enum.Tag},
		{input: `{"B":"g","D":"Tuesday","G":"theta","X":"PO","M":"November","P":"Koala Bear","C":"webshop"}`, rep: enum.Tag},

		{input: `{"B":347.2,"D":3,"G":8,"X":3,"M":11,"P":4,"C":1}`, rep: enum.Number},
		{input: `{"B":"G","D":"Tuesday","G":"theta","X":"po","M":"november","P":"Koala Bear","C":"webshop"}`, rep: enum.Number},

		{input: `{"B":2,"D":2,"G":7,"X":3,"M":10,"P":4,"C":1}`, rep: enum.Ordinal},
		{input: `{"B":"G","D":"Tuesday","G":"Θήτα","X":"POST","M":"NOVEMBER","P":"Koala Bear","C":"webshop"}`, rep: enum.Ordinal},
		{input: `{"B":"G","D":"Tuesday","G":"Θήτα","X":"POST","M":"NOVEMBER","P":"Cuddly","C":"webshop"}`, rep: enum.Identifier},
	}
	for i, c := range cases {
		setMarshalReps(c.rep)

		var v Group
		err := json.Unmarshal([]byte(c.input), &v)
		g.Expect(err).NotTo(gomega.HaveOccurred(), "%d %s %d", i, c.input, c.rep)
		g.Expect(v).Should(gomega.Equal(Group{G, Tuesday, Θήτα, POST, November, MyKoala_Bear, OnlineSales}), "%d %s %d", i, c.input, c.rep)
	}
}

func TestMethodScan(t *testing.T) {
	g := gomega.NewWithT(t)
	methodStoreRep = enum.Ordinal
	cases := []interface{}{
		int64(3), int64(3), float64(3), "POST", "PO", "post", "po", []byte("POST"), []byte("PO"),
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
		g.Expect(*m).Should(gomega.Equal(MyKoala_Bear))
	}
}

func TestSalesChannelScan(t *testing.T) {
	g := gomega.NewWithT(t)
	saleschannelStoreRep = enum.Ordinal
	expected := TelephoneSales // ordinal = 2
	cases := []interface{}{
		int64(2), int64(2), float64(2), "s", []byte("s"),
	}
	for i, s := range cases {
		if i > 0 {
			saleschannelStoreRep = enum.Identifier
			expected = InstoreSales // value = 2
		}
		var m = new(SalesChannel)
		err := m.Scan(s)
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(*m).Should(gomega.Equal(expected))
	}
}

func TestValueUsingNumber(t *testing.T) {
	g := gomega.NewWithT(t)

	setStoreRep(enum.Number)

	g.Expect(G.Value()).To(gomega.Equal(float64(347.20001220703125)))
	g.Expect(Tuesday.Value()).To(gomega.Equal(int64(3)))
	g.Expect(Θήτα.Value()).To(gomega.Equal(int64(8)))
	g.Expect(POST.Value()).To(gomega.Equal(int64(3)))
	g.Expect(November.Value()).To(gomega.Equal(int64(11)))
	g.Expect(MyKoala_Bear.Value()).To(gomega.Equal(int64(4)))
}

func TestValueUsingOrdinal(t *testing.T) {
	g := gomega.NewWithT(t)

	setStoreRep(enum.Ordinal)

	g.Expect(G.Value()).To(gomega.Equal(int64(2)))
	g.Expect(Tuesday.Value()).To(gomega.Equal(int64(2)))
	g.Expect(Θήτα.Value()).To(gomega.Equal(int64(7)))
	g.Expect(POST.Value()).To(gomega.Equal(int64(3)))
	g.Expect(November.Value()).To(gomega.Equal(int64(10)))
	g.Expect(MyKoala_Bear.Value()).To(gomega.Equal(int64(4)))
}

func TestValueUsingTag(t *testing.T) {
	g := gomega.NewWithT(t)

	setStoreRep(enum.Tag)

	g.Expect(G.Value()).To(gomega.Equal("g"))
	g.Expect(Tuesday.Value()).To(gomega.Equal("Tuesday"))
	g.Expect(POST.Value()).To(gomega.Equal("PO"))
	g.Expect(November.Value()).To(gomega.Equal("November"))
	g.Expect(MyKoala_Bear.Value()).To(gomega.Equal("Phascolarctos Cinereus"))
}

func TestValueUsingIdentifier(t *testing.T) {
	g := gomega.NewWithT(t)

	setStoreRep(enum.Identifier)

	g.Expect(G.Value()).To(gomega.Equal("g"))
	g.Expect(Tuesday.Value()).To(gomega.Equal("Tuesday"))
	g.Expect(POST.Value()).To(gomega.Equal("POST"))
	g.Expect(November.Value()).To(gomega.Equal("November"))
	g.Expect(MyKoala_Bear.Value()).To(gomega.Equal("koala bear"))
}

func TestGobEncodeAndDecode(t *testing.T) {
	g := gomega.NewWithT(t)
	v1 := Group{B: G, D: Tuesday, G: Θήτα, X: POST, M: November, P: MyKoala_Bear}
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

func setMarshalReps(rep enum.Representation) {
	baseMarshalTextRep = rep
	dayMarshalTextRep = rep
	greekalphabetMarshalTextRep = rep
	methodMarshalTextRep = rep
	monthMarshalTextRep = rep
	petMarshalTextRep = rep
}

func setStoreRep(rep enum.Representation) {
	baseStoreRep = rep
	dayStoreRep = rep
	greekalphabetStoreRep = rep
	methodStoreRep = rep
	monthStoreRep = rep
	petStoreRep = rep
}
