package csv

type Rows struct {
	Count   int
	Indexes map[string]int
	Header  []string
	Body    [][]string // Body without header
}
