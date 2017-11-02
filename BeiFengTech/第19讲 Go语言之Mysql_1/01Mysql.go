package main

import (
	"database/sql"
	//仅用于调用指定包的init方法
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8");
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO T_USER (NAME,add_time) values(?,?)");
	res,err := stmt.Exec("Goular22","2017-11-02 23:10:09")
	id,err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

}
