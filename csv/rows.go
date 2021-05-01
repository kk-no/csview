package csv

type Rows struct {
	Header []string
	Body   [][]string // Body without header
}
