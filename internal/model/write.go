package model

import (
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"text/template"
)

const toStringMethod = `
func (v <<.MainType>>) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}
`

var vtoStringUnit = Unit{
	Declares: "v.toString",
	Template: toStringMethod,
}

func buildToStringMethod(units Units) {
	units.Add(vtoStringUnit)
}

func (m Model) writeToStringMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.toString", toStringMethod)
}

//-------------------------------------------------------------------------------------------------

const parseStringMethod = `
func (v *<<.MainType>>) parseString(s string, concats string, indexes []uint16) (ok bool) {
	var i0 uint16 = 0

	for j := 1; j < len(indexes); j++ {
		i1 := indexes[j]
		p := concats[i0:i1]
		if s == p {
			*v = All<<.Plural>>[j-1]
			return true
		}
		i0 = i1
	}
<<- if .AliasTable>>
	*v, ok = <<.AliasTable>>[s]
	return ok
<<- else>>
	return false
<<- end>>
}
`

var vparseStringUnit = Unit{
	Declares: "v.parseString",
	Template: parseStringMethod,
}

func buildParseStringMethod(units Units) {
	units.Add(vparseStringUnit)
}

func (m Model) writeParseStringMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.parseString", parseStringMethod)
}

//-------------------------------------------------------------------------------------------------

const stringMethod = `
// String returns the literal string representation of a <<.MainType>>, which is
// the same as the const identifier but without prefix or suffix.
func (v <<.MainType>>) String() string {
	o := v.Ordinal()
	return v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
}
`

var vStringUnit = Unit{
	Declares: "v.String",
	Requires: []string{"v.Ordinal", "v.toString"},
	Template: stringMethod,
}

func buildStringMethod(units Units) {
	units.Add(vStringUnit)
}

func (m Model) writeStringMethod(w DualWriter) {
	m.execTemplate(w, stringMethod)
	m.writeToStringMethod(w)
}

//-------------------------------------------------------------------------------------------------

const ordinalMethod = `
// Ordinal returns the ordinal number of a <<.MainType>>. This is an integer counting
// from zero. It is *not* the same as the const number assigned to the value.
func (v <<.MainType>>) Ordinal() int {
	switch v {
	<<- range $i, $v := .Values>>
	case <<$v.Identifier>>:
		return <<$i>>
	<<- end>>
	}
	return -1
}
`

var vOrdinalUnit = Unit{
	Declares: "v.Ordinal",
	Template: ordinalMethod,
}

func buildOrdinalMethod(units Units) {
	units.Add(vOrdinalUnit)
}

func (m Model) writeOrdinalMethod(w DualWriter) {
	m.execTemplate(w, ordinalMethod)
}

//-------------------------------------------------------------------------------------------------

const intMethod = `
// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v <<.MainType>>) Int() int {
	return int(v)
}
`

var vIntUnit = Unit{
	Declares: "v.Int",
	Template: intMethod,
}

const floatMethod = `
// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (v <<.MainType>>) Float() float64 {
	return float64(v)
}
`

var vFloatUnit = Unit{
	Declares: "v.Float",
	Template: floatMethod,
}

func buildBaseMethod(units Units, m Model) {
	if m.IsFloat() {
		units.Add(vFloatUnit)
	} else {
		units.Add(vIntUnit)
	}
}

func (m Model) writeBaseMethod(w DualWriter) {
	if m.IsFloat() {
		m.execTemplate(w, floatMethod)
	} else {
		m.execTemplate(w, intMethod)
	}
}

//-------------------------------------------------------------------------------------------------

const mainTypeOfFunction = `
// <<.MainType>>Of returns a <<.MainType>> based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid <<.MainType>> is returned.
func <<.MainType>>Of(v int) <<.MainType>> {
	if 0 <= v && v < len(All<<.Plural>>) {
		return All<<.Plural>>[v]
	}
	// an invalid result
	return <<.ValuesJoined 0 " + ">> + 1
}
`

var mainTypeOfUnit = Unit{
	Declares: "OfFunction",
	Template: mainTypeOfFunction,
}

func buildOfMethod(units Units) {
	units.Add(mainTypeOfUnit)
}

func (m Model) writeOfMethod(w DualWriter) {
	m.execTemplate(w, mainTypeOfFunction)
}

//-------------------------------------------------------------------------------------------------

const isValidMethod = `
// IsValid determines whether a <<.MainType>> is one of the defined constants.
func (v <<.MainType>>) IsValid() bool {
	return v.Ordinal() >= 0
}
`

