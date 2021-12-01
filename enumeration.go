package main

import (
	"flag"
	"fmt"
	"github.com/rickb777/enumeration/v2/transform"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var input1, output1, pkg, mainType, plural, prefix, suffix, usingTable string
var force, lowercase, uppercase, ignorecase, unsnake, verbose, dbg, showVersion bool

func defineFlags() {
	flag.StringVar(&mainType, "type", "", "Name of the enumeration type (required).")
	flag.StringVar(&prefix, "prefix", "", "Optional prefix to be stripped from the identifiers.")
	flag.StringVar(&suffix, "suffix", "", "Optional suffix to be stripped from the identifiers.")
	flag.StringVar(&input1, "i", "", "Name of the input file. May be '-' for stdin. Default is enumeration type in lower case.")
	flag.StringVar(&output1, "o", "", "Name of the output file. May be '-' for stdout. Default is enumeration type in lower case plus '_enum'.")
	flag.StringVar(&plural, "plural", "", "Plural name of the enumeration type (optional).")
	flag.StringVar(&usingTable, "using", "", "Uses your own map[Type]string instead of generating one.")
	flag.StringVar(&pkg, "package", "", "Name of the output package (optional). Defaults to the output directory.")

	flag.BoolVar(&force, "f", false, "Force writing the output file even if up to date (not used when piping stdin or stdout).")
	flag.BoolVar(&lowercase, "lc", false, "Convert strings to lowercase and ignore case when parsing")
	flag.BoolVar(&uppercase, "uc", false, "Convert strings to uppercase and ignore case when parsing.")
	flag.BoolVar(&ignorecase, "ic", false, "Ignore case when parsing but keep the mixed case when outputting.")
	flag.BoolVar(&unsnake, "unsnake", false, "Convert underscores in identifiers to spaces.")
	flag.BoolVar(&verbose, "v", false, "Verbose progress messages.")
	flag.BoolVar(&dbg, "z", false, "Debug messages.")
	flag.BoolVar(&showVersion, "version", false, "Print version number.")
}

func choosePackage(outputFile string) string {
	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	dir := filepath.Base(filepath.Dir(outputFile))

	if dir != "." {
		return dir
	}

	return filepath.Base(filepath.FromSlash(wd))
}

func notUpToDate() bool {
	if input1 != "-" && output1 != "-" {
		xi, err := os.Stat(input1)
		if err == nil {
			xo, err := os.Stat(output1)
			if err == nil && xo.ModTime().After(xi.ModTime()) {
				info("Skipped %s.\n", output1)
				return false
			}
		}
	}
	return true
}

func generate(mainType, plural string) {
	debug("ReadFile %s\n", input1)
	var err error

	var in io.Reader = os.Stdin
	if input1 != "-" {
		in, err = os.Open(input1)
		if err != nil {
			fail(err)
		}
	}

	var out io.Writer = os.Stdout
	if output1 != "-" {
		out, err = os.Create(output1)
		if err != nil {
			fail(err)
		}
		pkg = choosePackage(output1)
		stdout = os.Stdout // ok because it's not going to be interleaved now
	} else {
		if pkg == "" {
			fail("-pkg is required when piping the output.")
		}
	}
	debug("pkg=%s\n", pkg)

	xCase := transform.Of(lowercase, uppercase)

	m, err := convert(in, input1, mainType, plural, pkg, xCase, ignorecase, unsnake)
	if err != nil {
		fail(err)
	}
	m.write(out)
	info("Generated %s.\n", output1)
}

func main() {
	defineFlags()
	flag.Parse()

	if showVersion {
		fmt.Fprintln(os.Stderr, version)
		os.Exit(1)
	}

	if mainType == "" {
		fail("Must specify -type.")
	}

	if plural == "" {
		plural = mainType + "s"
	}

	if input1 == "" {
		input1 = strings.ToLower(mainType) + ".go"
	}

	if output1 == "" {
		output1 = strings.ToLower(mainType) + "_enum.go"
	}

	if output1 == "-" {
		stdout = os.Stderr // avoiding interleaving with the output of generated code
	}

	debug("type=%s\n", mainType)
	debug("plural=%s\n", plural)
	debug("input=%s\n", input1)
	debug("output=%s\n", output1)

	if force || notUpToDate() {
		generate(mainType, plural)
	}
}
