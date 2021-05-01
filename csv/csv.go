package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

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

	return &Rows{
		Header: header,
		Body:   body,
	}, nil
}
