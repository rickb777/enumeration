package parse

import (
	"github.com/rickb777/enumeration/v2/internal/util"
	goscanner "go/scanner"
	"go/token"
)

// scanner implements a one-place lookahead wrapper around the Go scanner
type scanner struct {
	gs      *goscanner.Scanner
	Tok     token.Token
	Lit     string
	prevTok token.Token // ideally this buffer should be replaced by a small ring buffer
	prevLit string
	nextTok token.Token
	nextLit string
}

func (s *scanner) Unscan() token.Token {
	if s.nextTok != 0 {
		panic(s.nextTok.String() + s.nextLit)
	}
	s.nextTok = s.Tok
	s.nextLit = s.Lit
	s.Tok = s.prevTok
	s.Lit = s.prevLit
	s.prevTok = 0
	s.prevLit = ""
	return s.Tok
}

func (s *scanner) Scan() token.Token {
	if s.nextTok != 0 {
		s.prevTok = s.Tok
		s.prevLit = s.Lit
		s.Tok = s.nextTok
		s.Lit = s.nextLit
		s.nextTok = 0
		s.nextLit = ""
		return s.Tok
	}

	if s.Tok == token.EOF {
		return token.EOF
	}

	s.prevTok = s.Tok
	s.prevLit = s.Lit

	var pos token.Pos
	pos, s.Tok, s.Lit = s.gs.Scan()
	if s.Lit == "" {
		util.Debug("%-18s %s\n", fset.Position(pos), s.Tok)
	} else {
		util.Debug("%-18s %-8s %q\n", fset.Position(pos), s.Tok, s.Lit)
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
