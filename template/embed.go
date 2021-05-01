package template

import (
	"embed"
	"html/template"
)

//go:embed index.tmpl
var temp embed.FS

func New() (*template.Template, error) {
	return template.ParseFS(temp, "index.tmpl")
}
