package codegen

import (
	"github.com/rickb777/enumeration/v4/internal/collection"
	"slices"
	"unicode"
)

type Unit struct {
	Declares   string
	Requires   []string
	Extra      map[string]any
	Imports    collection.Set[string]
	Transforms bool
	Template   string
}

func (u Unit) Exported() bool {
	d := u.Declares
	if d[:2] == "v." {
		d = d[2:]
	}
	return unicode.IsUpper(rune(d[0]))
}

type Units struct {
	m map[string]Unit
	l []string
}

func New() *Units {
	return &Units{m: make(map[string]Unit)}
}

func (units *Units) Add(unit Unit) *Units {
	_, exists := units.m[unit.Declares]
	if !exists {
		units.l = append(units.l, unit.Declares)
		units.m[unit.Declares] = unit
	}
	return units
}

func (units *Units) Take(identifier string) (u Unit, found bool) {
	units.l = slices.DeleteFunc(units.l, func(id string) bool {
		return id == identifier
	})
	u, found = units.m[identifier]
	//delete(units.m, identifier)
	return u, found
}

func (units *Units) Slice() []Unit {
	us := make([]Unit, 0, len(units.m))
	for _, id := range units.l {
		us = append(us, units.m[id])
	}
	return us
}
