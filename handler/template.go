package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/kk-no/csview/csv"
)

type templateHandler struct {
	tmp  *template.Template
	rows *csv.Rows
}

func NewTemplateHandler(tmp *template.Template, rows *csv.Rows) http.Handler {
	return &templateHandler{tmp: tmp, rows: rows}
}

func (h templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.tmp.Execute(w, h.rows); err != nil {
		log.Fatalf("failed to template execute: %s", err)
	}
}
