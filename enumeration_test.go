package main

import (
	"fmt"
	"github.com/onsi/gomega"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMainApp(t *testing.T) {
	g := gomega.NewWithT(t)
	err := os.Remove("example/day_enum.go")
	if err != nil {
		t.Logf("rm example/day_enum.go: %s", err.Error())
		// continue anyway
	}

	os.Args = []string{"", "-f", "-type", "Day", "-i", "example/day.go", "-o", "example/day_enum.go"}
	main()

	f, err := os.Open("example/day_enum.go")
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer f.Close()

	src, err := ioutil.ReadAll(f)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// check just the first few lines of the generated Go source code

	pos, tok, lit := s.Scan()
	g.Expect(tok).To(gomega.Equal(token.COMMENT))

	_, tok, lit = s.Scan()
	g.Expect(tok).To(gomega.Equal(token.COMMENT))

	_, tok, lit = s.Scan()
	g.Expect(tok).To(gomega.Equal(token.PACKAGE))

	_, tok, lit = s.Scan()
	g.Expect(tok).To(gomega.Equal(token.IDENT))
	g.Expect(lit).To(gomega.Equal("example"))

	if testing.Verbose() {
		for {
			pos, tok, lit = s.Scan()
			if tok == token.EOF {
				break
			}
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
		}
	}
}

func TestScannerTryOut(t *testing.T) {
	g := gomega.NewWithT(t)

	if testing.Verbose() {
		for _, n := range []string{"example/base.go", "example/day.go", "example/month.go"} {
			f, err := os.Open(n)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			defer f.Close()

			fmt.Printf("-- %s\n", n)
			src, err := ioutil.ReadAll(f)
			g.Expect(err).NotTo(gomega.HaveOccurred())

			var s scanner.Scanner
			fset := token.NewFileSet()                      // positions are relative to fset
			file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
			s.Init(file, src, nil /* no error handler */, 0)

			for {
				pos, tok, lit := s.Scan()
				if tok == token.EOF {
					break
				}
				fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
			}
			fmt.Printf("%s\n", strings.Repeat("-", 80))
		}
	}
}
