package markdown

import (
	"fmt"
	"go/ast"
	"reflect"
	"text/template"
	"unicode"

	"github.com/iancoleman/strcase"
	"github.com/numtide/godoc/pkg/parse"
)

var funcMap = template.FuncMap{
	"typeLink":   typeLink,
	"jsonTag":    jsonTag,
	"formatType": formatType,
}

func typeLink(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		if unicode.IsUpper(rune(t.Name[0])) {
			return fmt.Sprintf("[%s](%s.md)", t.Name, strcase.ToSnake(t.Name))
		} else {
			return t.Name
		}
	case *ast.SelectorExpr:
		// For qualified types like time.Time
		return fmt.Sprintf("[%s](../%s/%s.md)", formatType(expr), strcase.ToSnake(formatType(t.X)), strcase.ToSnake(t.Sel.Name))
	case *ast.StarExpr:
		// For pointer types
		return "*" + typeLink(t.X)
	case *ast.ArrayType:
		// For array types
		return "[]" + typeLink(t.Elt)
	case *ast.MapType:
		// For map types
		return fmt.Sprintf("map[%s]%s", typeLink(t.Key), typeLink(t.Value))
	case *ast.FuncType:
		return "func"
	default:
		return reflect.TypeOf(expr).String()
	}
}

// formatType extracts and formats the type of a field
func formatType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		// For qualified types like time.Time
		return fmt.Sprintf("%s.%s", formatType(t.X), t.Sel.Name)
	case *ast.StarExpr:
		// For pointer types
		return "*" + formatType(t.X)
	case *ast.ArrayType:
		// For array types
		return "[]" + formatType(t.Elt)
	case *ast.MapType:
		// For map types
		return fmt.Sprintf("map[%s]%s", formatType(t.Key), formatType(t.Value))
	case *ast.FuncType:
		return "func"
	default:
		return reflect.TypeOf(expr).String()
	}
}

func jsonTag(tags parse.TagMap) string {
	return tags[parse.TagJson]
}
