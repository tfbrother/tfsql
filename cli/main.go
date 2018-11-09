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

	stmt, err := db.Prepare("delete from agent_statis where id=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	if result, err := stmt.Exec(1000); err == nil {
		if affected, err := result.RowsAffected(); err == nil {
			if insertID, err := result.LastInsertId(); err == nil {
				fmt.Println("RowsAffected : ", affected, "LastInsertId : ", insertID)
			} else {
				fmt.Println("RowsAffected : ", affected)
			}

		}
	}
	db.Close()
}
