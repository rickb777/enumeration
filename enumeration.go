package main

import (
	"flag"
	"fmt"
	"github.com/rickb777/enumeration/v2/internal/model"
	"github.com/rickb777/enumeration/v2/internal/parse"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"github.com/rickb777/enumeration/v2/internal/util"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var config model.Config
var input1, output1 string
var force, lowercase, uppercase, showVersion bool

func defineFlags() {
	flag.StringVar(&config.MainType, "type", "", "Name of the enumeration type (required).")
	flag.StringVar(&config.Prefix, "prefix", "", "Optional prefix to be stripped from the identifiers.")
	flag.StringVar(&config.Suffix, "suffix", "", "Optional suffix to be stripped from the identifiers.")
	flag.StringVar(&input1, "i", "", "Name of the input file. May be '-' for stdin. Default is enumeration type in lower case.")
	flag.StringVar(&output1, "o", "", "Name of the output file. May be '-' for stdout. Default is enumeration type in lower case plus '_enum'.")
	flag.StringVar(&config.Plural, "plural", "", "Plural name of the enumeration type (optional).")
	flag.StringVar(&parse.UsingTable, "using", "", "Uses your own map[Type]string instead of generating one.")
	flag.StringVar(&config.Pkg, "package", "", "Name of the output package (optional). Defaults to the output directory.")

	flag.BoolVar(&force, "f", false, "Force writing the output file even if up to date (not used when piping stdin or stdout).")
	flag.BoolVar(&lowercase, "lc", false, "Convert strings to lowercase and ignore case when parsing")
	flag.BoolVar(&uppercase, "uc", false, "Convert strings to uppercase and ignore case when parsing.")
	flag.BoolVar(&config.IgnoreCase, "ic", false, "Ignore case when parsing but keep the mixed case when outputting.")
	flag.BoolVar(&config.Unsnake, "unsnake", false, "Convert underscores in identifiers to spaces.")
	flag.BoolVar(&util.Verbose, "v", false, "Verbose progress messages.")
	flag.BoolVar(&util.Dbg, "z", false, "Debug messages.")
	flag.BoolVar(&showVersion, "version", false, "Print version number.")
}

func choosePackage(outputFile string) string {
	wd, err := os.Getwd()
	if err != nil {
		util.Fail(err)
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
				util.Info("Skipped %s.\n", output1)
				return false
			}
		}
	}
	return true
}

func generate() {
	util.Debug("ReadFile %s\n", input1)
	var err error

	var in io.Reader = os.Stdin
	if input1 != "-" {
		in, err = os.Open(input1)
		if err != nil {
			util.Fail(err)
		}
	}

	var out io.Writer = os.Stdout
	if output1 == "-" {
		if config.Pkg == "" {
			util.Fail("-pkg is required when piping the output.")
		}
	} else {
		out, err = os.Create(output1)
		if err != nil {
			util.Fail(err)
		}
		config.Pkg = choosePackage(output1)
		util.Stdout = os.Stdout // ok because it's not going to be interleaved now
	}
	util.Debug("pkg=%s\n", config.Pkg)

	xCase := transform.Of(lowercase, uppercase)

	m, err := parse.Convert(in, input1, xCase, config)
	if err != nil {
		util.Fail(err)
	}

	m.Write(out)
	util.Info("Generated %s.\n", output1)
}

func main() {
	defineFlags()
	flag.Parse()

	if showVersion {
		fmt.Fprintln(os.Stderr, util.Version)
		os.Exit(1)
	}

	if config.MainType == "" {
		util.Fail("Must specify -type.")
	}

	if config.Plural == "" {
		config.Plural = config.MainType + "s"
	}

	if input1 == "" {
		input1 = strings.ToLower(config.MainType) + ".go"
	}

	if output1 == "" {
		output1 = strings.ToLower(config.MainType) + "_enum.go"
	}

	if output1 == "-" {
		util.Stdout = os.Stderr // avoiding interleaving with the output of generated code
	}

	util.Debug("type=%s\n", config.MainType)
	util.Debug("plural=%s\n", config.Plural)
	util.Debug("input=%s\n", input1)
	util.Debug("output=%s\n", output1)

	if force || notUpToDate() {
		generate()
	}
}
