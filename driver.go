// Go TfSQL Driver - A TfSQL-Driver for Go's database/sql package
package tfsql

import (
	"database/sql"
	"database/sql/driver"
)

type tfsqlDriver struct{}

// 打开一个新连接
func (d *tfsqlDriver) Open(dsn string) (driver.Conn, error) {
	var err error

	return nil, err
}

func init() {
	sql.Register("tfsql", &tfsqlDriver{})
}
