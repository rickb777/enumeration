package example

import (
	. "github.com/rickb777/terst"
	"testing"
)

func TestString(t *testing.T) {
	Terst(t)
	if Sunday.String() != "Sunday" {
		t.Errorf("Want Sunday, got %s", Sunday.String())
	}
	if Monday.String() != "Monday" {
		t.Errorf("Want Monday, got %s", Monday.String())
	}
	if numberOfDays.String() != "numberOfDays" {
		t.Errorf("Want numberOfDays, got %s", numberOfDays.String())
	}
}

func TestOrdinal(t *testing.T) {
	Terst(t)
	if Sunday.Ordinal() != 0 {
		t.Errorf("Want 0, got %d", Sunday.Ordinal())
	}
	if Monday.Ordinal() != 1 {
		t.Errorf("Want 1, got %d", Monday.Ordinal())
	}
	if numberOfDays.Ordinal() != 7 {
		t.Errorf("Want 7, got %d", numberOfDays.Ordinal())
	}
}

func TestAllDays(t *testing.T) {
	Terst(t)
	if AllDays[0] != "Sunday" {
		t.Errorf("Want Sunday, got %s", AllDays[0])
	}
	if AllDays[5] != "Friday" {
		t.Errorf("Want Friday, got %s", AllDays[5])
	}
}

func TestAsDay(t *testing.T) {
	Terst(t)
	v, err := AsDay("Tuesday")
	if err != nil {
		t.Errorf("got %v", err)
	}
	if v != Tuesday {
		t.Errorf("got %v", v)
	}
	_, err = AsDay("Nosuchday")
	if err == nil {
		t.Errorf("exxpected error")
	}
}
