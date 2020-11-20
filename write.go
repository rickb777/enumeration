package main

import (
	"fmt"
	"io"
	"strings"
)

type printer struct {
	w   io.Writer
	err error
}

func (p *printer) Printf(message string, args ...interface{}) {
	if p.err == nil {
		_, p.err = fmt.Fprintf(p.w, message, args...)
	}
}

//-------------------------------------------------------------------------------------------------

const head = `// generated code - do not edit
// github.com/rickb777/enumeration %s

package %s

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/rickb777/enumeration/enum"
)
`

func (m model) writeHead(p *printer) {
	p.Printf(head, version, m.pkg)
}

//-------------------------------------------------------------------------------------------------

func (m model) writeConst(p *printer, names string) {
	p.Printf("\nconst %s = \"", names)

	for _, s := range m.values {
		for _, f := range m.xf {
			s = f.Func()(s)
		}
		p.Printf(s)
	}

	p.Printf("\"\n")
}

//-------------------------------------------------------------------------------------------------

func (m model) writeIndexes(p *printer, index string) {
	p.Printf("\nvar %s = [...]uint16{0", index)

	n := 0
	for _, s := range m.values {
		n += len(s)
		p.Printf(", %d", n)
	}

	p.Printf("}\n")
}

//-------------------------------------------------------------------------------------------------

func (m model) writeAllItemsSlice(p *printer, name, mainType string) {
	p.Printf("\n// All%s lists all %d values in order.\n", name, len(m.values))
	p.Printf("var All%s = []%s{", name, mainType)

	comma := ""
	for _, s := range m.values {
		p.Printf("%s%s", comma, s)
		comma = ", "
	}

	p.Printf("}\n")
}

//-------------------------------------------------------------------------------------------------

const stringMethod = `
// String returns the string representation of a %s.
func (i %s) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(All%s) {
		return fmt.Sprintf("%s(%s)", i)
	}
	return %s[%s[o]:%s[o+1]]
}
`

func (m model) writeFuncString(p *printer, names, indexes string) {
	placeholder := "%s"
	switch m.baseType {
	case "int", "uint",
		"int8", "uint8",
		"int16", "uint16",
		"int32", "uint32",
		"int64", "uint64":
		placeholder = "%d"
	case "float32", "float64":
		placeholder = "%g"
	}
	p.Printf(stringMethod, m.mainType, m.mainType, m.plural, m.mainType, placeholder, names, indexes, indexes)
}

//-------------------------------------------------------------------------------------------------

const ordinalMethodStart = `
// Ordinal returns the ordinal number of a %s.
func (i %s) Ordinal() int {
	switch i {
`
const ordinalMethodEnd = `	}
	return -1
}
`

func (m model) writeFuncOrdinal(p *printer) {
	p.Printf(ordinalMethodStart, m.mainType, m.mainType)

	for i, s := range m.values {
		p.Printf("\tcase %s:\n\t\treturn %d\n", s, i)
	}

	p.Printf(ordinalMethodEnd)
}

//-------------------------------------------------------------------------------------------------

const ofMethodStart = `
// %sOf returns a %s based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid %s is returned.
func %sOf(i int) %s {
	if 0 <= i && i < len(All%s) {
		return All%s[i]
	}
	// an invalid result
	return `

const ofMethodEnd = `
}
`

func (m model) writeFuncOf(p *printer) {
	p.Printf(ofMethodStart, m.mainType, m.mainType, m.mainType, m.mainType, m.mainType, m.plural, m.plural)

	for i, s := range m.values {
		if i > 0 {
			p.Printf(" + ")
		}
		p.Printf(s)
	}

	p.Printf(ofMethodEnd)
}

//-------------------------------------------------------------------------------------------------

const isValidMethodStart = `
// IsValid determines whether a %s is one of the defined constants.
func (i %s) IsValid() bool {
	switch i {
	case `

const isValidMethodEnd = `:
		return true
	}
	return false
}
`

