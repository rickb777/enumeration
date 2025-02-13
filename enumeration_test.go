package main

import (
	"fmt"
	"github.com/onsi/gomega"
	"github.com/rickb777/enumeration/v4/internal/model"
	"github.com/rickb777/enumeration/v4/internal/parse"
	"go/scanner"
	"go/token"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainApp_Day(t *testing.T) {
	g := gomega.NewWithT(t)
	outputFile := "example/day_enum.go"
	err := os.Remove(outputFile)
	if err != nil {
		t.Logf("rm %s: %s", outputFile, err.Error())
		// continue anyway
	}

	os.Args = []string{"", "-f", "-type", "Day", "-i", "example/day.go", "-o", outputFile}

	main()

	compareGeneratedFile(g, outputFile)
}

func TestMainApp_Channel(t *testing.T) {
	g := gomega.NewWithT(t)

	inputGo = "example/channel.go"
	outputGo = "example/channel_enum.go"

	err := os.Remove(outputGo)
	if err != nil {
		t.Logf("rm %s: %s", "example/channel_enum.go", err.Error())
		// continue anyway
	}

	force = true
	lowercase, uppercase, showVersion = false, false, false
	outputJSON, marshalTextRep, marshalJSONRep, storeRep = "", "None", "None", "None"

	model.Prefix = ""
	model.Suffix = "Sales"
	parse.AliasTable = ""

	config = model.Config{
		MainType: "SalesChannel",
	}

	doMain()

	compareGeneratedFile(g, outputGo)
}

func TestMainApp_Country(t *testing.T) {
	g := gomega.NewWithT(t)

	inputGo = "example/country.go"
	outputGo = "example/country_enum.go"

	err := os.Remove(outputGo)
	if err != nil {
		t.Logf("rm %s: %s", outputGo, err.Error())
		// continue anyway
	}

	force = true
	lowercase, uppercase, showVersion = false, false, false
	outputJSON, marshalTextRep, marshalJSONRep, storeRep = "", "None", "None", "Number"

	model.Prefix = ""
	model.Suffix = ""
	parse.AliasTable = "iso3166_3LetterCodes"

	config = model.Config{
		MainType:   "Country",
		Plural:     "Countries",
		IgnoreCase: true,
		Unsnake:    true,
	}

	doMain()

	compareGeneratedFile(g, outputGo)
}

func TestMainApp_Method(t *testing.T) {
	g := gomega.NewWithT(t)

	inputGo = "example/method.go"
	outputGo = "example/method_enum.go"

	err := os.Remove(outputGo)
	if err != nil {
		t.Logf("rm %s: %s", "example/method_enum.go", err.Error())
		// continue anyway
	}

	force = true
	lowercase, uppercase, showVersion = false, false, false
	outputJSON, marshalTextRep, marshalJSONRep, storeRep = "", "None", "None", "Number"

	model.Prefix = ""
	model.Suffix = ""
	parse.AliasTable = ""

	config = model.Config{
		MainType:   "Method",
		IgnoreCase: true,
	}

	doMain()

	compareGeneratedFile(g, outputGo)
}

func compareGeneratedFile(g *gomega.WithT, fileName string) {
	f, err := os.Open(fileName)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer f.Close()

	src, err := io.ReadAll(f)
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
		for _, n := range []string{"example/base.go", "example/day.go", "example/country.go", "example/month.go"} {
			f, err := os.Open(n)
			g.Expect(err).NotTo(gomega.HaveOccurred())
			defer f.Close()

			fmt.Printf("-- %s\n", n)
			src, err := io.ReadAll(f)
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
