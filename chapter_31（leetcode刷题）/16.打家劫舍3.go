package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//后序遍历+动态规划
func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//递归
//1.每个结点都有两种状态：取或者不取，我们设置2个变量来存储这两种情况的值
//2.如果取了当前结点，那么左右子树结点都不能取
//3.如果没取当前结点，那么左右子树都可以取或者不取，使用max函数来取最大值

func TreasureHunt(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil { //如果当前结点为空，则返回0，这也是递归的边界！！！
			return 0, 0
		} else {
			isleft, notleft := dfs(root.Left)                                                         //左子树取或不取
			isright, notright := dfs(root.Right)                                                      //右子树取或不取
			isroot, notroot := root.Val+notleft+notright, Max(isleft, notleft)+Max(isright, notright) //取或不取根节点，其子树的状态也会不同
			return isroot, notroot
		}
	}
	isroot, notroot := dfs(root)
	return Max(isroot, notroot) //返回最大值
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
