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
)

`

func writeHead(w io.Writer, pkg string) error {
	_, err := fmt.Fprintf(w, head, pkg)
	return err
}

//-------------------------------------------------------------------------------------------------

func writeConst(w io.Writer, names string, values []string) error {
	_, err := fmt.Fprintf(w, "const %s = \"", names)
	if err != nil {
		return err
	}

	for _, s := range values {
		_, err = fmt.Fprintf(w, s)
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
	if i < 0 || i >= %s(len(%s)-1) {
		return fmt.Sprintf("%s(%%d)", i)
	}
	return %s[%s[i]:%s[i+1]]
}
`

func writeFuncString(w io.Writer, mainType, names, indexes string) error {
	_, err := fmt.Fprintf(w, stringMethod, mainType, mainType, mainType, indexes, mainType, names, indexes, indexes)
	return err
}

//-------------------------------------------------------------------------------------------------

const ordinalMethod1 = `
// Ordinal returns the ordinal number of a %s
func (i %s) Ordinal() int {
	switch i {
`
const ordinalMethod2 = `	}
	panic(fmt.Errorf("%%d: unknown %s", i))
}
`

func writeFuncOrdinal(w io.Writer, mainType string, values []string) error {
	_, err := fmt.Fprintf(w, ordinalMethod1, mainType, mainType)
	if err != nil {
		return err
	}

	for i, s := range values {
		_, err = fmt.Fprintf(w, "\tcase %s:\n\t\treturn %d\n", s, i)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, ordinalMethod2, mainType)
	return err
}

//-------------------------------------------------------------------------------------------------

const asMethod = `
// Parse parses a string to find the corresponding %s
func (v *%s) Parse(s string) error {
	var i0 uint16 = 0
	for j := 1; j < len(%s); j++ {
		i1 := %s[j]
		p := %s[i0:i1]
		if s == p {
			*v = %s(j - 1)
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised %s")
}

// As%s parses a string to find the corresponding %s
func As%s(s string) (%s, error) {
	var i = new(%s)
	err := i.Parse(s)
	return *i, err
}
`

func writeFuncAsEnum(w io.Writer, mainType, names, indexes string, values []string) error {
	_, err := fmt.Fprintf(w, asMethod, mainType, mainType, indexes, indexes, names, mainType, mainType, mainType, mainType, mainType, mainType, mainType)
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

func write(w io.Writer, mainType, baseType, plural, pkg string, values []string) error {

	lc := strings.ToLower(mainType)
	names := fmt.Sprintf("%sEnumStrings", lc)
	indexes := fmt.Sprintf("%sEnumIndex", lc)

	err := writeHead(w, pkg)
	if err != nil {
		return err
	}

	err = writeConst(w, names, values)
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

	err = writeFuncString(w, mainType, names, indexes)
	if err != nil {
		return err
	}

	err = writeFuncOrdinal(w, mainType, values)
	if err != nil {
		return err
	}

	err = writeFuncAsEnum(w, mainType, names, indexes, values)
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
