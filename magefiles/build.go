// See https://magefile.org/

//go:build mage

// Build steps for the enumeration tool:
package main

import (
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Default = Build

func Build() error {
	if err := os.MkdirAll("bin", 0777); err != nil {
		return err
	}

	if err := run(
		"go build -o bin/enumeration .",
		"go clean -testcache",
		"rm -f internal/test/*_enum.go internal/test/simple/*_enum.go",
		"go test ./internal/parse",
		"go generate ./internal/test",
		"rm -f example/*_enum.go",
		"go generate -x ./example",
		"gofmt -l -w -s .",
	); err != nil {
		return err
	}

	time.Sleep(250 * time.Millisecond) // wait for the files to be stable

	//--- the following section is used in ./enumeration_test.go ---

	if err := os.MkdirAll("temp/example", 0777); err != nil {
		return err
	}

	if err := run(
		"rm -f temp/example/*_enum.go temp/example/*_test.go",
		"cp example/*.go temp/example",
		"go test ./...",
		"go vet ./...",
	); err != nil {
		return err
	}

	return nil
}

func run(cmds ...string) error {
	for _, cmd := range cmds {
		parts := dropBlanks(strings.Split(cmd, " "))
		if len(parts) > 0 {
			expanded, err := expandGlob(parts[1:])
			if err != nil {
				return err
			}
			err = sh.RunWith(env, parts[0], expanded...)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func dropBlanks(value []string) []string {
	words := make([]string, 0, len(value))
	for _, word := range value {
		word = strings.TrimSpace(word)
		if word != "" {
			words = append(words, word)
		}
	}
	return words
}

func expandGlob(value []string) ([]string, error) {
	var files []string
	for _, v := range value {
		matches, err := filepath.Glob(v)
		if err != nil {
			return nil, err
		}
		if len(matches) > 0 {
			for _, x := range matches {
				files = append(files, x)
			}
		} else {
			files = append(files, v)
		}
	}
	return files, nil
}

var env = setupPathWithBinDir()

func setupPathWithBinDir() map[string]string {
	dir, _ := os.Getwd()
	bindir := filepath.Join(dir, "bin")
	path := bindir + ":" + os.Getenv("PATH")
	return map[string]string{"PATH": path}
}
