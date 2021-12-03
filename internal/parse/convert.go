package parse

import (
	"fmt"
	"github.com/rickb777/enumeration/v2/internal/model"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"github.com/rickb777/enumeration/v2/internal/util"
	"go/scanner"
	"go/token"
	"io"
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

func discardToEndOfLine(s *scanner.Scanner, tok token.Token) {
	for tok != token.SEMICOLON && tok != token.EOF {
		_, tok, _ = scan(s)
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

func newFileScanner(input string, src []byte) *scanner.Scanner {
	s := new(scanner.Scanner)
	fset = token.NewFileSet()                          // positions are relative to fset
	file := fset.AddFile(input, fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, 0)
	return s
}

func Convert(in io.Reader, input string, xCase transform.Case, config model.Config) (model.Model, error) {
	src, err := io.ReadAll(in)
	if err != nil {
		return model.Model{}, err
	}

	m := model.Model{
		Config:      config,
		LcType:      strings.ToLower(config.MainType),
		BaseType:    "int",
		Version:     util.Version,
		Case:        xCase,
		LookupTable: UsingTable,
	}

	s := newFileScanner(input, src)

	var foundMainType = false
	var tok token.Token
	var lit string

	for tok != token.EOF {
		_, tok, lit = scan(s)
		switch tok {
		case token.TYPE:
			_, tok, lit = scan(s)
			if tok == token.IDENT && lit == config.MainType {
				foundMainType = true

				_, tok, lit = scan(s)
				if tok == token.IDENT {
					m.BaseType = lit
					util.Debug("type %s %s\n", config.MainType, m.BaseType)
				}
			}

		case token.CONST:
			m.Values = parseConst(config.MainType, s, m.Values)
		}
	}

	if s.ErrorCount > 0 {
		return model.Model{}, fmt.Errorf("Syntax error in %s", input)
	}

	if !foundMainType || len(m.Values) == 0 {
		return model.Model{}, fmt.Errorf("Failed to find %s in %s", config.MainType, input)
	}

	if e2 := m.CheckBadPrefixSuffix(); e2 != nil {
		return model.Model{}, e2
	}

	return m, nil
}
