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
	var result = &tfsqlResult{1, 1000}
	return result, nil
}

// 执行真正的查询操作，然后返回结果级
func (stmt *tfsqlStmt) Query(args []driver.Value) (driver.Rows, error) {
	var rowDatas [2][2]string
	rowDatas[0][0] = "天府兄弟"
	rowDatas[0][1] = "22"
	rowDatas[1][0] = "天府少年"
	rowDatas[1][1] = "12"
	//fmt.Println(args)
	return &tfsqlRows{2, 0, &rowDatas}, nil
}
