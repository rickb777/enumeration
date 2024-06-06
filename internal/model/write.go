package model

import (
	"fmt"
	"github.com/rickb777/enumeration/v3/enum"
	"github.com/rickb777/enumeration/v3/internal/codegen"
	"github.com/rickb777/enumeration/v3/internal/collection"
	"text/template"
)

var vtoStringUnit = codegen.Unit{
	Declares: "v.toString",
	Template: `
func (v <<.MainType>>) toString(o int, concats string, indexes []uint16) string {
	if o < 0 || o >= len(All<<.Plural>>) {
		return fmt.Sprintf("<<.MainType>>(<<.Placeholder>>)", v)
	}
	return concats[indexes[o]:indexes[o+1]]
}
`,
}

func buildToStringMethod(units *codegen.Units) {
	units.Add(vtoStringUnit)
}

//-------------------------------------------------------------------------------------------------

var vparseStringUnit = codegen.Unit{
	Declares: "v.parseString",
	Template: `
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
`,
}

func buildParseStringMethod(units *codegen.Units) {
	units.Add(vparseStringUnit)
}

//-------------------------------------------------------------------------------------------------

var vOrdinalUnit = codegen.Unit{
	Declares: "v.Ordinal",
	Template: `
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
`,
}

func buildOrdinalMethod(units *codegen.Units) {
	units.Add(vOrdinalUnit)
}

//-------------------------------------------------------------------------------------------------

var vStringUnit = codegen.Unit{
	Declares: "v.String",
	Requires: []string{"v.Ordinal", "v.toString"},
	Template: `
// String returns the literal string representation of a <<.MainType>>, which is
// the same as the const identifier but without prefix or suffix.
func (v <<.MainType>>) String() string {
	o := v.Ordinal()
	return v.toString(o, <<.LcType>>EnumStrings, <<.LcType>>EnumIndex[:])
}
`,
}

func buildStringMethod(units *codegen.Units) {
	units.Add(vStringUnit)
	buildToStringMethod(units)
	buildOrdinalMethod(units)
}

//-------------------------------------------------------------------------------------------------

var vIntUnit = codegen.Unit{
	Declares: "v.Number",
	Template: `
// Int returns the int value, which is not necessarily the same as the ordinal.
// This facilitates polymorphism (see enum.IntEnum).
func (v <<.MainType>>) Int() int {
	return int(v)
}
`,
}

var vFloatUnit = codegen.Unit{
	Declares: "v.Number",
	Template: `
// Float returns the float64 value. It serves to facilitate polymorphism (see enum.FloatEnum).
func (v <<.MainType>>) Float() float64 {
	return float64(v)
}
`,
}

func buildNumberMethod(units *codegen.Units, m Model) {
	if m.IsFloat() {
		units.Add(vFloatUnit)
	} else {
		units.Add(vIntUnit)
	}
}

//-------------------------------------------------------------------------------------------------

var mainTypeOfUnit = codegen.Unit{
	Declares: "OfFunction",
	Imports:  collection.NewStringSet("slices"),
	Template: `
var invalid<<.MainType>>Value = func() <<.MainType>> {
	var v <<.MainType>>
	for {
		if !slices.Contains(All<<.Plural>>, v) {
			return v
		}
		v++
	} // All<<.Plural>> is a finite set so loop will terminate eventually
}()

// <<.MainType>>Of returns a <<.MainType>> based on an ordinal number. This is the inverse of Ordinal.
// If the ordinal is out of range, an invalid <<.MainType>> is returned.
func <<.MainType>>Of(v int) <<.MainType>> {
	if 0 <= v && v < len(All<<.Plural>>) {
		return All<<.Plural>>[v]
	}
	return invalid<<.MainType>>Value
}
`,
}

func buildOfMethod(units *codegen.Units) {
	units.Add(mainTypeOfUnit)
}

//-------------------------------------------------------------------------------------------------

var vIsValidUnit = codegen.Unit{
	Declares: "v.IsValid",
	Requires: []string{"v.Ordinal"},
	Template: `
// IsValid determines whether a <<.MainType>> is one of the defined constants.
func (v <<.MainType>>) IsValid() bool {
	return v.Ordinal() >= 0
}
`,
}

