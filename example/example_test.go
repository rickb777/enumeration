package example

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"github.com/rickb777/expect"
	"testing"
)

func TestString(t *testing.T) {
	expect.String(Sunday.String()).ToBe(t, "Sunday")
	expect.String(Monday.String()).ToBe(t, "Monday")
	expect.String(MyDog.String()).ToBe(t, "dog")
	expect.String(OnlineSales.String()).ToBe(t, "online")
}

func TestOrdinal(t *testing.T) {
	expect.Number(int(Sunday)).ToBe(t, 1)
	expect.Number(Sunday.Ordinal()).ToBe(t, 0)
	expect.Number(int(Monday)).ToBe(t, 2)
	expect.Number(Monday.Ordinal()).ToBe(t, 1)
	expect.Number(Friday.Ordinal()).ToBe(t, 5)
	expect.Number(DayOf(0)).ToBe(t, Sunday)
	expect.Number(DayOf(3)).ToBe(t, Wednesday)
	expect.Number(DayOf(6)).ToBe(t, Saturday)
	expect.Number(DayOf(7)).ToBe(t, 0)
	expect.Number(DayOf(13)).ToBe(t, 0)
	expect.Bool(Wednesday.IsValid()).ToBeTrue(t)
	expect.Bool(Day(10).IsValid()).ToBeFalse(t)
	expect.Bool(DayOf(10).IsValid()).ToBeFalse(t)
	expect.Number(numberOfDays).ToBe(t, 7)
}

func TestIntOrFloat(t *testing.T) {
	expect.Number(G.Float()).ToBe(t, 347.20001220703125)
	expect.Number(Sunday.Int()).ToBe(t, 1)
	expect.Number(Wednesday.Int()).ToBe(t, 4)
	expect.Number(Θήτα.Int()).ToBe(t, 8)
	expect.Number(POST.Int()).ToBe(t, 3)
	expect.Number(November.Int()).ToBe(t, 11)
	expect.Number(MyKoala_Bear.Int()).ToBe(t, 4)
}

func TestAllDays(t *testing.T) {
	expect.Any(AllDays[0]).ToBe(t, Sunday)
	expect.Any(AllDays[5]).ToBe(t, Friday)
}

func TestAsDay(t *testing.T) {
	v, err := AsDay("Tuesday")
	expect.Any(v, err).ToBe(t, Tuesday)
	_, err = AsDay("Nosuchday")
	expect.Error(err).ToHaveOccurred(t)
}

func TestAsMethod(t *testing.T) {
	//methodMarshalTextRep = enum.Identifier
	expect.Number(AsMethod("POST")).ToBe(t, POST)
	//expect.Number(AsMethod("PO")).ToBe(POST))
	expect.Number(AsMethod("3")).ToBe(t, POST)

	expect.Number(AsMethod("PUT")).ToBe(t, PUT)
	//expect.Number( AsMethod("PU")).ToBe(PUT))
	expect.Number(AsMethod("2")).ToBe(t, PUT)
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
	//setMarshalReps(enum.Number)
	v := Group{G, Tuesday, Θήτα, POST, November, MyKoala_Bear, OnlineSales}
	s, err := json.Marshal(v)
	expect.Error(err).Not().ToHaveOccurred(t)
	x, err := xml.Marshal(v)
	expect.Error(err).Not().ToHaveOccurred(t)
	expect.String(s).ToEqual(t, `{"B":347.2,"D":3,"G":"theta","X":"PO","M":"November","P":"Phascolarctos Cinereus","C":"webshop"}`)
	expect.String(x).ToEqual(t, `<Group><B>347.2</B><D>3</D><G>theta</G><X>3</X><M>November</M><P>Phascolarctos Cinereus</P><C>3</C></Group>`)
}

func TestMethodScan(t *testing.T) {
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
		expect.Error(err).Info(i).Not().ToHaveOccurred(t)
		expect.Number(*m).Info(i).ToBe(t, POST)
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
//		expect.String(t, err).NotTo(gomega.HaveOccurred())
//		expect.String(t, *m).ToBe(November))
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
//		expect.String(t, err).NotTo(gomega.HaveOccurred())
//		expect.String(t, *m).ToBe(MyKoala_Bear))
//	}
//}

func TestSalesChannelScan(t *testing.T) {
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
		expect.Error(err).Info("%#v", c.in).ToBeNil(t)
		expect.Any(*m).Info("%#v", c.in).ToBe(t, c.expected)
	}
}

func TestValue(t *testing.T) {
	expect.Any(TelephoneSales.Value()).ToEqual(t, "t")
	expect.Any(Egypt.Value()).ToEqual(t, "eg")
	expect.Any(Ζήτα.Value()).ToEqual(t, "\u0396")
	expect.Any(POST.Value()).ToEqual(t, int64(3))
}

func TestGobEncodeAndDecode(t *testing.T) {
	v1 := Group{B: G, D: Tuesday, G: Θήτα, X: POST, M: November, P: MyKoala_Bear}
	gob.Register(v1)

	// gob-encode
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(v1)
	expect.Error(err).Not().ToHaveOccurred(t)

	// gob-decode
	var v2 Group
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&v2)
	expect.Error(err).Not().ToHaveOccurred(t)
	expect.Any(v2).ToBe(t, v1)
}
