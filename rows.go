package tfsql

import (
	"database/sql/driver"
	"io"
)

type tfsqlRows struct {
	num    int
	nowPtr int
}

func (rows *tfsqlRows) Columns() (columns []string) {
	return []string{"exchange", "code"}
}

func (rows *tfsqlRows) Close() (err error) {
	return nil
}

// 实际上是对dest进行处理，因为是指针传递的。并且修改偏移量，下次就读下一条数据
func (rows *tfsqlRows) Next(dest []driver.Value) error {
	if rows.nowPtr >= rows.num {
		return io.EOF
	}
	//fmt.Println(dest)
	dest[0] = "aaaaa"
	dest[1] = 11111
	rows.nowPtr++
	return nil
	//return io.EOF
}
