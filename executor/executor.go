package executor

import "io"

type Executor interface {
	Exec(w io.Writer, data interface{}) error
}
