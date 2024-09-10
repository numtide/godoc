package parse

import "go/ast"

type Package struct {
	Name  string
	Types []Struct
}

type Struct struct {
	Name   string
	Doc    string
	Fields []Field
	Source string
}

type Field struct {
	Name string
	Type ast.Expr
	Doc  string
	Tags TagMap
}
