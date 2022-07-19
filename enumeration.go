package main

import (
	"flag"
	"fmt"
	"github.com/rickb777/enumeration/v2/enum"
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
var inputGo, outputGo, outputJSON, marshalTextRep, storeRep string
var force, lowercase, uppercase, showVersion bool

func defineFlags() {
	flag.StringVar(&config.MainType, "type", "", "Name of the enumeration type (required).")
	flag.StringVar(&model.Prefix, "prefix", "", "Optional prefix to be stripped from the identifiers.")
	flag.StringVar(&model.Suffix, "suffix", "", "Optional suffix to be stripped from the identifiers.")
	flag.StringVar(&inputGo, "i", "", "Name of the input file. May be '-' for stdin. Default is enumeration type in lower case.")
	flag.StringVar(&outputGo, "o", "", "Name of the output file. May be '-' for stdout. Default is enumeration type in lower case plus '_enum'.")
	flag.StringVar(&config.Plural, "plural", "", "Plural name of the enumeration type (optional).")
	flag.StringVar(&parse.UsingTable, "using", "", "Uses your own map[Type]string instead of generating one.")
	flag.StringVar(&parse.AliasTable, "alias", "", "Uses your own map[string]Type as aliases during parsing.")
	flag.StringVar(&config.Pkg, "package", "", "Name of the output package (optional). Defaults to the output directory.")
	flag.StringVar(&marshalTextRep, "marshaltext", "Identifier", "Marshal values using Identifier, Tag, Number or Ordinal")
	flag.StringVar(&storeRep, "store", "Identifier", "Store values in a DB using Identifier, Tag, Number or Ordinal")

	flag.BoolVar(&config.Lenient, "lenient", false, "Allow parsing to yield invalid values.")
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
	if inputGo != "-" && outputGo != "-" {
		xi, err := os.Stat(inputGo)
		if err == nil {
			xo, err := os.Stat(outputGo)
			if err == nil && xo.ModTime().After(xi.ModTime()) {
				util.Info("Skipped %s.\n", outputGo)
				return false
			}
		}
	}
	return true
}

func generate() {
	util.Debug("ReadFile %s\n", inputGo)
	var err error
	config.MarshalTextRep, err = enum.AsRepresentation(marshalTextRep)
	util.Must(err, "(-marshaltext)")

	config.StoreRep, err = enum.AsRepresentation(storeRep)
	util.Must(err, "(-store)")

	var in io.Reader = os.Stdin
	if inputGo != "-" {
		inf, e2 := os.Open(inputGo)
		util.Must(e2)
		defer inf.Close()
		in = inf
	}

	var out io.Writer = os.Stdout
	if outputGo == "-" {
		if config.Pkg == "" {
			util.Fail("-pkg is required when piping the output.")
		}
	} else {
		outf, e2 := os.Create(outputGo)
		util.Must(e2)
		defer outf.Close()
		out = outf
		config.Pkg = choosePackage(outputGo)
		util.Stdout = os.Stdout // ok because it's not going to be interleaved now
	}
	util.Debug("pkg=%s\n", config.Pkg)

	xCase := transform.Of(lowercase, uppercase)

	m, err := parse.Convert(in, inputGo, xCase, config)
	util.Must(err)

	m.WriteGo(out)
	util.Info("Generated %s.\n", outputGo)
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

	if inputGo == "" {
		inputGo = strings.ToLower(config.MainType) + ".go"
	}

	if outputGo == "" {
		outputGo = strings.ToLower(config.MainType) + "_enum.go"
	} else if outputGo == "-" {
		util.Stdout = os.Stderr // avoiding interleaving with the output of generated code
	}

	if outputJSON == "" {
		outputJSON = strings.ToLower(config.MainType) + "_enum.json"
	}

	util.Debug("type=%s\n", config.MainType)
	util.Debug("plural=%s\n", config.Plural)
	util.Debug("inputGo=%s\n", inputGo)
	util.Debug("outputGo=%s\n", outputGo)
	util.Debug("outputJSON=%s\n", outputJSON)

	if force || notUpToDate() {
		generate()
	}
}
