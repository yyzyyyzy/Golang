package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1-33个红球取6个随机球, 蓝球1-16个球随机选择一个， 红球和蓝球不能重复
func main() {
	rand.Seed(time.Now().UnixNano())

	var redball [6]int

	for i := 0; i < len(redball); i++ {
		v := rand.Intn(33) + 1 // 有可能初始值为0，所以需要加1

		for j := 0; j < i; j++ { //
			if v == redball[j] { //如果取的6个随机红球号码重复，需要重新随机取
				v = rand.Intn(33) + 1 //rand.Intn()随即生成0-32的数字
				j = -1
			}
		}
		redball[i] = v
	}
	fmt.Println("红球", redball, "蓝球", rand.Intn(16)+1)
}
