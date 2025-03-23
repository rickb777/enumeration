package main

import (
	"fmt"
	"github.com/rickb777/enumeration/v4/internal/model"
	"github.com/rickb777/enumeration/v4/internal/parse"
	"github.com/rickb777/expect"
	"go/scanner"
	"go/token"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainApp_Day(t *testing.T) {
	outputFile := "temp/example/day_enum.go"
	err := os.Remove(outputFile)
	if err != nil {
		t.Logf("rm %s: %s", outputFile, err.Error())
		// continue anyway
	}

	os.Args = []string{"", "-f", "-type", "Day", "-i", "temp/example/day.go", "-o", outputFile}

	main()

	compareGeneratedFile(t, outputFile)
}

func TestMainApp_Channel(t *testing.T) {
	inputGo = "temp/example/channel.go"
	outputGo = "temp/example/channel_enum.go"

	err := os.Remove(outputGo)
	if err != nil {
		t.Logf("rm %s: %s", "temp/example/channel_enum.go", err.Error())
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

	compareGeneratedFile(t, outputGo)
}

func TestMainApp_Country(t *testing.T) {
	inputGo = "temp/example/country.go"
	outputGo = "temp/example/country_enum.go"

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

	compareGeneratedFile(t, outputGo)
}

func TestMainApp_Method(t *testing.T) {
	inputGo = "temp/example/method.go"
	outputGo = "temp/example/method_enum.go"

	err := os.Remove(outputGo)
	if err != nil {
		t.Logf("rm %s: %s", "temp/example/method_enum.go", err.Error())
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

	compareGeneratedFile(t, outputGo)
}

func compareGeneratedFile(t *testing.T, fileName string) {
	f, err := os.Open(fileName)
	expect.Error(err).Not().ToHaveOccurred(t)
	defer f.Close()

	src, err := io.ReadAll(f)
	expect.Error(err).Not().ToHaveOccurred(t)

	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// check just the first few lines of the generated Go source code

	pos, tok, lit := s.Scan()
	expect.Number(tok).ToBe(t, token.COMMENT)

	_, tok, lit = s.Scan()
	expect.Number(tok).ToBe(t, token.COMMENT)

	_, tok, lit = s.Scan()
	expect.Number(tok).ToBe(t, token.PACKAGE)

	_, tok, lit = s.Scan()
	expect.Number(tok).ToBe(t, token.IDENT)
	expect.String(lit).ToBe(t, "example")

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
	if testing.Verbose() {
		for _, n := range []string{"example/base.go", "example/day.go", "example/country.go", "example/month.go"} {
			f, err := os.Open(n)
			expect.Error(err).Not().ToHaveOccurred(t)
			defer f.Close()

			fmt.Printf("-- %s\n", n)
			src, err := io.ReadAll(f)
			expect.Error(err).Not().ToHaveOccurred(t)

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
