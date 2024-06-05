package example

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/onsi/gomega"
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

func TestIntOrFloat(t *testing.T) {
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
	//methodMarshalTextRep = enum.Identifier
	g.Expect(AsMethod("POST")).Should(gomega.Equal(POST))
	//g.Expect(AsMethod("PO")).Should(gomega.Equal(POST))
	g.Expect(AsMethod("3")).Should(gomega.Equal(POST))

	g.Expect(AsMethod("PUT")).Should(gomega.Equal(PUT))
	//g.Expect(AsMethod("PU")).Should(gomega.Equal(PUT))
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

func TestMarshal(t *testing.T) {
	g := gomega.NewWithT(t)

	//setMarshalReps(enum.Number)

	v := Group{G, Tuesday, Θήτα, POST, November, MyKoala_Bear, OnlineSales}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(string(s)).Should(gomega.Equal(`{"B":347.2,"D":3,"G":"theta","X":"PO","M":"November","P":"Phascolarctos Cinereus","C":"webshop"}`))
	g.Expect(string(x)).Should(gomega.Equal(`<Group><B>347.2</B><D>3</D><G>theta</G><X>3</X><M>November</M><P>Phascolarctos Cinereus</P><C>3</C></Group>`), string(x))
}

func TestMethodScan(t *testing.T) {
	g := gomega.NewWithT(t)

	//methodStoreRep = enum.Ordinal
	cases := []interface{}{
		int64(3), int64(3), float64(3), "POST", "post", []byte("POST"),
	}
	for i, s := range cases {
		if i > 0 {
			//methodStoreRep = enum.Identifier
		}
		var m = new(Method)
		err := m.Scan(s)
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(*m).Should(gomega.Equal(POST))
	}
}

//func TestMonthScan(t *testing.T) {
//	g := gomega.NewWithT(t)
//
//	//monthStoreRep = enum.Ordinal
//	cases := []interface{}{
//		int64(10), int64(11), float64(11), "november", []byte("NOVEMBER"),
//	}
//	for i, s := range cases {
//		if i > 0 {
//			//monthStoreRep = enum.Identifier
//		}
//		var m = new(Month)
//		err := m.Scan(s)
//		g.Expect(err).NotTo(gomega.HaveOccurred())
//		g.Expect(*m).Should(gomega.Equal(November))
//	}
//}

//func TestPetScan(t *testing.T) {
//	g := gomega.NewWithT(t)
//
//	//petStoreRep = enum.Ordinal
//	cases := []interface{}{
//		int64(4), int64(4), float64(4), "Koala Bear", "koala bear", "koala_bear", []byte("Koala Bear"), "Phascolarctos Cinereus",
//	}
//	for i, s := range cases {
//		if i > 0 {
//			//petStoreRep = enum.Identifier
//		}
//		var m = new(Pet)
//		err := m.Scan(s)
//		g.Expect(err).NotTo(gomega.HaveOccurred())
//		g.Expect(*m).Should(gomega.Equal(MyKoala_Bear))
//	}
//}

func TestSalesChannelScan(t *testing.T) {
	g := gomega.NewWithT(t)

	cases := []struct {
		in       interface{}
		expected SalesChannel
	}{
		{in: int64(7), expected: TelephoneSales},
		{in: float64(7), expected: TelephoneSales},
		{in: "5", expected: InstoreSales},
		{in: "s", expected: InstoreSales},
		{in: []byte("o"), expected: OnlineSales},
		{in: nil, expected: 0},
	}
	for _, c := range cases {
		var m = new(SalesChannel)
		err := m.Scan(c.in)
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(*m).Should(gomega.Equal(c.expected), "%#v", c.in)
	}
}

func TestValue(t *testing.T) {
	g := gomega.NewWithT(t)

	g.Expect(TelephoneSales.Value()).To(gomega.Equal("t"))
	g.Expect(Egypt.Value()).To(gomega.Equal("eg"))
	g.Expect(Ζήτα.Value()).To(gomega.Equal("\u0396"))
	g.Expect(POST.Value()).To(gomega.Equal(int64(3)))
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
