package handler

import (
	"log"
	"net/http"

	"github.com/kk-no/csview/csv"
	"github.com/kk-no/csview/executor"
)

type templateHandler struct {
	executor executor.Executor
}

func NewTemplateHandler(executor executor.Executor) http.Handler {
	return &templateHandler{executor: executor}
}

func (h templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.executor.Exec(w, csv.LoadedRows); err != nil {
		log.Printf("template executor error: %s", err)
		return
	}
}
