package parse

import (
	"fmt"
	"github.com/rickb777/enumeration/v4/internal/util"
	goscanner "go/scanner"
	"go/token"
)

// scanner implements a one-place lookahead wrapper around the Go scanner.
// It also coalesces SEMICOLON and COMMENT into one apparent SEMICOLON
// with the literal from the COMMENT.
type scanner struct {
	gs      *goscanner.Scanner
	errs    []string
	Pos     token.Pos
	Tok     token.Token
	Lit     string
	nextPos token.Pos
	nextTok token.Token
	nextLit string
}

func (s *scanner) Peek() token.Token {
	return s.nextTok
}

func (s *scanner) doScan() {
	s.nextPos, s.nextTok, s.nextLit = s.gs.Scan()
	s.debug()
}

func (s *scanner) debug() {
	if s.Lit == "" {
		util.Debug("%-18s %s\n", fset.Position(s.Pos), s.Tok)
	} else {
		util.Debug("%-18s %-8s %q\n", fset.Position(s.Pos), s.Tok, s.Lit)
	}
}

func (s *scanner) Position() token.Position {
	return fset.Position(s.Pos)
}

func (s *scanner) Scan() token.Token {
	switch s.Tok {
	case token.EOF:
		return token.EOF
	}

	s.Pos = s.nextPos
	s.Tok = s.nextTok
	s.Lit = s.nextLit

	s.doScan()

	return s.Tok
}

func newFileScanner(input string, src []byte) *scanner {
	gs := new(goscanner.Scanner)
	sc := &scanner{gs: gs}
	eh := func(pos token.Position, msg string) {
		sc.errs = append(sc.errs, fmt.Sprintf("%s: %s", pos, msg))
	}

	fset = token.NewFileSet()                          // positions are relative to fset
	file := fset.AddFile(input, fset.Base(), len(src)) // register input "file"
	gs.Init(file, src, eh, goscanner.ScanComments)
	return sc
}
