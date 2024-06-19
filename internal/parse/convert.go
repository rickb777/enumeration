package parse

import (
	"fmt"
	"go/token"
	"go/types"
	"io"
	"regexp"
	"strings"

	"github.com/rickb777/enumeration/v3/internal/collection"
	"github.com/rickb777/enumeration/v3/internal/model"
	"github.com/rickb777/enumeration/v3/internal/transform"
	"github.com/rickb777/enumeration/v3/internal/util"
)

var (
	AliasTable string
	MainType   string
)

var fset *token.FileSet

var tagRE = regexp.MustCompile(`[a-z]:"`)

var basicImports = []string{"fmt", "github.com/rickb777/enumeration/v3/enum"}

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
		AliasTable: AliasTable,
		Extra:      make(map[string]interface{}),
		Imports:    collection.NewStringSet(basicImports...),
	}

	s := newFileScanner(input, src)

	var numFound = 0
	var constItems []constItem
	var baseType string
	var baseKind types.BasicKind

	for s.Scan() != token.EOF {
		switch s.Tok {
		case token.TYPE:
			baseType, baseKind, err = parseType(s, MainType, numFound)
			if err != nil {
				return m, fmt.Errorf("%s: %w", s.Position(), err)
			}
			if baseKind != types.Invalid {
				numFound++
				m.BaseType = baseType
				m.BaseKind = baseKind
			}

		case token.CONST:
			constItems = parseConst(s, constItems)
		}
	}

	debugConstItems(constItems)

	m.Values, _ = filterExportedItems(config.MainType, constItems)

	debugValues(m.Values)

	if s.gs.ErrorCount > 0 {
		return model.Model{}, fmt.Errorf("%s: syntax error\n%s", input, strings.Join(s.errs, "\n"))
	}

	if numFound == 0 {
		return model.Model{}, fmt.Errorf("%s: failed to find type %s", input, config.MainType)
	}

	if len(m.Values) == 0 {
		return model.Model{}, fmt.Errorf("%s: failed to find any values for %s", input, config.MainType)
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