var vIsValidUnit = Unit{
	Declares: "v.IsValid",
	Requires: []string{"v.Ordinal"},
	Template: isValidMethod,
}

func buildIsValidMethod(units Units) {
	units.Add(vIsValidUnit)
}

func (m Model) writeIsValidMethod(w DualWriter) {
	m.execTemplate(w, isValidMethod)
}

//-------------------------------------------------------------------------------------------------

const parse_body = `
<< if .Extra.Doc ->>
// Parse parses a string to find the corresponding <<.MainType>>, accepting one of the string values or
// a number. The input representation is determined by <<.MarshalTextRep>>. It is used by As<<.MainType>>.
<<- if .IgnoreCase>>
// The input case does not matter.
<<- end>>
//
// Usage Example
//
//    v := new(<<.MainType>>)
//    err := v.Parse(s)
//    ...  etc
//
<< end ->>
func (v *<<.MainType>>) <<.Extra.Method>>(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := <<.LcType>>TransformInput(in)
<<- if .Extra.Enum>>
<<- if .Asymmetric>>

	if v.parseString(s, <<.LcType>><<.Extra.Table>>Inputs, <<.LcType>><<.Extra.Table>>Index[:]) {
		return nil
	}
<<- else >>

	if v.parseString(s, <<.LcType>><<.Extra.Table>>Strings, <<.LcType>><<.Extra.Table>>Index[:]) {
		return nil
	}
<<- end >>
<<- end >>

	return v.parseFallback(in, s)
}
`

func buildParseHelperMethod(units Units, method, table string) {
	required := []string{"v.parseNumber", "parseFallback", "lctypeTransformInput"}
	if table != "Enum" {
		required = append(required, "v.parseString")
	}
	units.Add(
		Unit{
			Declares: "v." + method,
			Requires: required,
			Template: parse_body,
			Extra: map[string]any{
				"Method": method,
				"Table":  table,
				"Doc":    method == "Parse",
				"Enum":   table != "Enum",
			},
		},
	)
}

func (m Model) writeParseHelperMethod(w DualWriter, method, table string) {
	if !done.Contains("v." + method) {
		done.Add("v." + method)

		m.Extra["Method"] = method
		m.Extra["Table"] = table
		m.Extra["Doc"] = method == "Parse"

		if table != "Enum" {
			m.Extra["Enum"] = true
		}

		m.writeParseNumberMethod(w)
		m.execTemplate(w, parse_body)

		m.writeParseFallback(w)
		m.writeTransformInputFunction(w)
		m.writeParseStringMethod(w)
	}
}

func (m Model) writeParseMethod(w DualWriter) {
	m.writeParseHelperMethod(w, "Parse", "Enum")
}

//-------------------------------------------------------------------------------------------------

const parseFallbackMethod = `
func (v *<<.MainType>>) parseFallback(in, s string) error {
	<<- if .Asymmetric>>
	if v.parseString(s, <<.LcType>>EnumInputs, <<.LcType>>EnumIndex[:]) {
		return nil
	}
	<<- else >>
	if v.parseString(s, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]) {
		return nil
	}
	<<- end >>
	<<- if .AliasTable>>

	var ok bool
	*v, ok = <<.AliasTable>>[s]
	if ok {
		return nil
	}
	<<- end>>

	return errors.New(in + ": unrecognised <<.LcType>>")
}
`

var vparseFallbackUnit = Unit{
	Declares: "v.parseFallback",
	Requires: []string{"v.parseString"},
	Template: parseFallbackMethod,
}

func buildParseFallback(units Units) {
	units.Add(vparseFallbackUnit)
}

func (m Model) writeParseFallback(w DualWriter) {
	m.writeUnexportedFunc(w, "v.parseFallback", parseFallbackMethod)
}

//-------------------------------------------------------------------------------------------------

const parseNumberMethod = `
// parseNumber attempts to convert a decimal value.
// << if .Lenient >>Any number is allowed, even if the result is invalid.<< else ->>
Only numbers that correspond to the enumeration are valid.<< end >>
func (v *<<.MainType>>) parseNumber(s string) (ok bool) {
<<- if .IsFloat>>
	num, err := strconv.ParseFloat(s, 64)
<<- else>>
	num, err := strconv.ParseInt(s, 10, 64)
<<- end>>
	if err == nil {
		*v = <<.MainType>>(num)
		return << if .Lenient >>true<< else >>v.IsValid()<< end >>
	}
	return false
}
`

