package executor

import (
	"fmt"
	"html/template"
	"io"

	"github.com/kk-no/csview/csv"
)

type csvExecutor struct {
	tmp *template.Template
}

func NewCSVExecutor(tmp *template.Template) Executor {
	return &csvExecutor{tmp: tmp}
}

func (e *csvExecutor) Exec(w io.Writer, data interface{}) error {
	rows, ok := data.(*csv.Rows)
	if !ok {
		return fmt.Errorf("invalid rows format %v", data)
	}
	csv.ExecutionRows = rows
	if err := e.tmp.Execute(w, csv.ExecutionRows); err != nil {
		return fmt.Errorf("failed to template execute: %s", err)
	}
	return nil
}
