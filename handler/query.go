package handler

import (
	"log"
	"net/http"

	"github.com/kk-no/csview/csv"
	"github.com/kk-no/csview/executor"
	"github.com/kk-no/csview/query"
)

type queryHandler struct {
	executor executor.Executor
}

func NewQueryHandler(executor executor.Executor) http.Handler {
	return &queryHandler{executor: executor}
}

func (h queryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	if qs == nil {
		log.Print("no query specified\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q, err := query.Parse(qs.Get("query"))
	if err != nil {
		log.Printf("failed to parse query: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rowIndex, ok := csv.LoadedRows.Indexes[q.Target]
	if !ok {
		log.Printf("colum does not exist: %s\n", q.Target)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := query.NewExecutor(rowIndex, q).Exec(csv.LoadedRows.Body)
	if err != nil {
		log.Printf("execute query error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	execRows := &csv.Rows{
		Count:   1,
		Indexes: csv.LoadedRows.Indexes,
		Header:  csv.LoadedRows.Header,
		Body:    result,
	}

	if err := h.executor.Exec(w, execRows); err != nil {
		log.Printf("execute template error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
