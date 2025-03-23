package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"github.com/rickb777/expect"
	"testing"
)

func TestString(t *testing.T) {
	expect.String(Spring1.String()).ToBe(t, "Spring")
	expect.String(Spring_Nc_Ji.String()).ToBe(t, `Spring`)
	expect.String(Autumn_Nc_Jj.String()).ToBe(t, `Autumn`)
	expect.String(Spring_Ic_Ji.String()).ToBe(t, `Spring`)
	expect.String(Autumn_Ic_Jj.String()).ToBe(t, `Autumn`)
	expect.String(Spring_Uc_Ji.String()).ToBe(t, `SPRING`)
	expect.String(Autumn_Uc_Jj.String()).ToBe(t, `AUTUMN`)
}

func TestOrdinal(t *testing.T) {
	expect.Number(Spring1.Ordinal()).ToBe(t, 0)
}

func TestIntOrFloat(t *testing.T) {
	expect.Number(Spring1.Int()).ToBe(t, 1)
}

func TestAllDays(t *testing.T) {
	expect.Number(AllSeason1s[0]).ToBe(t, Spring1)
}

func TestAsSeason(t *testing.T) {
	v, err := AsSeason1("Spring")
	expect.Number(v, err).ToBe(t, Spring1)
	_, err = AsSeason1("Nosuchday")
	expect.Error(err).ToHaveOccurred(t)
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

	v := Group{A: Spring1, B: Summer2}
	s, err := json.Marshal(v)
	expect.Error(err).Not().ToHaveOccurred(t)
	x, err := xml.Marshal(v)
	expect.Error(err).Not().ToHaveOccurred(t)
	expect.String(s).ToEqual(t, `{"A":1,"B":2}`)
	expect.String(x).ToEqual(t, `<Group><A>1</A><B>2</B></Group>`)
}

func TestMarshal_for_Text(t *testing.T) {
	vs := []Group{
		{A: Spring_Nc_Ti, B: Summer_Nc_Tn, C: Autumn_Nc_Tt, D: Autumn_Nc_Ta},
		{A: Spring_Ic_Ti, B: Summer_Ic_Tn, C: Autumn_Ic_Tt, D: Autumn_Ic_Ta},
	}
	for _, v := range vs {
		s, err := json.Marshal(v)
		expect.Error(err).Not().ToHaveOccurred(t)
		x, err := xml.Marshal(v)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.String(s).ToEqual(t, `{"A":"Spring","B":"2","C":"Autm","D":"Autm"}`)
		expect.String(x).ToEqual(t, `<Group><A>Spring</A><B>2</B><C>Autm</C><D>Autm</D></Group>`)
	}

	expect.String(Spring_Nc_Ti.Text()).ToBe(t, `Spring`)
	expect.String(Autumn_Nc_Tt.Text()).ToBe(t, `Autm`)
	expect.String(Spring_Ic_Ti.Text()).ToBe(t, `Spring`)
	expect.String(Autumn_Ic_Tt.Text()).ToBe(t, `Autm`)
	expect.String(Spring_Uc_Ti.Text()).ToBe(t, `SPRING`)
	expect.String(Autumn_Uc_Ta.Text()).ToBe(t, `Autm`) // ignores UC
}

func TestMarshal_for_JSON(t *testing.T) {
	//v := Group{A: Spring_Nc_Ji, B: Summer_Nc_Jn, C: Autumn_Nc_Jo}
	vs := []Group{
		{A: Spring_Nc_Ji, B: Summer_Nc_Jn, C: Autumn_Nc_Jj},
		{A: Spring_Ic_Ji, B: Summer_Ic_Jn, C: Autumn_Ic_Jj},
	}
	for _, v := range vs {
		s, err := json.Marshal(v)
		expect.Error(err).Not().ToHaveOccurred(t)
		x, err := xml.Marshal(v)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.String(s).ToEqual(t, `{"A":"Spring","B":2,"C":"Autm"}`)
		expect.String(x).ToEqual(t, `<Group><A>1</A><B>2</B><C>3</C></Group>`)

		var v2 Group
		err = json.Unmarshal(s, &v2)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Any(v2.A).ToBe(t, "Spring")
		expect.Any(v2.B).ToEqual(t, 2)
		expect.Any(v2.C).ToBe(t, "Autm")
	}

	expect.String(Spring_Nc_Ji.JSON()).ToBe(t, `Spring`)
	expect.String(Autumn_Nc_Jj.JSON()).ToBe(t, `Autm`)
	expect.String(Spring_Ic_Ji.JSON()).ToBe(t, `Spring`)
	expect.String(Autumn_Ic_Jj.JSON()).ToBe(t, `Autm`)
	expect.String(Spring_Uc_Ji.JSON()).ToBe(t, `SPRING`)
	expect.String(Autumn_Uc_Jj.JSON()).ToBe(t, `Autm`) // ignores UC
}

