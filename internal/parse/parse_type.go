package parse

import (
	"fmt"
	"go/token"
	"math"
)

func parseType(s *scanner, mainType string) (string, error) {
	for s.Tok != token.EOF {
		switch s.Scan() {
		case token.IDENT:
			return parseTypeSpec(s, mainType, 1)

		case token.LPAREN:
			switch s.Scan() {
			case token.IDENT:
				return parseTypeSpec(s, mainType, math.MaxInt)
			}
		}
	}

	return "", nil
}

func parseTypeSpec(s *scanner, mainType string, inBlock int) (string, error) {
	for s.Tok != token.EOF && inBlock > 0 {
		if s.Tok == token.IDENT && s.Lit == mainType {
			switch s.Scan() {
			case token.IDENT:
				switch s.Lit {
				case "int", "uint",
					"int8", "uint8",
					"int16", "uint16",
					"int32", "uint32",
					"int64", "uint64":
					return s.Lit, nil
				case "float32", "float64":
					return s.Lit, nil
				default:
					return "", fmt.Errorf("type %s must be an integer or float type", mainType)
				}

				//util.Debug("type %s %s\n", MainType, baseType)
			default:

			}
		}
		s.Scan()
		inBlock--
	}
	return "", nil
}
