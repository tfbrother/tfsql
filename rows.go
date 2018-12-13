package tfsql

import (
	"database/sql/driver"
	"io"
)

// TODO 实现的是driver.Rows接口迭代器，这种迭代器思想在后面的数据结构中遍历时用处非常大
type tfsqlRows struct {
	totalNum int           //保存当前当前结果级中的行总数，也是用于Next方法使用
	nowPtr   int           //保存当前遍历到了结果级中的那一行数据，主要是用于Next方法使用
	rowDatas *[2][2]string //保存了结果级中所有的数据，此处定义为二维数组
}

// 返回结果级中所有的字段名
func (rows *tfsqlRows) Columns() (columns []string) {
	return []string{"name", "age"}
}

// 释放结果级资源
func (rows *tfsqlRows) Close() (err error) {
	rows.rowDatas = nil
	return nil
}

// 实际上是对dest进行处理，因为是指针传递的。并且修改偏移量，下次就读下一条数据
func (rows *tfsqlRows) Next(dest []driver.Value) error {
	if rows.nowPtr >= rows.totalNum {
		return io.EOF
	}
	dest[0] = rows.rowDatas[rows.nowPtr][0]
	dest[1] = rows.rowDatas[rows.nowPtr][1]

	rows.nowPtr++
	return nil
}
