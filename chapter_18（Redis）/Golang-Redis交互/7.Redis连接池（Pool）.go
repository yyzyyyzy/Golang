package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func main() {
	pool := &redis.Pool{
		// 最大闲置连接数
		MaxIdle: 20,

		//最大活动连接数，0=无限
		MaxActive: 80,

		//闲置连接的超时时间，就剔除
		IdleTimeout: time.Second * 100,

		//定义拨号连接的函数
		Dial: DialFunc,
	}
	//延时关闭线程池
	defer pool.Close()

	//IO并发连接
	for i := 0; i < 10; i++ {
		go getconnFromPool(pool, i)
	}

	//保持主协程存活
	time.Sleep(time.Second * 3)
}

func getconnFromPool(pool *redis.Pool, i int) {
	// 通过连接池获取连接
	conn := pool.Get()

	//延迟关闭连接
	defer conn.Close()

	//使用连接操作数据库
	reply, err := conn.Do("SET", "conn"+strconv.Itoa(i), i)
	result, _ := redis.String(reply, err)
	fmt.Println(result)
}

func DialFunc() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, err
	}
	//此处root对应redis密码
	if _, err := conn.Do("AUTH", "root"); err != nil {
		conn.Close()
		return nil, err
	}
	return conn, err
}
