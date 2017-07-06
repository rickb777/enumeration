package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func removeComments(line string) string {
	return removeAfterS(line, "//")
}

func removeBlanks(words []string) []string {
	cp := make([]string, 0, len(words))
	for _, w := range words {
		if w != "" {
			cp = append(cp, w)
		}
	}
	return cp
}

func removeCommentsAndSplitWords(line string) []string {
	content := strings.TrimSpace(removeComments(line))
	return removeBlanks(strings.Split(content, " "))
}

func scanValues(s *bufio.Scanner, mainType string) (result []string) {
	debug("scanValues\n")
	found := false
	for s.Scan() {
		words := removeCommentsAndSplitWords(s.Text())
		debug("%#v\n", words)

		if len(words) == 1 && words[0] == ")" {
			if found {
				return
			}
		}

		eq := listIndexOf(words, "=")
		if eq >= 2 && len(words) >= 3 && words[eq-1] == mainType {
			found = true
			for i := 0; i < eq-1; i++ {
				names := removeBlanks(strings.Split(words[i], ","))
				debug("started with %s\n", names)
				result = append(result, names...)
			}
		} else if found && eq < 0 && len(words) >= 1 {
			debug("added %s\n", words[0])
			result = append(result, words[0])
		}
	}

	return
}

func convert(w io.Writer, in io.Reader, input, mainType, plural, pkg string, xf func(string) string) error {
	foundMainType := false
	baseType := "int"
	s := bufio.NewScanner(in)

	for s.Scan() {
		words := removeCommentsAndSplitWords(s.Text())
		debug("%#v\n", words)

		if len(words) == 3 && words[0] == "type" && words[1] == mainType {
			foundMainType = true
			baseType = words[2]
			debug("type %s %s\n", mainType, baseType)

		} else if foundMainType && len(words) == 2 && words[0] == "const" && words[1] == "(" {
			values := scanValues(s, mainType)
			if values != nil {
				return write(w, mainType, baseType, plural, pkg, values, xf)
			}
		}
	}

	return fmt.Errorf("Failed to find %s in %s", mainType, input)
}
