// See https://magefile.org/

//go:build mage

// Build steps for the enumeration tool:
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() error {
	if err := os.MkdirAll("bin", 0777); err != nil {
		return err
	}

	x := ternary(mg.Verbose(), "-x ", "")

	if err := run(
		"go build -o bin/enumeration .",
		"go clean -testcache",
		"rm -f internal/test/*_enum.go internal/test/simple/*_enum.go",
		"go test ./internal/parse",
		"go generate "+x+"./internal/test",
		"rm -f example/*_enum.go",
		"go generate "+x+"./example",
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
		"cp -r example temp",
		"rm -f temp/example/*_*.go",
		"go test -cover ./... -coverprofile coverage.out -coverpkg ./...",
		"go vet ./...",
		"rm -rf temp/",
	); err != nil {
		return err
	}

	return nil
}

func Coverage() error {
	if err := sh.RunV("go", "tool", "cover", "-func", "coverage.out", "-o", "report.out"); err != nil {
		return err
	}
	return nil
}

// tests the module on both amd64 and i386 architectures for Linux and Windows
func CrossCompile() error {
	win := "build"
	linux := "test"
	if os.Getenv("GOOS") == "windows" {
		win = "test"
		linux = "build"
	}
	log.Printf("Testing on Windows\n")
	if err := sh.RunWithV(map[string]string{"GOOS": "windows"}, "go", win, "./..."); err != nil {
		return err
	}
	for _, arch := range []string{"amd64", "386"} {
		log.Printf("Testing on Linux/%s\n", arch)
		env := map[string]string{"GOOS": "linux", "GOARCH": arch}
		if _, err := sh.Exec(env, os.Stdout, os.Stderr, "go", linux, "./..."); err != nil {
			return err
		}
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
			fmt.Println(strings.Join(parts, " "))
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

func ternary(pred bool, a, b string) string {
	if pred {
		return a
	}
	return b
}
