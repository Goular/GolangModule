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
	stmt,err := db.Prepare("UPDATE t_user set name=? where id =?")
	res,err := stmt.Exec("jack",1)
	id,err := res.RowsAffected()
	if err!=nil{
		panic(err)
	}
	fmt.Println(id)
}