func buildIsValidMethod(units *codegen.Units) {
	units.Add(vIsValidUnit)
	units.Add(vOrdinalUnit)
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

func buildParseHelperMethod(units *codegen.Units, method, table string) {
	required := []string{"v.parseNumber", "v.parseFallback", "lctypeTransformInput"}
	buildParseFallback(units)
	units.Add(vparseNumberUnit)
	units.Add(lctypeTransformInputUnit)

	if table != "Enum" {
		required = append(required, "v.parseString")
		units.Add(vparseStringUnit)
	}

	units.Add(
		codegen.Unit{
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

//-------------------------------------------------------------------------------------------------

var vparseFallbackUnit = codegen.Unit{
	Declares: "v.parseFallback",
	Requires: []string{"v.parseString"},
	Template: `
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
`,
}

func buildParseFallback(units *codegen.Units) {
	units.Add(vparseFallbackUnit)
	units.Add(vparseStringUnit)
}

//-------------------------------------------------------------------------------------------------

var vparseNumberUnit = codegen.Unit{
	Declares: "v.parseNumber",
	Imports:  collection.NewStringSet("strconv"),
	Template: `
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
`,
}

//func buildParseNumberMethod(codegen *codegen.Units) {
//	codegen.Add(vparseNumberUnit)
//}

//-------------------------------------------------------------------------------------------------

var asFunctionUnit = codegen.Unit{
	Declares: "AsFunction",
	Requires: []string{"v.Parse"},
	Template: `
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
`,
}

func buildAsMethod(units *codegen.Units) {
	units.Add(asFunctionUnit)
}

//-------------------------------------------------------------------------------------------------

var mustParseFunctionUnit = codegen.Unit{
	Declares: "MustParseFunction",
	Requires: []string{"AsFunction"},
	Template: `
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
`,
}

func buildMustParseMethod(units *codegen.Units) {
	units.Add(mustParseFunctionUnit)
	units.Add(asFunctionUnit)
}

//-------------------------------------------------------------------------------------------------

var vmarshalText_Main_Unit = codegen.Unit{
	Declares: "v.MarshalText",
	Requires: []string{"v.marshalText"},
	Template: `
// MarshalText converts values to bytes suitable for transmission via XML, JSON etc.
func (v <<.MainType>>) MarshalText() ([]byte, error) {
	s, err := v.marshalText()
	return []byte(s), err
}
`,
}

var vmarshalText_struct_tags_Unit = codegen.Unit{
	Declares: "v.marshalText",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: `
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
`,
}

var vmarshalText_identifier_Unit = codegen.Unit{
	Declares: "v.marshalText",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: `
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
`,
}

var vmarshalText_number_Unit = codegen.Unit{
	Declares: "v.marshalText",
	Requires: []string{"v.IsValid", "v.marshalNumberStringOrError", "v.toString"},
	Template: `
// marshalText converts values to a form suitable for transmission via XML, JSON etc.
// The number representation is chosen according to -marshaltext.
func (v <<.MainType>>) marshalText() (string, error) {
	if !v.IsValid() {
		return v.marshalNumberStringOrError()
	}

	return <<.LcType>>MarshalNumber(v), nil
}
`,
}

func buildMarshalText(units *codegen.Units, m Model) {
	buildToStringMethod(units)

	if m.HasTextTags() {
		units.Add(vmarshalText_Main_Unit)
		units.Add(vmarshalText_struct_tags_Unit)
		buildMarshalNumberOrErrorMethod(units, m)
		return
	}

	switch m.MarshalTextRep {
	case enum.Identifier:
		units.Add(vmarshalText_Main_Unit)
		units.Add(vmarshalText_identifier_Unit)
		buildMarshalNumberOrErrorMethod(units, m)
	case enum.Number:
		units.Add(vmarshalText_Main_Unit)
		units.Add(vmarshalText_number_Unit)
		buildMarshalNumberVarFunc(units)
		buildMarshalNumberOrErrorMethod(units, m)
	}
}

//-------------------------------------------------------------------------------------------------

var vJSON_struct_tags_Unit = codegen.Unit{
	Declares: "v.JSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: `
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
`,
}

const enumImportPath = "github.com/rickb777/enumeration/v3/enum"

var vMarshalJSON_struct_tags_Unit = codegen.Unit{
	Declares: "v.MarshalJSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberOrError", "v.toString"},
	Imports:  collection.NewStringSet(enumImportPath),
	Template: `
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
`,
}

var vJSON_identifier_Unit = codegen.Unit{
	Declares: "v.JSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberStringOrError", "v.toString"},
	Template: `
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
`,
}

var vMarshalJSON_identifier_Unit = codegen.Unit{
	Declares: "v.MarshalJSON",
	Requires: []string{"v.Ordinal", "v.marshalNumberOrError", "v.toString"},
	Imports:  collection.NewStringSet(enumImportPath),
	Template: `
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
`,
}

var vMarshalJSON_number_Unit = codegen.Unit{
	Declares: "v.MarshalJSON",
	Requires: []string{"v.IsValid", "v.marshalNumberOrError", "lctypeMarshalNumber"},
	Template: `
// MarshalJSON converts values to bytes suitable for transmission via JSON.
// The number representation is chosen according to -marshaljson.
func (v <<.MainType>>) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return v.marshalNumberOrError()
	}

	s := <<.LcType>>MarshalNumber(v)
	return []byte(s), nil
}
`,
}

func buildMarshalJSON(units *codegen.Units, m Model) {
	buildToStringMethod(units)

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

//-------------------------------------------------------------------------------------------------

var lctypeMarshalNumberUnit = codegen.Unit{
	Declares: "lctypeMarshalNumber",
	Imports:  collection.NewStringSet("strconv"),
	Template: `
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
`,
}

func buildMarshalNumberVarFunc(units *codegen.Units) {
	units.Add(lctypeMarshalNumberUnit)
}

//-------------------------------------------------------------------------------------------------

var vmarshalNumberStringOrErrorUnit = codegen.Unit{
	Declares: "v.marshalNumberStringOrError",
	Requires: []string{"v.marshalNumberOrError"},
	Template: `
func (v <<.MainType>>) marshalNumberStringOrError() (string, error) {
	bs, err := v.marshalNumberOrError()
	return string(bs), err
}
`,
}

var vmarshalNumberOrError_strict_Unit = codegen.Unit{
	Declares: "v.marshalNumberOrError",
	Requires: []string{"v.invalidError"},
	Template: `
func (v <<.MainType>>) marshalNumberOrError() ([]byte, error) {
	// disallow lenient marshaling
	return nil, v.invalidError()
}
`,
}

var vmarshalNumberOrError_lenient_Unit = codegen.Unit{
	Declares: "v.marshalNumberOrError",
	Requires: []string{"lctypeMarshalNumber"},
	Template: `
func (v <<.MainType>>) marshalNumberOrError() ([]byte, error) {
	// allow lenient marshaling
	return []byte(<<.LcType>>MarshalNumber(v)), nil
}
`,
}

func buildMarshalNumberOrErrorMethod(units *codegen.Units, m Model) {
	units.Add(vmarshalNumberStringOrErrorUnit)
	if m.Lenient {
		units.Add(vmarshalNumberOrError_lenient_Unit)
		buildMarshalNumberVarFunc(units)
	} else {
		units.Add(vmarshalNumberOrError_strict_Unit)
		buildInvalidErrorMethod(units)
	}
}

//-------------------------------------------------------------------------------------------------

var vinvalidErrorUnit = codegen.Unit{
	Declares: "v.invalidError",
	Template: `
func (v <<.MainType>>) invalidError() error {
<<- if .IsFloat>>
	return fmt.Errorf("%g is not a valid <<.LcType>>", v)
<<- else>>
	return fmt.Errorf("%d is not a valid <<.LcType>>", v)
<<- end>>
}
`,
}

func buildInvalidErrorMethod(units *codegen.Units) {
	units.Add(vinvalidErrorUnit)
}

//-------------------------------------------------------------------------------------------------

var verrorIfInvalidUnit = codegen.Unit{
	Declares: "v.errorIfInvalid",
	Requires: []string{"v.IsValid", "v.invalidError"},
	Template: `
func (v <<.MainType>>) errorIfInvalid() error {
	if v.IsValid() {
		return nil
	}
	return v.invalidError()
}
`,
}

func buildErrorIfInvalid(units *codegen.Units) {
	units.Add(verrorIfInvalidUnit)
	buildInvalidErrorMethod(units)
}

//-------------------------------------------------------------------------------------------------

var vUnmarshalTextUnit = codegen.Unit{
	Declares: "v.UnmarshalText",
	Requires: []string{"v.unmarshalText"},
	Template: `
// UnmarshalText converts transmitted values to ordinary values.
func (v *<<.MainType>>) UnmarshalText(bs []byte) error {
	return v.unmarshalText(string(bs))
}
`,
}

func buildUnmarshalText(units *codegen.Units, m Model) {
	if m.MarshalTextRep > 0 || m.HasTextTags() {
		units.Add(vUnmarshalTextUnit)
		if m.HasTextTags() {
			buildParseHelperMethod(units, "unmarshalText", "Text")
		} else {
			buildParseHelperMethod(units, "unmarshalText", "Enum")
		}
	}
}

//-------------------------------------------------------------------------------------------------

var lctypeTransformInputUnit = codegen.Unit{
	Declares: "lctypeTransformInput",
	Template: `
// <<.LcType>>TransformInput may alter input strings before they are parsed.
// This function is pluggable and is initialised using command-line flags
// -ic -lc -uc -unsnake.
var <<.LcType>>TransformInput = func(in string) string {
	return << transform "in" >>
}
`,
}

//func buildTransformInputFunction(codegen *codegen.Units) {
//	codegen.Add(lctypeTransformInputUnit)
//}

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

var vunmarshalJSON_plain_Unit = codegen.Unit{
	Declares: "v.UnmarshalJSON",
	Requires: []string{"v.parseNumber", "v.parseString", "lctypeTransformInput"},
	Template: unmarshalJSON_plain,
	Imports:  collection.NewStringSet("strings", "errors"),
}

func buildUnmarshalJSON(units *codegen.Units, m Model) {
	if m.MarshalJSONRep > 0 || m.HasJSONTags() {
		units.Add(vunmarshalJSON_plain_Unit)
		units.Add(vparseNumberUnit)
		units.Add(vparseStringUnit)
		units.Add(lctypeMarshalNumberUnit)
		units.Add(lctypeTransformInputUnit)
		buildParseStringMethod(units)
	}
}

//-------------------------------------------------------------------------------------------------

const awsDynamoDBImportPath = "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

var vMarshalDynamoDBAttributeValueUnit = codegen.Unit{
	Declares: "v.JSON",
	Requires: []string{"v.toString"},
	Imports:  collection.NewStringSet(awsDynamoDBImportPath),
	Template: `
// MarshalDynamoDBAttributeValue handles writing dates as DynamoDB attributes.
func (v <<.MainType>>) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: v.String(),
	}, nil
}`,
}

var vUnmarshalDynamoDBAttributeValueUnit = codegen.Unit{
	Declares: "v.JSON",
	Requires: []string{"v.Parse"},
	Imports:  collection.NewStringSet(awsDynamoDBImportPath),
	Template: `
// UnmarshalDynamoDBAttributeValue handles reading dates from DynamoDB attributes.
func (v <<.MainType>>) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) (err error) {
	avS, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return fmt.Errorf("wrong data type %T for a date; expecting types.AttributeValueMemberS", av)
	}

	*v, err = Parse<<.MainType>>(avS.Value)
	return err
}`,
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

var vScanUnit = codegen.Unit{
	Declares: "v.Scan",
	Requires: []string{"v.errorIfInvalid", "v.scanParse"},
	Template: scan_all,
}

func buildScanMethod(units *codegen.Units, m Model) {
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

var vValue_identifier_Unit = codegen.Unit{
	Declares: "v.Value",
	Requires: []string{"v.IsValid", "v.String"},
	Template: value_identifier,
	Imports:  collection.NewStringSet("database/sql/driver"),
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

var vValue_number_Unit = codegen.Unit{
	Declares: "v.Value",
	Requires: []string{"v.IsValid", "v.String"},
	Template: value_number,
	Imports:  collection.NewStringSet("database/sql/driver"),
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

var vValue_struct_tags_Unit = codegen.Unit{
	Declares: "v.Value",
	Requires: []string{"v.Ordinal", "v.toString"},
	Template: value_struct_tags,
	Imports:  collection.NewStringSet("database/sql/driver"),
}

func buildValueMethod(units *codegen.Units, m Model) {
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

//-------------------------------------------------------------------------------------------------

func writeUnit(w DualWriter, units *codegen.Units, done collection.StringSet, m Model, identifier, parent string) {
	if !done.Contains(identifier) {
		unit, ok := units.Take(identifier)
		if !ok {
			panic(fmt.Sprintf("Missing dependency %s required by %s", identifier, parent))
		}
		done.Add(identifier)
		m.Extra = unit.Extra
		m.execTemplate(w, unit.Template)

		for _, req := range unit.Requires {
			writeUnit(w, units, done, m, req, identifier)
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
