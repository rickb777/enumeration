package parse

import (
	"go/token"
	"reflect"
	"strings"
)

type constItem struct {
	id, typ, expression string
	tag                 reflect.StructTag
}

func appendConstItems(items []constItem, ids []string, typ string, number string, tag reflect.StructTag) []constItem {
	if len(ids) == 1 {
		// include the tag
		return append(items, constItem{id: ids[0], typ: typ, expression: number, tag: tag})
	}

	// don't include the tag
	for _, id := range ids {
		items = append(items, constItem{id: id, typ: typ, expression: number})
	}
	return items
}

//-------------------------------------------------------------------------------------------------
// https://go.dev/doc/go1.17_spec#Constant_declarations
// ConstDecl      = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" ) .
// ConstSpec      = IdentifierList [ [ Type ] "=" ExpressionList ] .
//
// IdentifierList = identifier { "," identifier } .
// ExpressionList = Expression { "," Expression } .

func parseConst(s *scanner, items []constItem) []constItem {
	for s.Tok != token.EOF {
		switch s.Scan() {
		case token.IDENT:
			items = parseConstSpec(s, items)

		case token.LPAREN:
			return parseConstBlock(s, items)
		}
	}
	return items
}

func parseConstSpec(s *scanner, items []constItem) []constItem {
	ids := parseIdentifierList(s)

	// parse the Type and the ExpressionList
	for s.Scan() != token.EOF {
		switch s.Tok {
		case token.IDENT:
			typeName := s.Lit
			switch s.Scan() {
			case token.ASSIGN:
				restOfLine, tag := readToEndOfLine(s)
				return appendConstItems(items, ids, typeName, restOfLine, tag)
			}

		case token.ASSIGN:
			restOfLine, tag := readToEndOfLine(s)
			return appendConstItems(items, ids, "", restOfLine, tag)
		}
	}

	return items
}

func parseIdentifierList(s *scanner) []string {
	var ids []string
	for s.Tok == token.IDENT {
		ids = append(ids, s.Lit)

		if s.Peek() != token.COMMA {
			return ids
		}

		s.Scan() // the comma
		s.Scan() // the next ident?
	}

	return ids
}

func parseConstBlock(s *scanner, items []constItem) []constItem {
	for s.Scan() != token.EOF {
		switch s.Tok {
		case token.IDENT:
			ids := parseIdentifierList(s)

			switch s.Scan() {
			case token.IDENT:
				typeName := s.Lit
				if s.Scan() == token.ASSIGN {
					restOfLine, tag := readToEndOfLine(s)
					items = appendConstItems(items, ids, typeName, restOfLine, tag)
					ids = nil
				} else {
					readToEndOfLine(s) // discard likely compilation error
				}

			case token.SEMICOLON:
				restOfLine, tag := readToEndOfLine(s)
				items = appendConstItems(items, ids, "", restOfLine, tag)
				ids = nil
			}

		case token.RPAREN, token.EOF:
			return items

			//default:
			//	_, _ = readToEndOfLine(s)
		}
	}

	return items
}

func readToEndOfLine(s *scanner) (string, reflect.StructTag) {
	if s.Tok == token.ASSIGN {
		s.Scan()
	}

	var rest string
	for s.Tok != token.SEMICOLON && s.Tok != token.EOF {
		if rest != "" {
			rest += " "
		}
		if s.Lit != "" {
			rest += s.Lit
		} else {
			rest += s.Tok.String()
		}
		s.Scan()
	}

	comment := strings.TrimSpace(s.Lit)
	if strings.HasPrefix(comment, "//") {
		comment = strings.TrimSpace(comment[2:])
	}

	var commentTag reflect.StructTag
	if tagRE.MatchString(comment) {
		commentTag = reflect.StructTag(comment)
	}

	return rest, commentTag
}
