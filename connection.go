// Go TfSQL Driver - A TfSQL-Driver for Go's database/sql package
package tfsql

import (
	"database/sql/driver"
)

type tfsqlConn struct {
}

func (t *tfsqlConn) Prepare(query string) (driver.Stmt, error) {
	return &tfsqlStmt{}, nil
}

func (t *tfsqlConn) Close() error {
	return nil
}

func (t *tfsqlConn) Begin() (driver.Tx, error) {
	return nil, nil
}
