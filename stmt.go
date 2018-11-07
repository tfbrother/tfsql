package tfsql

import (
	"database/sql/driver"
)

type tfsqlStmt struct {
}

func (stmt *tfsqlStmt) Close() (err error) {
	return nil
}

func (stmt *tfsqlStmt) NumInput() int {
	return 1
}

func (stmt *tfsqlStmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func (stmt *tfsqlStmt) Query(args []driver.Value) (driver.Rows, error) {
	return &tfsqlRows{}, nil
}
