package parse

import (
	"github.com/rickb777/enumeration/v2/internal/util"
	goscanner "go/scanner"
	"go/token"
)

// scanner implements a one-place lookahead wrapper around the Go scanner.
// It also coalesces SEMICOLON and COMMENT into one apparent SEMICOLON
// with the literal from the COMMENT.
type scanner struct {
	gs      *goscanner.Scanner
	Tok     token.Token
	Lit     string
	nextTok token.Token
	nextLit string
}

func (s *scanner) Peek() token.Token {
	return s.nextTok
}

func (s *scanner) doScan() {
	var pos token.Pos
	pos, s.nextTok, s.nextLit = s.gs.Scan()
	if s.Lit == "" {
		util.Debug("%-18s %s\n", fset.Position(pos), s.Tok)
	} else {
		util.Debug("%-18s %-8s %q\n", fset.Position(pos), s.Tok, s.Lit)
	}
}

func (s *scanner) Scan() token.Token {
	switch s.Tok {
	case token.EOF:
		return token.EOF
	}

	s.Tok = s.nextTok
	s.Lit = s.nextLit

	s.doScan()

	switch s.nextTok {
	case token.COMMENT:
		s.Lit = s.nextLit
		s.doScan()
	}
	return s.Tok
}

func newFileScanner(input string, src []byte) *scanner {
	gs := new(goscanner.Scanner)
	fset = token.NewFileSet()                          // positions are relative to fset
	file := fset.AddFile(input, fset.Base(), len(src)) // register input "file"
	gs.Init(file, src, nil /* no error handler */, goscanner.ScanComments)
	return &scanner{gs: gs}
}
