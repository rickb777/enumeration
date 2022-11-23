package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"testing"

	. "github.com/onsi/gomega"
)

func TestString(t *testing.T) {
	g := NewWithT(t)
	g.Expect(Spring1.String()).Should(Equal("Spring"))
	g.Expect(Spring_Nc_Ji.String()).Should(Equal(`Spring`))
	g.Expect(Autumn_Nc_Jj.String()).Should(Equal(`Autumn`))
	g.Expect(Spring_Ic_Ji.String()).Should(Equal(`Spring`))
	g.Expect(Autumn_Ic_Jj.String()).Should(Equal(`Autumn`))
	g.Expect(Spring_Uc_Ji.String()).Should(Equal(`SPRING`))
	g.Expect(Autumn_Uc_Jj.String()).Should(Equal(`AUTUMN`))
}

func TestOrdinal(t *testing.T) {
	g := NewWithT(t)
	g.Expect(Spring1.Ordinal()).Should(Equal(0))
}

func TestIntOrFloat(t *testing.T) {
	g := NewWithT(t)
	g.Expect(Spring1.Int()).Should(Equal(1))
}

func TestAllDays(t *testing.T) {
	g := NewWithT(t)
	g.Expect(AllSeason1s[0]).Should(Equal(Spring1))
}

func TestAsSeason(t *testing.T) {
	g := NewWithT(t)
	v, err := AsSeason1("Spring")
	g.Expect(v, err).Should(Equal(Spring1))
	_, err = AsSeason1("Nosuchday")
	g.Expect(err).Should(HaveOccurred())
}

//-------------------------------------------------------------------------------------------------

type Group struct {
	A interface{} `json:"A,omitempty"`
	B interface{} `json:"B,omitempty"`
	C interface{} `json:"C,omitempty"`
	D interface{} `json:"D,omitempty"`
	E interface{} `json:"E,omitempty"`
}

func TestMarshal_plain(t *testing.T) {
	g := NewWithT(t)

	v := Group{A: Spring1, B: Summer2}
	s, err := json.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	x, err := xml.Marshal(v)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(s)).Should(Equal(`{"A":1,"B":2}`))
	g.Expect(string(x)).Should(Equal(`<Group><A>1</A><B>2</B></Group>`), string(x))
}

func TestMarshal_for_Text(t *testing.T) {
	g := NewWithT(t)

	vs := []Group{
		{A: Spring_Nc_Ti, B: Summer_Nc_Tn, C: Autumn_Nc_Tt, D: Autumn_Nc_Ta},
		{A: Spring_Ic_Ti, B: Summer_Ic_Tn, C: Autumn_Ic_Tt, D: Autumn_Ic_Ta},
	}
	for _, v := range vs {
		s, err := json.Marshal(v)
		g.Expect(err).NotTo(HaveOccurred())
		x, err := xml.Marshal(v)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(s)).Should(Equal(`{"A":"Spring","B":"2","C":"Autm","D":"Autm"}`))
		g.Expect(string(x)).Should(Equal(`<Group><A>Spring</A><B>2</B><C>Autm</C><D>Autm</D></Group>`), string(x))
	}

	g.Expect(Spring_Nc_Ti.Text()).Should(Equal(`Spring`))
	g.Expect(Autumn_Nc_Tt.Text()).Should(Equal(`Autm`))
	g.Expect(Spring_Ic_Ti.Text()).Should(Equal(`Spring`))
	g.Expect(Autumn_Ic_Tt.Text()).Should(Equal(`Autm`))
	g.Expect(Spring_Uc_Ti.Text()).Should(Equal(`SPRING`))
	g.Expect(Autumn_Uc_Ta.Text()).Should(Equal(`Autm`)) // ignores UC
}

