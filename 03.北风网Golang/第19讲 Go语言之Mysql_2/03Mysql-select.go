package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)


func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	rows,err := db.Query("SELECT * from t_user")
	if err!=nil{
		panic(err)
	}
	for rows.Next(){
		var id int
		var name string
		var time string
		err := rows.Scan(&id,&name,&time)
		if err != nil {
			panic(err)
		}
		fmt.Println(id,name,time)
	}
}
