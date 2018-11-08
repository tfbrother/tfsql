package tfsql

import (
	"database/sql/driver"
)

type tfsqlStmt struct {
}

func (stmt *tfsqlStmt) Close() (err error) {
	return nil
}

// 返回参数个数
func (stmt *tfsqlStmt) NumInput() int {
	return 1
}

func (stmt *tfsqlStmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, nil
}

func (stmt *tfsqlStmt) Query(args []driver.Value) (driver.Rows, error) {
	//fmt.Println(args)
	return &tfsqlRows{1, 0}, nil
}
