package transform

import (
	"fmt"
	"strings"
)

//go:generate enumeration -type Case

type Case int

const (
	Stet Case = iota
	Upper
	Lower
)

func Of(lowercase, uppercase bool) Case {
	if lowercase {
		return Lower
	} else if uppercase {
		return Upper
	}
	return Stet
}

func (c Case) Transform(s string) string {
	switch c {
	case Upper:
		return strings.ToUpper(s)
	case Lower:
		return strings.ToLower(s)
	}
	return s
}

func (c Case) Expression(s string) string {
	switch c {
	case Upper:
		return fmt.Sprintf("strings.ToUpper(%s)", s)
	case Lower:
		return fmt.Sprintf("strings.ToLower(%s)", s)
	}
	return s
}
