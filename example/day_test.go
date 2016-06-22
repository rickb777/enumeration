package example

import (
	. "github.com/rickb777/terst"
	"testing"
)

func TestString(t *testing.T) {
	Terst(t)
	Is(Sunday.String(), "Sunday")
	Is(Monday.String(), "Monday")
	Is(numberOfDays.String(), "numberOfDays")
}

func TestOrdinal(t *testing.T) {
	Terst(t)
	Is(Sunday.Ordinal(), 0)
	Is(Monday.Ordinal(), 1)
	Is(numberOfDays.Ordinal(), 7)
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

func TestBinary(t *testing.T) {
	Terst(t)
	//err := Monday.MarshalBinary()
}
