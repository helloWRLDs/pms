package htmlmodule

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/*.tmpl.html
var templates embed.FS

type Template[T any] struct {
	Title   string
	Name    string
	Content T
}

func Render[T any](tmplate Template[T]) ([]byte, error) {
	t, err := template.ParseFS(
		templates,
		"templates/base.tmpl.html",
		fmt.Sprintf("templates/%s.tmpl.html", tmplate.Name),
	)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "base", tmplate); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
