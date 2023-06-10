package main

import "fmt"

// 递归算法的特点：
// 1.递归就是在过程或函数里面调用自身。
// 2.在使用递归策略时，必须有一个明确的递归结束条件，称为递归出口。

// 递归算法的优缺点：
// 优点：代码的表达力很强，写起来简洁。
// 缺点：空间复杂度高、有堆栈溢出的风险、存在重复计算、过多的函数调用会耗时较多等。

// 递归算法的应用场景：
// 1.数据定义是按递归定义的（Fibonacci数列、汉诺塔问题等）。
// 2.问题解法按递归实现（分治法）。
// 3.数据结构形式是按递归定义的（如链表）。

// 递归算法的实现要点：
// 1.把大问题转化为小问题的解决方案。
// 2.找到递归公式。
// 3.找到终止条件。

func func1(x int) {
	print(x)
	func1(x - 1) // 递归调用且有终止条件
}

func func2(x int) {
	if x > 0 {
		print(x)
		func2(x + 1) // 递归调用但无终止条件
	}
}

func func3(x int) {
	if x > 0 {
		print(x)
		func3(x - 1) // 递归调用且有终止条件
	}
}

// 汉诺塔问题
// 1.把A上的n-1个盘子借助C移动到B上
// 2.把A上的第n个盘子移动到C上
// 3.把B上的n-1个盘子借助A移动到C上
func hanoi(n int, A, B, C string) {
	if n == 1 {
		fmt.Println(A, "->", C)
	} else {
		hanoi(n-1, A, C, B)
		fmt.Println(A, "->", C)
		hanoi(n-1, B, A, C)
	}
}

// 斐波那契数列
// 1.第n个数等于前两个数之和
// 2.第1个数和第2个数都为1
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
