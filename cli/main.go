package main

import (
	"database/sql"
	"fmt"
	_ "github.com/tfbrother/tfsql"
)

func main() {
	db, err := sql.Open("tfsql", "root:password@tcp(ip:port)/test?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(db)
	rows, err := db.Query("select exchange, code from stock where status = ?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows.Columns())
	defer rows.Close()
	db.Close()
}
