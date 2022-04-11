package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //_ 加载数据库驱动
	"github.com/jmoiron/sqlx"
	"os"
)

func HandleSQLXERR(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/sqlx_database")
	HandleSQLXERR(err, "sqlx.open")
	defer db.Close()

	result, err := db.Exec("insert into person(Name,Age,Gender,Email,Phone) values (?,?,?,?,?);", "李子康", 18, "男", "916990143@qq.com", "13610850940")
	HandleSQLXERR(err, "db.Exec")

	rowAffected, _ := result.RowsAffected()
	lastInsertID, _ := result.LastInsertId()
	fmt.Println("受影响的行数=", rowAffected)   //受影响的行数
	fmt.Println("最后一行的ID=", lastInsertID) //最后一行的ID
}
