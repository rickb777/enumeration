package parse

import (
	"go/scanner"
	"go/token"
)

//TODO allow for var ( .. ) syntax

func parseVar(mainType string, s *scanner.Scanner, tags map[string]string) map[string]string {
	_, tok, lit := scan(s)
	if lit != UsingTable {
		discardToEndOfLine(s, tok)
		return nil
	}

	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil
	case token.IDENT: // continue
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil
	case token.ASSIGN: // continue
	}

	_, tok, lit = scan(s)
	if lit != "map" {
		discardToEndOfLine(s, tok)
		return nil
	}
	//if lit2 == mainType {
	//	tags = addIdentifier(tags, lit1)
	//}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil
	case token.LBRACK: // continue
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil
	case token.IDENT: // continue
		if lit != mainType {
			discardToEndOfLine(s, tok)
			return nil
		}
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil
	case token.RBRACK: // continue
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil
	case token.IDENT:
		if lit != "string" {
			discardToEndOfLine(s, tok)
			return nil
		}
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil
	case token.LBRACE: // continue
	}

	var more bool
	tags, more = parseTags(s, tags)
	for more {
		tags, more = parseTags(s, tags)
	}
	return tags
}

func parseTags(s *scanner.Scanner, tags map[string]string) (map[string]string, bool) {
	var key string
	_, tok, lit := scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil, false
	case token.IDENT:
		key = lit
	case token.RBRACE: // finished
		return tags, false
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil, false
	case token.COLON: // continue
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil, false
	case token.STRING:
		unquoted := lit[1 : len(lit)-1]
		tags[key] = unquoted
	}

	_, tok, lit = scan(s)
	switch tok {
	default:
		discardToEndOfLine(s, tok)
		return nil, false
	case token.COMMA: // continue
	}

	return tags, true
}
