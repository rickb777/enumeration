package parse

import (
	"fmt"
	"go/token"
	"go/types"
	"math"
)

func parseType(s *scanner, mainType string, numFound int) (string, types.BasicKind, error) {
	for s.Tok != token.EOF {
		switch s.Scan() {
		case token.IDENT:
			return parseTypeSpec(s, mainType, numFound, 1)

		case token.LPAREN:
			switch s.Scan() {
			case token.IDENT:
				return parseTypeSpec(s, mainType, numFound, math.MaxInt)
			}
		}
	}

	return "", 0, nil
}

func parseTypeSpec(s *scanner, mainType string, numFound int, inBlock int) (string, types.BasicKind, error) {
	for s.Tok != token.EOF && inBlock > 0 {
		if s.Tok == token.IDENT && s.Lit == mainType {
			if numFound > 0 {
				return "", types.Invalid, fmt.Errorf("found multiple type %s declarations", mainType)
			}
			switch s.Scan() {
			case token.IDENT:
				switch s.Lit {
				case "int", "uint",
					"int8", "uint8",
					"int16", "uint16",
					"int32", "uint32",
					"int64", "uint64",
					"byte", "rune":
					return s.Lit, types.Int, nil
				case "float32", "float64":
					return s.Lit, types.Float64, nil
				default:
					return "", types.Invalid, fmt.Errorf("enumeration type %s must be an integer or float type", mainType)
				}

			case token.ASSIGN:
				return "", types.Invalid, fmt.Errorf("type %s is a type alias (not supported)", mainType)

			default:
				return "", types.Invalid, fmt.Errorf("syntax error in type %s declaration", mainType)
			}
		}
		s.Scan()
		inBlock--
	}

	return "", 0, nil
}