var vparseNumberUnit = Unit{
	Declares: "v.parseNumber",
	Template: parseNumberMethod,
}

func buildParseNumberMethod(units Units) {
	units.Add(vparseNumberUnit)
}

func (m Model) writeParseNumberMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.parseNumber", parseNumberMethod)
}

//-------------------------------------------------------------------------------------------------

const asFunction = `
// As<<.MainType>> parses a string to find the corresponding <<.MainType>>, accepting either one of the string values or
// a number. The input representation is determined by <<.LcType>>MarshalTextRep. It wraps Parse.
<<- if .IgnoreCase>>
// The input case does not matter.
<<- end>>
func As<<.MainType>>(s string) (<<.MainType>>, error) {
	var v = new(<<.MainType>>)
	err := v.Parse(s)
	return *v, err
}
`

var asFunctionUnit = Unit{
	Declares: "AsFunction",
	Requires: []string{"v.Parse"},
	Template: asFunction,
}

func buildAsMethod(units Units) {
	units.Add(asFunctionUnit)
}

func (m Model) writeAsMethod(w DualWriter) {
	m.execTemplate(w, asFunction)
}

//-------------------------------------------------------------------------------------------------

const mustParseFunction = `
// MustParse<<.MainType>> is similar to As<<.MainType>> except that it panics on error.
<<- if .IgnoreCase>>
// The input case does not matter.
<<- end>>
func MustParse<<.MainType>>(s string) <<.MainType>> {
	v, err := As<<.MainType>>(s)
	if err != nil {
		panic(err)
	}
	return v
}
`

var mustParseFunctionUnit = Unit{
	Declares: "MustParseFunction",
	Requires: []string{"AsFunction"},
	Template: mustParseFunction,
}

func buildMustParseMethod(units Units) {
	units.Add(mustParseFunctionUnit)
}

func (m Model) writeMustParseMethod(w DualWriter) {
	m.execTemplate(w, mustParseFunction)
}

//-------------------------------------------------------------------------------------------------

const marshalText_Main = `
// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v <<.MainType>>) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}
`

var vmarshalText_Main_Unit = Unit{
	Declares: "v.MarshalText",
	Requires: []string{"v.marshalText"},
	Template: marshalText_Main,
}

const marshalText_struct_tags = `
// Text returns the representation used for transmission via XML, JSON etc.
func (v <<.MainType>>) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to bytes suitable for transmission via XML, JSON etc.
// The representation is chosen according to 'text' struct tags.
func (v <<.MainType>>) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, <<.LcType>>TextStrings, <<.LcType>>TextIndex[:]), nil
}
`

var vmarshalText_struct_tags_Unit = Unit{
	Declares: "v.marshalText",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: marshalText_struct_tags,
}

const marshalText_identifier = `
// Text returns the representation used for transmission via XML, JSON etc.
func (v <<.MainType>>) Text() string {
	s, _ := v.marshalText()
	return s
}

// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The identifier representation is chosen according to -marshaltext.
func (v <<.MainType>>) marshalText() (string, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberStringOrError()
	}

	return v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]), nil
}
`

var vmarshalText_identifier_Unit = Unit{
	Declares: "v.marshalText",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: marshalText_identifier,
}

const marshalText_number = `
// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The number representation is chosen according to -marshaltext.
func (v <<.MainType>>) marshalText() (string, error) {
	if !v.IsValid() {
		return v.marshalNumberStringOrError()
	}

	return <<.LcType>>MarshalNumber(v), nil
}
`

var vmarshalText_number_Unit = Unit{
	Declares: "v.marshalText",
	Requires: []string{"v.IsValid", "v.marshalNumberStringOrError", "v.toString"},
	Template: marshalText_number,
}

func buildMarshalText(units Units, m Model) {
	if m.HasTextTags() {
		units.Add(vmarshalText_Main_Unit)
		units.Add(vmarshalText_struct_tags_Unit)
		buildMarshalNumberOrErrorMethod(units, m)
		buildToStringMethod(units)
		return
	}

	switch m.MarshalTextRep {
	case enum.Identifier:
		units.Add(vmarshalText_Main_Unit)
		units.Add(vmarshalText_identifier_Unit)
		buildMarshalNumberOrErrorMethod(units, m)
		buildToStringMethod(units)
	case enum.Number:
		units.Add(vmarshalText_Main_Unit)
		units.Add(vmarshalText_number_Unit)
		buildMarshalNumberVarFunc(units)
		buildMarshalNumberOrErrorMethod(units, m)
	}
}

