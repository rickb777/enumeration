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

func writeConst(w io.Writer, name string, values []string) error {
	_, err := fmt.Fprintf(w, "const %s = \"", name)
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

const stringMethod = `// String returns the string representation of a %s
func (i %s) String() string {
	if i < 0 || i >= %s(len(%s)-1) {
		return fmt.Sprintf("%s(%%d)", i)
	}
	return %s[%s[i]:%s[i+1]]
}

`

func writeFuncString(w io.Writer, mainType, name, index string) error {
	_, err := fmt.Fprintf(w, stringMethod, mainType, mainType, mainType, index, mainType, name, index, index)
	return err
}

//-------------------------------------------------------------------------------------------------

const ordinalMethod1 = `// Ordinal returns the ordinal number of a %s
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
		_, err = fmt.Fprintf(w, "\tcase %s: return %d\n", s, i)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, ordinalMethod2, mainType)
	return err
}

//-------------------------------------------------------------------------------------------------

const asMethod = `// As%s parses a string to find the corresponding %s
func As%s(s string) (%s, error) {
	i0 := 0
	for j := 1; j < len(%s); j++ {
		i1 := %s[j]
		p := %s[i0:i1]
		if s == p {
			return %s(j-1), nil
		}
		i0 = i1
	}
	return %s, errors.New(s + ": unrecognised %s")
}

`

func writeFuncAsEnum(w io.Writer, mainType, name, index string, values []string) error {
	_, err := fmt.Fprintf(w, asMethod, mainType, mainType, mainType, mainType, index, index, name, mainType, values[0], mainType)
	return err
}

//-------------------------------------------------------------------------------------------------

func write(w io.Writer, mainType, plural, pkg string, values []string) error {

	lc := strings.ToLower(mainType)
	name := fmt.Sprintf("%sEnumStrings", lc)
	index := fmt.Sprintf("%sEnumIndex", lc)

	err := writeHead(w, pkg)
	if err != nil {
		return err
	}

	err = writeConst(w, name, values)
	if err != nil {
		return err
	}

	err = writeIndexes(w, index, values)
	if err != nil {
		return err
	}

	err = writeFuncString(w, mainType, name, index)
	if err != nil {
		return err
	}

	err = writeFuncOrdinal(w, mainType, values)
	if err != nil {
		return err
	}

	err = writeFuncAsEnum(w, mainType, name, index, values)
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
