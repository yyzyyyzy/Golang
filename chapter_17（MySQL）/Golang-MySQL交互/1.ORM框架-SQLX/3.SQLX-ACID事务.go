package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, _ := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/sqlx_database")
	defer db.Close()

	// 开启事务
	affairs, _ := db.Begin()

	// 事务的CURD
	result1, err1 := affairs.Exec("insert into person(Name,Age,Gender,Email,Phone) values (?,?,?,?,?);", "王浩川", 18, "男", "wanghaochuan@163.com", "13804056806")
	result2, err2 := affairs.Exec("insert into person(Name,Age,Gender,Email,Phone) values (?,?,?,?,?);", "赵杰", 17, "女", "zhaojie@yahoo.com", "13951972662")
	result3, err3 := affairs.Exec("insert into person(Name,Age,Gender,Email,Phone) values (?,?,?,?,?);", "张坤", 20, "男", "zhangkun@gmail.com", "13951901007")

	// 有一个错误就回滚
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("err1/err2/err3/:", err1, err2, err3)
		// 回滚事务
		affairs.Rollback()
	} else {
		//提交事务
		affairs.Commit()

		ra1, _ := result1.RowsAffected()
		ra2, _ := result2.RowsAffected()
		ra3, _ := result3.RowsAffected()
		fmt.Println("受到事务影响的行数为：", ra1+ra2+ra3)
	}

}
