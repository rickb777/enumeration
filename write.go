package main

import (
	"io"
	"fmt"
	"strings"
)

const head = `// generated code - do not edit

package %s

import "fmt"

const %s = "`

const stringMethod = `
func (i %s) String() string {
	if i < 0 || i >= %s(len(%s)-1) {
		return fmt.Sprintf("%s(%%d)", i)
	}
	return %s[%s[i]:%s[i+1]]
}
`
const asMethod = `
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

func write(w io.Writer, mainType, pkg string, values []string) error {

	lc := strings.ToLower(mainType)
	name := fmt.Sprintf("%sEnumStrings", lc)
	index := fmt.Sprintf("%sEnumIndex", lc)

	_, err := fmt.Fprintf(w, head, pkg, name)
	if err != nil {
		return err
	}

	for _, s := range values {
		_, err = fmt.Fprintf(w, s)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, "\"\n\nvar %s = [...]uint16{0", index)
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

	_, err = fmt.Fprintf(w, "}\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, stringMethod, mainType, mainType, index, mainType, name, index, index)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, asMethod, mainType, mainType, 	index, index, name, mainType, values[0], mainType)
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

