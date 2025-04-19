// See https://magefile.org/

//go:build mage

// Build steps for enumeration:
package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"time"
)

var Default = Build

var (
	date    = time.Now().Format("2006-01-02")
	ldFlags = fmt.Sprintf(`-s -X main.version=%s -X main.date=%s`, gitDescribe(), date)
)

func Build() {
	os.Mkdir("bin", 0777)

	sh.RunV("go", "mod", "download")
	sh.RunV("go", "build", "-o", "bin/enumeration", "-ldflags", ldFlags, ".")
	sh.RunV("go", "clean", "testcache")
	rmGlob("internal/test/*_enum.go")
	sh.RunV("go", "test", "./internal/parse")

	sh.RunV("./internal/test/generate.sh")
	sh.RunV("./example/generate.sh")

	os.Mkdir("temp/example", 0777) // used in ./enumeration_test.go
	glob(func(files ...string) {
		args := append(files, "temp/example")
		sh.RunV("cp", args...)
	}, "example/*.go")

	rmGlob("temp/example/*_enum.go")
	rmGlob("temp/example/*_test.go")
	sh.RunV("go", "test", "./...")

	Gofmt()
}

func Gofmt() {
	glob(func(files ...string) {
		args := append([]string{"-l", "-w", "-s"}, files...)
		sh.Run("gofmt", args...)
	}, "*.go", "*/*.go", "*/*/*.go", "*/*/*/*.go")
	time.Sleep(250 * time.Millisecond) // wait for the files to be stable
}

func rmGlob(pattern string) {
	glob(func(file ...string) {
		for _, f := range file {
			os.Remove(f)
		}
	}, pattern)
}

func glob(cmd func(file ...string), pattern ...string) {
	for _, p := range pattern {
		files, err := filepath.Glob(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		cmd(files...)
	}
}

func gitDescribe() string {
	s, err := sh.Output("git", "describe", "--tags", "--always", "--dirty")
	if err != nil {
		return "dev"
	}
	return s
}
