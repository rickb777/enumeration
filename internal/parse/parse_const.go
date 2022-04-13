package parse

import (
	"go/scanner"
	"go/token"
)

func parseConstBlock(mainType string, s *scanner.Scanner, values []string) []string {
	foundType := false
	var ss []string
	for {
		_, tok, lit := scan(s)
		switch tok {
		case token.IDENT:
			ss = addIdentifier(ss, lit)

			_, tok, lit = scan(s)
			switch tok {
			case token.IDENT:
				if lit == mainType {
					foundType = true
					values = append(values, ss...)
				} else {
					foundType = false
				}
				ss = nil

			case token.COMMA:
				for tok == token.COMMA {
					_, tok, lit = scan(s)
					switch tok {
					case token.IDENT:
						ss = addIdentifier(ss, lit)
						_, tok, lit = scan(s)
					default:
						discardToEndOfLine(s, tok)
					}
				}

				if tok == token.IDENT && lit == mainType {
					foundType = true
					values = append(values, ss...)
				} else {
					foundType = false
				}
				ss = nil

			default:
				discardToEndOfLine(s, tok)
			}

		case token.RPAREN, token.EOF:
			if foundType {
				values = append(values, ss...)
			}
			return values

		default:
			discardToEndOfLine(s, tok)
		}
	}
}

func parseConst(mainType string, s *scanner.Scanner, values []string) []string {
	var tok token.Token
	var lit1, lit2 string
	_, tok, lit1 = scan(s)
	switch tok {
	case token.IDENT:
		_, tok, lit2 = scan(s)
		switch tok {
		case token.IDENT:
			if lit2 == mainType {
				values = addIdentifier(values, lit1)
			}
			discardToEndOfLine(s, tok)
		}
	case token.LPAREN:
		return parseConstBlock(mainType, s, values)
	}
	return values
}
