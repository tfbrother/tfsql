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

	//fmt.Println(db)
	rows, err := db.Query("select name, age from stock where status = ?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		name string
		age  int
	)

	for rows.Next() {
		err := rows.Scan(&name, &age)
		if err != nil {
			//fmt.Println(err)
			break
		}
		fmt.Println(name, age)
	}
	//fmt.Println(rows.Columns())
	defer rows.Close()
	db.Close()
}
