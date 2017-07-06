package example

import (
	. "github.com/rickb777/terst"
	"testing"
)

func TestString(t *testing.T) {
	Terst(t)
	Is(Sunday.String(), "Sunday")
	Is(Monday.String(), "Monday")
}

func TestOrdinal(t *testing.T) {
	Terst(t)
	Is(int(Sunday), 1)
	Is(Sunday.Ordinal(), 0)
	Is(int(Monday), 2)
	Is(Monday.Ordinal(), 1)
	Is(numberOfDays, 7)
}

func TestAllDays(t *testing.T) {
	Terst(t)
	Is(AllDays[0], Sunday)
	Is(AllDays[5], Friday)
}

func TestAsDay(t *testing.T) {
	Terst(t)
	v, err := AsDay("Tuesday")
	Is(err, nil)
	Is(v, Tuesday)
	_, err = AsDay("Nosuchday")
	Is(err, "!=", nil)
}

func TestMarshalText(t *testing.T) {
	Terst(t)
	tt, err := Monday.MarshalText()
	Is(err, nil)
	Is(tt, []byte("Monday"))
}

func TestUnmarshalText(t *testing.T) {
	Terst(t)
	var d = new(Day)
	err := d.UnmarshalText([]byte("Monday"))
	Is(err, nil)
	Is(*d, Monday)
}
