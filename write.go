package main

import (
	"fmt"
	"go/types"
	"io"
	"strings"
	"text/template"
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
// github.com/rickb777/enumeration <<.Version>>

package <<.Pkg>>

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/rickb777/enumeration/enum"
)
`

func (m model) writeHead(p *printer) {
	m.execTemplate(p, head)
}

//-------------------------------------------------------------------------------------------------

func (m model) writeConst(p *printer, names string) {
	p.Printf("\nconst %s = \"", names)

	for _, s := range m.Values {
		for _, f := range m.XF {
			s = f.Fn(s)
		}
		p.Printf(s)
	}

	p.Printf("\"\n")
}

//-------------------------------------------------------------------------------------------------

func (m model) writeIndexes(p *printer, index string) {
	p.Printf("\nvar %s = [...]uint16{0", index)

	n := 0
	for _, s := range m.Values {
		n += len(s)
		p.Printf(", %d", n)
	}

	p.Printf("}\n")
}

//-------------------------------------------------------------------------------------------------

const allItems = `
// All<<.S1>> lists all <<len .Values>> values in order.
var All<<.S1>> = <<.S2>>{<<.ValuesJoined>>}
`

func (m model) doWriteAllItemsSlice(p *printer, name, enumsType string) {
	m.S1 = name
	m.S2 = enumsType
	m.execTemplate(p, allItems)
}

func (m model) writeAllItemsSlice(p *printer) {
	m.doWriteAllItemsSlice(p, m.Plural, "[]"+m.MainType)

	enumsType := "enum.Enums"
	switch m.BaseKind() {
	case types.Int:
		enumsType = "enum.IntEnums"
	case types.Float64:
		enumsType = "enum.FloatEnums"
	}

	m.doWriteAllItemsSlice(p, m.MainType+"Enums", enumsType)
}

//-------------------------------------------------------------------------------------------------

const methods = `
// String returns the string representation of a <<.MainType>>.
func (i <<.MainType>>) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", i)
	}
	return <<.LcType>>EnumStrings[<<.LcType>>EnumIndex[o]:<<.LcType>>EnumIndex[o+1]]
}

// Ordinal returns the ordinal number of a <<.MainType>>.
func (i <<.MainType>>) Ordinal() int {
	switch i {
	<<- range $i, $v := .Values>>
	case <<$v>>:
		return <<$i>>
	<<- end>>
	}
	return -1
}

// <<.BaseApproxUC>> returns the <<.BaseApproxLC>> value. This is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.<<.BaseApproxUC>>Enum).
func (i <<.MainType>>) <<.BaseApproxUC>>() <<.BaseApproxLC>> {
	return <<.BaseApproxLC>>(i)
}
`

func (m model) writeFuncValue(p *printer) {
	m.execTemplate(p, methods)
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
	p.Printf(ofMethodStart, m.MainType, m.MainType, m.MainType, m.MainType, m.MainType, m.Plural, m.Plural)

	for i, s := range m.Values {
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
	p.Printf(isValidMethodStart, m.MainType, m.MainType)

	nl := 5
	for i, s := range m.Values {
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
	p.Printf(parseMethodStart, m.MainType)
	for _, f := range m.XF {
		if f.Info != "" {
			p.Printf("// %s\n", f.Info)
		}
	}
	p.Printf("func (v *%s) Parse(s string) error {\n", m.MainType)
	for _, f := range m.XF {
		if f.Str != "" {
			p.Printf("\ts = %s\n", f.Str)
		}
	}
	p.Printf(parseMethodEnd, m.Plural, m.Plural, indexes, indexes, names, m.Plural, m.MainType)
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
	p.Printf(asMethodStart, m.MainType, m.MainType)
	for _, f := range m.XF {
		if f.Info != "" {
			p.Printf("// %s\n", f.Info)
		}
	}
	p.Printf(asMethodEnd, m.MainType, m.MainType, m.MainType)
}

//-------------------------------------------------------------------------------------------------

const marshalJson = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i <<.MainType>>) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *<<.MainType>>) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}

// <<.MainType>>MarshalJSONUsingString controls whether generated JSON uses ordinals or strings. By default,
// it is false and ordinals are used. Set it true to cause quoted strings to be used instead,
// these being easier to read but taking more resources.
var <<.MainType>>MarshalJSONUsingString = false

// MarshalJSON converts values to bytes suitable for transmission via JSON. By default, the
// ordinal integer is emitted, but a quoted string is emitted instead if
// <<.MainType>>MarshalJSONUsingString is true.
func (i <<.MainType>>) MarshalJSON() ([]byte, error) {
	if <<.MainType>>MarshalJSONUsingString {
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
func (i *<<.MainType>>) UnmarshalJSON(text []byte) error {
	if len(text) >= 2 && text[0] == '"' && text[len(text)-1] == '"' {
		s := string(text[1:len(text)-1])
		return i.Parse(s)
	}

	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	s := strings.Trim(string(text), "\"")
	return i.Parse(s)
}
`

func (m model) writeMarshal(p *printer) {
	m.execTemplate(p, marshalJson)
}

//-------------------------------------------------------------------------------------------------

func (m model) write(w io.Writer) error {
	lc := strings.ToLower(m.MainType)
	names := fmt.Sprintf("%sEnumStrings", lc)
	indexes := fmt.Sprintf("%sEnumIndex", lc)

	p := &printer{w: w}
	m.writeHead(p)
	m.writeConst(p, names)
	m.writeIndexes(p, indexes)
	m.writeAllItemsSlice(p)
	m.writeFuncValue(p)
	m.writeFuncOf(p)
	m.writeFuncIsValid(p)
	m.writeFuncParse(p, names, indexes)
	m.writeFuncAs(p)
	m.writeMarshal(p)

	if p.err != nil {
		return p.err
	}

	if c, ok := w.(io.Closer); ok {
		return c.Close()
	}

	return nil
}

func (m model) execTemplate(p *printer, tpl string) {
	tmpl, err := template.New("t").Delims("<<", ">>").Parse(tpl)
	if err != nil {
		panic(err)
	}
	p.err = tmpl.Execute(p.w, m)
}