func (m Model) writeMarshalText(w DualWriter) {
	if m.HasTextTags() {
		m.execTemplate(w, marshalText_Main)
		m.execTemplate(w, marshalText_struct_tags)
		m.writeMarshalNumberOrErrorMethod(w)
		m.writeToStringMethod(w)
		return
	}

	switch m.MarshalTextRep {
	case enum.Identifier:
		m.execTemplate(w, marshalText_Main)
		m.execTemplate(w, marshalText_identifier)
		m.writeMarshalNumberOrErrorMethod(w)
		m.writeToStringMethod(w)
	case enum.Number:
		m.execTemplate(w, marshalText_Main)
		m.execTemplate(w, marshalText_number)
		m.writeMarshalNumberVarFunc(w)
		m.writeMarshalNumberOrErrorMethod(w)
	}
}

//-------------------------------------------------------------------------------------------------

const json_struct_tags = `
// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v <<.MainType>>) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, <<.LcType>>JSONStrings, <<.LcType>>JSONIndex[:])
}
`

var vJSON_struct_tags_Unit = Unit{
	Declares: "v.JSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: json_struct_tags,
}

const marshalJSON_struct_tags = `
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The representation is chosen according to 'json' struct tags.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, <<.LcType>>JSONStrings, <<.LcType>>JSONIndex[:])
	return enum.QuotedString(s), nil
}
`

var vMarshalJSON_struct_tags_Unit = Unit{
	Declares: "v.MarshalJSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberOrError", "v.toString"},
	Template: marshalJSON_struct_tags,
}

const json_identifier = `
// JSON returns an approximation to the representation used for transmission via JSON.
// However, strings are not quoted.
func (v <<.MainType>>) JSON() string {
	o := v.Ordinal()
	if o < 0 {
		s, _ := v.marshalNumberStringOrError()
		return s
	}

	return v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
}
`

var vJSON_identifier_Unit = Unit{
	Declares: "v.JSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: json_identifier,
}

const marshalJSON_identifier = `
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The identifier representation is chosen according to -marshaljson.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	o := v.Ordinal()
	if o < 0 {
		return v.marshalNumberOrError()
	}

	s := v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
	return enum.QuotedString(s), nil
}
`

var vMarshalJSON_identifier_Unit = Unit{
	Declares: "v.MarshalJSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberOrError", "v.toString"},
	Template: marshalJSON_identifier,
}

const marshalJSON_number = `
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The number representation is chosen according to -marshaljson.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	s := <<.LcType>>MarshalNumber(v)
	return []byte(s), nil
}
`

var vMarshalJSON_number_Unit = Unit{
	Declares: "v.MarshalJSON",
	Requires: []string{"v.IsValid", "v.marshalNumberOrError", "lctypeMarshalNumber"},
	Template: marshalJSON_number,
}

func buildMarshalJSON(units Units, m Model) {
	if m.HasJSONTags() {
		units.Add(vJSON_struct_tags_Unit)
		units.Add(vMarshalJSON_struct_tags_Unit)
		buildMarshalNumberOrErrorMethod(units, m)

	} else {
		switch m.MarshalJSONRep {
		case enum.Identifier:
			units.Add(vJSON_identifier_Unit)
			units.Add(vMarshalJSON_identifier_Unit)
			buildMarshalNumberOrErrorMethod(units, m)
		case enum.Number:
			units.Add(vMarshalJSON_number_Unit)
			buildMarshalNumberOrErrorMethod(units, m)
			buildMarshalNumberVarFunc(units)
		}
	}
}

func (m Model) writeMarshalJSON(w DualWriter) {
	if m.HasJSONTags() {
		m.execTemplate(w, json_struct_tags)
		m.execTemplate(w, marshalJSON_struct_tags)
		m.writeMarshalNumberOrErrorMethod(w)

	} else {
		switch m.MarshalJSONRep {
		case enum.Identifier:
			m.execTemplate(w, json_identifier)
			m.execTemplate(w, marshalJSON_identifier)
			m.writeMarshalNumberOrErrorMethod(w)
		case enum.Number:
			m.execTemplate(w, marshalJSON_number)
			m.writeMarshalNumberOrErrorMethod(w)
			m.writeMarshalNumberVarFunc(w)
		}
	}
}

