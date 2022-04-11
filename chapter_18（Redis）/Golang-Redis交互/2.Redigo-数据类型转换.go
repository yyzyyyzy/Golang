package main

import (
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleRedisErr(err, "redis.Dial")
	err = conn.Send("auth", "root")
	HandleRedisErr(err, "conn.send")
	defer conn.Close()

	//reply1, err := conn.Do("SET", "name", "LZK")
	//reply2, err := conn.Do("SETEX", "name", "10", "LZK")
	//reply3, err := conn.Do("PERSIST", "name")
	//reply4, err := conn.Do("MSET", "age", "60", "gender", "男")
	//reply5, err := conn.Do("HMSET", "hashtable", "name", "LZK", "age", 18)	//返回值需要redis.String转化
	//reply6, err := conn.Do("HGETALL", "hashtable")							//返回值需要redis.Strings转化
	//reply7, err := conn.Do("RPUSH", "mylist", 11, 22, 33)
	//reply8, err := conn.Do("LRANGE", "mylist", 0, -1) 						//返回值需要redis.Strings转化
	//reply9, err := conn.Do("SADD", "myset", 1, 2, 3, 4)
	//reply10, err := conn.Do("SMEMBERS", "myset")
	//reply11, err := conn.Do("ZADD", "myzset", 10, "LZK", 9, "WHC")			//返回值需要redis.Int转化
	//reply12, err := conn.Do("ZRANGE", "myzset", 0, -1)						//返回值需要redis.Strings转化

	//fmt.Println("reply原始数据类型：", reply, err)

	//具体转换的数据类型需要根据业务实际需求转化
	//ret, _ := redis.Strings(reply, err)
	//ret, _ := redis.String(reply, err)
	//ret, _ := redis.Int(reply, err)
	//ret, _ := redis.Bool(reply, err)
	//fmt.Println(ret)
}
