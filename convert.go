package main

import (
	"bufio"
	"io"
	"strings"
)

func removeComments(line string) string {
	return removeAfterS(line, "//")
}

func doRemoveBlanks(words []string) []string {
	copy := make([]string, 0)
	for _, w := range words {
		if w != " " {
			copy = append(copy, w)
		}
	}
	return copy
}

func removeBlanks(words []string) []string {
	for _, w := range words {
		if w == " " {
			return doRemoveBlanks(words)
		}
	}
	return words
}

func scanValues(s *bufio.Scanner, mainType string) (result []string) {
	debug("scanValues\n")
	found := false
	for s.Scan() {
		line := s.Text()
		content := strings.TrimSpace(removeComments(line))
		words := removeBlanks(strings.Split(content, " "))
		debug("%#v\n", words)

		if len(words) == 1 && words[0] == ")" {
			return
		} else if found && len(words) >= 1 && len(words[0]) > 0 {
			debug("added %s\n", words[0])
			result = append(result, words[0])
		} else if len(words) >= 3 && len(words[0]) > 0 && words[1] == mainType {
			found = true
			debug("started with %s\n", words[0])
			result = append(result, words[0])
		}
	}

	return
}

func convert(w io.Writer, in io.Reader, input, mainType, plural, pkg string) error {
	foundMainType := false
	baseType := "int"
	s := bufio.NewScanner(in)

	for s.Scan() {
		line := strings.TrimSpace(removeComments(s.Text()))
		words := removeBlanks(strings.Split(line, " "))
		debug("%#v\n", words)

		if len(words) == 3 && words[0] == "type" && words[1] == mainType {
			debug("Found type %s\n", mainType)
			foundMainType = true
			baseType = words[2]

		} else if foundMainType && len(words) == 2 && words[0] == "const" && words[1] == "(" {
			values := scanValues(s, mainType)
			if values != nil {
				return write(w, mainType, baseType, plural, pkg, values)
			}
		}
	}

	fail("Failed to find", mainType, "in", input)
	return nil
}
