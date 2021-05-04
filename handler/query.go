package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/kk-no/csview/csv"
	"github.com/kk-no/csview/executor"
)

type queryHandler struct {
	executor executor.Executor
}

func NewQueryHandler(executor executor.Executor) http.Handler {
	return &queryHandler{executor: executor}
}

func (h queryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if q == nil {
		return
	}

	sq := strings.Split(q.Get("query"), " ")
	if len(sq) < 4 {
		return
	}

	targetRow := sq[1]
	conditionalState := sq[2]
	conditionalValue := sq[3]

	rowIndex, ok := csv.LoadedRows.Indexes[targetRow]
	if !ok {
		return
	}
	comparator := makeComparator(conditionalState, conditionalValue)

	execRows := &csv.Rows{
		Count:   1,
		Indexes: csv.LoadedRows.Indexes,
		Header:  csv.LoadedRows.Header,
		Body:    make([][]string, 0, len(csv.LoadedRows.Body)),
	}

	for _, row := range csv.LoadedRows.Body {
		if comparator(row[rowIndex]) {
			execRows.Body = append(execRows.Body, row)
		}
	}

	if err := h.executor.Exec(w, execRows); err != nil {
		log.Printf("query executor error: %s", err)
		return
	}
}

func makeComparator(state, value string) func(string) bool {
	switch state {
	case "=":
		return func(target string) bool {
			if target == value {
				return true
			}
			return false
		}
	case "<":
		return func(target string) bool {
			if target < value {
				return true
			}
			return false
		}
	case ">":
		return func(target string) bool {
			if target > value {
				return true
			}
			return false
		}
	case "<=":
		return func(target string) bool {
			if target <= value {
				return true
			}
			return false
		}
	case ">=":
		return func(target string) bool {
			if target >= value {
				return true
			}
			return false
		}
	}
	return nil
}
