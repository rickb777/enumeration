package main

import (
	"io"
	"bufio"
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

func convertHere(w io.Writer, s *bufio.Scanner, mainType string) (bool, error) {
	done := false

	for s.Scan() {
		line := strings.TrimSpace(removeComments(s.Text()))
		words := removeBlanks(strings.Split(line, " "))
		if len(words) == 1 && words[0] == ")" {
			return done, nil
		} else if len(words) >= 2 && words[1] == mainType {
		}
	}

	return false, nil
}

func convert(w io.Writer, s *bufio.Scanner, mainType string) error {
	//foundMainType := false

	for s.Scan() {
		line := strings.TrimSpace(removeComments(s.Text()))
		words := removeBlanks(strings.Split(line, " "))

		if len(words) == 3 && words[0] == "type" && words[1] == mainType {
			//foundMainType = true

		} else if len(words) == 2 && words[0] == "const" && words[1] == "(" {
			done, err := convertHere(w, s, mainType)
			if err != nil {
				return err
			}
			if done {
				return nil
			}
		}
	}

	return nil
}


