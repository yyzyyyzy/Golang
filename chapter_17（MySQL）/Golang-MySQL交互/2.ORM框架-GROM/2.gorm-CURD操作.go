package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int `gorm:"primary key"`
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
	Gender   string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Address  string `gorm:"index:addr"`
}

func main() {
	//打开数据库
	DNS := "root:root@tcp(127.0.0.1:3306)/china?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	//创建表自动迁移（结构体和数据表进行对应）
	db.AutoMigrate(&User{})

	//创建表记录
	user1 := User{
		Id:       0,
		Name:     "LZK",
		Age:      sql.NullInt64{18, true},
		Birthday: &time.Time{},
		Gender:   "男",
		Email:    "916990143@qq.com",
		Address:  "江苏省南京市",
	}
	db.Create(&user1)

	//查询表记录
	var u User
	db.First(&u)
	fmt.Printf("u:%#v\n", u)
}
