package template

import (
	"embed"
	"html/template"
)

//go:embed index.html
var temp embed.FS

func New() (*template.Template, error) {
	return template.ParseFS(temp, "index.html")
}
