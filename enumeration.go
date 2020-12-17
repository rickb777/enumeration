package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var input1 = flag.String("i", "", "Name of the input file. May be '-' for stdin. Default is enumeration type in lower case.")
var output1 = flag.String("o", "", "Name of the output file. May be '-' for stdout. Default is enumeration type in lower case plus '_enum'.")

var pMainType = flag.String("type", "", "Name of the enumeration type (required).")
var pPlural = flag.String("plural", "", "Plural name of the enumeration type (optional).")
var usingTable = flag.String("using", "", "Uses your own map[Type]string instead of generating one.")

//var force = flag.Bool("f", false, "Force output generation, even if up to date.")
var pPkg = flag.String("package", "", "Name of the output package (optional). Defaults to the output directory).")
var force = flag.Bool("f", false, "Force writing the output file even if up to date (not used when piping stdin or stdout).")
var lowercase = flag.Bool("lc", false, "Convert strings to lowercase.")
var uppercase = flag.Bool("uc", false, "Convert strings to uppercase.")
var unsnake = flag.Bool("unsnake", false, "Convert underscores in identifiers to spaces.")
var verbose = flag.Bool("v", false, "Verbose progress messages.")
var dbg = flag.Bool("z", false, "Debug messages.")
var showVersion = flag.Bool("version", false, "Print version number.")

func choosePackage(outputFile string) string {
	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	pkg := removeBeforeC(wd, '/')

	if strings.IndexByte(outputFile, '/') > 0 {
		dir, _ := divideC(outputFile, '/')
		if strings.IndexByte(dir, '/') > 0 {
			dir = removeBeforeC(dir, '/')
		}
		if dir != "." {
			pkg = dir
		}
	}

	return pkg
}

func notUpToDate() bool {
	if *input1 != "-" && *output1 != "-" {
		xi, err := os.Stat(*input1)
		if err == nil {
			xo, err := os.Stat(*output1)
			if err == nil && xo.ModTime().After(xi.ModTime()) {
				info("Skipped %s.\n", *output1)
				return false
			}
		}
	}
	return true
}

func generate(mainType, plural string) {
	debug("ReadFile %s\n", *input1)
	var err error

	var in io.Reader = os.Stdin
	if *input1 != "-" {
		in, err = os.Open(*input1)
		if err != nil {
			fail(err)
		}
	}

	var pkg = ""
	var out io.Writer = os.Stdout
	if *output1 != "-" {
		out, err = os.Create(*output1)
		if err != nil {
			fail(err)
		}
		pkg = choosePackage(*output1)
		stdout = os.Stdout // ok because it's not going to be interleaved now
	} else {
		if pPkg == nil || *pPkg == "" {
			fail("-pkg is required when piping the output.")
		}
		pkg = *pPkg
	}
	debug("pkg=%s\n", pkg)

	var transforms []Transformer
	if *lowercase {
		transforms = append(transforms, toLower)
	} else if *uppercase {
		transforms = append(transforms, toUpper)
	}
	if *unsnake {
		transforms = append(transforms, xUnsnake)
	}

	err = convert(out, in, *input1, mainType, plural, pkg, transforms...)
	if err != nil {
		fail(err)
	}
	info("Generated %s.\n", *output1)
}

func sPtr(s string) *string {
	return &s
}

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Fprintln(os.Stderr, version)
		os.Exit(1)
	}

	if pMainType == nil || *pMainType == "" {
		fail("Must specify -type.")
	}

	mainType := *pMainType
	plural := mainType + "s"
	if pPlural != nil && *pPlural != "" {
		plural = *pPlural
	}

	if input1 == nil || *input1 == "" {
		input1 = sPtr(strings.ToLower(mainType) + ".go")
	}

	if output1 == nil || *output1 == "" {
		output1 = sPtr(strings.ToLower(mainType) + "_enum.go")
	}

	if *output1 == "-" {
		stdout = os.Stderr // avoiding interleaving with the output of generated code
	}

	debug("type=%s\n", mainType)
	debug("plural=%s\n", plural)
	debug("input=%s\n", *input1)
	debug("output=%s\n", *output1)

	if *force || notUpToDate() {
		generate(mainType, plural)
	}
}
