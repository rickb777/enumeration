package main

import (
	"io"
	"flag"
	"os"
	"strings"
	"bufio"
)

var input1 = flag.String("i", "", "Name of the input file (optional short form). May be '-' for stdin.")
var input2 = flag.String("input", "", "Name of the input file (optional long form).")
var output1 = flag.String("o", "", "Name of the output file (optional short form). May be '-' for stdout.")
var output2 = flag.String("output", "", "Name of the output file (optional long form).")
var pMainType = flag.String("type", "", "Name of the enumeration type (required).")
var force = flag.Bool("f", false, "Force output generation, even if up to date.")
var verbose = flag.Bool("v", false, "Verbose progress messages.")
var dbg = flag.Bool("z", false, "Debug messages.")

func choosePackage(outputFile string) (string, string) {
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

	return wd, pkg
}

func generate(mainType string) {
	debug("ReadFile %s\n", *input1)
	var err error

	var in io.Reader = os.Stdin
	if *input1 != "-" {
		in, err = os.Open(*input1)
		if err != nil {
			fail(err)
		}
	}

	var out io.Writer = os.Stdout
	if *output1 != "-" {
		out, err = os.Create(*output1)
		if err != nil {
			fail(err)
		}
		stdout = os.Stdout // ok because it's not going to be interleaved now
	}
	convert(out, bufio.NewScanner(in), mainType)
	info("Generated %s.\n", *output1)
}

func sPtr(s string) *string {
	return &s
}

func main() {
	flag.Parse()
	if pMainType == nil || *pMainType == "" {
		fail("Must specify -type.")
	}

	if input2 != nil {
		input1 = input2
	}
	if (input1 == nil || *input1 == "") && (input2 == nil || *input2 == "") {
		input1 = sPtr(strings.ToLower(*pMainType) + ".go")
	}

	if output2 != nil {
		output1 = output2
	}
	if (output1 == nil || *output1 == "") && (output2 == nil || *output2 == "") {
		output1 = sPtr(strings.ToLower(*pMainType) + "_enum.go")
	}

	generate(*pMainType)
}
