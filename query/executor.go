package query

import (
	"fmt"
)

type execFunc func([][]string) ([][]string, error)

type Executor interface {
	Exec(data [][]string) ([][]string, error)
}

type executor struct {
	query   *Query
	execute execFunc
}

func NewExecutor(targetIndex int, q *Query) Executor {
	return &executor{
		query:   q,
		execute: makeExecuteFunc(targetIndex, q),
	}
}

func (e *executor) Exec(data [][]string) ([][]string, error) {
	return e.execute(data)
}

func makeExecuteFunc(targetIndex int, q *Query) execFunc {
	switch q.Keyword {
	case "where":
		return makeWhereFunc(targetIndex, q)
	}
	return func([][]string) ([][]string, error) {
		return nil, fmt.Errorf("execute error: invalid keyword %s", q.Keyword)
	}
}
