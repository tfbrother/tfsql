package tfsql

import (
	"database/sql/driver"
)

type tfsqlRows struct {
}

func (rows *tfsqlRows) Columns() (columns []string) {
	return []string{"111", "2222", "3333"}
}

func (rows *tfsqlRows) Close() (err error) {
	return nil
}

func (rows *tfsqlRows) Next(dest []driver.Value) error {
	return nil
}