//-------------------------------------------------------------------------------------------------

const marshalNumberVarFunc = `
// <<.LcType>>MarshalNumber handles marshaling where a number is required or where
// the value is out of range.
// This function can be replaced with any bespoke function than matches signature.
var <<.LcType>>MarshalNumber = func(v <<.MainType>>) string {
<<- if .IsFloat>>
	return strconv.FormatFloat(float64(v), 'g', 7, 64)
<<- else>>
	return strconv.FormatInt(int64(v), 10)
<<- end>>
}
`

var lctypeMarshalNumberUnit = Unit{
	Declares: "lctypeMarshalNumber",
	Template: marshalNumberVarFunc,
}

func buildMarshalNumberVarFunc(units Units) {
	units.Add(lctypeMarshalNumberUnit)
}

func (m Model) writeMarshalNumberVarFunc(w DualWriter) {
	m.writeUnexportedFunc(w, "marshalNumberVarFunc", marshalNumberVarFunc)
}

//-------------------------------------------------------------------------------------------------

const marshalNumberStringOrError = `
func (v <<.MainType>>) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}
`

var vmarshalNumberStringOrErrorUnit = Unit{
	Declares: "v.marshalNumberStringOrError",
	Requires: []string{"v.marshalNumberOrError"},
	Template: marshalNumberStringOrError,
}

const marshalNumberOrError_strict = `
func (v <<.MainType>>) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}
`

var vmarshalNumberOrError_strict_Unit = Unit{
	Declares: "v.marshalNumberOrError",
	Requires: []string{"v.invalidError"},
	Template: marshalNumberOrError_strict,
}

const marshalNumberOrError_lenient = `
func (v <<.MainType>>) marshalNumberOrError() ([]byte, error) {
	// allow lenient marshaling
	return []byte(<<.LcType>>MarshalNumber(v)), nil
}
`

var vmarshalNumberOrError_lenient_Unit = Unit{
	Declares: "v.marshalNumberOrError",
	Requires: []string{"lctypeMarshalNumber"},
	Template: marshalNumberOrError_lenient,
}

func buildMarshalNumberOrErrorMethod(units Units, m Model) {
	units.Add(vmarshalNumberStringOrErrorUnit)
	if m.Lenient {
		units.Add(vmarshalNumberOrError_lenient_Unit)
		buildMarshalNumberVarFunc(units)
	} else {
		units.Add(vmarshalNumberOrError_strict_Unit)
		buildInvalidErrorMethod(units)
	}
}

func (m Model) writeMarshalNumberOrErrorMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.marshalNumberStringOrError", marshalNumberStringOrError)
	if m.Lenient {
		m.writeUnexportedFunc(w, "v.marshalNumberOrError", marshalNumberOrError_lenient)
		m.writeMarshalNumberVarFunc(w)
	} else {
		m.writeUnexportedFunc(w, "v.marshalNumberOrError", marshalNumberOrError_strict)
		m.writeInvalidErrorMethod(w)
	}
}

//-------------------------------------------------------------------------------------------------

const invalidError = `
func (v <<.MainType>>) invalidError() error {
<<- if .IsFloat>>
	return fmt.Errorf("%g is not a valid <<.LcType>>", v)
<<- else>>
	return fmt.Errorf("%d is not a valid <<.LcType>>", v)
<<- end>>
}
`

var vinvalidErrorUnit = Unit{
	Declares: "v.invalidError",
	Template: invalidError,
}

func buildInvalidErrorMethod(units Units) {
	units.Add(vinvalidErrorUnit)
}

func (m Model) writeInvalidErrorMethod(w DualWriter) {
	m.writeUnexportedFunc(w, "v.invalidError", invalidError)
}

//-------------------------------------------------------------------------------------------------

const errorIfInvalid = `
func (v <<.MainType>>) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}
`

var verrorIfInvalidUnit = Unit{
	Declares: "v.errorIfInvalid",
	Requires: []string{"v.IsValid", "v.invalidError"},
	Template: errorIfInvalid,
}

func buildErrorIfInvalid(units Units) {
	units.Add(verrorIfInvalidUnit)
	buildInvalidErrorMethod(units)
}

func (m Model) writeErrorIfInvalid(w DualWriter) {
	m.writeUnexportedFunc(w, "v.errorIfInvalid", errorIfInvalid)
	m.writeInvalidErrorMethod(w)
}

