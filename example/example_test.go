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
	g.Expect(Sunday.Ordinal(), 0)
	g.Expect(int(Monday)).Should(Equal(2))
	g.Expect(Monday.Ordinal(), 1)
	g.Expect(numberOfDays).Should(Equal(7))
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
	g.Expect(err).ShouldNot(BeNil())
}

func TestMarshalText(t *testing.T) {
	g := NewGomegaWithT(t)
	tt, err := Monday.MarshalText()
	g.Expect(tt, err).Should(Equal([]byte("Monday")))
}

func TestUnmarshalText(t *testing.T) {
	g := NewGomegaWithT(t)
	var d = new(Day)
	err := d.UnmarshalText([]byte("Monday"))
	g.Expect(*d, err).Should(Equal(Monday))
}

//-------------------------------------------------------------------------------------------------

type Three struct {
	B Base
	D Day
	M Month
}

func TestMarshalJSON1(t *testing.T) {
	g := NewGomegaWithT(t)

	BaseMarshalJSONUsingString = false
	DayMarshalJSONUsingString = false
	MonthMarshalJSONUsingString = false

	v := Three{G, Tuesday, November}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":2,"D":2,"M":10}`))
}

func TestMarshalJSON2(t *testing.T) {
	g := NewGomegaWithT(t)

	BaseMarshalJSONUsingString = true
	DayMarshalJSONUsingString = true
	MonthMarshalJSONUsingString = true

	g.Expect(G.MarshalJSON()).Should(Equal([]byte{'"', 'G', '"'}))

	v := Three{G, Tuesday, November}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"B":"G","D":"Tuesday","M":"November"}`))
}

func TestUnmarshalJSON1(t *testing.T) {
	g := NewGomegaWithT(t)
	s := `{"B":2,"D":2,"M":10}`
	var v Three
	err := json.Unmarshal([]byte(s), &v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(v).Should(Equal(Three{G, Tuesday, November}))
}

func TestUnmarshalJSON2(t *testing.T) {
	g := NewGomegaWithT(t)
	s := `{"B":2,"D":"Tuesday","M":"November"}`
	var v Three
	err := json.Unmarshal([]byte(s), &v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(v).Should(Equal(Three{G, Tuesday, November}))
}

func TestGobEncodeAndDecode(t *testing.T) {
	g := NewGomegaWithT(t)
	v1 := Three{G, Tuesday, November}
	gob.Register(v1)

	// gob-encode
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(v1)
	g.Expect(err).NotTo(HaveOccurred())

	// gob-decode
	var v2 Three
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&v2)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(v2).Should(Equal(v1))
}