func TestMethodScan_Nc_string_ok(t *testing.T) {
	cases := []interface{}{
		"Autumn", []byte("Autumn"),
	}
	for _, s := range cases {
		var mi = new(Season_Nc_Si)
		err := mi.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mi).ToBe(t, Autumn_Nc_Si)

		var mn = new(Season_Nc_Sn)
		err = mn.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mn).ToBe(t, Autumn_Nc_Sn)
	}
}

func TestMethodScan_Nc_number_ok(t *testing.T) {
	cases := []interface{}{
		int64(3), float64(3),
	}
	for _, s := range cases {
		var mi = new(Season_Nc_Si)
		err := mi.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mi).ToBe(t, Autumn_Nc_Si)

		var mn = new(Season_Nc_Sn)
		err = mn.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mn).ToBe(t, Autumn_Nc_Sn)
	}
}

func TestMethodScan_Ic_string_ok(t *testing.T) {
	cases := []interface{}{
		"Autumn", "AUTUMN", "autumn", []byte("Autumn"), []byte("AUTUMN"), []byte("autumn"),
	}
	for _, s := range cases {
		var mi = new(Season_Ic_Si)
		err := mi.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mi).ToBe(t, Autumn_Ic_Si)

		var mn = new(Season_Ic_Sn)
		err = mn.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mn).ToBe(t, Autumn_Ic_Sn)

		var ms = new(Season_Ic_Ss)
		err = ms.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*ms).ToBe(t, Autumn_Ic_Ss)

		var ma = new(Season_Ic_Ta)
		err = ma.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*ma).ToBe(t, Autumn_Ic_Ta)
	}
}

func TestMethodScan_Ic_number_ok(t *testing.T) {
	cases := []interface{}{
		int64(3), float64(3),
	}
	for _, s := range cases {
		var mi = new(Season_Ic_Si)
		err := mi.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mi).ToBe(t, Autumn_Ic_Si)

		var mn = new(Season_Ic_Sn)
		err = mn.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*mn).ToBe(t, Autumn_Ic_Sn)

		var ms = new(Season_Ic_Ss)
		err = ms.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*ms).ToBe(t, Autumn_Ic_Ss)

		var ma = new(Season_Ic_Ta)
		err = ma.Scan(s)
		expect.Error(err).Not().ToHaveOccurred(t)
		expect.Number(*ma).ToBe(t, Autumn_Ic_Ta)
	}
}

func TestValue(t *testing.T) {
	expect.Any(Spring_Nc_Si.Value()).ToBe(t, "Spring")
	expect.Any(Spring_Nc_Sn.Value()).ToBe(t, int64(1))
	expect.Any(Spring_Nc_Ss.Value()).ToBe(t, "Sprg")

	expect.Any(Spring_Ic_Si.Value()).ToBe(t, "Spring")
	expect.Any(Spring_Ic_Sn.Value()).ToBe(t, int64(1))
	expect.Any(Spring_Ic_Ss.Value()).ToBe(t, "Sprg")
}

func TestGobEncodeAndDecode(t *testing.T) {
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
	expect.Error(err).Not().ToHaveOccurred(t)

	// gob-decode
	var v2 Group
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&v2)
	expect.Error(err).Not().ToHaveOccurred(t)
	expect.Any(v2).ToBe(t, v1)
}