//-------------------------------------------------------------------------------------------------

const unmarshalText = `
// UnmarshalText converts transmitted values to ordinary values.
func (v *<<.MainType>>) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}
`

var vUnmarshalTextUnit = Unit{
	Declares: "v.UnmarshalText",
	Requires: []string{"v.unmarshalText"},
	Template: unmarshalText,
}

func buildUnmarshalText(units Units, m Model) {
	if m.MarshalTextRep > 0 || m.HasTextTags() {
		units.Add(vUnmarshalTextUnit)
		if m.HasTextTags() {
			buildParseHelperMethod(units, "unmarshalText", "Text")
		} else {
			buildParseHelperMethod(units, "unmarshalText", "Enum")
		}
	}
}

func (m Model) writeUnmarshalText(w DualWriter) {
	if m.MarshalTextRep > 0 || m.HasTextTags() {
		m.execTemplate(w, unmarshalText)
		if m.HasTextTags() {
			m.writeParseHelperMethod(w, "unmarshalText", "Text")
		} else {
			m.writeParseHelperMethod(w, "unmarshalText", "Enum")
		}
	}
}

//-------------------------------------------------------------------------------------------------

const transformFunction = `
// <<.LcType>>TransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var <<.LcType>>TransformInput = func(in string) string {
	return << transform "in" >>
}
`

var transformFunctionUnit = Unit{
	Declares: "lctypeTransformInput",
	Template: transformFunction,
}

func buildTransformInputFunction(units Units) {
	units.Add(transformFunctionUnit)
}

func (m Model) writeTransformInputFunction(w DualWriter) {
	m.writeUnexportedFunc(w, "xTransformInput", transformFunction)
}

//-------------------------------------------------------------------------------------------------

const unmarshalJSON_plain = `
// UnmarshalJSON converts transmitted JSON values to ordinary values. It allows both
// ordinals and strings to represent the values.
func (v *<<.MainType>>) UnmarshalJSON(text []byte) error {
	s := string(text)
	if s == "null" {
		// Ignore null, like in the main JSON package.
		return nil
	}
	s = strings.Trim(s, "\"")
	return v.unmarshalJSON(s)
}

func (v *<<.MainType>>) unmarshalJSON(in string) error {
	if v.parseNumber(in) {
		return nil
	}

	s := <<.LcType>>TransformInput(in)
<<- if .HasJSONTags>>
<<- if .Asymmetric>>

	if v.parseString(s, <<.LcType>>JSONInputs, <<.LcType>>JSONIndex[:]) {
		return nil
	}
<<- else >>

	if v.parseString(s, <<.LcType>>JSONStrings, <<.LcType>>JSONIndex[:]) {
		return nil
	}
<<- end >>
<<- end >>
<<- if .Asymmetric>>

	if v.parseString(s, <<.LcType>>EnumInputs, <<.LcType>>EnumIndex[:]) {
		return nil
	}
<<- else >>

	if v.parseString(s, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:]) {
		return nil
	}
<<- end >>
<<- if .AliasTable>>

	var ok bool
	*v, ok = <<.AliasTable>>[s]
	if ok {
		return nil
	}
<<- end>>

	return errors.New(in + ": unrecognised <<.LcType>>")
}
`

var vunmarshalJSON_plain_Unit = Unit{
	Declares: "v.UnmarshalJSON",
	Requires: []string{"v.parseNumber", "v.parseString", "lctypeTransformInput"},
	Template: unmarshalJSON_plain,
}

func buildUnmarshalJSON(units Units, m Model) {
	if m.MarshalJSONRep > 0 || m.HasJSONTags() {
		units.Add(vunmarshalJSON_plain_Unit)
		buildParseStringMethod(units)
	}
}

func (m Model) writeUnmarshalJSON(w DualWriter) {
	if m.MarshalJSONRep > 0 || m.HasJSONTags() {
		m.execTemplate(w, unmarshalJSON_plain)
		m.writeParseStringMethod(w)
	}
}

//-------------------------------------------------------------------------------------------------

const scan_all = `
// Scan parses some value, which can be a number, a string or []byte.
// It implements sql.Scanner, https://golang.org/pkg/database/sql/#Scanner
func (v *<<.MainType>>) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var s string
	switch x := value.(type) {
	case int64:
		*v = <<.MainType>>(x)
		return v.errorIfInvalid()
	case float64:
		*v = <<.MainType>>(x)
		return v.errorIfInvalid()
	case []byte:
		s = string(x)
	case string:
		s = x
	default:
		return fmt.Errorf("%T %+v is not a meaningful <<.LcType>>", value, value)
	}

	return v.scanParse(s)
}
`

