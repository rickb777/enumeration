package main

import (
	"fmt"
	"go/types"
	"io"
	"strings"
	"text/template"
)

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

//-------------------------------------------------------------------------------------------------

const joinedStringAndIndexes = `
const <<.LcType>>EnumStrings = "<<.TransformedValues>>"

var <<.LcType>>EnumIndex = [...]uint16{<<.Indexes>>}
`

func (m model) TransformedValues() string {
	buf := &strings.Builder{}
	for _, s := range m.Values {
		for _, f := range m.XF {
			s = f.Fn(s)
		}
		fmt.Fprintf(buf, s)
	}
	return buf.String()
}

func (m model) Indexes() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "0")
	n := 0
	for _, s := range m.Values {
		n += len(s)
		fmt.Fprintf(buf, ", %d", n)
	}
	return buf.String()
}

//-------------------------------------------------------------------------------------------------

const allItems = `
// All<<.Plural>> lists all <<len .Values>> values in order.
var All<<.Plural>> = []<<.MainType>>{<<.ValuesJoined 0 ", ">>}

// All<<.MainType>>Enums lists all <<len .Values>> values in order.
var All<<.MainType>>Enums = <<.AllItemsSlice>>{<<.ValuesJoined 0 ", ">>}
`

func (m model) AllItemsSlice() string {
	switch m.BaseKind() {
	case types.Int:
		return "enum.IntEnums"
	case types.Float64:
		return "enum.FloatEnums"
	}
	panic("undefined")
}

//-------------------------------------------------------------------------------------------------

const stringMethod = `
// String returns the string representation of a <<.MainType>>.
func (i <<.MainType>>) String() string {
	o := i.Ordinal()
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", i)
	}
	return <<.LcType>>EnumStrings[<<.LcType>>EnumIndex[o]:<<.LcType>>EnumIndex[o+1]]
}
`

//-------------------------------------------------------------------------------------------------

const ordinalMethod = `
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
`

//-------------------------------------------------------------------------------------------------

const baseMethod = `
// <<.BaseApproxUC>> returns the <<.BaseApproxLC>> value. This is not necessarily the same as the ordinal.
// It serves to facilitate polymorphism (see enum.<<.BaseApproxUC>>Enum).
func (i <<.MainType>>) <<.BaseApproxUC>>() <<.BaseApproxLC>> {
	return <<.BaseApproxLC>>(i)
}
`

//-------------------------------------------------------------------------------------------------

const ofMethod = `
// <<.MainType>>Of returns a <<.MainType>> based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid <<.MainType>> is returned.
func <<.MainType>>Of(i int) <<.MainType>> {
	if 0 <= i && i < len(All<<.Plural>>) {
		return All<<.Plural>>[i]
	}
	// an invalid result
	return <<.ValuesJoined 0 " + ">>
}
`

//-------------------------------------------------------------------------------------------------

const isValidMethod = `
// IsValid determines whether a <<.MainType>> is one of the defined constants.
func (i <<.MainType>>) IsValid() bool {
	switch i {
	case <<.ValuesWithWrapping>>:
		return true
	}
	return false
}
`

func (m model) ValuesWithWrapping() string {
	buf := &strings.Builder{}
	nl := 5
	for i, s := range m.Values {
		if i > 0 {
			fmt.Fprintf(buf, ",")
		}
		nl--
		if nl == 0 {
			fmt.Fprintf(buf, "\n\t\t")
			nl = 5
		} else if i > 0 {
			fmt.Fprintf(buf, " ")
		}
		fmt.Fprintf(buf, "%s", s)
	}
	return buf.String()
}

//-------------------------------------------------------------------------------------------------

const parseMethod = `
// Parse parses a string to find the corresponding <<.MainType>>, accepting either one of the string
// values or an ordinal number.
<<- range .XF>><<if ne .Info "">>
// <<.Info>>
<<- end>>
<<- end>>
func (v *<<.MainType>>) Parse(s string) error {
<<- range .XF>><<if ne .Str "">>
	s = <<.Str>>
<<- end>>
<<- end>>
	ord, err := strconv.Atoi(s)
	if err == nil && 0 <= ord && ord < len(All<<.Plural>>) {
		*v = All<<.Plural>>[ord]
		return nil
	}
	var i0 uint16 = 0
	for j := 1; j < len(<<.LcType>>EnumIndex); j++ {
		i1 := <<.LcType>>EnumIndex[j]
		p := <<.LcType>>EnumStrings[i0:i1]
		if s == p {
			*v = All<<.Plural>>[j-1]
			return nil
		}
		i0 = i1
	}
	return errors.New(s + ": unrecognised <<.MainType>>")
}
`

//-------------------------------------------------------------------------------------------------

const asMethod = `
// As<<.MainType>> parses a string to find the corresponding <<.MainType>>, accepting either one of the string
// values or an ordinal number.
<<- range .XF>><<if ne .Info "">>
// <<.Info>>
<<- end>>
<<- end>>
func As<<.MainType>>(s string) (<<.MainType>>, error) {
	var i = new(<<.MainType>>)
	err := i.Parse(s)
	return *i, err
}
`

//-------------------------------------------------------------------------------------------------

const marshalText = `
// MarshalText converts values to a form suitable for transmission via JSON, XML etc.
func (i <<.MainType>>) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalText converts transmitted values to ordinary values.
func (i *<<.MainType>>) UnmarshalText(text []byte) error {
	return i.Parse(string(text))
}
`

//-------------------------------------------------------------------------------------------------

const marshalJSON = `
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

//-------------------------------------------------------------------------------------------------

func (m model) write(w io.Writer) error {
	err := m.execTemplate(w,
		head+
			joinedStringAndIndexes+
			allItems+
			stringMethod+
			ordinalMethod+
			baseMethod+
			ofMethod+
			isValidMethod+
			parseMethod+
			asMethod+
			marshalText+
			marshalJSON)

	if err != nil {
		return err
	}

	if c, ok := w.(io.Closer); ok {
		return c.Close()
	}

	return nil
}

func (m model) execTemplate(w io.Writer, tpl string) error {
	tmpl, err := template.New("t").Delims("<<", ">>").Parse(tpl)
	if err != nil {
		panic(err)
	}
	return tmpl.Execute(w, m)
}
