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
)

var UsingTable string
var fset *token.FileSet

func scan(s *scanner.Scanner) (token.Pos, token.Token, string) {
	pos, tok, lit := s.Scan()
	if lit == "" {
		util.Debug("%-18s %s\n", fset.Position(pos), tok)
	} else {
		util.Debug("%-18s %-8s %q\n", fset.Position(pos), tok, lit)
	}
	return pos, tok, lit
}

func discardToEndOfLine(s *scanner.Scanner, tok token.Token, lit string) (rest string) {
	for tok != token.SEMICOLON && tok != token.EOF {
		if lit != "" {
			rest += lit
		} else {
			rest += tok.String()
		}
		_, tok, lit = scan(s)
	}
	return rest
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
	var constItems []constItem

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
			constItems = parseConst(config.MainType, s, constItems)

		case token.VAR:
			if len(m.Tags) == 0 {
				m.Tags = parseVar(config.MainType, s, make(map[string]string))
			}
		}
	}

	m.Values, m.DefaultValue = filterExported(config.MainType, constItems)

	if s.ErrorCount > 0 {
		return model.Model{}, fmt.Errorf("Syntax error in %s", input)
	}

	if !foundMainType {
		return model.Model{}, fmt.Errorf("Failed to find %s in %s", config.MainType, input)
	}

	if len(m.Values) == 0 {
		return model.Model{}, fmt.Errorf("Failed to find any values for %s in %s", config.MainType, input)
	}

	if len(m.Tags) > 0 {
		if len(m.Tags) < len(m.Values) {
			return model.Model{}, fmt.Errorf("Too few values in %s for %s (%s)", UsingTable, config.MainType, input)
		}
		var blanks []string
		for key, tag := range m.Tags {
			if tag == "" {
				blanks = append(blanks, key)
			}
		}
		if len(blanks) > 1 {
			return model.Model{}, fmt.Errorf("Blank tag for %s %v in %s (%s)", config.MainType, blanks, UsingTable, input)
		}
	}

	if e2 := m.CheckBadPrefixSuffix(); e2 != nil {
		return model.Model{}, e2
	}

	return m, nil
}
