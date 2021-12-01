package parse

import (
	"fmt"
	"github.com/rickb777/enumeration/v2/internal/model"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"github.com/rickb777/enumeration/v2/internal/util"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"strings"
	"unicode"
)

var UsingTable string
var fset *token.FileSet

func isExported(s string) bool {
	for i, r := range s {
		if i > 0 {
			return false
		}
		return unicode.IsUpper(r)
	}
	return false
}

func addIdentifier(ss []string, id string) []string {
	if isExported(id) {
		ss = append(ss, id)
	}
	return ss
}

func scan(s *scanner.Scanner) (token.Pos, token.Token, string) {
	pos, tok, lit := s.Scan()
	if lit == "" {
		util.Debug("%-18s %s\n", fset.Position(pos), tok)
	} else {
		util.Debug("%-18s %-8s %q\n", fset.Position(pos), tok, lit)
	}
	return pos, tok, lit
}

func parseConstBlock(mainType string, s *scanner.Scanner, m *model.Model) error {
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
					m.Values = append(m.Values, ss...)
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
					m.Values = append(m.Values, ss...)
				} else {
					foundType = false
				}
				ss = nil

			default:
				discardToEndOfLine(s, tok)
			}

		case token.RPAREN, token.EOF:
			if foundType {
				m.Values = append(m.Values, ss...)
			}
			return nil

		default:
			discardToEndOfLine(s, tok)
		}
	}
}

func discardToEndOfLine(s *scanner.Scanner, tok token.Token) {
	for tok != token.SEMICOLON && tok != token.EOF {
		_, tok, _ = scan(s)
	}
}

func parseConst(mainType string, s *scanner.Scanner, m *model.Model) error {
	var tok token.Token
	var lit1, lit2 string
	_, tok, lit1 = scan(s)
	switch tok {
	case token.IDENT:
		_, tok, lit2 = scan(s)
		switch tok {
		case token.IDENT:
			if lit2 == mainType {
				m.Values = addIdentifier(m.Values, lit1)
			}
			discardToEndOfLine(s, tok)
		}
	case token.LPAREN:
		return parseConstBlock(mainType, s, m)
	}
	return nil
}

func Convert(in io.Reader, input, mainType, plural, pkg string, xCase transform.Case, ignoreCase, unsnake bool) (model.Model, error) {
	foundMainType := false
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return model.Model{}, err
	}

	s := new(scanner.Scanner)
	fset = token.NewFileSet()                          // positions are relative to fset
	file := fset.AddFile(input, fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, 0)

	m := &model.Model{
		MainType:    mainType,
		LcType:      strings.ToLower(mainType),
		BaseType:    "int",
		Plural:      plural,
		Pkg:         pkg,
		Version:     util.Version,
		IgnoreCase:  ignoreCase,
		Unsnake:     unsnake,
		Case:        xCase,
		LookupTable: UsingTable,
	}

	var tok token.Token
	var lit string

	for tok != token.EOF {
		_, tok, lit = scan(s)
		switch tok {
		case token.TYPE:
			_, tok, lit = scan(s)
			if tok == token.IDENT && lit == mainType {
				foundMainType = true

				_, tok, lit = scan(s)
				if tok == token.IDENT {
					m.BaseType = lit
					util.Debug("type %s %s\n", mainType, m.BaseType)
				}
			}

		case token.CONST:
			_ = parseConst(mainType, s, m)
		}
	}

	if s.ErrorCount > 0 {
		return model.Model{}, fmt.Errorf("Syntax error in %s", input)
	}

	if !foundMainType || len(m.Values) == 0 {
		return model.Model{}, fmt.Errorf("Failed to find %s in %s", mainType, input)
	}

	return *m, nil
}
