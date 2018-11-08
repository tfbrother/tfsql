// Go TfSQL Driver - A TfSQL-Driver for Go's database/sql package
package tfsql

import (
	"database/sql/driver"
)

type tfsqlConn struct {
}

// 这个query就是发送的sql语句，比如：select exchange, code from stock where status = ?
func (t *tfsqlConn) Prepare(query string) (driver.Stmt, error) {
	//fmt.Println(query)
	return &tfsqlStmt{}, nil
}

func (t *tfsqlConn) Close() error {
	return nil
}

func (t *tfsqlConn) Begin() (driver.Tx, error) {
	return nil, nil
}