func TestMarshal_for_JSON(t *testing.T) {
	g := NewWithT(t)

	//v := Group{A: Spring_Nc_Ji, B: Summer_Nc_Jn, C: Autumn_Nc_Jo}
	vs := []Group{
		{A: Spring_Nc_Ji, B: Summer_Nc_Jn, C: Autumn_Nc_Jj},
		{A: Spring_Ic_Ji, B: Summer_Ic_Jn, C: Autumn_Ic_Jj},
	}
	for _, v := range vs {
		s, err := json.Marshal(v)
		g.Expect(err).NotTo(HaveOccurred())
		x, err := xml.Marshal(v)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(string(s)).Should(Equal(`{"A":"Spring","B":2,"C":"Autm"}`))
		g.Expect(string(x)).Should(Equal(`<Group><A>1</A><B>2</B><C>3</C></Group>`), string(x))

		var v2 Group
		err = json.Unmarshal(s, &v2)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(v2.A).Should(Equal("Spring"), string(x))
		g.Expect(v2.B).Should(BeEquivalentTo(2), string(x))
		g.Expect(v2.C).Should(Equal("Autm"), string(x))
	}

	g.Expect(Spring_Nc_Ji.JSON()).Should(Equal(`Spring`))
	g.Expect(Autumn_Nc_Jj.JSON()).Should(Equal(`Autm`))
	g.Expect(Spring_Ic_Ji.JSON()).Should(Equal(`Spring`))
	g.Expect(Autumn_Ic_Jj.JSON()).Should(Equal(`Autm`))
	g.Expect(Spring_Uc_Ji.JSON()).Should(Equal(`SPRING`))
	g.Expect(Autumn_Uc_Jj.JSON()).Should(Equal(`Autm`)) // ignores UC
}

func TestMethodScan_Nc_string_ok(t *testing.T) {
	g := NewWithT(t)

	cases := []interface{}{
		"Autumn", []byte("Autumn"),
	}
	for _, s := range cases {
		var mi = new(Season_Nc_Si)
		err := mi.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mi).Should(Equal(Autumn_Nc_Si))

		var mn = new(Season_Nc_Sn)
		err = mn.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mn).Should(Equal(Autumn_Nc_Sn))
	}
}

func TestMethodScan_Nc_number_ok(t *testing.T) {
	g := NewWithT(t)

	cases := []interface{}{
		int64(3), float64(3),
	}
	for _, s := range cases {
		var mi = new(Season_Nc_Si)
		err := mi.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mi).Should(Equal(Autumn_Nc_Si))

		var mn = new(Season_Nc_Sn)
		err = mn.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mn).Should(Equal(Autumn_Nc_Sn))
	}
}

func TestMethodScan_Ic_string_ok(t *testing.T) {
	g := NewWithT(t)

	cases := []interface{}{
		"Autumn", "AUTUMN", "autumn", []byte("Autumn"), []byte("AUTUMN"), []byte("autumn"),
	}
	for _, s := range cases {
		var mi = new(Season_Ic_Si)
		err := mi.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mi).Should(Equal(Autumn_Ic_Si))

		var mn = new(Season_Ic_Sn)
		err = mn.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mn).Should(Equal(Autumn_Ic_Sn))

		var ms = new(Season_Ic_Ss)
		err = ms.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*ms).Should(Equal(Autumn_Ic_Ss))

		var ma = new(Season_Ic_Ta)
		err = ma.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*ma).Should(Equal(Autumn_Ic_Ta))
	}
}

func TestMethodScan_Ic_number_ok(t *testing.T) {
	g := NewWithT(t)

	cases := []interface{}{
		int64(3), float64(3),
	}
	for _, s := range cases {
		var mi = new(Season_Ic_Si)
		err := mi.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mi).Should(Equal(Autumn_Ic_Si))

		var mn = new(Season_Ic_Sn)
		err = mn.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*mn).Should(Equal(Autumn_Ic_Sn))

		var ms = new(Season_Ic_Ss)
		err = ms.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*ms).Should(Equal(Autumn_Ic_Ss))

		var ma = new(Season_Ic_Ta)
		err = ma.Scan(s)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(*ma).Should(Equal(Autumn_Ic_Ta))
	}
}

func TestValue(t *testing.T) {
	g := NewWithT(t)

	g.Expect(Spring_Nc_Si.Value()).To(Equal("Spring"))
	g.Expect(Spring_Nc_Sn.Value()).To(Equal(int64(1)))
	g.Expect(Spring_Nc_Ss.Value()).To(Equal("Sprg"))

	g.Expect(Spring_Ic_Si.Value()).To(Equal("Spring"))
	g.Expect(Spring_Ic_Sn.Value()).To(Equal(int64(1)))
	g.Expect(Spring_Ic_Ss.Value()).To(Equal("Sprg"))
}

func TestGobEncodeAndDecode(t *testing.T) {
	g := NewWithT(t)
	v1 := Group{A: Spring_Nc_Ti, B: Summer_Nc_Tn, C: Autumn_Nc_Tt, D: Autumn_Nc_Ta}
	gob.Register(v1)
	gob.Register(Spring_Nc_Ti)
	gob.Register(Summer_Nc_Tn)
	gob.Register(Autumn_Nc_Tt)
	gob.Register(Autumn_Nc_Ta)

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
