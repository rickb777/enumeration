package parse

import (
	"fmt"
	"github.com/rickb777/enumeration/v2/internal/model"
	"github.com/rickb777/enumeration/v2/internal/transform"
	"github.com/rickb777/enumeration/v2/internal/util"
	"go/token"
	"io"
	"regexp"
	"strings"
)

var UsingTable string
var AliasTable string
var MainType string
var fset *token.FileSet

var tagRE = regexp.MustCompile(`[a-z]:"`)

// https://go.dev/doc/go1.17_spec#Type_declarations (without type parameters)
// TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
// TypeSpec = AliasDecl | TypeDef .
//
// TypeDef  = identifier Type .
//
// Type     = TypeName | TypeLit | "(" Type ")" .
// TypeName = identifier | QualifiedIdent .
// TypeLit  = ArrayType | StructType | PointerType | FunctionType | InterfaceType | SliceType | MapType | ChannelType .

func Convert(in io.Reader, input string, xCase transform.Case, config model.Config) (model.Model, error) {
	src, err := io.ReadAll(in)
	if err != nil {
		return model.Model{}, err
	}

	MainType = config.MainType

	m := model.Model{
		Config:     config,
		LcType:     strings.ToLower(config.MainType),
		BaseType:   "int",
		Version:    util.Version,
		Case:       xCase,
		TagTable:   UsingTable,
		AliasTable: AliasTable,
		Extra:      make(map[string]string),
	}

	s := newFileScanner(input, src)

	var foundMainType = false
	var constItems []constItem

	for s.Scan() != token.EOF {
		switch s.Tok {
		case token.TYPE:
			s.Scan()
			if s.Tok == token.IDENT && s.Lit == MainType {
				foundMainType = true

				s.Scan()
				if s.Tok == token.IDENT {
					m.BaseType = s.Lit
					util.Debug("type %s %s\n", MainType, m.BaseType)
				}
			}

		case token.CONST:
			constItems = parseConst(s, constItems)
		}
	}

	debugConstItems(constItems)

	m.Values, _ = filterExportedItems(config.MainType, constItems)

	debugValues(m.Values)

	if s.gs.ErrorCount > 0 {
		return model.Model{}, fmt.Errorf("Syntax error in %s", input)
	}

	if !foundMainType {
		return model.Model{}, fmt.Errorf("Failed to find %s in %s", config.MainType, input)
	}

	if len(m.Values) == 0 {
		return model.Model{}, fmt.Errorf("Failed to find any values for %s in %s", config.MainType, input)
	}

	if e2 := m.CheckBadPrefixSuffix(); e2 != nil {
		return model.Model{}, e2
	}

	if e2 := m.CheckBadTags(); e2 != nil {
		return model.Model{}, e2
	}

	return m, nil
}

func filterExportedItems(mainType string, ids []constItem) (exported model.Values, defaultValue string) {
	var currentType string
	exported = make(model.Values, 0, len(ids))

	for _, v := range ids {
		if v.typ == mainType {
			if token.IsExported(v.id) {
				exported = exported.Append(v.id, v.tag)
				switch v.expression {
				case "0", "iota":
					defaultValue = v.id
				}
			}

		} else if v.typ == "" && v.expression == "" && currentType == mainType {
			if token.IsExported(v.id) {
				exported = exported.Append(v.id, v.tag)
			}
		}

		if v.typ != "" {
			currentType = v.typ
		}
	}

	return exported, defaultValue
}

func debugConstItems(ids []constItem) {
	if util.Dbg {
		util.Debug("\n   %-25s %-25s %-25s %s\n", "id", "type", "expression", "tag")
		for i, v := range ids {
			util.Debug("%-2d %-25s %-25s %-25s %s\n", i, v.id, v.typ, v.expression, v.tag)
		}
	}
}

func debugValues(values model.Values) {
	if util.Dbg {
		util.Debug("\n   %-25s %-25s %-25s %s\n", "id", "short", "json", "sql")
		for i, v := range values {
			util.Debug("%-2d %-25s %-25s %-25s %s\n", i, v.Identifier, v.Shortened, v.JSON, v.SQL)
		}
	}
}