func (m model) writeFuncIsValid(p *printer) {
	p.Printf(isValidMethodStart, m.mainType, m.mainType)

	nl := 5
	for i, s := range m.values {
		if i > 0 {
			p.Printf(",")
		}
		nl--
		if nl == 0 {
			p.Printf("\n\t\t")
			nl = 5
		} else if i > 0 {
			p.Printf(" ")
		}
		p.Printf("%s", s)

	}

	p.Printf(isValidMethodEnd)
}

//-------------------------------------------------------------------------------------------------

const parseMethodStart = `
// Parse parses a string to find the corresponding %s, accepting either one of the string
// values or an ordinal number.
`

const parseMethodEnd = `	ord, err := strconv.Atoi(s)
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

func (m model) writeFuncParse(p *printer, names, indexes string) {
	p.Printf(parseMethodStart, m.mainType)
	for _, f := range m.xf {
		if f.Info() != "" {
			p.Printf("// %s\n", f.Info())
		}
	}
	p.Printf("func (v *%s) Parse(s string) error {\n", m.mainType)
	for _, f := range m.xf {
		if f != NoChange {
			p.Printf("\ts = %s\n", f)
		}
	}
	p.Printf(parseMethodEnd, m.plural, m.plural, indexes, indexes, names, m.plural, m.mainType)
}

//-------------------------------------------------------------------------------------------------

const asMethodStart = `
// As%s parses a string to find the corresponding %s, accepting either one of the string
// values or an ordinal number.
`

const asMethodEnd = `func As%s(s string) (%s, error) {
	var i = new(%s)
	err := i.Parse(s)
	return *i, err
}
`

func (m model) writeFuncAs(p *printer) {
	p.Printf(asMethodStart, m.mainType, m.mainType)
	for _, f := range m.xf {
		if f.Info() != "" {
			p.Printf("// %s\n", f.Info())
		}
	}
	p.Printf(asMethodEnd, m.mainType, m.mainType, m.mainType)
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

func (m model) writeMarshalText(p *printer) {
	p.Printf(marshalText, m.mainType, m.mainType)
}

//-------------------------------------------------------------------------------------------------

const marshalJson = `
// %sMarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var %sMarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// %sMarshalJSONUsingString is true.
func (i %s) MarshalJSON() ([]byte, error) {
	if %sMarshalJSONUsingString {
		s := []byte(i.String())
		b := make([]byte, len(s)+2)
		b[0] = '"'
		copy(b[1:], s)
		b[len(s)+1] = '"'
		return b, nil
	}
	// else use the ordinal
	s := strconv.Itoa(i.Ordinal())
	return []byte(s), nil
}

// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (i *%s) UnmarshalJSON(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
`

func (m model) writeMarshalJson(p *printer) {
	p.Printf(marshalJson, m.mainType, m.mainType, m.mainType, m.mainType, m.mainType, m.mainType)
}

//-------------------------------------------------------------------------------------------------

func (m model) write(w io.Writer) error {
	lc := strings.ToLower(m.mainType)
	names := fmt.Sprintf("%sEnumStrings", lc)
	indexes := fmt.Sprintf("%sEnumIndex", lc)

	p := &printer{w: w}
	m.writeHead(p)
	m.writeConst(p, names)
	m.writeIndexes(p, indexes)
	m.writeAllItemsSlice(p, m.plural, m.mainType)
	m.writeAllItemsSlice(p, m.mainType+"Enums", "enum.Enum")
	m.writeFuncString(p, names, indexes)
	m.writeFuncOrdinal(p)
	m.writeFuncOf(p)
	m.writeFuncIsValid(p)
	m.writeFuncParse(p, names, indexes)
	m.writeFuncAs(p)
	m.writeMarshalText(p)
	m.writeMarshalJson(p)

	if p.err != nil {
		return p.err
	}

	if c, ok := w.(io.Closer); ok {
		return c.Close()
	}

	return nil
}
