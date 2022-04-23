package parse

import (
	"github.com/rickb777/enumeration/v2/internal/model"
	"go/scanner"
	"go/token"
)

type constItem struct {
	id, typ, number string
}

func appendConstItems(values []constItem, ids []string, typ string, number string) []constItem {
	for _, id := range ids {
		values = append(values, constItem{id: id, typ: typ, number: number})
	}
	return values
}

func parseConstBlock(mainType string, s *scanner.Scanner, values []constItem) []constItem {
	var iotaType, restOfLine string
	var tok token.Token
	var lit1, lit2, lit3, lit4 string
	var ids []string

	for {
		_, tok, lit1 = scan(s)
		switch tok {
		case token.IDENT:
			ids = append(ids, lit1)
			_, tok, lit2 = scan(s)
			switch tok {
			case token.IDENT:
				_, tok, _ = scan(s)
				if tok == token.ASSIGN {
					_, tok, lit3 = scan(s)
					switch tok {
					case token.IDENT, token.INT, token.FLOAT:
						tok, restOfLine, _ = readToEndOfLine(s, tok, lit3)
						values = appendConstItems(values, ids, lit2, restOfLine)
						if lit3 == "iota" {
							iotaType = lit2
						} else {
							iotaType = ""
						}
					}
					ids = nil
				}

			case token.COMMA:
				for tok == token.COMMA {
					_, tok, lit2 = scan(s)
					switch tok {
					case token.IDENT:
						ids = append(ids, lit2)
						_, tok, lit3 = scan(s)
					default:
						readToEndOfLine(s, tok, lit2)
					}
				}

				if tok == token.IDENT {
					_, tok, _ = scan(s)
					if tok == token.ASSIGN {
						_, tok, lit4 = scan(s)
						values = appendConstItems(values, ids, lit3, lit4)
					}
				}
				ids = nil

			default:
				readToEndOfLine(s, tok, lit2)
			}

		case token.RPAREN, token.EOF:
			if len(ids) > 0 && iotaType != "" {
				values = appendConstItems(values, ids, iotaType, "")
			}
			return values

		default:
			readToEndOfLine(s, tok, lit1)
		}
	}
}

func parseConst(mainType string, s *scanner.Scanner, values []constItem) []constItem {
	var tok token.Token
	var lit1, lit2, lit3 string
	_, tok, lit1 = scan(s)
	switch tok {
	case token.IDENT:
		_, tok, lit2 = scan(s)
		switch tok {
		case token.IDENT:
			if lit2 == mainType {
				_, tok, _ = scan(s)
				if tok == token.ASSIGN {
					_, tok, lit3 = scan(s)
					values = append(values, constItem{id: lit1, typ: lit2, number: lit3})
				}
			}
			readToEndOfLine(s, tok, "")
		}
	case token.LPAREN:
		return parseConstBlock(mainType, s, values)
	}
	return values
}

func filterExported(mainType string, ids []constItem) (exported model.Values, defaultValue string) {
	var currentType string
	var hasIota bool
	exported = make(model.Values, 0, len(ids))

	for _, v := range ids {
		if v.typ == mainType {
			if token.IsExported(v.id) {
				exported = exported.Append(v.id)
				switch v.number {
				case "0", "iota":
					defaultValue = v.id
				}
			}

			if v.number == "iota" {
				hasIota = true
			} else if v.typ != "" {
				hasIota = false
			}

		} else if v.typ == "" && currentType == mainType {
			if token.IsExported(v.id) {
				exported = exported.Append(v.id)
			}
		}

		if hasIota && v.typ != "" {
			currentType = v.typ
			hasIota = false
		}
	}

	return exported, defaultValue
}
