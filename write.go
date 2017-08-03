package main

import (
	"fmt"
	"io"
	"strings"
)

//-------------------------------------------------------------------------------------------------

const head = `// generated code - do not edit

package %s

import (
	"errors"
	"fmt"
	"strconv"
)

`

func writeHead(w io.Writer, pkg string) error {
	_, err := fmt.Fprintf(w, head, pkg)
	return err
}

//-------------------------------------------------------------------------------------------------

func writeConst(w io.Writer, names string, values []string, xf func(string) string) error {
	_, err := fmt.Fprintf(w, "const %s = \"", names)
	if err != nil {
		return err
	}

	for _, s := range values {
		_, err = fmt.Fprintf(w, xf(s))
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, "\"\n\n")
	return err
}

//-------------------------------------------------------------------------------------------------

func writeIndexes(w io.Writer, index string, values []string) error {
	_, err := fmt.Fprintf(w, "var %s = [...]uint16{0", index)
	if err != nil {
		return err
	}

	n := 0
	for _, s := range values {
		n += len(s)
		_, err = fmt.Fprintf(w, ", %d", n)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, "}\n\n")
	return err
}

//-------------------------------------------------------------------------------------------------

func writeAllItemsSlice(w io.Writer, mainType, plural string, values []string) error {
	_, err := fmt.Fprintf(w, "var All%s = []%s{", plural, mainType)
	if err != nil {
		return err
	}

	comma := ""
	for _, s := range values {
		_, err = fmt.Fprintf(w, "%s%s", comma, s)
		if err != nil {
			return err
		}
		comma = ", "
	}

	_, err = fmt.Fprintf(w, "}\n")
	return err
}

//-------------------------------------------------------------------------------------------------

const stringMethod = `
// String returns the string representation of a %s
func (i %s) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(All%s) {
		return fmt.Sprintf("%s(%%v)", i)
	}
	return %s[%s[o]:%s[o+1]]
}
`

func writeFuncString(w io.Writer, mainType, plural, names, indexes string) error {
	_, err := fmt.Fprintf(w, stringMethod, mainType, mainType, plural, mainType, names, indexes, indexes)
	return err
}

//-------------------------------------------------------------------------------------------------

const ordinalMethodStart = `
// Ordinal returns the ordinal number of a %s
func (i %s) Ordinal() int {
	switch i {
`
const ordinalMethodEnd = `	}
	return -1
}
`

func writeFuncOrdinal(w io.Writer, mainType string, values []string) error {
	_, err := fmt.Fprintf(w, ordinalMethodStart, mainType, mainType)
	if err != nil {
		return err
	}

	for i, s := range values {
		_, err = fmt.Fprintf(w, "\tcase %s:\n\t\treturn %d\n", s, i)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, ordinalMethodEnd)
	return err
}

//-------------------------------------------------------------------------------------------------

const parseMethod = `
// Parse parses a string to find the corresponding %s, accepting either one of the string
// values or an ordinal number.
func (v *%s) Parse(s string) error {
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(All%s) {
		*v = All%s[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(%s); j++ {
		i1 := %s[j]
		p := %s[i0:i1]
		if s == p {
			*v = All%s[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised %s")
}
`
func writeFuncParse(w io.Writer, mainType, plural, names, indexes string) error {
	_, err := fmt.Fprintf(w, parseMethod, mainType, mainType, plural, plural, indexes, indexes, names, plural, mainType)
	return err
}

//-------------------------------------------------------------------------------------------------

const asMethod = `
// As%s parses a string to find the corresponding %s, accepting either one of the string
// values or an ordinal number.
func As%s(s string) (%s, error) {
	var i = new(%s)
	err := i.Parse(s)
	return *i, err
}
`

func writeFuncAs(w io.Writer, mainType, plural, names, indexes string) error {
	_, err := fmt.Fprintf(w, asMethod, mainType, mainType, mainType, mainType, mainType)
	return err
}

//-------------------------------------------------------------------------------------------------

const marshalText = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i %s) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *%s) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

func writeMarshalText(w io.Writer, mainType string) error {
	_, err := fmt.Fprintf(w, marshalText, mainType, mainType)
	return err
}

//-------------------------------------------------------------------------------------------------

func write(w io.Writer, mainType, baseType, plural, pkg string, values []string, xf func(string) string) error {

	lc := strings.ToLower(mainType)
	names := fmt.Sprintf("%sEnumStrings", lc)
	indexes := fmt.Sprintf("%sEnumIndex", lc)

	err := writeHead(w, pkg)
	if err != nil {
		return err
	}

	err = writeConst(w, names, values, xf)
	if err != nil {
		return err
	}

	err = writeIndexes(w, indexes, values)
	if err != nil {
		return err
	}

	err = writeAllItemsSlice(w, mainType, plural, values)
	if err != nil {
		return err
	}

	err = writeFuncString(w, mainType, plural, names, indexes)
	if err != nil {
		return err
	}

	err = writeFuncOrdinal(w, mainType, values)
	if err != nil {
		return err
	}

	err = writeFuncParse(w, mainType, plural, names, indexes)
	if err != nil {
		return err
	}

	err = writeFuncAs(w, mainType, plural, names, indexes)
	if err != nil {
		return err
	}

	err = writeMarshalText(w, mainType)
	if err != nil {
		return err
	}

	if c, ok := w.(io.Closer); ok {
		err = c.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
