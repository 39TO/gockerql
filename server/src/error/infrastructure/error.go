package infrastructure

import "errors"

var (
	StatementError  = errors.New("statement error")
	ExecError       = errors.New("exec error")
	QueryrowError   = errors.New("queryrow error")
	RowsScanError   = errors.New("Rows Scan error")
	RowsLoopError   = errors.New("Rows Loop error")
	LastInsertError = errors.New("last insert id get error")
)