var vScanUnit = Unit{
	Declares: "v.Scan",
	Requires: []string{"v.errorIfInvalid", "v.scanParse"},
	Template: scan_all,
}

func buildScanMethod(units Units, m Model) {
	if m.StoreRep > 0 || m.HasSQLTags() {
		units.Add(vScanUnit)
		if m.HasSQLTags() {
			buildParseHelperMethod(units, "scanParse", "SQL")
		} else {
			buildParseHelperMethod(units, "scanParse", "Enum")
		}
		buildErrorIfInvalid(units)
	}
}

func (m Model) writeScanMethod(w DualWriter) {
	if m.StoreRep > 0 || m.HasSQLTags() {
		m.execTemplate(w, scan_all)
		if m.HasSQLTags() {
			m.writeParseHelperMethod(w, "scanParse", "SQL")
		} else {
			m.writeParseHelperMethod(w, "scanParse", "Enum")
		}
		m.writeErrorIfInvalid(w)
	}
}

//-------------------------------------------------------------------------------------------------

const value_identifier = `
// Value converts the <<.MainType>> to a string  (based on '-store identifier').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.String(), nil
}
`

var vValue_identifier_Unit = Unit{
	Declares: "v.Value",
	Requires: []string{"v.IsValid", "v.String"},
	Template: value_identifier,
}

const value_number = `
// Value converts the <<.MainType>> to a number (based on '-store number').
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
<<- if .IsFloat>>
	return float64(v), nil
<<- else>>
	return int64(v), nil
<<- end>>
}
`

var vValue_number_Unit = Unit{
	Declares: "v.Value",
	Requires: []string{"v.IsValid", "v.String"},
	Template: value_number,
}

const value_struct_tags = `
// Value converts the <<.MainType>> to a string.
// The representation is chosen according to 'sql' struct tags.
// It implements driver.Valuer, https://golang.org/pkg/database/sql/driver/#Valuer
func (v <<.MainType>>) Value() (driver.Value, error) {
	o := v.Ordinal()
	if o < 0 {
		return nil, fmt.Errorf("%v: cannot be stored", v)
	}

	return v.toString(o, <<.LcType>>SQLStrings, <<.LcType>>SQLIndex[:]), nil
}
`

var vValue_struct_tags_Unit = Unit{
	Declares: "v.Value",
	Requires: []string{"v.Ordinal", "v.toString"},
	Template: value_struct_tags,
}

func buildValueMethod(units Units, m Model) {
	if m.HasSQLTags() {
		units.Add(vValue_struct_tags_Unit)
		return
	}

	switch m.StoreRep {
	case enum.Identifier:
		units.Add(vValue_identifier_Unit)
	case enum.Number:
		units.Add(vValue_number_Unit)
	}
}

func (m Model) writeValueMethod(w DualWriter) {
	if m.HasSQLTags() {
		m.execTemplate(w, value_struct_tags)
		return
	}

	switch m.StoreRep {
	case enum.Identifier:
		m.execTemplate(w, value_identifier)
	case enum.Number:
		m.execTemplate(w, value_number)
	}
}

//-------------------------------------------------------------------------------------------------

type Unit struct {
	Declares string
	Requires []string
	Template string
	Extra    map[string]any
}

type Units map[string]Unit

func (units Units) Add(unit Unit) Units {
	if _, exists := units[unit.Declares]; exists {
		panic(unit.Declares + " already exists")
	}
	units[unit.Declares] = unit
	return units
}

func (units Units) writeUnit(w DualWriter, m Model, identifier string, unit Unit) {
	if !done.Contains(identifier) {
		done.Add(identifier)
		m.execTemplate(w, unit.Template)

		for _, req := range unit.Requires {
			dep, ok := units[req]
			if !ok {
				panic(fmt.Sprintf("Missing dependency %s required by %s", req, identifier))
			}
			units.writeUnit(w, m, req, dep)
		}
	}
}

func (m Model) execTemplate(w DualWriter, tpl string) {
	tmpl, err := template.New("t").Funcs(m.FnMap()).Delims("<<", ">>").Parse(tpl)
	checkErr(err)
	checkErr(tmpl.Execute(w, m))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
