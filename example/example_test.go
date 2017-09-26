package example

import (
	. "github.com/onsi/gomega"
	"testing"
	"encoding/json"
	"encoding/gob"
	"bytes"
)

func TestString(t *testing.T) {
	RegisterTestingT(t)
	Ω(Sunday.String()).Should(Equal("Sunday"))
	Ω(Monday.String()).Should(Equal("Monday"))
}

func TestOrdinal(t *testing.T) {
	RegisterTestingT(t)
	Ω(int(Sunday)).Should(Equal(1))
	Ω(Sunday.Ordinal(), 0)
	Ω(int(Monday)).Should(Equal(2))
	Ω(Monday.Ordinal(), 1)
	Ω(numberOfDays).Should(Equal(7))
}

func TestAllDays(t *testing.T) {
	RegisterTestingT(t)
	Ω(AllDays[0]).Should(Equal(Sunday))
	Ω(AllDays[5]).Should(Equal(Friday))
}

func TestAsDay(t *testing.T) {
	RegisterTestingT(t)
	v, err := AsDay("Tuesday")
	Ω(v, err).Should(Equal(Tuesday))
	_, err = AsDay("Nosuchday")
	Ω(err).ShouldNot(BeNil())
}

func TestMarshalText(t *testing.T) {
	RegisterTestingT(t)
	tt, err := Monday.MarshalText()
	Ω(tt, err).Should(Equal([]byte("Monday")))
}

func TestUnmarshalText(t *testing.T) {
	RegisterTestingT(t)
	var d = new(Day)
	err := d.UnmarshalText([]byte("Monday"))
	Ω(*d, err).Should(Equal(Monday))
}

//-------------------------------------------------------------------------------------------------

type Three struct {
	B Base
	D Day
	M Month
}

func TestMarshalJSON1(t *testing.T) {
	RegisterTestingT(t)

	BaseMarshalJSONUsingString = false
	DayMarshalJSONUsingString = false
	MonthMarshalJSONUsingString = false

	v := Three{G, Tuesday, November}
	s, err := json.Marshal(v)
	Ω(err).Should(BeNil())
	Ω(string(s)).Should(Equal(`{"B":2,"D":2,"M":10}`))
}

func TestMarshalJSON2(t *testing.T) {
	RegisterTestingT(t)

	BaseMarshalJSONUsingString = true
	DayMarshalJSONUsingString = true
	MonthMarshalJSONUsingString = true

	Ω(G.MarshalJSON()).Should(Equal([]byte{'"', 'G', '"'}))

	v := Three{G, Tuesday, November}
	s, err := json.Marshal(v)
	Ω(err).Should(BeNil())
	Ω(string(s)).Should(Equal(`{"B":"G","D":"Tuesday","M":"November"}`))
}

func TestUnmarshalJSON1(t *testing.T) {
	RegisterTestingT(t)
	s := `{"B":2,"D":2,"M":10}`
	var v Three
	err := json.Unmarshal([]byte(s), &v)
	Ω(err).Should(BeNil())
	Ω(v).Should(Equal(Three{G, Tuesday, November}))
}

func TestUnmarshalJSON2(t *testing.T) {
	RegisterTestingT(t)
	s := `{"B":2,"D":"Tuesday","M":"November"}`
	var v Three
	err := json.Unmarshal([]byte(s), &v)
	Ω(err).Should(BeNil())
	Ω(v).Should(Equal(Three{G, Tuesday, November}))
}

func TestGobEncodeAndDecode(t *testing.T) {
	RegisterTestingT(t)
	v1 := Three{G, Tuesday, November}
	gob.Register(v1)

	// gob-encode
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(v1)
	Ω(err).Should(BeNil())

	// gob-decode
	var v2 Three
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&v2)
	Ω(err).Should(BeNil())
	Ω(v2).Should(Equal(v1))
}

