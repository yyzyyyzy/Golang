package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //执行驱动包的init方法，不执行任何API，用_导入
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"os"
)

/*
程序运行后，显示“请输入命令”，输入getall，查询并显示所有人员信息
第一次查询的结果缓存在redis，设置60秒的过期时间
以后每次查询，redis有数据就从redis加载，没有就重复以上步骤
*/

func main() {
	//定义用户输入
	var cmd string

	/*
		循环接收用户输入，用户输入getall：读取mysql数据存入redis；
		                      exit：退出死循环，结束程序
	*/
	for {
		fmt.Println("请输入命令：")

		//阻塞等待用户输入，存入cmd的内存地址
		fmt.Scan(&cmd)

		if cmd == "getall" {
			GetAllPeople()
		} else if cmd == "exit" {
			break
		} else {
			fmt.Println("请重新输入命令(1.getall 2.exit)")
		}
	}
	fmt.Println("OFF")
}

/*
结构体用来存储mysql读出的数据
*/

type Person struct {
	ID     int            `db:"id"`
	Name   sql.NullString `db:"name"`
	Age    sql.NullInt64  `db:"age"`
	Gender sql.NullString `db:"gender"`
	Email  sql.NullString `db:"email"`
	Phone  sql.NullString `db:"phone"`
}

/*
用来处理错误的函数
参数：传入错误、出错的语句
只要出错就打印错误和暴力退出
*/

func HandleERR(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

/*
从mysql中获取所有人员信息
1.判断redis内是否为空，如果是空，可以从mysql读取信息存入redis
2.不为空则不再存储
*/

func GetAllPeople() {
	gotRedisInfo := GetInfoFromRedis() //从Redis取出数据

	if gotRedisInfo {
		fmt.Println("Redis 已有数据,无法继续存储")
	} else {
		people := GetInfoFromMySQL() //从mysql查询数据
		CachePeopletoRedis(&people)  //将mysql数据缓存到redis
	}

}

/*
查询redis people(list)是否为空，返回布尔值
*/

func GetInfoFromRedis() (ok bool) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	HandleERR(err, "redis.Dial")
	err = conn.Send("auth", "root")
	HandleERR(err, "conn.send")
	defer conn.Close()

	reply, err := conn.Do("lrange", "people", 0, -1)
	HandleERR(err, "lrange people all")
	peopleStrs, err := redis.Strings(reply, err)
	fmt.Println("读取缓存的数据：", peopleStrs)

	// 返回redis.Strings(conn.Do("lrange","people",0,-1)的len，len>0为true)
	return len(peopleStrs) > 0
}

/*
使用切片容器存储mysql查询的信息，返回切片
*/

func GetInfoFromMySQL() []Person {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/sqlx_database")
	HandleERR(err, "sqlx.connect")
	defer db.Close()

	var people []Person
	err = db.Select(&people, "select id,name,age,gender,email,phone from person;")
	HandleERR(err, "db.select")
	fmt.Printf("%#v\n", people)
	return people
}

/*
将mysql读取的数据缓存redis内作为二级缓存
*/

func CachePeopletoRedis(people *[]Person) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	HandleERR(err, "redis.Dial")
	err = conn.Send("auth", "root")
	HandleERR(err, "conn.send")
	defer conn.Close()

	// 删除之前缓存的people(list)信息
	conn.Do("del", "people")

	//循环读取切片内的结构体数据，转换为字符串格式后，从尾部插入people(list)队列
	for _, human := range *people {
		humanStr := fmt.Sprint(human)
		_, err := conn.Do("rpush", "people", humanStr)
		HandleERR(err, "rpush people err")
	}

	//设置people的过期时间
	_, err = conn.Do("expire", "people", 60)
	HandleERR(err, "expire people err")
	fmt.Println("Redis 缓存成功")
}
