package example

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	. "github.com/onsi/gomega"
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

func TestDayMarshalText(t *testing.T) {
	g := NewGomegaWithT(t)
	tt, err := Friday.MarshalText()
	g.Expect(tt, err).Should(Equal([]byte("Friday")))
}

func TestDayUnmarshalText(t *testing.T) {
	g := NewGomegaWithT(t)
	var d = new(Day)
	err := d.UnmarshalText([]byte("Friday"))
	g.Expect(*d, err).Should(Equal(Friday))
	err = d.UnmarshalText([]byte("5"))
	g.Expect(*d, err).Should(Equal(Friday))
}

func TestMethodMarshalText(t *testing.T) {
	g := NewGomegaWithT(t)
	methodMarshalTextUsingLiteral = true
	tt, err := POST.MarshalText()
	g.Expect(tt, err).Should(Equal([]byte("POST")))

	methodMarshalTextUsingLiteral = false
	tt, err = POST.MarshalText()
	g.Expect(tt, err).Should(Equal([]byte("PO")))
}

func TestMethodUnmarshalText(t *testing.T) {
	g := NewGomegaWithT(t)
	var m = new(Method)
	methodMarshalTextUsingLiteral = false
	err := m.UnmarshalText([]byte("POST"))
	g.Expect(*m, err).Should(Equal(POST))

	methodMarshalTextUsingLiteral = true
	err = m.UnmarshalText([]byte("POST"))
	g.Expect(*m, err).Should(Equal(POST))
}

func TestAsMethod(t *testing.T) {
	g := NewGomegaWithT(t)
	methodMarshalTextUsingLiteral = false
	g.Expect(AsMethod("POST")).Should(Equal(POST))
	g.Expect(AsMethod("PO")).Should(Equal(POST))
	g.Expect(AsMethod("3")).Should(Equal(POST))

	methodMarshalTextUsingLiteral = true
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

func TestMarshalJSONUsingOrdinal(t *testing.T) {
	g := NewGomegaWithT(t)

	BaseMarshalJSONUsingString = false
	DayMarshalJSONUsingString = false
	MethodMarshalJSONUsingString = false
	MonthMarshalJSONUsingString = false
	PetMarshalJSONUsingString = false

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":2,"D":2,"X":3,"M":10,"P":4}`))
}

func TestMarshalJSONUsingString(t *testing.T) {
	g := NewGomegaWithT(t)

	BaseMarshalJSONUsingString = true
	DayMarshalJSONUsingString = true
	MethodMarshalJSONUsingString = true
	methodMarshalTextUsingLiteral = true
	MonthMarshalJSONUsingString = true
	PetMarshalJSONUsingString = true

	g.Expect(G.MarshalJSON()).Should(Equal([]byte{'"', 'g', '"'}))

	v := Group{G, Tuesday, POST, November, Koala_Bear}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":"g","D":"Tuesday","X":"POST","M":"November","P":"koala bear"}`))
}

func TestUnmarshalJSON1(t *testing.T) {
	g := NewGomegaWithT(t)
	methodMarshalTextUsingLiteral = false
	cases := []string{
		`{"B":2,"D":2,"X":3,"M":10,"P":4}`,
		`{"B":2,"D":"Tuesday","X":"PO","M":"November","P":"Koala Bear"}`,
	}
	for _, s := range cases {
		var v Group
		err := json.Unmarshal([]byte(s), &v)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(v).Should(Equal(Group{G, Tuesday, POST, November, Koala_Bear}))
	}
}

func TestMethodScan(t *testing.T) {
	g := NewGomegaWithT(t)
	methodMarshalTextUsingLiteral = false
	cases := []interface{}{
		int64(4), "POST", "PO", []byte("POST"), []byte("PO"),
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
