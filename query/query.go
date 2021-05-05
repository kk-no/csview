package query

import (
	"fmt"
	"strings"
)

type Query struct {
	Keyword   string
	Target    string
	Statement string
	Value     string
}

func Parse(q string) (*Query, error) {
	qs := strings.Split(q, " ")
	if len(qs) < 4 {
		return nil, fmt.Errorf("invalid query: %s", qs)
	}

	e := &Query{}
	keyword := strings.ToLower(qs[0])
	if !isKeyword(keyword) {
		return nil, fmt.Errorf("invalid query keyword: %s", qs[0])
	}
	e.Keyword = keyword

	target := make([]string, 0, len(qs))
	for i, s := range qs[1:] {
		if isStmt(s) {
			e.Statement = s
			e.Value = strings.Join(qs[i+2:], " ") // Add the index in range.
			break
		}
		target = append(target, s)
	}
	e.Target = strings.Join(target, " ")

	return e, nil
}

func isKeyword(s string) bool {
	s = strings.ToLower(s)
	return s == "where" // TODO: add other keyword
}

func isStmt(s string) bool {
	return s == "=" || s == "<" || s == ">" || s == "<=" || s == ">="
}
