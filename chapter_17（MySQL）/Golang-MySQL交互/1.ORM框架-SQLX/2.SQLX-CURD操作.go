package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //_ 加载数据库驱动
	"github.com/jmoiron/sqlx"
	"os"
)

type Person struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Age    int    `db:"age"`
	Gender string `db:"gender"`
	Email  string `db:"email"`
	Phone  string `db:"phone"`
}

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/sqlx_database")
	HandleSQLXError(err, "25.查询三表，排名前三.Open")
	defer db.Close()

	////增
	//_, err = db.Exec("insert into person(Name,Age,Gender,Email,Phone) values (?,?,?,?,?);", "戴超", 20, "男", "857644244@qq.com", "15195920313")
	//HandleSQLXError(err, "db.Exec insert")
	//
	////删
	//_, err = db.Exec("delete from person where name not like ?;", "%康")
	//HandleSQLXError(err, "db.Exec delete")
	//
	////改
	//_, err = db.Exec("update person set Age =? where id =?", 20, 1)
	//HandleSQLXError(err, "db.Exec update")
	//
	////查一条记录（不能用select * from）
	//var ps []Person
	//_, err = db.Exec("select id, name, age from person where id > ?", 0)
	//HandleSQLXError(err, "db.Get")
	//fmt.Printf("person:%#v\n", ps)
	//
	////查多条记录
	//err = db.Select(&ps, "select id,name,age,gender,email,phone from person;")
	//HandleSQLXError(err, "db.select")
	//fmt.Printf("person:%#v\n", ps)

	var ps Person
	db.Get(&ps, "select id,name,age,gender,email,phone from sqlx_database.person where ID = ?;", 1)
	fmt.Printf("person:%#v\n", ps)
}

func HandleSQLXError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
