package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// cache loaded csv data
var LoadedRows *Rows

// template execution target rows
var ExecutionRows *Rows

func Load(path string, withHeader bool) (*Rows, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read rows in file: %w", err)
	}

	// Error if no rows, no columns or header only.
	if len(lines) == 0 || len(lines[0]) == 0 || (withHeader && len(lines) == 1) {
		return nil, fmt.Errorf("the file rows invalid: %s", path)
	}

	header := lines[0]
	body := lines

	if withHeader {
		// ignore header
		body = body[1:]
	} else {
		for i := range header {
			header[i] = strconv.Itoa(i)
		}
	}

	indexes := make(map[string]int, len(header))
	for i, v := range header {
		indexes[v] = i
	}

	rows := &Rows{
		Count:   len(lines),
		Indexes: indexes,
		Header:  header,
		Body:    body,
	}

	LoadedRows = rows
	ExecutionRows = rows

	return rows, nil
}
