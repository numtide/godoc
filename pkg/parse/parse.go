package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/packages"
)

func File(name string) (*Package, error) {
	// parse the source file
	fset := token.NewFileSet()

	fileNode, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}

	// load and type-check the package using go/packages
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
		Dir:  ".",
		Fset: fset,
	}

	pkgs, err := packages.Load(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}

	// ensure we loaded a package
	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found")
	} else if len(pkgs) > 1 {
		return nil, fmt.Errorf("multiple packages found")
	}

	// extract the package data
	result := Package{
		Name:  fileNode.Name.Name,
		Types: []Struct{},
	}

	for _, decl := range fileNode.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec := spec.(*ast.TypeSpec)
			typeName := typeSpec.Name.Name
			doc := ""
			if genDecl.Doc != nil {
				doc = genDecl.Doc.Text()
			}

			// only handle structs for now
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			typeDoc := Struct{
				Name:   typeName,
				Doc:    doc,
				Fields: []Field{},
			}

			// process the fields of the struct
			for _, field := range structType.Fields.List {
				for _, fieldName := range field.Names {
					// parse field tags
					tags, err := parseTags(field)
					if err != nil {
						return nil, fmt.Errorf("failed to parse tags: %w", err)
					}

					fieldDoc := Field{
						Name: fieldName.Name,
						Type: field.Type,
						Tags: tags,
					}

					if field.Doc != nil {
						fieldDoc.Doc = field.Doc.Text()
					}
					typeDoc.Fields = append(typeDoc.Fields, fieldDoc)
				}
			}

			// append to types list
			result.Types = append(result.Types, typeDoc)
		}
	}

	return &result, nil
}
