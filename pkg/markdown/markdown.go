package markdown

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/numtide/godoc/pkg/parse"
)

//go:embed templates
var templatesFS embed.FS

func Write(dir string, pkg *parse.Package) error {
	// load template
	tmpl, err := template.New("struct.tmpl").Funcs(funcMap).ParseFS(templatesFS, "**/*.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// iterate the types in the package
	for _, t := range pkg.Types {

		// ensure the directory exists, using the package name as the leaf directory
		baseDir := filepath.Join(dir, pkg.Name)
		if err := os.MkdirAll(baseDir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", baseDir, err)
		}

		// output markdown
		if err := writeType(filepath.Join(baseDir, strcase.ToSnake(t.Name)+".md"), tmpl, t); err != nil {
			return fmt.Errorf("failed to write doc for %s: %w", t.Name, err)
		}
	}
	return nil
}

func writeType(filename string, tmpl *template.Template, t parse.Struct) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()
	return tmpl.Execute(file, t)
}
